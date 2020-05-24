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
	return raw, err
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
