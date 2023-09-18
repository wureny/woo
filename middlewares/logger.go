package Woo

import (
	"fmt"
	"github.com/wureny/Woo"
	"log/slog"
	"os"
	"time"
)

func (c Context) Logger() HandlerFunc {
	return func(c *Woo.Context) {
		t := time.Now()
		c.Next()
		slog1 := slog.NewJSONHandler(os.Stderr, nil)
		slog2 := slog.New(slog1)
		slog2.Info("Logger", "Status", fmt.Sprintf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t)))
	}
}
