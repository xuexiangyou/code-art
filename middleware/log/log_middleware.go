package log

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const TransactionLogKey = "request-id"

func JsonLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"}) //设置日志格式
		logrus.SetLevel(logrus.TraceLevel)                                                 //设置日志级别
		logEntry := logrus.WithFields(logrus.Fields{
			"method":   c.Request.Method,
			"path":     c.Request.RequestURI,
			"referrer": c.Request.Referer(),
			"Request-Id": c.GetHeader(TransactionLogKey),
		})

		//日志设置到上下文中
		c.Set("logger", logEntry)

		c.Next()
	}
}
