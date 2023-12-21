package apirequest

import(
	"net/http"
	"net/url"
	"github.com/Didjacome/hmac-veracode/hmac"
	"io"
	"os"
	"bytes"
)

func ApiRequest(apiKeyID, apiKeySecret, apiUrl, httpMethod string) string {
	parsedUrl, err := url.Parse(apiUrl)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}

	req, err := http.NewRequest(httpMethod, parsedUrl.String(), nil)
	if err != nil {
		panic(err)
	}

	authorizationHeader, err := hmac.CalculateAuthorizationHeader(parsedUrl, httpMethod, apiKeyID, apiKeySecret)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorization", authorizationHeader)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != http.StatusOK {
		panic("Expected status 200. Status was: " + resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}


	file, err := os.Create("response.json")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	_, err = io.Copy(file, bytes.NewReader(body))
	if err != nil {
		panic(err)
	}

	return string(body[:])
}