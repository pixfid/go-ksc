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
	"encoding/json"
	"net/http"
)

//	SubnetMasks Class Reference
//	Subnets provider.
//
//	List of all members.
type SubnetMasks service

type PSubnetSettings struct {
	PSubnetSettings *PSubnetSettingsClass `json:"pSubnetSettings,omitempty"`
}

type PSubnetSettingsClass struct {
	NIPAddress     *int64  `json:"nIpAddress,omitempty"`
	NMask          *int64  `json:"nMask,omitempty"`
	WstrSubnetName *string `json:"wstrSubnetName,omitempty"`
	WstrComment    *string `json:"wstrComment,omitempty"`
}

func (sm *SubnetMasks) CreateSubnet(ctx context.Context, params PSubnetSettings) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", sm.client.Server+"/api/v1.0/SubnetMasks.CreateSubnet", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := sm.client.Do(ctx, request, nil)
	return raw, err
}
