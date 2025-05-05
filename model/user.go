// model/user.go
package model

type User struct {
	Username string
	Password string
	Role     string // "user" or "admin"
}
