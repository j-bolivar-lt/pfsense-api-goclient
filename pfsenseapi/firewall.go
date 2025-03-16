package pfsenseapi

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/markphelps/optional"
)

const (
	firewallAliasEndpoint    = "api/v2/firewall/alias"
	firewallAliasesEndpoint  = "api/v2/firewall/aliases"
	firewallRuleEndpoint     = "api/v2/firewall/rule"
	firewallRulesEndpoint    = "api/v2/firewall/rules"
	firewallApplyEndpoint    = "api/v2/firewall/apply"
	firewallStateEndpoint    = "api/v2/firewall/states"
	firewallStatesEndpoint   = "api/v2/firewall/states/size"
	firewallScheduleEndpoint = "api/v2/firewall/schedule"
)

// FirewallService provides firewall API methods
type FirewallService service

// FirewallAlias represents a firewall alias
type FirewallAlias struct {
	Name    string   `json:"name"`
	Type    string   `json:"type"`
	Descr   string   `json:"descr"`
	Address []string `json:"address"`
	Detail  []string `json:"detail,omitempty"`
}

type firewallAliasListResponse struct {
	apiResponse
	Data []*FirewallAlias `json:"data"`
}

// ListAliases returns a list of firewall aliases
func (s FirewallService) ListAliases(ctx context.Context) ([]*FirewallAlias, error) {
	response, err := s.client.get(ctx, firewallAliasesEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(firewallAliasListResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

type firewallAliasResponse struct {
	apiResponse
	Data *FirewallAlias `json:"data"`
}

// GetAlias returns a single firewall alias by name
func (s FirewallService) GetAlias(ctx context.Context, name string) (*FirewallAlias, error) {
	response, err := s.client.get(
		ctx,
		firewallAliasEndpoint,
		map[string]string{
			"name": name,
		},
	)
	if err != nil {
		return nil, err
	}

	resp := new(firewallAliasResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// CreateAlias creates a new firewall alias
func (s FirewallService) CreateAlias(ctx context.Context, alias FirewallAlias) (*FirewallAlias, error) {
	jsonData, err := json.Marshal(alias)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	response, err := s.client.post(ctx, firewallAliasEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(firewallAliasResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// UpdateAlias updates an existing firewall alias
func (s FirewallService) UpdateAlias(ctx context.Context, alias FirewallAlias) (*FirewallAlias, error) {
	jsonData, err := json.Marshal(alias)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	response, err := s.client.put(
		ctx,
		firewallAliasEndpoint,
		map[string]string{
			"name": alias.Name, // Use name as ID for update
		},
		jsonData,
	)
	if err != nil {
		return nil, err
	}

	resp := new(firewallAliasResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// DeleteAlias deletes a firewall alias by name
func (s FirewallService) DeleteAlias(ctx context.Context, name string) (*FirewallAlias, error) {
	response, err := s.client.delete(
		ctx,
		firewallAliasEndpoint,
		map[string]string{
			"name": name,
		},
	)
	if err != nil {
		return nil, err
	}

	resp := new(firewallAliasResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// FirewallRule represents a firewall rule
type FirewallRule struct {
	Type            string           `json:"type"`
	Interface       []string         `json:"interface"`
	Ipprotocol      string           `json:"ipprotocol"`
	Protocol        string           `json:"protocol,omitempty"`
	Source          string           `json:"source"`
	Destination     string           `json:"destination"`
	SourcePort      string           `json:"source_port,omitempty"`
	DestinationPort string           `json:"destination_port,omitempty"`
	Descr           string           `json:"descr"`
	Disabled        bool             `json:"disabled"`
	Tracker         int              `json:"tracker,omitempty"` // Use int for tracker
	Log             optional.Bool    `json:"log,omitempty"`
	Tag             optional.String  `json:"tag,omitempty"`
	StatePolicy     optional.String  `json:"statetype,omitempty"`
	Gateway         optional.String  `json:"gateway,omitempty"`
	Quick           optional.Bool    `json:"quick,omitempty"`
	Created         *Time            `json:"created_time,omitempty"`
	Updated         *Time            `json:"updated_time,omitempty"`
}

// Time is a custom time type to handle unmarshalling of unix timestamps
// returned by the pfSense API.
type Time time.Time

type firewallRuleListResponse struct {
	apiResponse
	Data []*FirewallRule `json:"data"`
}

// ListRules returns a list of firewall rules
func (s FirewallService) ListRules(ctx context.Context) ([]*FirewallRule, error) {
	response, err := s.client.get(ctx, firewallRulesEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(firewallRuleListResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

type firewallRuleResponse struct {
	apiResponse
	Data *FirewallRule `json:"data"`
}

// GetRule returns a single firewall rule by tracker ID
func (s FirewallService) GetRule(ctx context.Context, tracker int) (*FirewallRule, error) {
	response, err := s.client.get(
		ctx,
		firewallRuleEndpoint,
		map[string]string{
			"tracker": strconv.Itoa(tracker),
		},
	)
	if err != nil {
		return nil, err
	}

	resp := new(firewallRuleResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// CreateRule creates a new firewall rule
func (s FirewallService) CreateRule(ctx context.Context, rule FirewallRule) (*FirewallRule, error) {
	jsonData, err := json.Marshal(rule)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	response, err := s.client.post(ctx, firewallRuleEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(firewallRuleResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// UpdateRule updates an existing firewall rule
func (s FirewallService) UpdateRule(ctx context.Context, rule FirewallRule) (*FirewallRule, error) {
	jsonData, err := json.Marshal(rule)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	response, err := s.client.put(
		ctx,
		firewallRuleEndpoint,
		map[string]string{
			"tracker": strconv.Itoa(rule.Tracker),
		},
		jsonData,
	)
	if err != nil {
		return nil, err
	}

	resp := new(firewallRuleResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// DeleteRule deletes a firewall rule by tracker ID
func (s FirewallService) DeleteRule(ctx context.Context, tracker int) (*FirewallRule, error) {
	response, err := s.client.delete(
		ctx,
		firewallRuleEndpoint,
		map[string]string{
			"tracker": strconv.Itoa(tracker),
		},
	)
	if err != nil {
		return nil, err
	}

	resp := new(firewallRuleResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// FirewallApply represents the response from applying firewall changes
type FirewallApply struct {
	Applied          bool     `json:"applied"`
	PendingSubsystems []string `json:"pending_subsystems,omitempty"`
}

// Apply applies pending firewall changes
func (s FirewallService) Apply(ctx context.Context) (*FirewallApply, error) {
	response, err := s.client.post(ctx, firewallApplyEndpoint, nil, nil)
	if err != nil {
		return nil, err
	}

	resp := new(struct {
		apiResponse
		Data *FirewallApply `json:"data"`
	})
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// FirewallState represents a firewall state
type FirewallState struct {
	Interface    string `json:"interface"`
	Protocol     string `json:"protocol"`
	Direction    string `json:"direction"`
	Source       string `json:"source"`
	Destination  string `json:"destination"`
	State        string `json:"state"`
	Age          string `json:"age,omitempty"`
	ExpiresIn    string `json:"expires_in,omitempty"`
	PacketsTotal int    `json:"packets_total"`
	PacketsIn    int    `json:"packets_in,omitempty"`
	PacketsOut   int    `json:"packets_out,omitempty"`
	BytesTotal   int    `json:"bytes_total"`
	BytesIn      int    `json:"bytes_in,omitempty"`
	BytesOut     int    `json:"bytes_out,omitempty"`
}

// FirewallStatesSize represents the firewall states size configuration
type FirewallStatesSize struct {
	MaximumStates       int `json:"maximumstates"`
	DefaultMaximumStates int `json:"defaultmaximumstates,omitempty"`
	CurrentStates       int `json:"currentstates,omitempty"`
}

// GetStatesSize returns the current firewall states size configuration
func (s FirewallService) GetStatesSize(ctx context.Context) (*FirewallStatesSize, error) {
	response, err := s.client.get(ctx, firewallStatesEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(struct {
		apiResponse
		Data *FirewallStatesSize `json:"data"`
	})
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// UpdateStatesSize updates the firewall states size configuration
func (s FirewallService) UpdateStatesSize(ctx context.Context, size FirewallStatesSize) (*FirewallStatesSize, error) {
	jsonData, err := json.Marshal(size)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	response, err := s.client.put(ctx, firewallStatesEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(struct {
		apiResponse
		Data *FirewallStatesSize `json:"data"`
	})
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}
