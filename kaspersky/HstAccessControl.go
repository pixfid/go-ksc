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
//	for administration groups and non-group objects.
//
//	List of all members.
type HstAccessControl service

//HstAccessControl.AccessCheckToAdmGroup
//Checks if current user session has access to the administration group.
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

//HstAccessControl.AddRole
//Add user role.
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

//HstAccessControl.DeleteRole
//Delete user role.
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

//HstAccessControl.DeleteScObjectAcl
//Deletes ACL for the specified object.
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

//HstAccessControl.DeleteScVServerAcl
//Deletes ACL for the specified virtual server.
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

//Roles attributes to use in the HstAccessControl.FindRoles method.
type RolesAttributes struct {
	RolesPChunk *RolesPChunk `json:"pChunk,omitempty"`
	PxgRetVal   *int64       `json:"PxgRetVal,omitempty"`
}

type RolesPChunk struct {
	RolesIteratorArray []RolesIteratorArray `json:"KLCSP_ITERATOR_ARRAY"`
}

type RolesIteratorArray struct {
	Type      *string    `json:"type,omitempty"`
	RoleValue *RoleValue `json:"value,omitempty"`
}

type RoleValue struct {
	KlhstACLRoleBuiltIn   *bool   `json:"KLHST_ACL_ROLE_BUILT_IN,omitempty"`
	KlhstACLRoleDN        *string `json:"KLHST_ACL_ROLE_DN,omitempty"`
	KlhstACLRoleID        *int64  `json:"KLHST_ACL_ROLE_ID,omitempty"`
	KlhstACLRoleInherited *bool   `json:"KLHST_ACL_ROLE_INHERITED,omitempty"`
	KlhstACLRoleName      *string `json:"KLHST_ACL_ROLE_NAME,omitempty"`
	KlhstACLTrusteeID     *int64  `json:"KLHST_ACL_TRUSTEE_ID,omitempty"`
	DN                    *string `json:"dn,omitempty"`
	ObjectGUID            *string `json:"objectGUID,omitempty"`
	UserPrincipalName     *string `json:"userPrincipalName,omitempty"`
}

//HstAccessControl.FindRoles
//Find roles by filter string.
//
//	Parameters:
//	- strFilter	(string) filter string, see Search filter syntax.
//	- pFieldsToReturn	(array) array of strings with role attribute names to return.
//	See Roles attributes for allowed attributes.
//	- pFieldsToOrder	(array) array of containers each of them containing two attributes:
//		|- "Name" of type (paramString), name of attribute used for sorting
//		|- "Asc" of type (paramBool), ascending if true descending otherwise
//	- lMaxLifeTime	(int64) max lifetime of accessor (sec)
//
//	Return:
//	- strAccessor	(string) result-set ID, identifier of the server-side ordered collection of found roles.
//	The result-set is destroyed and associated memory is freed in following cases:
//	Passed lMaxLifeTime seconds after last access to the result-set
//	(by methods ChunkAccessor.GetItemsCount and ChunkAccessor.GetItemsChunk
//	Session to the Administration Server has been closed.
//	ChunkAccessor.Release has been called.
//
//	Returns:
//	- (int64) number of records found
func (hac *HstAccessControl) FindRoles(ctx context.Context, params PFindParams) (*Accessor, []byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", hac.client.Server+"/api/v1.0/HstAccessControl.FindRoles", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	accessor := new(Accessor)
	raw, err := hac.client.Do(ctx, request, &accessor)
	return accessor, raw, err
}

//Trustee struct
type Trustees struct {
	TrusteePChunk *TrusteePChunk `json:"pChunk,omitempty"`
	PxgRetVal     *int64         `json:"PxgRetVal,omitempty"`
}

type TrusteePChunk struct {
	TrusteeIteratorArray []TrusteeIteratorArray `json:"KLCSP_ITERATOR_ARRAY"`
}

type TrusteeIteratorArray struct {
	Type         *string       `json:"type,omitempty"`
	TrusteeValue *TrusteeValue `json:"value,omitempty"`
}

type TrusteeValue struct {
	KlhstACLTrusteeID  *int64              `json:"KLHST_ACL_TRUSTEE_ID,omitempty"`
	KlhstACLTrusteeSid *KlhstACLTrusteeSid `json:"KLHST_ACL_TRUSTEE_SID,omitempty"`
	DN                 *string             `json:"dn,omitempty"`
	ObjectGUID         *string             `json:"objectGUID,omitempty"`
	UserPrincipalName  *string             `json:"userPrincipalName,omitempty"`
	KscInternalUserID  *int64              `json:"kscInternalUserId,omitempty"`
}

