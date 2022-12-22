// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"context"
	"github.com/timwmillard/fishing"
	"sync"
)

// Ensure, that CompetitorRepoMock does implement fishing.CompetitorRepo.
// If this is not the case, regenerate this file with moq.
var _ fishing.CompetitorRepo = &CompetitorRepoMock{}

// CompetitorRepoMock is a mock implementation of fishing.CompetitorRepo.
//
//	func TestSomethingThatUsesCompetitorRepo(t *testing.T) {
//
//		// make and configure a mocked fishing.CompetitorRepo
//		mockedCompetitorRepo := &CompetitorRepoMock{
//			CreateFunc: func(ctx context.Context, c fishing.CompetitorParams) (fishing.Competitor, error) {
//				panic("mock out the Create method")
//			},
//			DeleteFunc: func(ctx context.Context, id fishing.HashID) error {
//				panic("mock out the Delete method")
//			},
//			GetFunc: func(ctx context.Context, id fishing.HashID) (fishing.Competitor, error) {
//				panic("mock out the Get method")
//			},
//			ListFunc: func(ctx context.Context) ([]fishing.Competitor, error) {
//				panic("mock out the List method")
//			},
//			UpdateFunc: func(ctx context.Context, id fishing.HashID, c fishing.CompetitorParams) (fishing.Competitor, error) {
//				panic("mock out the Update method")
//			},
//		}
//
//		// use mockedCompetitorRepo in code that requires fishing.CompetitorRepo
//		// and then make assertions.
//
//	}
type CompetitorRepoMock struct {
	// CreateFunc mocks the Create method.
	CreateFunc func(ctx context.Context, c fishing.CompetitorParams) (fishing.Competitor, error)

	// DeleteFunc mocks the Delete method.
	DeleteFunc func(ctx context.Context, id fishing.HashID) error

	// GetFunc mocks the Get method.
	GetFunc func(ctx context.Context, id fishing.HashID) (fishing.Competitor, error)

	// ListFunc mocks the List method.
	ListFunc func(ctx context.Context) ([]fishing.Competitor, error)

	// UpdateFunc mocks the Update method.
	UpdateFunc func(ctx context.Context, id fishing.HashID, c fishing.CompetitorParams) (fishing.Competitor, error)

	// calls tracks calls to the methods.
	calls struct {
		// Create holds details about calls to the Create method.
		Create []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// C is the c argument value.
			C fishing.CompetitorParams
		}
		// Delete holds details about calls to the Delete method.
		Delete []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID fishing.HashID
		}
		// Get holds details about calls to the Get method.
		Get []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID fishing.HashID
		}
		// List holds details about calls to the List method.
		List []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// Update holds details about calls to the Update method.
		Update []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID fishing.HashID
			// C is the c argument value.
			C fishing.CompetitorParams
		}
	}
	lockCreate sync.RWMutex
	lockDelete sync.RWMutex
	lockGet    sync.RWMutex
	lockList   sync.RWMutex
	lockUpdate sync.RWMutex
}

// Create calls CreateFunc.
func (mock *CompetitorRepoMock) Create(ctx context.Context, c fishing.CompetitorParams) (fishing.Competitor, error) {
	if mock.CreateFunc == nil {
		panic("CompetitorRepoMock.CreateFunc: method is nil but CompetitorRepo.Create was just called")
	}
	callInfo := struct {
		Ctx context.Context
		C   fishing.CompetitorParams
	}{
		Ctx: ctx,
		C:   c,
	}
	mock.lockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	mock.lockCreate.Unlock()
	return mock.CreateFunc(ctx, c)
}

// CreateCalls gets all the calls that were made to Create.
// Check the length with:
//
//	len(mockedCompetitorRepo.CreateCalls())
func (mock *CompetitorRepoMock) CreateCalls() []struct {
	Ctx context.Context
	C   fishing.CompetitorParams
} {
	var calls []struct {
		Ctx context.Context
		C   fishing.CompetitorParams
	}
	mock.lockCreate.RLock()
	calls = mock.calls.Create
	mock.lockCreate.RUnlock()
	return calls
}

