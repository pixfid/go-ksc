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

//	Multitenancy Class Reference
//
//	Multitenancy product managing.
//
//	List of all members.
type Multitenancy service

//	Retrieves tenant identity.
//
//	Identity is unique for each tenant
//
//	Returns:
//	- (string) tenant id
func (m *Multitenancy) GetTenantId(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", m.client.Server+"/api/v1.0/Multitenancy.GetTenantId", nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	raw, err := m.client.Do(ctx, request, nil)
	return raw, err
}

//	Retrieves multitenancy products available for current tenant.
//
//	Parameters:
//	- strProdName	(string) product name, if set then result will be filtered by this value
//	- strProdVersion	(string) product version, if set then result will be filtered by this value
//
//	Returns:
//	- (array) each element of array contains information about multitenancy product:
//
//	+---------------------------+--------+-------------------------+
//	|           Value           |  Type  |       Description       |
//	+---------------------------+--------+-------------------------+
//	| MTNC_PRODUCT_NAME         | string | product name            |
//	| MTNC_PRODUCT_VERSION      | string | product version         |
//	| MTNC_PRODUCT_DISP_NAME    | string | display product name    |
//	| MTNC_PRODUCT_DISP_VERSION | string | display product version |
//	+---------------------------+--------+-------------------------+
func (m *Multitenancy) GetProducts(ctx context.Context, strProdName, strProdVersion string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`
	{
	"strProdName": "%s",
	"strProdVersion": "%s"
	}`, strProdName, strProdVersion))
	request, err := http.NewRequest("POST", m.client.Server+"/api/v1.0/Multitenancy.GetProducts", bytes.NewBuffer(postData))
	if err != nil {
		log.Fatal(err.Error())
	}

	raw, err := m.client.Do(ctx, request, nil)
	return raw, err
}

//	Get new binary token for current tennant
//
//	Returns:
//	- new token for current tennant
//
//	NotWoking on KSC < 12
func (m *Multitenancy) GetAuthToken(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", m.client.Server+"/api/v1.0/Multitenancy.GetAuthToken", nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	raw, err := m.client.Do(ctx, request, nil)
	return raw, err
}

//TODO func (m *Multitenancy) CheckAuthToken(ctx context.Context, wstrTenantId string, binToken []byte) ([]byte, error)
