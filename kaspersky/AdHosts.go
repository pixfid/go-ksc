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

// AdHosts Class service allows to enumerate scanned active directory OU structure.
type AdHosts service

// FindAdGroupsParams struct
type FindAdGroupsParams struct {
	VecFieldsToReturn []string        `json:"vecFieldsToReturn,omitempty"`
	VecFieldsToOrder  []FieldsToOrder `json:"vecFieldsToOrder,empty"`
	POptions          POptions        `json:"pOptions,omitempty"`
	LMaxLifeTime      int64           `json:"lMaxLifeTime,omitempty"`
}

type POptions struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value"`
}

//	ADHostIterator struct
type ADHostIterator struct {
	WstrIterator string `json:"wstrIterator"`
	PxgRetVal    int64  `json:"PxgRetVal"`
}

// FindAdGroups Enumerates AD groups.
func (ah *AdHosts) FindAdGroups(ctx context.Context, params FindAdGroupsParams) (*ADHostIterator, []byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", ah.client.Server+"/api/v1.0/AdHosts.FindAdGroups", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	aDHostIterator := new(ADHostIterator)
	raw, err := ah.client.Request(ctx, request, &aDHostIterator)
	return aDHostIterator, raw, err
}

// ChildComputersParams struct
type ChildComputersParams struct {
	IDOU              int64    `json:"idOU"`
	VecFieldsToReturn []string `json:"vecFieldsToReturn"`
	LMaxLifeTime      int64    `json:"lMaxLifeTime"`
}

// ChildComputerParams struct
type ChildComputerParams struct {
	IDAdhst           int64    `json:"idAdhst,omitempty"`
	VecFieldsToReturn []string `json:"vecFieldsToReturn,omitempty"`
}

// AdHstIDParent struct contain AD host attributes.
type AdHstIDParent struct {
	PxgRetVal AdHstIDParentPxgRetVal `json:"PxgRetVal"`
}

type AdHstIDParentPxgRetVal struct {
	AdhstIDParent int64 `json:"adhst_idParent"`
}

// GetChildComputer Retrieves AD host attributes.
func (ah *AdHosts) GetChildComputer(ctx context.Context, params ChildComputerParams) (*AdHstIDParent,
	[]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", ah.client.Server+"/api/v1.0/AdHosts.GetChildComputer", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	adHstIdParent := new(AdHstIDParent)
	raw, err := ah.client.Request(ctx, request, &adHstIdParent)
	return adHstIdParent, raw, err
}

// GetChildComputers Returns list of hosts located in "Unassigned computers" for specified organization unit.
func (ah *AdHosts) GetChildComputers(ctx context.Context, params ChildComputersParams) (*PxgValStr, []byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", ah.client.Server+"/api/v1.0/AdHosts.GetChildComputers", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := ah.client.Request(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

// ChildOUParams struct
type ChildOUParams struct {
	IDOU         int64    `json:"idOU,omitempty"`
	PFields      []string `json:"pFields,omitempty"`
	LMaxLifeTime int64    `json:"lMaxLifeTime,omitempty"`
}

// GetChildOUs Returns list of child organization units for specified organization unit
func (ah *AdHosts) GetChildOUs(ctx context.Context, params ChildOUParams) (*PxgValStr,
	[]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", ah.client.Server+"/api/v1.0/AdHosts.GetChildOUs", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := ah.client.Request(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

// OUAttributesParams struct
type OUAttributesParams struct {
	IDOU    int64    `json:"idOU,omitempty"`
	PFields []string `json:"pFields"`
}

// OUAttributes struct
type OUAttributes struct {
	Attributes Attributes `json:"PxgRetVal"`
}

type Attributes struct {
	AdhstBinOu            AdhstBinOu `json:"adhst_binOu"`
	AdhstChildSubunitsNum int64      `json:"adhst_childSubunitsNum"`
	AdhstEnableAdScan     bool       `json:"adhst_enable_ad_scan"`
	AdhstHostsNum         int64      `json:"adhst_hostsNum"`
	AdhstID               int64      `json:"adhst_id"`
	AdhstIDComputer       string     `json:"adhst_idComputer"`
	AdhstIDParent         int64      `json:"adhst_idParent"`
}

type AdhstBinOu struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// GetOU Returns attributes of specified OU
func (ah *AdHosts) GetOU(ctx context.Context, params OUAttributesParams) (*OUAttributes, []byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", ah.client.Server+"/api/v1.0/AdHosts.GetOU", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	oUAttributes := new(OUAttributes)
	raw, err := ah.client.Request(ctx, request, &oUAttributes)
	return oUAttributes, raw, err
}

// UpdateOUParams struct
type UpdateOUParams struct {
	//Id of organization unit
	IDOU int64 `json:"idOU,omitempty"`

	//Params
	OUPData *OUPData `json:"pData,omitempty"`
}

// OUPData struct
type OUPData struct {
	//If scanning of this OU is allowed
	AdhstEnableAdScan bool `json:"adhst_enable_ad_scan,omitempty"`
}

// UpdateOU Updates OU properties.
func (ah *AdHosts) UpdateOU(ctx context.Context, params UpdateOUParams) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", ah.client.Server+"/api/v1.0/AdHosts.UpdateOU", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ah.client.Request(ctx, request, nil)
	return raw, err
}
