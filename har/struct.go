package har

import (
	"net"
	"time"
)

// The HAR specification defines an archival format for HTTP transactions.
// Further detail can be found here:
// https://w3c.github.io/web-performance/specs/HAR/Overview.html
type HAR struct {
	Log Log `json:"log,omitempty"`
}

// Log represents the root of the exported data.
// This object MUST be present and its name MUST be "log".
type Log struct {
	Version string  `json:"version"`
	Creator Creator `json:"creator"`
	Browser Browser `json:"browser,omitempty"`
	Pages   []Page  `json:"pages,omitempty"`
	Entries []Entry `json:"entries"`
	Comment string  `json:"comment,omitempty"`
}

// Creator contains information about the log creator application.
type Creator struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Comment string `json:"comment,omitempty"`
}

// Browser contains information about the browser that created the log.
type Browser struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Comment string `json:"comment,omitempty"`
}

// Page represents an exported page.
type Page struct {
	StartedDateTime time.Time   `json:"startedDateTime"`
	ID              string      `json:"id"`
	Title           string      `json:"title"`
	PageTimings     PageTimings `json:"pageTimings"`
	Comment         string      `json:"comment,omitempty"`
}

// PageTimings describes timings for various events (states) fired during the
// page load. All times are specified in milliseconds. If a time info is not
// available appropriate field is set to -1
type PageTimings struct {
	OnContentLoad int    `json:"onContentLoad,omitempty"`
	OnLoad        int    `json:"onLoad,omitempty"`
	Comment       string `json:"comment,omitempty"`
}

// Entry represents an exported HTTP request.
type Entry struct {
	PageRef         string    `json:"pageRef,omitempty"`
	StartedDateTime time.Time `json:"startedDateTime"`
	Time            int       `json:"time"`
	Request         Request   `json:"request"`
	Response        Response  `json:"response"`
	Cache           Cache     `json:"cache"`
	Timings         Timings   `json:"timings"`
	ServerIPAddress net.IP    `json:"serverIPAddress,omitempty"`
	Connection      string    `json:"connection,omitempty"`
	Comment         string    `json:"comment,omitempty"`
}

// Request contains detailed info about performed request.
type Request struct {
	Method      string        `json:"method"`
	URL         string        `json:"url"`
	HTTPVersion string        `json:"httpVersion"`
	Cookies     []Cookie      `json:"cookies"`
	Headers     []Header      `json:"headers"`
	QueryString []QueryString `json:"queryString"`
	PostData    PostData      `json:"postData,omitempty"`
	HeadersSize int           `json:"headersSize"`
	BodySize    int           `json:"bodySize"`
	TotalSize   int           `json:"totalSize"`
	Comment     string        `json:"comment,omitempty"`
}

// Response contains detailed info about performed response.
type Response struct {
	Status      int      `json:"status"`
	StatusText  string   `json:"statusText"`
	HTTPVersion string   `json:"httpVersion"`
	Cookies     []Cookie `json:"cookies"`
	Headers     []Header `json:"headers"`
	Content     Content  `json:"content"`
	RedirectURL string   `json:"redirectURL"`
	HeadersSize int      `json:"headersSize"`
	BodySize    int      `json:"bodySize"`
	TotalSize   int      `json:"totalSize"`
	Comment     string   `json:"comment,omitempty"`
}

// Cookie contains individual cookie information
type Cookie struct {
	Name     string    `json:"name"`
	Value    string    `json:"value"`
	Path     string    `json:"path,omitempty"`
	Domain   string    `json:"domain,omitempty"`
	Expires  time.Time `json:"expires,omitempty"`
	HTTPOnly bool      `json:"httpOnly,omitempty"`
	Secure   bool      `json:"secure,omitempty"`
	Comment  string    `json:"comment,omitempty"`
}

// Header contains individual header information
type Header struct {
	Name    string `json:"name"`
	Value   string `json:"value"`
	Comment string `json:"comment,omitempty"`
}

// QueryString contains an individual querystring parameter
type QueryString struct {
	Name    string `json:"name"`
	Value   string `json:"value"`
	Comment string `json:"comment,omitempty"`
}

// PostData describes posted data
type PostData struct {
	MimeType string  `json:"mimeType"`
	Params   []Param `json:"params"`
	Text     string  `json:"text"`
	Comment  string  `json:"comment,omitempty"`
}

// Param contains an individual parameter
type Param struct {
	Name        string `json:"name"`
	Value       string `json:"value"`
	FileName    string `json:"fileName"`
	ContentType string `json:"contentType"`
	Comment     string `json:"comment,omitempty"`
}

// Content
type Content struct {
	Size        int    `json:"size"`
	Compression int    `json:"compression,omitempty"`
	MimeTye     string `json:"mimeType"`
	Text        string `json:"text,omitempty"`
	Encoding    string `json:"encoding,omitempty"`
	Comment     string `json:"comment,omitempty"`
}

// Cache contains info about a request coming from browser cache
type Cache struct {
	BeforeRequest CacheRequest `json:"beforeRequest,omitempty"`
	AfterRequest  CacheRequest `json:"afterRequest,omitempty"`
	Comment       string       `json:"comment,omitempty"`
}

// CacheRequest contains details of the cache for a request
type CacheRequest struct {
	Expires    time.Time `json:"expires,omitempty"`
	LastAccess time.Time `json:"lastAccess"`
	ETag       string    `json:"eTag"`
	HitCount   int       `json:"hitCount"`
	Comment    string    `json:"comment,omitempty"`
}

// Timings describes various phases within the request-response round trip
type Timings struct {
	Blocked int    `json:"blocked,omitempty"`
	DNS     int    `json:"dns,omitempty"`
	Connect int    `json:"connect,omitempty"`
	Send    int    `json:"send"`
	Wait    int    `json:"wait"`
	Receive int    `json:"receive"`
	SSL     int    `json:"ssl,omitempty"`
	Comment string `json:"comment,omitempty"`
}
