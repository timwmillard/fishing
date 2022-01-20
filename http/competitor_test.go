package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/matryer/is"
	"github.com/timwmillard/fishing"
	"github.com/timwmillard/fishing/mock"
	"github.com/timwmillard/fishing/postgres"
	pmock "github.com/timwmillard/fishing/postgres/mock"
	"github.com/timwmillard/fishing/postgres/sqlc"
)

func TestCompetitorsHandler_List(t *testing.T) {
	is := is.New(t)

	want := mock.Competitors(2)

	mockRepo := &mock.CompetitorRepo{}
	mockRepo.ListFunc = func(ctx context.Context) ([]fishing.Competitor, error) {
		return want, nil
	}

	req := httptest.NewRequest(http.MethodGet, "/competitors/", nil)
	w := httptest.NewRecorder()

	compHandler := NewCompetitorHandler(mockRepo)

	// SUT
	compHandler.List(w, req)

	is.Equal(http.StatusOK, w.Code)
	is.True(mockRepo.ListInvoked)

	got := make([]fishing.Competitor, 2)
	json.Unmarshal(w.Body.Bytes(), &got)

	if !reflect.DeepEqual(got, want) {
		t.Logf("got %v\n", got)
		t.Logf("want %v\n", want)
		t.Errorf("deep equal failed")
	}
}

func TestCompetitorsHandler_List_empty(t *testing.T) {
	is := is.New(t)

	mockRepo := &mock.CompetitorRepo{}
	mockRepo.ListFunc = func(ctx context.Context) ([]fishing.Competitor, error) {
		return []fishing.Competitor{}, nil
	}

	req := httptest.NewRequest(http.MethodGet, "/competitors/", nil)
	w := httptest.NewRecorder()

	compHandler := NewCompetitorHandler(mockRepo)

	// SUT
	compHandler.List(w, req)

	is.Equal(http.StatusOK, w.Code)
	is.True(mockRepo.ListInvoked)

	got := make([]fishing.Competitor, 0)
	json.Unmarshal(w.Body.Bytes(), &got)

	is.Equal(len(got), 0)
}

func TestCompetitorsHandler_List_error(t *testing.T) {
	is := is.New(t)

	// want := fake.Competitors(2)

	mockRepo := &mock.CompetitorRepo{}
	mockRepo.ListFunc = func(ctx context.Context) ([]fishing.Competitor, error) {
		return nil, fmt.Errorf("test")
	}

	req := httptest.NewRequest(http.MethodGet, "/competitors/", nil)
	w := httptest.NewRecorder()

	compHandler := NewCompetitorHandler(mockRepo)

	// SUT
	compHandler.List(w, req)

	is.Equal(http.StatusNotFound, w.Code)
	is.True(mockRepo.ListInvoked)
}

func TestCompetitorsHandler_List_sqlc(t *testing.T) {
	is := is.New(t)

	want := pmock.Competitors(1)

	mockQueries := &pmock.CompetitorQueries{}
	mockQueries.ListCompetitorsFunc = func(ctx context.Context) ([]sqlc.FishingCompetitor, error) {
		return want, nil
	}

	req := httptest.NewRequest(http.MethodGet, "/competitors/", nil)
	w := httptest.NewRecorder()

	compRepo := postgres.NewCompetitorRepoWithQueries(mockQueries)
	compHandler := NewCompetitorHandler(compRepo)

	// SUT
	compHandler.List(w, req)

	is.Equal(http.StatusOK, w.Code)
	is.True(mockQueries.ListCompetitorsInvoked)

	got := make([]fishing.Competitor, 1)
	json.Unmarshal(w.Body.Bytes(), &got)

	is.Equal(got[0].ID, want[0].ID)
	is.Equal(got[0].Firstname, want[0].Firstname)
	is.Equal(got[0].Lastname, want[0].Lastname)
	is.Equal(got[0].Email, want[0].Email)
}
