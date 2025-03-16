package pfsenseapi

import (
	"context"
	"encoding/json"
	"fmt"
)

const (
	arpTableEndpoint        = "api/v2/diagnostics/arp_table"
	commandPromptEndpoint   = "api/v2/diagnostics/command_prompt"
	configHistoryEndpoint = "api/v2/diagnostics/config_history/revision"
	configHistoryRevisionsEndpoint = "api/v2/diagnostics/config_history/revisions"
	haltSystemEndpoint = "api/v2/diagnostics/halt_system"
	rebootSystemEndpoint = "api/v2/diagnostics/reboot"
)

// DiagnosticsService provides diagnostics API methods
type DiagnosticsService service

// ARPTable represents an ARP table entry
type ARPTable struct {
	Hostname    string `json:"hostname,omitempty"`
	IPAddress   string `json:"ip_address,omitempty"`
	MACAddress  string `json:"mac_address,omitempty"`
	Interface   string `json:"interface,omitempty"`
	Type        string `json:"type,omitempty"`
	Permanent   bool   `json:"permanent,omitempty"`
	DNSResolve  string `json:"dnsresolve,omitempty"`
	Expires     string `json:"expires,omitempty"`
}

type arpTableListResponse struct {
	apiResponse
	Data []*ARPTable `json:"data"`
}

// ListARPEntries returns a list of ARP table entries
func (s DiagnosticsService) ListARPEntries(ctx context.Context) ([]*ARPTable, error) {
	response, err := s.client.get(ctx, arpTableEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(arpTableListResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// CommandPrompt represents a command prompt execution
type CommandPrompt struct {
	Command    string `json:"command"`
	Output     string `json:"output,omitempty"`
	ResultCode int    `json:"result_code,omitempty"`
}

// ExecuteCommand executes a command prompt
func (s DiagnosticsService) ExecuteCommand(ctx context.Context, command string) (*CommandPrompt, error) {
	cmd := CommandPrompt{
		Command: command,
	}

	jsonData, err := json.Marshal(cmd)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	response, err := s.client.post(ctx, commandPromptEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(struct {
		apiResponse
		Data *CommandPrompt `json:"data"`
	})
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// ConfigHistoryRevision represents a config history revision
type ConfigHistoryRevision struct {
	Time        int    `json:"time"`
	Description string `json:"description"`
	Version     string `json:"version"`
	FileSize    int    `json:"filesize"`
}

type configHistoryListResponse struct {
	apiResponse
	Data []*ConfigHistoryRevision `json:"data"`
}

// ListConfigHistoryRevisions returns a list of config history revisions
func (s DiagnosticsService) ListConfigHistoryRevisions(ctx context.Context) ([]*ConfigHistoryRevision, error) {
	response, err := s.client.get(ctx, configHistoryRevisionsEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(configHistoryListResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

type configHistoryResponse struct {
	apiResponse
	Data *ConfigHistoryRevision `json:"data"`
}

// GetConfigHistoryRevision returns a config history revision by ID
func (s DiagnosticsService) GetConfigHistoryRevision(ctx context.Context, id string) (*ConfigHistoryRevision, error) {
	response, err := s.client.get(
		ctx,
		configHistoryEndpoint,
		map[string]string{
			"id": id,
		},
	)
	if err != nil {
		return nil, err
	}

	resp := new(configHistoryResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}
