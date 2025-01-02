package main

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

func Get(arg string) (string, io.ReadCloser, error) {
	if fileInfo, err := os.Stat(arg); os.IsNotExist(err) || fileInfo.IsDir() {
		parsed, err := url.Parse(arg)
		if err != nil {
			return arg, nil, err
		}
		if parsed.Scheme == "" {
			arg = "http://" + arg
		}

		req, err := http.NewRequest("GET", arg, nil)
		if err != nil {
			return arg, nil, err
		}
		req.Header.Add("User-Agent", "reto/"+Version)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return arg, nil, errors.Unwrap(err)
		}
		return arg, resp.Body, nil
	} else {
		path, err := filepath.Abs(arg)
		if err != nil {
			return arg, nil, err
		}
		file, err := os.Open(path)
		return path, file, err
	}
}
