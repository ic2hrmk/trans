package collector_http_client

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-resty/resty"

	"trans/server/shared/communication/representation"
)

type reportClient struct {
	address string
}

func NewCollectorServiceClient(address string) *reportClient {
	return &reportClient{address: address}
}

func (rcv *reportClient) CreateReport(
	in *representation.CreateReportRequest,
) (
	*representation.CreateReportResponse, error,
) {
	out := &representation.CreateReportResponse{}
	errResp := &representation.ErrorResponse{}

	response, err := resty.R().
		SetBody(in).
		SetResult(&out).
		SetError(&errResp).
		Post(fmt.Sprintf("%s/api/reports", rcv.address))

	if err != nil {
		return nil, err
	}

	if response.StatusCode() != http.StatusCreated {
		err = errors.New("empty error message")

		if errResp != nil {
			err = errors.New(errResp.Message)
		}

		return nil, err
	}

	return out, nil
}
