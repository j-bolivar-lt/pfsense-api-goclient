package pfsenseapi

import (
	"context"
	"encoding/json"
	"fmt"
)

const (
	dhcpServerEndpoint  = "api/v2/services/dhcpserver"
	dhcpServersEndpoint = "api/v2/services/dhcpserver/interfaces" // Plural for listing
	dhcpLeasesEndpoint  = "api/v2/services/dhcp/leases"
	dhcpApplyEndpoint   = "api/v2/services/dhcp/apply"
    dhcpStaticMappingEndpoint = "api/v2/services/dhcpserver/static_mapping" //singular
    dhcpStaticMappingsEndpoint = "api/v2/services/dhcpserver/static_mappings" //plural
)

// DHCPService provides DHCP API methods
type DHCPService service

// DHCPServer represents a DHCP server configuration
type DHCPServer struct {
	Interface       string                     `json:"interface"`
	Enable          bool                       `json:"enable"`
	RangeFrom       string                     `json:"range_from"`
	RangeTo         string                     `json:"range_to"`
	Domain          string                     `json:"domain,omitempty"`
	DefaultLeasetime int                       `json:"defaultleasetime,omitempty"`
	MaxLeasetime    int                        `json:"maxleasetime,omitempty"`
	Gateway         string                     `json:"gateway,omitempty"`
	DNSServer       []string                   `json:"dnsserver,omitempty"`
	StaticMappings  []DHCPServerStaticMapping  `json:"staticmap,omitempty"`
}

// DHCPServerStaticMapping represents a static DHCP mapping
type DHCPServerStaticMapping struct {
	MAC       string   `json:"mac"`
	IPAddr    string   `json:"ipaddr,omitempty"`
	Hostname  string   `json:"hostname,omitempty"`
	Domain    string   `json:"domain,omitempty"`
	Gateway   string   `json:"gateway,omitempty"`
	DNSServer []string `json:"dnsserver,omitempty"`
	Descr     string   `json:"descr,omitempty"`
}

type dhcpServerListResponse struct {
	apiResponse
	Data []*DHCPServer `json:"data"`
}

// ListDHCPServers returns a list of DHCP servers
func (s DHCPService) ListDHCPServers(ctx context.Context) ([]*DHCPServer, error) {
	response, err := s.client.get(ctx, dhcpServersEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(dhcpServerListResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

type dhcpServerResponse struct {
	apiResponse
	Data *DHCPServer `json:"data"`
}

// GetDHCPServer returns a DHCP server configuration for a specific interface
func (s DHCPService) GetDHCPServer(ctx context.Context, interfaceID string) (*DHCPServer, error) {
	response, err := s.client.get(
		ctx,
		dhcpServerEndpoint,
		map[string]string{
			"interface": interfaceID,
		},
	)
	if err != nil {
		return nil, err
	}

	resp := new(dhcpServerResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// CreateDHCPServer creates a new DHCP server configuration
func (s DHCPService) CreateDHCPServer(ctx context.Context, server DHCPServer) (*DHCPServer, error) {
	jsonData, err := json.Marshal(server)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	response, err := s.client.post(ctx, dhcpServerEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(dhcpServerResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// UpdateDHCPServer updates an existing DHCP server configuration
func (s DHCPService) UpdateDHCPServer(ctx context.Context, server DHCPServer) (*DHCPServer, error) {
	jsonData, err := json.Marshal(server)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	response, err := s.client.put(
		ctx,
		dhcpServerEndpoint,
		map[string]string{
			"interface": server.Interface,
		},
		jsonData,
	)
	if err != nil {
		return nil, err
	}

	resp := new(dhcpServerResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// DeleteDHCPServer deletes a DHCP server configuration
func (s DHCPService) DeleteDHCPServer(ctx context.Context, interfaceID string) (*DHCPServer, error) {
	response, err := s.client.delete(
		ctx,
		dhcpServerEndpoint,
		map[string]string{
			"interface": interfaceID,
		},
	)
	if err != nil {
		return nil, err
	}

	resp := new(dhcpServerResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// DHCPLease represents a DHCP lease
type DHCPLease struct {
	IP           string `json:"ip"`
	MAC          string `json:"mac"`
	Hostname     string `json:"hostname,omitempty"`
	Interface    string `json:"if,omitempty"`
	Starts       string `json:"starts,omitempty"`
	Ends         string `json:"ends,omitempty"`
	ActiveStatus string `json:"active_status,omitempty"`
	OnlineStatus string `json:"online_status,omitempty"`
	Descr        string `json:"descr,omitempty"`
}

// ListLeases returns a list of DHCP leases
func (s DHCPService) ListLeases(ctx context.Context) ([]*DHCPLease, error) {
	response, err := s.client.get(ctx, dhcpLeasesEndpoint, nil)
	if err != nil {
		return nil, err
	}

    resp := new(struct {
        apiResponse
        Data struct {
            Leases []*DHCPLease `json:"leases"` // Correct nesting
        } `json:"data"`
    })
    if err = json.Unmarshal(response, resp); err != nil {
        return nil, fmt.Errorf("error unmarshalling response: %w", err)
    }

    return resp.Data.Leases, nil
}

// DHCPApply represents the response from applying DHCP changes
type DHCPApply struct {
	Applied bool `json:"applied"`
}

// Apply applies pending DHCP changes
func (s DHCPService) Apply(ctx context.Context) (*DHCPApply, error) {
	response, err := s.client.post(ctx, dhcpApplyEndpoint, nil, nil)
	if err != nil {
		return nil, err
	}

	resp := new(struct {
		apiResponse
		Data *DHCPApply `json:"data"`
	})
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

type dhcpStaticMappingListResponse struct {
	apiResponse
	Data []*DHCPServerStaticMapping `json:"data"` // Corrected: Directly an array
}

// ListStaticMappings returns a list of DHCP static mappings
func (s DHCPService) ListStaticMappings(ctx context.Context) ([]*DHCPServerStaticMapping, error) {
	response, err := s.client.get(ctx, dhcpStaticMappingsEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(dhcpStaticMappingListResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

type dhcpStaticMappingResponse struct {
	apiResponse
	Data *DHCPServerStaticMapping `json:"data"`
}

func (s *DHCPService) GetStaticMapping(ctx context.Context, id string) (*DHCPServerStaticMapping, error) {
	response, err := s.client.get(ctx, dhcpStaticMappingEndpoint, map[string]string{"id": id})
	if err != nil {
		return nil, err
	}
	resp := new(dhcpStaticMappingResponse)

	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	if resp.Code != 200 {
		return nil, fmt.Errorf("non 2xx response code received: %d", resp.Code)
	}

	return resp.Data, nil
}
