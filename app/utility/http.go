package utility

import (
	"io"
	"net/http"
)

func Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
        return nil, err
    }
	return io.ReadAll(resp.Body)
}