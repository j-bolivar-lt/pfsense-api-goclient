package pfsenseapi

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/j-bolivar-lt/pfsense-api-goclient/v2/pfsenseapi"
	"github.com/stretchr/testify/require"
)

var (
	defaultTimeout = 5 * time.Second

	// noAuthEndpoints is a list of endpoints that require no authentication
	noAuthEndpoints = []string{}

	// localAuthEndpoints is a list of endpoints that always require local
	// authentication. This overrides the default behavior of authenticating with
	// whatever client the Client is constructed with.
	localAuthEndpoints = []string{}

	// responseCodeErrorMap maps HTTP status codes to errors
	responseCodeErrorMap = map[int]error{
		400: errors.New("bad request"),
		401: errors.New("unauthorized"),
		403: errors.New("forbidden"),
		404: errors.New("not found"),
		500: errors.New("internal server error"),
	}
)

// Config provides configuration for the client. These values are only read in
// when NewClient is called.
type Config struct {
	Host string

	LocalAuthEnabled bool
	User             string
	Password         string

	JWTAuthEnabled bool
	JWTToken       string

	TokenAuthEnabled bool
	ApiClientID      string
	ApiClientToken   string

	SkipTLS bool
	Timeout time.Duration
}

func TestClientErrors(t *testing.T) {
	for code, expectedErr := range responseCodeErrorMap {
		t.Run(strconv.Itoa(code), func(t *testing.T) {
			apiRes := new(pfsenseapi.apiResponse)
			apiRes.Code = code
			apiRes.Message = "test message"

			response, err := json.Marshal(apiRes)
			require.NoError(t, err)

			ctx := context.Background()

			handler := func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(code)
				_, err := io.WriteString(w, string(response))
				require.NoError(t, err)
			}

			server := httptest.NewServer(http.HandlerFunc(handler))
			defer server.Close()

			client := NewClientWithNoAuth(server.URL)
			// Since the methods get, post, put, delete are unexported, we can't directly test them.
			// Instead, we test them indirectly through the exported service methods.
			// This requires more setup and potentially different assertions.
			// The following is a placeholder and needs to be replaced with actual service method tests.
			_, err = client.System.GetConfigRaw(ctx)
			require.ErrorIs(t, err, expectedErr)
		})
	}
}
