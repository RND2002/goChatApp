package main

import (
	"fmt"
	"github.com/RND2002/goChatApp/db"
	_ "github.com/RND2002/goChatApp/docs"
	"github.com/RND2002/goChatApp/routers"
	"github.com/RND2002/goChatApp/ws"
)

func main() {
	fmt.Println("Hello to Chat App")
	db.InitializeDB()

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	routers.SetupRoutes(wsHandler)

}
