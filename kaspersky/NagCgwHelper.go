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
	"fmt"

	"net/http"
)

//	NagCgwHelper Class Reference
//
//	Nagent CGW (Connection Gateway) API
//
//	This interface is implemented at Network Agent side,
//	so use gateway connection to connect Network Agent and call interface methods.
//
//	List of all members.
type NagCgwHelper service

//	Retrieves product's component location.
//
//	Parameters:
//	- szwProduct	Product name *
//	- szwVersion	Product version *
//	- szwComponent	Component name
//
//	+-----------------------------------------------------------------+--------------------------+-----------------------------+
//	|                             Product                             | "KLHST_WKS_PRODUCT_NAME" | "KLHST_WKS_PRODUCT_VERSION" |
//	+-----------------------------------------------------------------+--------------------------+-----------------------------+
//	| Kaspersky Administration Server                                 | "1093"                   | "1.0.0.0"                   |
//	| Kaspersky Network Agent                                         | "1103"                   | "1.0.0.0"                   |
//	| Kaspersky Anti-Virus 5.0 for Windows Workstation                | "Workstation"            | "5.0.0.0"                   |
//	| Kaspersky Anti-Virus 5.0 for Windows File Server                | "Fileserver"             | "5.0.0.0"                   |
//	| Kaspersky Anti-Virus 6.0 for Windows Workstation                | "KAVWKS6"                | "6.0.0.0"                   |
//	| Kaspersky Anti-Virus 6.0 for Windows File Server                | "KAVFS6"                 | "6.0.0.0"                   |
//	| Kaspersky Anti-Virus 6.0 for Windows Workstation MP1            | "KAVWKS6"                | "6.0.4.0"                   |
//	| Kaspersky Anti-Virus 6.0 for Windows File Server MP1            | "KAVFS6"                 | "6.0.4.0"                   |
//	| Kaspersky Anti-Virus 6.0 for Windows Servers Enterprise Edition | "KAVFSEE"                | "6.0.0.0"                   |
//	| Kaspersky Anti-Virus 8.0 for Windows Servers Enterprise Edition | "KAVFSEE"                | "8.0.0.0"                   |
//	| Kaspersky Anti-Virus 6.0 for Linux File Server                  | "KLIN"                   | "6.0.0.0"                   |
//	| Kaspersky Enterprise Security 8.0                               | "KES"                    | "8.1.0.0"                   |
//	| Kaspersky Enterprise Security 10.0                              | "KES"                    | "10.1.0.0"                  |
//	| Kaspersky Enterprise Security 10.0 Maintenance Release 1        | "KES"                    | "10.2.1.0"                  |
//	| Kaspersky Enterprise Security 10.0 Service Pack 1               | "KES"                    | "10.2.2.0"                  |
//	| Kaspersky Security for Virtualization Light Agent 3.0           | "KSVLA"                  | "3.0.0.0"                   |
//	| Kaspersky Endpoint Security 10 MR1                              | "KES"                    | "10.2.1.0"                  |
//	| Kaspersky Endpoint Security 10 SP1                              | "KES"                    | "10.2.2.0"                  |
//	| Kaspersky Endpoint Security 10 SP1 MR1                          | "KES"                    | "10.2.4.0"                  |
//	| Kaspersky Endpoint Security 11                                  | "KES"                    | "11.0.0.0"                  |
//	+-----------------------------------------------------------------+--------------------------+-----------------------------+
//
//	Returns:
//	- (params) Location params (non-transparent for a user).
func (nc *NagCgwHelper) GetProductComponentLocation(ctx context.Context, szwProduct, szwVersion,
	szwComponent string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{
		"szwProduct": "%s",
		"szwVersion": "%s",
		"szwComponent": "%s"
	}`, szwProduct, szwVersion, szwComponent))
	request, err := http.NewRequest("POST", nc.client.Server+"/api/v1.0/NagCgwHelper.GetProductComponentLocation", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := nc.client.Do(ctx, request, nil)
	return raw, err
}
