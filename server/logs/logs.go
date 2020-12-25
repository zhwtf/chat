/******************************************************************************
 *
 *  Description :
 *    Package exposes info, warning and error loggers.
 *
 *****************************************************************************/
package logs

import (
	"io"
	"log"
	"strings"
)

var (
	Info *log.Logger
	Warn *log.Logger
	Err  *log.Logger
)

func parseFlags(logFlags string) int {
	flags := 0
	for _, v := range strings.Split(logFlags, ",") {
		switch {
		case v == "date":
			flags |= log.Ldate
		case v == "time":
			flags |= log.Ltime
		case v == "microseconds":
			flags |= log.Lmicroseconds
		case v == "longfile":
			flags |= log.Llongfile
		case v == "shortfile":
			flags |= log.Lshortfile
		case v == "UTC":
			flags |= log.LUTC
		case v == "msgprefix":
			flags |= log.Lmsgprefix
		case v == "stdFlags":
			flags |= log.LstdFlags
		default:
			log.Fatalln("Invalid log flags string: ", logFlags)
		}
	}
	if flags == 0 {
		flags = log.LstdFlags
	}
	return flags
}

// Initializes info, warning and error loggers given the flags
// and the output.
func Init(output io.Writer, logFlags string) {
	flags := parseFlags(logFlags)
	Info = log.New(output, "I", flags)
	Warn = log.New(output, "W", flags)
	Err = log.New(output, "E", flags)
}
