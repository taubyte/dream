package common

import (
	"context"

	client "github.com/taubyte/dreamland/http"
)

type Context struct {
	Ctx        context.Context
	Multiverse *client.Client
}
