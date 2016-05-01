package api

import (
	"golang.org/x/net/context"

	"github.com/go-kit/kit/endpoint"
	jujuratelimit "github.com/juju/ratelimit"

	"crypto/md5"
	"strings"
	"errors"
	"fmt"
)

func bucketLimiterMW(qps int64) endpoint.Middleware {
	tb := jujuratelimit.NewBucketWithRate(float64(qps), int64(qps))
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			if tb.TakeAvailable(1) == 0 {
				return createNodeEntityResponse{false, errors.New("Rate limit exceeded")}, nil
			}
			return next(ctx, request)
		}
	}
}

func cvcCheckerMW(cvc string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			req := request.(createNodeEntityRequest)
			err := checkCvcSum(req.Cvc, req.Nid, req.Uid, cvc); if err != nil {
				return createNodeEntityResponse{false, err}, nil
			}
			return next(ctx, request)
		}
	}
}


func checkCvcSum(cvc string, vals...interface{}) error {
	chunks := make([]string, 0, cap(vals))
	for _, v := range vals {

		var val string
		switch x := v.(type) {
		case string:
			val = x
		case fmt.Stringer:
			val = x.String()
		default:
			val = fmt.Sprint(x)
		}
		chunks = append(chunks, val)
	}

	key := strings.Join(chunks, "")
	sum := fmt.Sprintf("%x", md5.Sum([]byte(key)))

	if sum != cvc {
		fmt.Println("checkCvcSumErr", "key", key, "expectedSum", sum, "requestSum", cvc)
		return errors.New("Invalid request")
	}

	return nil
}