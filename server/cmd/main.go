package main

import (
	"log"

	"github.com/IlnarAhm/notebook-app/db"
	"github.com/IlnarAhm/notebook-app/internal/user"
	"github.com/IlnarAhm/notebook-app/router"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("Could not initialized database connection: %s", err)
	}

	userRep := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRep)
	userHanlder := user.NewHandler(userSvc)

	router.InitRouter(userHanlder)
	router.Start("0.0.0.0:8080")
}
