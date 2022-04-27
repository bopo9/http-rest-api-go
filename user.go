package socialApp

type User struct {
	Id         int    `json:"-" db:"id" `
	Name       string `json:"name" binding:"required"`
	SecondName string `json:"secondName" binding:"required"`
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
}
