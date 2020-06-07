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

//	HstAccessControl Class Reference
//
//	Security policy Allows to specify permissions
//	for administrration groups and non-group objects.
//
//	List of all members.
type HstAccessControl service

//	Checks if current user session has access to the administration group.
//
//	Parameters:
//	- lGroupId	(int64) id of the group
//	- dwAccessMask	(int64) access mask, see Access rights Access rights
//	- szwFuncArea	(string) functional area, see Functional areas
//	- szwProduct	(string) product, see Functional areas
//	- szwVersion	(string) version of product, see Functional areas
//
//	Returns:
//	- (bool) true if the user has access else false
func (hac *HstAccessControl) AccessCheckToAdmGroup(ctx context.Context,
	lGroupId, dwAccessMask int64, szwFuncArea, szwProduct, szwVersion string) (*PxgValBool, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"lGroupId":%d,"dwAccessMask":%d,"szwFuncArea":"%s","szwProduct":"%s", 
	"szwVersion":"%s"}`, lGroupId, dwAccessMask, szwFuncArea, szwProduct, szwVersion))
	request, err := http.NewRequest("POST", hac.client.Server+"/api/v1.0/HstAccessControl.AccessCheckToAdmGroup", bytes.NewBuffer(postData))

	if err != nil {
		return nil, nil, err
	}

	pxgValBool := new(PxgValBool)
	raw, err := hac.client.Do(ctx, request, &pxgValBool)
	return pxgValBool, raw, err
}

//	Add user role.
//
//	A role can be added only at a main server.
//
//	Parameters:
//	- pRoleData	(params) container with role attributes, see Role. May contain attributes:
//	- KLHST_ACL_ROLE_DN (mandatory)
//	- role_products (mandatory)
//
//	Returns:
//	- (params) container params with atttibutes:
//	+--- (PARAMS_T)
//  	+---KLHST_ACL_ROLE_DN = (STRING_T)\<display name\>
//     	+---KLHST_ACL_ROLE_ID = (INT_T)\<id\>
//     	+---KLHST_ACL_ROLE_NAME = (STRING_T)<guid name>
func (hac *HstAccessControl) AddRole(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", hac.client.Server+"/api/v1.0/HstAccessControl.AddRole", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hac.client.Do(ctx, request, nil)
	return raw, err
}

//	Delete user role.
//
//	Parameters:
//	- nId	(int64) id of a role
//	- bProtection	(bool) if true then it checks that the user does not reduce rights for himself
func (hac *HstAccessControl) DeleteRole(ctx context.Context, nId int64, bProtection bool) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nId":%d,"bProtection":%v}`, nId, bProtection))
	request, err := http.NewRequest("POST", hac.client.Server+"/api/v1.0/HstAccessControl.DeleteRole", bytes.NewBuffer(postData))

	if err != nil {
		return nil, err
	}

	raw, err := hac.client.Do(ctx, request, nil)
	return raw, err
}

//	Deletes ACL for the specified object.
//
//	Parameters:
//	- nObjId	(int64) object id
//	- nObjType	(int64) object type, see Object types
func (hac *HstAccessControl) DeleteScObjectAcl(ctx context.Context, nObjId, nObjType int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nObjId":%d,"nObjType":%d}`, nObjId, nObjType))
	request, err := http.NewRequest("POST", hac.client.Server+"/api/v1.0/HstAccessControl.DeleteScObjectAcl", bytes.NewBuffer(postData))

	if err != nil {
		return nil, err
	}

	raw, err := hac.client.Do(ctx, request, nil)
	return raw, err
}

//	Deletes ACL for the specified virtual server.
//
//	Parameters:
//	- nId	(int64) server id
func (hac *HstAccessControl) DeleteScVServerAcl(ctx context.Context, nId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nId":%d}`, nId))
	request, err := http.NewRequest("POST", hac.client.Server+"/api/v1.0/HstAccessControl.DeleteScVServerAcl", bytes.NewBuffer(postData))

	if err != nil {
		return nil, err
	}

	raw, err := hac.client.Do(ctx, request, nil)
	return raw, err
}

