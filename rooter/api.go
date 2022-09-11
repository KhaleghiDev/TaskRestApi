package rooter

import (
	controller "RestApi/controller"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Api(Router *mux.Router, db *gorm.DB) {
	api := Router.PathPrefix("/api/v1").Subrouter()
	//api.Methods("POST").Path("/operation:(?:register|login)").HandlerFunc(controller.AuthController())
	api.Methods("POST").Path("/register").HandlerFunc(controller.Register())
	api.Methods("POST").Path("/login").HandlerFunc(controller.Login())
	api.Methods("GET").Path("/allTicket").HandlerFunc(controller.AllTicket())
	api.Methods("POst").Path("/createTicket").HandlerFunc(controller.CreateTicket())
	api.Methods("Post").Path("/updateTicket").HandlerFunc(controller.UpdateTicket())
}
