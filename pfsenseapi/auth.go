package pfsenseapi

import (
	"context"
	"encoding/json"
	"fmt"
)

const (
	authJWTEndpoint  = "api/v2/auth/jwt"
	authKeyEndpoint  = "api/v2/auth/key"
	authKeysEndpoint = "api/v2/auth/keys"
)

// AuthService provides authentication API methods
type AuthService service

// RESTAPIJWT represents a REST API JWT
type RESTAPIJWT struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Token    string `json:"token,omitempty"`
}

// RESTAPIKey represents a REST API Key
type RESTAPIKey struct {
	Descr       string `json:"descr,omitempty"`
	Username    string `json:"username,omitempty"`
	HashAlgo    string `json:"hash_algo,omitempty"`
	LengthBytes int    `json:"length_bytes,omitempty"`
	Hash        string `json:"hash,omitempty"`
	Key         string `json:"key,omitempty"`
}

type restAPIKeyListResponse struct {
	apiResponse
	Data []*RESTAPIKey `json:"data"`
}

// ListKeys returns a list of REST API Keys
func (s AuthService) ListKeys(ctx context.Context) ([]*RESTAPIKey, error) {
	response, err := s.client.get(ctx, authKeysEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(restAPIKeyListResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

type restAPIKeyResponse struct {
	apiResponse
	Data *RESTAPIKey `json:"data"`
}

// GetKey returns a REST API Key
func (s AuthService) GetKey(ctx context.Context, id string) (*RESTAPIKey, error) {
	response, err := s.client.get(
		ctx,
		authKeyEndpoint,
		map[string]string{
			"id": id,
		},
	)
	if err != nil {
		return nil, err
	}

	resp := new(restAPIKeyResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// CreateKey creates a new REST API Key
func (s AuthService) CreateKey(ctx context.Context, key RESTAPIKey) (*RESTAPIKey, error) {
	jsonData, err := json.Marshal(key)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	response, err := s.client.post(ctx, authKeyEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(restAPIKeyResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// DeleteKey deletes a REST API Key
func (s AuthService) DeleteKey(ctx context.Context, id string) (*RESTAPIKey, error) {
	response, err := s.client.delete(
		ctx,
		authKeyEndpoint,
		map[string]string{
			"id": id,
		},
	)
	if err != nil {
		return nil, err
	}

	resp := new(restAPIKeyResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

type restAPIJWTResponse struct {
	apiResponse
	Data *RESTAPIJWT `json:"data"`
}

// CreateJWT creates a new REST API JWT
func (s AuthService) CreateJWT(ctx context.Context, jwt RESTAPIJWT) (*RESTAPIJWT, error) {
	jsonData, err := json.Marshal(jwt)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	response, err := s.client.post(ctx, authJWTEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(restAPIJWTResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}