//TODO FindRoles
//TODO FindTrustees
//TODO GetAccessibleFuncAreas

//	Returns mapping functional area to policies.
//
//	Parameters:
//	- szwProduct	(string) product, see Functional areas
//	- szwVersion	(string) version of product, see Functional areas
//
//	Returns:
//	- (params) mapping of functional areas to settings
//        +--- (paramParams)
//            +---<functional area> (paramArray)
//            |   +---0 = (paramString)"\<policy name\>"
func (hac *HstAccessControl) GetMappingFuncAreaToPolicies(ctx context.Context, szwProduct, szwVersion string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"szwProduct": "%s","szwVersion": "%s"}`, szwProduct, szwVersion))
	request, err := http.NewRequest("POST", hac.client.Server+"/api/v1.0/HstAccessControl.GetMappingFuncAreaToPolicies",
		bytes.NewBuffer(postData))

	if err != nil {
		return nil, err
	}

	raw, err := hac.client.Do(ctx, request, nil)
	return raw, err
}

//	Returns mapping functional area to reports.
//
//	Parameters:
//	- szwProduct	(string) product, see Functional areas
//	- szwVersion	(string) version of product, see Functional areas
//
//	Returns:
//	- (params) mapping of functional areas to reports
//        +--- (paramParams)
//            +---<functional area> (paramArray)
//            |   +---0 = (paramInt)<report template id>
//
func (hac *HstAccessControl) GetMappingFuncAreaToReports(ctx context.Context, szwProduct, szwVersion string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"szwProduct": "%s","szwVersion": "%s"}`, szwProduct, szwVersion))
	request, err := http.NewRequest("POST", hac.client.Server+"/api/v1.0/HstAccessControl.GetMappingFuncAreaToReports",
		bytes.NewBuffer(postData))

	if err != nil {
		return nil, err
	}

	raw, err := hac.client.Do(ctx, request, nil)
	return raw, err
}

//	Returns mapping functional area to settings.
//
//	Parameters:
//	- szwProduct	(string) product, see Functional areas
//	- szwVersion	(string) version of product, see Functional areas
//
//	Returns:
//	- (params) mapping of functional areas to settings
//        +--- (paramParams)
//            +---<functional area> (paramArray)
//            |   +---0 = (paramString)"\<setting name\>"
//
func (hac *HstAccessControl) GetMappingFuncAreaToSettings(ctx context.Context, szwProduct, szwVersion string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"szwProduct": "%s","szwVersion": "%s"}`, szwProduct, szwVersion))
	request, err := http.NewRequest("POST", hac.client.Server+"/api/v1.0/HstAccessControl.GetMappingFuncAreaToSettings",
		bytes.NewBuffer(postData))

	if err != nil {
		return nil, err
	}

	raw, err := hac.client.Do(ctx, request, nil)
	return raw, err
}

//	Returns mapping functional area to tasks.
//
//	Parameters:
//	- szwProduct	(string) product, see Functional areas
//	- szwVersion	(string) version of product, see Functional areas
//
//	Returns:
//	- (params) mapping of functional areas to tasks
//        +--- (paramParams)
//            +---<functional area> (paramArray)
//            |   +---0 = (paramString)"\<task name\>"
//
func (hac *HstAccessControl) GetMappingFuncAreaToTasks(ctx context.Context, szwProduct, szwVersion string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"szwProduct": "%s","szwVersion": "%s"}`, szwProduct, szwVersion))
	request, err := http.NewRequest("POST", hac.client.Server+"/api/v1.0/HstAccessControl.GetMappingFuncAreaToTasks",
		bytes.NewBuffer(postData))

	if err != nil {
		return nil, err
	}

	raw, err := hac.client.Do(ctx, request, nil)
	return raw, err
}

//TODO GetPolicyReadonlyNodes
//TODO GetRole

