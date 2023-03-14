package send

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandleHTTPResponseError(t *testing.T) {
	for _, test := range []struct {
		name     string
		resp     *http.Response
		hasErr   bool
		contains string
	}{
		{
			name: "NilResponse",
		},
		{
			name: "100StatusCode",
			resp: &http.Response{
				StatusCode: http.StatusContinue,
				Body:       io.NopCloser(strings.NewReader("body")),
			},
		},
		{
			name: "200StatusCode",
			resp: &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader("body")),
			},
		},
		{
			name: "308StatusCode",
			resp: &http.Response{
				StatusCode: http.StatusPermanentRedirect,
				Body:       io.NopCloser(strings.NewReader("body")),
			},
		},
		{
			name: "400StatusCode",
			resp: &http.Response{
				StatusCode: http.StatusBadRequest,
				Body:       io.NopCloser(strings.NewReader("invalid request")),
			},
			hasErr:   true,
			contains: "invalid request",
		},
		{
			name: "500StatusCode",
			resp: &http.Response{
				StatusCode: http.StatusInternalServerError,
				Body:       io.NopCloser(strings.NewReader("internal error")),
			},
			hasErr:   true,
			contains: "internal error",
		},
		{
			name: "511StatusCode",
			resp: &http.Response{
				StatusCode: http.StatusNetworkAuthenticationRequired,
				Body:       io.NopCloser(strings.NewReader("auth required")),
			},
			hasErr:   true,
			contains: "auth required",
		},
		{
			name: "ReaderError",
			resp: &http.Response{
				StatusCode: http.StatusInternalServerError,
				Body:       io.NopCloser(&errorReader{}),
			},
			hasErr:   true,
			contains: "failed to read response body",
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			err := handleHTTPResponseError(test.resp)
			if test.hasErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), test.contains)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

type errorReader struct{}

func (r *errorReader) Read(_ []byte) (int, error) {
	return 0, errors.New("read error")
}
