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
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type LicenseKeys struct {
	client *Client
}

//TODO Download license key files from activation key servers V1.
//
//Parameters:
//strActivationCode	(string) activation code in format XXXXX-XXXXX-XXXXX-XXXXX, mandatory.
//Returns:
//(array) of paramBinary, array of license key files related to the specified code.
//Exceptions:
//Throws	exception in case of error.
func (lk *LicenseKeys) DownloadKeyFiles(ctx context.Context, wstrActivationCode string) bool {
	postData := []byte(fmt.Sprintf(`
	{
	"wstrActivationCode": "%s"
	}`, wstrActivationCode))

	request, err := http.NewRequest("POST", lk.client.Server+"/api/v1.0/LicenseKeys.DownloadKeyFiles",
		bytes.NewBuffer(postData))

	if err != nil {
		log.Fatal(err.Error())
	}

	jsonData, err := lk.client.Do(ctx, request, nil)

	if jsonData != nil {
		return true
	}
	return false
}

//	EnumKeysParams struct
type EnumKeysParams struct {
	PFields     []string `json:"pFields"`
	LTimeoutSEC int64    `json:"lTimeoutSec,omitempty"`
}

//	LicenseKeysIterator struct
type LicenseKeysIterator struct {
	LKeyCount    int64  `json:"lKeyCount,omitempty"`
	WstrIterator string `json:"wstrIterator,omitempty"`
}

//	{
//	"pFields": [
//			"KLLIC_APP_ID",
//			"KLLIC_APP_IDS",
//			"KLLIC_CREATION_DATE",
//			"KLLIC_CUSTOMER_INFO",
//			"KLLIC_KEY_TYPE",
//			"KLLIC_LIC_COUNT",
//			"KLLIC_LICENSE_PERIOD",
//			"KLLIC_LICINFO",
//			"KLLIC_LICTYPE_IS_SUBSCRIPTION",
//			"KLLIC_LIMIT_DATE",
//			"KLLIC_MAJ_VER",
//			"KLLIC_NEAREST_EXPIRATION_DATE",
//			"KLLIC_NHOSTS_ASCURRENT",
//			"KLLIC_NHOSTS_ASNEXT",
//			"KLLIC_PROD_NAME",
//			"KLLIC_PROD_SUITE_ID",
//			"KLLIC_SERIAL",
//			"KLLIC_SUBSCRINFO_ENDDATE",
//			"KLLIC_SUBSCRINFO_ENDDATETYPE",
//			"KLLIC_SUBSCRINFO_GRACETERM",
//			"KLLIC_SUBSCRINFO_PROVIDERURL",
//			"KLLIC_SUBSCRINFO_STATE",
//			"KLLIC_SUPPORT_INFO",
//			"KLLICSRV_AUTOKEY",
//			"KLLICSRV_KEY_INSTALLED",
//			"KLLICSRV_KEYDATA"
//	],
//	"lTimeoutSec": 130
//	}
func (lk *LicenseKeys) EnumKeys(ctx context.Context, ekp EnumKeysParams) (*LicenseKeysIterator, []byte, error) {
	postData, _ := json.Marshal(ekp)
	request, err := http.NewRequest("POST", lk.client.Server+"/api/v1.0/LicenseKeys.EnumKeys",
		bytes.NewBuffer(postData))
	if err != nil {
		log.Fatal(err.Error())
	}
	licenseKeysIterator := new(LicenseKeysIterator)
	raw, err := lk.client.Do(ctx, request, &licenseKeysIterator)

	return licenseKeysIterator, raw, err
}

//KeyDataParams struct
type KeyDataParams struct {
	PKeyInfo PKeyInfo `json:"pKeyInfo,omitempty"`
}

type PKeyInfo struct {
	KllicSerial     string `json:"KLLIC_SERIAL,omitempty"`
	KllicsrvKeydata bool   `json:"KLLICSRV_KEYDATA,omitempty"`
}

//LicenseData struct
type LicenseData struct {
	Data *LData `json:"PxgRetVal,omitempty"`
}

