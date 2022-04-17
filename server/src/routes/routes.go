package routes

import (
	"new_go_app/server/src/controllers"
	"new_go_app/server/src/middleware"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	api := app.Group("api")
	api.Get("posts", controllers.Posts)
	api.Get("post_index/:page", controllers.PostIndex)
	api.Get("post/:id", controllers.GetPost)
	api.Get("comment/:id", controllers.GetReply)
	api.Get("tags", controllers.Tags)
	api.Get("users/:id", controllers.GetUser)

	user := api.Group("user")
	user.Post("register", controllers.Register)
	user.Post("login", controllers.Login)
	user.Get("logout", controllers.Logout)
	user.Get("me", controllers.Me)
	user.Put("bio/:id", controllers.ChangeBio)

	loginUser := api.Use(middleware.IsAuthenticated)
	loginUser.Post("create_post", controllers.CreatePost)
	loginUser.Delete("delete_post/:id", controllers.DeletePost)
	loginUser.Post("comment/:id", controllers.CreateComment)
	loginUser.Post("create_tag", controllers.CreateTag)
	loginUser.Put("add_tag/:id/tags/:tag_id", controllers.AddTags)

	loginUser.Post("vote", controllers.Vote)
}
