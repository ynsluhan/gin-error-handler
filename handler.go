package gin_error_handler

import (
	"github.com/gin-gonic/gin"
	R "github.com/ynsluhan/go-r"
	"log"
	"runtime/debug"
)

//全局捕获
func ErrorRecover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				// 打印错误堆栈信息
				log.Printf("[error] message: %v\n  ", r)
				// 打印详细错误
				debug.PrintStack()
				//封装通用json返回
				R.Error(c, r.(error).Error())
				//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
				c.Abort()
			}
		}()
		//加载完 defer recover，继续后续接口调用
		c.Next()
	}
}