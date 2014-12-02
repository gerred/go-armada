package command

import (
	"fmt"
	"strings"
)

func hostLog(opts *CommandHandlerOptions, msgs ...string) {
	fmt.Println("["+opts.EnvName+"@"+opts.Endpoint+"]:", strings.Join(msgs, " "))
}

func contains(s []string, p string) bool {
	for _, v := range s {
		if v == p {
			return true
		}
	}
	return false
}
