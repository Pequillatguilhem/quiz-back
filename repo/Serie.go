package repo

import (
	"Quiz-back/model"
	"database/sql"
	"fmt"
)

type Serie struct {
	db *sql.DB
}

func NewSerie(host string, port int, user string, password string, dbname string) (*Serie, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	return &Serie{
		db: db,
	}, nil
}

func (s *Serie) InsertSerie(name string) error {
	_, err := s.db.Exec("INSERT INTO Serie (name) VALUES ($1)", name)
	return err
}

func (s *Serie) InsertQuestion(name string, response string, serie int) error {
	_, err := s.db.Exec("INSERT INTO Question	 (name, response, serie_id) VALUES ($1,$2,$3)", name, response, serie)
	return err
}

func (s *Serie) SelectSeries() ([]model.Serie, error) {
	res, err := s.db.Query("SELECT id, name FROM serie")
	if err != nil {
		return nil, err
	}
	allSeries := []model.Serie{}
	for res.Next() {
		var id int
		var name string
		err := res.Scan(&id, &name)
		s := model.Serie{
			Id:   id,
			Name: name,
		}
		allSeries = append(allSeries, s)
		if err != nil {
			fmt.Println("Scan  get psw error error %v", err)
			return nil, err
		}
		//uuidRes, err := uuid.FromBytes(id)
	}
	return allSeries, nil
}

func (s *Serie) SelectQuestions(idSerie int) ([]model.Question, error) {
	res, err := s.db.Query("SELECT id, name, response FROM question WHERE serie_id = $1", idSerie)
	if err != nil {
		return nil, err
	}
	questions := []model.Question{}
	for res.Next() {
		var id *int
		var name string
		var response string

		err := res.Scan(&id, &name, &response)
		question := model.Question{
			Id:   id,
			Name: name, Response: response,
		}
		questions = append(questions, question)
		if err != nil {
			fmt.Println("Scan  get psw error error %v", err)
			return nil, err
		}
	}
	return questions, nil
}

func (s *Serie) DeleteQuestion(id int) error {
	_, err := s.db.Exec("DELETE FROM question WHERE id = $1", id)
	return err
}

func (s *Serie) DeleteSerie(id int) error {
	_, err := s.db.Exec("DELETE FROM serie WHERE id = $1", id)
	return err
}
