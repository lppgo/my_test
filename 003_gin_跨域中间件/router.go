package Router
import (
    "github.com/gin-gonic/gin"
    "gin/Controllers"
    "gin/Middlewares"
)
func InitRouter() {
    router := gin.Default()
    // 要在路由组之前全局使用「跨域中间件」, 否则OPTIONS会返回404
    router.Use(Middlewares.Cors())
    v1 := router.Group("v1")
    {
        v1.GET("/jsontest", Controllers.JsonTest)
        v1.POST("/jsonposttest", Controllers.JsonPostTest)
    }
    
    router.Run(":8080")
}
