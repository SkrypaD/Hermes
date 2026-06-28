package http

import (
	"Hermes/internal/delivery/http/middleware"
	"Hermes/internal/delivery/http/response"
	"database/sql"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Db *sql.DB
}

func RegisterUserRoutes(g *gin.Engine, db *sql.DB) {
	usr_controller := UserController{Db: db}
	usr_group := g.Group("/users")

	usr_group.Use(middleware.JWTMiddleware())

	usr_group.GET("", usr_controller.GetAll)
	usr_group.GET("/:id", usr_controller.GetByID)
	usr_group.POST("", usr_controller.Create)
}

func (usr_contr *UserController) GetAll(g *gin.Context) {

}

func (usr_contr *UserController) Create(g *gin.Context) {

}

func (usr_contr *UserController) GetByID(g *gin.Context) {
	urs_id := strings.Trim(g.Param("id"), " ")

	if urs_id == "" {
		g.JSON(http.StatusBadRequest, response.Failure("User id should not be empty", ""))
	}
}
