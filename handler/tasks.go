package handler

import (
	"RADserver/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func TasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getTasks(w,r)

		w.Header().Set("Content-Type", "application/json")
		fmt.Println("'GET'-response sent to /tasks on", time.Now().Format(time.RFC850))
	case "POST":
		createTask(w,r)

		reqMessage := "'POST'-request sent."
		w.Header().Set("Content-Type", "application/json") 
		fmt.Fprintf(w,`{"request": "%s"}`, reqMessage)
		fmt.Println("'POST'-response sent to /tasks on", time.Now().Format(time.RFC850))
	case "PUT":
		updateTask(w,r)

		reqMessage := "'PUT'-request sent."
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w,`{"request": "%s"}`, reqMessage)
		fmt.Println("'PUT'-response sent to /tasks on", time.Now().Format(time.RFC850))
	case "DELETE":
		deleteTask(w,r)	

		reqMessage := "'DELETE'-request sent."
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w,`{"request": "%s"}`, reqMessage)
		fmt.Println("'DELETE'-response sent to /tasks on", time.Now().Format(time.RFC850))
	}
}
func getTasks(w http.ResponseWriter, r *http.Request){
	var tasksData []models.Task
	byteData, _ := os.ReadFile("db/tasks.json")
	json.Unmarshal(byteData, &tasksData)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(tasksData)
}
func createTask(w http.ResponseWriter, r *http.Request){
	var newTask models.Task
	json.NewDecoder(r.Body).Decode(&newTask)

	var tasksData []models.Task
	byteData,_:= os.ReadFile("db/tasks.json")
	json.Unmarshal(byteData, &tasksData)

	var userData []models.User
	userByte,_ := os.ReadFile("db/users.json")
	json.Unmarshal(userByte, &userData)

	var userFound bool

	for i := 0; i < len(userData); i++ {
		if userData[i].Id == newTask.UserID {
			newTask.ID = len(tasksData)+1
			newTask.CreatedTime = time.Now().Format(time.RFC850)
			newTask.UpdatedTime = time.Now().Format(time.RFC850)
			userFound = true
			break
		}
	}
	if !userFound {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "no user found with this UserID.")
		return
	}

	tasksData = append(tasksData, newTask)

	res, _ := json.Marshal(tasksData)
	os.WriteFile("db/tasks.json", res, 0)
	
	w.WriteHeader(http.StatusCreated)
	fmt.Println("\n_____________________________________________")
	fmt.Println("Created new task at", time.Now().Format(time.RFC850))
	fmt.Println("_____________________________________________")
	fmt.Println("____________________________")
	fmt.Println("ID is: ", newTask.ID)
	fmt.Println("____________________________")
	fmt.Println("User ID is: ", newTask.UserID)
	fmt.Println("____________________________")	
}
func updateTask(w http.ResponseWriter, r *http.Request){
	var updateTask models.Task
	json.NewDecoder(r.Body).Decode(&updateTask)

	var tasksData []models.Task
	byteData,_ := os.ReadFile("db/tasks.json")
	json.Unmarshal(byteData, &tasksData)

	var userData []models.User
	userByte, _ := os.ReadFile("db/users.json")
	json.Unmarshal(userByte, &userData)

	var userFound bool
	var taskFound bool

	for i := 0; i < len(userData); i++ {
		if userData[i].Id == updateTask.UserID {
			for j := 0; j < len(tasksData); j++ {
				if tasksData[j].ID == updateTask.ID {
					tasksData[j].TaskContain = updateTask.TaskContain
					tasksData[j].UpdatedTime = time.Now().Format(time.RFC850)
					taskFound = true
					break
				}
			}
			if !taskFound {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "Task with such an ID not found.")
			}
			userFound = true
			break
		}
	}
	if !userFound {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w,"Task with such an UserID not found.")
		return
	}

	res, _ := json.Marshal(tasksData)
	os.WriteFile("db/tasks.json", res, 0)

	w.WriteHeader(http.StatusAccepted)
	fmt.Println("\n_____________________________________________")
	fmt.Println("Updated task at", time.Now().Format(time.RFC850))
	fmt.Println("_____________________________________________")
	fmt.Println("____________________________")
	fmt.Println("ID is: ", updateTask.ID)
	fmt.Println("____________________________")
	fmt.Println("User ID is: ", updateTask.UserID)
	fmt.Println("____________________________")
}
func deleteTask(w http.ResponseWriter, r *http.Request){
	var deleteTask models.Task
	json.NewDecoder(r.Body).Decode(&deleteTask)
	
	var tasksData []models.Task
	byteData,_ := os.ReadFile("db/tasks.json")
	json.Unmarshal(byteData, &tasksData)

	var userData []models.User
	userByte,_ := os.ReadFile("db/users.json")
	json.Unmarshal(userByte, &userData)

	var userFound bool
	var taskfound bool

	for i := 0; i < len(userData); i++ {
		if userData[i].Id == deleteTask.UserID {
			for j := 0; j < len(tasksData); j++ {
				if tasksData[j].ID == deleteTask.ID {
					tasksData = append(tasksData[:j], tasksData[j+1:]...)
					taskfound = true
					break
				}
			}
			if !taskfound {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "Task wich such an ID not found")
				return
			}
			userFound = true
			break
		}
	}
	if !userFound {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Task witch such an UserID not found.")
		return
	}

	res,_ := json.Marshal(tasksData)
	os.WriteFile("db/tasks.json", res, 0)

	w.WriteHeader(http.StatusAccepted)
	fmt.Println("\n_____________________________________________")
	fmt.Println("Deleted task at", time.Now().Format(time.RFC850))
	fmt.Println("_____________________________________________")
	fmt.Println("____________________________")
	fmt.Println("ID was: ", deleteTask.ID)
	fmt.Println("____________________________")
	fmt.Println("User ID was: ", deleteTask.UserID)
	fmt.Println("____________________________")
}