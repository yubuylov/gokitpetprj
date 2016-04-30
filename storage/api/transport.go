package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/net/context"

	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	b "github.com/yubuylov/gokitpetprj/storage/backend"
	"strconv"
)

func Handler(ctx context.Context, bs b.Service, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(encodeError),
	}

	getNodeEntity := kithttp.NewServer(
		ctx,
		makeGetNodeEntityEndpoint(bs),
		decodeGetNodeEntityRequest,
		encodeResponse,
		opts...,
	)

	getNodeEntities := kithttp.NewServer(
		ctx,
		makeGetNodeEntitiesEndpoint(bs),
		decodeGetNodeRequest,
		encodeResponse,
		opts...,
	)

	getNodeEntitiesCount := kithttp.NewServer(
		ctx,
		makeGetNodeEntitiesCountEndpoint(bs),
		decodeGetNodeRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/api/v1/{nid:[0-9]+}/entities", getNodeEntities).Methods("GET")
	r.Handle("/api/v1/{nid:[0-9]+}/entities/{id:[0-9]+}", getNodeEntity).Methods("GET")
	r.Handle("/api/v1/{nid:[0-9]+}/entities/count", getNodeEntitiesCount).Methods("GET")

	return r
}

var errBadRoute = errors.New("bad route")

func decodeGetNodeRequest(r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	nid, ok := vars["nid"]
	if !ok {
		return nil, errBadRoute
	}

	NidI64, err := strconv.ParseInt(nid, 10, 64); if err != nil {
		return nil, errBadRoute
	}
	NodeID := b.NodeId(NidI64)
	return getNodeRequest{NodeID: NodeID}, nil
}

func decodeGetNodeEntityRequest(r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	nid, ok := vars["nid"]
	if !ok {
		return nil, errBadRoute
	}

	id, ok := vars["id"]
	if !ok {
		return nil, errBadRoute
	}

	NidI64, err := strconv.ParseInt(nid, 10, 64); if err != nil {
		return nil, errBadRoute
	}
	IdI64, err := strconv.ParseInt(id, 10, 64); if err != nil {
		return nil, errBadRoute
	}

	return getNodeEntityRequest{NodeID: b.NodeId(NidI64), EntityId: b.EntityId(IdI64)}, nil
}

func encodeResponse(w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(w, e.error())
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errorer interface {
	error() error
}

// encode errors from business-logic
func encodeError(w http.ResponseWriter, err error) {
	switch err {
	//case cargo.ErrUnknown:
	//	w.WriteHeader(http.StatusNotFound)
	//case ErrInvalidArgument:
	//	w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
