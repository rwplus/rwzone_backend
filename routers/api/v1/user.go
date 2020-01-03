package v1

import (
    "github.com/gin-gonic/gin"
    "github.com/astaxie/beego/validation"

    "fmt"
    "strconv"
    "net/http"

    "rwplus-api/models"
    "rwplus-api/pkg/setting"
    "rwplus-api/pkg/e"
)

// 新增用户
func AddUser(c *gin.Context) {
    // url里&name传入的参数
    name := c.Query("name")
    //如果有&state参数，现转换为int
    state, err :=  strconv.Atoi(c.DefaultQuery("state", "0"))
    // 转换时异常，打印失败
    if err != nil {
        fmt.Printf("%v 转换失败！", c.DefaultQuery("state", "0"))
    }
    // url里&created_date传入的参数
    // createdDate := c.Query("created_date")
    // 调用beego的校验包
    valid := validation.Validation{}
    // Ruquired不允许为空
    valid.Required(name, "name").Message("用户名不允许为空")
    // Maxsize表示字符串长度
    valid.MaxSize(name, 100, "name").Message("名字长度不允许超过100")
    //valid.Required(createdDate, "created_date").Message("创建时间不能为空")
    // valid.MaxSize(createdBy, 100, "created_date").Message("创建人时间不允许超过100")
    // range表示枚举，只能用设置0或者1
    valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
    // 验证失败则抛出错误号
    code := e.INVALID_PARAMS
    if ! valid.HasErrors() {
        // 如果根据用户名检索不到用户，则新增用户
        if ! models.ExistUserByName(name) {
            code = e.SUCCESS
            // 传入url中参数如的三个参数name，status，created_date
            models.AddUser(name, state)
        } else {
            // 已经存在此用户
            code = e.ERROR_EXIST_USER
        }
    }
    // 返回接口请求数据
    c.JSON(http.StatusOK, gin.H{
        "code" : code,
        "msg" : e.GetMsg(code),
        "data" : make(map[string]string),
    })
}

// 获取用户信息
func GetUser(c *gin.Context) {
    // &name传入的信息
    name := c.Query("name")

    maps := make(map[string]interface{})
    data := make(map[string]interface{})

    if name != "" {
        maps["name"] = name
    }


    if arg := c.Query("state"); arg != "" {
        state,err := strconv.Atoi(arg)
        if err != nil {
            fmt.Printf("%v 转换失败！", arg)
        }

        maps["state"] = state
    }


    code := e.SUCCESS

    data["lists"] = models.GetUser(setting.PageSize, maps)
    data["total"] = models.GetUserTotal(maps)

    c.JSON(http.StatusOK, gin.H{
        "code" : code,
        "msg" : e.GetMsg(code),
        "data" : data,
    })

}


//修改用户信息
func EditUser(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))

    if err != nil {
        fmt.Println("%v 转换失败！", c.Param("id"))
    }
    name := c.Query("name")
    phone := c.Query("phone")

    valid := validation.Validation{}

    var state int = -1
    if arg := c.Query("state"); arg != "" {
        state, err := strconv.Atoi(arg)
        if err != nil {
            fmt.Println("%v 转换失败", arg)
        }

        valid.Range(state, 0, 1, "state").Message("状态只允许0或者1")
    }

    valid.Required(id, "id").Message("ID不能为空")
    valid.Required(phone, "phone").Message("手机号码不允许为空")
    valid.MaxSize(phone, 11, "phone").Message("手机号码为11位")
    valid.MaxSize(name, 20, "name").Message("用户名长度不允许超过20")

    code := e.INVALID_PARAMS

    if ! valid.HasErrors() {
        code = e.SUCCESS
        if models.ExistUserByID(id) {
            data := make(map[string]interface{})
            if name != "" {
                data["name"] = name
            }
            if state != -1 {
                data["state"] = state
            }

            models.EditUser(id, data)
        } else {
            code = e.ERROR_NOT_EXIST_USER
        }
    }

    c.JSON(http.StatusOK, gin.H{
        "code" : code,
        "msg" : e.GetMsg(code),
        "data" : make(map[string]string),
    })
}

//删除用户
func DeleteUser(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))

    if err != nil {
        fmt.Println("%v 转换失败！", c.Param("id"))
    }

    valid := validation.Validation{}
    valid.Min(id, 1, "id").Message("ID必须大于0")

    code := e.INVALID_PARAMS
    if ! valid.HasErrors() {
        code = e.SUCCESS
        if models.ExistUserByID(id) {
            models.DeleteUser(id)
        } else {
            code = e.ERROR_NOT_EXIST_USER
        }
    }

    c.JSON(http.StatusOK, gin.H{
        "code" : code,
        "msg" : e.GetMsg(code),
        "data" : make(map[string]string),
    })
}
