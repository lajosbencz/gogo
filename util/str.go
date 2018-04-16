package util

import (
	"strings"
	"fmt"
)

type FmtArgs map[string]interface{}

func StrFmt(format string, arguments FmtArgs) string {
	args, i := make([]string, len(arguments)*2), 0
	for k, v := range arguments {
		args[i] = "%{" + k + "}"
		args[i+1] = fmt.Sprint(v)
		i += 2
	}
	return strings.NewReplacer(args...).Replace(format)
}
