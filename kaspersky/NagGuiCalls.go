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

//	NagGuiCalls Class Reference
//
//	Remote host caller. More...
//
//	This interface is implemented at Network Agent side,
//	so use gateway connection to connect Network Agent and call interface methods.
//
//	List of all members.
type NagGuiCalls service

//	Asynchronously connects to the remote host (if hasn't connected yet), and makes call with the specified name szwCallName
//
//	Parameters:
//	- szwProduct	product name
//	- szwVersion	product settings compatibility version
//	- szwCallName	remote call name
//	- pInData	call-specific input data
//
//	Returns:
//	- asynchronous request ID, used to get the result
//
//	Remarks:
//	Check the operation state by calling AsyncActionStateChecker::CheckActionState periodically until it's
//	finalized. If the operation succedes then AsyncActionStateChecker::CheckActionState returns call-results in pStateData. Otherwise, a call to AsyncActionStateChecker::CheckActionState returns error in pStateData.
//
//	Exceptions:
//	Throws	exception in case of error, see Some error definitions
//
//	 +------------------------------------------------+------+--------------------------+--------------------------------------------------+
//	 |                  Module name                   | Code |      Mnemonic name       |                Error description                 |
//	 +------------------------------------------------+------+--------------------------+--------------------------------------------------+
//	 | KLPRSS                                         | 1125 | KLPRSS::NOT_EXIST        | Parameter with the specified name does not exist |
//	 | KLSTD                                          | 1184 | KLSTD::STDE_NOACCESS     | Access denied                                    |
//	 | KLSTD                                          | 1186 | KLSTD::STDE_NOTFOUND     | Object not found                                 |
//	 | KLSTD                                          | 1193 | KLSTD::STDE_NOTPERM      | Operation is not permitted                       |
//	 | KLCONN                                         | 1194 | KLCONN::GCR_NO_SUCH_CALL |  No such GUI call (GUI call is not implemented)  |
//	 +------------------------------------------------+------+--------------------------+--------------------------------------------------+
func (sd *NagGuiCalls) CallConnectorAsync(ctx context.Context, params interface{}) ([]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", sd.client.Server+"/api/v1.0/NagGuiCalls.CallConnectorAsync", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := sd.client.Do(ctx, request, nil)
	return raw, err
}
