package routes

import (
	"new_go_app/src/controllers"
	"new_go_app/src/middleware"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	api := app.Group("api")
	user := api.Group("user")

	user.Post("register", controllers.Register)
	user.Post("login", controllers.Login)
	user.Get("me", controllers.Me)

	loginUser := api.Use(middleware.IsAuthenticated)
	loginUser.Post("create_post", controllers.CreatePost)
	loginUser.Delete("delete_post/:id", controllers.DeletePost)
	loginUser.Post("comment/:id", controllers.CreateComment)
	api.Get("posts", controllers.Posts)
	api.Get("post/:id", controllers.GetPost)

	loginUser.Post("vote", controllers.Vote)
}
