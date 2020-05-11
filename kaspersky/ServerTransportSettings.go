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
	"context"
	"net/http"
)

//	ServerTransportSettings Class Reference
//
//	Server transport settings proxy class. More...
//
//	List of all members.
type ServerTransportSettings service

//	GetNumberOfManagedDevicesAgentless.
//	Returns number of agentless managed devices.
//
//	Note: It can be called from main server only !
//
//	Returns:
//	It returns total number of managed devices for main server and all virtual servers.
func (sts *ServerTransportSettings) GetNumberOfManagedDevicesAgentless(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", sts.client.Server+"/api/v1.0/ServerTransportSettings.GetNumberOfManagedDevicesAgentless", nil)

	raw, err := sts.client.Do(ctx, request, nil)
	return raw, err
}

//	GetNumberOfManagedDevicesKSM.
//	Returns number of managed devices for KSM (Kaspersky for Mobile).
//
//	Note: It can be called from main server only !
//
//	Returns:
//	It returns total number of managed devices for main server and all virtual servers.
func (sts *ServerTransportSettings) GetNumberOfManagedDevicesKSM(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", sts.client.Server+"/api/v1.0/ServerTransportSettings.GetNumberOfManagedDevicesKSM", nil)

	raw, err := sts.client.Do(ctx, request, nil)
	return raw, err
}
