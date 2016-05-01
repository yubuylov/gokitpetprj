package api

import (
	"golang.org/x/net/context"

	"github.com/go-kit/kit/endpoint"

	b "github.com/yubuylov/gokitpetprj/public-api/backend"
)

type createNodeEntityRequest struct {
	Nid int64 `json:"nid"`
	Uid int64 `json:"uid"`
	Cvc string `json:"cvc"`
}

type createNodeEntityResponse struct {
	Success bool `json:"entity"`
	Err     error   `json:"error,omitempty"`
}

func (r createNodeEntityResponse) error() error {
	return r.Err
}

func makeCreateNodeEntityEndpoint(s b.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createNodeEntityRequest)
		success, err := s.CreateNodeEntity(req.Nid, req.Uid)
		return createNodeEntityResponse{Success:success, Err: err}, nil
	}
}