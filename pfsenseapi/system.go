package pfsenseapi

import (
	"context"
	"encoding/json"
	"fmt"
)

const (
	systemHaltEndpoint   = "api/v2/system/halt"
	systemRebootEndpoint = "api/v2/system/reboot"
	systemVersionEndpoint = "api/v2/system/version"
	systemConfigEndpoint = "api/v2/system/config"
	systemCronEndpoint   = "api/v2/system/cron"
)

// SystemService provides system API methods
type SystemService service

// SystemHalt represents a system halt request
type SystemHalt struct {
	DryRun bool `json:"dry_run"`
}

// Halt halts the system
func (s SystemService) Halt(ctx context.Context, dryRun bool) error {
	halt := SystemHalt{
		DryRun: dryRun,
	}

	jsonData, err := json.Marshal(halt)
	if err != nil {
		return fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	_, err = s.client.post(ctx, systemHaltEndpoint, nil, jsonData)
	return err
}

// SystemReboot represents a system reboot request
type SystemReboot struct {
	DryRun bool `json:"dry_run"`
}

// Reboot reboots the system
func (s SystemService) Reboot(ctx context.Context, dryRun bool) error {
	reboot := SystemReboot{
		DryRun: dryRun,
	}

	jsonData, err := json.Marshal(reboot)
	if err != nil {
		return fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	_, err = s.client.post(ctx, systemRebootEndpoint, nil, jsonData)
	return err
}

// SystemVersion represents the system version information
type SystemVersion struct {
	Version     string `json:"version"`
	BuildTime   string `json:"build_time,omitempty"`
	Platform    string `json:"platform,omitempty"`
	Architecture string `json:"arch,omitempty"`
	PatchLevel  string `json:"patch,omitempty"`
	Firmware    string `json:"firmware,omitempty"`
}

// GetVersion returns the system version information
func (s SystemService) GetVersion(ctx context.Context) (*SystemVersion, error) {
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

// SystemConfig represents the system configuration
type SystemConfig struct {
	ConfigText string `json:"config_text"`
}

// GetConfig returns the system configuration
func (s SystemService) GetConfig(ctx context.Context) (*SystemConfig, error) {
	response, err := s.client.get(ctx, systemConfigEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(struct {
		apiResponse
		Data *SystemConfig `json:"data"`
	})
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
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
func (s SystemService) ListCronJobs(ctx context.Context) ([]*CronJob, error) {
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
func (s SystemService) GetCronJob(ctx context.Context, id string) (*CronJob, error) {
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
func (s SystemService) CreateCronJob(ctx context.Context, job CronJob) (*CronJob, error) {
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
func (s SystemService) UpdateCronJob(ctx context.Context, id string, job CronJob) (*CronJob, error) {
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
func (s SystemService) DeleteCronJob(ctx context.Context, id string) (*CronJob, error) {
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
