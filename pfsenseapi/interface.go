package pfsenseapi

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/markphelps/optional"
)

const (
	interfaceEndpoint        = "api/v2/interface"
	interfacesEndpoint       = "api/v2/interfaces"
	interfaceVLANEndpoint    = "api/v2/interface/vlan"
	interfaceVLANsEndpoint   = "api/v2/interface/vlans"
	interfaceGroupEndpoint   = "api/v2/interface/group"
	interfaceGroupsEndpoint  = "api/v2/interface/groups"
	interfaceBridgeEndpoint  = "api/v2/interface/bridge"
	interfaceBridgesEndpoint = "api/v2/interface/bridges"
	interfaceApplyEndpoint   = "api/v2/interface/apply"
)

// InterfaceService provides interface API methods
type InterfaceService service

type Interface struct {
	InterfaceRequest
	Id string `json:"id"`
}

type interfaceListResponse struct {
	apiResponse
	Data []*Interface `json:"data"`
}

// GetInterface returns a single interface.
func (s InterfaceService) GetInterface(ctx context.Context, interfaceID string) (*Interface, error) {
	response, err := s.client.get(
		ctx,
		interfaceEndpoint,
		map[string]string{
			"if": interfaceID,
		},
	)
	if err != nil {
		return nil, err
	}

	resp := new(createInterfaceResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return resp.Data, nil
}

// ListInterfaces returns a list of the interfaces.
func (s InterfaceService) ListInterfaces(ctx context.Context) ([]*Interface, error) {
	response, err := s.client.get(ctx, interfacesEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(interfaceListResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// DeleteInterface deletes the interface. The interfaceID can be specified in
// either the interface's descriptive name, the pfSense ID (wan, lan, optx), or
// the physical interface id (e.g. igb0).
func (s InterfaceService) DeleteInterface(ctx context.Context, interfaceID string) (*Interface, error) {
	response, err := s.client.delete(
		ctx,
		interfaceEndpoint,
		map[string]string{
			"if": interfaceID,
		},
	)
	if err != nil {
		return nil, err
	}

	resp := new(createInterfaceResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

type InterfaceRequest struct {
	If                            string           `json:"if"`
	Enable                        *optional.Bool   `json:"enable,omitempty"`
	Descr                         string           `json:"descr"`
	Spoofmac                      *optional.String `json:"spoofmac,omitempty"`
	Mtu                           *optional.Int32  `json:"mtu,omitempty"`
	Mss                           *optional.Int32  `json:"mss,omitempty"`
	Media                         *optional.String `json:"media,omitempty"`
	Mediaopt                      *optional.String `json:"mediaopt,omitempty"`
	Blockpriv                     *optional.Bool   `json:"blockpriv,omitempty"`
	Blockbogons                   *optional.Bool   `json:"blockbogons,omitempty"`
	Typev4                        string           `json:"typev4"`
	Ipaddr                        string           `json:"ipaddr"`
	Subnet                        int32            `json:"subnet"`
	Gateway                       *optional.String `json:"gateway,omitempty"`
	AliasSubnet                   *optional.Int32  `json:"alias_subnet,omitempty"`
	AdvDhcpPtTimeout              *optional.Int32  `json:"adv_dhcp_pt_timeout,omitempty"`
	AdvDhcpPtRetry                *optional.Int32  `json:"adv_dhcp_pt_retry,omitempty"`
	AdvDhcpPtSelectTimeout        *optional.Int32  `json:"adv_dhcp_pt_select_timeout,omitempty"`
	AdvDhcpPtReboot               *optional.Int32  `json:"adv_dhcp_pt_reboot,omitempty"`
	AdvDhcpPtBackoffCutoff        *optional.Int32  `json:"adv_dhcp_pt_backoff_cutoff,omitempty"`
	AdvDhcpPtInitialInterval      *optional.Int32  `json:"adv_dhcp_pt_initial_interval,omitempty"`
	AdvDhcpSendOptions            *optional.String `json:"adv_dhcp_send_options,omitempty"`
	AdvDhcpRequestOptions         *optional.String `json:"adv_dhcp_request_options,omitempty"`
	AdvDhcpRequiredOptions        *optional.String `json:"adv_dhcp_required_options,omitempty"`
	AdvDhcpOptionModifiers        *optional.String `json:"adv_dhcp_option_modifiers,omitempty"`
	AdvDhcpConfigFileOverridePath *optional.String `json:"adv_dhcp_config_file_override_path,omitempty"`
	Typev6                        *optional.String `json:"typev6,omitempty"`
	Ipaddrv6                      string           `json:"ipaddrv6"`
	Subnetv6                      int32            `json:"subnetv6"`
	Gatewayv6                     *optional.String `json:"gatewayv6,omitempty"`
	Prefix6Rd                     string           `json:"prefix_6rd"`
	Gateway6Rd                    string           `json:"gateway_6rd"`
	Prefix6RdV4Plen               int32            `json:"prefix_6rd_v4plen"`
	Track6Interface               string           `json:"track6_interface"`
}

type createInterfaceResponse struct {
	apiResponse
	Data *Interface `json:"data"`
}

// CreateInterface creates a new interface.
func (s InterfaceService) CreateInterface(
	ctx context.Context,
	newInterface InterfaceRequest,
) (*Interface, error) {
	jsonData, err := json.Marshal(newInterface)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	response, err := s.client.post(ctx, interfaceEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(createInterfaceResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return resp.Data, nil
}

// UpdateInterface modifies an existing interface.
func (s InterfaceService) UpdateInterface(
	ctx context.Context,
	idToUpdate string,
	interfaceData InterfaceRequest,
) (*Interface, error) {
	requestData := Interface{
		InterfaceRequest: interfaceData,
		Id:               idToUpdate,
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	response, err := s.client.patch(ctx, interfaceEndpoint, nil, jsonData)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	resp := new(createInterfaceResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}
	return nil
}

pfsenseapi/system.go
````
<<<<<<< SEARCH
	return resp.Data, nil
}

