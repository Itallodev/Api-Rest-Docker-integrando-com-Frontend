package routes

import (
	"log"
	"net/http"

	"github.com/devitallo/api-go/controllers"
	"github.com/devitallo/api-go/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// UTILIZAMOS O GORILLA MUX PARA CONTROLAR AS ROTAS INVÉS DE USAR HTTP
func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPesonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CriaPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletaPersonalidade).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controllers.AtualizaPersonalidade).Methods("Put")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))

}
