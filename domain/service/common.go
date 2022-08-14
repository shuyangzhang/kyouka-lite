package service

import (
	"github.com/lonelyevil/khl"
	"github.com/shuyangzhang/kyouka-light/tools"
)

func Ping(ctx *khl.KmarkdownMessageContext, parameters ...string) {
	tools.Send(ctx, "コスモブルーフラッシュ！")
}
