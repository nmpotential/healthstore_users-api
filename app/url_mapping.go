package app

import (
	"github.com/nmpotential/career-development/projects/healthstore/healthstore_users-api/controllers/ping"
	"github.com/nmpotential/career-development/projects/healthstore/healthstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:users_id", users.GetUser)
	router.GET("/users/search", users.SearchUser)
	router.POST("/users", users.CreateUser)
}
