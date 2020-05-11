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

//	VServers2 Class Reference
//	Virtual servers processing. More...
//
//	List of all members:
type VServers2 service

//VServerStatistic struct
type VServerStatistic struct {
	VSStatistic *VSStatistic `json:"PxgRetVal,omitempty"`
}

type VSStatistic struct {
	KlvsrvCreated   *Klvsrv       `json:"KLVSRV_CREATED,omitempty"`
	KlvsrvGroups    *int64        `json:"KLVSRV_GROUPS,omitempty"`
	KlvsrvHosts     *int64        `json:"KLVSRV_HOSTS,omitempty"`
	KlvsrvLicenses  []interface{} `json:"KLVSRV_LICENSES"`
	KlvsrvMdmios    *int64        `json:"KLVSRV_MDMIOS,omitempty"`
	KlvsrvMobilies  *int64        `json:"KLVSRV_MOBILIES,omitempty"`
	KlvsrvProducts  *Klvsrv       `json:"KLVSRV_PRODUCTS,omitempty"`
	KlvsrvProducts2 *Klvsrv       `json:"KLVSRV_PRODUCTS_2,omitempty"`
	KlvsrvUsers     *int64        `json:"KLVSRV_USERS,omitempty"`
}

type Klvsrv struct {
	Type  *string `json:"type,omitempty"`
	Value *string `json:"value"`
}

//Acquire info on virtual server.
//
//Returns info about the specified virtual server
//
//Parameters:
//	- lVsId	(int64) virtual server id
//Returns:
//	- (params) a container, see Virtual server statistic.
func (vs *VServers2) GetVServerStatistic(ctx context.Context, lVsId int) (*VServerStatistic, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"lVsId": %d}`, lVsId))
	request, err := http.NewRequest("POST", vs.client.Server+"/api/v1.0/VServers2.GetVServerStatistic", bytes.NewBuffer(postData))

	vServerStatistic := new(VServerStatistic)
	raw, err := vs.client.Do(ctx, request, &vServerStatistic)
	return vServerStatistic, raw, err
}
