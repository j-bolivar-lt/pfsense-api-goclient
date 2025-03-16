package pfsenseapi

import (
	"context"
	"encoding/json"
	"fmt"
)

const (
	routingGatewayEndpoint      = "api/v2/routing/gateway"
	routingGatewaysEndpoint     = "api/v2/routing/gateways"
	routingGatewayGroupEndpoint = "api/v2/routing/gateway_group"
	routingGatewayGroupsEndpoint = "api/v2/routing/gateway_groups"
	routingStaticRouteEndpoint  = "api/v2/routing/static_route"
	routingStaticRoutesEndpoint = "api/v2/routing/static_routes"
	routingApplyEndpoint        = "api/v2/routing/apply"
)

// RoutingService provides routing API methods
type RoutingService service

// RoutingGateway represents a routing gateway
type RoutingGateway struct {
	Name                      string `json:"name"`
	Descr                     string `json:"descr,omitempty"`
	Disabled                  bool   `json:"disabled"`
	IPProtocol                string `json:"ipprotocol"`
	Interface                 string `json:"interface"`
	Gateway                   string `json:"gateway"`
	MonitorDisable            bool   `json:"monitor_disable"`
	Monitor                   string `json:"monitor,omitempty"`
	Weight                    int    `json:"weight"`
	DataPayload               int    `json:"data_payload,omitempty"`
	LatencyLow                int    `json:"latencylow,omitempty"`
	LatencyHigh               int    `json:"latencyhigh,omitempty"`
	LossLow                   int    `json:"losslow,omitempty"`
	LossHigh                  int    `json:"losshigh,omitempty"`
	Interval                  int    `json:"interval,omitempty"`
	LossInterval              int    `json:"loss_interval,omitempty"`
	TimePeriod                int    `json:"time_period,omitempty"`
	AlertInterval             int    `json:"alert_interval,omitempty"`
}

type routingGatewayListResponse struct {
	apiResponse
	Data []*RoutingGateway `json:"data"`
}

