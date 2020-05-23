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

//	InvLicenseProducts Class Reference
//
//	Interface to manage License Management (third party) Functionality..
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

//	Removes specified License Key.
//
//	Parameters:
//	- nLicKeyId	(int64) id of License Key to remove
func (ilp *InvLicenseProducts) DeleteLicenseKey(ctx context.Context, nLicKeyId int64) (*PxgRetError, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"nLicKeyId": %d}`, nLicKeyId))
	request, err := http.NewRequest("POST", ilp.client.Server+"/api/v1.0/InvLicenseProducts.DeleteLicenseKey", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgRetError := new(PxgRetError)
	raw, err := ilp.client.Do(ctx, request, &pxgRetError)
	return pxgRetError, raw, err
}

//	Removes specified License Product.
//
//	Parameters:
//	- nLicProdId	(int64) id of License Product to remove
func (ilp *InvLicenseProducts) DeleteLicenseProduct(ctx context.Context, nLicProdId int64) (*PxgRetError, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"nLicProdId": %d}`, nLicProdId))
	request, err := http.NewRequest("POST", ilp.client.Server+"/api/v1.0/InvLicenseProducts.DeleteLicenseProduct", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgRetError := new(PxgRetError)
	raw, err := ilp.client.Do(ctx, request, &pxgRetError)
	return pxgRetError, raw, err
}

//LicenseKeyParams struct
type LicenseKeyParams struct {
	PLicKeyData *PLicKeyData `json:"pLicKeyData,omitempty"`
}

type PLicKeyData struct {
	KlinvlicKeyName            *string      `json:"KLINVLIC_KEY_NAME,omitempty"`
	KlinvlicKeyLiccount        *int64       `json:"KLINVLIC_KEY_LICCOUNT,omitempty"`
	KlinvlicKeyCreation        *KlinvlicKey `json:"KLINVLIC_KEY_CREATION,omitempty"`
	KlinvlicKeyExpirationlimit *KlinvlicKey `json:"KLINVLIC_KEY_EXPIRATIONLIMIT,omitempty"`
	KlinvlicKeyInfo            *string      `json:"KLINVLIC_KEY_INFO,omitempty"`
}

type KlinvlicKey struct {
	Type  *string `json:"type,omitempty"`
	Value *string `json:"value,omitempty"`
}

//	Add a new License Key.
//
//	Returns attributes for License Key.
//
//	Parameters:
//	- pLicKeyData	(params) object containing License Key attributes,
//	see List of attributes of software inventory License Key List
//	of attributes of software inventory License Key Following attributes are required:
//	- KLINVLIC_KEY_NAME
//
//	╔════════════════════════════════╦═══════════════╦═══════════════════════════════════════╦══════════╗
//	║              Name              ║     Type      ║              Description              ║ Remarks  ║
//	╠════════════════════════════════╬═══════════════╬═══════════════════════════════════════╬══════════╣
//	║ "KLINVLIC_KEY_NAME"            ║ paramString   ║ License Key name                      ║          ║
//	║ "KLINVLIC_KEY_LICCOUNT"        ║ paramInt      ║ Count of installations allowed by key ║ Optional ║
//	║ "KLINVLIC_KEY_CREATION"        ║ paramDateTime ║ Indicates time when key become active ║ Optional ║
//	║ "KLINVLIC_KEY_EXPIRATIONLIMIT" ║ paramDateTime ║ Indicates time when key expires       ║ Optional ║
//	║ "KLINVLIC_KEY_INFO"            ║ paramString   ║ Description                           ║          ║
//	╚════════════════════════════════╩═══════════════╩═══════════════════════════════════════╩══════════╝
//
//	Returns:
//	- (int64) id of created License Key.
func (ilp *InvLicenseProducts) AddLicenseKey(ctx context.Context, params LicenseKeyParams) (*PxgValInt, []byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", ilp.client.Server+"/api/v1.0/InvLicenseProducts.AddLicenseKey", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValInt := new(PxgValInt)
	raw, err := ilp.client.Do(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

//LicenseProductParams struct
type LicenseProductParams struct {
	PLicProdData *PLicProdData `json:"pLicProdData,omitempty"`
}

type PLicProdData struct {
	KlinvlicActive *bool           `json:"KLINVLIC_ACTIVE,omitempty"`
	KlinvlicDesc   *string         `json:"KLINVLIC_DESC,omitempty"`
	KlinvlicLikeys []int64         `json:"KLINVLIC_LIKEYS"`
	KlinvlicMasks  []KlinvlicMasks `json:"KLINVLIC_MASKS,omitempty"`
	KlinvlicName   *string         `json:"KLINVLIC_NAME,omitempty"`
	KlinvlicType   *int64          `json:"KLINVLIC_TYPE,omitempty"`
}

type KlinvlicMasks struct {
	Type               *string            `json:"type,omitempty"`
	KlinvlicMasksValue KlinvlicMasksValue `json:"value"`
}

type KlinvlicMasksValue struct {
	KlinvlicMaskProdDispNameFilter  *string `json:"KLINVLIC_MASK_PROD_DISP_NAME_FILTER,omitempty"`
	KlinvlicMaskProdDispVerFilter   *string `json:"KLINVLIC_MASK_PROD_DISP_VER_FILTER,omitempty"`
	KlinvlicMaskProdPublisherFilter *string `json:"KLINVLIC_MASK_PROD_PUBLISHER_FILTER,omitempty"`
	KlinvlicMaskProdTagValue        *string `json:"KLINVLIC_MASK_PROD_TAG_VALUE,omitempty"`
}

//	Add a new License Product.
//
//	Returns attributes for License Products.
//
//	- Parameters:
//	- pLicProdData	(params) object containing License Product attributes,
//	see List of attributes of software inventory License Product List
//	of attributes of software inventory License Product. Following attributes are required:
//	- KLINVLIC_NAME
//	- KLINVLIC_ACTIVE
//	- KLINVLIC_MASKS
//	- KLINVLIC_LIKEYS
//
//	Returns:
//	- (int64) id of created License Product.
func (ilp *InvLicenseProducts) AddLicenseProduct(ctx context.Context, params LicenseProductParams) (*PxgValInt, []byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", ilp.client.Server+"/api/v1.0/InvLicenseProducts.AddLicenseProduct", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValInt := new(PxgValInt)
	raw, err := ilp.client.Do(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

//TODO InvLicenseProducts::UpdateLicenseKey
//TODO InvLicenseProducts::UpdateLicenseProduct
