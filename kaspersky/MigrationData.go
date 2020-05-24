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

//MDExportParams struct
type MDExportParams struct {
	//Export options, which include the following (see below). All options are mandatory!
	MDPOptions *MDPOptions `json:"pOptions,omitempty"`
}

//MDPOptions struct
type MDPOptions struct {
	//Root group identifier
	KlmigrRootGroupID *int64 `json:"KLMIGR_ROOT_GROUP_ID,omitempty"`

	//Identifiers of report templates
	KlmigrArrIDReports []int64 `json:"KLMIGR_ARR_ID_REPORTS"`

	//Non-group tasks identifiers
	KlmigrArrIDCmnTasks []int64 `json:"KLMIGR_ARR_ID_CMN_TASKS"`

	//Extra (user-created) queries identifiers
	KlmigrArrIDExtraQrs []int64 `json:"KLMIGR_ARR_ID_EXTRA_QRS"`

	//Array of params, where each element contains the following:
	KlmigrArrProductsInfo []KlmigrArrProductsInfo `json:"KLMIGR_ARR_PRODUCTS_INFO"`

	//Skip import of custom roles
	KlmigrSkipCustomRoles *bool `json:"KLMIGR_SKIP_CUSTOM_ROLES,omitempty"`

	//Skip import of internal users and security groups
	KlmigrSkipUsersGroups *bool `json:"KLMIGR_SKIP_USERS_GROUPS,omitempty"`

	//Skip import of custom application categories
	KlmigrSkipCustomAppCategories *bool `json:"KLMIGR_SKIP_CUSTOM_APP_CATEGORIES,omitempty"`
}

type KlmigrArrProductsInfo struct {
	Type              *string           `json:"Type,omitempty"`
	ProductsInfoValue ProductsInfoValue `json:"Value"`
}

type ProductsInfoValue struct {
	//Name of product, for which tasks and policies should be exported
	KlmigrProductInfoName *string `json:"KLMIGR_PRODUCT_INFO_NAME,omitempty"`

	//Version of product, for which tasks and policies should be exported
	KlmigrProductInfoVersion *string `json:"KLMIGR_PRODUCT_INFO_VERSION,omitempty"`
}

//	TODO Fix Request params...
//	Performs export of objects.
//
//	Exports all objects, specified in pOptions, and returns async action GUID,
//	which can be used later to retrieve file URL. To provide more consistency,
//	method can also export some additional "child" objects,
//	which are referred to by "parent" objects, specified by user.
//	As a result, user receives GUID to async action, containing URL to zip-archive with all exported data
func (md *MigrationData) Export(ctx context.Context, params MDExportParams) (*PxgValStr, []byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", md.client.Server+"/api/v1.0/MigrationData.Export", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := md.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

//	Retrieves URL for zip archive upload.
//
//	Retrieves URL for upload zip archive with exported data. To correctly upload zip archive,
//	which was received after export, one should do the following:
//
//	1. Acquire async action GUID from MigrationData::Export method
//	2. Get URL from async action GUID by calling AsyncActionStateChecker.CheckActionState - URL-path to zip archive
//	will be in in pStateData
//	3. Download zip file from URL, received in previous step
//	4. Call method MigrationData.InitFileUpload to receive URL for upload (this method)
//	5. Upload zip archive to URL, retrieved in previous step, using HTTP PUT request.
//
//	After all above is done, you can call MigrationData::Import to perform import
//
//	Returns:
//	- (string) resulting URL for zip-archive
func (md *MigrationData) InitFileUpload(ctx context.Context) (*PxgValStr, []byte, error) {
	request, err := http.NewRequest("POST", md.client.Server+"/api/v1.0/MigrationData.InitFileUpload",
		nil)
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := md.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

//TODO MigrationData.Import
