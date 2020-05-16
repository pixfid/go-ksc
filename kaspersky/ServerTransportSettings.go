/*
 *
 * 	Copyright (C) 2020  <Semchenko Aleksandr>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful, but
 * WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.See the GNU
 * General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.If not, see <http://www.gnu.org/licenses/>.
 * /
 */

package kaspersky

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
)

//	ServerTransportSettings Class Reference
//
//	Server transport settings proxy class..
//
//	List of all members.
type ServerTransportSettings service

//	GetNumberOfManagedDevicesAgentless.
//	Returns number of agentless managed devices.
//
//	Note: It can be called from main server only !
//
//	Returns:
//	It returns total number of managed devices for main server and all virtual servers.
func (sts *ServerTransportSettings) GetNumberOfManagedDevicesAgentless(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", sts.client.Server+"/api/v1.0/ServerTransportSettings.GetNumberOfManagedDevicesAgentless", nil)

	raw, err := sts.client.Do(ctx, request, nil)
	return raw, err
}

//	GetNumberOfManagedDevicesKSM.
//	Returns number of managed devices for KSM (Kaspersky for Mobile).
//
//	Note: It can be called from main server only !
//
//	Returns:
//	It returns total number of managed devices for main server and all virtual servers.
func (sts *ServerTransportSettings) GetNumberOfManagedDevicesKSM(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", sts.client.Server+"/api/v1.0/ServerTransportSettings.GetNumberOfManagedDevicesKSM", nil)

	raw, err := sts.client.Do(ctx, request, nil)
	return raw, err
}

//	IsFeatureActive. Checks if feature is activated and certificate can be changed to some custom value.
//
//	Parameters:
//	- szwCertType	Certificate type (can be "CERT_TYPE_MOBILE", "CERT_TYPE_EMBEDDED")
//
//	Returns:
//	- True if feature is active; false otherwise.
func (sts *ServerTransportSettings) IsFeatureActive(ctx context.Context, szwCertType string) (*PxgValBool, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"szwCertType": "%s"}`, szwCertType))
	request, err := http.NewRequest("POST", sts.client.Server+"/api/v1.0/ServerTransportSettings.IsFeatureActive", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValBool := new(PxgValBool)
	raw, err := sts.client.Do(ctx, request, &pxgValBool)
	return pxgValBool, raw, err
}

//	SetFeatureActive. Sets feature active.
//
//	Parameters:
//	- szwCertType	Certificate type. For "CERT_TYPE_MOBILE" only!
//	- bFeatureActive	Should be true to activate feature.
func (sts *ServerTransportSettings) SetFeatureActive(ctx context.Context, szwCertType string,
	bFeatureActive bool) (*PxgValBool, []byte,
	error) {
	postData := []byte(fmt.Sprintf(`{"szwCertType": "%s", "bFeatureActive" : %v}`, szwCertType, bFeatureActive))
	request, err := http.NewRequest("POST", sts.client.Server+"/api/v1.0/ServerTransportSettings.SetFeatureActive", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValBool := new(PxgValBool)
	raw, err := sts.client.Do(ctx, request, &pxgValBool)
	return pxgValBool, raw, err
}

//	CheckDefaultCertificateExists. It checks if default certificate exists.
//
//	Parameters:
//	- szwCertType	Certificate type (can be "CERT_TYPE_MOBILE", "CERT_TYPE_EMBEDDED")
//
//	Returns:
//	- True if certificate exists; false otherwise.
func (sts *ServerTransportSettings) CheckDefaultCertificateExists(ctx context.Context, szwCertType string) (*PxgValBool, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"szwCertType": "%s"}`, szwCertType))
	request, err := http.NewRequest("POST", sts.client.Server+"/api/v1.0/ServerTransportSettings.CheckDefaultCertificateExists", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValBool := new(PxgValBool)
	raw, err := sts.client.Do(ctx, request, &pxgValBool)
	return pxgValBool, raw, err
}

//CurrentConnectionSettings struct
type CurrentConnectionSettings struct {
	CCSettings *CCSettings `json:"PxgRetVal,omitempty"`
}

type CCSettings struct {
	CERTPub                   *CERTPub `json:"CERT_PUB,omitempty"`
	TrspSettingsFQDN          *string  `json:"TRSP_SETTINGS_FQDN,omitempty"`
	TrspSettingsIsdefcertused *bool    `json:"TRSP_SETTINGS_ISDEFCERTUSED,omitempty"`
	TrspSettingsOpenPort      *bool    `json:"TRSP_SETTINGS_OPEN_PORT,omitempty"`
	TrspSettingsPort          *int64   `json:"TRSP_SETTINGS_PORT,omitempty"`
}

type CERTPub struct {
	Type  *string `json:"type,omitempty"`
	Value *string `json:"value,omitempty"`
}

