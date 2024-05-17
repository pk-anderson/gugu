package routes

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router, db *sql.DB) {
	SetupUserRoutes(router, db)
	SetupAccessRoutes(router, db)
}
