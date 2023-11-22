package repo

import (
	"bug-free/internal/model"
	
)

func (p *postgres) Create(user model.User) (int, error) {
	var id int
//	query := fmt.Sprintf("insert into \"users\" (name,password) values($1,$2) RETURNING id")

	query := `
	insert into users (name,password) values($1,$2) RETURNING id
	`

	row := p.db.QueryRow(query, user.Name, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
func (p *postgres) Read() []string {
	// query:=fmt.Sprintf("select * from %s",_const.Database_name)

	return nil
}
func (p *postgres) Update() error {
	return nil
}
func (p *postgres) Delete() error {
	return nil
}
