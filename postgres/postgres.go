package postgres

//go:generate sqlc generate

import "database/sql"

func nullIntP(i sql.NullInt32) *int {
	if _, err := i.Value(); err != nil {
		return nil
	}
	r := int(i.Int32)
	return &r
}

func nullInt(i sql.NullInt32) int {
	if _, err := i.Value(); err != nil {
		return 0
	}
	r := int(i.Int32)
	return r
}

func nullStringP(s sql.NullString) *string {
	if _, err := s.Value(); err != nil {
		return nil
	}
	return &s.String
}

func nullString(s sql.NullString) string {
	if _, err := s.Value(); err != nil {
		return ""
	}
	return s.String
}
