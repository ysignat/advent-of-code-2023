package calendar

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

const ADVENT_OF_CODE_BASE_URL = "https://adventofcode.com"

type DatasetGetter interface {
	GetDataset(year uint, day uint) (string, error)
}

type WebDataset struct {
	session_cookie string
}

func NewWebDataset(sessionCookie string) (WebDataset, error) {
	return WebDataset{session_cookie: sessionCookie}, nil
}

func (wd WebDataset) GetDataset(year uint, day uint) (string, error) {
	content := ""

	request_url := fmt.Sprintf("%s/%d/day/%d/input", ADVENT_OF_CODE_BASE_URL, year, day)
	req, err := http.NewRequest("GET", request_url, nil)
	if err != nil {
		return content, err
	}

	cookie := new(http.Cookie)
	cookie.Name, cookie.Value = "session", wd.session_cookie

	req.AddCookie(cookie)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return content, err
	}

	if resp.StatusCode != 200 {
		return content, errors.New(resp.Status)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return content, err
	}

	return string(respBody), nil

}

type FilesystemDataset struct {
	root_path string
}

func NewFilesystemDataset(rootPath string) (FilesystemDataset, error) {
	return FilesystemDataset{root_path: rootPath}, nil
}

func (fd FilesystemDataset) GetDataset(year uint, day uint) (string, error) {
	path := path.Join(fd.root_path, fmt.Sprint(year), fmt.Sprintf("day%d", day))

	content := ""

	byte_content, err := os.ReadFile(path)
	if err != nil {
		return content, err
	}

	return string(byte_content), nil
}
