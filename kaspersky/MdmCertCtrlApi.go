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
	"log"
	"net/http"
)

// MdmCertCtrlApi Mobile devices certificates and restore data management.
type MdmCertCtrlApi service

// CancelGeneratePackage2 Cancels asynchronous operation SetCertificateAsync.GeneratePackageAsync2.
//	wstrRequestId - identity of asynchronous operation
func (mca *MdmCertCtrlApi) CancelGeneratePackage2(ctx context.Context, wstrRequestId string) error {
	postData := []byte(fmt.Sprintf(`{"wstrRequestId": "%s"}`, wstrRequestId))
	request, err := http.NewRequest("POST", mca.client.Server+"/api/v1.0/MdmCertCtrlApi.CancelGeneratePackage2",
		bytes.NewBuffer(postData))
	if err != nil {
		return err
	}

	raw, err := mca.client.Request(ctx, request, nil)

	if mca.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return err
}

// CancelSetCertificate2 Cancels asynchronous operation SetCertificateAsync.SetCertificateAsync2.
//	wstrRequestId - identity of asynchronous operation
func (mca *MdmCertCtrlApi) CancelSetCertificate2(ctx context.Context, wstrRequestId string) error {
	postData := []byte(fmt.Sprintf(`{"wstrRequestId": "%s"}`, wstrRequestId))
	request, err := http.NewRequest("POST", mca.client.Server+"/api/v1.0/MdmCertCtrlApi.CancelSetCertificate2",
		bytes.NewBuffer(postData))
	if err != nil {
		return err
	}

	raw, err := mca.client.Request(ctx, request, nil)

	if mca.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return err
}

