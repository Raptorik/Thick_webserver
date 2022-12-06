package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"mvc/app/models"
	"net/http"
)

type ArtController struct {
	arts   []*models.Art
	router *mux.Router
}

func (ac *ArtController) RegisterRouter(r *mux.Router) {
	ac.router = r
}

func (ac *ArtController) RegisterActions() {
	ac.router.HandleFunc("/createart/{art}", ac.ArtCreation)
	ac.router.HandleFunc("/artist/assign/{artist}/{art}", ac.AssignArt)
}

func (ac *ArtController) CreateArt(a *models.Art) {
	ac.arts = append(ac.arts, a)
}

func (ac *ArtController) FindArt(name string) *models.Art {
	for _, art := range ac.arts {
		if art.Name == name {
			return art
		}
	}
	return nil
}

func (ac *ArtController) AssignedArtToArtist(art *models.Art, artist *models.Artist) *models.Art {
	if art.IsNotAssigned() {
		art.Owner = artist.Name
		artist.Arts = append(artist.Arts, art)
		fmt.Println(artist.Name)
		return art
	} else {
		fmt.Println("This art already has an owner! You'd better make your own art")
	}
	return nil
}

func (ac *ArtController) ArtCreation(rw http.ResponseWriter, r *http.Request) {
	var vars map[string]string = mux.Vars(r)
	var artName string = vars["art"]
	resp := make(map[string]string)
	resp["message"] = `Art ` + artName + ` created successfully`
	art := &models.Art{Name: artName}
	ac.CreateArt(art)
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	rw.Write(jsonResp)
}

func (ac *ArtController) AssignArt(rw http.ResponseWriter, r *http.Request) {
	var vars map[string]string = mux.Vars(r)
	var artistName string = vars["artist"]
	var artName string = vars["art"]
	art := ac.FindArt(artName)
	if err := ac.FindArt(artName); err != nil {
		artistC := &ArtistController{}
		artist := artistC.FindArtist(artistName)
		if err := artistC.FindArtist(artistName); err != nil {
			ac.AssignedArtToArtist(art, artist)
		}
	}
	resp := make(map[string]string)
	resp["message"] = `Art: ` + artName + ` is assigned to Artist:` + artistName
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	rw.Write(jsonResp)
}
