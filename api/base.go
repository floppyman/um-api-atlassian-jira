package api

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/umbrella-sh/um-common/logging/ulog"
)

var doLogging = false
var jiraUrl string
var jiraUsername string
var jiraToken string

func Init(url string, username string, token string) {
	jiraUrl = url
	jiraUsername = username
	jiraToken = token
}

//goland:noinspection GoUnusedExportedFunction
func EnableLogging() {
	doLogging = true
}

func doGetRequest(urlPath string) ([]byte, error) {
	client := http.Client{
		Timeout: 60 * time.Second,
	}

	url := jiraUrl + urlPath
	//goland:noinspection GoBoolExpressions
	if doLogging {
		ulog.Console.Trace().Msg(url)
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return make([]byte, 0), err
	}

	req.Header.Add("Authorization", generateBase64BasicString())

	res, err := client.Do(req)
	if err != nil {
		return make([]byte, 0), err
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return make([]byte, 0), err
	}

	_ = res.Body.Close()

	return bodyBytes, nil
}

func doPostRequest(urlPath string, body []byte) ([]byte, error) {
	client := http.Client{
		Timeout: 60 * time.Second,
	}

	url := jiraUrl + urlPath
	//goland:noinspection GoBoolExpressions
	if doLogging {
		ulog.Console.Trace().Msg(url)
		ulog.Console.Trace().Msg(string(body))
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return make([]byte, 0), err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", generateBase64BasicString())

	res, err := client.Do(req)
	if err != nil {
		return make([]byte, 0), err
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return make([]byte, 0), err
	}

	_ = res.Body.Close()

	return bodyBytes, nil
}

func generateBase64BasicString() string {
	return fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", jiraUsername, jiraToken))))
}
