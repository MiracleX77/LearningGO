package di

import (
	"database/sql"
	"testing"
)

type mockDB struct {
	query        string
	lastInsertID int64
	rowsAffected int64
}

func (m *mockDB) LastInsertId() (int64, error) {
	return m.lastInsertID, nil
}
func (m *mockDB) RowsAffected() (int64, error) {
	return m.rowsAffected, nil
}

func (m *mockDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	m.query = query
	return m, nil
}
func TestExecQuery(t *testing.T) {
	mock := &mockDB{
		rowsAffected: 32,
	}

	r, _ := execQuery(mock, "SELECT * FROM sql")

	if mock.query != "SELECT * FROM sql" {
		t.Error("should have been called with SELECT * FROM sql")
	}

	if r != 32 {
		t.Error("should have returned 32")
	}
}
