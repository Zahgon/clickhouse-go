package clickhouse

import (
	_ "embed"
)

func (c *connect) handshake(auth Auth) error { _ = "STUB: not implemented"; return nil }

// set a read deadline - alternative to context.Read operation will fail if no data is received after deadline.

// context level deadlines override any read deadline

//nolint:govet

func (c *connect) sendAddendum() error { _ = "STUB: not implemented"; return nil }

// todo quota key support
