package log

import (
	"errors"
	"fmt"
	"mango/pkg/util"
	"os"
	"path"
	"runtime"
	"syscall"
	"time"
)

var (
	kernel32    *syscall.LazyDLL  = syscall.NewLazyDLL(`kernel32.dll`)
	proc        *syscall.LazyProc = kernel32.NewProc(`SetConsoleTextAttribute`)
	CloseHandle *syscall.LazyProc = kernel32.NewProc(`CloseHandle`)

	// 给字体颜色对象赋值
	FontColor Color = Color{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
)

type Color struct {
	black       int // 黑色
	blue        int // 蓝色
	green       int // 绿色
	cyan        int // 青色
	red         int // 红色
	purple      int // 紫色
	yellow      int // 黄色
	lightGray   int // 淡灰色（系统默认值）
	gray        int // 灰色
	lightBlue   int // 亮蓝色
	lightGreen  int // 亮绿色
	lightCyan   int // 亮青色
	lightRed    int // 亮红色
	lightPurple int // 亮紫色
	lightYellow int // 亮黄色
	white       int // 白色
}

// 输出有颜色的字体
func ColorPrint(s string, i int) {
	handle, _, _ := proc.Call(uintptr(syscall.Stdout), uintptr(i))
	print(s)
	CloseHandle.Call(handle)
}

type LogInfo struct {
	File      string
	Line      int
	Classname string
	Level     int
	LogStr    string
	TimeNs    int64
}

var (
	logger            *Logger                                    = nil
	screenPrint       int                                        = 1
	chanPrint         chan LogInfo                               = make(chan LogInfo, 100)
	GetMinLevelConfig func(key string, defaultValue int64) int64 = nil
	cb                func(i LogInfo)                            = nil
	tempLogInfo       []LogInfo
)

// levels
const (
	trace      = 0
	debugLevel = 1
	info       = 2
	warning    = 3
	errorLevel = 4
	fatalLevel = 5
)

const (
	printDebugLevel   = "[ debug ] "
	printInfoLevel    = "[ info  ] "
	printWarningLevel = "[warning] "
	printErrorLevel   = "[ error ] "
	printFatalLevel   = "[ fatal ] "
)

type Logger struct {
	level    int
	baseFile *os.File
	pathname string
}

func init() {
	go func() {
		for {
			i := <-chanPrint
			logStr := i.LogStr
			if i.Level >= warning && runtime.GOOS == `windows` {
				if i.Level == warning {
					ColorPrint(logStr, FontColor.yellow)
				} else if i.Level == errorLevel {
					ColorPrint(logStr, FontColor.red)
				} else {
					ColorPrint(logStr, FontColor.lightRed)
				}
				ColorPrint("\n", FontColor.lightGray)
			} else {
				fmt.Println(logStr)
			}
		}
	}()
}

func New(appName string) (*Logger, error) {
	pathname := ""
	curPath, err := util.GetCurrentPath()
	if err == nil {
		pathname = curPath + "log/" + appName + "/"
		_, err := os.Stat(pathname)
		if err != nil && os.IsNotExist(err) {
			err = os.MkdirAll(pathname, os.ModePerm)
			if err != nil {
				pathname = ""
			}
		}
	}
	if pathname == "" {
		return nil, errors.New("文件路径创建失败")
	}

	file, err := createNewLogFile(pathname)
	if err != nil {
		return nil, err
	}

	// new
	l := new(Logger)
	l.baseFile = file
	l.pathname = pathname

	return l, nil
}

func createNewLogFile(pathName string) (*os.File, error) {
	now := time.Now()
	filename := fmt.Sprintf("%d%02d%02d_%02d_%02d_%02d.log",
		now.Year(),
		now.Month(),
		now.Day(),
		now.Hour(),
		now.Minute(),
		now.Second())
	file, err := os.Create(path.Join(pathName, filename))
	if err != nil {
		return nil, err
	}
	return file, nil
}

// It's dangerous to call the method on logging
func (l *Logger) Close() {
	if l.baseFile != nil {
		l.baseFile.Close()
	}

	l.baseFile = nil
}

func (l *Logger) doPrintf(level int, logInfo string) {
	if level < l.level {
		return
	}
	if l.baseFile == nil {
		panic("logger closed")
	}

	l.baseFile.WriteString(logInfo + "\n")

	//60M分割文件 1024*1024*60
	fi, err := l.baseFile.Stat()
	if err == nil && fi.Size() >= 1024*1024*60 {
		file, err := createNewLogFile(l.pathname)
		if err == nil {
			l.baseFile.Close()
			l.baseFile = file
		}
	}

	if level == fatalLevel {
		fatalExit()
	}
}

// It's dangerous to call the method on logging
func Export(l *Logger) {
	if l != nil {
		logger = l
	}
}

func SetCallback(c func(i LogInfo)) {
	if cb != nil && c != nil {
		return
	}
	cb = c
	if cb != nil {
		for _, logInfo := range tempLogInfo {
			c(logInfo)
		}

		//清空缓存
		tempLogInfo = append([]LogInfo{})
	}
}

func SetScreenPrint(print int) {
	screenPrint = print
}

func nowTimeString() string {
	now := time.Now()
	timeStr := fmt.Sprintf("%v-%02d-%02d %02d:%02d:%02d.%09d",
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond())
	return timeStr
}

func printLog(classname, file, format string, line, level int, a ...interface{}) {
	if GetMinLevelConfig != nil && level < int(GetMinLevelConfig("最低日志级别", 0)) {
		return
	}

	//组装格式
	if screenPrint != 0 || level >= errorLevel || cb == nil {
		logStr := fmt.Sprintf(nowTimeString()+GetLogLevelStr(level)+format, a...)
		chanPrint <- LogInfo{
			LogStr: logStr,
			Level:  level,
		}

		//失去连接时写入文件保存
		if cb == nil && logger != nil {
			logStr = logStr + fmt.Sprintf(" << %s, line #%d ", file, line)
			logger.doPrintf(level, logStr)
		}
	}

	//日志变量
	logInfo := LogInfo{
		File:      file,
		Line:      line,
		Classname: classname,
		Level:     level,
		LogStr:    fmt.Sprintf(format, a...),
		TimeNs:    time.Now().UnixNano(),
	}
	if cb != nil {
		cb(logInfo)
	} else {
		tempLogInfo = append(tempLogInfo, logInfo)
	}
}

func Debug(classname, format string, a ...interface{}) {
	_, file, line, _ := runtime.Caller(2)
	printLog(classname, file, format, line, debugLevel, a...)
}

func Info(classname, format string, a ...interface{}) {
	_, file, line, _ := runtime.Caller(2)
	printLog(classname, file, format, line, info, a...)
}

func Warning(classname, format string, a ...interface{}) {
	_, file, line, _ := runtime.Caller(2)
	printLog(classname, file, format, line, warning, a...)
}

func Error(classname, format string, a ...interface{}) {
	_, file, line, _ := runtime.Caller(2)
	printLog(classname, file, format, line, errorLevel, a...)
}

func Fatal(classname, format string, a ...interface{}) {
	_, file, line, _ := runtime.Caller(2)
	printLog(classname, file, format, line, fatalLevel, a...)
	fatalExit()
}

func fatalExit() {
	time.Sleep(time.Second / 2)
	os.Exit(1)
}

func GetLogLevelStr(level int) string {
	logLevel := ""
	if level == debugLevel {
		logLevel = printDebugLevel
	} else if level == info {
		logLevel = printInfoLevel
	} else if level == warning {
		logLevel = printWarningLevel
	} else if level == errorLevel {
		logLevel = printErrorLevel
	}
	return logLevel
}

func Close() {
	if logger != nil {
		logger.Close()
	}
}