type LData struct {
	KllicsrvAutokey                  bool            `json:"KLLICSRV_AUTOKEY,omitempty"`
	KllicsrvKeydata                  KllicsrvKeydata `json:"KLLICSRV_KEYDATA,omitempty"`
	KllicsrvKeyEffectiveLicnumberMax int64           `json:"KLLICSRV_KEY_EFFECTIVE_LICNUMBER_MAX,omitempty"`
	KllicsrvKeyEffectiveLicnumberMin int64           `json:"KLLICSRV_KEY_EFFECTIVE_LICNUMBER_MIN,omitempty"`
	KllicsrvKeyInstalled             bool            `json:"KLLICSRV_KEY_INSTALLED,omitempty"`
	KllicAppID                       int64           `json:"KLLIC_APP_ID,omitempty"`
	KllicAppIDS                      KllicAppIDS     `json:"KLLIC_APP_IDS,omitempty"`
	KllicCreationDate                Kllic           `json:"KLLIC_CREATION_DATE,omitempty"`
	KllicCustomerInfo                string          `json:"KLLIC_CUSTOMER_INFO,omitempty"`
	KllicEffectiveLicensePeriod      int64           `json:"KLLIC_EFFECTIVE_LICENSE_PERIOD,omitempty"`
	KllicEffectiveLicCount           int64           `json:"KLLIC_EFFECTIVE_LIC_COUNT,omitempty"`
	KllicKeyType                     int64           `json:"KLLIC_KEY_TYPE,omitempty"`
	KllicLicensePeriod               int64           `json:"KLLIC_LICENSE_PERIOD,omitempty"`
	KllicLicinfo                     string          `json:"KLLIC_LICINFO,omitempty"`
	KllicLicobject1                  int64           `json:"KLLIC_LICOBJECT_1,omitempty"`
	KllicLicobject102                int64           `json:"KLLIC_LICOBJECT_102,omitempty"`
	KllicLicobject110                int64           `json:"KLLIC_LICOBJECT_110,omitempty"`
	KllicLicobject54                 int64           `json:"KLLIC_LICOBJECT_54,omitempty"`
	KllicLicobject58                 int64           `json:"KLLIC_LICOBJECT_58,omitempty"`
	KllicLicobject60                 int64           `json:"KLLIC_LICOBJECT_60,omitempty"`
	KllicLicobject62                 int64           `json:"KLLIC_LICOBJECT_62,omitempty"`
	KllicLicobject68                 int64           `json:"KLLIC_LICOBJECT_68,omitempty"`
	KllicLicobject70                 int64           `json:"KLLIC_LICOBJECT_70,omitempty"`
	KllicLicobject72                 int64           `json:"KLLIC_LICOBJECT_72,omitempty"`
	KllicLictypeIsSubscription       bool            `json:"KLLIC_LICTYPE_IS_SUBSCRIPTION,omitempty"`
	KllicLicCount                    int64           `json:"KLLIC_LIC_COUNT,omitempty"`
	KllicLicIsPendingSubscription    bool            `json:"KLLIC_LIC_IS_PENDING_SUBSCRIPTION,omitempty"`
	KllicLimitDate                   Kllic           `json:"KLLIC_LIMIT_DATE,omitempty"`
	KllicMajVer                      string          `json:"KLLIC_MAJ_VER,omitempty"`
	KllicNhostsAscurrent             int64           `json:"KLLIC_NHOSTS_ASCURRENT,omitempty"`
	KllicNhostsAsnext                int64           `json:"KLLIC_NHOSTS_ASNEXT,omitempty"`
	KllicProdName                    string          `json:"KLLIC_PROD_NAME,omitempty"`
	KllicProdSuiteID                 int64           `json:"KLLIC_PROD_SUITE_ID,omitempty"`
	KllicSerial                      string          `json:"KLLIC_SERIAL,omitempty"`
	KllicSubscrinfoEnddatetype       int64           `json:"KLLIC_SUBSCRINFO_ENDDATETYPE,omitempty"`
	KllicSubscrinfoGraceterm         int64           `json:"KLLIC_SUBSCRINFO_GRACETERM,omitempty"`
	KllicSubscrinfoProviderurl       string          `json:"KLLIC_SUBSCRINFO_PROVIDERURL,omitempty"`
	KllicSubscrinfoState             int64           `json:"KLLIC_SUBSCRINFO_STATE,omitempty"`
	KllicSubscrinfoStatereason       int64           `json:"KLLIC_SUBSCRINFO_STATEREASON,omitempty"`
	KllicSupportInfo                 string          `json:"KLLIC_SUPPORT_INFO,omitempty"`
}

type KllicAppIDS struct {
	Value KLLICAPPIDSValue `json:"value,omitempty"`
}

type KLLICAPPIDSValue struct {
	The1105 The1105 `json:"1105,omitempty"`
}

type The1105 struct {
	Value map[string]int64 `json:"value,omitempty"`
}

type Kllic struct {
	Value string `json:"value,omitempty"`
}

type KllicsrvKeydata struct {
	Value KLLICSRVKEYDATAValue `json:"value,omitempty"`
}

type KLLICSRVKEYDATAValue struct {
	KllicLicfile *Kllic `json:"KLLIC_LICFILE,omitempty"`
}

func (u KeyData) ExpirationDate() string {
	if u.KllicNearestExpirationDate.Value == "" {
		return "Not Received"
	} else {
		return ParseTime(u.KllicNearestExpirationDate.Value)
	}
}

//{
//  			"pKeyInfo" : {
//    		"KLLIC_SERIAL" : "%s",
//    		"KLLICSRV_KEYDATA" : true
//  			}
//		}
func (lk *LicenseKeys) GetKeyData(ctx context.Context, kdp KeyDataParams) (*LicenseData, []byte, error) {
	postData, _ := json.Marshal(kdp)

	request, err := http.NewRequest("POST", lk.client.Server+"/api/v1.0/LicenseKeys.GetKeyData",
		bytes.NewBuffer(postData))

	licenseData := new(LicenseData)
	raw, err := lk.client.Do(ctx, request, &licenseData)
	return licenseData, raw, err
}
