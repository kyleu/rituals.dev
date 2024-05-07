// Package websocket - Content managed by Project Forge, see [projectforge.md] for details.
package websocket

import (
	"context"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/lib/telemetry"
	"github.com/kyleu/rituals/app/lib/user"
	dbuser "github.com/kyleu/rituals/app/user"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) Register(u *dbuser.User, profile *user.Profile, accts user.Accounts, c *websocket.Conn, h Handler, logger util.Logger) (*Connection, error) {
	conn := NewConnection("system", u, profile, accts, c, h)
	s.connectionsMu.Lock()
	defer s.connectionsMu.Unlock()
	s.connections[conn.ID] = conn
	if s.onOpen != nil {
		err := s.onOpen(s, conn, logger)
		if err != nil {
			return nil, err
		}
	}
	s.WriteTap(NewMessage(nil, "connection", "open", conn), logger)
	return conn, nil
}

func OnMessage(ctx context.Context, s *Service, connID uuid.UUID, message *Message, logger util.Logger) error {
	ctx, span, logger := telemetry.StartSpan(ctx, "message::"+message.Cmd, logger)
	defer span.Complete()

	if connID == systemID {
		logger.Warnf("admin message received: %s", util.ToJSON(message))
		return nil
	}
	s.connectionsMu.Lock()
	c, ok := s.connections[connID]
	s.connectionsMu.Unlock()
	if !ok {
		return invalidConnection(connID)
	}
	s.WriteTap(message, logger)
	if c.handler == nil {
		return nil
	} else {
		return c.handler(ctx, s, c, message.Channel, message.Cmd, message.Param, logger)
	}
}

func (s *Service) Disconnect(connID uuid.UUID, logger util.Logger) (bool, error) {
	conn, ok := s.connections[connID]
	if !ok {
		return false, errors.Errorf("no connection found with id [%s]", connID.String())
	}
	left := false
	if conn.Channels != nil {
		left = true
		for _, x := range conn.Channels {
			_, err := s.Leave(connID, x, logger)
			if err != nil {
				return left, errors.Wrapf(err, "error leaving channel [%s]", x)
			}
		}
	}
	s.WriteTap(NewMessage(nil, "connection", "close", conn), logger)
	s.connectionsMu.Lock()
	defer s.connectionsMu.Unlock()
	delete(s.connections, connID)
	return left, nil
}

func invalidConnection(id uuid.UUID) error {
	return errors.Errorf("no connection found with id [%s]", id.String())
}
