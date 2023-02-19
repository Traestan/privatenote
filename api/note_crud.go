package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"github.com/traestan/privatenote/internal/model"
)

func (s *Server) create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var response map[string]interface{}
	newNote := s.factory.NewNote(&model.Note{})
	err := json.NewDecoder(r.Body).Decode(&newNote)
	switch {
	case err == io.EOF:
		s.logger.Error().Msg(err.Error())
		s.WriteErrorResponse(w, 500, "Internal Server Error")
		return
	case err != nil:
		s.logger.Error().Msg(err.Error())
		s.WriteErrorResponse(w, 500, "Internal Server Error")
		return
	}
	user := r.Context().Value("Email").(jwt.MapClaims)
	newNote.User = user["email"].(string)

	err = newNote.AddNote()
	if err != nil {
		s.logger.Error().Msg(err.Error())
		s.WriteErrorResponse(w, 500, "Internal Server Error")
		return
	}

	// send response
	response = map[string]interface{}{
		"note": newNote,
	}
	s.WriteOKResponse(w, response)
}

func (s *Server) list(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var response map[string]interface{}

	note := s.factory.NewNote(&model.Note{})
	user := r.Context().Value("Email").(jwt.MapClaims)
	note.User = user["email"].(string)
	list, err := note.ListNotes()

	if err != nil {
		s.logger.Error().Msg(err.Error())
		s.WriteErrorResponse(w, 500, "Internal Server Error")
		return
	}
	// send response
	response = map[string]interface{}{
		"note": list,
	}
	s.WriteOKResponse(w, response)
}
func (s *Server) get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	s.logger.Debug().Msg("Get note")
	var response map[string]interface{}

	note := s.factory.NewNote(&model.Note{})
	numberNote := ps.ByName("shorturl")
	if numberNote == "" {
		s.logger.Error().Str("shorturl", numberNote).Msg("Not found")
		s.WriteErrorResponse(w, 404, "Not found")
		return
	}
	user := r.Context().Value("Email").(jwt.MapClaims)
	note.User = user["email"].(string)

	noteResponse, err := note.GetNote(numberNote)

	if note.User != noteResponse.User {
		s.logger.Error().
			Str("shorturl", numberNote).
			Str("user auth", note.User).
			Str("user note", noteResponse.User).Msg("Not found")

		s.WriteErrorResponse(w, 404, "Not found")
		return
	}
	if err != nil {
		s.logger.Error().Str("shorturl", numberNote).Str("err", err.Error()).Msg("Not found")
		s.WriteErrorResponse(w, 404, "Not found")
		return
	}

	// send response
	response = map[string]interface{}{
		"note": noteResponse,
	}
	s.WriteOKResponse(w, response)
}

func (s *Server) view(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	s.logger.Debug().Msg("View note")
	note := s.factory.NewNote(&model.Note{})

	numberNote := ps.ByName("shorturl")

	if numberNote == "" {
		s.logger.Error().Str("shorturl", numberNote).Msg("Not found")
		s.WriteErrorResponse(w, 404, "Not found")
		return
	}
	noteResponse, err := note.GetNote(numberNote)

	if err != nil {
		s.logger.Error().Str("shorturl", numberNote).Str("err", err.Error()).Msg("Not found")
		s.WriteErrorResponse(w, 404, "Not found")
		return
	}

	s.WriteTextResponse(w, noteResponse)
}

func (s *Server) edit(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	s.logger.Debug().Msg("Edit note")
	var response map[string]interface{}

	noteEdit := s.factory.NewNote(&model.Note{})

	err := json.NewDecoder(r.Body).Decode(&noteEdit)
	switch {
	case err == io.EOF:
		s.logger.Error().Msg(err.Error())
		s.WriteErrorResponse(w, 500, "Internal Server Error")
		return
	case err != nil:
		s.logger.Error().Msg(err.Error())
		s.WriteErrorResponse(w, 500, "Internal Server Error")
		return
	}
	user := r.Context().Value("Email").(jwt.MapClaims)
	noteEdit.User = user["email"].(string)
	noteEdit.Number = ps.ByName("shorturl")

	noteResponse, err := noteEdit.EditNote()

	if err != nil {
		s.logger.Error().Str("shorturl", noteEdit.Number).Str("err", err.Error()).Msg("Not found")
		s.WriteErrorResponse(w, 404, "Not found")
		return
	}

	// send response
	response = map[string]interface{}{
		"note": noteResponse,
	}
	s.WriteOKResponse(w, response)
}
