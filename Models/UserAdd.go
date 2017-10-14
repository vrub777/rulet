package Models

type UserAdd struct {
	Id           int
	Name         string
	Email        string
	Password     string
	IdsRole      []int
	Roles        []UserRole
	IsAddUser    bool
	NameOkButton string
	NameAction   string
}
