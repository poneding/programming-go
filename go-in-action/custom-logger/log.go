package cuzlog

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

var (
	traceLogger *log.Logger
	debuglogger *log.Logger
	infologger  *log.Logger
	warnlogger  *log.Logger
	errorlogger *log.Logger
	fatallogger *log.Logger

	logout   *os.File
	logFile  bool
	filelock sync.RWMutex

	loglevel LogLevel // Default TRACE level

	curDay int
)

type LogLevel int

const (
	TRACE LogLevel = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

func init() {
	logout = os.Stdout
	setLogger(logout)
}

func setLogger(logout *os.File) {
	curDay = time.Now().YearDay()
	fmt.Printf("curDay: %d\n", curDay)

	traceLogger = log.New(logout, "[TRACE] ", log.LstdFlags)
	debuglogger = log.New(logout, "[DEBUG] ", log.LstdFlags)
	infologger = log.New(logout, "[INFO] ", log.LstdFlags)
	warnlogger = log.New(logout, "[WARN] ", log.LstdFlags)
	errorlogger = log.New(logout, "[ERROR] ", log.LstdFlags)
	fatallogger = log.New(logout, "[FATAL] ", log.LstdFlags)
}

func onDayChanged() {
	if logFile {
		filelock.Lock()
		defer filelock.Unlock()
		day := time.Now().YearDay()
		if curDay != day {
			logout.Close()
			// rename log file
			os.Rename(logout.Name(), time.Now().Format("20060102")+"-"+logout.Name())
			var err error
			logout, err = os.OpenFile(logout.Name(), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				log.Fatalln("Open log file error:", err)
			}
			setLogger(logout)

			curDay = day
		}
	}
}

func SetLevel(level LogLevel) {
	loglevel = level
}

func SetFile(file string) {
	logFile = true
	var err error
	logout, err = os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("Open log file error:", err)
	}
	setLogger(logout)
}

func Trace(v ...interface{}) {
	onDayChanged()
	if loglevel <= TRACE {
		traceLogger.Print(getCaller(), v)
	}
}

func Traceln(v ...interface{}) {
	onDayChanged()
	if loglevel <= TRACE {
		traceLogger.Println(getCaller(), v)
	}
}

func Tracef(format string, v ...interface{}) {
	onDayChanged()
	if loglevel <= TRACE {
		traceLogger.Printf(getCaller()+format, v...)
	}
}

func Debug(v ...interface{}) {
	onDayChanged()
	if loglevel <= DEBUG {
		debuglogger.Print(getCaller(), v)
	}
}

func Debugln(v ...interface{}) {
	onDayChanged()
	if loglevel <= DEBUG {
		debuglogger.Println(getCaller(), v)
	}
}

func Debugf(format string, v ...interface{}) {
	onDayChanged()
	if loglevel <= DEBUG {
		debuglogger.Printf(getCaller()+format, v...)
	}
}

func Info(v ...interface{}) {
	onDayChanged()
	if loglevel <= INFO {
		infologger.Print(getCaller(), v)
	}
}

func Infoln(v ...interface{}) {
	onDayChanged()
	if loglevel <= INFO {
		infologger.Println(getCaller(), v)
	}
}

func Infof(format string, v ...interface{}) {
	onDayChanged()
	if loglevel <= INFO {
		infologger.Printf(getCaller()+format, v...)
	}
}

func Warn(v ...interface{}) {
	onDayChanged()
	if loglevel <= WARN {
		warnlogger.Print(getCaller(), v)
	}
}

func Warnln(v ...interface{}) {
	onDayChanged()
	if loglevel <= WARN {
		warnlogger.Println(getCaller(), v)
	}
}

func Warnf(format string, v ...interface{}) {
	onDayChanged()
	if loglevel <= WARN {
		warnlogger.Printf(getCaller()+format, v...)
	}
}

func Error(v ...interface{}) {
	onDayChanged()
	if loglevel <= ERROR {
		errorlogger.Print(getCaller(), v)
	}
}

func Errorln(v ...interface{}) {
	onDayChanged()
	if loglevel <= ERROR {
		errorlogger.Println(getCaller(), v)
	}
}

func Errorf(format string, v ...interface{}) {
	onDayChanged()
	if loglevel <= ERROR {
		errorlogger.Printf(getCaller()+format, v...)
	}
}

func Fatal(v ...interface{}) {
	onDayChanged()
	if loglevel <= FATAL {
		fatallogger.Fatal(getCaller(), v)
	}
}

func Fatalln(v ...interface{}) {
	onDayChanged()
	if loglevel <= FATAL {
		fatallogger.Fatalln(getCaller(), v)
	}
}

func Fatalf(format string, v ...interface{}) {
	onDayChanged()
	if loglevel <= FATAL {
		fatallogger.Fatalf(getCaller()+format, v)
	}
}

// getCaller 返回调用者的文件名和行号
func getCaller() string {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return ""
	}
	fmt.Printf("file: %s\n", file)

	executable, err := os.Executable()
	if err != nil {
		return ""
	}
	fmt.Printf("executable: %s\n", executable)

	lastSlash := strings.LastIndex(executable, "/")

	file = strings.Replace(file, executable[0:lastSlash+1], "", 1)

	return fmt.Sprintf("%s: %d ", file, line)
}
