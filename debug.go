package debug

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	logFilepath = "/tmp/debug.log"
)

func Print(a ...interface{}) {
	f, err := os.OpenFile(logFilepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	pid := os.Getpid()
	goid := getGoId()
	_, file, line, _ := runtime.Caller(1)
	fmt.Fprintf(f, "[%s %d:%d %s:%d] ", time.Now().Format("2006-01-02 15:04:05.000000"), pid, goid, chopPath(file), line)

	if _, err := fmt.Fprint(f, a...); err != nil {
		panic(err)
	}
}

func Printf(format string, a ...interface{}) {
	f, err := os.OpenFile(logFilepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	pid := os.Getpid()
	goid := getGoId()
	_, file, line, _ := runtime.Caller(1)
	fmt.Fprintf(f, "[%s %d:%d %s:%d] ", time.Now().Format("2006-01-02 15:04:05.000000"), pid, goid, chopPath(file), line)

	if _, err := fmt.Fprintf(f, format, a...); err != nil {
		panic(err)
	}
}

func Println(a ...interface{}) {
	f, err := os.OpenFile(logFilepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	pid := os.Getpid()
	goid := getGoId()
	_, file, line, _ := runtime.Caller(1)
	fmt.Fprintf(f, "[%s %d:%d %s:%d] ", time.Now().Format("2006-01-02 15:04:05.000000"), pid, goid, chopPath(file), line)

	if _, err := fmt.Fprintln(f, a...); err != nil {
		panic(err)
	}
}

func chopPath(original string) string {
	i := strings.LastIndex(original, "src/")
	if i == -1 {
		return original
	} else {
		return original[i+4:]
	}
}

func getGoId() int64 {
	var buf [64]byte
	length := runtime.Stack(buf[:], false)
	stack := buf[:length]
	stack = stack[len("goroutine "):]
	stack = stack[:bytes.IndexByte(stack, ' ')]
	goid, _ := strconv.ParseInt(string(stack), 10, 64)
	return goid
}
