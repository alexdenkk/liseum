package controllers

import "alexdenkk/liseum/db"

// Handler - gateway layer struct
type Handler struct {
	SignKey    []byte
	Repository *db.Repository
}

// New - function for creating new handler
func New(repo *db.Repository, key []byte) *Handler {
	return &Handler{
		Repository: repo,
		SignKey:    key,
	}
}
