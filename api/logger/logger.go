package logger

import (
	"fmt"
	"os"
	"strings"

	"github.com/withmandala/go-log"
)

var titleCharNum = 67
var titleChar = "-"
var logger = log.New(os.Stderr).WithoutDebug().WithoutTimestamp().WithColor()

func ShowLogGtdfTitle() {
	fmt.Println(strings.Repeat(titleChar, titleCharNum))
	fmt.Println(`
	Getting Things Done Framework CLI by @akrck02																		
	`)
	fmt.Println(strings.Repeat(titleChar, titleCharNum))
}

func Log(msg string) {
	logger.Info(msg)
}

func Error(msg string) {
	logger.Error(msg)
}

func Warn(msg string) {
	logger.Warn(msg)
}

func Jump() {
	Log("")
}
