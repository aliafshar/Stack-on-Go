package stackongo

import (
	"http"
	"json"
	"io/ioutil"
	"url"
	"os"
)

var host string = "http://api.stackexchange.com"

type Session struct {
	Site string
}

func NewSession(site string) *Session {
	return &Session{Site: site}
}

func setHost(url string) {
	host = url
}

// construct the endpoint URL
func setupEndpoint(path string, params map[string]string) *url.URL {
	base_url, _ := url.Parse(host)
	endpoint, _ := base_url.Parse("/2.0/" + path)

	query := endpoint.Query()
	for key, value := range params {
		query.Set(key, value)
	}

	endpoint.RawQuery = query.Encode()

	return endpoint
}

// parse the response
func parseResponse(response *http.Response, result interface{}) (error os.Error) {
	// close the body when done reading
	defer response.Body.Close()

	//read the response
	bytes, error := ioutil.ReadAll(response.Body)

	if error != nil {
		return
	}

	//parse JSON
	error = json.Unmarshal(bytes, result)

	if error != nil {
		print(error.String())
	}

	//check whether the response is a bad request
	if response.StatusCode == 400 {
		error = os.NewError("Bad Request")
	}

	return
}

// make the request
func (session Session) get_old(section string, params map[string]string) (*http.Response, os.Error) {
	client := new(http.Client)

	//set parameters for querystring
	params["site"] = session.Site

	return client.Get(setupEndpoint(section, params).String())
}

func (session Session) get(section string, params map[string]string, collection interface{}) (error os.Error) {
	//set parameters for querystring
	params["site"] = session.Site

	return get(section, params, collection)
}

func get(section string, params map[string]string, collection interface{}) (error os.Error) {
	client := new(http.Client)

	response, error := client.Get(setupEndpoint(section, params).String())

	if error != nil {
		return
	}

	error = parseResponse(response, collection)

	return

}
