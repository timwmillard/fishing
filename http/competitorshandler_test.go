package http

import (
	"context"
	"fishing"
	"fishing/mock"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
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

	ctx := context.Background()

	mockRepo := &mock.CompetitorsRepo{}
	mockRepo.On("List", ctx).Return([]fishing.Competitor{comp1, comp2}, nil)

	// Create the handler
	compHandler := NewCompetitorsHandler(mockRepo)
	_ = httptest.NewServer(compHandler)

	req := httptest.NewRequest(http.MethodGet, "/competitors/", nil)
	w := httptest.NewRecorder()

	// Call the handler
	compHandler.List(w, req)

	// competitors := make([]fishing.Competitor, 2)
	resp := w.Result()
	// decoder := json.NewDecoder(resp.Body)
	// err := decoder.Decode(&competitors)
	// if err != nil {
	// 	t.Fatalf("json.Decode: %v", err)
	// }
	output, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("ioutil.ReadAll(resp.Body) err = %v", err)
	}
	fmt.Printf("competitors %v", string(output))

	// respBody, _ := io.ReadAll(resp.Body)
}

func TestCompetitorsHanderCreate(t *testing.T) {
	reqBody, err := os.Open("testdata/competitors-create-request.json")
	if err != nil {
		t.Fatal("Unable to open testdata/competitors-create-request.json")
	}
	httptest.NewRequest(http.MethodGet, "/competitors/", reqBody)
}
