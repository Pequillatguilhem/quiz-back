package main

import (
	"Quiz-back/model"
	"Quiz-back/repo"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	portInt, err := strconv.Atoi(port)
	if err != nil {
		panic(err)
	}
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_DBNAME")

	serieRepo, err := repo.NewSerie(host, portInt, user, password, dbname)
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	router.GET("/series", func(c *gin.Context) {
		series, err := serieRepo.SelectSeries()
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, series)
	})
	router.GET("/questions/:serieId", func(c *gin.Context) {
		serieId := c.Param("serieId")
		sid, err := strconv.Atoi(serieId)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		series, err := serieRepo.SelectQuestions(sid)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, series)
	})

	router.POST("/question", func(c *gin.Context) {
		var question model.Question
		err := c.BindJSON(&question)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		err = serieRepo.InsertQuestion(question.Name, question.Response, question.SerieId)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.String(http.StatusOK, "created")
	})

	router.POST("/serie", func(c *gin.Context) {
		var serie model.Serie
		err := c.BindJSON(&serie)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		err = serieRepo.InsertSerie(serie.Name)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.String(http.StatusOK, "created")
	})

	router.DELETE("/question/:id", func(c *gin.Context) {
		id := c.Param("id")
		sid, err := strconv.Atoi(id)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		err = serieRepo.DeleteQuestion(sid)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.String(http.StatusOK, "deleted")
	})

	router.DELETE("/serie/:id", func(c *gin.Context) {
		id := c.Param("id")
		sid, err := strconv.Atoi(id)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		err = serieRepo.DeleteSerie(sid)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.String(http.StatusOK, "deleted")
	})
	router.Run(":8079")
}
