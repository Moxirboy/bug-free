package usecase

import "bug-free/internal/model"

func (c *User) Create(user model.User) (int, error) {
	return c.repo.Create(user)
}
func (c *User) Read(id model.UserGetBYId) (string,error) {
	return c.repo.ReadById(id)
}
func (c *User) Update(user model.User)(int, error ){
	return c.repo.Update(user )
}
func (c *User) Delete(user model.User) (int,error) {
	return c.repo.Delete(user)
}
