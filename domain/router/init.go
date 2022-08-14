package router

import "github.com/lonelyevil/khl"

var commandRouter = make(map[string]func(*khl.KmarkdownMessageContext, ...string))

func init() {
	register()
}
