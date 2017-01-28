package har

import (
	"net/http"
	"net/url"
	"strings"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestFromHTTPRequest(t *testing.T) {
	t.Log("given the need to test creating a HAR from an http.Request")

	// create sample request
	form := url.Values{}
	form.Add("name", "form")
	req, err := http.NewRequest(http.MethodPost, "https://demo.local/request/sample", strings.NewReader(form.Encode()))
	if err != nil {
		t.Fatal("\tshould be able to create a request", ballotX, err)
	}
	t.Log("\tshould be able to create a request", checkMark)

	// convert to HAR object
	h, err := FromHTTPRequest(req)
	if err != nil {
		t.Fatal("\tshould be able to create a HAR object", ballotX, err)
	}
	t.Log("\tshould be able to create a HAR object", checkMark)

	// validate HAR object
	if h.Version != "1.2" {
		t.Fatal("\tshould contain expected version number", ballotX, h.Version)
	}
	t.Log("\tshould contain expected version number", checkMark)

	if h.Creator.Name == "" {
		t.Fatal("\tshould contain the creator name", ballotX, h.Creator.Name)
	}
	t.Log("\tshould contain the creator name", checkMark)

	if h.Creator.Name == "" {
		t.Fatal("\tshould contain the creator version", ballotX, h.Creator.Version)
	}
	t.Log("\tshould contain the creator version", checkMark)

	if len(h.Entries) == 0 {
		t.Fatal("\tshould include at least one entry", ballotX, len(h.Entries))
	}
	t.Log("\tshould include at least one entry", checkMark)

	// validate the first entry
	t.Log("\tgiven the need to test the first entry in the log")
	e := h.Entries[0]

	if e.StartedDateTime.IsZero() {
		t.Fatal("\t\tshould include a valid startedDateTime", ballotX, e.StartedDateTime)
	}
	t.Log("\t\tshould include a valid startedDateTime", checkMark)

	if e.Time != 0 {
		t.Fatal("\t\tshould indicate a total processing time of 0 for all requests", ballotX, e.Time)
	}
	t.Log("\t\tshould indicate a total processing time of 0 for all requests", checkMark)
}
