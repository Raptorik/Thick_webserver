package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"mvc/app/models"
	"net/http"
)

type ArtistController struct {
	artists []*models.Artist
	router  *mux.Router
}

func (ac *ArtistController) RegisterRouter(r *mux.Router) {
	ac.router = r
}

func (ac *ArtistController) RegisterActions() {
	ac.router.HandleFunc("/createartist/{artist}", ac.Registration)
	ac.router.HandleFunc("/artist/register/{artist}/{exhibition}", ac.ArtistRegistration)
}

func (ac *ArtistController) InviteArtist(a *models.Artist) {
	ac.artists = append(ac.artists, a)
}

func (ac *ArtistController) FindArtist(name string) *models.Artist {
	for _, artist := range ac.artists {
		if artist.Name == name {
			return artist
		}
	}
	return nil
}
func (ac *ArtistController) Registration(rw http.ResponseWriter, r *http.Request) {
	var vars map[string]string = mux.Vars(r)
	var artistName string = vars["artist"]
	resp := make(map[string]string)
	resp["message"] = `Artist ` + artistName + ` created successfully`
	artist := &models.Artist{Name: artistName, AtExhibition: false}
	ac.InviteArtist(artist)
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	rw.Write(jsonResp)
}

func (ac *ArtistController) ArtistRegistration(rw http.ResponseWriter, r *http.Request) {
	var vars map[string]string = mux.Vars(r)
	var artistName string = vars["artist"]
	var exhibitionName string = vars["exhibition"]
	artist := ac.FindArtist(artistName)
	if err := ac.FindArtist(artistName); err != nil {
		exhibitionC := &ExhibitionController{}
		exhibition := exhibitionC.FindExhibition(exhibitionName)
		exhibitionC.AddArtist(exhibition, artist)
	}
	resp := make(map[string]string)
	resp["message"] = `Artist: ` + artistName + `is registered on the Exhibition:` + exhibitionName
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	rw.Write(jsonResp)
}
