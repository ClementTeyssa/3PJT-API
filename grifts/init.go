package grifts

import (
	"github.com/ClementTeyssa/3_pjt_api/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
