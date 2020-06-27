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

// MigrationData service provide to:
//
// Migration of data between KSC On-Premise and KSCHosted.
//
// Migration of data from KSC On-Premise to KSCHosted
type MigrationData service

// KnownProducts struct
type KnownProducts struct {
	KnownProductsVal []KnownProductsVal `json:"PxgRetVal"`
}

type KnownProductsVal struct {
	Type            string           `json:"type,omitempty"`
	KnownProductVal *KnownProductVal `json:"value,omitempty"`
}

type KnownProductVal struct {
	KlmigrProductInfoDN      string `json:"KLMIGR_PRODUCT_INFO_DN,omitempty"`
	KlmigrProductInfoName    string `json:"KLMIGR_PRODUCT_INFO_NAME,omitempty"`
	KlmigrProductInfoVersion string `json:"KLMIGR_PRODUCT_INFO_VERSION,omitempty"`
}

// AcquireKnownProducts Acquire list of known products for migration.
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

// CancelExport Interrupts and cancels export operation at any time by async action GUID of export operation,
// returned by MigrationData.Export method
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

// MDExportParams struct
type MDExportParams struct {
	// MDPOptions Export options, which include the following (see below). All options are mandatory!
	MDPOptions *MDPOptions `json:"pOptions,omitempty"`
}

// MDPOptions struct
type MDPOptions struct {
	// KlmigrRootGroupID Root group identifier
	KlmigrRootGroupID int64 `json:"KLMIGR_ROOT_GROUP_ID,omitempty"`

	// KlmigrArrIDReports Identifiers of report templates
	KlmigrArrIDReports []int64 `json:"KLMIGR_ARR_ID_REPORTS"`

	// KlmigrArrIDCmnTasks Non-group tasks identifiers
	KlmigrArrIDCmnTasks []int64 `json:"KLMIGR_ARR_ID_CMN_TASKS"`

	// KlmigrArrIDExtraQrs Extra (user-created) queries identifiers
	KlmigrArrIDExtraQrs []int64 `json:"KLMIGR_ARR_ID_EXTRA_QRS"`

	// KlmigrArrProductsInfo Array of params, where each element contains the following:
	KlmigrArrProductsInfo []KlmigrArrProductsInfo `json:"KLMIGR_ARR_PRODUCTS_INFO"`

	// KlmigrSkipCustomRoles Skip import of custom roles
	KlmigrSkipCustomRoles bool `json:"KLMIGR_SKIP_CUSTOM_ROLES,omitempty"`

	// KlmigrSkipUsersGroups Skip import of internal users and security groups
	KlmigrSkipUsersGroups bool `json:"KLMIGR_SKIP_USERS_GROUPS,omitempty"`

	// KlmigrSkipCustomAppCategories Skip import of custom application categories
	KlmigrSkipCustomAppCategories bool `json:"KLMIGR_SKIP_CUSTOM_APP_CATEGORIES,omitempty"`
}

type KlmigrArrProductsInfo struct {
	Type              string            `json:"Type,omitempty"`
	ProductsInfoValue ProductsInfoValue `json:"Value"`
}

type ProductsInfoValue struct {
	// KlmigrProductInfoName Name of product, for which tasks and policies should be exported
	KlmigrProductInfoName string `json:"KLMIGR_PRODUCT_INFO_NAME,omitempty"`

	// KlmigrProductInfoVersion Version of product, for which tasks and policies should be exported
	KlmigrProductInfoVersion string `json:"KLMIGR_PRODUCT_INFO_VERSION,omitempty"`
}

// Export Performs export of objects. Exports all objects, specified in pOptions, and returns async action GUID,
// which can be used later to retrieve file URL. To provide more consistency, method can also export some additional "child" objects,
// which are referred to by "parent" objects, specified by user.
//
// As a result, user receives GUID to async action, containing URL to zip-archive with all exported data
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

// InitFileUpload Retrieves URL for zip archive upload. Retrieves URL for upload zip archive with exported data.
// To correctly upload zip archive, which was received after export, one should do the following:
//
// 1. Acquire async action GUID from MigrationData.Export method
//
// 2. Get URL from async action GUID by calling AsyncActionStateChecker.CheckActionState - URL-path to zip archive will be in in pStateData
//
// 3. Download zip file from URL, received in previous step
//
// 4. Call method MigrationData.InitFileUpload to receive URL for upload (this method)
//
// 5. Upload zip archive to URL, retrieved in previous step, using HTTP PUT request.
//
// After all above is done, you can call MigrationData.Import to perform import
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

// ImportMDParams struct using in MigrationData.Import
type ImportMDParams struct {
	// WstrURL upload URL. Use MigrationData.InitFileUpload() method to obtain it
	WstrURL string `json:"wstrUrl"`

	// IOptions import options
	IOptions IOptions `json:"pOptions"`
}

// IOptions struct
type IOptions struct {
	// KlmigrRootGroupID root group identifier
	KlmigrRootGroupID int64 `json:"KLMIGR_ROOT_GROUP_ID"`
}

// Import Performs import of objects. Imports all objects, specified in pOptions, from file, pointed by upload URL.
//
// Method is asynchronious. To correctly use this method, first call InitFileUpload() to obtain file URL.
// If wstrUrl is invalid, method fails with error.
//
// Remarks:
// Check the operation state by calling AsyncActionStateChecker.CheckActionState periodically until it's finalized.
// If the operation succeeds then AsyncActionStateChecker. CheckActionState returns bFinalized=true and lStateCode=1 in pStateData.
// Otherwise, a call to AsyncActionStateChecker.CheckActionState returns error in pStateData.
func (md *MigrationData) Import(ctx context.Context, params ImportMDParams) (*PxgValStr, []byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", md.client.Server+"/api/v1.0/MigrationData.Import", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := md.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}
