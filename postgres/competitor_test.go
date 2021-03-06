package postgres

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/matryer/is"
	"github.com/timwmillard/fishing"
	"github.com/timwmillard/fishing/postgres/mock"
	"github.com/timwmillard/fishing/postgres/sqlc"
)

func TestCompetitorRepo_List(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()

	numCompetitors := 10
	want := mock.Competitors(numCompetitors)

	queries := &mock.CompetitorQueries{}
	queries.ListCompetitorsFunc = func(ctx context.Context) ([]sqlc.FishingCompetitor, error) { return want, nil }

	repo := newCompetitorRepoWithQuerier(queries)

	// SUT
	got, err := repo.List(ctx)
	is.NoErr(err)
	is.Equal(len(got), len(want))
	for i := 0; i < numCompetitors; i++ {
		is.Equal(got[i].ID, want[i].ID)
		is.Equal(got[i].Firstname, want[i].Firstname)
		is.Equal(got[i].Lastname, want[i].Lastname)
	}
	is.True(queries.ListCompetitorsInvoked)
}

func TestCompetitorRepo_List_error(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()

	queries := &mock.CompetitorQueries{}
	queries.ListCompetitorsFunc = func(ctx context.Context) ([]sqlc.FishingCompetitor, error) { return nil, errors.New("test") }

	repo := newCompetitorRepoWithQuerier(queries)

	// SUT
	_, err := repo.List(ctx)
	if err == nil {
		t.Errorf("err == nil")
	}
	is.True(queries.ListCompetitorsInvoked)
}

func TestCompetitorRepo_Get(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()

	compID := uuid.New()
	want := sqlc.FishingCompetitor{ID: compID, Firstname: "Tim", Lastname: "Millard"}

	queries := &mock.CompetitorQueries{}
	queries.GetCompetitorFunc = func(ctx context.Context, id uuid.UUID) (sqlc.FishingCompetitor, error) {
		is.Equal(compID, id)
		return want, nil
	}

	repo := newCompetitorRepoWithQuerier(queries)

	// SUT
	got, err := repo.Get(ctx, compID)
	is.NoErr(err)
	is.Equal(got.ID, want.ID)
	is.Equal(got.Firstname, want.Firstname)
	is.Equal(got.Lastname, want.Lastname)
	is.True(queries.GetCompetitorInvoked)
}

func TestCompetitorRepo_Get_error(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()

	compID := uuid.New()

	queries := &mock.CompetitorQueries{}
	queries.GetCompetitorFunc = func(ctx context.Context, id uuid.UUID) (sqlc.FishingCompetitor, error) {
		is.Equal(compID, id)
		return sqlc.FishingCompetitor{}, errors.New("test")
	}

	repo := newCompetitorRepoWithQuerier(queries)

	// SUT
	_, err := repo.Get(ctx, compID)
	if err == nil {
		t.Errorf("err == nil")
	}
	is.True(queries.GetCompetitorInvoked)
}

func TestCompetitorRepo_Get_not_found(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()

	compID := uuid.New()

	queries := &mock.CompetitorQueries{}
	queries.GetCompetitorFunc = func(ctx context.Context, id uuid.UUID) (sqlc.FishingCompetitor, error) {
		is.Equal(compID, id)
		return sqlc.FishingCompetitor{}, sql.ErrNoRows
	}

	repo := newCompetitorRepoWithQuerier(queries)

	// SUT
	_, err := repo.Get(ctx, compID)
	if err == nil {
		t.Errorf("err == nil")
	}
	is.Equal(err, fishing.ErrCompetitorNotFound)
	is.True(queries.GetCompetitorInvoked)
}

func TestCompetitorRepo_Update(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()

	compID := uuid.New()
	want := fishing.Competitor{ID: compID, Firstname: "Tim", Lastname: "Millard"}
	wantDB := sqlc.FishingCompetitor{ID: compID, Firstname: "Tim", Lastname: "Millard"}

	queries := &mock.CompetitorQueries{}
	queries.UpdateCompetitorFunc = func(ctx context.Context, arg sqlc.UpdateCompetitorParams) (sqlc.FishingCompetitor, error) {
		is.Equal(wantDB.ID, arg.ID)
		is.Equal(wantDB.Firstname, arg.Firstname)
		is.Equal(wantDB.Lastname, arg.Lastname)
		return wantDB, nil
	}

	repo := newCompetitorRepoWithQuerier(queries)

	// SUT
	got, err := repo.Update(ctx, want)
	is.NoErr(err)
	is.Equal(got.ID, want.ID)
	is.Equal(got.Firstname, want.Firstname)
	is.Equal(got.Lastname, want.Lastname)
	is.True(queries.UpdateCompetitorInvoked)
}

