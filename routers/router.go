package routers

import (
    "github.com/gin-gonic/gin"
    
    "rwplus-api/routers/api/v1"
    "rwplus-api/pkg/setting"
    // 引入auth.go
    "rwplus-api/routers/api"
    // 引入jwt
    "rwplus-api/middleware/jwt"
)


func InitRouter() *gin.Engine {
    r := gin.New()

    r.Use(gin.Logger())

    r.Use(gin.Recovery())

    gin.SetMode(setting.RunMode)

    // 获取token
    r.GET("/auth", api.GetAuth)

    apiv1 := r.Group("/api/v1")
    apiv1.Use(jwt.JWT())
    {
        //获取标签列表
        apiv1.GET("/users", v1.GetUser)
        //新建标签
        apiv1.POST("/users", v1.AddUser)
        //更新指定标签
        apiv1.PUT("/users/:id", v1.EditUser)
        //删除指定标签
        apiv1.DELETE("/users/:id", v1.DeleteUser)
    }

    return r
}
