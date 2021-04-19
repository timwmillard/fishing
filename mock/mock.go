package mock

//go:generate mockgen -destination competitor.go -package mock -mock_names CompetitorsRepo=CompetitorsRepo github.com/timwmillard/fishing CompetitorsRepo
