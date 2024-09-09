package handler

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	server := New()

	server.get("/", func(context *Context) {
		context.JSON(200, H{
			"message": "hello from vercel",
		})
	})

	server.get("/name", func(context *Context) {
		name := context.Query("name")
		if name == "" {
			context.JSON(400, H{
				"message": "name not found",
			})
		} else {
			context.JSON(200, H{
				"message": fmt.Sprintf("hello from %s", name),
			})
		}
	})
}
