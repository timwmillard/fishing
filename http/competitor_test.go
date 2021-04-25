package http

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/matryer/is"
	"github.com/timwmillard/fishing"
	"github.com/timwmillard/fishing/fake"
	"github.com/timwmillard/fishing/mock"
)

var (
	comp1 = fishing.Competitor{
		Firstname: "Tim",
		Lastname:  "Millard",
	}
	comp2 = fishing.Competitor{
		Firstname: "John",
		Lastname:  "Smith",
	}
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
	is.True(mockRepo.ListInvoked)

	got := make([]fishing.Competitor, 2)
	json.Unmarshal(w.Body.Bytes(), &got)

	if !reflect.DeepEqual(got, want) {
		t.Logf("got %v\n", got)
		t.Logf("want %v\n", want)
		t.Errorf("deep equal failed")
	}
}
