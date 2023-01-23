package router

import (
	"github.com/lonelyevil/khl"
	"github.com/shuyangzhang/kyouka-lite/domain/service"
)

func register() {
	registerCommand(service.Ping, "ping")
	registerCommand(service.Version, "version")
}

func registerCommand(service func(*khl.KmarkdownMessageContext, ...string), commandAliases ...string) {
	for _, command := range commandAliases {
		commandRouter[command] = service
	}
}
