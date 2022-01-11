package util

import (
	"os"
	"runtime"
	"strconv"
	"strings"
)

func CurMemory() int64 {
	var rtm runtime.MemStats
	runtime.ReadMemStats(&rtm)
	return int64(rtm.Alloc / 1024)
}

func ParseArgsUint32(name string) (uint32, bool) {
	args := os.Args
	for i := 0; i < len(args); i++ {
		a := strings.Split(args[i], "=")
		if len(a) != 2 {
			continue
		}
		if a[0] == name {
			v, err := strconv.Atoi(a[1])
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
		a := strings.Split(args[i], "=")
		if len(a) != 2 {
			continue
		}
		if a[0] == name {
			return a[1], true
		}
	}
	return "", false
}

func MakeUint64FromUint32(high, low uint32) uint64 {
	return uint64(high)<<32 | uint64(low)
}

func Get2Uint32FromUint64(v uint64) (uint32, uint32) {
	return GetHUint32FromUint64(v), GetLUint32FromUint64(v)
}

func GetHUint32FromUint64(v uint64) uint32 {
	return uint32(v >> 32)
}

func GetLUint32FromUint64(v uint64) uint32 {
	return uint32(v & 0xFFFFFFFF)
}
