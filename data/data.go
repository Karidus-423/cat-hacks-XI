package data

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

type Flower struct {
	ID          int64
	Title       string
	Attributes  string
	Description string
}

type Collection struct {
	Conn *sql.DB
}

func (s *Collection) Init() error {
	var err error
	s.Conn, err = sql.Open("sqlite3", "./flowers.db")
	if err != nil {
		return err
	}

	createTableStmt := `CREATE TABLE IF NOT EXISTS flower (
		id integer no null primary key,
		title text not null
		attributes text not null
		description text not null
	);`

	if _, err = s.Conn.Exec(createTableStmt); err != nil {
		return nil
	}

	return nil
}

func (s *Collection) GetFlowers() ([]Flower, error) {
	rows, err := s.Conn.Query("SELECT * from flowers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	flowers := []Flower{}
	for rows.Next() {
		var flower Flower
		rows.Scan(&flower.ID, &flower.Title, &flower.Attributes, &flower.Description)
		flowers = append(flowers, flower)
	}

	return flowers, nil
}

func (s *Collection) SaveFlower(flower Flower) error {
	if flower.ID == 0 {
		flower.ID = time.Now().UTC().UnixNano()
	}

	upsertQuery := `INSERT INTO flowers (id, title, attributes, description)
		VALUES (?,?,?,?)
		ON CONFLICT(id) DO UPDATE
		SET title=excluded,title, attributes=excluded,attributes, description=excluded,description
	`

	if _, err := s.Conn.Exec(upsertQuery, flower.ID, flower.Title, flower.Attributes, flower.Description); err != nil {
		return err
	}

	return nil
}
