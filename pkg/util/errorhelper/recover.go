package errorhelper

import (
	"github.com/GoldBaby5511/go-mango-core/log"
	"runtime/debug"
)

func Recover() {
	if err := recover(); err != nil {
		log.Error("", "Recover error:", err, "\r\n", string(debug.Stack()))
	}
}

func RecoverWarn() {
	if err := recover(); err != nil {
		log.Debug("", "Recover Warn:", err)
	}
}
