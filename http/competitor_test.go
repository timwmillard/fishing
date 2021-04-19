package http

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
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

	ctx := context.Background()

	fakeCompetitors := fake.Competitors(2)

	ctrl := gomock.NewController(t)

	mockRepo := mock.NewCompetitorRepo(ctrl)
	mockRepo.EXPECT().List(ctx).Return(fakeCompetitors, nil)

	// Create the handler
	compHandler := NewCompetitorHandler(mockRepo)
	_ = httptest.NewServer(compHandler)

	req := httptest.NewRequest(http.MethodGet, "/competitors/", nil)
	w := httptest.NewRecorder()

	// Call the handler
	compHandler.List(w, req)

	// competitors := make([]fishing.Competitor, 2)
	// resp := w.Result()
	// decoder := json.NewDecoder(resp.Body)
	// err := decoder.Decode(&competitors)
	// if err != nil {
	// 	t.Fatalf("json.Decode: %v", err)
	// }
	// output, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	t.Fatalf("ioutil.ReadAll(resp.Body) err = %v", err)
	// }
	// fmt.Printf("competitors %v", string(output))

	// respBody, _ := io.ReadAll(resp.Body)
}

func TestCompetitorHanderCreate(t *testing.T) {
	reqBody, err := os.Open("testdata/competitors-create-request.json")
	if err != nil {
		t.Fatal("Unable to open testdata/competitors-create-request.json")
	}
	httptest.NewRequest(http.MethodGet, "/competitors/", reqBody)
}
