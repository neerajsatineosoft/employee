package controllers

import (
	"net/http"

	"github.com/neerajsatineosoft/neostore/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To Neostore")

}
