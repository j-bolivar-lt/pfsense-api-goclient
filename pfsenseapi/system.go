package pfsenseapi

import (
	"context"
	"encoding/json"
	"fmt"
)

const (
	systemHaltEndpoint    = "api/v2/system/halt"
	systemRebootEndpoint  = "api/v2/system/reboot"
	systemVersionEndpoint = "api/v2/system/version"
	systemConfigEndpoint  = "api/v2/system/config"
	systemCronEndpoint    = "api/v2/system/cron"
	systemDNSEndpoint     = "api/v2/system/dns"
	systemStatusEndpoint  = "api/v2/status/system"
)

type dnsSystem struct {
	DnsAllowOverride bool     `json:"dnsallowoverride"`
	DnsLocalhost     string   `json:"dnslocalhost,omitempty"`
	DnsServer        []string `json:"dnsserver"`
}

type SystemStatusResponse struct {
	Platform      string `json:"platform"`
	Serial        string `json:"serial"`
	NetgateDevice string `json:"netgate_id"`
	Uptime        string `json:"uptime"`
	Bios_vendor   string `json:"bios_vendor"`
	Bios_version  string `json:"bios_version"`
	Kernel_pti    bool   `json:"kernel_pti"`
	Cpu_model     string `json:"cpu_model"`
}

type SystemStatus struct {
	apiResponse
	Data SystemStatusResponse `json:"data"`
}

type dnsSystemResponse struct {
	apiResponse
	Data dnsSystem `json:"data"`
}

func (s *SystemService) GetSystemStatus(ctx context.Context) (*SystemStatus, error) {
	response, err := s.client.get(ctx, systemStatusEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(SystemStatus)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return resp, nil
}

func (s *SystemService) GetDNSResolverSettings(ctx context.Context) (*dnsSystemResponse, error) {
	response, err := s.client.get(ctx, systemDNSEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(dnsSystemResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return resp, nil
}

func (s *SystemService) PatchDNSResolverSettings(ctx context.Context, d dnsSystem) (*dnsSystemResponse, error) {
	// Convert the dnsSystem struct to JSON
	body, err := json.Marshal(d)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request: %w", err)
	}

	// Call patch with the correct parameter order: ctx, endpoint, queryMap, body
	response, err := s.client.patch(ctx, systemDNSEndpoint, nil, body)
	if err != nil {
		return nil, err
	}

	resp := new(dnsSystemResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return resp, nil
}

// SystemService provides system API methods
type SystemService service

// SystemVersion represents the system version information
type SystemVersion struct {
	Version      string `json:"version"`
	BuildTime    string `json:"build_time,omitempty"`
	Platform     string `json:"platform,omitempty"`
	Architecture string `json:"arch,omitempty"`
	PatchLevel   string `json:"patch,omitempty"`
	Firmware     string `json:"firmware,omitempty"`
}

// GetVersion returns the system version information
func (s *SystemService) GetVersion(ctx context.Context) (*SystemVersion, error) {
	response, err := s.client.get(ctx, systemVersionEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(struct {
		apiResponse
		Data *SystemVersion `json:"data"`
	})
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// GetConfigRaw returns the system configuration as a raw string
func (s *SystemService) GetConfigRaw(ctx context.Context) (string, error) {
	responseBytes, err := s.client.get(ctx, systemConfigEndpoint, nil)
	if err != nil {
		return "", err
	}

	resp := new(apiResponse)
	if err = json.Unmarshal(responseBytes, resp); err != nil {
		return "", fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data.(map[string]interface{})["config_text"].(string), nil
}

// CronJob represents a cron job
type CronJob struct {
	Minute  string `json:"minute"`
	Hour    string `json:"hour,omitempty"`
	Mday    string `json:"mday,omitempty"`
	Month   string `json:"month,omitempty"`
	Wday    string `json:"wday,omitempty"`
	Who     string `json:"who"`
	Command string `json:"command"`
}

type cronJobListResponse struct {
	apiResponse
	Data []*CronJob `json:"data"`
}

// ListCronJobs returns a list of cron jobs
func (s *SystemService) ListCronJobs(ctx context.Context) ([]*CronJob, error) {
	response, err := s.client.get(ctx, systemCronEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(cronJobListResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

type cronJobResponse struct {
	apiResponse
	Data *CronJob `json:"data"`
}

// GetCronJob returns a cron job by ID
func (s *SystemService) GetCronJob(ctx context.Context, id string) (*CronJob, error) {
	response, err := s.client.get(
		ctx,
		systemCronEndpoint,
		map[string]string{
			"id": id,
		},
	)
	if err != nil {
		return nil, err
	}

	resp := new(cronJobResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// CreateCronJob creates a new cron job
func (s *SystemService) CreateCronJob(ctx context.Context, job CronJob) (*CronJob, error) {
	jsonData, err := json.Marshal(job)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	response, err := s.client.post(ctx, systemCronEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(cronJobResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// UpdateCronJob updates an existing cron job
func (s *SystemService) UpdateCronJob(ctx context.Context, id string, job CronJob) (*CronJob, error) {
	jsonData, err := json.Marshal(job)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	response, err := s.client.put(
		ctx,
		systemCronEndpoint,
		map[string]string{
			"id": id,
		},
		jsonData,
	)
	if err != nil {
		return nil, err
	}

	resp := new(cronJobResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// DeleteCronJob deletes a cron job by ID
func (s *SystemService) DeleteCronJob(ctx context.Context, id string) (*CronJob, error) {
	response, err := s.client.delete(
		ctx,
		systemCronEndpoint,
		map[string]string{
			"id": id,
		},
	)
	if err != nil {
		return nil, err
	}

	resp := new(cronJobResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}
