package models

import (
    "time"
    "fmt"
    "github.com/jinzhu/gorm"
)

type User struct {
    // 嵌套入models.go定义的Model结构体
    Model
    // 自定义用户表的字段
    Name string `json:"name"`
    Password string `json:"password"`
    Address string `json:"address"`
    Phone string `json:"phone"`
    State int `json:"state"`
    Vip int `json:"vip"`
    CreatedDate int `json:"created_date"`
}

// 获取所有的用户用心
func GetUser(pageSize int, maps interface {}) (users []User) {
    db.Where(maps).Find(&users)
    // db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
    return
}

// 获取用户总数
func GetUserTotal(maps interface{}) (count int) {
    db.Model(&User{}).Where(maps).Count(&count)

    return
}

// 检索用户是否存在
func ExistUserByName(name string) bool {
    var user User
    db.Select("id").Where("name = ?", name).First(&user)
    if user.ID > 0 {
        return true
    }

    return false
}
// 设置插入用户数据的时间
func (user *User) BeforeCreate(scope *gorm.Scope) error {
    scope.SetColumn("CreatedDate", time.Now().Unix())
    fmt.Println(time.Now().Unix())
    return nil
}

// 新增用户
// func AddUser(name string, state int, createdDate string) bool{
func AddUser(name string, state int) bool{
    db.Create(&User {
        Name : name,
        State : state,
        //CreatedDate : createdDate,
    })
    return true
}

// 根据id检索用户
func ExistUserByID(id int) bool {
    var user User
    db.Select("id").Where("id = ?", id).First(&user)
    if user.ID > 0 {
        return true
    }

    return false
}


func DeleteUser(id int) bool {
    db.Where("id = ?", id).Delete(&User{})
    return true
}

func EditUser(id int, data interface{}) bool {
    db.Model(&User{}).Where("id = ?", id).Updates(data)

    return true
}