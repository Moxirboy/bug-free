package repo

import (
	"fmt"
	"golang-project-template/internal/model"
)

func (p *postgres) Create(user model.User) (int, error) {
	var id int
	query := fmt.Sprintf("insert into user(name,password) values(?,?)")
	row := p.db.QueryRow(query, user.Name, user.Name, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
func (p *postgres) Read() []string {
	return nil
}
func (p *postgres) Update() error {
	return nil
}
func (p *postgres) Delete() error {
	return nil
}
