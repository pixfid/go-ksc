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

// SsContents service provide access to settings storage.
type SsContents service

// SsContent struct
type SsContent struct {
	// WstrID identifier of opened SsContents
	WstrID string `json:"wstrID"`

	// WstrProduct product name string, non-empty string, not longer than 31 character,
	//and cannot contain characters /\:*?"<>.
	WstrProduct string `json:"wstrProduct"`

	// WstrVersion version string, non-empty string, not longer than 31 character,
	//and cannot contain characters /\:*?"<>.
	WstrVersion string `json:"wstrVersion"`

	// WstrSection section name string, non-empty string, not longer than 31 character,
	//and cannot contain characters /\:*?"<>.
	WstrSection string `json:"wstrSection"`

	// PNewData new data to write
	PNewData interface{} `json:"pNewData"`
}

// SsAdd Add new data to settings storage. Adds new variables to the specified section.
//	Changes are not saved until method SsApply is called. Unsaved data is not available by methods SsRead and SSGetNames.
func (sc *SsContents) SsAdd(ctx context.Context, params SsContent) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", sc.client.Server+"/api/v1.0/SsContents.Ss_Add", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := sc.client.Request(ctx, request, nil)
	return raw, err
}

// SsApply	Saves changes made by methods: SsUpdate, SsAdd, SsReplace, SsClear, SsDelete, SsCreateSection, SsDeleteSection
func (sc *SsContents) SsApply(ctx context.Context, wstrID string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`	{ "wstrID": "%s"}`, wstrID))
	request, err := http.NewRequest("POST", sc.client.Server+"/api/v1.0/SsContents.Ss_Apply", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := sc.client.Request(ctx, request, nil)
	return raw, err
}

// SsClear Clear and write data in settings storage.
//	Replaces existing section contents with pData, i.e. existing section contents will deleted and variables
//	from pData will be written to the section.
//	Changes are not saved until method SsApply is called. Unsaved data is not available by methods SsRead and SSGetNames.
func (sc *SsContents) SsClear(ctx context.Context, params SsContent) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", sc.client.Server+"/api/v1.0/SsContents.Ss_Clear", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := sc.client.Request(ctx, request, nil)
	return raw, err
}

// SsCreateSection Create section in settings storage.
//	Creates empty section in settings storage
//	Changes are not saved until method SsApply is called. Unsaved data is not available by methods SsRead and SSGetNames.
func (sc *SsContents) SsCreateSection(ctx context.Context, params SsContentD) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", sc.client.Server+"/api/v1.0/SsContents.Ss_CreateSection", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := sc.client.Request(ctx, request, nil)
	return raw, err
}

// SsContentD struct
type SsContentD struct {
	// WstrID identifier of opened SsContents
	WstrID string `json:"wstrID"`

	// WstrProduct product name string, non-empty string, not longer than 31 character,
	//and cannot contain characters /\:*?"<>.
	WstrProduct string `json:"wstrProduct"`

	// WstrVersion version string, non-empty string, not longer than 31 character,
	//and cannot contain characters /\:*?"<>.
	WstrVersion string `json:"wstrVersion,omitempty"`

	// WstrSection section name string, non-empty string, not longer than 31 character,
	//and cannot contain characters /\:*?"<>.
	WstrSection string `json:"wstrSection,omitempty"`

	// PData data
	PData interface{} `json:"pData,omitempty"`
}

// SsDelete Delete data from settings storage.
//	Deletes variables specified in pData from the specified section.
//	Changes are not saved until method SsApply is called. Unsaved data is not available by methods SsRead and SSGetNames.
func (sc *SsContents) SsDelete(ctx context.Context, params SsContentD) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", sc.client.Server+"/api/v1.0/SsContents.Ss_Delete", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := sc.client.Request(ctx, request, nil)
	return raw, err
}

// SsDeleteSection Delete section from settings storage.
//	Deletes section with all contents from settings storage.
//	Changes are not saved until method SsApply is called.
//	Unsaved data is not available by methods SsRead and SSGetNames.
func (sc *SsContents) SsDeleteSection(ctx context.Context, params SsContentD) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", sc.client.Server+"/api/v1.0/SsContents.Ss_DeleteSection", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := sc.client.Request(ctx, request, nil)
	return raw, err
}

// SSGetNames Enumerate contents of settings storage
//	Retrieves list of sections.
func (sc *SsContents) SSGetNames(ctx context.Context, params SsContentD) (*PxgValArrayOfString, []byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", sc.client.Server+"/api/v1.0/SsContents.SS_GetNames", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValArrayOfString := new(PxgValArrayOfString)
	raw, err := sc.client.Request(ctx, request, &pxgValArrayOfString)
	return pxgValArrayOfString, raw, err
}

// SsRead Reads saved data from the specified section of settings storage.
func (sc *SsContents) SsRead(ctx context.Context, params SsContentD, v interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", sc.client.Server+"/api/v1.0/SsContents.Ss_Read", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := sc.client.Request(ctx, request, &v)
	return raw, err
}

// SsRelease Closes opened SsContents and releases associated server resources.
//	After calling this method wstrID is not longer valid.
func (sc *SsContents) SsRelease(ctx context.Context, wstrID string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"wstrID": "%s"}`, wstrID))
	request, err := http.NewRequest("POST", sc.client.Server+"/api/v1.0/SsContents.Ss_Release", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := sc.client.Request(ctx, request, nil)
	return raw, err
}

// SsReplace Replaces variables in the specified section.
//	If a variable already exists it will be updated, if a variable does not exist it will be added.
//	Changes are not saved until method Ss_Apply is called. Unsaved data is not available by methods SsRead and SSGetNames.
func (sc *SsContents) SsReplace(ctx context.Context, params SsContent) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", sc.client.Server+"/api/v1.0/SsContents.Ss_Replace", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := sc.client.Request(ctx, request, nil)
	return raw, err
}

// SsUpdate Updates existing variables in the specified section.
//	Changes are not saved until method SsApply is called. Unsaved data is not available by methods SsRead and SSGetNames.
func (sc *SsContents) SsUpdate(ctx context.Context, params SsContent) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", sc.client.Server+"/api/v1.0/SsContents.Ss_Update", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := sc.client.Request(ctx, request, nil)
	return raw, err
}