//	Returns ACL for the specified object.
//
//	Parameters:
//	- nObjId	(int64) object id
//	- nObjType	(int64) object type, see Object types
//	Return:
//	- pAclParams	(params) see ACL structure 2
func (hac *HstAccessControl) GetScObjectAcl(ctx context.Context, nObjId, nObjType int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nObjId":%d,"nObjType":%d}`, nObjId, nObjType))
	request, err := http.NewRequest("POST", hac.client.Server+"/api/v1.0/HstAccessControl.GetScObjectAcl", bytes.NewBuffer(postData))

	if err != nil {
		return nil, err
	}

	raw, err := hac.client.Do(ctx, request, nil)
	return raw, err
}

//	Returns ACL for the server.
//
//	Parameters:
//	- nId	(int) -1 means 'current server', otherwise virtual server id
//
//	Return:
//	- pAclParams	(params) see ACL structure 2
func (hac *HstAccessControl) GetScVServerAcl(ctx context.Context, nId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nId":%d}`, nId))
	request, err := http.NewRequest("POST", hac.client.Server+"/api/v1.0/HstAccessControl.GetScVServerAcl", bytes.NewBuffer(postData))

	if err != nil {
		return nil, err
	}

	raw, err := hac.client.Do(ctx, request, nil)
	return raw, err
}

//TODO GetSettingsReadonlyNodes
//TODO GetTrustee

//	Returns descriptions of visual view for access rights in KSC.
//
//	Parameters:
//	- wstrLangCode	(string) IETF language tag (for example: en-us)
//	- nObjId	(int64) object id (for administration groups is ignored).
//	- nObjType	(int64) object type, see Object types
//	Return:
//	- pViewParams	(params) see ACL view structure:
//
//	"ACL view structure"
//    +---(paramParams)
//                +--- <guid>(paramParams)
//                    +---KLCONN_ACE_PRODUCT = "<product>" // see Functional areas
//                    +---KLCONN_ACE_VERSION = "<version>"
//                    +---<functional area> (paramParams) // group of functional areas
//                        +--KLCONN_FUNC_AREA_DISP_NAME = (paramString)"<display name>"
//                        +---<functional area> (paramParams) // see Functional areas
//                            +---KLCONN_FUNC_AREA_DISP_NAME = (paramString)"<display name>"
//                            +---KLCONN_ACE_OPERATION_MASK = (paramInt)<access mask> see Access rights
func (hac *HstAccessControl) GetVisualViewForAccessRights(ctx context.Context, wstrLangCode string, nObjId, nObjType int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"wstrLangCode": "%s","nObjId": %d,"nObjType": %d}`, wstrLangCode, nObjId, nObjType))
	request, err := http.NewRequest("POST", hac.client.Server+"/api/v1.0/HstAccessControl.GetVisualViewForAccessRights",
		bytes.NewBuffer(postData))

	if err != nil {
		return nil, err
	}

	raw, err := hac.client.Do(ctx, request, nil)
	return raw, err
}

//	Determines read only attribute by product's task type.
//
//	Parameters:
//	- lGroupId	(int64) group id
//	- szwProduct	(string) name of product
//	- szwVersion	(string) product's version
//	- szwTaskTypeName	(string) name of product's task type
//
//	Returns:
//	- (bool) true if task is readonly
func (hac *HstAccessControl) IsTaskTypeReadonly(ctx context.Context, lGroupId int64, szwProduct, szwVersion,
	szwTaskTypeName string) (*PxgValBool, []byte,
	error) {
	postData := []byte(fmt.Sprintf(`{"lGroupId": %d,"szwProduct": "%s","szwVersion": "%s","szwTaskTypeName": "%s"}`,
		lGroupId, szwProduct, szwVersion, szwTaskTypeName))
	request, err := http.NewRequest("POST", hac.client.Server+"/api/v1.0/HstAccessControl.IsTaskTypeReadonly",
		bytes.NewBuffer(postData))

	if err != nil {
		return nil, nil, err
	}

	pxgValBool := new(PxgValBool)
	raw, err := hac.client.Do(ctx, request, &pxgValBool)
	return pxgValBool, raw, err
}

//TODO ModifyScObjectAcl
//TODO SetScObjectAcl
//TODO SetScVServerAcl
//TODO UpdateRole
