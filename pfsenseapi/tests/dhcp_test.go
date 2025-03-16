package pfsenseapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDHCPService_ListLeases(t *testing.T) {
	data := mustReadFileString(t, "testdata/multipledhcplease.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	leases, err := newClient.DHCP.ListLeases(context.Background())
	require.NoError(t, err)
	require.Len(t, leases, 2)
    require.Equal(t, "192.168.1.100", leases[0].IP)
    require.Equal(t, "host2", leases[1].Hostname)
}

func TestDHCPService_ListStaticMappings(t *testing.T) {
	data := mustReadFileString(t, "testdata/multiplestaticmapping.json") // Use the new file
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	mappings, err := newClient.DHCP.ListStaticMappings(context.Background())
	require.NoError(t, err)
	require.Len(t, mappings, 2)
    require.Equal(t, "00:11:22:33:44:55", mappings[0].MAC)
    require.Equal(t, "static-host-2", mappings[1].Hostname)
}

func TestDHCPService_GetStaticMapping(t *testing.T) {
	data := mustReadFileString(t, "testdata/singlestaticmapping.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	mapping, err := newClient.DHCP.GetStaticMapping(context.Background(), "example")
	require.NoError(t, err)
	require.NotNil(t, mapping)
	require.Equal(t, "00:11:22:33:44:55", mapping.MAC)
	require.Equal(t, "192.168.1.10", mapping.IPAddr)
	require.Equal(t, "static-host-1", mapping.Hostname)
}

func TestDHCPService_CreateStaticMapping(t *testing.T) {
	data := mustReadFileString(t, "testdata/singlestaticmapping.json")
	server := setupTestServer(t, data)
	defer server.Close()
	newClient := NewClientWithNoAuth(server.URL)
	mapping, err := newClient.DHCP.GetStaticMapping(context.Background(), "example")
	require.NoError(t, err)
	require.NotNil(t, mapping)
}
