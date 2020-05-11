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

//	CertPoolCtrl Class Reference
//
//	Interface to manage the pool of certificates used by the
//	Kaspersky Security Center Server.
//
//	Public Member Functions
type CertPoolCtrl service

//	Returns information about certificate from server's certificates pool.
//
//	Parameters:
//	- nVServerId	[in] (int64) Virtual server id (-1 for current, 0 for main server)
//	- nFunction	[in] (int64) Certificate function (see "KLCERTP::CertificateFunction enum values")
//
//	Returns:
//	- (params) If certificate present then it returns params
//	with "CPublic" (paramBinary) field only.
func (cp *CertPoolCtrl) GetCertificateInfo(ctx context.Context, nVServerId, nFunction int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nVServerId": %d, "nFunction" : %d }`, nVServerId, nFunction))
	request, err := http.NewRequest("POST", cp.client.Server+"/api/v1.0/CertPoolCtrl.GetCertificateInfo",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := cp.client.Do(ctx, request, nil)
	return raw, err
}

//TODO func (cp *SetCertificate) GetCertificateInfo(ctx context.Context, nVServerId, nFunction int64,
//pCertData interface{}) ([]byte,
//error)
