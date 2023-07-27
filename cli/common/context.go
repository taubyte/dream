package common

import (
	"context"

	client "github.com/taubyte/dreamland/service"
)

type Context struct {
	Ctx        context.Context
	Multiverse *client.Client
}
