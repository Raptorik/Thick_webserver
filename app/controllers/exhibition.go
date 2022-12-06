package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"mvc/app/models"
	"net/http"
)

type ExhibitionController struct {
	exhibitions []*models.Exhibition
	router      *mux.Router
}

func (ec *ExhibitionController) RegisterRouter(r *mux.Router) {
	ec.router = r
}

func (ec *ExhibitionController) RegisterActions() {
	ec.router.HandleFunc("/createexhibition]/{exhibition}", ec.ExhibitionCreation)
	ec.router.HandleFunc("/artist/delete/{artist}/{exhibition}", ec.KickArtistOffExhibition)
}

func (ec *ExhibitionController) OrganizeExhibition(e *models.Exhibition) {
	ec.exhibitions = append(ec.exhibitions, e)
}

func (gc *ExhibitionController) FindExhibition(name string) *models.Exhibition {
	for _, e := range gc.exhibitions {
		if name == e.Name {
			return e
		}
	}
	return nil
}

func (ec *ExhibitionController) AddArtist(exhibition *models.Exhibition, artist *models.Artist) {
	if len(artist.Arts) > 0 {
		ec.exhibitions = append(ec.exhibitions, exhibition)
		exhibition.Artists = append(exhibition.Artists, artist)
		artist.AtExhibition = true
		return
	}
	if len(artist.Arts) == 0 {
		fmt.Println("We can't register an Artist without arts")
	}
	if artist.AtExhibition {
		fmt.Println("Artist are already at the exhibition")
	}
}
func (ec *ExhibitionController) DeleteArtist(exhibition *models.Exhibition, artist *models.Artist) {
	for _, exhibition := range ec.exhibitions {
		if exhibition == exhibition {
			exhibition.DeleteArtist(artist)
		}
	}
}
func (ec *ExhibitionController) ExhibitionCreation(rw http.ResponseWriter, r *http.Request) {
	var vars map[string]string = mux.Vars(r)
	var exhibitionName string = vars["exhibition"]
	resp := make(map[string]string)
	resp["message"] = `Exhibition ` + exhibitionName + ` created successfully`
	exhibition := &models.Exhibition{Name: exhibitionName}
	ec.OrganizeExhibition(exhibition)
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	rw.Write(jsonResp)
}

func (ec *ExhibitionController) KickArtistOffExhibition(rw http.ResponseWriter, r *http.Request) {
	var vars map[string]string = mux.Vars(r)
	var artistName string = vars["artist"]
	var exhibitionName string = vars["exhibition"]
	artistC := &ArtistController{}
	artist := artistC.FindArtist(artistName)
	if err := artistC.FindArtist(artistName); err != nil {
		exhibition := ec.FindExhibition(exhibitionName)
		ec.DeleteArtist(exhibition, artist)
	}
	resp := make(map[string]string)
	resp["message"] = `Artist:` + artistName + `is deleted from the Exhibition:` + exhibitionName
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	rw.Write(jsonResp)
}
