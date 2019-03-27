package models

import (
	db "newland/datebase"
)

type News struct {
	Id        int    `json:"id" form:"id"`
	Title string `json:"title" form:"title"`
	TypeId  string `json:"type_id" form:"type_id"`
	Content  string `json:"content" form:"content"`
	//Token  string `json:"token" form:"token"`
}
func (n *News) AddNews() (id int64, err error) {
	rs, err := db.SqlDB.Exec("INSERT INTO news(title, type_id, content) VALUES (?, ?)", n.Title, n.TypeId, n.Content)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}
func (n *News) GetNews()(news []News, err error)  {
	news = make([]News, 0)
	rows, err := db.SqlDB.Query("SELECT id, title, type_id, content FROM users")
	defer rows.Close()
	if err != nil {
		return
	}

	for rows.Next() {
		var newOne News
		rows.Scan(&newOne.Id, &newOne.Title, &newOne.TypeId, &newOne.Content)
		news = append(news, newOne)
	}

	if err = rows.Err(); err!=nil {
		return
	}
	return
}

func (n *News) GetNewsId(id int, err error) {
	err = db.SqlDB.QueryRow("SELECT id, title, type_id, content FROM users where id = ?", id).Scan(
		&n.Id, &n.Title, &n.TypeId, &n.Content,
	)
	if err != nil {
		return
	}
	return
}