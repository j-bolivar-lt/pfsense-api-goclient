package pfsenseapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSystemService_GetConfigRaw(t *testing.T) {
	data := `{
		"status": "ok",
		"code": 200,
		"return": 0,
		"message": "Configuration fetched successfully",
		"data": {
			"config_text": "PGNvbmZpZz4KICA8c3lzdGVtPjwv\nc3lzdGVtPgo8L2NvbmZpZz4K"
		}
	}`
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	config, err := newClient.System.GetConfigRaw(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, config)
    require.Equal(t, "PGNvbmZpZz4KICA8c3lzdGVtPjwv\nc3lzdGVtPgo8L2NvbmZpZz4K", config)

	config, err = newClient.System.GetConfigRaw(context.Background())
	require.Error(t, err)
	require.Empty(t, config)

	config, err = newClient.System.GetConfigRaw(context.Background())
	require.Error(t, err)
	require.Empty(t, config)
}
