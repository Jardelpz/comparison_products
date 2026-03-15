package handler

import (
	"log"
	"time"

	"comparison/internal/trace"

	"github.com/gin-gonic/gin"
)

const TraceIDHeader = "X-Trace-Id"

func TraceMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := c.GetHeader(TraceIDHeader)
		if traceID == "" {
			traceID = trace.NewTraceID()
		}

		start := time.Now()

		// coloca no context da request
		ctx := trace.WithTraceID(c.Request.Context(), traceID)
		c.Request = c.Request.WithContext(ctx)

		c.Header(TraceIDHeader, traceID)

		c.Next()

		dur := time.Since(start)
		log.Printf("trace end: method=%s path=%s status=%d trace_id=%s duration_ms=%d", c.Request.Method, c.Request.URL.Path, c.Writer.Status(), traceID, dur.Milliseconds())
	}
}
