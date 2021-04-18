package fake

import (
	"github.com/google/uuid"
	"github.com/icrowley/fake"
	"github.com/timwmillard/fishing"
)

func Competitor() fishing.Competitor {
	return fishing.Competitor{
		ID:        uuid.New(),
		Firstname: fake.FirstName(),
		Lastname:  fake.LastName(),
		Email:     fake.EmailAddress(),
		Address1:  fake.StreetAddress(),
		Suburb:    fake.City(),
		Postcode:  fake.DigitsN(4),
		Phone:     fake.Phone(),
		Mobile:    fake.Phone(),
	}
}

func Competitors(n int) []fishing.Competitor {
	comps := make([]fishing.Competitor, n)
	for i := 0; i < n; i++ {
		comps[i] = Competitor()
	}
	return comps
}
