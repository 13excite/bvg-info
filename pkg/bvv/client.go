package bvv

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/13excite/bvg-info/pkg/store"
	"go.uber.org/zap"
)

// BvgClient is a struct of bvg api client
type BvgClient struct {
	httpClient *http.Client
	APIURL     string
	logger     *zap.SugaredLogger
}

const (
	// nearbyDepartesPath is a template for getting path to nearby departures
	nearbyDepartesPath = "/stops/%d/departures?duration=20"
)

func NewClent(apiUrl string) *BvgClient {
	return &BvgClient{
		APIURL: apiUrl,
		logger: zap.S().With("package", "bvv-client"),
		httpClient: &http.Client{
			Timeout: time.Second * 5,
		},
	}
}

// SetHttpClient redefines default http client
func (c *BvgClient) SetHTTPClient(client *http.Client) {
	c.httpClient = client
}

// GetNearbyDepartes returns departures from nearby stops by stopID
func (c *BvgClient) GetNearbyDepartes(stopID int) (*store.Departures, error) {
	urlString := c.APIURL + fmt.Sprintf(nearbyDepartesPath, stopID)
	req, _ := http.NewRequest(http.MethodGet, urlString, nil)

	res, err := c.httpClient.Do(req)
	if err != nil {
		c.logger.Error("GetNearbyDepartes got a client error", "error", err, "address", urlString)
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		c.logger.Error("GetNearbyDepartes finished with error status code. ", "StatusCode:", res.StatusCode,
			"address", urlString,
		)
		return nil, fmt.Errorf("Unexpected status code of GetNearbyDepartes: %d", res.StatusCode)
	}

	defer res.Body.Close()

	var departures store.Departures

	if err := json.NewDecoder(res.Body).Decode(&departures); err != nil {
		c.logger.Error("GetNearbyDepartes decoding error", "error", err)
		return nil, fmt.Errorf("failed to decode body into departes slice: %s", err.Error())
	}

	return &departures, nil
}
