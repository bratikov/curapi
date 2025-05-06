package utils

import (
	"bytes"
	"compress/gzip"
	"io"
	"net/http"
	"os"
)

// Take in mind, that it will works fine only with small files
func DownloadFile(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func GunzipData(data []byte) ([]byte, error) {
	gzBuffer := bytes.NewBuffer(data)

	gzReader, err := gzip.NewReader(gzBuffer)
	if err != nil {
		return nil, err
	}

	var buffer bytes.Buffer
	_, err = buffer.ReadFrom(gzReader)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func GzipData(data []byte) ([]byte, error) {
	var gzBuffer bytes.Buffer
	gz := gzip.NewWriter(&gzBuffer)

	_, err := gz.Write(data)
	if err != nil {
		return nil, err
	}

	if err = gz.Flush(); err != nil {
		return nil, err
	}

	if err = gz.Close(); err != nil {
		return nil, err
	}

	return gzBuffer.Bytes(), nil
}

func SaveToFile(filepath string, buffer *[]byte) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, bytes.NewBuffer(*buffer))
	if err != nil {
		return err
	}

	return nil
}
