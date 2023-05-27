package main

import (
	"log"
	"net/http"

	"github.com/Epritka/gokit/example/internal/core/entity"
	"github.com/Epritka/gokit/example/internal/core/usecase/user"
	"github.com/Epritka/gokit/wrapper"
	"github.com/gin-gonic/gin"
)

type handlers struct {
	userInterceptor *user.Interceptor
}

func New(userInterceptor *user.Interceptor) *handlers {
	return &handlers{
		userInterceptor: userInterceptor,
	}
}

func (h *handlers) Create(c *gin.Context) {
	request := entity.User{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(wrapper.FailedHttpResponse(err))
		return
	}

	err := h.userInterceptor.Create(&request)
	if err != nil {
		c.AbortWithStatusJSON(wrapper.FailedHttpResponse(err))
		return
	}

	c.Status(200)
}

func main() {
	userInterceptor := user.New()

	gin.SetMode(gin.DebugMode)
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	userHandlers := New(userInterceptor)

	example := router.Group("/example/")
	{
		user := example.Group("/user/")
		{
			user.POST("/", userHandlers.Create)
		}
	}

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
