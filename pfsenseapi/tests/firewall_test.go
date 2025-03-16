package pfsenseapi

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFirewallService_ListRules(t *testing.T) {
	data := mustReadFileString(t, "testdata/multiplefirewallrule.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	rules, err := newClient.Firewall.ListRules(context.Background())
	require.NoError(t, err)
	require.Len(t, rules, 2)

	rules, err = newClient.Firewall.ListRules(context.Background())
	require.Error(t, err)
	require.Nil(t, rules)

	rules, err = newClient.Firewall.ListRules(context.Background())
	require.Error(t, err)
	require.Nil(t, rules)
}

func TestFirewallService_GetRule(t *testing.T) {
	data := mustReadFileString(t, "testdata/singlefirewallrule.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	rule, err := newClient.Firewall.GetRule(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, rule)

	rule, err = newClient.Firewall.GetRule(context.Background(), 1)
	require.Error(t, err)
	require.Nil(t, rule)

	rule, err = newClient.Firewall.GetRule(context.Background(), 1)
	require.Error(t, err)
	require.Nil(t, rule)
}

func TestFirewallService_CreateRule(t *testing.T) {
	data := mustReadFileString(t, "testdata/singlefirewallrule.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	newRule := pfsenseapi.FirewallRule{
		Type:        "pass",
		Interface:   []string{"wan"},
		Ipprotocol:  "inet",
		Protocol:    "tcp",
		Source:      "any",
		Destination: "any",
		Descr:       "Test Rule",
	}
	rule, err := newClient.Firewall.CreateRule(context.Background(), newRule)
	require.NoError(t, err)
	require.NotNil(t, rule)

	rule, err = newClient.Firewall.CreateRule(context.Background(), newRule)
	require.Error(t, err)
	require.Nil(t, rule)

	rule, err = newClient.Firewall.CreateRule(context.Background(), newRule)
	require.Error(t, err)
	require.Nil(t, rule)
}

func TestFirewallService_UpdateRule(t *testing.T) {
	data := mustReadFileString(t, "testdata/singlefirewallrule.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	updatedRule := pfsenseapi.FirewallRule{
		Tracker:     1,
		Type:        "pass",
		Interface:   []string{"wan"},
		Ipprotocol:  "inet",
		Protocol:    "tcp",
		Source:      "any",
		Destination: "any",
		Descr:       "Updated Test Rule",
	}
	rule, err := newClient.Firewall.UpdateRule(context.Background(), updatedRule)
	require.NoError(t, err)
	require.NotNil(t, rule)

	rule, err = newClient.Firewall.UpdateRule(context.Background(), updatedRule)
	require.Error(t, err)
	require.Nil(t, rule)

	rule, err = newClient.Firewall.UpdateRule(context.Background(), updatedRule)
	require.Error(t, err)
	require.Nil(t, rule)
}

func TestFirewallService_DeleteRule(t *testing.T) {
	data := mustReadFileString(t, "testdata/singlefirewallrule.json")
	server := setupTestServer(t, data)
	defer server.Close()

	newClient := NewClientWithNoAuth(server.URL)
	rule, err := newClient.Firewall.DeleteRule(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, rule)

	rule, err = newClient.Firewall.DeleteRule(context.Background(), 1)
	require.Error(t, err)
	require.Nil(t, rule)

	rule, err = newClient.Firewall.DeleteRule(context.Background(), 1)
	require.Error(t, err)
	require.Nil(t, rule)
}
