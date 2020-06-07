package sandbox

import (
	"github.com/kyleu/rituals.dev/app/web"
)

var Testbed = Sandbox{
	Key:         "testbed",
	Title:       "Testbed",
	Description: "This could do anything, be careful",
	Resolve: func(ctx *web.RequestContext) (string, interface{}, error) {
		var rsp = "Testbed Response!"
		return "Testbed!", rsp, nil
	},
}
