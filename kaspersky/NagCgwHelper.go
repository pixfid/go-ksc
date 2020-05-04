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
	"log"
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
type NagCgwHelper struct {
	client *Client
}

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
		log.Fatal(err.Error())
	}

	raw, err := nc.client.Do(ctx, request, nil)
	return raw, err
}