func TestCompetitorRepo_Update_error(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()

	compID := uuid.New()

	queries := &mock.CompetitorQueries{}
	queries.UpdateCompetitorFunc = func(ctx context.Context, arg sqlc.UpdateCompetitorParams) (sqlc.FishingCompetitor, error) {
		is.Equal(compID, arg.ID)
		return sqlc.FishingCompetitor{}, errors.New("test")
	}

	repo := newCompetitorRepoWithQuerier(queries)
	want := fishing.Competitor{ID: compID, Firstname: "Tim", Lastname: "Millard"}

	// SUT
	_, err := repo.Update(ctx, want)
	if err == nil {
		t.Errorf("err == nil")
	}
	is.True(queries.UpdateCompetitorInvoked)
}

func TestCompetitorRepo_Update_not_found(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()

	compID := uuid.New()

	queries := &mock.CompetitorQueries{}
	queries.UpdateCompetitorFunc = func(ctx context.Context, arg sqlc.UpdateCompetitorParams) (sqlc.FishingCompetitor, error) {
		is.Equal(compID, arg.ID)
		return sqlc.FishingCompetitor{}, sql.ErrNoRows
	}

	repo := newCompetitorRepoWithQuerier(queries)
	want := fishing.Competitor{ID: compID, Firstname: "Tim", Lastname: "Millard"}

	// SUT
	_, err := repo.Update(ctx, want)
	if err == nil {
		t.Errorf("err == nil")
	}
	is.Equal(err, fishing.ErrCompetitorNotFound)
	is.True(queries.UpdateCompetitorInvoked)
}

func TestCompetitorRepo_Create(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()

	want := fishing.Competitor{Firstname: "Tim", Lastname: "Millard"}
	wantDB := sqlc.FishingCompetitor{ID: uuid.New(), Firstname: "Tim", Lastname: "Millard"}

	queries := &mock.CompetitorQueries{}
	queries.CreateCompetitorFunc = func(ctx context.Context, arg sqlc.CreateCompetitorParams) (sqlc.FishingCompetitor, error) {
		is.Equal(wantDB.Firstname, arg.Firstname)
		is.Equal(wantDB.Lastname, arg.Lastname)
		return wantDB, nil
	}

	repo := newCompetitorRepoWithQuerier(queries)

	// SUT
	got, err := repo.Create(ctx, want)
	is.NoErr(err)
	is.Equal(got.ID, wantDB.ID)
	is.Equal(got.Firstname, want.Firstname)
	is.Equal(got.Lastname, want.Lastname)
	is.True(queries.CreateCompetitorInvoked)
}

func TestCompetitorRepo_Create_error(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()

	queries := &mock.CompetitorQueries{}
	queries.CreateCompetitorFunc = func(ctx context.Context, arg sqlc.CreateCompetitorParams) (sqlc.FishingCompetitor, error) {
		return sqlc.FishingCompetitor{}, errors.New("test")
	}

	repo := newCompetitorRepoWithQuerier(queries)
	want := fishing.Competitor{Firstname: "Tim", Lastname: "Millard"}

	// SUT
	_, err := repo.Create(ctx, want)
	if err == nil {
		t.Errorf("err == nil")
	}
	is.True(queries.CreateCompetitorInvoked)
}

func TestCompetitorRepo_Delete(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()

	compID := uuid.New()

	queries := &mock.CompetitorQueries{}
	queries.DeleteCompetitorFunc = func(ctx context.Context, id uuid.UUID) (int64, error) {
		is.Equal(compID, id)
		return 1, nil
	}

	repo := newCompetitorRepoWithQuerier(queries)

	// SUT
	err := repo.Delete(ctx, compID)
	is.NoErr(err)
	is.True(queries.DeleteCompetitorInvoked)
}

func TestCompetitorRepo_Delete_error(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()

	compID := uuid.New()

	queries := &mock.CompetitorQueries{}
	queries.DeleteCompetitorFunc = func(ctx context.Context, id uuid.UUID) (int64, error) {
		is.Equal(compID, id)
		return 0, errors.New("test")
	}

	repo := newCompetitorRepoWithQuerier(queries)

	// SUT
	err := repo.Delete(ctx, compID)
	if err == nil {
		t.Errorf("err == nil")
	}
	is.True(queries.DeleteCompetitorInvoked)
}

func TestCompetitorRepo_Delete_not_found(t *testing.T) {
	is := is.New(t)
	ctx := context.Background()

	compID := uuid.New()

	queries := &mock.CompetitorQueries{}
	queries.DeleteCompetitorFunc = func(ctx context.Context, id uuid.UUID) (int64, error) {
		is.Equal(compID, id)
		return 0, nil
	}

	repo := newCompetitorRepoWithQuerier(queries)

	// SUT
	err := repo.Delete(ctx, compID)
	is.Equal(err, fishing.ErrCompetitorNotFound)
	is.True(queries.DeleteCompetitorInvoked)
}

func TestCompetitor(t *testing.T) {
	is := is.New(t)

	comp := sqlc.FishingCompetitor{ID: uuid.New(), Firstname: "Tim", Lastname: "Millard"}

	// SUT
	fcomp := competitor(comp)

	is.Equal(comp.ID, fcomp.ID)
	is.Equal(comp.Firstname, fcomp.Firstname)
	is.Equal(comp.Lastname, fcomp.Lastname)
}
