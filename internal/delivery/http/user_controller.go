package http

import (
	"Hermes/internal/delivery/http/middleware"
	"Hermes/internal/delivery/http/response"
	"Hermes/internal/domain"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	usr_usecase domain.UserUseCase
}

func RegisterUserRoutes(g *gin.Engine, u_usecase domain.UserUseCase) {
	log.Print("Registering user routes")
	usr_controller := UserController{
		usr_usecase: u_usecase,
	}
	usr_group := g.Group("/users")

	usr_group.Use(middleware.JWTMiddleware())

	usr_group.GET("", usr_controller.GetAll)
	usr_group.GET("/:id", usr_controller.GetByID)
	usr_group.POST("", usr_controller.Create)
}

func (usr_contr *UserController) GetAll(g *gin.Context) {
	usrs, err := usr_contr.usr_usecase.GetAll(g.Request.Context())
	if err != nil {
		g.JSON(http.StatusInternalServerError, response.Failure("Unable to get users.", ""))
		return
	}

	g.JSON(http.StatusOK, response.Success(usrs, "Users found."))
}

func (usr_contr *UserController) Create(g *gin.Context) {

}

func (u *UserController) GetByID(g *gin.Context) {
	usr_id := strings.Trim(g.Param("id"), " ")
	log.Print(usr_id)

	if usr_id == "" {
		g.JSON(http.StatusBadRequest, response.Failure("User id should not be empty", ""))
		return
	}

	var id domain.ID
	id, err := id.Convert(usr_id)
	if err != nil {
		log.Print("Error during user ID parsing: ", err)
		g.JSON(http.StatusBadRequest, response.Failure("", "Unable to parse user ID."))
		return
	}

	user, err := u.usr_usecase.GetById(g.Request.Context(), id)
	if err != nil {
		log.Print("Error during user search: ", err)
		g.JSON(http.StatusBadRequest, response.Failure("", "Unable to find user by id."))
		return
	}
	g.JSON(http.StatusOK, response.Success(user, "User found."))
}
