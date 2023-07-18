// Copyright 2023 The Bucketeer Authors.
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

package log

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLoggerOK(t *testing.T) {
	for _, level := range Levels {
		logger, err := NewLogger(
			WithLevel(level),
			WithServiceContext("test-service", "1.0.0"),
		)
		des := fmt.Sprintf("level: %s", level)
		assert.Nil(t, err, des)
		assert.NotNil(t, logger, des)
	}
}

func TestNewLoggerFailed(t *testing.T) {
	logger, err := NewLogger(WithLevel("foo"))
	assert.NotNil(t, err)
	assert.Nil(t, logger)
}
