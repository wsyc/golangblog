package models

import (
	db "newland/datebase"
)

type Person struct {
	Id        int    `json:"id" form:"id"`
	UserName string `json:"user_name" form:"user_name"`
	Password  string `json:"password" form:"password"`
	//Token  string `json:"token" form:"token"`
}
func (p *Person) AddPerson() (id int64, err error) {
	rs, err := db.SqlDB.Exec("INSERT INTO users(user_name, password) VALUES (?, ?)", p.UserName, p.Password)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}
func (p *Person) GetPersons()(person []Person, err error)  {
	persons := make([]Person, 0)
	rows, err := db.SqlDB.Query("SELECT id, user_name, password FROM users")
	defer rows.Close()
	if err != nil {
		return
	}

	for rows.Next() {
		var person Person
		rows.Scan(&person.Id, &person.UserName, &person.Password)
		persons = append(persons, person)
	}

	if err = rows.Err(); err!=nil {
		return
	}
	return
}

func (p *Person) GetUserId(id int, err error) {
	err = db.SqlDB.QueryRow("SELECT id, user_name, password FROM users where id = ?", id).Scan(
		&p.Id, &p.UserName, &p.Password,
	)
	if err != nil {
		return
	}
	return
}

//func (p *Person) LoginUserId(Uid int, err error) string {
//	//err = db.SqlDB.QueryRow("SELECT id, frist_name, last_name, token FROM person where token = ?", Token).Scan(
//	//	&p.Id, &p.FirstName, &p.LastName,
//	//)
//	person := Person{Id:Uid}
//	stmt, err := db.SqlDB.Prepare("UPDATE person SET token=? WHERE id=?")
//	defer stmt.Close()
//	if err != nil {
//		log.Fatalln(err)
//	}
//	token:=p.setToken()
//	rs, err := stmt.Exec(token, person.Id)
//	ra, err := rs.RowsAffected()
//	if err != nil {
//		log.Fatalln(err)
//	}
//	if ra >0 {
//		return token
//	}
//	if err != nil {
//		return ""
//	}
//	return ""
//}

func (p *Person) SetToken() string{
	//util := util2.Util{}
	return ""
	//token := jwt.New(jwt.SigningMethodHS256)
	//claims := make(jwt.MapClaims)
	//claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	//claims["iat"] = time.Now().Unix()
	////claims["uid"] = uid
	//token.Claims = claims
	//Token, err := token.SignedString([]byte("key"))
	//if err != nil {
	//	return Token
	//}
	//return ""
}