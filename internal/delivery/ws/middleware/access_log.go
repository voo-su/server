package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
	clickhouseModel "voo.su/internal/infrastructure/clickhouse/model"
	clickhouseRepo "voo.su/internal/infrastructure/clickhouse/repository"
)

func AccessLogMiddleware(accessLogRepo *clickhouseRepo.AccessLogRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer = responseWriter{c.Writer, bytes.NewBuffer([]byte{})}
		access := newAccessLogStore(c)
		if err := access.init(); err != nil {
			c.Abort()
			c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error: %s", err.Error()))
			return
		}
		c.Next()
		access.load()
		if c.Request.Method != "OPTIONS" {
			if err := accessLogRepo.Create(context.Background(), access.data); err != nil {
				log.Printf("Failed to save access log: %s", err)
			}
		}
	}
}

type responseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

type AccessLogStore struct {
	ctx       *gin.Context
	startTime time.Time
	data      *clickhouseModel.AccessLog
}

func newAccessLogStore(c *gin.Context) *AccessLogStore {
	return &AccessLogStore{
		ctx:       c,
		startTime: time.Now(),
		data:      &clickhouseModel.AccessLog{},
	}
}

func (a *AccessLogStore) init() error {
	hostname, _ := os.Hostname()
	headers := fmt.Sprintf("%v", a.ctx.Request.Header)
	body, err := io.ReadAll(a.ctx.Request.Body)
	if err != nil {
		return err
	}

	a.data = &clickhouseModel.AccessLog{
		RequestId:       a.ctx.Request.Header.Get("X-Request-ID"),
		RequestMethod:   a.ctx.Request.Method,
		RequestURI:      a.ctx.Request.URL.Path,
		RequestHeader:   headers,
		RequestBody:     string(body),
		RequestTime:     a.startTime,
		RequestQuery:    a.ctx.Request.URL.RawQuery,
		ResponseHeader:  "",
		ResponseBodyRaw: "",
		ResponseTime:    time.Now(),
		HttpUserAgent:   a.ctx.Request.UserAgent(),
		HttpStatus:      0,
		HostName:        hostname,
		ServerName:      a.ctx.Request.Host,
		RemoteAddr:      a.ctx.RemoteIP(),
	}

	if a.data.RequestId == "" {
		a.data.RequestId = strings.ReplaceAll(uuid.New().String(), "-", "")
	}

	if strings.HasPrefix(a.ctx.GetHeader("Content-Type"), "application/json") {
		var jsonBody map[string]interface{}
		_ = json.Unmarshal(body, &jsonBody)
		a.data.RequestBody = fmt.Sprintf("%v", jsonBody)
	}

	a.ctx.Request.Body = io.NopCloser(bytes.NewReader(body))

	return nil
}

func (a *AccessLogStore) load() {
	writer := a.ctx.Writer.(responseWriter)
	headers := fmt.Sprintf("%v", writer.Header())

	a.data.ResponseHeader = headers
	a.data.ResponseTime = time.Now()
	a.data.RequestDuration = fmt.Sprintf("%.3f", time.Since(a.startTime).Seconds())
	a.data.HttpStatus = writer.Status()
	a.data.ResponseBodyRaw = writer.body.String()

	if strings.HasPrefix(writer.Header().Get("Content-Type"), "application/json") {
		var body map[string]interface{}
		_ = json.Unmarshal(writer.body.Bytes(), &body)
		a.data.ResponseBodyRaw = fmt.Sprintf("%v", body)
	}
}
