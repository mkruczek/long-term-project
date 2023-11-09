package healthReady

import (
	"context"
	"sync"
)

type Observer struct {
	ServiceContextName string
	Service            Pinger
}

type Pinger interface {
	Ping(ctx context.Context) error
}

type Service struct {
	observers []Observer
}

func New(observers ...Observer) *Service {
	return &Service{
		observers: observers,
	}
}

func (s Service) Health(ctx context.Context) bool {
	return true
}

func (s Service) Ready(ctx context.Context) map[string]error {
	wg := &sync.WaitGroup{}
	wg.Add(len(s.observers))

	type pingResponse struct {
		provider string
		err      error
	}

	pingCh := make(chan pingResponse, len(s.observers))

	for _, p := range s.observers {
		go func(p Observer) {
			e := p.Service.Ping(ctx)
			pingCh <- pingResponse{provider: p.ServiceContextName, err: e}
			wg.Done()
		}(p)
	}

	wg.Wait()
	close(pingCh)

	result := make(map[string]error)
	for e := range pingCh {
		result[e.provider] = e.err
	}

	return result
}
