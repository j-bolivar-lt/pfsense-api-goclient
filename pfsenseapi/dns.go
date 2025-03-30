package pfsenseapi

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
)

const (
	dnsResolverEndpoint               = "api/v2/services/unbound"
	dnsResolverApplyEndpoint          = "api/v2/services/unbound/apply"
	dnsResolverHostOverrideEndpoint   = "api/v2/services/dns_resolver/host_override"
	dnsResolverHostOverrideEndpoints  = "api/v2/services/dns_resolver/host_overrides"
	dnsResolverDomainOverrideEndpoint = "api/v2/services/dns_resolver/domain_overrides"
	dnsForwarderEndpoint              = "api/v2/services/dnsforwarder"
	dnsForwarderApplyEndpoint         = "api/v2/services/dnsforwarder/apply"
	dnsForwarderHostOverrideEndpoint  = "api/v2/services/dnsforwarder/host_override"
)

// DNSService provides DNS API methods
type DNSService service

// DNSResolverSettings represents the DNS resolver settings
type DNSResolverSettings struct {
	Enable            bool     `json:"enable"`
	Port              string   `json:"port,omitempty"`
	EnableSSL         bool     `json:"enablessl"`
	SSLCertRef        string   `json:"sslcertref,omitempty"`
	TLSPort           string   `json:"tlsport,omitempty"`
	ActiveInterface   []string `json:"active_interface,omitempty"`
	OutgoingInterface []string `json:"outgoing_interface,omitempty"`
	DNSSEC            bool     `json:"dnssec"`
	Forwarding        bool     `json:"forwarding"`
	RegDHCP           bool     `json:"regdhcp"`
	RegDHCPStatic     bool     `json:"regdhcpstatic"`
	RegOpenVPNClients bool     `json:"regovpnclients"`
	CustomOptions     string   `json:"custom_options,omitempty"`
}

