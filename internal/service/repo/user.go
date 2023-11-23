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
func (p *postgres) ReadById(id model.UserGetBYId) (string,error) {
	query:=`
	select name from users where id=$1 
	`
	var name string
	rows:=p.db.QueryRow(query,id.Id)

        if err := rows.Scan(&name); err != nil {
          
            return "", err
        }
	return name,nil
}
func (p *postgres) Update(user model.User) (int,error) {
	var id int
	query:=`
	update users
	set password=$1
	where name=$2
	RETURNING id
	`
	row:=p.db.QueryRow(query,user.Password,user.Name)
	if err:=row.Scan(&id);err!=nil{
		return 0,err
	}
	return id,nil
}
func (p *postgres) Delete(user model.User) (int,error) {
	var id int 
	query:=`
	delete from users where name=$1 returning id
	`
	row:=p.db.QueryRow(query,user.Name)
	if err:=row.Scan(&id);err!=nil{
		return 0,err
	}
	return id ,  nil
}
