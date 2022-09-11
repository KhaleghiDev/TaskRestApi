package rooter

import (
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

// Rooter  call http and create api restFul
func Rooter(address string, db *gorm.DB) {
	router := mux.NewRouter()
	//start router api
	Api(router, db)

	srv := &http.Server{
		Handler: router,
		Addr:    address,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("server", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
