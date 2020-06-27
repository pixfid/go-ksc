/*
 * MIT License
 *
 * Copyright (c) [2020] [Semchenko Aleksandr]
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package kaspersky

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"net/http"
)

// UaControl service Update agents and Connection gateways management.
type UaControl service

// GetAssignUasAutomatically Check if Update agents automatic assignment is enabled, see uactl_ua_assignment.
func (uc *UaControl) GetAssignUasAutomatically(ctx context.Context) (*PxgValBool, []byte, error) {
	request, err := http.NewRequest("POST", uc.client.Server+"/api/v1.0/UaControl.GetAssignUasAutomatically", nil)
	if err != nil {
		return nil, nil, err
	}

	pxgValBool := new(PxgValBool)
	raw, err := uc.client.Do(ctx, request, &pxgValBool)
	return pxgValBool, raw, err
}

// GetDefaultUpdateAgentRegistrationInfo Fill in default Update agent settings.
func (uc *UaControl) GetDefaultUpdateAgentRegistrationInfo(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", uc.client.Server+"/api/v1.0/UaControl.GetDefaultUpdateAgentRegistrationInfo", nil)
	if err != nil {
		return nil, err
	}

	raw, err := uc.client.Do(ctx, request, nil)
	return raw, err
}

// GetUpdateAgentInfo Get Update agent info.
func (uc *UaControl) GetUpdateAgentInfo(ctx context.Context, wstrUaHostId string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"wstrUaHostId": "%s"}`, wstrUaHostId))
	request, err := http.NewRequest("POST", uc.client.Server+"/api/v1.0/UaControl.GetUpdateAgentInfo",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uc.client.Do(ctx, request, nil)
	return raw, err
}

// GetUpdateAgentsDisplayInfoForHost Returns list of Update agents assigned to specified host.
func (uc *UaControl) GetUpdateAgentsDisplayInfoForHost(ctx context.Context, wstrHostId string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"wstrHostId": "%s"}`, wstrHostId))
	request, err := http.NewRequest("POST", uc.client.Server+"/api/v1.0/UaControl.GetUpdateAgentsDisplayInfoForHost",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uc.client.Do(ctx, request, nil)
	return raw, err
}

// GetUpdateAgentsList List all registered Update agents.
func (uc *UaControl) GetUpdateAgentsList(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", uc.client.Server+"/api/v1.0/UaControl.GetUpdateAgentsList", nil)
	if err != nil {
		return nil, err
	}

	raw, err := uc.client.Do(ctx, request, nil)
	return raw, err
}

// ModifyUpdateAgent Modify update agent info of an existing Update agent.
func (uc *UaControl) ModifyUpdateAgent(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", uc.client.Server+"/api/v1.0/UaControl.ModifyUpdateAgent",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uc.client.Do(ctx, request, nil)
	return raw, err
}

// RegisterDmzGateway Register Connection gateway located in DMZ.
//
// See Connection gateways to know about DMZ-located connection gateways. To register CG in DMZ, one should:
// on target host install Network agent with option 'Use as connection gateway' in installer or in installation package settings
// determine scope of hosts which will be using this CG call this method providing ip address by which CG host is available to KSC server
//
// After successful registration CG host will appear in Unassigned computers group and will be assigned as CG.
//
// After that move CG host to appropriate Administration group.
// If Network agents which will be using this CG cannot access KSC server without using this CG (which is likely)
// they must be installed with 'Use CG' installer option.
func (uc *UaControl) RegisterDmzGateway(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", uc.client.Server+"/api/v1.0/UaControl.RegisterDmzGateway",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uc.client.Do(ctx, request, nil)
	return raw, err
}

// RegisterUpdateAgent Register host as Update agent or Connection gateway.
//
// To register host as an Update agent one should:
// call GetDefaultUpdateAgentRegistrationInfo() to get default UA settings set host id via UaHostId attribute (mandatory)
// define Update agent scope via UaScope attribute (mandatory) modify other options if required (optional),
// see Update agent settings perform this call
func (uc *UaControl) RegisterUpdateAgent(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", uc.client.Server+"/api/v1.0/UaControl.RegisterUpdateAgent",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uc.client.Do(ctx, request, nil)
	return raw, err
}

// SetAssignUasAutomatically Enable or disable automatic Update agents assignment, see uactl_ua_assignment.
func (uc *UaControl) SetAssignUasAutomatically(ctx context.Context, bEnabled bool) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"bEnabled": %v}`, bEnabled))
	request, err := http.NewRequest("POST", uc.client.Server+"/api/v1.0/UaControl.UnregisterUpdateAgent",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := uc.client.Do(ctx, request, nil)
	return raw, err
}

// UnregisterUpdateAgent Unregister host as Update agent.
func (uc *UaControl) UnregisterUpdateAgent(ctx context.Context, wstrUaHostId string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"wstrUaHostId": "%s"}`, wstrUaHostId))
	request, err := http.NewRequest("POST", uc.client.Server+"/api/v1.0/UaControl.UnregisterUpdateAgent",
		bytes.NewBuffer(postData))

	if err != nil {
		return nil, err
	}

	raw, err := uc.client.Do(ctx, request, nil)
	return raw, err
}
