package router

import (
	"fmt"

	"github.com/lonelyevil/khl"
)

func RouteCommand(ctx *khl.KmarkdownMessageContext, command string, params []string) {
	commandService, ok := commandRouter[command]

	if ok {
		commandService(ctx, params...)
	} else {
		fmt.Printf("Unknown command: %v \n", command)
	}
}