// ListGateways returns a list of routing gateways
func (s RoutingService) ListGateways(ctx context.Context) ([]*RoutingGateway, error) {
	response, err := s.client.get(ctx, routingGatewaysEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(routingGatewayListResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

type routingGatewayResponse struct {
	apiResponse
	Data *RoutingGateway `json:"data"`
}

// GetGateway returns a routing gateway by name
func (s RoutingService) GetGateway(ctx context.Context, name string) (*RoutingGateway, error) {
	response, err := s.client.get(
		ctx,
		routingGatewayEndpoint,
		map[string]string{
			"name": name,
		},
	)
	if err != nil {
		return nil, err
	}

	resp := new(routingGatewayResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// CreateGateway creates a new routing gateway
func (s RoutingService) CreateGateway(ctx context.Context, gateway RoutingGateway) (*RoutingGateway, error) {
	jsonData, err := json.Marshal(gateway)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	response, err := s.client.post(ctx, routingGatewayEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(routingGatewayResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// UpdateGateway updates an existing routing gateway
func (s RoutingService) UpdateGateway(ctx context.Context, gateway RoutingGateway) (*RoutingGateway, error) {
	jsonData, err := json.Marshal(gateway)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	response, err := s.client.put(
		ctx,
		routingGatewayEndpoint,
		map[string]string{
			"name": gateway.Name,
		},
		jsonData,
	)
	if err != nil {
		return nil, err
	}

	resp := new(routingGatewayResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// DeleteGateway deletes a routing gateway by name
func (s RoutingService) DeleteGateway(ctx context.Context, name string) (*RoutingGateway, error) {
	response, err := s.client.delete(
		ctx,
		routingGatewayEndpoint,
		map[string]string{
			"name": name,
		},
	)
	if err != nil {
		return nil, err
	}

	resp := new(routingGatewayResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// RoutingGatewayGroup represents a routing gateway group
type RoutingGatewayGroup struct {
	Name        string                        `json:"name"`
	Trigger     string                        `json:"trigger"`
	Descr       string                        `json:"descr,omitempty"`
	IPProtocol  string                        `json:"ipprotocol,omitempty"`
	Priorities  []RoutingGatewayGroupPriority `json:"priorities"`
}

// RoutingGatewayGroupPriority represents a priority in a routing gateway group
type RoutingGatewayGroupPriority struct {
	Gateway   string `json:"gateway"`
	Tier      int    `json:"tier"`
	VirtualIP string `json:"virtual_ip,omitempty"`
}

type routingGatewayGroupListResponse struct {
	apiResponse
	Data []*RoutingGatewayGroup `json:"data"`
}

// ListGatewayGroups returns a list of routing gateway groups
func (s RoutingService) ListGatewayGroups(ctx context.Context) ([]*RoutingGatewayGroup, error) {
	response, err := s.client.get(ctx, routingGatewayGroupsEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(routingGatewayGroupListResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

type routingGatewayGroupResponse struct {
	apiResponse
	Data *RoutingGatewayGroup `json:"data"`
}

// GetGatewayGroup returns a routing gateway group by name
func (s RoutingService) GetGatewayGroup(ctx context.Context, name string) (*RoutingGatewayGroup, error) {
	response, err := s.client.get(
		ctx,
		routingGatewayGroupEndpoint,
		map[string]string{
			"name": name,
		},
	)
	if err != nil {
		return nil, err
	}

	resp := new(routingGatewayGroupResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// CreateGatewayGroup creates a new routing gateway group
func (s RoutingService) CreateGatewayGroup(ctx context.Context, group RoutingGatewayGroup) (*RoutingGatewayGroup, error) {
	jsonData, err := json.Marshal(group)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	response, err := s.client.post(ctx, routingGatewayGroupEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(routingGatewayGroupResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// UpdateGatewayGroup updates an existing routing gateway group
func (s RoutingService) UpdateGatewayGroup(ctx context.Context, group RoutingGatewayGroup) (*RoutingGatewayGroup, error) {
	jsonData, err := json.Marshal(group)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	response, err := s.client.put(
		ctx,
		routingGatewayGroupEndpoint,
		map[string]string{
			"name": group.Name,
		},
		jsonData,
	)
	if err != nil {
		return nil, err
	}

	resp := new(routingGatewayGroupResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// DeleteGatewayGroup deletes a routing gateway group by name
func (s RoutingService) DeleteGatewayGroup(ctx context.Context, name string) (*RoutingGatewayGroup, error) {
	response, err := s.client.delete(
		ctx,
		routingGatewayGroupEndpoint,
		map[string]string{
			"name": name,
		},
	)
	if err != nil {
		return nil, err
	}

	resp := new(routingGatewayGroupResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// StaticRoute represents a static route
type StaticRoute struct {
	Network  string `json:"network"`
	Gateway  string `json:"gateway"`
	Descr    string `json:"descr,omitempty"`
	Disabled bool   `json:"disabled"`
}

type staticRouteListResponse struct {
	apiResponse
	Data []*StaticRoute `json:"data"`
}

// ListStaticRoutes returns a list of static routes
func (s RoutingService) ListStaticRoutes(ctx context.Context) ([]*StaticRoute, error) {
	response, err := s.client.get(ctx, routingStaticRoutesEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := new(staticRouteListResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

type staticRouteResponse struct {
	apiResponse
	Data *StaticRoute `json:"data"`
}

// GetStaticRoute returns a static route by ID
func (s RoutingService) GetStaticRoute(ctx context.Context, id string) (*StaticRoute, error) {
	response, err := s.client.get(
		ctx,
		routingStaticRouteEndpoint,
		map[string]string{
			"id": id,
		},
	)
	if err != nil {
		return nil, err
	}

	resp := new(staticRouteResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// CreateStaticRoute creates a new static route
func (s RoutingService) CreateStaticRoute(ctx context.Context, route StaticRoute) (*StaticRoute, error) {
	jsonData, err := json.Marshal(route)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	response, err := s.client.post(ctx, routingStaticRouteEndpoint, nil, jsonData)
	if err != nil {
		return nil, err
	}

	resp := new(staticRouteResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// UpdateStaticRoute updates an existing static route
func (s RoutingService) UpdateStaticRoute(ctx context.Context, id string, route StaticRoute) (*StaticRoute, error) {
	jsonData, err := json.Marshal(route)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request payload into json: %w", err)
	}

	response, err := s.client.put(
		ctx,
		routingStaticRouteEndpoint,
		map[string]string{
			"id": id,
		},
		jsonData,
	)
	if err != nil {
		return nil, err
	}

	resp := new(staticRouteResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// DeleteStaticRoute deletes a static route by ID
func (s RoutingService) DeleteStaticRoute(ctx context.Context, id string) (*StaticRoute, error) {
	response, err := s.client.delete(
		ctx,
		routingStaticRouteEndpoint,
		map[string]string{
			"id": id,
		},
	)
	if err != nil {
		return nil, err
	}

	resp := new(staticRouteResponse)
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}

// RoutingApply represents the response from applying routing changes
type RoutingApply struct {
	Applied bool `json:"applied"`
}

// Apply applies pending routing changes
func (s RoutingService) Apply(ctx context.Context) (*RoutingApply, error) {
	response, err := s.client.post(ctx, routingApplyEndpoint, nil, nil)
	if err != nil {
		return nil, err
	}

	resp := new(struct {
		apiResponse
		Data *RoutingApply `json:"data"`
	})
	if err = json.Unmarshal(response, resp); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return resp.Data, nil
}
