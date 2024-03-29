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

//	InvLicenseProducts service to manage License Management (third party) Functionality.
type InvLicenseProducts service

type LicenseKeysResponse struct {
	PxgRetVal *LKResponse `json:"PxgRetVal,omitempty"`
}

type LKResponse struct {
	KlinvlicLicProductsArray []KlinvlicLicProductsArray `json:"KLINVLIC_LIC_PRODUCTS_ARRAY,omitempty"`
}

type KlinvlicLicProductsArray struct {
	Type  *string                        `json:"type,omitempty"`
	Value *KLINVLICLICPRODUCTSARRAYValue `json:"value,omitempty"`
}

type KLINVLICLICPRODUCTSARRAYValue struct {
	KlinvlicData *KlinvlicData `json:"KLINVLIC_DATA,omitempty"`
	KlinvlicID   *int64        `json:"KLINVLIC_ID,omitempty"`
}

type KlinvlicData struct {
	Type  *string       `json:"type,omitempty"`
	Value *PLicProdData `json:"value,omitempty"`
}

// GetLicenseProducts Acquire License Products data.
func (ilp *InvLicenseProducts) GetLicenseProducts(ctx context.Context) (*LicenseKeysResponse, error) {
	request, err := http.NewRequest("POST", ilp.client.Server+"/api/v1.0/InvLicenseProducts.GetLicenseProducts", nil)
	if err != nil {
		return nil, err
	}

	licenseKeysResponse := new(LicenseKeysResponse)
	_, err = ilp.client.Request(ctx, request, &licenseKeysResponse)
	return licenseKeysResponse, err
}

