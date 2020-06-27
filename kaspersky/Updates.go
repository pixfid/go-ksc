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

// Updates processing.
type Updates service

// AvailableUpdates containing well-known retranslated update components list
type AvailableUpdates struct {
	PAvailableUpdateComps *PAvailableUpdateComps `json:"pAvailableUpdateComps,omitempty"`
}

// PAvailableUpdateComps
type PAvailableUpdateComps struct {
	KLUpdates []KLUpdate `json:"KLUpdates"`
}

// KLUpdate
type KLUpdate struct {
	Type *string `json:"type,omitempty"`
	// KLUpdateValue containing well-known retranslated update components list,
	// which can be used for update retranslation task tuning
	KLUpdateValue *KLUpdateValue `json:"value,omitempty"`
}

type KLUpdateValue struct {
	// AppID Application identity, only anti-spam components, e.g.: "1184"
	AppID *string `json:"AppId,omitempty"`

	// AppSuperID Application family identity, for anti-spam components only, possible values:
	//	0 - Unknown
	//	1 - AntiSpam Personal 1.0
	//	2 - AntiSpam Personal 1.1
	//	3 - Security for SMTPGateway 5.5
	//	4 - Security for Exchange 2003
	//	5 - AntiSpam 3.0
	AppSuperID *string `json:"AppSuperId,omitempty"`

	// CompID Component internal name, e.g.: "KDB"
	CompID *string `json:"CompId,omitempty"`

	// Descr Component description
	Descr *string `json:"Descr,omitempty"`

	// Name Component title
	Name *string `json:"Name,omitempty"`
}

// GetAvailableUpdatesInfo Get available updates info.
// strLocalization can be one of "ru", "en", "fr", "de"; for other values "en" localization will be used.
func (upd *Updates) GetAvailableUpdatesInfo(ctx context.Context, strLocalization string) (*AvailableUpdates, error) {
	postData := []byte(fmt.Sprintf(`{"strLocalization": "%s"}`, strLocalization))
	request, err := http.NewRequest("POST", upd.client.Server+"/api/v1.0/Updates.GetAvailableUpdatesInfo",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	availableUpdates := new(AvailableUpdates)
	_, err = upd.client.Do(ctx, request, &availableUpdates)
	return availableUpdates, err
}

// UpdatesInfoParams
type UpdatesInfoParams struct {
	PFilter []string `json:"pFilter"`
}

// UpdatesInfos
type UpdatesInfos struct {
	UpdatesInfo []UpdatesInfo `json:"PxgRetVal"`
}

// UpdatesInfo
type UpdatesInfo struct {
	Type             *string           `json:"type,omitempty"`
	UpdatesInfoValue *UpdatesInfoValue `json:"value,omitempty"`
}

// UpdatesInfoValue
type UpdatesInfoValue struct {
	ChangeStatus            *string   `json:"ChangeStatus,omitempty"`
	CompID                  *string   `json:"CompId,omitempty"`
	Date                    *DateTime `json:"Date,omitempty"`
	FileName                *string   `json:"FileName,omitempty"`
	Index                   *string   `json:"Index,omitempty"`
	KlupdsrvBundleDwlDate   *DateTime `json:"KLUPDSRV_BUNDLE_DWL_DATE,omitempty"`
	KlupdsrvBundleID        *string   `json:"KLUPDSRV_BUNDLE_ID,omitempty"`
	KlupdsrvBundleTypeDescr *string   `json:"KLUPDSRV_BUNDLE_TYPE_DESCR,omitempty"`
	LocalPath               *string   `json:"LocalPath,omitempty"`
	RelativeSrvPath         *string   `json:"RelativeSrvPath,omitempty"`
	Type                    *string   `json:"Type,omitempty"`
	Size                    *Size     `json:"Size,omitempty"`
	Stt                     *string   `json:"Stt,omitempty"`
}

// GetUpdatesInfo lists of retransmissions.
func (upd *Updates) GetUpdatesInfo(ctx context.Context, params UpdatesInfoParams) (*UpdatesInfos, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", upd.client.Server+"/api/v1.0/Updates.GetUpdatesInfo", bytes.NewBuffer(postData))

	if err != nil {
		return nil, err
	}

	updatesInfos := new(UpdatesInfos)
	_, err = upd.client.Do(ctx, request, &updatesInfos)
	return updatesInfos, err
}

// RemoveUpdates Asynchronously remove updates.
func (upd *Updates) RemoveUpdates(ctx context.Context) (*PxgValStr, error) {
	request, err := http.NewRequest("POST", upd.client.Server+"/api/v1.0/Updates.RemoveUpdates", nil)
	if err != nil {
		return nil, err
	}

	pxgValStr := new(PxgValStr)
	_, err = upd.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, err
}

// RemoveUpdatesCancel Cancel asynchronous operation RemoveUpdates.
func (upd *Updates) RemoveUpdatesCancel(ctx context.Context, strRequestId string) error {
	postData := []byte(fmt.Sprintf(`{"strRequestId": "%s"}`, strRequestId))
	request, err := http.NewRequest("POST", upd.client.Server+"/api/v1.0/Updates.RemoveUpdatesCancel", bytes.NewBuffer(postData))
	if err != nil {
		return err
	}

	_, err = upd.client.Do(ctx, request, nil)
	return err
}
