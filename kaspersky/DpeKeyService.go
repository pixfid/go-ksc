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

//	DpeKeyService Class Reference
//	Interface for working with encrypted devices. More...
//
//	List of all members.
type DpeKeyService service

//TODO: Not working now!!!
//
//	Returns information about host and key for chosen encrypted device.
//
//	Parameters:
//	- wstrDeviceId	[in] (string) Device id
//	Return:
//	- pKeyInfos	[out] (array) Array of params with key info. See Srvview encrypted devices on hosts.
//	The difference is that the key is decrypted.
func (di *DpeKeyService) GetDeviceKeys3(ctx context.Context, wstrDeviceId string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"wstrDeviceId": "%s"}`, wstrDeviceId))

	request, err := http.NewRequest("POST", di.client.Server+"/api/v1.0/DpeKeyService.GetDeviceKeys3", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := di.client.Do(ctx, request, nil)
	return raw, err
}
