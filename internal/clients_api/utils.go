package clients_api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

func Get(url string, seconds int64) (io.ReadCloser, error) {

	duration := time.Duration(seconds) * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		if errors.Is(err, ctx.Err()) {
			return nil, fmt.Errorf("timeout")
		}
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code %d", resp.StatusCode)
	}

	return resp.Body, nil
}

func Decode(body io.ReadCloser, v interface{}) error {

	err := json.NewDecoder(body).Decode(&v)
	if err != nil {
		return err
	}

	return nil
}
