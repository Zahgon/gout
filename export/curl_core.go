package export

import (
	"io"
	"net/http"
)

type curl struct {
	Header []string
	Method string
	Data   string
	URL    string

	FormData []string
}

const boundary = "boundary="

func isExists(path string) bool { _ = "STUB: not implemented"; return false }

func getFileName(fName string) string { _ = "STUB: not implemented"; return "" }

func (c *curl) formData(req *http.Request) error { _ = "STUB: not implemented"; return nil }

func (c *curl) header(req *http.Request) { _ = "STUB: not implemented"; return }

// GenCurl used to generate curl commands
func GenCurl(req *http.Request, long bool, w io.Writer) error {
	_ = "STUB: not implemented"
	return nil
}
