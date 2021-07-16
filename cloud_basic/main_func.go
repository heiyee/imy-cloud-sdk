package cloud_basic

import (
	"github.com/gin-gonic/gin"
)

// run web server
func StartServer(initializer Initializer) {
	r := gin.Default()
	r.NoRoute(func(context *gin.Context) {
		context.AbortWithStatus(404)
	})
	handlers := initializer.GetHandlers()
	if len(handlers) > 0 {
		for _, v := range handlers {
			switch v.Method {
			case 1:
				r.GET(v.Url, func(context *gin.Context) {
					v.Handler(v.ApiId, context)
				})
			case 2:
				r.POST(v.Url, func(context *gin.Context) {
					v.Handler(v.ApiId, context)
				})
			}
		}
	}
	r.Run(initializer.GetAddr())
}
