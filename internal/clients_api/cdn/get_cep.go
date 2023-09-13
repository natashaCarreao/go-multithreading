package cdn

import (
	"fmt"
	"io"
	"net/http"
)

const (
	urlBase = "https://cdn.apicep.com/file/apicep/%s.json"
)

func (c client) GetCEP(cep string) (string, error) {
	url := fmt.Sprintf(urlBase, cep)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		return "", err
	}

	defer resp.Body.Close()

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return url, nil
}
