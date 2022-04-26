package socialApp

type User struct {
	Id         int    `json:"-"'`
	Name       string `json:"name"`
	SecondName string `json:"secondName"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}
