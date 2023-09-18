package Woo

import (
	"fmt"
	"log/slog"
	"runtime"
	"strings"
)

// print stack trace for debug
func trace(message string) string {
	var pcs [32]uintptr
	n := runtime.Callers(3, pcs[:]) // skip first 3 caller

	var str strings.Builder
	str.WriteString(message + "\nTraceback:")
	for _, pc := range pcs[:n] {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}
	return str.String()
}

func Recovery() HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				message := fmt.Sprintf("%s", err)
				//log.Printf("%s\n\n", trace(message))
				slog.Info("recovery", "error", trace(message))
				//以下代码缺少相关实现，暂且先注释掉
				//		c.Fail(http.StatusInternalServerError, "Internal Server Error")
			}
		}()
		c.Next()
	}
}
