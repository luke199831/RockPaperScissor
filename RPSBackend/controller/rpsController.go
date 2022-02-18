package controller

import (
	"RPSBackend/model"
	"RPSBackend/service"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var rpsService = service.RPSService{}

func Run(port string) {

	ginEngine := gin.Default()

	ginEngine.Use(cors.Default())

	router := ginEngine.Group("/rps")
	router.GET("/scoreboard", scoreBoard)
	router.POST("/user", createUser)
	router.PUT("/:id/:rps", playGame)

	ginEngine.Run(port)

}

func scoreBoard(context *gin.Context) {

	score, err := rpsService.ScoreBoard()

	if err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
	} else {
		context.JSON(http.StatusOK, score)
	}

}

func createUser(context *gin.Context) {

	var user model.User

	if err := context.BindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		context.Abort()
		return
	}

	if id, err := rpsService.CreateUser(user); err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
	} else {
		message := "user registered with id: " + strconv.Itoa(id)
		context.JSON(http.StatusOK, message)
	}

}

func playGame(context *gin.Context) {

	id, convErr := strconv.Atoi(context.Param("id"))
	rps := context.Param("rps")

	if convErr != nil {
		context.JSON(http.StatusBadRequest, convErr.Error())
	}

	if score, err := rpsService.UpdateScore(id, rps); err != nil {
		context.JSON(http.StatusInternalServerError, err.Error())
	} else {
		message := "score:" + strconv.Itoa(score)
		context.JSON(http.StatusOK, message)
	}

}