type KlhstACLTrusteeSid struct {
	Type  *string `json:"type,omitempty"`
	Value *string `json:"value,omitempty"`
}

//HstAccessControl.FindTrustees
//Searches for trustees meeting specified criteria.
//
//	Parameters:
//	- strFilter	(string) search filter, see Search filter syntax, following values may be used in the filter string:
//		+ KLHST_ACL_TRUSTEE_ID (int64) // trustee id
//		+ kscInternalUserId (int64) // Id of KSC internal user (optional)
//		+ dn (string) // account display name, as user-friendly string for account, for example,
//	"LocalSystem" or "MYCOMPUTER\sidorov" or "sidorov@avp.ru" and so on
//		+ userPrincipalName (string) // AD userPrincipalName attribute
//		+ objectGUID (paramBinary) // AD objectGUID attribute, this attribute is mandatory for non-builtin AD accounts
//	- pFieldsToReturn	(array) names of attributes to return, following attributes are possible
//		+ KLHST_ACL_TRUSTEE_ID (int64) // trustee id
//		+ KLHST_ACL_TRUSTEE_SID (paramBinary) // trustee SID
//		+ kscInternalUserId (int64) // Id of KSC internal user (optional)
//		+ dn (string) // account display name, as user-friendly string for account, for example,
//	"LocalSystem" or "MYCOMPUTER\sidorov" or "sidorov@avp.ru" and so on
//		+ userPrincipalName (string) // AD userPrincipalName attribute
//		+ objectGUID (paramBinary) // AD objectGUID attribute, this attribute is mandatory for non-builtin AD accounts
//	- pFieldsToOrder	(array) array of containers each of them containing two attributes:
//		+ "Name" of type (string), name of attribute used for sorting
//		+ "Asc" of type (paramBool), ascending if true descending otherwise
//	- lMaxLifeTime	(int) max lifetime of accessor (sec)
//
//	Return:
//	- strAccessor	(string) result-set ID, identifier of the server-side ordered collection of found roles.
//	- number of records found
//	The result-set is destroyed and associated memory is freed in following cases:
//	Passed lMaxLifeTime seconds after last access to the result-set (by methods ChunkAccessor.GetItemsCount and ChunkAccessor.GetItemsChunk
//	Session to the Administration Server has been closed.
//	ChunkAccessor.Release has been called.
func (hac *HstAccessControl) FindTrustees(ctx context.Context, params PFindParams) (*Accessor, []byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", hac.client.Server+"/api/v1.0/HstAccessControl.FindTrustees", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	accessor := new(Accessor)
	raw, err := hac.client.Do(ctx, request, &accessor)
	return accessor, raw, err
}

//FuncAreas struct using in HstAccessControl.GetAccessibleFuncAreas
type FuncAreas struct {
	//PFuncAreasArray array of functionality areas.
	//Each element of the array is a string: "<product>|<version>|<functional area>".
	//See Functional areas.
	PFuncAreasArray []string `json:"pFuncAreasArray"`
}

