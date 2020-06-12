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
	"fmt"
	"net/http"
)

//	PatchParameters Class Reference
//	Patch parameters processing.
//
//	List of all members.
type PatchParameters service

//	Get template for command.
//
//	Parameters:
//	- patchID	(int64) Identification of patch in vapm
//	- locID		(int64) Patch LCID
//
//	Returns:
//	- parTemplate (params)
func (pp *PatchParameters) GetTemplate(ctx context.Context, patchID, locID int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"patchID": %d, "locID": %d}`, patchID, locID))
	request, err := http.NewRequest("POST", pp.client.Server+"/api/v1.0/PatchParameters.GetTemplate", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pp.client.Do(ctx, request, nil)
	return raw, err
}

//	Get install command for patch.
//
//	Parameters:
//	- patchID	(int64) Identification of patch in vapm
//	- locID		(int64) Identification of localization of patch in vapm
//
//	Returns:
//	- parsValues (params)
func (pp *PatchParameters) GetValues(ctx context.Context, patchID, locID int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"patchID": %d, "locID": %d}`, patchID, locID))
	request, err := http.NewRequest("POST", pp.client.Server+"/api/v1.0/PatchParameters.GetValues", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pp.client.Do(ctx, request, nil)
	return raw, err
}

//	Get install command for patch.
//
//	Parameters:
//	- packageId	(int64) Identification of package
//
//	Returns:
//	- parsValues (params)
func (pp *PatchParameters) GetValuesByPkg(ctx context.Context, packageId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"packageId": %d}`, packageId))
	request, err := http.NewRequest("POST", pp.client.Server+"/api/v1.0/PatchParameters.GetValuesByPkg", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pp.client.Do(ctx, request, nil)
	return raw, err
}

//	Set values for parameters of command.
//
//	Parameters:
//	- patchID	(int64) Identification of patch in vapm
//	- locID		(int64) Identification of localization of patch in vapm
//	- parsValues	(params)
func (pp *PatchParameters) SetValues(ctx context.Context, params interface{}) ([]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", pp.client.Server+"/api/v1.0/PatchParameters.SetValues", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pp.client.Do(ctx, request, nil)
	return raw, err
}

//	Set values for parameters of command.
//
//	Parameters:
//	- packageId		(int64) Identification of package
//	- parsValues	(params)
func (pp *PatchParameters) SetValuesByPkg(ctx context.Context, params interface{}) ([]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", pp.client.Server+"/api/v1.0/PatchParameters.SetValuesByPkg", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pp.client.Do(ctx, request, nil)
	return raw, err
}
