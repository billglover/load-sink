package har

import (
	"net"
	"time"
)

// The HAR specification defines an archival format for HTTP transactions.
// Further detail can be found here:
// https://w3c.github.io/web-performance/specs/HAR/Overview.html

// Log represents the root of the exported data.
// This object MUST be present and its name MUST be "log".
type Log struct {
	Version string
	Creator Creator
	Browser Browser `json:"omitempty"`
	Pages   []Page  `json:"omitempty"`
	Entries []Entry
	Comment string `json:"omitempty"`
}

// Creator contains information about the log creator application.
type Creator struct {
	Name    string
	Version string
	Comment string `json:"omitempty"`
}

// Browser contains information about the browser that created the log.
type Browser struct {
	Name    string
	Version string
	Comment string `json:"omitempty"`
}

// Page represents an exported page.
type Page struct {
	StartedDateTime time.Time
	ID              string
	Title           string
	PageTimings     PageTimings
	Comment         string `json:"omitempty"`
}

// PageTimings describes timings for various events (states) fired during the
// page load. All times are specified in milliseconds. If a time info is not
// available appropriate field is set to -1
type PageTimings struct {
	OnContentLoad int    `json:"omitempty"`
	OnLoad        int    `json:"omitempty"`
	Comment       string `json:"omitempty"`
}

// Entry represents an exported HTTP request.
type Entry struct {
	PageRef         string `json:"omitempty"`
	StartedDateTime time.Time
	Time            int
	Request         Request
	Response        Response
	Cache           Cache
	Timings         Timings
	ServerIPAddress net.IP `json:"omitempty"`
	Connection      string `json:"omitempty"`
	Comment         string `json:"omitempty"`
}

// Request contains detailed info about performed request.
type Request struct {
	Method      string
	URL         string
	HTTPVersion string
	Cookies     []Cookie
	Headers     []Header
	QueryString []QueryString
	PostData    PostData `json:"omitempty"`
	HeaderSize  int
	BodySize    int
	TotalSize   int
	Comment     string `json:"omitempty"`
}

// Response contains detailed info about performed response.
type Response struct {
	Status      int
	StatusText  string
	HTTPVersion string
	Cookies     []Cookie
	Headers     []Header
	Content     Content
	RedirectURL string
	HeaderSize  int
	BodySize    int
	TotalSize   int
	Comment     string `json:"omitempty"`
}

// Cookie contains individual cookie information
type Cookie struct {
	Name     string
	Value    string
	Path     string    `json:"omitempty"`
	Domain   string    `json:"omitempty"`
	Expires  time.Time `json:"omitempty"`
	HTTPOnly bool      `json:"omitempty"`
	Secure   bool      `json:"omitempty"`
	Comment  string    `json:"omitempty"`
}

// Header contains individual header information
type Header struct {
	Name    string
	Value   string
	Comment string `json:"omitempty"`
}

// QueryString contains an individual querystring parameter
type QueryString struct {
	Name    string
	Value   string
	Comment string `json:"omitempty"`
}

// PostData describes posted data
type PostData struct {
	MimeType string
	Params   []Param
	Text     string
	Comment  string `json:"omitempty"`
}

// Param contains an individual parameter
type Param struct {
	Name        string
	Value       string
	FileName    string
	ContentType string
	Comment     string `json:"omitempty"`
}

// Content
type Content struct {
	Size        int
	Compression int `json:"omitempty"`
	MimeTye     string
	Text        string `json:"omitempty"`
	Encoding    string `json:"omitempty"`
	Comment     string `json:"omitempty"`
}

// Cache contains info about a request coming from browser cache
type Cache struct {
	BeforeRequest CacheRequest `json:"omitempty"`
	AfterRequest  CacheRequest `json:"omitempty"`
	Comment       string       `json:"omitempty"`
}

// CacheRequest contains details of the cache for a request
type CacheRequest struct {
	Expres     time.Time `json:"omitempty"`
	LastAccess time.Time
	ETag       string
	HitCount   int
	Comment    string `json:"omitempty"`
}

// Timings describes various phases within the request-response round trip
type Timings struct {
	Blocked int `json:"omitempty"`
	DNS     int `json:"omitempty"`
	Connect int `json:"omitempty"`
	Send    int
	Wait    int
	Receive int
	SSL     int    `json:"omitempty"`
	Comment string `json:"omitempty"`
}
