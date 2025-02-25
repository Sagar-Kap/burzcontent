/*
Package routes defines the application's routing configuration.

This package is responsible for setting up the routes and their associated handlers
for various HTTP endpoints. The routes are mapped to functions defined in the
handlers package, which process incoming requests and generate appropriate responses.

The main function in this package, `SetupRoutes`, configures the application's
routes and binds them to specific handlers for resource management, such as
user-related routes.
*/
package routes

import (
	"github.com/go-chi/chi/v5"

	"github.com/Weburz/burzcontent/server/internal/api/handlers"
)

/*
SetupRoutes sets up the application's HTTP routes and maps them to their corresponding
handlers.

This function performs the following steps:

 1. Configures the `/users` route using `r.Route("/users", ...)` for handling
    user-related requests.
 2. Binds the `GET` method for the `/users` route to the `GetUsers` method of
    the `UserHandler` defined in the `handlers.Handlers` instance.

The routes are now ready to process incoming requests related to users.
*/
func SetupRoutes(r *chi.Mux, h *handlers.Handlers) {
	// Mount all handlers related to the users
	r.Route("/users", func(r chi.Router) {
		r.Get("/", h.UserHandler.GetUsers)
		r.Put("/new", h.UserHandler.CreateUser)
		r.Get("/{id}", h.UserHandler.GetUser)
		r.Post("/{id}/edit", h.UserHandler.EditUser)
		r.Delete("/{id}/delete", h.UserHandler.DeleteUser)
	})

	// Mount all handlers related to the articles
	r.Route("/articles", func(r chi.Router) {
		r.Get("/", h.ArticleHandler.GetArticles)
		r.Put("/new", h.ArticleHandler.CreateArticle)
		r.Get("/{id}", h.ArticleHandler.GetArticle)
		r.Post("/{id}/edit", h.ArticleHandler.EditArticle)
		r.Delete("/{id}/delete", h.ArticleHandler.DeleteArticle)
	})
}
