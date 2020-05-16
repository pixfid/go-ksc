/*
 *
 * 	Copyright (C) 2020  <Semchenko Aleksandr>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful, but
 * WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.See the GNU
 * General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.If not, see <http://www.gnu.org/licenses/>.
 * /
 */

package kaspersky

import (
	"bytes"
	"context"
	"fmt"

	"net/http"
)

//	UaControl Class Reference
//	Update agents and Connection gateways management.
//
//	List of all members.
type UaControl service

//	Check if Update agents automatic assignment is enabled, see uactl_ua_assignment.
//
//	Returns:
//	- true if UAs assigned automatically
func (uc *UaControl) GetAssignUasAutomatically(ctx context.Context) (*PxgValBool, []byte, error) {
	request, err := http.NewRequest("POST", uc.client.Server+"/api/v1.0/UaControl.GetAssignUasAutomatically", nil)

	pxgValBool := new(PxgValBool)
	raw, err := uc.client.Do(ctx, request, &pxgValBool)
	return pxgValBool, raw, err
}

//	Fill in default Update agent settings.
//
//	Parameters:
//
//	Return:
//	- pUaInfo	default settings, see Update agent settings
func (uc *UaControl) GetDefaultUpdateAgentRegistrationInfo(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", uc.client.Server+"/api/v1.0/UaControl.GetDefaultUpdateAgentRegistrationInfo", nil)
	raw, err := uc.client.Do(ctx, request, nil)
	return raw, err
}

//TODO func (uc *UaControl) GetUpdateAgentInfo(ctx context.Context)  ([]byte, error) {}

//	Returns list of Update agents assigned to specified host.
//
//	Parameters:
//	- wstrHostId(string)	Host id to find UAs
//
//	Return:
//	- pUaInfo	(array) of Update agents display info containers
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

//	List all registered Update agents.
//
//	Parameters:
//
//	Return:
//	- pUasArr	(array) of Update agents info containers.
func (uc *UaControl) GetUpdateAgentsList(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", uc.client.Server+"/api/v1.0/UaControl.GetUpdateAgentsList", nil)

	raw, err := uc.client.Do(ctx, request, nil)
	return raw, err
}

//TODO func (uc *UaControl) ModifyUpdateAgent(ctx context.Context)  ([]byte, error) {}
//TODO func (uc *UaControl) RegisterDmzGateway(ctx context.Context)  ([]byte, error) {}
//TODO func (uc *UaControl) RegisterUpdateAgent(ctx context.Context)  ([]byte, error) {}

//	Enable or disable automatic Update agents assignment, see uactl_ua_assignment.
//
//	Parameters:
//	- bEnabled	(bool)	enable or disable bool
func (uc *UaControl) SetAssignUasAutomatically(ctx context.Context, bEnabled bool) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"bEnabled": %v}`, bEnabled))
	request, err := http.NewRequest("POST", uc.client.Server+"/api/v1.0/UaControl.UnregisterUpdateAgent",
		bytes.NewBuffer(postData))

	raw, err := uc.client.Do(ctx, request, nil)
	return raw, err
}

//	Unregister host as Update agent.
//
//	Parameters:
//	- wstrUaHostId	(string)	UA host id
func (uc *UaControl) UnregisterUpdateAgent(ctx context.Context, wstrUaHostId string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"wstrUaHostId": "%s"}`, wstrUaHostId))
	request, err := http.NewRequest("POST", uc.client.Server+"/api/v1.0/UaControl.UnregisterUpdateAgent",
		bytes.NewBuffer(postData))

	raw, err := uc.client.Do(ctx, request, nil)
	return raw, err
}
