package models

type Auth struct {
    ID int `gorm:"primary_key" json:"id"`
    Name string `json:"name"`
    Password string `json:"password"`
}

func CheckAuth(name, password string) bool {
    var auth Auth
    db.Select("id").Where(Auth{Name : name, Password : password}).First(&auth)

    if auth.ID > 0 {
        return true
    }

    return false
}