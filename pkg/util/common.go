package util

import (
	"os"
	"runtime"
	"strconv"
)

func CurMemory() int64 {
	var rtm runtime.MemStats
	runtime.ReadMemStats(&rtm)
	return int64(rtm.Alloc / 1024)
}

func ParseArgsUint32(name string) (uint32, bool) {
	args := os.Args
	for i := 0; i < len(args); i++ {
		if args[i] == name && i+1 < len(args) {
			v, err := strconv.Atoi(args[i+1])
			if err == nil {
				return uint32(v), true
			}
		}
	}
	return 0, false
}

func ParseArgsString(name string) (string, bool) {
	args := os.Args
	for i := 0; i < len(args); i++ {
		if args[i] == name && i+1 < len(args) {
			return args[i+1], true
		}
	}
	return "", false
}

func MakeUint64FromUint32(high, low uint32) uint64 {
	return uint64(high)<<32 | uint64(low)
}
