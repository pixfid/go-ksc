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
	"log"
	"net/http"
)

type GroupSyncIterator struct {
	client *Client
}

//Releases the result-set.
//
//Releases the specified result-set and frees associated memory
//
//Parameters:
//	- szwIterator	(string) result-set ID
func (ca *GroupSyncIterator) ReleaseIterator(ctx context.Context, szwIterator string) bool {
	postData := []byte(fmt.Sprintf(`
	{
	"szwIterator": "%s"
	}`, szwIterator))

	request, err := http.NewRequest("POST", ca.client.Server+"/api/v1.0/GroupSyncIterator.ReleaseIterator", bytes.NewBuffer(postData))

	if err != nil {
		log.Fatal(err.Error())
	}

	jsonData, err := ca.client.Do(ctx, request, nil, false)

	if jsonData != nil {
		return true
	}
	return false
}

type LicenseKeysData struct {
	BEOF      *bool     `json:"bEOF,omitempty"`
	PData     *KeysData `json:"pData,omitempty"`
	PxgRetVal *int64    `json:"PxgRetVal,omitempty"`
}

type KeysData struct {
	KeysDataArray []KeysDataArray `json:"KLCSP_ITERATOR_ARRAY"`
}

type KeysDataArray struct {
	Type    *string  `json:"type,omitempty"`
	KeyData *KeyData `json:"value,omitempty"`
}

type KeyData struct {
	KllicsrvAutokey            bool      `json:"KLLICSRV_AUTOKEY,omitempty"`
	KllicsrvKeyInstalled       bool      `json:"KLLICSRV_KEY_INSTALLED,omitempty"`
	KllicAppID                 int64     `json:"KLLIC_APP_ID,omitempty"`
	KllicCreationDate          KllicDate `json:"KLLIC_CREATION_DATE,omitempty"`
	KllicCustomerInfo          string    `json:"KLLIC_CUSTOMER_INFO,omitempty"`
	KllicKeyType               int64     `json:"KLLIC_KEY_TYPE,omitempty"`
	KllicLicensePeriod         int64     `json:"KLLIC_LICENSE_PERIOD,omitempty"`
	KllicLicinfo               string    `json:"KLLIC_LICINFO,omitempty"`
	KllicLictypeIsSubscription bool      `json:"KLLIC_LICTYPE_IS_SUBSCRIPTION,omitempty"`
	KllicLicCount              int64     `json:"KLLIC_LIC_COUNT,omitempty"`
	KllicLimitDate             KllicDate `json:"KLLIC_LIMIT_DATE,omitempty"`
	KllicMajVer                string    `json:"KLLIC_MAJ_VER,omitempty"`
	KllicNearestExpirationDate KllicDate `json:"KLLIC_NEAREST_EXPIRATION_DATE,omitempty"`
	KllicNhostsAscurrent       int64     `json:"KLLIC_NHOSTS_ASCURRENT,omitempty"`
	KllicNhostsAsnext          int64     `json:"KLLIC_NHOSTS_ASNEXT,omitempty"`
	KllicProdName              string    `json:"KLLIC_PROD_NAME,omitempty"`
	KllicProdSuiteID           int64     `json:"KLLIC_PROD_SUITE_ID,omitempty"`
	KllicSerial                string    `json:"KLLIC_SERIAL,omitempty"`
	KllicSubscrinfoEnddatetype int64     `json:"KLLIC_SUBSCRINFO_ENDDATETYPE,omitempty"`
	KllicSubscrinfoGraceterm   int64     `json:"KLLIC_SUBSCRINFO_GRACETERM,omitempty"`
	KllicSubscrinfoProviderurl string    `json:"KLLIC_SUBSCRINFO_PROVIDERURL,omitempty"`
	KllicSubscrinfoState       int64     `json:"KLLIC_SUBSCRINFO_STATE,omitempty"`
	KllicSupportInfo           string    `json:"KLLIC_SUPPORT_INFO,omitempty"`
}

type KllicDate struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

func (u KeyData) KeyType() string {
	switch u.KllicKeyType {
	case 1:
		return "Commercial"
	case 2:
		return "Beta"
	case 3:
		return "Trial"
	case 4:
		return "Test"
	case 5:
		return "OEM"
	case 6:
		return "Subscription"
	default:
		return "Unknown"
	}
}

func (ca *GroupSyncIterator) GetNextItems(ctx context.Context, szwIterator string, nCount int64) (*LicenseKeysData,
	[]byte, error) {
	postData := []byte(fmt.Sprintf(`
	{
		"szwIterator": "%s",
		"nCount": %d
	}`, szwIterator, nCount))

	request, err := http.NewRequest("POST", ca.client.Server+"/api/v1.0/GroupSyncIterator.GetNextItems", bytes.NewBuffer(postData))

	if err != nil {
		log.Fatal(err.Error())
	}

	licenseKeysData := new(LicenseKeysData)

	raw, err := ca.client.Do(ctx, request, licenseKeysData, false)

	return licenseKeysData, raw, err
}
