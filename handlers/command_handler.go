package handlers

import (
	"github.com/lonelyevil/khl"
	"github.com/shuyangzhang/kyouka-lite/domain/router"
	"github.com/shuyangzhang/kyouka-lite/tools"
)

func commandHandler(ctx *khl.KmarkdownMessageContext) {
	if ctx.Common.Type != khl.MessageTypeKMarkdown || ctx.Extra.Author.Bot {
		return
	}

	withPrefix, command, params := tools.ParseCommandWithParameters(ctx.Common.Content)

	if withPrefix {
		router.RouteCommand(ctx, command, params)
	}
}
