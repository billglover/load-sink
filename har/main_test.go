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

	// add a custom header
	req.Header.Add("header_key", "header_value")

	// add a sample cookie
	c := &http.Cookie{
		Name:  "cookie_key",
		Value: "cookie_value",
	}
	req.AddCookie(c)

	// convert to HAR object
	h, err := FromHTTPRequest(req)
	l := &h.Log
	if err != nil {
		t.Fatal("\tshould be able to create a HAR object", ballotX, err)
	}
	t.Log("\tshould be able to create a HAR object", checkMark)

	// validate HAR object
	if l.Version != "1.2" {
		t.Fatal("\tshould contain expected version number", ballotX, l.Version)
	}
	t.Log("\tshould contain expected version number", checkMark)

	if l.Creator.Name == "" {
		t.Fatal("\tshould contain the creator name", ballotX, l.Creator.Name)
	}
	t.Log("\tshould contain the creator name", checkMark)

	if l.Creator.Name == "" {
		t.Fatal("\tshould contain the creator version", ballotX, l.Creator.Version)
	}
	t.Log("\tshould contain the creator version", checkMark)

	if len(l.Entries) == 0 {
		t.Fatal("\tshould include at least one entry", ballotX, len(l.Entries))
	}
	t.Log("\tshould include at least one entry", checkMark)

	// validate the first entry
	t.Log("\tgiven the need to test the first entry in the HAR Log")
	e := l.Entries[0]

	// Date and time stamp of the request start (ISO 8601 - YYYY-MM-DDThh:mm:ss.sTZD).
	if e.StartedDateTime.IsZero() {
		t.Fatal("\t\tshould include a valid startedDateTime", ballotX, e.StartedDateTime)
	}
	t.Log("\t\tshould include a valid startedDateTime", checkMark)

	// Total elapsed time of the request in milliseconds. This is the sum of
	// all timings available in the timings object (i.e. not including -1 values).
	if e.Time != 0 {
		t.Fatal("\t\tshould indicate a total processing time of 0 for all requests", ballotX, e.Time)
	}
	t.Log("\t\tshould indicate a total processing time of 0 for all requests", checkMark)

	// Request method (GET, POST, ...).
	if e.Request.Method != req.Method {
		t.Fatal("\t\tshould have request method of POST", ballotX, e.Request.Method)
	}
	t.Log("\t\tshould have request method of POST", checkMark)

	// Absolute URL of the request (fragments are not included).
	if e.Request.URL != req.URL.String() {
		t.Fatal("\t\tshould contain the correct URL", ballotX, e.Request.URL)
	}
	t.Log("\t\tshould contain the correct URL", checkMark)

	// Request HTTP Version.
	if e.Request.HTTPVersion != req.Proto {
		t.Fatal("\t\tshould contain the HTTP version", ballotX, e.Request.HTTPVersion)
	}
	t.Log("\t\tshould contain the HTTP version", checkMark)

	// List of cookie objects.
	if len(e.Request.Cookies) != 1 {
		t.Fatal("\t\tshould contain one Cookie", ballotX, len(e.Request.Cookies))
	}
	t.Log("\t\tshould contain one Cookie", checkMark)

	// The name of the cookie.
	if e.Request.Cookies[0].Name != c.Name {
		t.Fatal("\t\tshould contain the name of the Cookie", ballotX, e.Request.Cookies[0].Name)
	}
	t.Log("\t\tshould contain the name of the Cookie", checkMark)

	// The value of the cookie.
	if e.Request.Cookies[0].Value != c.Value {
		t.Fatal("\t\tshould contain the value of the Cookie", ballotX, e.Request.Cookies[0].Value)
	}
	t.Log("\t\tshould contain the value of the Cookie", checkMark)

	// List of header objects.
	if len(e.Request.Headers) != 2 {
		t.Fatal("\t\tshould contain two Headers", ballotX, len(e.Request.Headers))
	}
	t.Log("\t\tshould contain two Headers", checkMark)

	hFound := false
	for _, h := range e.Request.Headers {
		if strings.ToLower(h.Name) == "header_key" && h.Value == "header_value" {
			hFound = true
		}
	}
	if hFound == false {
		t.Fatal("\t\tshould contain the custom Header", ballotX)
	}
	t.Fatal("\t\tshould contain the custom Header", checkMark)

}
