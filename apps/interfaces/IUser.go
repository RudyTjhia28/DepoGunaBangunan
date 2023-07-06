package interfaces

import "depogunabangunan/apps/model"

// IUserRepository defines the interface for user repository operations
type IUserService interface {
	GetUserByID(id int64) (*model.User, error)
	GetUserByUsername(username string) (*model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUser(id int64) error
}
