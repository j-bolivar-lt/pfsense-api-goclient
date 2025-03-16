package pfsenseapi

import (
	"context"
	"encoding/json"
	"fmt"
)

// BridgeService provides bridge API methods
type BridgeService service

// InterfaceBridge represents a single bridge.
type InterfaceBridge struct {
	InterfaceBridgeRequest
	Id string `json:"id"`
}

type interfaceBridgeListResponse struct {
	apiResponse
	Data []*InterfaceBridge `json:"data"`
}

// ListInterfaceBridges returns the bridges.
func (s BridgeService) ListInterfaceBridges(ctx context.Context) ([]*InterfaceBridge, error) {
	response, err := s.client.get(ctx, interfaceBridgesEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(interfaceBridgeListResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

type interfaceBridgeResponse struct {
	apiResponse
	Data *InterfaceBridge `json:"data"`
}

// GetInterfaceBridge returns the bridge with the given ID.
func (s BridgeService) GetInterfaceBridge(ctx context.Context, id string) (*InterfaceBridge, error) {
	response, err := s.client.get(
		ctx,
		interfaceBridgeEndpoint,
		map[string]string{
			"id": id,
		},
	)
	if err != nil {
		return nil, err
	}

	resp := new(interfaceBridgeResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return resp.Data, nil
}

// DeleteInterfaceBridge deletes a bridge.
func (s BridgeService) DeleteInterfaceBridge(ctx context.Context, idToDelete string) (*InterfaceBridge, error) {
	response, err := s.client.delete(
		ctx,
		interfaceBridgeEndpoint,
		map[string]string{
			"id": idToDelete,
		},
	)
	if err != nil {
		return nil, err
	}

	resp := new(interfaceBridgeResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return resp.Data, nil
}

// CreateInterfaceBridge creates a new bridge.
func (s BridgeService) CreateInterfaceBridge(
	ctx context.Context,
	newBridge InterfaceBridgeRequest,
) (*InterfaceBridge, error) {
	jsonData, err := json.Marshal(newBridge)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	response, err := s.client.post(ctx, interfaceBridgeEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(interfaceBridgeResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return resp.Data, nil
}

// UpdateInterfaceBridge updates an existing bridge.
func (s BridgeService) UpdateInterfaceBridge(
	ctx context.Context,
	idToUpdate string,
	bridgeData InterfaceBridgeRequest,
) (*InterfaceBridge, error) {
	requestData := InterfaceBridge{
		InterfaceBridgeRequest: bridgeData,
		Id:                     idToUpdate,
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	response, err := s.client.patch(ctx, interfaceBridgeEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(interfaceBridgeResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return resp.Data, nil
}

// InterfaceBridgeRequest represents the request to create or update a bridge.
type InterfaceBridgeRequest struct {
	Members  []string `json:"members"`
	Descr    string   `json:"descr"`
	Bridgeif string   `json:"bridgeif"`
}
