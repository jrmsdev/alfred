// Copyright (c) Jeremías Casteglione <jrmsdev@gmail.com>
// See LICENSE file.

package log

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type logger struct {
	Debug   func(io.Writer, string)
	Print   func(io.Writer, string)
	Warn    func(io.Writer, string)
	Error   func(io.Writer, string)
	Format  func(string, ...interface{}) string
	Formatf func(string, string, ...interface{}) string
}

var (
	l        *logger
	shortIdx int
	colored  bool
)

var lvlTag = map[string]string{
	"debug": "D: ",
	"error": "E: ",
	"warn":  "W: ",
	"msg":   "",
}

// colors

var (
	//~ cyan = "\033[0;36m"
	//~ blue = "\033[0;34m"
	red    = "\033[0;31m"
	yellow = "\033[0;33m"
	green  = "\033[0;32m"
	grey   = "\033[1;30m"
	reset  = "\033[0m"
)

var lvlColor = map[string]string{
	"debug": grey,
	"error": red,
	"warn":  yellow,
	"msg":   green,
}

func init() {
	colored = false
	if istty(os.Stdout) && istty(os.Stderr) {
		colored = true
	}
}

func istty(fh *os.File) bool {
	if st, err := fh.Stat(); err == nil {
		m := st.Mode()
		if m&os.ModeDevice != 0 && m&os.ModeCharDevice != 0 {
			return true
		}
	}
	return false
}

// init level

func Init(level string) {
	l = new(logger)
	l.Debug = quiet
	l.Print = print
	l.Warn = print
	l.Error = print
	if level == "debug" {
		l.Debug = print
		l.Print = print
		l.Warn = print
		l.Error = print
	} else if level == "warn" {
		l.Debug = quiet
		l.Print = quiet
		l.Warn = print
		l.Error = print
	} else if level == "error" {
		l.Debug = quiet
		l.Print = quiet
		l.Warn = quiet
		l.Error = print
	} else if level == "quiet" {
		l.Debug = quiet
		l.Print = quiet
		l.Warn = quiet
		l.Error = quiet
	}
	shortIdx = getShortIdx()
	if colored {
		l.Format = color
		l.Formatf = colorf
	} else {
		l.Format = mfmt
		l.Formatf = mfmtf
	}
	//~ Debug("level %s", level)
	//~ Debug("stderr %s", os.Stderr.Name())
	//~ if st, err := os.Stderr.Stat(); err == nil {
	//~ Debug("stderr %s", st.Mode())
	//~ } else {
	//~ Error(err)
	//~ }
}

func getShortIdx() int {
	_, fn, _, ok := runtime.Caller(0)
	if ok {
		idx := strings.Index(fn, "src")
		idx += 3 + len(string(filepath.Separator))
		return idx
	}
	return 0
}

// message format/color functions

func mfmt(lvl string, args ...interface{}) string {
	tag := lvlTag[lvl]
	return fmt.Sprintf("%s%s", tag, fmt.Sprint(args...))
}

func mfmtf(lvl string, format string, args ...interface{}) string {
	tag := lvlTag[lvl]
	return fmt.Sprintf("%s%s", tag, fmt.Sprintf(format, args...))
}

func color(lvl string, args ...interface{}) string {
	clr := lvlColor[lvl]
	return fmt.Sprint(clr, fmt.Sprint(args...), reset)
}

func colorf(lvl string, format string, args ...interface{}) string {
	clr := lvlColor[lvl]
	return fmt.Sprint(clr, fmt.Sprintf(format, args...), reset)
}

// logger helpers

func quiet(outs io.Writer, msg string) {
}

func print(outs io.Writer, msg string) {
	fmt.Fprintf(outs, "%s\n", msg)
}

// public functions

func Debug(format string, args ...interface{}) {
	finfo := ""
	if _, fn, ln, ok := runtime.Caller(1); ok {
		finfo = fmt.Sprintf("%s:%d: ", fn[shortIdx:], ln)
	}
	l.Debug(os.Stderr, l.Format("debug",
		fmt.Sprintf("%s%s", finfo, fmt.Sprintf(format, args...))))
}

func Errorf(format string, args ...interface{}) {
	l.Error(os.Stderr, l.Formatf("error", format, args...))
}

func Error(err error) {
	l.Error(os.Stderr, l.Format("error", err))
}

func Warnf(format string, args ...interface{}) {
	l.Warn(os.Stderr, l.Formatf("warn", format, args...))
}

func Warn(err error) {
	l.Warn(os.Stderr, l.Format("warn", err))
}

func Printf(format string, args ...interface{}) {
	l.Print(os.Stdout, l.Formatf("msg", format, args...))
}

func Print(args ...interface{}) {
	l.Print(os.Stdout, l.Format("msg", args...))
}
