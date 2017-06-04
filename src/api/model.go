package main

import (
	"database/sql"
	"errors"
)

// Heritage represents an Heritage object
type Heritage struct {
	ID      int    `json:"id"`
	PlaceId int    `json:"place_id"`
	File    string `json:"file"`
	Name    string `json:"name"`
	Class   string `json:"class"`
	Status  string `json:"status"`
	State   string `json:"state"`
	Source  string `json:"source"`
	Url     string `json:"url"`
	Address string `json:"address"`
}

func (h *Heritage) getHeritage(db *sql.DB) error {
  return db.QueryRow("SELECT id, file, name, class, status, address FROM national_heritage_area WHERE id=$1",
        h.ID).Scan(&h.ID, &h.File, &h.Name, &h.Class, &h.Status, &h.Address)
}

func (h *Heritage) updateHeritage(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (h *Heritage) deleteHeritage(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (h *Heritage) createHeritage(db *sql.DB) error {
	return errors.New("Not implemented")
}

func getHeritages(db *sql.DB, start, count int) ([]Heritage, error) {
	rows, err := db.Query(
		"SELECT id, file, name, class, status, address FROM national_heritage_area ORDER BY id LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	heritages := []Heritage{}

	for rows.Next() {
		var h Heritage
		if err := rows.Scan(&h.ID, &h.File, &h.Name, &h.Class, &h.Status, &h.Address); err != nil {
			return nil, err
		}
		heritages = append(heritages, h)
	}

	return heritages, nil
}
