package models

import (
	db "newland/datebase"
)

type Types struct {
	Id        int    `json:"id" form:"id"`
	Title string `json:"title" form:"title"`
	Del  string `json:"del" form:"del"`
	//Token  string `json:"token" form:"token"`
}
func (t *Types) AddType() (id int64, err error) {
	rs, err := db.SqlDB.Exec("INSERT INTO types(title, del) VALUES (?, ?)", t.Title, 0)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}
func (t *Types) GetTypes()(types []Types, err error)  {
	typess := make([]Types, 0)
	rows, err := db.SqlDB.Query("SELECT id, title, del FROM types")
	defer rows.Close()
	if err != nil {
		return
	}

	for rows.Next() {
		var types Types
		rows.Scan(&types.Id, &types.Title, &types.Del)
		typess = append(typess, types)
	}

	if err = rows.Err(); err!=nil {
		return
	}
	return
}

func (t *Types) GetTypeId(id int, err error) {
	err = db.SqlDB.QueryRow("SELECT id, title, del FROM users where id = ?", id).Scan(
		&t.Id, &t.Title, &t.Del,
	)
	if err != nil {
		return
	}
	return
}