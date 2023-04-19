package repos

import (
	"gorm-viper/orm"
	"gorm-viper/sample/orm/with_repos/entities"
)

type userRepository struct {
	*orm.Repository
}

func User() *userRepository {
	return &userRepository{r}
}

// Override Get method
func (r *userRepository) Get(id string) (entities.User, error) {
	var res entities.User
	err := r.Database.First(&res, "id = ?", id).Error
	return res, err
}

// Add GetByName method
func (r *userRepository) GetByName(name string) (entities.User, error) {
	var res entities.User
	err := r.Database.First(&res, "name = ?", name).Error
	return res, err
}
