package usecase

import "bug-free/internal/model"

func (c *User) Create(user model.User) (int, error) {
	return c.repo.Create(user)
}
func (c *User) Read() []string {
	return c.repo.Read()
}
func (c *User) Update() error {
	return c.repo.Update()
}
func (c *User) Delete() error {
	return c.repo.Delete()
}
