package interfaces

import "github.com/username/depogunabangunan/apps/model"

// IUserRepository defines the interface for user repository operations
type IUserRepository interface {
	GetUserByID(id int64) (*model.User, error)
	GetUserByUsername(username string) (*model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUser(id int64) error
}
