package fake

import (
	"strconv"

	ifake "github.com/icrowley/fake"
	"github.com/timwmillard/fishing"
)

func Competitor() fishing.Competitor {
	id, _ := strconv.Atoi(ifake.Digits())
	return fishing.Competitor{
		ID:        fishing.NewHashID(id),
		FirstName: ifake.FirstName(),
		LastName:  ifake.LastName(),
		Email:     ifake.EmailAddress(),
		Address1:  ifake.StreetAddress(),
		Suburb:    ifake.City(),
		Postcode:  ifake.DigitsN(4),
		Phone:     ifake.Phone(),
		Mobile:    ifake.Phone(),
	}
}

func Competitors(n int) []fishing.Competitor {
	comps := make([]fishing.Competitor, n)
	for i := 0; i < n; i++ {
		comps[i] = Competitor()
	}
	return comps
}
