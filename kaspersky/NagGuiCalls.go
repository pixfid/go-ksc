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

//	NagGuiCalls Class Reference
//
//	Remote host caller..
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
//	Check the operation state by calling AsyncActionStateChecker.CheckActionState periodically until it's
//	finalized. If the operation succeeds then AsyncActionStateChecker.CheckActionState returns call-results in pStateData. Otherwise, a call to AsyncActionStateChecker.CheckActionState returns error in pStateData.
//
//	Exceptions:
//	Throws	exception in case of error, see Some error definitions
//
//	 +------------------------------------------------+------+--------------------------+--------------------------------------------------+
//	 |                  Module name                   | Code |      Mnemonic name       |                Error description                 |
//	 +------------------------------------------------+------+--------------------------+--------------------------------------------------+
//	 | KLPRSS                                         | 1125 | KLPRSS.NOT_EXIST        | Parameter with the specified name does not exist |
//	 | KLSTD                                          | 1184 | KLSTD.STDE_NOACCESS     | Access denied                                    |
//	 | KLSTD                                          | 1186 | KLSTD.STDE_NOTFOUND     | Object not found                                 |
//	 | KLSTD                                          | 1193 | KLSTD.STDE_NOTPERM      | Operation is not permitted                       |
//	 | KLCONN                                         | 1194 | KLCONN.GCR_NO_SUCH_CALL |  No such GUI call (GUI call is not implemented)  |
//	 +------------------------------------------------+------+--------------------------+--------------------------------------------------+
func (ngc *NagGuiCalls) CallConnectorAsync(ctx context.Context, params interface{}) ([]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", ngc.client.Server+"/api/v1.0/NagGuiCalls.CallConnectorAsync", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ngc.client.Do(ctx, request, nil)
	return raw, err
}
