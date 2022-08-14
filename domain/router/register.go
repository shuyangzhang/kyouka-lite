package router

import (
	"github.com/lonelyevil/khl"
	"github.com/shuyangzhang/kyouka-light/domain/service"
)

func register() {
	registerCommand(service.Ping, "ping")
}

func registerCommand(service func(*khl.KmarkdownMessageContext, ...string), commandAliases ...string) {
	for _, command := range commandAliases {
		commandRouter[command] = service
	}
}
