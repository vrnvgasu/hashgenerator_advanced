package lg

import (
	"github.com/vrnvgasu/logwrapper"
)

var Logger *logwrapper.StandardLogger

func Error(op string, pack string, err error) {
	Logger.Payload(newPayload(op, pack)).Error(err)
}

func Fatal(op string, pack string, err error) {
	Logger.Payload(newPayload(op, pack)).Fatal(err)
}

func Info(op string, pack string, message string) {
	Logger.Payload(newPayload(op, pack)).Info(message)
}

func newPayload(op string, pack string) *logwrapper.Payload {
	return logwrapper.NewPayload().
		Op(op).
		Package(pack)
}
