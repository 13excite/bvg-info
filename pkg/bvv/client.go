package bvv

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

// BvgClient is a struct of bvg api client
type BvgClient struct {
	HttpClient *http.Client
	APIURL     string
	logger     *zap.SugaredLogger
}

const (
	nearbyDepartesPath = "/api/blbla"
)

func NewClent(httpClient *http.Client, apiUrl string) *BvgClient {
	return &BvgClient{
		HttpClient: httpClient,
		APIURL:     apiUrl,
		logger:     zap.S().With("package", "client"),
	}
}

func (c *BvgClient) GetNearbyDepartes(requestParams string) error {
	req, err := http.NewRequest("GET", c.APIURL+nearbyDepartesPath+requestParams, nil)
	if err != nil {
		c.logger.Error("Bvg http client error", "error", err, "address", c.APIURL+nearbyDepartesPath)
		return fmt.Errorf("Bvg http client error. Address: %s", c.APIURL+nearbyDepartesPath)
	}

	defer req.Body.Close()

	return nil
}
