package api

import (
	"golang.org/x/net/context"

	"github.com/go-kit/kit/endpoint"
	b "github.com/yubuylov/gokitpetprj/storage/backend"
)

type getNodeRequest struct {
	NodeID b.NodeId
}
type getNodeEntityRequest struct {
	NodeID   b.NodeId
	EntityId b.EntityId
}

type getNodeEntityResponse struct {
	Entity b.Entity `json:"entity"`
	Err    error   `json:"error,omitempty"`
}

func (r getNodeEntityResponse) error() error {
	return r.Err
}

type getNodeEntitiesResponse struct {
	Entities []b.Entity `json:"entities"`
	Err      error   `json:"error,omitempty"`
}

func (r getNodeEntitiesResponse) error() error {
	return r.Err
}

type getNodeEntitiesCountResponse struct {
	Count int64 `json:"count"`
	Err   error   `json:"error,omitempty"`
}

func (r getNodeEntitiesCountResponse) error() error {
	return r.Err
}

func makeGetNodeEntityEndpoint(s b.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getNodeEntityRequest)
		m, err := s.GetNodeEntity(req.NodeID, req.EntityId)
		return getNodeEntityResponse{Entity: m, Err: err}, nil
	}
}

func makeGetNodeEntitiesCountEndpoint(s b.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getNodeRequest)
		count, err := s.GetNodeEntitiesCount(req.NodeID)
		return getNodeEntitiesCountResponse{Count: count, Err: err}, nil
	}
}

func makeGetNodeEntitiesEndpoint(s b.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getNodeRequest)
		list, err := s.GetNodeEntities(req.NodeID)
		return getNodeEntitiesResponse{Entities: list, Err: err}, nil
	}
}