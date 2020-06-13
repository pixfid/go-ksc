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

//FileCategorizer2 Class Reference
//
//Interface for working with FileCategorizer subsystem.
//
//There are 3 types of categories: simple, autoupdate and silverimage.
//
//	- Simple category can be created by user manually.
//	- Autoupdate category is working on server side and calculating hashes of files from chosen directory.
//	- SilverImage category type accumulates hashes of files from chosen hosts.
//
//List of all members.
type FileCategorizer2 service

//AddExpressions
//Add some expressions to category.
//
//	Parameters:
//	- nCategoryId	(int64) Category id
//	- arrNewExpressions	(array) Array of new expressions. Max size is 100 elements. See Custom category format.
//	- bInclusions	(bool) If true then add to inclusions. Otherwise add to exclusions.
//
//	Return:
//	- wstrAsyncId	(string) Id of async operation.
//
//See also:
//AsyncActionStateChecker
func (fc *FileCategorizer2) AddExpressions(ctx context.Context, params interface{}) (*PxgValStr, []byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", fc.client.Server+"/api/v1.0/FileCategorizer2.AddExpressions", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := fc.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

//CancelFileMetadataOperations
//Cancel file metadata operations.
//
//Method cancels operation (GetFileMetadata, GetFilesMetadata, GetFilesMetadataFromMSI) initialized using current connection.
func (fc *FileCategorizer2) CancelFileMetadataOperations(ctx context.Context) (*PxgValInt, []byte, error) {
	request, err := http.NewRequest("POST", fc.client.Server+"/api/v1.0/FileCategorizer2.CancelFileMetadataOperations", nil)
	if err != nil {
		return nil, nil, err
	}

	pxgValInt := new(PxgValInt)
	raw, err := fc.client.Do(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

//CancelFileUpload
//Cancel file upload for file categorizer subsystem.
//
//This methode cancels file upload.
//Call FileCategorizer2.InitFileUpload to start new upload.
func (fc *FileCategorizer2) CancelFileUpload(ctx context.Context) (*PxgValInt, []byte, error) {
	request, err := http.NewRequest("POST", fc.client.Server+"/api/v1.0/FileCategorizer2.CancelFileUpload", nil)
	if err != nil {
		return nil, nil, err
	}

	pxgValInt := new(PxgValInt)
	raw, err := fc.client.Do(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

//CategoryParams struct
type CategoryParams struct {
	Category *Category `json:"pCategory,omitempty"`
}

type Category struct {
	CategoryType                int64        `json:"CategoryType"`
	CustomCategoryCipCompatible bool         `json:"CustomCategoryCipCompatible"`
	Md5WithoutSha256Exists      bool         `json:"Md5WithoutSha256Exists"`
	Exclusions                  []Exclusions `json:"exclusions"`
	FromMaster                  bool         `json:"fromMaster"`
	Inclusions                  []Inclusion  `json:"inclusions"`
	Name                        string       `json:"name"`
	Descr                       string       `json:"descr"`
	Version                     int64        `json:"version"`
}

type Exclusions struct {
	Type            string     `json:"type,omitempty"`
	ExclusionsValue *Exclusion `json:"value,omitempty"`
}

type Inclusion struct {
	Type           string     `json:"type,omitempty"`
	InclusionValue *Exclusion `json:"value,omitempty"`
}

type Exclusion struct {
	ExType      int64  `json:"ex_type,omitempty"`
	Str         string `json:"str,omitempty"`
	Str2        string `json:"str2,omitempty"`
	StrOp       int64  `json:"str_op,omitempty"`
	VerMajor    int64  `json:"ver_major,omitempty"`
	VerMinor    int64  `json:"ver_minor,omitempty"`
	VerBuild    int64  `json:"ver_build,omitempty"`
	VerRevision int64  `json:"ver_revision,omitempty"`
	VerSuffix   string `json:"ver_suffix,omitempty"`
	VerRaw      string `json:"ver_raw,omitempty"`
	VerOp       int64  `json:"ver_op,omitempty"`
	//uuid
	MediaType int64 `json:"media_type,omitempty"`
	//l_expr
	//r_expr
	//expr
	//Certificate CertificateParams `json:"certificate,omitempty"`
}

type CertificateParams struct {
	Type        string       `json:"type,omitempty"`
	Certificate *Certificate `json:"value,omitempty"`
}

type Certificate struct {
}

//CreateCategory
//Create category (simple, autoupdate or silverimage)
//
//	Parameters:
//	- pCategory	(params) Category body (see Custom category format)
//
//	Returns:
//	- (int64) Category id
//
//	Exceptions:
//	- KLSTD.STDE_OBJ_EXISTS	- name or UUID is not unique
func (fc *FileCategorizer2) CreateCategory(ctx context.Context, params CategoryParams) (*PxgValStr, []byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", fc.client.Server+"/api/v1.0/FileCategorizer2.CreateCategory", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := fc.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

//DeleteCategory
//Delete category.
//
//	Parameters:
//	- nCategoryId (int) Category id
//
//	Exceptions:
//	- KLSTD.STDE_NOTFOUND	- category not found
//	- KLSTD.STDE_NOACCESS
func (fc *FileCategorizer2) DeleteCategory(ctx context.Context, nCategoryId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nCategoryId": %d}`, nCategoryId))
	request, err := http.NewRequest("POST", fc.client.Server+"/api/v1.0/FileCategorizer2.DeleteCategory",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := fc.client.Do(ctx, request, nil)
	return raw, err
}

//ExpressionParams struct using in FileCategorizer2.DeleteExpression
type ExpressionParams struct {
	NCategoryID int64   `json:"nCategoryId,omitempty"`
	ArrIDS      []int64 `json:"arrIds"`
	BInclusions bool    `json:"bInclusions,omitempty"`
}

//DeleteExpression
//Delete some expressions from category.
//
//	Parameters:
//	- nCategoryId	(int64) Category id
//	- arrIds	(array) Array of identifiers (int). Max size is 100 elements.
//	- bInclusions	(bool) If true then delete from inclusions. Otherwise delete from exclusions.
//
//	Return:
//	- wstrAsyncId	(string) Id of async operation.
//
//See also:
//AsyncActionStateChecker
func (fc *FileCategorizer2) DeleteExpression(ctx context.Context, params ExpressionParams) (*PxgValStr, []byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", fc.client.Server+"/api/v1.0/FileCategorizer2.DeleteExpression", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := fc.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

//DoStaticAnalysisAsync
//Start static analysis.
//
//	Deprecated:
//	for using in OpenAPI. Use FileCategorizer2.DoStaticAnalysisAsync2 instead.
//
//	Parameters:
//	- wstrRequestId	(string) Async id
//	- nPolicyId	(int64) Policy id
//
//	Exceptions:
//	KLSTD::STDE_NOACCESS	- access denied
func (fc *FileCategorizer2) DoStaticAnalysisAsync(ctx context.Context, wstrRequestId string, nPolicyId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"wstrRequestId": "%s", "nPolicyId": %d}`, wstrRequestId, nPolicyId))
	request, err := http.NewRequest("POST", fc.client.Server+"/api/v1.0/FileCategorizer2.DoStaticAnalysisAsync",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := fc.client.Do(ctx, request, nil)
	return raw, err
}

//AsyncID struct
type AsyncID struct {
	WstrAsyncID string `json:"wstrAsyncId,omitempty"`
}

//DoStaticAnalysisAsync2
//Start static analysis.
//
//	Parameters:
//	- nPolicyId	(int64) Policy id
//
//	Return:
//	- wstrAsyncId	(string) Id of async operation.
//	To get status use AsyncActionStateChecker.CheckActionState, lStateCode "0" means OK.
//
//	Result of operation is placed into views:
//
//	Srvview results of static analysis - common information per category (count of executable files)
//	Srvview results of static analysis with exefiles info - list of executable files
//
//See also:
//AsyncActionStateChecker
//
//Static analysis of application categories
func (fc *FileCategorizer2) DoStaticAnalysisAsync2(ctx context.Context, nPolicyId int64) (*AsyncID, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"nPolicyId": %d}`, nPolicyId))
	request, err := http.NewRequest("POST", fc.client.Server+"/api/v1.0/FileCategorizer2.DoStaticAnalysisAsync2",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	asyncID := new(AsyncID)
	raw, err := fc.client.Do(ctx, request, &asyncID)
	return asyncID, raw, err
}

//DoTestStaticAnalysisAsync
//Start static analysis for test ACL.
//
//	Deprecated:
//	- Deprecated for using in OpenAPI. Use FileCategorizer2.DoTestStaticAnalysisAsync2 instead.
//
//	Parameters:
//	- wstrRequestId		(string) Async id
//	- nPolicyId			(int64) Policy id
//	- pTestACL			(params) Test ACL (see "ACL structure 2")
func (fc *FileCategorizer2) DoTestStaticAnalysisAsync(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", fc.client.Server+"/api/v1.0/FileCategorizer2.DoTestStaticAnalysisAsync", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := fc.client.Do(ctx, request, nil)
	return raw, err
}

//DoTestStaticAnalysisAsync2
//Start static analysis for test ACL.
//
//	Parameters:
//	- nPolicyId		(int64) Policy id
//	- pTestACL		(params) Test ACL (see "ACL structure 2")
//
//	Return:
//	- wstrAsyncId	(string) Id of async operation. To get status use AsyncActionStateChecker.
//	CheckActionState, lStateCode "0" means OK.
//Result of operation is placed into views:
//
//Srvview results of static analysis - common information per category (count of executable files)
//Srvview results of static analysis with exefiles info - list of executable files
func (fc *FileCategorizer2) DoTestStaticAnalysisAsync2(ctx context.Context, params interface{}) (*WActionGUID, []byte,
	error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", fc.client.Server+"/api/v1.0/FileCategorizer2.DoTestStaticAnalysisAsync2", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	wActionGUID := new(WActionGUID)
	raw, err := fc.client.Do(ctx, request, &wActionGUID)
	return wActionGUID, raw, err
}

//FinishStaticAnalysis
func (fc *FileCategorizer2) FinishStaticAnalysis(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", fc.client.Server+"/api/v1.0/FileCategorizer2.FinishStaticAnalysis", nil)
	if err != nil {
		return nil, err
	}

	raw, err := fc.client.Do(ctx, request, nil)
	return raw, err
}

//ForceCategoryUpdate
//Force process of automatic update (for autoupdate and silverimage)
//
//	Parameters:
//	- nCategoryId (int64) Category id
//
//	Exceptions:
//	KLSTD.STDE_NOTPERM	- wrong category type
func (fc *FileCategorizer2) ForceCategoryUpdate(ctx context.Context, nCategoryId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nCategoryId": %d}`, nCategoryId))
	request, err := http.NewRequest("POST", fc.client.Server+"/api/v1.0/FileCategorizer2.ForceCategoryUpdate",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := fc.client.Do(ctx, request, nil)
	return raw, err
}

//GetCategoriesModificationCounter
//Returns modification counter. It increments on every category change.
//
//	Returns:
//	- (int64) Modification counter
func (fc *FileCategorizer2) GetCategoriesModificationCounter(ctx context.Context) (*PxgValInt, []byte, error) {
	request, err := http.NewRequest("POST", fc.client.Server+"/api/v1.0/FileCategorizer2.GetCategoriesModificationCounter", nil)
	if err != nil {
		return nil, nil, err
	}

	pxgValInt := new(PxgValInt)
	raw, err := fc.client.Do(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

//GetCategory
//Get category by id.
//
//	Parameters:
//	 - nCategoryId	(int64) Category id
//
//	Return:
//	- pCategory	(params) Category body (see Custom category format)
//
//	Exceptions:
//	- KLSTD.STDE_NOTFOUND	- category not found
//	- KLSTD.STDE_BADFORMAT	- format of category is wrong
func (fc *FileCategorizer2) GetCategory(ctx context.Context, nCategoryId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nCategoryId": %d}`, nCategoryId))
	request, err := http.NewRequest("POST", fc.client.Server+"/api/v1.0/FileCategorizer2.GetCategory",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := fc.client.Do(ctx, request, nil)
	return raw, err
}

//GetCategoryByUUID
//Get category by uuid.
//
//	Parameters:
//	- pCategoryUUID	(string) Category uuid
//
//	Return:
//		- pCategory	(params) Category body (see Custom category format)
//
//	Exceptions:
//	- KLSTD.STDE_NOTFOUND	- category not found
func (fc *FileCategorizer2) GetCategoryByUUID(ctx context.Context, pCategoryUUID string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"pCategoryUUID": "%s"}`, pCategoryUUID))
	request, err := http.NewRequest("POST", fc.client.Server+"/api/v1.0/FileCategorizer2.GetCategoryByUUID",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := fc.client.Do(ctx, request, nil)
	return raw, err
}

//GetFileMetadata
//Get file metadata.
//
//	Parameters:
//	- ulFlag (int64) Requested meta information. Use like a mask of flags. See File metadata flags.
//
//	Return:
//	- wstrAsyncId (string) Id of async operation.
//
//	To get result use AsyncActionStateChecker.CheckActionState.
//	It returns params with requested attributes.
//	See list of attributes File metadata flags.
//
//See also:
//AsyncActionStateChecker
func (fc *FileCategorizer2) GetFileMetadata(ctx context.Context, ulFlag int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"ulFlag": %d}`, ulFlag))
	request, err := http.NewRequest("POST", fc.client.Server+"/api/v1.0/FileCategorizer2.GetFileMetadata",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := fc.client.Do(ctx, request, nil)
	return raw, err
}

//GetFilesMetadata
//Get files metadata from zip-archive.
//
//	Parameters:
//	- ulFlag (int64) Requested meta information. Use like a mask of flags. See File metadata flags.
//
//	Return:
//	- wstrAsyncId (string) Id of async operation.
//
//	To get action status use AsyncActionStateChecker.CheckActionState.
//	When action is not finished and lStateCode equals 2 then task in progress
//	and pStateData may contain attribute "Progress" (int).
//	When action is successfully finished it returns pStateData with an array "FilesMetadata".
//	Each element is a params with requested attributes. See list of attributes File metadata flags.
//
//See also:
//AsyncActionStateChecker
func (fc *FileCategorizer2) GetFilesMetadata(ctx context.Context, ulFlag int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"ulFlag": %d}`, ulFlag))
	request, err := http.NewRequest("POST", fc.client.Server+"/api/v1.0/FileCategorizer2.GetFilesMetadata",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := fc.client.Do(ctx, request, nil)
	return raw, err
}

//GetFilesMetadataFromMSI
//Get files metadata from MSI.
//
//	Parameters:
//	- ulFlag (int64) Requested meta information. Use like a mask of flags. See File metadata flags.
//
//	Return:
//	- wstrAsyncId	(string) Id of async operation.
//	To get action status use AsyncActionStateChecker.CheckActionState.
//	When action is not finished and lStateCode equals 2 then task in progress
//	and pStateData may contain attribute "Progress" (int).
//	When action is successfully finished it returns pStateData with an array "FilesMetadata".
//	Each element is a params with requested attributes. See list of attributes File metadata flags.
//
//See also:
//AsyncActionStateChecker
func (fc *FileCategorizer2) GetFilesMetadataFromMSI(ctx context.Context, ulFlag int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"ulFlag": %d}`, ulFlag))
	request, err := http.NewRequest("POST", fc.client.Server+"/api/v1.0/FileCategorizer2.GetFilesMetadataFromMSI",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := fc.client.Do(ctx, request, nil)
	return raw, err
}

//RefPolicies struct
type RefPolicies struct {
	PPolicies *PPolicies `json:"pPolicies,omitempty"`
}

type PPolicies struct {
	RefPolicies []RefPolicy `json:"RefPolicies"`
}

type RefPolicy struct {
	Type           string          `json:"type,omitempty"`
	RefPolicyValue *RefPolicyValue `json:"value,omitempty"`
}

type RefPolicyValue struct {
	FromMaster  bool   `json:"FromMaster,omitempty"`
	GroupID     int64  `json:"GroupId,omitempty"`
	GroupName   string `json:"GroupName,omitempty"`
	PolID       int64  `json:"PolId,omitempty"`
	PolName     string `json:"PolName,omitempty"`
	VServerID   int64  `json:"VServerId,omitempty"`
	VServerName string `json:"VServerName,omitempty"`
}

//GetRefPolicies
//Returns array of policies with references to specified category.
//
//	Parameters:
//	- nCatId (int64) Category id
//
//	Return:
//	- pPolicies	(params) See Policies array
//
//	Exceptions:
//	- KLSTD.STDE_NOACCESS	- access denied
func (fc *FileCategorizer2) GetRefPolicies(ctx context.Context, nCatId int64) (*RefPolicies, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"nCatId": %d}`, nCatId))
	request, err := http.NewRequest("POST", fc.client.Server+"/api/v1.0/FileCategorizer2.GetRefPolicies",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	refPolicies := new(RefPolicies)
	raw, err := fc.client.Do(ctx, request, &refPolicies)
	return refPolicies, raw, err
}

//GetSerializedCategoryBody
//Returns serialized category body for plugin.
//
//	Warning:
//	Deprecated for using in OpenAPI. Use FileCategorizer2.GetSerializedCategoryBody2 instead.
//
//	Parameters:
//	- nCategoryId (int) Category id
//
//	Return:
//	 - pCategory (params) Category serialized body
func (fc *FileCategorizer2) GetSerializedCategoryBody(ctx context.Context, nCategoryId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nCategoryId": %d}`, nCategoryId))
	request, err := http.NewRequest("POST", fc.client.Server+"/api/v1.0/FileCategorizer2.GetSerializedCategoryBody",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := fc.client.Do(ctx, request, nil)
	return raw, err
}

//GetSerializedCategoryBody2
//Returns serialized category body for plugin.
//
//	Parameters:
//	- nCategoryId (int64) Category id
//
//	Return:
//	- pCategory	(params) Category serialized body
//
//See also:
//See Serialized category format
func (fc *FileCategorizer2) GetSerializedCategoryBody2(ctx context.Context, nCategoryId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nCategoryId": %d}`, nCategoryId))
	request, err := http.NewRequest("POST", fc.client.Server+"/api/v1.0/FileCategorizer2.GetSerializedCategoryBody2",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := fc.client.Do(ctx, request, nil)
	return raw, err
}

//GetSyncId
//Returns categories synchronization id.
//
//	Returns:
//	- (int64) Synchronization id
//
//See also:
//See GroupSync
func (fc *FileCategorizer2) GetSyncId(ctx context.Context) (*PxgValInt, []byte, error) {
	request, err := http.NewRequest("POST", fc.client.Server+"/api/v1.0/FileCategorizer2.GetSyncId", nil)
	if err != nil {
		return nil, nil, err
	}

	pxgValInt := new(PxgValInt)
	raw, err := fc.client.Do(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

type UploadParams struct {
	WstrUploadURL string `json:"wstrUploadUrl,omitempty"`
}

//InitFileUpload
//Initialize file upload for file categorizer subsystem.
//
//	Return:
//	- wstrUploadUrl	(string) Upload url. See Files upload.
//
//Only one upload url is allowed for connection.
func (fc *FileCategorizer2) InitFileUpload(ctx context.Context) (*UploadParams, []byte, error) {
	request, err := http.NewRequest("POST", fc.client.Server+"/api/v1.0/FileCategorizer2.InitFileUpload", nil)
	if err != nil {
		return nil, nil, err
	}

	uploadParams := new(UploadParams)
	raw, err := fc.client.Do(ctx, request, &uploadParams)
	return uploadParams, raw, err
}

//UpdateCategory
//Update category.
//
//	Parameters:
//	- nCategoryId	(int64) Category id
//	- pCategory	(params) New category body (see Custom category format)
//
//	Exceptions:
//	- KLSTD::STDE_NOTFOUND	- category not found
//	- KLSTD::STDE_BADFORMAT	- format of category is wrong
//	- KLSTD::STDE_NOTPERM	- can't change attribute
func (fc *FileCategorizer2) UpdateCategory(ctx context.Context, params interface{}) (*PxgValStr, []byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", fc.client.Server+"/api/v1.0/FileCategorizer2.UpdateCategory", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := fc.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

//UpdateExpressions
//Update some expressions in category.
//
//	Parameters:
//	- nCategoryId	(int64) Category id
//	- arrIdAndExpression	(array) Array of params.
//	Each element contains 2 attributes: "id" (int) - expression id,	"body" (params) - expression body.
//	Max size is 100 elements. See Custom category format.
//
//	- bInclusions	(boolean) If true then update inclusions. Otherwise update exclusions.
//
//	Return:
//	wstrAsyncId	 (string) Id of async operation.
//
//See also:
//AsyncActionStateChecker
func (fc *FileCategorizer2) UpdateExpressions(ctx context.Context, params interface{}) (*PxgValStr, []byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", fc.client.Server+"/api/v1.0/FileCategorizer2.UpdateExpressions", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := fc.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}
