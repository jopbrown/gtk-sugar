package gfmt

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	STRING_NIL               = "NULL"
	STRING_WHITE_SPACE_CHARS = " \t\n\r"
)

// FormatCommand format the cmd and args to raw gtk-sever command
func FormatCommand(cmd string, args ...interface{}) string {
	var sb strings.Builder
	sb.WriteString(cmd)

	for _, arg := range args {
		sb.WriteByte(' ')

		switch v := arg.(type) {
		case nil:
			sb.WriteString(STRING_NIL)
		case string:
			if len(v) == 0 {
				sb.WriteString(STRING_NIL)
			} else {
				if strings.IndexAny(v, STRING_WHITE_SPACE_CHARS) > 0 {
					v = strconv.Quote(v)
				}
				sb.WriteString(v)
			}

		case bool:
			if v {
				sb.WriteString(STRING_BOOL_TRUE)
			} else {
				sb.WriteString(STRING_BOOL_FALSE)
			}
		default:
			sb.WriteString(fmt.Sprintf("%v", v))
		}
	}

	return sb.String()
}