//HstAccessControl.GetAccessibleFuncAreas
//Returns accessible functional areas.
//
//	Parameters:
//	- lGroupId	(int64) id of the group
//	- dwAccessMask	(int64) access mask, see Access rights Access rights
//	- szwProduct	(string) product, see Functional areas
//	- szwVersion	(string) version of product, see Functional areas
//	- bInvert	(bool) if true then access mask is inverted
func (hac *HstAccessControl) GetAccessibleFuncAreas(ctx context.Context, lGroupId, dwAccessMask int64, szwProduct,
	szwVersion string, bInvert bool) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lGroupId" : %d, "dwAccessMask": %d, "szwProduct": "%s", "szwVersion": "%s", 
	"bInvert" : %v}`,
		lGroupId, dwAccessMask, szwProduct, szwVersion, bInvert))

	request, err := http.NewRequest("POST", hac.client.Server+"/api/v1.0/HstAccessControl.GetAccessibleFuncAreas",
		bytes.NewBuffer(postData))

	if err != nil {
		return nil, err
	}

	raw, err := hac.client.Do(ctx, request, nil)
	return raw, err
}

//HstAccessControl.GetMappingFuncAreaToPolicies
//Returns mapping functional area to policies.
//
//	Parameters:
//	- szwProduct	(string) product, see Functional areas
//	- szwVersion	(string) version of product, see Functional areas
//
//	Returns:
//	- (params) mapping of functional areas to settings
//        +--- (paramParams)
//            +---<functional area> (paramArray)
//            |   +---0 = (string)"\<policy name\>"
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

//HstAccessControl.GetMappingFuncAreaToReports
//Returns mapping functional area to reports.
//
//	Parameters:
//	- szwProduct	(string) product, see Functional areas
//	- szwVersion	(string) version of product, see Functional areas
//
//	Returns:
//	- (params) mapping of functional areas to reports
//        +--- (paramParams)
//            +---<functional area> (paramArray)
//            |   +---0 = (int64)<report template id>
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

//HstAccessControl.GetMappingFuncAreaToSettings
//Returns mapping functional area to settings.
//
//	Parameters:
//	- szwProduct	(string) product, see Functional areas
//	- szwVersion	(string) version of product, see Functional areas
//
//	Returns:
//	- (params) mapping of functional areas to settings
//        +--- (paramParams)
//            +---<functional area> (paramArray)
//            |   +---0 = (string)"\<setting name\>"
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

//HstAccessControl.GetMappingFuncAreaToTasks
//Returns mapping functional area to tasks.
//
//	Parameters:
//	- szwProduct	(string) product, see Functional areas
//	- szwVersion	(string) version of product, see Functional areas
//
//	Returns:
//	- (params) mapping of functional areas to tasks
//        +--- (paramParams)
//            +---<functional area> (paramArray)
//            |   +---0 = (string)"\<task name\>"
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

//HstAccessControl.GetPolicyReadonlyNodes
//Returns array of paths for all nodes actually located in the specified policy section, which are readonly for current user session.
//
//	Parameters:
//	- lGroupId			(int64) group id
//	- szwProduct		(string) name of product
//	- szwVersion		(string) product's version
//	- szwSectionName	(string) name of policy's section name
//	- pPolicySection	(params) policy's section data
//
//	Returns:
//	- (params) nodes in readonly mode
func (hac *HstAccessControl) GetPolicyReadonlyNodes(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", hac.client.Server+"/api/v1.0/HstAccessControl.GetPolicyReadonlyNodes", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hac.client.Do(ctx, request, nil)
	return raw, err
}

//HstAccessControl.GetRole
//Return parameters of a role.
//
//	Parameters:
//	- nId	(int) id of a role
//	- pFieldsToReturn	(array) array of strings with attribute names, see Role. May contain attributes:
//		|- KLHST_ACL_ROLE_DN
//		|- KLHST_ACL_ROLE_NAME
//		|- KLHST_ACL_ROLE_ID
//		|- KLHST_ACL_ROLE_BUILT_IN
//		|- KLHST_ACL_ROLE_INHERITED
//		|- KLHST_ACL_ROLE_READ_ONLY
//		|- role_products
//
//	Returns:
//	- (params) requested parameters
func (hac *HstAccessControl) GetRole(ctx context.Context, params TRParams) (*Trustee, []byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", hac.client.Server+"/api/v1.0/HstAccessControl.GetRole", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	trustee := new(Trustee)
	raw, err := hac.client.Do(ctx, request, &trustee)
	return trustee, raw, err
}

//HstAccessControl.GetScObjectAcl
//Returns ACL for the specified object.
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

//HstAccessControl.GetScVServerAcl
//Returns ACL for the server.
//
//	Parameters:
//	- nId			(int) -1 means 'current server', otherwise virtual server id
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

//HstAccessControl.GetSettingsReadonlyNodes
//Returns array of paths for nodes from product's setting section, which are readonly for current user session.
//
//	Parameters:
//	- lGroupId			(int64) group id
//	- szwProduct		(string) name of product
//	- szwVersion		(string) product's version
//	- szwSectionName	(string) name of setting's section name
//	- pSettingsSection	(params) setting's section data
//
//	Returns:
//	- (params) nodes in readonly mode
func (hac *HstAccessControl) GetSettingsReadonlyNodes(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", hac.client.Server+"/api/v1.0/HstAccessControl.GetSettingsReadonlyNodes", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hac.client.Do(ctx, request, nil)
	return raw, err
}

//TrusteeParams struct using in HstAccessControl.GetTrustee
type TRParams struct {
	//id of trustee\role
	NID int64 `json:"nId,omitempty"`

	//(array) array of strings with attribute names
	PFieldsToReturn []string `json:"pFieldsToReturn"`
}

//Trustee struct using in HstAccessControl.GetTrustee
type Trustee struct {
	Trustee *TrusteeValue `json:"PxgRetVal,omitempty"`
}

//HstAccessControl.GetTrustee
//Get trustee data.
//
//Returns trustee's data.
//
//	Parameters:
//	- nId	(int64) trustee id
//	- pFieldsToReturn	(array) array of strings with attribute names, following attrs are possible:
//	+ KLHST_ACL_TRUSTEE_ID (int64) // trustee id
//	+ KLHST_ACL_TRUSTEE_SID (binary) // trustee SID
//	+ id (string) // account id
//	+ dn (string) // account display name, as user-friendly string for account, for example,
//	"LocalSystem" or "MYCOMPUTER\sidorov" or "sidorov@avp.ru" and so on
//	+ userPrincipalName (string) // AD userPrincipalName attribute
//	+ objectGUID (binary) // AD objectGUID attribute, this attribute is mandatory for non-builtin AD accounts
//
//	Returns:
//	Trustee (params) trustee's data
func (hac *HstAccessControl) GetTrustee(ctx context.Context, params TRParams) (*Trustee, []byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", hac.client.Server+"/api/v1.0/HstAccessControl.GetTrustee", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	trustee := new(Trustee)
	raw, err := hac.client.Do(ctx, request, &trustee)
	return trustee, raw, err
}

//HstAccessControl.GetVisualViewForAccessRights
//Returns descriptions of visual view for access rights in KSC.
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
//                        +--KLCONN_FUNC_AREA_DISP_NAME = (string)"<display name>"
//                        +---<functional area> (paramParams) // see Functional areas
//                            +---KLCONN_FUNC_AREA_DISP_NAME = (string)"<display name>"
//                            +---KLCONN_ACE_OPERATION_MASK = (int64)<access mask> see Access rights
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

//HstAccessControl.IsTaskTypeReadonly
//Determines read only attribute by product's task type.
//
//	Parameters:
//	- lGroupId			(int64) group id
//	- szwProduct		(string) name of product
//	- szwVersion		(string) product's version
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

//HstAccessControl.ModifyScObjectAcl
//Modify ACL for the specified object.
//
//Method updates only Accounts, permissions and roles which presented in pAclParams. To delete Ace from Acl, it must be added to 'delete' list.
//
//	Parameters:
//	- nObjId				(int64) object id
//	- nObjType				(int64) object type, see Object types
//	- pAclParams			(params) ACL, see ACL structure 2
//	- bCheckCurrentUserAce	(bool) if true then it checks that the user does not reduce rights for himself
func (hac *HstAccessControl) ModifyScObjectAcl(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", hac.client.Server+"/api/v1.0/HstAccessControl.ModifyScObjectAcl", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hac.client.Do(ctx, request, nil)
	return raw, err
}

//HstAccessControl.SetScObjectAcl
//Sets ACL for the specified object.
//
//	Parameters:
//	- nObjId				(int64) object id
//	- nObjType				(int64) object type, see Object types
//	- pAclParams			(params) ACL, see ACL structure 2
//	- bCheckCurrentUserAce	(bool) if true then it checks that the user does not reduce rights for himself
func (hac *HstAccessControl) SetScObjectAcl(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", hac.client.Server+"/api/v1.0/HstAccessControl.SetScObjectAcl", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hac.client.Do(ctx, request, nil)
	return raw, err
}

//HstAccessControl.SetScVServerAcl
//Set ACL for virtual server.
//
//	Parameters:
//	- nId					(int64) -1 means 'current server', otherwise virtual server id
//	- pAclParams			(params) ACL, see ACL structure 2
//	- bCheckCurrentUserAce	(bool) if true then it checks that the user does not reduce rights for himself
func (hac *HstAccessControl) SetScVServerAcl(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", hac.client.Server+"/api/v1.0/HstAccessControl.SetScVServerAcl", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hac.client.Do(ctx, request, nil)
	return raw, err
}

//HstAccessControl.UpdateRole
//Update user role.
//
//	Parameters:
//	- nId			(int64) id of a role
//	- pRoleData		(params) parameters of a role, KLHST_ACL_ROLE_NAME is mandatory, see Role
//	- bProtection	(bool) if true then it checks that the user does not reduce rights for himself
//
//	Returns:
//	- (params) returned parameters of role (guid, id, display name).
func (hac *HstAccessControl) UpdateRole(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", hac.client.Server+"/api/v1.0/HstAccessControl.UpdateRole", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hac.client.Do(ctx, request, nil)
	return raw, err
}
