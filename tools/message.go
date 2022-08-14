package tools

import "github.com/lonelyevil/khl"

func Send(ctx *khl.KmarkdownMessageContext, content string) {
	ctx.Session.MessageCreate(
		&khl.MessageCreate{
			MessageCreateBase: khl.MessageCreateBase{
				TargetID: ctx.Common.TargetID,
				Content:  content,
				Type:     khl.MessageTypeKMarkdown,
			},
		},
	)
}

func Reply(ctx *khl.KmarkdownMessageContext, content string) {
	ctx.Session.MessageCreate(
		&khl.MessageCreate{
			MessageCreateBase: khl.MessageCreateBase{
				TargetID: ctx.Common.TargetID,
				Content:  content,
				Quote:    ctx.Common.MsgID,
				Type:     khl.MessageTypeKMarkdown,
			},
		},
	)
}
