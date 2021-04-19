package mock

//go:generate mockgen -destination competitor.go -package mock -mock_names CompetitorRepo=CompetitorRepo github.com/timwmillard/fishing CompetitorRepo
