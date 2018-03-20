package grifts

import (
	"github.com/bscott/k8s_buffalo_example/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
