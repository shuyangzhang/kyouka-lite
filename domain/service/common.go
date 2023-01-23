package service

import (
	"fmt"

	"github.com/lonelyevil/khl"
	"github.com/shuyangzhang/kyouka-lite/domain/constant"
	"github.com/shuyangzhang/kyouka-lite/tools"
)

func Ping(ctx *khl.KmarkdownMessageContext, parameters ...string) {
	tools.Send(ctx, constant.HelloMessage)
}

func Version(ctx *khl.KmarkdownMessageContext, parameters ...string) {
	tools.Reply(ctx, fmt.Sprintf("Version: %v", constant.Version))
}