// Delete calls DeleteFunc.
func (mock *CompetitorRepoMock) Delete(ctx context.Context, id fishing.HashID) error {
	if mock.DeleteFunc == nil {
		panic("CompetitorRepoMock.DeleteFunc: method is nil but CompetitorRepo.Delete was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  fishing.HashID
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockDelete.Lock()
	mock.calls.Delete = append(mock.calls.Delete, callInfo)
	mock.lockDelete.Unlock()
	return mock.DeleteFunc(ctx, id)
}

// DeleteCalls gets all the calls that were made to Delete.
// Check the length with:
//
//	len(mockedCompetitorRepo.DeleteCalls())
func (mock *CompetitorRepoMock) DeleteCalls() []struct {
	Ctx context.Context
	ID  fishing.HashID
} {
	var calls []struct {
		Ctx context.Context
		ID  fishing.HashID
	}
	mock.lockDelete.RLock()
	calls = mock.calls.Delete
	mock.lockDelete.RUnlock()
	return calls
}

// Get calls GetFunc.
func (mock *CompetitorRepoMock) Get(ctx context.Context, id fishing.HashID) (fishing.Competitor, error) {
	if mock.GetFunc == nil {
		panic("CompetitorRepoMock.GetFunc: method is nil but CompetitorRepo.Get was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  fishing.HashID
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockGet.Lock()
	mock.calls.Get = append(mock.calls.Get, callInfo)
	mock.lockGet.Unlock()
	return mock.GetFunc(ctx, id)
}

// GetCalls gets all the calls that were made to Get.
// Check the length with:
//
//	len(mockedCompetitorRepo.GetCalls())
func (mock *CompetitorRepoMock) GetCalls() []struct {
	Ctx context.Context
	ID  fishing.HashID
} {
	var calls []struct {
		Ctx context.Context
		ID  fishing.HashID
	}
	mock.lockGet.RLock()
	calls = mock.calls.Get
	mock.lockGet.RUnlock()
	return calls
}

// List calls ListFunc.
func (mock *CompetitorRepoMock) List(ctx context.Context) ([]fishing.Competitor, error) {
	if mock.ListFunc == nil {
		panic("CompetitorRepoMock.ListFunc: method is nil but CompetitorRepo.List was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockList.Lock()
	mock.calls.List = append(mock.calls.List, callInfo)
	mock.lockList.Unlock()
	return mock.ListFunc(ctx)
}

// ListCalls gets all the calls that were made to List.
// Check the length with:
//
//	len(mockedCompetitorRepo.ListCalls())
func (mock *CompetitorRepoMock) ListCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockList.RLock()
	calls = mock.calls.List
	mock.lockList.RUnlock()
	return calls
}

// Update calls UpdateFunc.
func (mock *CompetitorRepoMock) Update(ctx context.Context, id fishing.HashID, c fishing.CompetitorParams) (fishing.Competitor, error) {
	if mock.UpdateFunc == nil {
		panic("CompetitorRepoMock.UpdateFunc: method is nil but CompetitorRepo.Update was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  fishing.HashID
		C   fishing.CompetitorParams
	}{
		Ctx: ctx,
		ID:  id,
		C:   c,
	}
	mock.lockUpdate.Lock()
	mock.calls.Update = append(mock.calls.Update, callInfo)
	mock.lockUpdate.Unlock()
	return mock.UpdateFunc(ctx, id, c)
}

// UpdateCalls gets all the calls that were made to Update.
// Check the length with:
//
//	len(mockedCompetitorRepo.UpdateCalls())
func (mock *CompetitorRepoMock) UpdateCalls() []struct {
	Ctx context.Context
	ID  fishing.HashID
	C   fishing.CompetitorParams
} {
	var calls []struct {
		Ctx context.Context
		ID  fishing.HashID
		C   fishing.CompetitorParams
	}
	mock.lockUpdate.RLock()
	calls = mock.calls.Update
	mock.lockUpdate.RUnlock()
	return calls
}
