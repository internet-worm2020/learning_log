package middleware
import (
	"github.com/gin-gonic/gin"
	"github.com/internet-worm2020/go-pkg/log"
	"time"
)
func LoggerMiddleware1() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime).Seconds()
		// // 请求方式
		reqMethod := c.Request.Method
		// // 请求路由
		reqUrl := c.Request.RequestURI
		// // 状态码
		statusCode := c.Writer.Status()
		// // 请求ip
		clientIP := c.ClientIP()
		log.Info("999",
		log.Int("status_code",statusCode),
		log.Float64("latency_time", latencyTime),
		log.String("client_ip",clientIP),
		log.String("req_uri",      reqUrl),
		log.String("req_method",   reqMethod))
	}
}
