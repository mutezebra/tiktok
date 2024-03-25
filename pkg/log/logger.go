package log

import (
	"fmt"
	"runtime"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
	pool sync.Pool
}

// Panic refactors the panic function within logrus. It records
// the stack info and outputs it to logrus`s 'Out'
func (l *Logger) Panic(v any) {
	var str *string
	str = l.formatCallStack(v)
	l.Logger.Panic(*str)
}

// Error refactors the error function within logrus. It records
// the stack info and outputs it to logrus`s 'Out'
func (l *Logger) Error(v any) {
	var str *string
	str = l.formatCallStack(v)
	l.Logger.Error(str)
}

func (l *Logger) formatCallStack(v any) *string {
	builder := l.pool.Get().(*strings.Builder)

	err, ok := v.(error)
	if ok {
		builder.WriteString(fmt.Sprintf("%s  \n", err.Error()))
	} else {
		builder.WriteString(fmt.Sprintf("%v  \n", v))
	}

	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(3, pcs[:])
	st := make([]uintptr, n)
	st = pcs[0:n]
	for _, pc := range st {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		builder.WriteString(fmt.Sprintf("\t%s:%d %s\n", file, line, fn.Name()))
	}

	str := builder.String()
	l.clean(builder)
	l.pool.Put(builder)
	return &str
}

func (l *Logger) clean(str *strings.Builder) {
	str.Reset()
}

// Restore replaces "\\n" with "\n" in the panic msg.
func Restore(msg string) string {
	return strings.ReplaceAll(strings.ReplaceAll(msg, "\\n", "\n"), "\\t", "\t")
}
