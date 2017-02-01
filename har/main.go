package har

import "net/http"
import "time"
import "strings"

const version = "1.2"
const creator = "load-sink"

// FromHTTPRequest takes an http.Request and returns a HAR compatibale
// archive. An error is returned if creation of the archive fails validation
// or is otherwise impossible.
func FromHTTPRequest(r *http.Request) (h HAR, e error) {

	h.Log.Version = version
	h.Log.Creator = Creator{
		Name:    "load-sink",
		Version: "0.1",
	}

	h.Log.Entries = make([]Entry, 1)
	ent := &h.Log.Entries[0]
	ent.StartedDateTime = time.Now()
	ent.Time = 0
	ent.Request.Method = r.Method
	ent.Request.URL = r.URL.String()
	ent.Request.HTTPVersion = r.Proto

	ent.Request.Cookies = make([]Cookie, len(r.Cookies()))
	for i, c := range r.Cookies() {
		ent.Request.Cookies[i].Name = c.Name
		ent.Request.Cookies[i].Value = c.Value
	}

	ent.Request.Headers = make([]Header, len(r.Header))
	i := 0
	for key, value := range r.Header {
		ent.Request.Headers[i].Name = key
		ent.Request.Headers[i].Value = strings.Join(value, ",")
		i++
	}

	ent.Request.QueryString = make([]QueryString, len(r.URL.Query()))
	i = 0
	for key, value := range r.URL.Query() {
		ent.Request.QueryString[i].Name = key
		ent.Request.QueryString[i].Value = strings.Join(value, ",")
		i++
	}

	return
}
