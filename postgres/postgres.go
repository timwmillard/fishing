package postgres

//go:generate sqlc generate

import (
	"database/sql"
	"time"
)

func nullIntP(i sql.NullInt32) *int {
	if _, err := i.Value(); err != nil {
		return nil
	}
	r := int(i.Int32)
	return &r
}

func nullInt(i sql.NullInt32) int {
	if !i.Valid {
		return 0
	}
	return int(i.Int32)
}

func nullStringP(s sql.NullString) *string {
	if _, err := s.Value(); err != nil {
		return nil
	}
	return &s.String
}

func nullString(s sql.NullString) string {
	if !s.Valid {
		return ""
	}
	return s.String
}

func nullTime(t sql.NullTime) time.Time {
	if !t.Valid {
		return time.Time{}
	}
	return t.Time
}
