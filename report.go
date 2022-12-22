package fishing

import (
	"context"
)

type ReportService struct {
	competitorRepo CompetitorRepo
	catchRepo      CatchRepo
}

type CompetitorReport struct {
	Competitor
	Catches []Catch
}

func (rs *ReportService) CompetitorReport(ctx context.Context, id HashID) (CompetitorReport, error) {
	competitor, err := rs.competitorRepo.Get(ctx, id)
	if err != nil {
		return CompetitorReport{}, nil
	}

	catches, err := rs.catchRepo.ListByCompetitor(id)
	if err != nil {
		return CompetitorReport{}, nil
	}
	return CompetitorReport{
		Competitor: competitor,
		Catches:    catches,
	}, nil
}
