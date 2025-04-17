package api

import (
	db "github.com/a-paladini/bau-do-aventureiro/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/weapons", server.createWeapon)
	router.GET("/weapons/:id", server.getWeapon)
	router.GET("/weapons", server.listAllWeapons)
	router.GET("/weapons/category/:type", server.listWeaponsByCategory)
	router.PUT("/weapons/:id", server.updateWeapon)
	router.DELETE("/weapons/:id", server.deleteWeapon)

	router.POST("/armours", server.createArmour)
	router.GET("/armours/:id", server.getArmour)
	router.GET("/armours", server.listAllArmour)
	router.GET("/armours/category/:category", server.listArmoursByCategory)
	router.PUT("/armours/:id", server.updateArmour)
	router.DELETE("/armours/:id", server.deleteArmour)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"Error": err.Error()}
}
