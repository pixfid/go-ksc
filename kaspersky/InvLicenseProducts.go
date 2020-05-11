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
	"context"
	"net/http"
)

//	InvLicenseProducts Class Reference
//
//	Interface to manage License Management (third party) Functionality. More...
//
//	List of all members.
type InvLicenseProducts service

func (ilp *InvLicenseProducts) GetLicenseProducts(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", ilp.client.Server+"/api/v1.0/InvLicenseProducts.GetLicenseProducts", nil)
	if err != nil {
		return nil, err
	}

	raw, err := ilp.client.Do(ctx, request, nil)
	return raw, err
}
