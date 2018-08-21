package logger

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	_ logLevel = iota
	LFatal
	LDebug
	LInfo
	LError
	LOff
)

var (
	r *roga = New(os.Stdout, LInfo)
)

type roga struct {
	w     io.Writer
	Level logLevel
}

type logLevel int

func New(w io.Writer, level logLevel) *roga {
	return &roga{w: w, Level: level}
}

func SetLevel(level logLevel) {
	r.Level = level
}

func SetOut(w io.Writer) {
	r.w = w
}

func Error(v ...interface{}) {
	if r.Level <= LError {
		r.print(LError, v...)
	}
}

func Errorf(f string, v ...interface{}) {
	if r.Level <= LError {
		r.printf(LError, f, v...)
	}
}

func Info(v ...interface{}) {
	if r.Level <= LInfo {
		r.print(LInfo, v...)
	}
}

func Infof(f string, v ...interface{}) {
	if r.Level <= LInfo {
		r.printf(LInfo, f, v...)
	}
}

func Debug(v ...interface{}) {
	if r.Level <= LDebug {
		r.print(LDebug, v...)
	}
}

func Debugf(f string, v ...interface{}) {
	if r.Level <= LDebug {
		r.printf(LDebug, f, v...)
	}
}

func Fatal(v ...interface{}) {
	r.print(LFatal, v...)
	os.Exit(1)
}

func Fatalf(f string, v ...interface{}) {
	r.printf(LFatal, f, v...)
	os.Exit(1)
}

func (r *roga) writeHeader(level logLevel) {
	// _, f, l, ok := runtime.Caller(3)
	// if !ok {
	// 	f = "???"
	// 	l = 0
	// }
	// fmt.Fprintf(r.w, "%s %s:%d [%s]: ", time.Now().Format("2006/01/02 15:04:05"), f, l, getLevelString(level))

	// default: YYYY/MM/DD HH:MM:SS [LEVEL] this is the err message
	fmt.Fprintf(r.w, "%s [%s]: ", time.Now().Format("2006/01/02 15:04:05"), getLevelString(level))
}

func (r *roga) print(level logLevel, v ...interface{}) {
	r.writeHeader(level)
	fmt.Fprintln(r.w, v...)
}

func (r *roga) printf(level logLevel, f string, v ...interface{}) {
	r.writeHeader(level)
	fmt.Fprintf(r.w, f, v...)
}

func getLevelString(level logLevel) string {
	switch level {
	case LFatal:
		return "FATAL"
	case LDebug:
		return "DEBUG"
	case LInfo:
		return "INFO"
	case LError:
		return "ERROR"
	default:
		return ""
	}
}

func (r *roga) Error(v ...interface{}) {
	if r.Level <= LError {
		r.print(LError, v...)
	}
}

func (r *roga) Errorf(f string, v ...interface{}) {
	if r.Level <= LError {
		r.printf(LError, f, v...)
	}
}

func (r *roga) Info(v ...interface{}) {
	if r.Level <= LInfo {
		r.print(LInfo, v...)
	}
}

func (r *roga) Infof(f string, v ...interface{}) {
	if r.Level <= LInfo {
		r.printf(LInfo, f, v...)
	}
}

func (r *roga) Debug(v ...interface{}) {
	if r.Level <= LDebug {
		r.print(LDebug, v...)
	}
}

func (r *roga) Debugf(f string, v ...interface{}) {
	if r.Level <= LDebug {
		r.printf(LDebug, f, v...)
	}
}

func (r *roga) Fatal(v ...interface{}) {
	r.print(LFatal, v...)
	os.Exit(1)
}

func (r *roga) Fatalf(f string, v ...interface{}) {
	r.printf(LFatal, f, v...)
	os.Exit(1)
}
