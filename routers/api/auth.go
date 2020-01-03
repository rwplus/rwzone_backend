package api

import (
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/astaxie/beego/validation"

    "rwplus-api/pkg/e"
    "rwplus-api/pkg/util"
    "rwplus-api/models"
)

type auth struct {
    Name string `valid:"Required"; MaxSize(50)`
    Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
    name := c.Query("name")
    password := c.Query("password")

    valid := validation.Validation{}
    a := auth{Name: name, Password: password}
    ok, _ := valid.Valid(&a)

    data := make(map[string]interface{})
    code := e.INVALID_PARAMS
    if ok {
        isExist := models.CheckAuth(name, password)
        if isExist {
            token, err := util.GenerateToken(name, password)
            if err != nil {
                code = e.ERROR_AUTH_TOKEN
            } else {
                data["token"] = token

                code = e.SUCCESS
            }
        } else {
            code = e.ERROR_AUTH
        }
    } else {
        for _, err := range valid.Errors {
            log.Println(err.Key, err.Message)
        }
    }

    c.JSON(http.StatusOK, gin.H{
        "code" : code,
        "msg" : e.GetMsg(code),
        "data" : data,
    })
}