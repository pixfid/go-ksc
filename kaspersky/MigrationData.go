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
	"fmt"
	"net/http"
)

//	MigrationData Class Reference
//
//	Detailed Description:
//	Migration of data between KSC On-Premise and KSCHosted.
//	Migration of data from KSC On-Premise to KSCHosted
//
//	List of all members.
type MigrationData service

//KnownProducts struct
type KnownProducts struct {
	KnownProductsVal []KnownProductsVal `json:"PxgRetVal"`
}

type KnownProductsVal struct {
	Type            *string          `json:"type,omitempty"`
	KnownProductVal *KnownProductVal `json:"value,omitempty"`
}

type KnownProductVal struct {
	KlmigrProductInfoDN      *string `json:"KLMIGR_PRODUCT_INFO_DN,omitempty"`
	KlmigrProductInfoName    *string `json:"KLMIGR_PRODUCT_INFO_NAME,omitempty"`
	KlmigrProductInfoVersion *string `json:"KLMIGR_PRODUCT_INFO_VERSION,omitempty"`
}

//	Acquire list of known products for migration.
//
//	Acquire list of known products for migration
//
//	Returns:
//	- Array of (params). Each element contains:
//	|-KLMIGR_PRODUCT_INFO_NAME - (paramString) Product name
//	|-KLMIGR_PRODUCT_INFO_VERSION - (paramString) Product version
//	|-KLMIGR_PRODUCT_INFO_DN - (paramString) Product display name
//
//	Exceptions:
//	- Throws	exception in case of error
func (md *MigrationData) AcquireKnownProducts(ctx context.Context) (*KnownProducts, []byte, error) {
	request, err := http.NewRequest("POST", md.client.Server+"/api/v1.0/MigrationData.AcquireKnownProducts",
		nil)
	if err != nil {
		return nil, nil, err
	}

	knownProducts := new(KnownProducts)
	raw, err := md.client.Do(ctx, request, &knownProducts)
	return knownProducts, raw, err
}

//	Cancels export operation.
//
//	Interrupts and cancels export operation at any time
//
//	Parameters:
//	- wstrActionGuid	(string) - async action GUID of export operation, returned by MigrationData.Export() method
func (md *MigrationData) CancelExport(ctx context.Context, wstrActionGuid string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"wstrActionGuid": "%s"}`, wstrActionGuid))
	request, err := http.NewRequest("POST", md.client.Server+"/api/v1.0/MigrationData.CancelExport",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := md.client.Do(ctx, request, nil)
	return raw, err
}

//TODO MigrationData.Export
//TODO MigrationData.Import
//TODO MigrationData.InitFileUpload
