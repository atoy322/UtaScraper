package main

import (
	"net/http"
)

func Get(url string) ([]byte, error) {
	response, e := http.Get(url)
	if e != nil {
		return nil, e
	}

	content := make([]byte, 1024)
	for {
		buf := make([]byte, 1024)
		n, e := response.Body.Read(buf)
		content = append(content, buf[:n]...)

		if (n == 0) || (e != nil) {
			break
		}
	}

	return content, nil
}
