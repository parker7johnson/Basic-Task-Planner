package contorller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/service"

	"github.com/jmoiron/sqlx"
)

type Response struct {
	response map[string]interface{}
}

func NewRes() Response {
	return Response{
		make(map[string]interface{}),
	}
}

var DB *sqlx.DB

func BeginServer(db *sqlx.DB) {
	DB = db
	http.HandleFunc("/", handleInsertTask)
	http.HandleFunc("/update", handleUpdateTask)
	http.HandleFunc("/delete", handleDeleteTask)
	http.HandleFunc("/get", getAllTasks)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}

	fmt.Println("Listening on Port 8080")
}

func handleInsertTask(w http.ResponseWriter, r *http.Request) {
	res := NewRes()
	w.Header().Set("Content-Type", "application/json")
	status, err := service.Insert(r, DB)

	res.response["status"] = status
	if err != nil {
		fmt.Printf("Error inserting task: %v\n", err)
		res.response["message"] = "Error inserting task"
	} else {
		fmt.Printf("Successfully inserted task\n")
		res.response["message"] = "Success"
	}
	json.NewEncoder(w).Encode(res.response)
}

func handleUpdateTask(w http.ResponseWriter, r *http.Request) {
	res := NewRes()
	w.Header().Set("Content-Type", "application/json")
	status, err := service.Update(r, DB)

	res.response["status"] = status

	if err != nil {
		fmt.Printf("Error updating task: %v\n", err)
		res.response["message"] = "Error updating task"
	} else {
		fmt.Printf("Successfully updated task\n")
		res.response["message"] = "Success"
	}
}

func handleDeleteTask(w http.ResponseWriter, r *http.Request) {
	res := NewRes()
	w.Header().Set("Content-Type", "application/json")

	status, err := service.Delete(r, DB)

	res.response["status"] = status
	if err != nil {
		fmt.Printf("Error deleting task: %v\n", err)
		res.response["message"] = "Error deleting task"
	} else {
		fmt.Printf("Successfully deleted task\n")
		res.response["message"] = "Success"
	}

	json.NewEncoder(w).Encode(res.response)
}

func getAllTasks(w http.ResponseWriter, r *http.Request) {
	res := NewRes()
	w.Header().Set("Content-Type", "application/json")

	tasks, status, err := service.Get(r, DB)

	res.response["status"] = status
	if err != nil {
		fmt.Printf("Error getting tasks: %v\n", err)
		res.response["message"] = "Error getting tasks"
	} else {
		fmt.Printf("Successfully got tasks\n")
		res.response["tasks"] = tasks
	}

	json.NewEncoder(w).Encode(res.response)
}
