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

// NlaDefinedNetworks Network location awareness (NLA) defined networks.
//
// Used as a scope for Update agents.
//
// Each NLA-defined network is defined by list of NLA locations.
type NlaDefinedNetworks service

// AddNetwork Add NLA-defined network.
func (ndn *NlaDefinedNetworks) AddNetwork(ctx context.Context, wstrNetworkName string) (*PxgValInt, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"wstrNetworkName": "%s"}`, wstrNetworkName))
	request, err := http.NewRequest("POST", ndn.client.Server+"/api/v1.0/NlaDefinedNetworks.AddNetwork",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValInt := new(PxgValInt)
	raw, err := ndn.client.Request(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

// DeleteNetwork Delete NLA-defined network.
func (ndn *NlaDefinedNetworks) DeleteNetwork(ctx context.Context, nNetworkId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nNetworkId": %d}`, nNetworkId))
	request, err := http.NewRequest("POST", ndn.client.Server+"/api/v1.0/NlaDefinedNetworks.DeleteNetwork",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ndn.client.Request(ctx, request, nil)
	return raw, err
}

// PNetworkInfo struct
type PNetworkInfo struct {
	PNetwork *PNetwork `json:"pNetwork,omitempty"`
}

// PNetwork struct
type PNetwork struct {
	// NlantwkEnableAutoua Enable automatic Update agents assignment to this network flag
	NlantwkEnableAutoua bool `json:"NLANTWK_ENABLE_AUTOUA,omitempty"`

	// NlantwkNetworkID Network id
	NlantwkNetworkID int64 `json:"NLANTWK_NETWORK_ID,omitempty"`

	// NlantwkNetworkName Human-readable network name
	NlantwkNetworkName string `json:"NLANTWK_NETWORK_NAME,omitempty"`
}

// GetNetworkInfo Get NLA-defined network info.
func (ndn *NlaDefinedNetworks) GetNetworkInfo(ctx context.Context, nNetworkId int64) (*PNetworkInfo, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"nNetworkId": %d}`, nNetworkId))
	request, err := http.NewRequest("POST", ndn.client.Server+"/api/v1.0/NlaDefinedNetworks.GetNetworkInfo",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pNetworkInfo := new(PNetworkInfo)
	raw, err := ndn.client.Request(ctx, request, &pNetworkInfo)
	return pNetworkInfo, raw, err
}

// PNetworkList struct
type PNetworkList struct {
	//Each representing NLA-defined network
	PNetworks []PNetworks `json:"pNetworks"`
}

// PNetworks struct
type PNetworks struct {
	Type  string    `json:"type,omitempty"`
	Value *PNetwork `json:"value,omitempty"`
}

// GetNetworksList Get list of all NLA-defined networks.
func (ndn *NlaDefinedNetworks) GetNetworksList(ctx context.Context) (*PNetworkList, []byte, error) {
	request, err := http.NewRequest("POST", ndn.client.Server+"/api/v1.0/NlaDefinedNetworks.GetNetworksList", nil)
	if err != nil {
		return nil, nil, err
	}

	pNetworkList := new(PNetworkList)
	raw, err := ndn.client.Request(ctx, request, &pNetworkList)
	return pNetworkList, raw, err
}

// SetNetworkInfo Change NLA-defined network.
func (ndn *NlaDefinedNetworks) SetNetworkInfo(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", ndn.client.Server+"/api/v1.0/NlaDefinedNetworks.SetNetworkInfo", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ndn.client.Request(ctx, request, nil)
	return raw, err
}
