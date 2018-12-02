package vehicle

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-resty/resty"

	"trans/server/shared/communication/representation"
)

type vehicleClient struct {
	address string
}

func NewVehicleServiceClient(address string) *vehicleClient {
	return &vehicleClient{address: address}
}

func (rcv *vehicleClient) GetVehicleByUniqueIdentifier(
	in *representation.GetVehicleRequest,
) (
	*representation.GetVehicleResponse, error,
) {
	out := &representation.GetVehicleResponse{}
	errResp := &representation.ErrorResponse{}

	response, err := resty.R().
		SetQueryParam("uniqueIdentifier", in.UniqueIdentifier).
		SetResult(&out).
		SetError(&errResp).
		Get(fmt.Sprintf("%s/api/vehicles", rcv.address))

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
