package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
)

type Task struct {
	Id       string    `json:"id,omitempty"`
	Name     string    `json:"name"`
	Deadline time.Time `json:"deadline"`
	Status   string    `json:"status"`
}

const (
	SUCCESS               = 200
	BAD_REQUEST           = 400
	INTERNAL_SERVER_ERROR = 500
)

func Insert(r *http.Request, db *sqlx.DB) (int, error) {

	var task Task
	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		fmt.Printf("Error decoding json: %v\n", err)
		return BAD_REQUEST, err
	}
	sql := "insert into task (name, deadline, status) values (?,?,?)"

	stmt, err := db.Prepare(sql)

	if err != nil {
		fmt.Printf("Error preparing statement: %v\n", err)
		return INTERNAL_SERVER_ERROR, err
	}

	_, err = stmt.Exec(task.Name, task.Deadline, task.Status)

	if err != nil {
		fmt.Printf("Error inserting task: %v\n", err)
		return INTERNAL_SERVER_ERROR, err
	}

	return SUCCESS, nil

}

func Update(r *http.Request, db *sqlx.DB) (int, error) {

	var task Task
	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		fmt.Printf("Error decoding json: %v\n", err)
		return BAD_REQUEST, err
	}

	sql := "update task set name = ?, deadline = ?, status = ? where id = ?"

	stmt, err := db.Prepare(sql)

	if err != nil {
		fmt.Printf("Error preparing statement: %v\n", err)
		return INTERNAL_SERVER_ERROR, err
	}

	_, err = stmt.Exec(task.Name, task.Deadline, task.Status, task.Id)

	if err != nil {
		fmt.Printf("Error updating task: %v\n", err)
		return INTERNAL_SERVER_ERROR, err
	}

	return SUCCESS, nil
}

func Delete(r *http.Request, db *sqlx.DB) (int, error) {
	var task Task
	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		fmt.Printf("Error decoding json: %v\n", err)
		return BAD_REQUEST, err
	}

	sql := "delete from task where id = ?"

	stmt, err := db.Prepare(sql)

	if err != nil {
		fmt.Printf("Error preparing statement: %v\n", err)
		return INTERNAL_SERVER_ERROR, err
	}

	_, err = stmt.Exec(task.Id)

	if err != nil {
		fmt.Printf("Error deleting task: %v\n", err)
		return INTERNAL_SERVER_ERROR, err
	}

	return SUCCESS, nil

}

func Get(r *http.Request, db *sqlx.DB) ([]Task, int, error) {
	var tasks []Task

	err := db.Select(&tasks, "select * from task")

	if err != nil {
		fmt.Printf("Error querying tasks: %v\n", err)
		return nil, INTERNAL_SERVER_ERROR, err
	}

	return tasks, SUCCESS, nil
}
