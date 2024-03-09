package middleware

import (
	"bytes"
	"io"
	"jianji-server/entity"
	"jianji-server/utils"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
)

type responseLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w responseLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func addLog(log *entity.ResponseLog) {
	utils.DB.Create(log)
}

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		rlw := &responseLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = rlw
		// 将 request body 储存起来供多次使用
		bodyByte, _ := io.ReadAll(c.Request.Body)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyByte))
		c.Next()

		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyByte))
		dump, _ := httputil.DumpRequest(c.Request, true)
		sessionId, _ := c.Get("SessionId")
		traceId, _ := c.Get("TraceId")
		addLog(&entity.ResponseLog{
			Universal:  entity.Universal{},
			StatusCode: rlw.Status(),
			RequestURL: c.Request.URL.String(),
			SessionId:  sessionId.(string),
			TraceId:    traceId.(string),
			Method:     c.Request.Method,
			Request:    string(dump),
			Response:   rlw.body.String(),
		})
	}
}
