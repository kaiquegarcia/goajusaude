package shared

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type HttpClient interface {
	Get(url string) ([]byte, error)
}

type httpClient struct {
	baseURL        string
	internalClient *http.Client
}

func NewHttpClient(baseURL string) HttpClient {
	return httpClient{
		baseURL:        baseURL,
		internalClient: &http.Client{Timeout: 30 * time.Second},
	}
}

func (client httpClient) Get(url string) ([]byte, error) {
	return client.request(http.MethodGet, url, nil)
}

func (client httpClient) request(method string, url string, body io.Reader) ([]byte, error) {
	request, err := http.NewRequest(method, client.baseURL+url, body)
	if err != nil {
		return nil, err
	}

	//request.Header.Set("")

	response, err := client.internalClient.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		return nil, fmt.Errorf("HTTP Error: %s", response.Status)
	}

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
