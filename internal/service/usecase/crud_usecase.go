package usecase

import "bug-free/internal/model"

func (c *Crud) Create(user model.User) (int, error) {
	return c.repo.Create(user)
}
func (c *Crud) Read() []string {
	return c.repo.Read()
}
func (c *Crud) Update() error {
	return c.repo.Update()
}
func (c *Crud) Delete() error {
	return c.repo.Delete()
}
