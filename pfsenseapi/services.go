package pfsenseapi

import (
	"context"
	"encoding/json"
	"fmt"
)

// ACMEService provides ACME API methods
type ACMEService service

// BindService provides BIND API methods
type BindService service

const (
	acmeAccountKeysEndpoint = "api/v2/services/acme/account_keys"
	acmeAccountKeyEndpoint  = "api/v2/services/acme/account_key"
	acmeCertificatesEndpoint = "api/v2/services/acme/certificates"
	acmeCertificateEndpoint  = "api/v2/services/acme/certificate"
	bindSettingsEndpoint = "api/v2/services/bind/settings"
	bindAccessListsEndpoint = "api/v2/services/bind/access_lists"
	bindAccessListEndpoint = "api/v2/services/bind/access_list"
	bindViewsEndpoint = "api/v2/services/bind/views"
	bindViewEndpoint = "api/v2/services/bind/view"
	bindZonesEndpoint   = "api/v2/services/bind/zones"
	bindZoneEndpoint    = "api/v2/services/bind/zone"
)

// ACMESettings represents the ACME settings
type ACMESettings struct {
	Enable bool `json:"enable"`
	// Add other fields based on the OpenAPI spec
}

// BINDSettings represents the BIND settings
type BINDSettings struct {
	Enable bool `json:"enable"`
	// Add other fields based on the OpenAPI spec
}

// GetACMESettings returns the current ACME settings
func (s ACMEService) GetACMESettings(ctx context.Context) (*ACMESettings, error) {
	response, err := s.client.get(ctx, "api/v2/services/acme/settings", nil)
	if err != nil {
		return nil, err
	}

	resp := new(struct {
		apiResponse
		Data *ACMESettings `json:"data"`
	})
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// UpdateACMESettings updates the ACME settings
func (s ACMEService) UpdateACMESettings(ctx context.Context, settings ACMESettings) error {
	jsonData, err := json.Marshal(settings)
	if err != nil {
		return fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	_, err = s.client.patch(ctx, "api/v2/services/acme/settings", nil, jsonData)
	return err
}

// GetBINDSettings returns the current BIND settings
func (s BindService) GetBINDSettings(ctx context.Context) (*BINDSettings, error) {
	response, err := s.client.get(ctx, bindSettingsEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(struct {
		apiResponse
		Data *BINDSettings `json:"data"`
	})
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// UpdateBINDSettings updates the BIND settings
func (s BindService) UpdateBINDSettings(ctx context.Context, settings BINDSettings) error {
	jsonData, err := json.Marshal(settings)
	if err != nil {
		return fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	_, err = s.client.patch(ctx, bindSettingsEndpoint, nil, jsonData)
	return err
}
