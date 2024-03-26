package handler

import (
	"RADserver/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func NotesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getNotes(w,r)

		w.Header().Set("Content-Type","application/json")
		fmt.Println("\n'GET'-response sent to /notes on", time.Now().Format(time.RFC850))
	case "POST":
		createNote(w,r)

		reqMessage := " 'POST'-response sent"
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"request": "%s"}`, reqMessage)
		fmt.Println("\n'POST'-response sent to /notes on", time.Now().Format(time.RFC850))
	case "PUT":
		updateNote(w,r)

		reqMessage := " 'PUT'-response sent"
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"request": "%s"}`, reqMessage)
		fmt.Println("\n'PUT'-response sent to /notes on", time.Now().Format(time.RFC850))
	case "DELETE":
		deleteNote(w,r)

		reqMessage := " 'DELETE'-response sent"
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"request": "%s"}`, reqMessage)
		fmt.Println("\n'DELETE'-response sent to /notes on", time.Now().Format(time.RFC850))
	}
}

func getNotes(w http.ResponseWriter, r *http.Request) {
	var notesData []models.Notes
	byteData, _ := os.ReadFile("db/notes.json")
	json.Unmarshal(byteData, &notesData) 

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(notesData)
}
func createNote(w http.ResponseWriter, r *http.Request) {
	var newNote models.Notes
	json.NewDecoder(r.Body).Decode(&newNote)

	var notesData []models.Notes
	byteData, _ := os.ReadFile("db/notes.json")
	json.Unmarshal(byteData, &notesData)

	var usersData []models.User
	userByte, _ := os.ReadFile("db/users.json")
	json.Unmarshal(userByte, &usersData)

	var userFound bool

	for i := 0; i < len(usersData); i++ {
		if usersData[i].Id == newNote.UserID {
			newNote.ID = len(notesData)+1
			newNote.CreatedTime = time.Now().Format(time.RFC850)
			newNote.UpdatedTime = time.Now().Format(time.RFC850)
			userFound = true
			break
		}
	}
	if !userFound{
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "no user found with this UserID")
		return
	}

	notesData = append(notesData, newNote)

	res,_ := json.Marshal(notesData)
	os.WriteFile("db/notes.json", res, 0)

	w.WriteHeader(http.StatusOK)
	fmt.Println("\n_____________________________________________")
	fmt.Println("Created new note at", time.Now().Format(time.RFC850))
	fmt.Println("_____________________________________________")
	fmt.Println("____________________________")
	fmt.Println("ID is: ", newNote.ID)
	fmt.Println("____________________________")
	fmt.Println("User ID is: ", newNote.UserID)
	fmt.Println("____________________________")
}
func updateNote(w http.ResponseWriter, r *http.Request) {
	var updateNote models.Notes
	json.NewDecoder(r.Body).Decode(&updateNote)

	var notesData []models.Notes
	byteData,_ := os.ReadFile("db/notes.json")
	json.Unmarshal(byteData, &notesData)

	var usersData []models.User
	userByte,_ := os.ReadFile("db/users.json")
	json.Unmarshal(userByte, &usersData)

	var userFound bool
	var noteFound bool

	for i := 0; i < len(usersData); i++ {
		if usersData[i].Id == updateNote.UserID {
			for j := 0; j < len(notesData); j++ {
				if notesData[j].ID == updateNote.ID{
					notesData[j].Title = updateNote.Title
					notesData[j].Content = updateNote.Content
					notesData[j].UpdatedTime = time.Now().Format(time.RFC850)
					noteFound = true
					break 
				} 
			}
			if !noteFound {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "Note with such an ID not found.")
				return
			}
			userFound = true
			break
		}
 	}
	if !userFound {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Note with such an UserID not found.")
		return
	}

	res, _ := json.Marshal(notesData)
	os.WriteFile("db/notes.json", res, 0)

	w.WriteHeader(http.StatusOK)
	fmt.Println("\n_____________________________________________")
	fmt.Println("Updated note at", time.Now().Format(time.RFC850))
	fmt.Println("_____________________________________________")
	fmt.Println("____________________________")
	fmt.Println("ID is: ", updateNote.ID)
	fmt.Println("____________________________")
	fmt.Println("User ID is: ", updateNote.UserID)
	fmt.Println("____________________________")
}
func deleteNote(w http.ResponseWriter, r *http.Request) {
	var deleteNote models.Notes
	json.NewDecoder(r.Body).Decode(&deleteNote)

	var notesData []models.Notes
	byteData,_ := os.ReadFile("db/notes.json")
	json.Unmarshal(byteData, &notesData)

	var usersData []models.User
	userByte,_ := os.ReadFile("db/users.json")
	json.Unmarshal(userByte, &usersData)

	var userFound bool
	var noteFound bool

	for i := 0; i < len(usersData); i++ {
		if usersData[i].Id == deleteNote.UserID {
			for j := 0; j < len(notesData); j++ {
				if notesData[j].ID == deleteNote.ID {
					notesData = append(notesData[:j], notesData[j+1:]...)
					noteFound = true
					break
				} 
			}
			if !noteFound {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w,"Note with such kind of ID not found.")
				return
			}
			userFound = true
			break
		}
	}
	if !userFound {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Note with such kind of UserID not found.")
		return
	}

	res,_ := json.Marshal(notesData)
	os.WriteFile("db/notes.json", res, 0)

	w.WriteHeader(http.StatusOK)
	fmt.Println("\n_____________________________________________")
	fmt.Println("Deleted note at", time.Now().Format(time.RFC850))
	fmt.Println("_____________________________________________")
	fmt.Println("____________________________")
	fmt.Println("ID was: ", deleteNote.ID)
	fmt.Println("____________________________")
	fmt.Println("User ID was: ", deleteNote.UserID)
	fmt.Println("____________________________")
}
