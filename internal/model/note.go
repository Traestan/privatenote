package model

import (
	"errors"
	"time"
)

type Note struct {
	modelData
	User        string `json:"email"`
	Number      string `json:"number"`
	Text        string `json:"text"`
	Ttl         string `json:"ttl"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// newNote model note
func newNote(data modelData, note *Note) *Note {
	if note == nil {
		note = &Note{}
	}
	note.modelData = data
	return note
}

func (n *Note) AddNote() error {
	n.logger.Debug().Msg("Add note")

	number := n.Text
	n.Number = n.encoder.EncodeURL(n.encoder.DecodeURL(number))

	timeTtl, err := time.ParseDuration(n.Ttl)
	if err != nil {
		return err
	}
	newNote := map[string]interface{}{
		"title":       n.Title,
		"description": n.Description,
		"text":        n.Text,
		"ttl":         n.Ttl,
		"user":        n.User,
		"number":      n.Number,
	}
	err = n.repository.HMSet(n.Number, newNote)
	if err != nil {
		return err
	}

	status := n.repository.Expire(n.Number, timeTtl)
	if status != nil {
		n.logger.Debug().Str("Status ttl", status.Error()).Msg("Add note")
	}
	//добавляем связь юзер заметка
	newUserNote := map[string]interface{}{
		n.Number: n.Ttl,
	}

	err = n.repository.HMSet(n.User, newUserNote)
	if err != nil {
		n.logger.Error().Str("err", err.Error()).Msg("Add note hmset user")
		return err
	}

	return nil
}
func (n *Note) GetNote(shorturl string) (Note, error) {
	n.logger.Debug().Msg("Get note")
	note := Note{}

	result, err := n.repository.GetAll(shorturl)
	if err != nil {
		n.logger.Error().Str("err", err.Error()).Msg("List note")
		return note, err
	}

	note = Note{
		Title:       result["title"],
		Text:        result["text"],
		Description: result["description"],
		Ttl:         result["ttl"],
		Number:      shorturl,
		User:        result["user"],
	}

	return note, nil
}
func (n *Note) ListNotes() ([]Note, error) {
	n.logger.Debug().Msg("List note")
	var notes []Note
	userNotes, err := n.repository.GetAll(n.User)
	if err != nil {
		n.logger.Error().Str("err", err.Error()).Msg("List note")
		return nil, err
	}
	for k, _ := range userNotes {
		result, err := n.repository.GetAll(k)
		if err != nil {
			n.logger.Error().Str("err", err.Error()).Msg("List note")
			return nil, err
		}
		note := Note{
			Title:       result["title"],
			Text:        result["text"],
			Description: result["description"],
			Ttl:         result["ttl"],
			Number:      k,
		}
		notes = append(notes, note)
	}

	return notes, nil
}
func (n *Note) EditNote() (*Note, error) {
	n.logger.Debug().Msg("Get note")

	note := &Note{}

	result, err := n.repository.GetAll(n.Number)
	if err != nil {
		n.logger.Error().Str("err", err.Error()).Msg("List note")
		return note, err
	}

	note = &Note{
		Title:       result["title"],
		Text:        result["text"],
		Description: result["description"],
		Ttl:         result["ttl"],
		Number:      n.Number,
		User:        result["user"],
	}

	if note.User == n.User {
		changeTitle := n.repository.HSet(n.Number, "title", n.Title)
		changeDescription := n.repository.HSet(n.Number, "description", n.Description)
		changeText := n.repository.HSet(n.Number, "text", n.Text)
		changeTtl := n.repository.HSet(n.Number, "ttl", n.Ttl)

		n.logger.Debug().
			Bool("title", changeTitle).
			Bool("description", changeDescription).
			Bool("text", changeText).
			Bool("ttl", changeTtl).
			Msg("Changes")

		timeTtl, err := time.ParseDuration(n.Ttl)
		if err != nil {
			n.logger.Error().
				Str("ttl set", n.Ttl).
				Str("get error", err.Error()).Msg("Change note")
			return nil, err
		}

		status := n.repository.Expire(n.Number, timeTtl)
		if status != nil {
			n.logger.Debug().Str("Status ttl", status.Error()).Msg("Add note")
		}
		n.logger.Debug().
			Bool("title", changeTitle).
			Bool("description", changeDescription).
			Bool("text", changeText).
			Bool("ttl", changeTtl).
			Msg("Changes error")
	} else {
		n.logger.Error().
			Str("user auth", n.User).
			Str("user note", note.User).Msg("Not found")

		return nil, errors.New("Not found")
	}
	return note, nil
}
