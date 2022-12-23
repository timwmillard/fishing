package fake

import (
	"github.com/icrowley/fake"
	"github.com/timwmillard/fishing"
)

func Competitor() fishing.Competitor {
	hid, _ := fishing.NewHashID(fake.Digits())
	return fishing.Competitor{
		ID:        hid,
		FirstName: fake.FirstName(),
		LastName:  fake.LastName(),
		Email:     fake.EmailAddress(),
		Address1:  fake.StreetAddress(),
		Suburb:    fake.City(),
		Postcode:  fake.DigitsN(4),
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
