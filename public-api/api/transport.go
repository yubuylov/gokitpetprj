package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/net/context"

	"github.com/go-kit/kit/endpoint"
	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"

	b "github.com/yubuylov/gokitpetprj/public-api/backend"

	app_config "github.com/yubuylov/gokitpetprj/public-api/config"

)

func Handler(ctx context.Context, bs b.Service, logger kitlog.Logger, cfg app_config.AppConfig) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(encodeError),
	}

	var e endpoint.Endpoint
	{
		e = makeCreateNodeEntityEndpoint(bs)
		e = cvcCheckerMW(cfg.Server.CvcKey)(e)
		e = bucketLimiterMW(cfg.Server.Qps)(e)
	}

	createNodeEntity := kithttp.NewServer(ctx, e,
		decodeCreateNodeEntityRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/entities", createNodeEntity).Methods("POST")

	return r
}

func decodeCreateNodeEntityRequest(r *http.Request) (interface{}, error) {
	var request createNodeEntityRequest
	if err := decodeRequest(r, &request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeRequest(r *http.Request, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return err
	}
	return nil
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