// CheckMailNotificationSettings Check email settings It defines if virtual server's email settings
//should be checked in merge (as defaults) with main server's email settings
func (mca *MdmCertCtrlApi) CheckMailNotificationSettings(ctx context.Context, bCheckMainServerDefaults bool) (*PxgValBool, error) {
	postData := []byte(fmt.Sprintf(`{"bCheckMainServerDefaults": %v}`, bCheckMainServerDefaults))
	request, err := http.NewRequest("POST", mca.client.Server+"/api/v1.0/MdmCertCtrlApi.CheckMailNotificationSettings",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	smtpServerEmpty := new(PxgValBool)
	raw, err := mca.client.Request(ctx, request, &smtpServerEmpty)

	if mca.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return smtpServerEmpty, err
}

//TODO CheckPkiAccount -

// CheckPkiEnabled Check if PKI enabled.
func (mca *MdmCertCtrlApi) CheckPkiEnabled(ctx context.Context) (*PxgValStr, error) {
	request, err := http.NewRequest("POST", mca.client.Server+"/api/v1.0/MdmCertCtrlApi.CheckPkiEnabled", nil)
	if err != nil {
		return nil, err
	}

	pkiFlag := new(PxgValStr)
	raw, err := mca.client.Request(ctx, request, &pkiFlag)

	if mca.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return pkiFlag, err
}

//TODO DeleteCertificates
//TODO GeneratePackageAsync2

// GetCertificatePublic certificate id to get
//	nCertId - Retrieve certificate's public part.
func (mca *MdmCertCtrlApi) GetCertificatePublic(ctx context.Context, nCertId int64) error {
	postData := []byte(fmt.Sprintf(`{"nCertId": %d}`, nCertId))
	request, err := http.NewRequest("POST", mca.client.Server+"/api/v1.0/MdmCertCtrlApi.GetCertificatePublic",
		bytes.NewBuffer(postData))
	if err != nil {
		return err
	}

	raw, err := mca.client.Request(ctx, request, nil)

	if mca.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return err
}

type IssuanceSettings struct {
	// IssuanceSettingVal array of issuance settings
	PxgRetVal []IssuanceSettingVal `json:"PxgRetVal"`
}

type IssuanceSettingVal struct {
	Type  string               `json:"type"`
	Value IssuanceSettingValue `json:"value"`
}

type IssuanceSetting struct {
	// IssuanceSettingValue array of issuance settings
	PxgRetVal IssuanceSettingValue `json:"PxgRetVal"`
}

type IssuanceSettingValue struct {
	// CIBAutoIssuanceDisabled Is certificate auto-issuance allowed
	CIBAutoIssuanceDisabled bool `json:"CI_bAutoIssuanceDisabled"`
	// CIBAutorenewalEnabled If autorenewal enabled
	CIBAutorenewalEnabled bool `json:"CI_bAutorenewalEnabled"`
	// CIBEncryptionEnabled Is certificate encryption feature enabled
	CIBEncryptionEnabled bool `json:"CI_bEncryptionEnabled"`
	// CIBFeatureActive Is issuance feature active
	CIBFeatureActive bool `json:"CI_bFeatureActive"`
	// CINCERTType certificate type
	CINCERTType int64 `json:"CI_nCertType"`
	// CINEncryptionPwdLength certificate password length for certificate encryption feature
	CINEncryptionPwdLength int64 `json:"CI_nEncryptionPwdLength"`
	// CINExpiryPeriod Certificate expiry period
	CINExpiryPeriod int64 `json:"CI_nExpiryPeriod"`
	// CINRenewalPeriodSEC Certificate renewal period
	CINRenewalPeriodSEC int64 `json:"CI_nRenewalPeriodSec"`
	// CIWstrPKICERTTemplateName PKI certificate template name
	CIWstrPKICERTTemplateName string `json:"CI_wstrPkiCertTemplateName"`
}

// GetIssuanceSettings Retrieve saved issuance settings for certificate types.
func (mca *MdmCertCtrlApi) GetIssuanceSettings(ctx context.Context) (*IssuanceSettings, error) {
	request, err := http.NewRequest("POST", mca.client.Server+"/api/v1.0/MdmCertCtrlApi.GetIssuanceSettings", nil)
	if err != nil {
		return nil, err
	}

	issuanceSettings := new(IssuanceSettings)
	raw, err := mca.client.Request(ctx, request, &issuanceSettings)

	if mca.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return issuanceSettings, err
}

// GetIssuanceSettingsByType Retrieve saved issuance settings for specific certificate type.
//	nCertType Certificate type (0 - Unknown, 1 - User, 2 - EMail, 3 - VPN)
func (mca *MdmCertCtrlApi) GetIssuanceSettingsByType(ctx context.Context, nCertType int64) (*IssuanceSetting, error) {
	postData := []byte(fmt.Sprintf(`{"nCertType": %d}`, nCertType))
	request, err := http.NewRequest("POST", mca.client.Server+"/api/v1.0/MdmCertCtrlApi.GetIssuanceSettingsByType",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	issuanceSetting := new(IssuanceSetting)
	raw, err := mca.client.Request(ctx, request, &issuanceSetting)

	if mca.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return issuanceSetting, err
}

//TODO GetPackagesByPackageProduct

// GetPkiTemplates Get PKI templates.
//	bForceReload (bool) if True then templates will be reloaded from AD, if False then it will be loaded from cache
func (mca *MdmCertCtrlApi) GetPkiTemplates(ctx context.Context, bForceReload bool) error {
	postData := []byte(fmt.Sprintf(`{"bForceReload": %v}`, bForceReload))
	request, err := http.NewRequest("POST", mca.client.Server+"/api/v1.0/MdmCertCtrlApi.GetPkiTemplates",
		bytes.NewBuffer(postData))
	if err != nil {
		return err
	}

	raw, err := mca.client.Request(ctx, request, nil)

	if mca.client.Debug {
		log.Printf("raw response: %s", string(raw))
	}

	return err
}

//TODO GetPkiTemplatesByEKU
//TODO GetRestoreData
//TODO RenewCertificate
//TODO SetCertificateAsync2
//TODO SetCertificateTag
//TODO UpdateIssuanceSettings
