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

//	SrvCloud Class Reference
//
//	Interface to acquire info about public clouds.
//
//	List of all members.
type SrvCloud service

//	Returns list of clouds of the current server.
//
//	For main server also returns only clouds of main server.
//
//	Parameters:
//	- pParams	(params), For additional options. Reserved.
//
//	Returns:
//	- (array) collection of paramParams objects where each of them has the following structure:
//	each element contains attributes from List of server cloud attributes.
//
//	Note:
//	cloud containers may be listed my means of CloudContainersSrvViewName SrvView
//
//	See also:
//	List of server cloud attributes.
//	SrvView List of server cloud containers
//	SrvView List of server cloud hosts
func (sc *SrvCloud) GetCloudsInfo(ctx context.Context, params Null) ([]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", sc.client.Server+"/api/v1.0/SrvCloud.GetCloudsInfo",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := sc.client.Do(ctx, request, nil)
	return raw, nil
}

//	Returns properties of the cloud host.
//
//	Parameters:
//	- pCloudHostBinId	(binary). Internal id of the cloud host (see KLHST_CLOUD_HOST_BINID).
//	- pFields	(array) collection of cloud host attribute names that need to return see List of server cloud host
//	attributes..
//	- pParams	(params), For additional options. Reserved.
//
//	Returns:
//	- ppHostInfo (params) contains following attributes:
//
//	list of founded attributes that are specified in arrFields
//
//	Note:
//	cloud hosts may be listed my means of CloudHostsSrvViewName SrvView
//
//	See also:
//	List of server cloud host attributes.
//	SrvView List of server cloud hosts
//	SrvView List of server cloud containers
//TODO func (sc *SrvCloud) GetCloudHostInfo(ctx context.Context, params Params, params Params, params Null})
