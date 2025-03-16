package pfsenseapi

import (
	"context"
	"crypto/tls"
	"net/http"
	"testing"

	"github.com/j-bolivar-lt/pfsense-api-goclient/v2/pfsenseapi"
	"github.com/stretchr/testify/require"
)

func NewClientWithNoAuth(host string) *pfsenseapi.Client {
	config := pfsenseapi.Config{
		Host:    host,
		SkipTLS: true,
		Timeout: defaultTimeout,
	}

    newClient := pfsenseapi.NewClient(config)
    // We can't directly set the unexported field, so we'll use the existing System service
    return newClient
}

func TestBridgeService_ListInterfaceBridges(t *testing.T) {
    data := mustReadFileString(t, "testdata/multipleinterfacebridge.json")
    server := setupTestServer(t, data)  // Add this line to initialize the server
    defer server.Close()

    newClient := NewClientWithNoAuth(server.URL)
    bridges, err := newClient.Bridge.ListInterfaceBridges(context.Background())
	require.NoError(t, err)
	require.Len(t, bridges, 2)
	require.Equal(t, "test_bridge", bridges[0].Id)
	require.Equal(t, "bridge0", bridges[0].Bridgeif)
	require.Equal(t, "test_bridge_2", bridges[1].Id)
	require.Equal(t, "bridge1", bridges[1].Bridgeif)

	bridges, err = newClient.Bridge.ListInterfaceBridges(context.Background())
	require.Error(t, err)
	require.Nil(t, bridges)

	bridges, err = newClient.Bridge.ListInterfaceBridges(context.Background())
	require.Error(t, err)
	require.Nil(t, bridges)
}

func TestBridgeService_GetInterfaceBridge(t *testing.T) {
	data := mustReadFileString(t, "testdata/singleinterfacebridge.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	bridge, err := newClient.Bridge.GetInterfaceBridge(context.Background(), "test_bridge")
	require.NoError(t, err)
	require.NotNil(t, bridge)
	require.Equal(t, "test_bridge", bridge.Id)
	require.Equal(t, "bridge0", bridge.Bridgeif)
	require.Equal(t, "Test Bridge", bridge.Descr)
	require.ElementsMatch(t, []string{"em1", "em2"}, bridge.Members)

	bridge, err = newClient.Bridge.GetInterfaceBridge(context.Background(), "test_bridge")
	require.Error(t, err)
	require.Nil(t, bridge)

	bridge, err = newClient.Bridge.GetInterfaceBridge(context.Background(), "test_bridge")
	require.Error(t, err)
	require.Nil(t, bridge)
}

func TestBridgeService_CreateInterfaceBridge(t *testing.T) {
	data := mustReadFileString(t, "testdata/singleinterfacebridge.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	newBridge := pfsenseapi.InterfaceBridgeRequest{
		Members:  []string{"em1", "em2"},
		Descr:    "Test Bridge",
		Bridgeif: "bridge0",
	}
	bridge, err := newClient.Bridge.CreateInterfaceBridge(context.Background(), newBridge)
	require.NoError(t, err)
	require.NotNil(t, bridge)
	require.Equal(t, "test_bridge", bridge.Id)
	require.Equal(t, "bridge0", bridge.Bridgeif)
	require.Equal(t, "Test Bridge", bridge.Descr)
	require.ElementsMatch(t, []string{"em1", "em2"}, bridge.Members)

	bridge, err = newClient.Bridge.CreateInterfaceBridge(context.Background(), newBridge)
	require.Error(t, err)
	require.Nil(t, bridge)

	bridge, err = newClient.Bridge.CreateInterfaceBridge(context.Background(), newBridge)
	require.Error(t, err)
	require.Nil(t, bridge)
}

func TestBridgeService_UpdateInterfaceBridge(t *testing.T) {
	data := mustReadFileString(t, "testdata/singleinterfacebridge.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	updatedBridge := pfsenseapi.InterfaceBridgeRequest{
		Members:  []string{"em1", "em2"},
		Descr:    "Updated Test Bridge",
		Bridgeif: "bridge0",
	}
	bridge, err := newClient.Bridge.UpdateInterfaceBridge(context.Background(), "test_bridge", updatedBridge)
	require.NoError(t, err)
	require.NotNil(t, bridge)
	require.Equal(t, "test_bridge", bridge.Id)
	require.Equal(t, "bridge0", bridge.Bridgeif)
	require.Equal(t, "Test Bridge", bridge.Descr) // Note: The test server returns the original data, not the updated one

	bridge, err = newClient.Bridge.UpdateInterfaceBridge(context.Background(), "test_bridge", updatedBridge)
	require.Error(t, err)
	require.Nil(t, bridge)

	bridge, err = newClient.Bridge.UpdateInterfaceBridge(context.Background(), "test_bridge", updatedBridge)
	require.Error(t, err)
	require.Nil(t, bridge)
}

func TestBridgeService_DeleteInterfaceBridge(t *testing.T) {
	data := mustReadFileString(t, "testdata/singleinterfacebridge.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	bridge, err := newClient.Bridge.DeleteInterfaceBridge(context.Background(), "test_bridge")
	require.NoError(t, err)
	require.NotNil(t, bridge)
	require.Equal(t, "test_bridge", bridge.Id)
	require.Equal(t, "bridge0", bridge.Bridgeif)
	require.Equal(t, "Test Bridge", bridge.Descr)
	require.ElementsMatch(t, []string{"em1", "em2"}, bridge.Members)

	bridge, err = newClient.Bridge.DeleteInterfaceBridge(context.Background(), "test_bridge")
	require.Error(t, err)
	require.Nil(t, bridge)

	bridge, err = newClient.Bridge.DeleteInterfaceBridge(context.Background(), "test_bridge")
	require.Error(t, err)
	require.Nil(t, bridge)
}
