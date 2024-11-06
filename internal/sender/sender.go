package sender

import (
	"io"
	"net/http"
	"strings"
)

type Sender interface {
	Send(link string) (string, error)
}

type RequestSender struct {
	cookies string
}

func NewRequestSender(cookies string) *RequestSender {
	return &RequestSender{
		cookies: cookies,
	}
}

func (s *RequestSender) Send(link string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", link, strings.NewReader(s.cookies))
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
