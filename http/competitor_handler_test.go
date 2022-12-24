package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/matryer/is"
	"github.com/timwmillard/fishing"
	"github.com/timwmillard/fishing/mock"
	"github.com/timwmillard/fishing/mock/fake"
)

func TestCompetitorsHandler_List(t *testing.T) {
	is := is.New(t)

	want := fake.Competitors(2)

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
	is.True(len(mockRepo.ListCalls()) > 0)

	got := make([]fishing.Competitor, 2)
	json.Unmarshal(w.Body.Bytes(), &got)

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("List() mismatch (-want +got):\n%s", diff)
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
	is.True(len(mockRepo.ListCalls()) > 0)

	got := make([]fishing.Competitor, 0)
	json.Unmarshal(w.Body.Bytes(), &got)

	is.Equal(len(got), 0)
}

func TestCompetitorsHandler_List_error(t *testing.T) {
	is := is.New(t)

	mockRepo := &mock.CompetitorRepo{}
	mockRepo.ListFunc = func(ctx context.Context) ([]fishing.Competitor, error) {
		return nil, fmt.Errorf("test")
	}

	req := httptest.NewRequest(http.MethodGet, "/competitors/", nil)
	w := httptest.NewRecorder()

	compHandler := NewCompetitorHandler(mockRepo)

	// SUT
	compHandler.List(w, req)

	is.Equal(http.StatusInternalServerError, w.Code)
	is.True(len(mockRepo.ListCalls()) > 0)
}
