package interfaces

import (
	"github.com/alimahboubi/gochat/internal/user-service/domain/entities"
	"github.com/alimahboubi/gochat/internal/user-service/domain/valueobjects"
)

type UserRepository interface {
	Create(user *entities.User) error
	FindById(id *valueobjects.UserId) (*entities.User, error)
	FindByEmail(email *valueobjects.Email) (*entities.User, error)
	Exists(id *valueobjects.UserId) (bool, error)
	ExistsByEmail(email *valueobjects.Email) (bool, error)
	Update(user *entities.User) error
}
