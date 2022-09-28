// Copyright 2022 The Bucketeer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cacher

import (
	"context"
	"time"

	"go.uber.org/zap"
	kingpin "gopkg.in/alecthomas/kingpin.v2"

	cachev3 "github.com/bucketeer-io/bucketeer/pkg/cache/v3"
	"github.com/bucketeer-io/bucketeer/pkg/cli"
	tc "github.com/bucketeer-io/bucketeer/pkg/feature/cacher"
	featureclient "github.com/bucketeer-io/bucketeer/pkg/feature/client"
	"github.com/bucketeer-io/bucketeer/pkg/health"
	"github.com/bucketeer-io/bucketeer/pkg/metrics"
	"github.com/bucketeer-io/bucketeer/pkg/pubsub"
	"github.com/bucketeer-io/bucketeer/pkg/pubsub/puller"
	redisv3 "github.com/bucketeer-io/bucketeer/pkg/redis/v3"
	"github.com/bucketeer-io/bucketeer/pkg/rpc"
	"github.com/bucketeer-io/bucketeer/pkg/rpc/client"
)

const command = "tag-cacher"

type featureCacher struct {
	*kingpin.CmdClause
	port                         *int
	project                      *string
	subscription                 *string
	topic                        *string
	maxMPS                       *int
	numWorkers                   *int
	flushSize                    *int
	flushInterval                *time.Duration
	featureService               *string
	redisServerName              *string
	redisAddr                    *string
	certPath                     *string
	keyPath                      *string
	serviceTokenPath             *string
	pullerNumGoroutines          *int
	pullerMaxOutstandingMessages *int
	pullerMaxOutstandingBytes    *int
}

func RegisterCommand(r cli.CommandRegistry, p cli.ParentCommand) cli.Command {
	cmd := p.Command(command, "Start tag cacher")
	c := &featureCacher{
		CmdClause:        cmd,
		port:             cmd.Flag("port", "Port to bind to.").Default("9090").Int(),
		project:          cmd.Flag("project", "Google Cloud project name.").String(),
		subscription:     cmd.Flag("subscription", "Google PubSub subscription name.").String(),
		topic:            cmd.Flag("topic", "Google PubSub topic name.").String(),
		maxMPS:           cmd.Flag("max-mps", "Maximum messages should be handled in a second.").Default("5000").Int(),
		numWorkers:       cmd.Flag("num-workers", "Number of workers.").Default("2").Int(),
		flushSize:        cmd.Flag("flush-size", "Maximum number of messages in one flush.").Default("100").Int(),
		flushInterval:    cmd.Flag("flush-interval", "Maximum interval between two flushes.").Default("1m").Duration(),
		featureService:   cmd.Flag("feature-service", "bucketeer-feature-service address.").Default("feature:9090").String(),
		redisServerName:  cmd.Flag("redis-server-name", "Name of the redis.").Required().String(),
		redisAddr:        cmd.Flag("redis-addr", "Address of the redis.").Required().String(),
		certPath:         cmd.Flag("cert", "Path to TLS certificate.").Required().String(),
		keyPath:          cmd.Flag("key", "Path to TLS key.").Required().String(),
		serviceTokenPath: cmd.Flag("service-token", "Path to service token.").Required().String(),
		pullerNumGoroutines: cmd.Flag(
			"puller-num-goroutines",
			"Number of goroutines will be spawned to pull messages.",
		).Int(),
		pullerMaxOutstandingMessages: cmd.Flag(
			"puller-max-outstanding-messages",
			"Maximum number of unprocessed messages.",
		).Int(),
		pullerMaxOutstandingBytes: cmd.Flag("puller-max-outstanding-bytes", "Maximum size of unprocessed messages.").Int(),
	}
	r.RegisterCommand(c)
	return c
}

func (c *featureCacher) Run(ctx context.Context, metrics metrics.Metrics, logger *zap.Logger) error {
	registerer := metrics.DefaultRegisterer()

	puller, err := c.createPuller(ctx, logger)
	if err != nil {
		return err
	}

	creds, err := client.NewPerRPCCredentials(*c.serviceTokenPath)
	if err != nil {
		return err
	}

	featureClient, err := featureclient.NewClient(*c.featureService, *c.certPath,
		client.WithPerRPCCredentials(creds),
		client.WithDialTimeout(30*time.Second),
		client.WithBlock(),
		client.WithMetrics(registerer),
		client.WithLogger(logger),
	)
	if err != nil {
		return err
	}
	defer featureClient.Close()

	redisV3Client, err := redisv3.NewClient(
		*c.redisAddr,
		redisv3.WithServerName(*c.redisServerName),
		redisv3.WithMetrics(registerer),
		redisv3.WithLogger(logger),
	)
	if err != nil {
		return err
	}
	defer redisV3Client.Close()
	redisV3Cache := cachev3.NewRedisCache(redisV3Client)

	cacher := tc.NewFeatureCacher(puller, featureClient, redisV3Cache,
		tc.WithMaxMPS(*c.maxMPS),
		tc.WithNumWorkers(*c.numWorkers),
		tc.WithFlushSize(*c.flushSize),
		tc.WithFlushInterval(*c.flushInterval),
		tc.WithMetrics(registerer),
		tc.WithLogger(logger),
	)
	defer cacher.Stop()
	go cacher.Run() // nolint:errcheck

	healthChecker := health.NewGrpcChecker(
		health.WithTimeout(time.Second),
		health.WithCheck("metrics", metrics.Check),
		health.WithCheck("cacher", cacher.Check),
	)
	go healthChecker.Run(ctx)

	server := rpc.NewServer(healthChecker, *c.certPath, *c.keyPath,
		rpc.WithPort(*c.port),
		rpc.WithMetrics(registerer),
		rpc.WithLogger(logger),
		rpc.WithHandler("/health", healthChecker),
	)
	defer server.Stop(10 * time.Second)
	go server.Run()

	<-ctx.Done()
	return nil
}

func (c *featureCacher) createPuller(ctx context.Context, logger *zap.Logger) (puller.Puller, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	client, err := pubsub.NewClient(ctx, *c.project, pubsub.WithLogger(logger))
	if err != nil {
		return nil, err
	}
	return client.CreatePuller(*c.subscription, *c.topic,
		pubsub.WithNumGoroutines(*c.pullerNumGoroutines),
		pubsub.WithMaxOutstandingMessages(*c.pullerMaxOutstandingMessages),
		pubsub.WithMaxOutstandingBytes(*c.pullerMaxOutstandingBytes),
	)
}
