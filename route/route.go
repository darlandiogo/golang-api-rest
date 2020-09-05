package route

import(
	"log"
	"net/http"
	"app/controller"
	"github.com/gorilla/mux"
)

func HandleRoute(){
	route := mux.NewRouter().StrictSlash(true)
	route.HandleFunc("/", controller.Home).Methods("GET")
	route.HandleFunc("/user", controller.ListUsers).Methods("GET")
	route.HandleFunc("/user/{id}", controller.GetUserById).Methods("GET")
	route.HandleFunc("/user", controller.UpdateUser).Methods("PUT")
	route.HandleFunc("/user", controller.CreateUser).Methods("POST")
	route.HandleFunc("/user/{id}", controller.DeleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", route))
}