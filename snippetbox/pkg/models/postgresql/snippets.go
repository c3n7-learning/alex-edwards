package postgresql

import (
	"database/sql"
	"fmt"

	"github.com/c3n7/alex-edwards/snippetbox/pkg/models"
)

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title, content string, expires int) (int, error) {

	sttmt := fmt.Sprintf(`INSERT INTO snippets (title, content, created, expires) 
	VALUES ($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP + INTERVAL '%d days' ) RETURNING id;`, expires)

	var insertedId int
	err := m.DB.QueryRow(sttmt, title, content).Scan(&insertedId)
	if err != nil {
		return 0, err
	}

	return insertedId, nil
}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	sttmt := `SELECT id, title, content, created, expires FROM snippets 
	WHERE expires > CURRENT_TIMESTAMP AND id = $1`

	row := m.DB.QueryRow(sttmt, id)

	s := &models.Snippet{}
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err == sql.ErrNoRows {
		return nil, models.ErroNoRecord
	} else if err != nil {
		return nil, err
	}

	return s, nil
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	sttmt := `SELECT id, title, content, created, expires FROM snippets 
	WHERE expires > CURRENT_TIMESTAMP ORDER BY created DESC LIMIT 10`

	rows, err := m.DB.Query(sttmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	snippets := []*models.Snippet{}

	for rows.Next() {
		s := &models.Snippet{}
		err := rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}

	// if there was any error while iterating over data
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
