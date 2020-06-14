package slack

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const contentType = "application/json"

type Messenger struct {
	url string
}

func NewMessenger(url string) *Messenger {
	return &Messenger{url: url}
}

func (m *Messenger) Post(message []byte) error {
	resp, err := http.Post(m.url, contentType, bytes.NewReader(message))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return errors.New("failed to post message in stripe" + string(body))
	}

	return nil
}

func SimpleMessage(msg string) []byte {
	message := struct {
		Text string `json:"text"`
	}{Text: msg}

	data, _ := json.Marshal(&message)
	return data
}
