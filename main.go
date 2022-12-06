package main

import (
	"github.com/Raptorik/mywebserver"
	"mvc/app/controllers"
)

func main() {
	artistC := &controllers.ArtistController{}
	artC := &controllers.ArtController{}
	exhibitionC := &controllers.ExhibitionController{}
	mywebserver.StartServer(artistC, artC, exhibitionC)
}
