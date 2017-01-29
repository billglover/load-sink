package har

import "net/http"
import "time"

const version = "1.2"
const creator = "load-sink"

// FromHTTPRequest takes an http.Request and returns a HAR compatibale
// archive. An error is returned if creation of the archive fails validation
// or is otherwise impossible.
func FromHTTPRequest(r *http.Request) (h Log, e error) {

	h.Version = version
	h.Creator = Creator{
		Name:    "load-sink",
		Version: "0.1",
	}

	h.Entries = make([]Entry, 1)
	ent := &h.Entries[0]
	ent.StartedDateTime = time.Now()
	ent.Time = 0
	ent.Request.Method = r.Method
	ent.Request.URL = r.URL.String()

	return
}
