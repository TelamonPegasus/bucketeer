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

package command

import (
	"context"

	"github.com/golang/protobuf/proto" // nolint:staticcheck

	"github.com/bucketeer-io/bucketeer/pkg/account/domain"
	domainevent "github.com/bucketeer-io/bucketeer/pkg/domainevent/domain"
	"github.com/bucketeer-io/bucketeer/pkg/pubsub/publisher"
	accountproto "github.com/bucketeer-io/bucketeer/proto/account"
	eventproto "github.com/bucketeer-io/bucketeer/proto/event/domain"
)

type accountCommandHandler struct {
	editor               *eventproto.Editor
	account              *domain.Account
	publisher            publisher.Publisher
	environmentNamespace string
}

func NewAccountCommandHandler(
	editor *eventproto.Editor,
	account *domain.Account,
	p publisher.Publisher,
	environmentNamespace string,
) Handler {
	return &accountCommandHandler{
		editor:               editor,
		account:              account,
		publisher:            p,
		environmentNamespace: environmentNamespace,
	}
}

func (h *accountCommandHandler) Handle(ctx context.Context, cmd Command) error {
	switch c := cmd.(type) {
	case *accountproto.CreateAccountCommand:
		return h.create(ctx, c)
	case *accountproto.ChangeAccountRoleCommand:
		return h.changeRole(ctx, c)
	case *accountproto.EnableAccountCommand:
		return h.enable(ctx, c)
	case *accountproto.DisableAccountCommand:
		return h.disable(ctx, c)
	case *accountproto.DeleteAccountCommand:
		return h.delete(ctx, c)
	default:
		return ErrBadCommand
	}
}

func (h *accountCommandHandler) create(ctx context.Context, cmd *accountproto.CreateAccountCommand) error {
	return h.send(ctx, eventproto.Event_ACCOUNT_CREATED, &eventproto.AccountCreatedEvent{
		Id:        h.account.Id,
		Email:     h.account.Email,
		Name:      h.account.Name,
		Role:      h.account.Role,
		Disabled:  h.account.Disabled,
		CreatedAt: h.account.CreatedAt,
		UpdatedAt: h.account.UpdatedAt,
	})
}

func (h *accountCommandHandler) changeRole(ctx context.Context, cmd *accountproto.ChangeAccountRoleCommand) error {
	if err := h.account.ChangeRole(cmd.Role); err != nil {
		return err
	}
	return h.send(ctx, eventproto.Event_ACCOUNT_ROLE_CHANGED, &eventproto.AccountRoleChangedEvent{
		Id:   h.account.Id,
		Role: cmd.Role,
	})
}

func (h *accountCommandHandler) enable(ctx context.Context, cmd *accountproto.EnableAccountCommand) error {
	if err := h.account.Enable(); err != nil {
		return err
	}
	return h.send(ctx, eventproto.Event_ACCOUNT_ENABLED, &eventproto.AccountEnabledEvent{
		Id: h.account.Id,
	})
}

func (h *accountCommandHandler) disable(ctx context.Context, cmd *accountproto.DisableAccountCommand) error {
	if err := h.account.Disable(); err != nil {
		return err
	}
	return h.send(ctx, eventproto.Event_ACCOUNT_DISABLED, &eventproto.AccountDisabledEvent{
		Id: h.account.Id,
	})
}

func (h *accountCommandHandler) delete(ctx context.Context, cmd *accountproto.DeleteAccountCommand) error {
	if err := h.account.Delete(); err != nil {
		return err
	}
	return h.send(ctx, eventproto.Event_ACCOUNT_DELETED, &eventproto.AccountDeletedEvent{
		Id: h.account.Id,
	})
}

func (h *accountCommandHandler) send(ctx context.Context, eventType eventproto.Event_Type, event proto.Message) error {
	e, err := domainevent.NewEvent(
		h.editor,
		eventproto.Event_ACCOUNT,
		h.account.Id,
		eventType,
		event,
		h.environmentNamespace,
	)
	if err != nil {
		return err
	}
	if err := h.publisher.Publish(ctx, e); err != nil {
		return err
	}
	return nil
}
