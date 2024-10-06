package endpoints

import (
	"context"

	"github.com/Zerofisher/stringsvc/pkg/stringsvc"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	UppercaseEndpoint endpoint.Endpoint
	CountEndpoint     endpoint.Endpoint
}

func MakeEndpoints(s stringsvc.StringService) Endpoints {
	return Endpoints{
		UppercaseEndpoint: makeUppercaseEndpoint(s),
		CountEndpoint:     makeCountEndpoint(s),
	}
}

func makeUppercaseEndpoint(s stringsvc.StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(stringsvc.UppercaseRequest)
		v, err := s.Uppercase(req.S)
		if err != nil {
			return stringsvc.UppercaseResponse{V: v, Err: err.Error()}, nil
		}
		return stringsvc.UppercaseResponse{V: v}, nil
	}
}

func makeCountEndpoint(s stringsvc.StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(stringsvc.CountRequest)
		v := s.Count(req.S)
		return stringsvc.CountResponse{V: v}, nil
	}
}
