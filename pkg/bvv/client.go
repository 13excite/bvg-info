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
	nearbyDepartesPath = "/stops/123/departures"
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

func (c *BvgClient) GetNearbyDepartes() ([]store.StopDepartures, error) {
	req, _ := http.NewRequest(http.MethodGet, c.APIURL+nearbyDepartesPath, nil)

	res, err := c.httpClient.Do(req)
	if err != nil {
		c.logger.Error("GetNearbyDepartes clent error", "error", err, "address", c.APIURL+nearbyDepartesPath)
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		c.logger.Error("GetNearbyDepartes bad status code. ", "StatusCode:", res.StatusCode,
			"address", c.APIURL+nearbyDepartesPath,
		)
		return nil, fmt.Errorf("Unexpected status code of GetNearbyDepartes: %d", res.StatusCode)
	}

	defer res.Body.Close()

	var departes []store.StopDepartures

	if err := json.NewDecoder(res.Body).Decode(&departes); err != nil {
		c.logger.Error("GetNearbyDepartes decoding error", "error", err)
		return nil, fmt.Errorf("failed to decode body into departes slice: %w", err)
	}
	// bytes, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	c.logger.Error("GetNearbyDepartes reading body failed", "error", err)
	// 	return nil, err
	// }
	return departes, nil
}