// DeleteLicenseKey Removes specified License Key.
func (ilp *InvLicenseProducts) DeleteLicenseKey(ctx context.Context, nLicKeyId int64) (*PxgRetError, error) {
	postData := []byte(fmt.Sprintf(`{"nLicKeyId": %d}`, nLicKeyId))
	request, err := http.NewRequest("POST", ilp.client.Server+"/api/v1.0/InvLicenseProducts.DeleteLicenseKey", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	pxgRetError := new(PxgRetError)
	_, err = ilp.client.Request(ctx, request, &pxgRetError)
	return pxgRetError, err
}

// DeleteLicenseProduct Removes specified License Product.
func (ilp *InvLicenseProducts) DeleteLicenseProduct(ctx context.Context, nLicProdId int64) (*PxgRetError, error) {
	postData := []byte(fmt.Sprintf(`{"nLicProdId": %d}`, nLicProdId))
	request, err := http.NewRequest("POST", ilp.client.Server+"/api/v1.0/InvLicenseProducts.DeleteLicenseProduct", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	pxgRetError := new(PxgRetError)
	_, err = ilp.client.Request(ctx, request, &pxgRetError)
	return pxgRetError, err
}

// LicenseKeyParams struct
type LicenseKeyParams struct {
	PLicKeyData *PLicKeyData `json:"pLicKeyData,omitempty"`
}

type PLicKeyData struct {
	KlinvlicKeyName            string       `json:"KLINVLIC_KEY_NAME,omitempty"`
	KlinvlicKeyLiccount        int64        `json:"KLINVLIC_KEY_LICCOUNT,omitempty"`
	KlinvlicKeyCreation        *KlinvlicKey `json:"KLINVLIC_KEY_CREATION,omitempty"`
	KlinvlicKeyExpirationlimit *KlinvlicKey `json:"KLINVLIC_KEY_EXPIRATIONLIMIT,omitempty"`
	KlinvlicKeyInfo            string       `json:"KLINVLIC_KEY_INFO,omitempty"`
}

type KlinvlicKey struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

// AddLicenseKey Add a new License Key.
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
func (ilp *InvLicenseProducts) AddLicenseKey(ctx context.Context, params LicenseKeyParams) (*PxgValInt, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", ilp.client.Server+"/api/v1.0/InvLicenseProducts.AddLicenseKey", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	pxgValInt := new(PxgValInt)
	_, err = ilp.client.Request(ctx, request, &pxgValInt)
	return pxgValInt, err
}

// LicenseProductParams struct
type LicenseProductParams struct {
	PLicProdData PLicProdData `json:"pLicProdData,omitempty"`
}

type PLicProdData struct {
	KlinvlicActive bool            `json:"KLINVLIC_ACTIVE,omitempty"`
	KlinvlicDesc   string          `json:"KLINVLIC_DESC,omitempty"`
	KlinvlicLikeys []int64         `json:"KLINVLIC_LIKEYS"`
	KlinvlicMasks  []KlinvlicMasks `json:"KLINVLIC_MASKS,omitempty"`
	KlinvlicName   string          `json:"KLINVLIC_NAME,omitempty"`
	KlinvlicType   int64           `json:"KLINVLIC_TYPE,omitempty"`
}

type KlinvlicMasks struct {
	Type               string             `json:"type,omitempty"`
	KlinvlicMasksValue KlinvlicMasksValue `json:"value"`
}

type KlinvlicMasksValue struct {
	KlinvlicMaskProdDispNameFilter  string `json:"KLINVLIC_MASK_PROD_DISP_NAME_FILTER,omitempty"`
	KlinvlicMaskProdDispVerFilter   string `json:"KLINVLIC_MASK_PROD_DISP_VER_FILTER,omitempty"`
	KlinvlicMaskProdPublisherFilter string `json:"KLINVLIC_MASK_PROD_PUBLISHER_FILTER,omitempty"`
	KlinvlicMaskProdTagValue        string `json:"KLINVLIC_MASK_PROD_TAG_VALUE,omitempty"`
}

// AddLicenseProduct Add a new License Product.
func (ilp *InvLicenseProducts) AddLicenseProduct(ctx context.Context, params LicenseProductParams) (*PxgValInt, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", ilp.client.Server+"/api/v1.0/InvLicenseProducts.AddLicenseProduct", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	pxgValInt := new(PxgValInt)
	_, err = ilp.client.Request(ctx, request, &pxgValInt)
	return pxgValInt, err
}

// UpdateLicenseKeyParams struct using in InvLicenseProducts.UpdateLicenseKey
type UpdateLicenseKeyParams struct {
	//NLicKeyID id of License Key to modify
	NLicKeyID int64 `json:"nLicKeyId"`

	//PLicKeyData object containing License Key attributes to modify
	PLicKeyData PLicKeyData `json:"pLicKeyData"`
}

// UpdateLicenseKey Modifies attributes of specified License Key.
func (ilp *InvLicenseProducts) UpdateLicenseKey(ctx context.Context, params UpdateLicenseKeyParams) error {
	postData, err := json.Marshal(params)
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", ilp.client.Server+"/api/v1.0/InvLicenseProducts.UpdateLicenseKey", bytes.NewBuffer(postData))
	if err != nil {
		return err
	}

	_, err = ilp.client.Request(ctx, request, nil)
	return err
}

// UpdateLicenseProductParams struct using in InvLicenseProducts.UpdateLicenseProduct
type UpdateLicenseProductParams struct {
	// NLicProdID id of License Product to modify
	NLicProdID string `json:"nLicProdId"`

	// PLicProdData object containing License Product attributes to modify
	PLicProdData PLicProdData `json:"pLicProdData"`
}

// UpdateLicenseProduct Modifies attributes of specified License Product.
func (ilp *InvLicenseProducts) UpdateLicenseProduct(ctx context.Context, params UpdateLicenseProductParams) error {
	postData, err := json.Marshal(params)
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", ilp.client.Server+"/api/v1.0/InvLicenseProducts.UpdateLicenseProduct", bytes.NewBuffer(postData))
	if err != nil {
		return err
	}

	_, err = ilp.client.Request(ctx, request, nil)
	return err
}
