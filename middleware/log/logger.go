package log

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/xuexiangyou/code-art/config"
	"os"
	"time"
)


type Logs struct {
	AppLog *logrus.Logger
	WebLog *logrus.Logger
}

func NewLogs(config *config.Config) *Logs {
	appLog := initAppLog(config)
	webLog := initWebLog(config)

	return &Logs{
		AppLog: appLog,
		WebLog: webLog,
	}
}

func initAppLog(config *config.Config) *logrus.Logger {
	return initLog(config.Log.AppLogPath)
}

func initWebLog(config *config.Config) *logrus.Logger {
	return initLog(config.Log.WebLogPath)
}

func initLog(logName string) *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	var f *os.File
	var err error

	if _, err = os.Stat(logName); os.IsNotExist(err) {
		f, err = os.Create(logName)
	} else {
		f, err = os.OpenFile(logName,  os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	}

	if err != nil {
		fmt.Println("open logs file failed")
	}
	log.Out = f

	//Set logs level
	log.SetLevel(logrus.InfoLevel)

	return log
}


//LoggerToFile Log to file
func LoggerToFile(log *Logs) gin.HandlerFunc {

	/*//logFilePath := config.Log_FILE_PATH
	//logFileName := config.LOG_FILE_NAME

	//Log file
	//fileName := path.Join(logFilePath, logFileName)
	fileName := "gin.logs" //todo 先写死路径

	//Write to file
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err) //todo
	}

	// instantiation
	logger := logrus.New()

	//Set output
	logger.Out = src

	//Set logs level
	logger.SetLevel(logrus.DebugLevel)

	//Format logs
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})*/

	return func(c *gin.Context) {
		//Start time
		startTime := time.Now()

		//Process request
		c.Next()

		//End time
		endTime := time.Now()

		//Execution time
		latencyTime := endTime.Sub(startTime)

		//Request method
		reqMethod := c.Request.Method

		//Request routing
		reqUri := c.Request.RequestURI

		// status code
		statusCode := c.Writer.Status()

		// request IP
		clientIP := c.ClientIP()

		//Log format
		/*logger.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)*/

		log.WebLog.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqUri,
		}).Info()
	}
}
