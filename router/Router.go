package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"../controller"
)

type Router struct {
}

func NewRouter() *Router {
	return new(Router)
}

func Run() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(context *gin.Context) {
		controller, err := controller.NewTaskController()
		if err != nil {
			context.
		}
		context.HTML(http.StatusOK, "index.html", gin.H {
			tasks: controller.Index(),
		})
	})

	router.NoRoute(func(context *gin.Context) {
	})
}

func ErrorResponse(context *gin.Context) {
	context.
}

