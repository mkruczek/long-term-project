package statistics_v2

import "context"

// Service - i think this is some kind of input port
type Service struct {
	getTrades getTrades
}

// New - create new Service
func New(gt getTrades) Service {
	return Service{
		getTrades: gt,
	}
}

func (s Service) Calculate(ctx context.Context, f Filter) (Summary, error) {

	trades, err := s.getTrades.GetTrades(ctx, f)
	if err != nil {
		return Summary{}, err
	}

	return calculate(trades), nil
}
