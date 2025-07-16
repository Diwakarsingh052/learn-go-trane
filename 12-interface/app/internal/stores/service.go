package stores

import "app/internal/stores/models"

type UserRepository interface {
	Create(u models.User) (models.User, bool)
	Update(id int, name string) (models.User, bool)
	Delete(id int) bool
	FetchAll() (map[int]models.User, bool)
	FetchUser(id int) (models.User, bool)
}
