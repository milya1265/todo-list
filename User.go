package todo_app

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name" binding="required"`
	Username string `json:"username" binding="required"`
	Password string `json:"password" binding="required"`
}

type UserList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type ListItem struct {
	Id     int
	UserId int
	ListId int
}