//	GetCurrentConnectionSettings. Returns current connection settings.
//
//	Parameters:
//	- szwCertType	Certificate type (can be "CERT_TYPE_MOBILE", "CERT_TYPE_EMBEDDED")
//
//	Returns:
//	Returned data format:
//                "TRSP_SETTINGS_FQDN"                [paramString], actual endpoint FQDN (from certificate)
//                "TRSP_SETTINGS_PORT"                [paramInt], actual enpoint port
//                "TRSP_SETTINGS_ISDEFCERTUSED"       [paramBool], is default certificate used ?
//                "TRSP_SETTINGS_OPEN_PORT"           [paramBool], true if port should be opened, false otherwise.
//                "CERT_PUB"                          [paramBinary], current certificate's public key
//                "TRSP_RESERVE_CERT_PUB"             [paramBinary], optional, reserve certificate's public key
//                "TRSP_RESERVE_CERT_ACTIVATION_DATE" [DATETIME_T], optional, reserve certificate's activation date
func (sts *ServerTransportSettings) GetCurrentConnectionSettings(ctx context.Context, szwCertType string) (*CurrentConnectionSettings, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"szwCertType": "%s"}`, szwCertType))
	request, err := http.NewRequest("POST", sts.client.Server+"/api/v1.0/ServerTransportSettings.GetCurrentConnectionSettings", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	cCSettings := new(CurrentConnectionSettings)
	raw, err := sts.client.Do(ctx, request, &cCSettings)
	return cCSettings, raw, err
}

//	GetCustomSrvCertificateInfo. Returns information about custom certificate.
//
//	Parameters:
//	- szwCertType	Certificate type (can be "CERT_TYPE_MOBILE", "CERT_TYPE_EMBEDDED")
//
//	Returns:
//	If certificate present then params with the fields below will be returned:
//                "CERT_USE"                          [paramBool], it indicates if custom certificate feature is enabled or disabled for this certificate type (szwCertType);
//                "CERT_PUB"                          [paramBinary], certificate's public key;
//                "TRSP_RESERVE_CERT_PUB"             [paramBinary], optional, reserve certificate's public key;
//                "TRSP_RESERVE_CERT_ACTIVATION_DATE" [DATETIME_T], optional, reserve certificate's activation date;
//                In case if custom certificate was not set before then empty params will be returned with no any fields;
//                In case if custom certificate was set but disabled then params with ["CERT_USE"] field set to false will be returned.
//
//	See also:
//	Common format for certificate params.
func (sts *ServerTransportSettings) GetCustomSrvCertificateInfo(ctx context.Context, szwCertType string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"szwCertType": "%s"}`, szwCertType))
	request, err := http.NewRequest("POST", sts.client.Server+"/api/v1.0/ServerTransportSettings.GetCustomSrvCertificateInfo", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := sts.client.Do(ctx, request, nil)
	return raw, err
}

//GetDefaultConnectionSettings. Returns default connection settings.
//
//Parameters:
//szwCertType	Certificate type (can be "CERT_TYPE_MOBILE", "CERT_TYPE_EMBEDDED")
//Returns:
//Returned data format:
//                "TRSP_SETTINGS_FQDN"                [paramString], default endpoint FQDN (from certificate)
//                "TRSP_SETTINGS_PORT"                [paramInt], default enpoint port
//                "TRSP_SETTINGS_OPEN_PORT"           [paramBool], true if port should be opened, false otherwise.
//                "CERT_PUB"                          [paramBinary], current certificate's public key
//                "TRSP_RESERVE_CERT_PUB"             [paramBinary], optional, reserve certificate's public key
//                "TRSP_RESERVE_CERT_ACTIVATION_DATE" [DATETIME_T], optional, reserve certificate's activation date
//
func (sts *ServerTransportSettings) GetDefaultConnectionSettings(ctx context.Context, szwCertType string) (*CurrentConnectionSettings, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"szwCertType": "%s"}`, szwCertType))
	request, err := http.NewRequest("POST", sts.client.Server+"/api/v1.0/ServerTransportSettings.GetDefaultConnectionSettings", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	cCSettings := new(CurrentConnectionSettings)
	raw, err := sts.client.Do(ctx, request, &cCSettings)
	return cCSettings, raw, err
}

//	ResetCstmReserveCertificate. Resets custom reserve certificate.
//
//	Parameters:
//	- szwCertType	Certificate type (can be "CERT_TYPE_MOBILE")
func (sts *ServerTransportSettings) ResetCstmReserveCertificate(ctx context.Context, szwCertType string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"szwCertType": "%s"}`, szwCertType))
	request, err := http.NewRequest("POST", sts.client.Server+"/api/v1.0/ServerTransportSettings.ResetCstmReserveCertificate", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := sts.client.Do(ctx, request, nil)
	return raw, err
}

//	ResetDefaultReserveCertificate. Resets default reserve certificate.
//
//	Parameters:
//	- szwCertType	Certificate type (can be "CERT_TYPE_MOBILE")
func (sts *ServerTransportSettings) ResetDefaultReserveCertificate(ctx context.Context, szwCertType string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"szwCertType": "%s"}`, szwCertType))
	request, err := http.NewRequest("POST", sts.client.Server+"/api/v1.0/ServerTransportSettings.ResetDefaultReserveCertificate", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := sts.client.Do(ctx, request, nil)
	return raw, err
}

//TODO SetCustomSrvCertificate
//TODO SetOrCreateDefaultCertificate
