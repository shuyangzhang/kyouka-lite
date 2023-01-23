package tools

import (
	"strings"

	"github.com/shuyangzhang/kyouka-lite/configs"
	"golang.org/x/exp/slices"
)

func ParseCommandWithParameters(rawCommand string) (withPrefix bool, command string, params []string) {
	runeRawCommand := []rune(rawCommand)
	prefix := string(runeRawCommand[0])
	commandWithParameters := string(runeRawCommand[1:])

	if slices.Contains(configs.EnvConfigs.AllPrefixes, prefix) {
		commandSlice := strings.Fields(commandWithParameters)

		withPrefix = true
		command = commandSlice[0]
		params = commandSlice[1:]
	}

	return
}
