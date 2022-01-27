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
	return nil, nil
}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
