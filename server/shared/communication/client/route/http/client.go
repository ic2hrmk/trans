package route

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-resty/resty"

	"trans/server/shared/communication/representation"
)

type routeClient struct {
	address string
}

func NewRouteServiceClient(address string) *routeClient {
	return &routeClient{address: address}
}

func (rcv *routeClient) GetRouteByID(
	in *representation.GetRouteRequest,
) (
	*representation.GetRouteResponse, error,
) {
	out := &representation.GetRouteResponse{}
	errResp := &representation.ErrorResponse{}

	response, err := resty.R().
		SetQueryParam("routeId", in.RouteID).
		SetResult(&out).
		SetError(&errResp).
		Get(fmt.Sprintf("%s/api/routes", rcv.address))

	if err != nil {
		return nil, err
	}

	if response.StatusCode() != http.StatusOK {
		err = errors.New("empty error message")

		if errResp != nil {
			err = errors.New(errResp.Message)
		}

		return nil, err
	}

	return out, nil
}
