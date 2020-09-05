package controller

import (
    "fmt"
    "net/http"
    "encoding/json"
    "app/model"
    mux "github.com/gorilla/mux"
    "app/database"
)

func GetUserById(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json") 

    vars := mux.Vars(r)
    id   := vars["id"]

    var user model.User
    
    db := database.Conn.DB

    result := db.First(&user, id)

    if result.Error != nil {
        fmt.Fprintf(w, "Usuario não localizado!")
        return
    } 

    json.NewEncoder(w).Encode(user)
    
}

func CreateUser(w http.ResponseWriter, r *http.Request){
    var user model.User
    decoder := json.NewDecoder(r.Body)
    decoder.Decode(&user)

    db := database.Conn.DB

    result := db.Create(&user)

    if result.Error != nil {
        fmt.Fprintf(w, "Ocorreu um errro ao cadastro o usuarios")
        return
    } 

    fmt.Fprintf(w, "Usuario cadastro com sucesso!")

}

func ListUsers (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json") 

    var users [] model.User

    db := database.Conn.DB

    result := db.Find(&users)

    if result.Error != nil{
        fmt.Fprintf(w,"Ocorreu um errro ao listar os usuarios")
        return
    }  

    json.NewEncoder(w).Encode(users)
}

func UpdateUser(w http.ResponseWriter, r *http.Request){
    var user model.User
    decoder := json.NewDecoder(r.Body)
    decoder.Decode(&user)

    db := database.Conn.DB
    result := db.Save(&user)

    if result.Error != nil {
        fmt.Fprintf(w, "Ocorreu um errro ao atualizar o usuarios")
        return
    } 

    fmt.Fprintf(w, "Usuario atualizado com sucesso!")

}

func DeleteUser(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    id := vars["id"]

    db := database.Conn.DB

    result := db.Delete(&model.User{}, id)

    if result.Error != nil {
        fmt.Fprintf(w,"Ocorreu um errro ao excluir o usuarios")
        return
    } 

    fmt.Fprintf(w, "Usuario deletado com sucesso!")

}
