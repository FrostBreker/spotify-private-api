package spotifyprivateapi

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	spotifyprivateapi "github.com/FrostBreker/spotify-private-api/responseTypes"
)

type html struct {
	Body body `xml:"body"`
}
type body struct {
	Content string `xml:",innerxml"`
}

var HttpClient = &http.Client{
	Transport: &http.Transport{
		MaxIdleConnsPerHost: 100,
		IdleConnTimeout:     90 * time.Second,
	},
	Timeout: 30 * time.Second,
}

func GetSpotifyClientId() (spotifyprivateapi.ResponseClientId, error) {
	jsonBody := []byte(`{"client_data":{"client_version":"1.2.16.334.g29fe6bdc","client_id":"d8a5ed958d274c2e8ee717e6a4b0971d","js_sdk_data":{"device_brand":"unknown","device_model":"unknown","os":"windows","os_version":"NT 10.0","device_id":"653eea96ea4044e6725f27bc508ea9a5","device_type":"computer"}}}`)
	bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequest(http.MethodPost, "https://clienttoken.spotify.com/v1/clienttoken", bodyReader)
	if err != nil {
		return spotifyprivateapi.ResponseClientId{}, fmt.Errorf("client: could not create request: %s", err)
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("accept-language", "en-US,en;q=0.9")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("sec-ch-ua", "\"Not.A/Brand\";v=\"8\", \"Chromium\";v=\"114\", \"Brave\";v=\"114\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-site")
	req.Header.Add("sec-gpc", "1")
	req.Header.Add("Referer", "https://open.spotify.com/")
	req.Header.Add("Referrer-Policy", "strict-origin-when-cross-origin")

	res, err := HttpClient.Do(req)
	if err != nil {
		return spotifyprivateapi.ResponseClientId{}, fmt.Errorf("client: error making HTTP request: %s", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return spotifyprivateapi.ResponseClientId{}, fmt.Errorf("client: error reading response body: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		return spotifyprivateapi.ResponseClientId{}, fmt.Errorf("client: received non-OK status code: %d", res.StatusCode)
	}

	var response spotifyprivateapi.ResponseClientId
	err = json.Unmarshal(body, &response)
	if err != nil {
		return spotifyprivateapi.ResponseClientId{}, fmt.Errorf("client: error parsing JSON response: %s", err)
	}

	return response, nil
}

func GetSpotifyClientToken() (spotifyprivateapi.ResponseClientToken, error) {
	// Load saved response from file
	savedResponseBytes, err := ioutil.ReadFile("response.json")
	if err != nil {
		// Handle error while reading the file
		response, err := FetchSpotifyClientToken()
		if err != nil {
			// Handle error while fetching new token
			return spotifyprivateapi.ResponseClientToken{}, fmt.Errorf("client: error fetching new token: %s", err)
		}
		return response, nil
	}

	// Unmarshal saved response
	var savedResponse spotifyprivateapi.ResponseClientToken
	err = json.Unmarshal(savedResponseBytes, &savedResponse)
	if err != nil {
		// Handle error while parsing JSON
		return spotifyprivateapi.ResponseClientToken{}, fmt.Errorf("client: error parsing JSON: %s", err)
	}

	// Check if access token is expired
	currentTimeMs := time.Now().UnixNano() / int64(time.Millisecond)
	if savedResponse.AccessTokenExpirationMs <= currentTimeMs {
		// Access token is expired, fetch a new one
		response, err := FetchSpotifyClientToken()
		if err != nil {
			// Handle error while fetching new token
			return spotifyprivateapi.ResponseClientToken{}, fmt.Errorf("client: error fetching new token: %s", err)
		}

		return response, nil
	}

	// Access token is valid, return saved response
	return savedResponse, nil
}

func FetchSpotifyClientToken() (spotifyprivateapi.ResponseClientToken, error) {
	req, err := http.NewRequest(http.MethodGet, "https://open.spotify.com/intl-fr", nil)
	if err != nil {
		return spotifyprivateapi.ResponseClientToken{}, fmt.Errorf("client: could not create request: %s", err)
	}

	res, err := HttpClient.Do(req)
	if err != nil {
		return spotifyprivateapi.ResponseClientToken{}, fmt.Errorf("client: error making HTTP request: %s", err)
	}
	b, err := io.ReadAll(res.Body)
	if err != nil {
		return spotifyprivateapi.ResponseClientToken{}, fmt.Errorf("client: error reading response body: %s", err)
	}

	h := html{}
	err = xml.NewDecoder(bytes.NewBuffer(b)).Decode(&h)
	if err != nil {
		fmt.Println("error", err)
		return spotifyprivateapi.ResponseClientToken{}, fmt.Errorf("client: error reading response body: %s", err)
	}

	tokenObject := strings.Split(h.Body.Content, `<script id="session" data-testid="session" type="application/json">`)[1]
	tokenObject = strings.Split(tokenObject, `</script>`)[0]

	var response spotifyprivateapi.ResponseClientToken
	err = json.Unmarshal([]byte(tokenObject), &response)
	if err != nil {
		return spotifyprivateapi.ResponseClientToken{}, fmt.Errorf("client: error parsing JSON response: %s", err)
	}
	//Save reponse in file
	err = ioutil.WriteFile("response.json", []byte(tokenObject), 0644)
	if err != nil {
		fmt.Println("error", err)
		// Handle error while writing to file
	}
	return response, nil
}

func SetHeadersForSpotifyRequest(req *http.Request) (*http.Request, error) {
	responseClient, err := GetSpotifyClientId()
	if err != nil {
		log.Println(err)
		return req, err
	}
	token, err := GetSpotifyClientToken()
	if err != nil {
		fmt.Println(err)
		return req, err
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("accept-language", "en")
	req.Header.Add("app-platform", "WebPlayer")
	req.Header.Add("authorization", "Bearer "+token.AccessToken)
	req.Header.Add("client-token", responseClient.GrantedToken.Token)
	req.Header.Add("content-type", "application/json;charset=UTF-8")
	req.Header.Add("sec-ch-ua", "\"Not.A/Brand\";v=\"8\", \"Chromium\";v=\"114\", \"Brave\";v=\"114\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-site")
	req.Header.Add("sec-gpc", "1")
	req.Header.Add("spotify-app-version", "1.2.16.233.gc90b6d01")
	req.Header.Add("Referer", "https://open.spotify.com/")
	req.Header.Add("Referrer-Policy", "strict-origin-when-cross-origin")

	return req, err
}

func GetFormatedTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func LogInfo(message string) {
	log.Printf("[INFO] --> %v", message)
}

func LogError(message string) {
	log.Printf("[ERROR] --> %v", message)
}

func LogDebug(message string) {
	log.Printf("[DEBUG] --> %v", message)
}
