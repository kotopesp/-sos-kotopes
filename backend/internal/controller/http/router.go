package http

import (
	"gitflic.ru/spbu-se/sos-kotopes/internal/core"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Router struct {
	app           *fiber.App
	entityService core.EntityService
	authService   interface{}
	postService   core.PostService
}

func NewRouter(
	app *fiber.App,
	entityService core.EntityService,
	authService interface{},
	postService   core.PostService,
) {
	router := &Router{
		app:           app,
		entityService: entityService,
		authService:   authService,
		postService:   postService,
	}

	router.initRequestMiddlewares()

	router.initRoutes()

	router.initResponseMiddlewares()
}

func (r *Router) initRoutes() {
	r.app.Get("/ping", r.ping)

	v1 := r.app.Group("/api/v1")

	// entities
	v1.Get("/entities", r.getEntities)
	v1.Get("/entities/:id", r.getEntityByID)

	// posts
	v1.Get("/posts", r.getPosts)
	v1.Get("/posts/:id", r.getPostByID)
	v1.Get("/posts/:id/photo", r.getPostPhoto)
	v1.Post("/posts", r.createPost)
	v1.Put("/posts/:id", r.updatePost)
	v1.Delete("/posts/:id", r.deletePost)
}

// initRequestMiddlewares initializes all middlewares for http requests
func (r *Router) initRequestMiddlewares() {
	r.app.Use(logger.New())
}

// initResponseMiddlewares initializes all middlewares for http response
func (r *Router) initResponseMiddlewares() {}
