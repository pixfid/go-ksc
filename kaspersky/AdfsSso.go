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

//	AdfsSso Class Reference
//
//	Interface for working with ADFS SSO. More...
//
//	This interface allow you to manage ADFS SSO settings
//
//	List of all members.
type AdfsSso service

func (as *AdfsSso) GetSettings(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", as.client.Server+"/api/v1.0/AdfsSso.GetSettings", nil)
	if err != nil {
		return nil, err
	}

	raw, err := as.client.Do(ctx, request, nil)
	return raw, err
}

//	Set a ADFS SSO settings.
//
//	Parameters:
//	- pAdfsSettings	(params) ADFS SSO settings; "ADFS SSO Settings".
func (as *AdfsSso) SetSettings(ctx context.Context, params interface{}) ([]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", as.client.Server+"/api/v1.0/AdfsSso.SetSettings", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := as.client.Do(ctx, request, nil)
	return raw, err
}
