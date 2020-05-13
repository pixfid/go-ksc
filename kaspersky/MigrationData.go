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

//	MigrationData Class Reference
//	Detailed Description:
//
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
