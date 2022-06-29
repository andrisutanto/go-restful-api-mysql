package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func setupRoutesForVideoGames(router *mux.Router) {
	enableCORS(router)

	router.HandleFunc("/videogames", func(w http.ResponseWriter, r *http.Request) {
		videoGames, err := getVideoGames()
		if err == nil {
			respondWithSuccess(videoGames, w)
		} else {
			respondWithError(err, w)
		}
	}).Methods(http.MethodGet)
	router.HandleFunc("/videogame/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := stringToInt64(idAsString)
		if err != nil {
			respondWithError(err, w)
			return
		}
		videogame, err := getVideoGameById(id)
		if err != nil {
			respondWithError(err, w)
		} else {
			respondWithSuccess(videogame, w)
		}
	}).Methods(http.MethodGet)

	router.HandleFunc("/videogame", func(w http.ResponseWriter, r *http.Request) {
		var videoGame VideoGame
		err := json.NewDecoder(r.Body).Decode(&videoGame)
		if err != nil {
			respondWithError(err, w)
		} else {
			err := createVideoGame(videoGame)
			if err != nil {
				respondWithError(err, w)
			} else {
				respondWithSuccess(true, w)
			}
		}
	}).Methods(http.MethodPost)

	router.HandleFunc("/videogame", func(w http.ResponseWriter, r *http.Request) {
		var videoGame VideoGame
		err := json.NewDecoder(r.Body).Decode(&videoGame)
		if err != nil {
			respondWithError(err, w)
		} else {
			err := updateVideoGame(videoGame)
			if err != nil {
				respondWithError(err, w)
			} else {
				respondWithSuccess(true, w)
			}
		}
	}).Methods(http.MethodPut)
	router.HandleFunc("/videogame/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := stringToInt64(idAsString)
		if err != nil {
			respondWithError(err, w)
			return
		}
		err = deleteVideoGame(id)
		if err != nil {
			respondWithError(err, w)
		} else {
			respondWithSuccess(true, w)
		}
	}).Methods(http.MethodDelete)
}