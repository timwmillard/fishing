package memory

import (
	"context"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/timwmillard/fishing"
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

func TestCreateAndGet(t *testing.T) {
	ctx := context.Background()
	repo := NewCompetitorRepo()

	want, err := repo.Create(ctx, comp1)
	if err != nil {
		t.Fatalf("unable to create competitor: %v", err)
	}

	got, err := repo.Get(ctx, want.ID)
	if err != nil {
		t.Fatalf("get competitor repo: %v", err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got a different competitor: got %v, want %v", got, want)
	}
}

func TestCreateAndList(t *testing.T) {
	ctx := context.Background()
	repo := NewCompetitorRepo()

	_, err := repo.Create(ctx, comp1)
	if err != nil {
		t.Fatalf("unable to create competitor: %v", err)
	}
	_, err = repo.Create(ctx, comp2)
	if err != nil {
		t.Fatalf("unable to create competitor: %v", err)
	}

	competitors, err := repo.List(ctx)
	if err != nil {
		t.Fatalf("unable to create competitor: %v", err)
	}

	if len(competitors) != 2 {
		t.Errorf("should be 2 competitors but got %v competitors", len(competitors))
	}
}

func TestGetCompetitorNotFound(t *testing.T) {
	ctx := context.Background()
	repo := NewCompetitorRepo()

	id := uuid.New()
	_, err := repo.Get(ctx, id)
	if err != fishing.ErrCompetitorNotFound {
		t.Error("should have got ErrcompetitorNotFound")
	}
}

func TestCreateAndUpdate(t *testing.T) {
	ctx := context.Background()
	repo := NewCompetitorRepo()

	c, err := repo.Create(ctx, comp1)
	if err != nil {
		t.Fatalf("unable to create competitor: %v", err)
	}

	want := comp2
	want.ID = c.ID
	got, err := repo.Update(ctx, want)
	if err != nil {
		t.Fatalf("unable to update competitor: %v", err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("got a different competitor: got %v, want %v", got, want)
	}

}

func TestUpdateCompetitorNotFound(t *testing.T) {
	ctx := context.Background()
	repo := NewCompetitorRepo()
	comp1.ID = uuid.New()
	_, err := repo.Update(ctx, comp1)
	if err != fishing.ErrCompetitorNotFound {
		t.Error("should have got ErrCompetitorNotFound")
	}
}

func TestCreateAndDelete(t *testing.T) {
	ctx := context.Background()
	repo := NewCompetitorRepo()

	c, err := repo.Create(ctx, comp1)
	if err != nil {
		t.Fatalf("unable to create competitor: %v", err)
	}

	err = repo.Delete(ctx, c.ID)
	if err != nil {
		t.Fatalf("unable to delete competitor: %v", err)
	}

	competitors, err := repo.List(ctx)
	if err != nil {
		t.Fatalf("unable to list competitors: %v", err)
	}

	if len(competitors) != 0 {
		t.Errorf("competitor list should be empty but has %d competitors", len(competitors))
	}
}
