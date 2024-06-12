package cl

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"bharvest.io/beramon/utils"
)

func (c *Client) GetValidatorInfo(ctx context.Context, valAddr string) (*Validator, error) {
	utils.Info("Querying validator info")

	resp, err := http.Get(c.host)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	result := Response{}
	err = json.Unmarshal(respBytes, &result)
	if err != nil {
		return nil, err
	}

	for _, val := range result.Data {
		if val.Addr == valAddr {
			return &val, nil
		}
	}

	return nil, errors.New("validator not found")
}
