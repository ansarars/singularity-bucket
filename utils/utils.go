package utils

import (
	"errors"
	"fmt"
	log "github.com/hpe-storage/common-host-libs/logger"
	"strings"
)

const (
	ARG_KEY_INDEX   int = 0
	ARG_VALUE_INDEX int = 1
)

func MakeCommand(args []string) (map[string]string, error) {
	argsMap := map[string]string{}
	for _, arg := range args {
		arg = strings.TrimSpace(arg)
		tokens := strings.SplitN(arg, "=", 2)
		if len(tokens) < 2 {
			msg := fmt.Sprintf("value for key %s is not provided", tokens[ARG_KEY_INDEX])
			log.Errorf(msg)
			return nil, errors.New(msg)
		}
		argsMap[tokens[ARG_KEY_INDEX]] = tokens[ARG_VALUE_INDEX]
	}
	return argsMap, nil
}
