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

// ScanDiapasons Network subnets processing.
type ScanDiapasons service

// NotifyDpnsTask Restart the task scanning IP diapasons.
func (sd *ScanDiapasons) NotifyDpnsTask(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", sd.client.Server+"/api/v1.0/ScanDiapasons.NotifyDpnsTask", nil)
	if err != nil {
		return nil, err
	}

	raw, err := sd.client.Request(ctx, request, nil)
	return raw, err
}

//DiapasonsParams struct using in ScanDiapasons.GetDiapasons
type DiapasonsParams struct {
	// VecFieldsToReturn array of diapasons attribute names to return.
	VecFieldsToReturn []string `json:"vecFieldsToReturn"`

	//LMaxLifeTime max result-set lifetime in seconds, not more than 7200
	LMaxLifeTime int64 `json:"lMaxLifeTime,omitempty"`
}

// GetDiapasons Enumerate existing diapasons.
func (sd *ScanDiapasons) GetDiapasons(ctx context.Context, params DiapasonsParams) (*PxgValStr, []byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", sd.client.Server+"/api/v1.0/ScanDiapasons.GetDiapasons", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := sd.client.Request(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

// DiapasonParams struct using in ScanDiapasons.GetDiapason
type DiapasonParams struct {
	IDDiapason int64    `json:"idDiapason,omitempty"`
	PFields    []string `json:"pFields"`
}

// DiapasonAttributes struct
type DiapasonAttributes struct {
	DAttributes *DAttributes `json:"PxgRetVal,omitempty"`
}

//DAttributes struct network diapason attributes
type DAttributes struct {
	// KldpnsDN Diapason display name
	KldpnsDN *string `json:"KLDPNS_DN,omitempty"`

	// KldpnsID Diapason id
	KldpnsID *int64 `json:"KLDPNS_ID,omitempty"`

	// KldpnsIls Array of ip intervals or subnets descriptions.
	KldpnsIls []KldpnsIL `json:"KLDPNS_ILS"`

	// KldpnsLF Ip address validity period in seconds
	KldpnsLF *int64 `json:"KLDPNS_LF,omitempty"`

	// KLDPNSScanEnabled If diapason may be scanned by ip subnets scanning
	KLDPNSScanEnabled *bool `json:"KLDPNS_ScanEnabled,omitempty"`
}

// KldpnsIL struct ip intervals or subnets descriptions
type KldpnsIL struct {
	// Type "params"
	Type *string `json:"type,omitempty"`

	//KldpnsILValue Value
	KldpnsILValue *KldpnsILValue `json:"value,omitempty"`
}

type KldpnsILValue struct {
	// KldpnsILIssubnet true if subnet and false if ip interval
	KldpnsILIssubnet *bool `json:"KLDPNS_IL_ISSUBNET,omitempty"`

	// KldpnsILMaskorlow subnet mask (if KLDPNS_IL_ISSUBNET is true) or low ip of interval end (if KLDPNS_IL_ISSUBNET is false)
	// in TCP/IP network byte order. Subnet mask must be contigious.
	KldpnsILMaskorlow *int64 `json:"KLDPNS_IL_MASKORLOW,omitempty"`

	// KldpnsILSubnetorhi subnet address (if KLDPNS_IL_ISSUBNET is true) or high interval end (if KLDPNS_IL_ISSUBNET is false)
	// in TCP/IP network byte order. Subnet address must correspond subnet mask.
	KldpnsILSubnetorhi *int64 `json:"KLDPNS_IL_SUBNETORHI,omitempty"`
}

// GetDiapason Acquire specified diapason attributes.
//
// Returns specified attributes of given diapason.
func (sd *ScanDiapasons) GetDiapason(ctx context.Context, params DiapasonParams) (*DiapasonAttributes, []byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", sd.client.Server+"/api/v1.0/ScanDiapasons.GetDiapason", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	diapasonAttributes := new(DiapasonAttributes)
	raw, err := sd.client.Request(ctx, request, &diapasonAttributes)
	return diapasonAttributes, raw, err
}

// RemoveDiapason Removes specified diapason.
func (sd *ScanDiapasons) RemoveDiapason(ctx context.Context, idDiapason int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"idDiapason": %d}`, idDiapason))
	request, err := http.NewRequest("POST", sd.client.Server+"/api/v1.0/ScanDiapasons.RemoveDiapason", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := sd.client.Request(ctx, request, nil)
	return raw, err
}

//UpdateDiapasonParams struct using in ScanDiapasons.UpdateDiapason
type UpdateDiapasonParams struct {
	// IDDiapason id of diapason to modify
	IDDiapason int64 `json:"idDiapason,omitempty"`
	// UDPInfo container with diapason attributes
	UDPInfo UDPInfo `json:"pInfo,omitempty"`
}

type UDPInfo struct {
	// KldpnsDN Diapason display name
	KldpnsDN string `json:"KLDPNS_DN,omitempty"`

	// KldpnsLF Ip address validity period in seconds
	KldpnsLF int64 `json:"KLDPNS_LF,omitempty"`

	// KLDPNSScanEnabled If diapason may be scanned by ip subnets scanning
	KLDPNSScanEnabled bool `json:"KLDPNS_ScanEnabled,omitempty"`

	//TODO KLDPNS_ILS
}

type UpdateDiapasonRespond struct {
	PInvalidIntervals *PInvalidIntervals `json:"pInvalidIntervals,omitempty"`
	PxgRetVal         *bool              `json:"PxgRetVal,omitempty"`
}

type PInvalidIntervals struct {
	//TODO WTF???
}

// UpdateDiapason Change one or more attributes of diapason.
//
// Performs update of one or more attributes of diapason.
// If at least one of diapasons intersects with any of existing diapasons or is invalid then
// false is returned and such interval is added to "KLDPNS_ILS" array in ppInvalidIntervals.
func (sd *ScanDiapasons) UpdateDiapason(ctx context.Context, params UpdateDiapasonParams) (*UpdateDiapasonRespond, []byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", sd.client.Server+"/api/v1.0/ScanDiapasons.UpdateDiapason", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	updateDiapasonRespond := new(UpdateDiapasonRespond)
	raw, err := sd.client.Request(ctx, request, &updateDiapasonRespond)
	return updateDiapasonRespond, raw, err
}

// AddDiapason Adds new diapason with the specified attributes (see List of network diapason attributes).
// Error occurs if at least one of intervals/subnet intersects with any of existing intervals/subnets.
func (sd *ScanDiapasons) AddDiapason(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", sd.client.Server+"/api/v1.0/ScanDiapasons.AddDiapason", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := sd.client.Request(ctx, request, nil)
	return raw, err
}