// GetDNSResolverSettings returns the DNS resolver settings
func (s DNSService) GetDNSResolverSettings(ctx context.Context) (*DNSResolverSettings, error) {
	response, err := s.client.get(ctx, dnsResolverEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(struct {
		apiResponse
		Data *DNSResolverSettings `json:"data"`
	})
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// UpdateDNSResolverSettings updates the DNS resolver settings
func (s DNSService) UpdateDNSResolverSettings(ctx context.Context, settings DNSResolverSettings) (*DNSResolverSettings, error) {
	jsonData, err := json.Marshal(settings)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	response, err := s.client.put(ctx, dnsResolverEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(struct {
		apiResponse
		Data *DNSResolverSettings `json:"data"`
	})
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// DNSResolverApply represents the response from applying DNS resolver changes
type DNSResolverApply struct {
	Applied bool `json:"applied"`
}

// ApplyDNSResolver applies pending DNS resolver changes
func (s DNSService) ApplyDNSResolver(ctx context.Context) (*DNSResolverApply, error) {
	response, err := s.client.post(ctx, dnsResolverApplyEndpoint, nil, nil)
	if err != nil {
		return nil, err
	}

	resp := new(struct {
		apiResponse
		Data *DNSResolverApply `json:"data"`
	})
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// DNSResolverHostOverride represents a DNS resolver host override
type DNSResolverHostOverride struct {
	Id      int                            `json:"id,omitempty"`
	Host    string                         `json:"host"`
	Domain  string                         `json:"domain"`
	IP      []string                       `json:"ip"`
	Descr   string                         `json:"descr,omitempty"`
	Aliases []DNSResolverHostOverrideAlias `json:"aliases,omitempty"`
}

// DNSResolverHostOverrideAlias represents a DNS resolver host override alias
type DNSResolverHostOverrideAlias struct {
	Host   string `json:"host"`
	Domain string `json:"domain"`
	Descr  string `json:"descr,omitempty"`
}

// GetDNSResolverHostOverrides returns the DNS resolver host overrides
func (s DNSService) GetDNSResolverHostOverrides(ctx context.Context) ([]*DNSResolverHostOverride, error) {
	response, err := s.client.get(ctx, dnsResolverHostOverrideEndpoints, nil)
	if err != nil {
		return nil, err
	}

	resp := new(struct {
		apiResponse
		Data []*DNSResolverHostOverride `json:"data"`
	})
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// CreateDNSResolverHostOverride creates a new DNS resolver host override
func (s DNSService) DeleteDNSResolverHostOverride(ctx context.Context, override DNSResolverHostOverride) (*DNSResolverHostOverride, error) {

	queryMap := make(map[string]string)

	queryMap["id"] = strconv.Itoa(override.Id)
	queryMap["apply"] = "false"

	response, err := s.client.delete(ctx, dnsResolverHostOverrideEndpoint, queryMap)

	if err != nil {
		return nil, err
	}

	resp := new(struct {
		apiResponse
		Data *DNSResolverHostOverride `json:"data"`
	})
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// CreateDNSResolverHostOverride creates a new DNS resolver host override
func (s DNSService) CreateDNSResolverHostOverride(ctx context.Context, override DNSResolverHostOverride) (*DNSResolverHostOverride, error) {
	jsonData, err := json.Marshal(override)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	response, err := s.client.post(ctx, dnsResolverHostOverrideEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(struct {
		apiResponse
		Data *DNSResolverHostOverride `json:"data"`
	})
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// DNSResolverDomainOverride represents a DNS resolver domain override
type DNSResolverDomainOverride struct {
	Domain             string `json:"domain"`
	IP                 string `json:"ip"`
	Descr              string `json:"descr,omitempty"`
	ForwardTLSUpstream bool   `json:"forward_tls_upstream"`
	TLSHostname        string `json:"tls_hostname,omitempty"`
}

// GetDNSResolverDomainOverrides returns the DNS resolver domain overrides
func (s DNSService) GetDNSResolverDomainOverrides(ctx context.Context) ([]*DNSResolverDomainOverride, error) {
	response, err := s.client.get(ctx, dnsResolverDomainOverrideEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(struct {
		apiResponse
		Data []*DNSResolverDomainOverride `json:"data"`
	})
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// CreateDNSResolverDomainOverride creates a new DNS resolver domain override
func (s DNSService) CreateDNSResolverDomainOverride(ctx context.Context, override DNSResolverDomainOverride) (*DNSResolverDomainOverride, error) {
	jsonData, err := json.Marshal(override)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	response, err := s.client.post(ctx, dnsResolverDomainOverrideEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(struct {
		apiResponse
		Data *DNSResolverDomainOverride `json:"data"`
	})
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// DNSForwarderSettings represents the DNS forwarder settings
type DNSForwarderSettings struct {
	Enable        bool     `json:"enable"`
	Interfaces    []string `json:"interfaces,omitempty"`
	Port          string   `json:"port,omitempty"`
	DNSServers    []string `json:"dnsservers,omitempty"`
	StrictOrder   bool     `json:"strict_order"`
	RegDHCP       bool     `json:"regdhcp"`
	RegDHCPStatic bool     `json:"regdhcpstatic"`
}

// GetDNSForwarderSettings returns the DNS forwarder settings
func (s DNSService) GetDNSForwarderSettings(ctx context.Context) (*DNSForwarderSettings, error) {
	response, err := s.client.get(ctx, dnsForwarderEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(struct {
		apiResponse
		Data *DNSForwarderSettings `json:"data"`
	})
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// UpdateDNSForwarderSettings updates the DNS forwarder settings
func (s DNSService) UpdateDNSForwarderSettings(ctx context.Context, settings DNSForwarderSettings) (*DNSForwarderSettings, error) {
	jsonData, err := json.Marshal(settings)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	response, err := s.client.put(ctx, dnsForwarderEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(struct {
		apiResponse
		Data *DNSForwarderSettings `json:"data"`
	})
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// DNSForwarderApply represents the response from applying DNS forwarder changes
type DNSForwarderApply struct {
	Applied bool `json:"applied"`
}

// ApplyDNSForwarder applies pending DNS forwarder changes
func (s DNSService) ApplyDNSForwarder(ctx context.Context) (*DNSForwarderApply, error) {
	response, err := s.client.post(ctx, dnsForwarderApplyEndpoint, nil, nil)
	if err != nil {
		return nil, err
	}

	resp := new(struct {
		apiResponse
		Data *DNSForwarderApply `json:"data"`
	})
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// DNSForwarderHostOverride represents a DNS forwarder host override
type DNSForwarderHostOverride struct {
	Host    string                          `json:"host"`
	Domain  string                          `json:"domain"`
	IP      string                          `json:"ip"`
	Descr   string                          `json:"descr,omitempty"`
	Aliases []DNSForwarderHostOverrideAlias `json:"aliases,omitempty"`
}

// DNSForwarderHostOverrideAlias represents a DNS forwarder host override alias
type DNSForwarderHostOverrideAlias struct {
	Host   string `json:"host"`
	Domain string `json:"domain"`
	Descr  string `json:"descr,omitempty"`
}

// GetDNSForwarderHostOverrides returns the DNS forwarder host overrides
func (s DNSService) GetDNSForwarderHostOverrides(ctx context.Context) ([]*DNSForwarderHostOverride, error) {
	response, err := s.client.get(ctx, dnsForwarderHostOverrideEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(struct {
		apiResponse
		Data []*DNSForwarderHostOverride `json:"data"`
	})
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// CreateDNSForwarderHostOverride creates a new DNS forwarder host override
func (s DNSService) CreateDNSForwarderHostOverride(ctx context.Context, override DNSForwarderHostOverride) (*DNSForwarderHostOverride, error) {
	jsonData, err := json.Marshal(override)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	response, err := s.client.post(ctx, dnsForwarderHostOverrideEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(struct {
		apiResponse
		Data *DNSForwarderHostOverride `json:"data"`
	})
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}
