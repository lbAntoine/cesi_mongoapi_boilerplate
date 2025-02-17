package models

type User struct {
	BaseModel `bson:",inline"`
	Email     string `bson:"email" json:"email"`
	Name      string `bson:"name" json:"name"`
	Password  string `bson:"password" json:"password"`
	Active    bool   `bson:"active" json:"active"`
}

func NewUser(email, name, password string) *User {
	return &User{
		Email:    email,
		Name:     name,
		Password: password,
		Active:   true,
	}
}
