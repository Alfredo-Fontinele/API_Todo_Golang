package services

import (
	db "api-golang/database"
	models "api-golang/models"
)

func Create(todo models.Todo) (id int64, err error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer connection.Close()

	sql := `
		INSERT INTO 
			todos (title, description, done)
		VALUES 
			($1, $2, $3)
		RETURNING id`

	err = connection.QueryRow(sql, todo.Title, todo.Description, todo.Done).Scan(&id)
	return
}

func Get(id int64) (todo models.Todo, err error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer connection.Close()

	sql := `SELECT * FROM todos where id = $1`
	row := connection.QueryRow(sql, id)
	err = row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)

	return
}

func GetAll() (todos []models.Todo, err error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer connection.Close()
	sql := `SELECT * FROM todos`
	rows, err := connection.Query(sql)
	if err != nil {
		return
	}
	for rows.Next() {
		var todo models.Todo
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)
		if err != nil {
			continue
		}
	}

	return
}

func Update(id int64, todo models.Todo) (int64, error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer connection.Close()

	res, err := connection.Exec(`
		UPDATE todos
		SET title=$1, description=$2, done=$3
		WHERE id=$4
	`, todo.Title, todo.Description, todo.Done, id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func Delete(id int64) (int64, error) {
	connection, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer connection.Close()

	res, err := connection.Exec(`
		DELETE FROM todos WHERE id=$1
	`, id)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
