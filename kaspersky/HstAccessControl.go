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

// HstAccessControl Security policy Allows to specify permissions for administration groups and non-group objects.
type HstAccessControl service

// AccessCheckToAdmGroup Checks if current user session has access to the administration group.
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

// AddRole A role can be added only at a main server.
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

// DeleteRole Delete user role.
func (hac *HstAccessControl) DeleteRole(ctx context.Context, nId int64, bProtection bool) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nId":%d,"bProtection":%v}`, nId, bProtection))
	request, err := http.NewRequest("POST", hac.client.Server+"/api/v1.0/HstAccessControl.DeleteRole", bytes.NewBuffer(postData))

	if err != nil {
		return nil, err
	}

	raw, err := hac.client.Do(ctx, request, nil)
	return raw, err
}

// DeleteScObjectAcl Deletes ACL for the specified object.
func (hac *HstAccessControl) DeleteScObjectAcl(ctx context.Context, nObjId, nObjType int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nObjId":%d,"nObjType":%d}`, nObjId, nObjType))
	request, err := http.NewRequest("POST", hac.client.Server+"/api/v1.0/HstAccessControl.DeleteScObjectAcl", bytes.NewBuffer(postData))

	if err != nil {
		return nil, err
	}

	raw, err := hac.client.Do(ctx, request, nil)
	return raw, err
}

// DeleteScVServerAcl Deletes ACL for the specified virtual server.
func (hac *HstAccessControl) DeleteScVServerAcl(ctx context.Context, nId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nId":%d}`, nId))
	request, err := http.NewRequest("POST", hac.client.Server+"/api/v1.0/HstAccessControl.DeleteScVServerAcl", bytes.NewBuffer(postData))

	if err != nil {
		return nil, err
	}

	raw, err := hac.client.Do(ctx, request, nil)
	return raw, err
}

// RolesAttributes to use in the HstAccessControl.FindRoles method.
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

// FindRoles Find roles by filter string.
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

// Trustee struct
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

// FindTrustees Searches for trustees meeting specified criteria.
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

// FuncAreas struct using in HstAccessControl.GetAccessibleFuncAreas
type FuncAreas struct {
	// PFuncAreasArray array of functionality areas.
	// Each element of the array is a string: "<product>|<version>|<functional area>".
	PFuncAreasArray []string `json:"pFuncAreasArray"`
}

// GetAccessibleFuncAreas Returns accessible functional areas.
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

// GetMappingFuncAreaToPolicies Returns mapping functional area to policies.
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

// GetMappingFuncAreaToReports Returns mapping functional area to reports.
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

// GetMappingFuncAreaToSettings Returns mapping functional area to settings.
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

// GetMappingFuncAreaToTasks Returns mapping functional area to tasks.
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

// GetPolicyReadonlyNodes Returns array of paths for all nodes actually located in the specified policy section,
// which are readonly for current user session.
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

// GetRole Return parameters of a role.
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

// GetScObjectAcl Returns ACL for the specified object.
func (hac *HstAccessControl) GetScObjectAcl(ctx context.Context, nObjId, nObjType int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nObjId":%d,"nObjType":%d}`, nObjId, nObjType))
	request, err := http.NewRequest("POST", hac.client.Server+"/api/v1.0/HstAccessControl.GetScObjectAcl", bytes.NewBuffer(postData))

	if err != nil {
		return nil, err
	}

	raw, err := hac.client.Do(ctx, request, nil)
	return raw, err
}

// GetScVServerAcl Returns ACL for the server.
func (hac *HstAccessControl) GetScVServerAcl(ctx context.Context, nId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nId":%d}`, nId))
	request, err := http.NewRequest("POST", hac.client.Server+"/api/v1.0/HstAccessControl.GetScVServerAcl", bytes.NewBuffer(postData))

	if err != nil {
		return nil, err
	}

	raw, err := hac.client.Do(ctx, request, nil)
	return raw, err
}

// GetSettingsReadonlyNodes Returns array of paths for nodes from product's setting section, which are readonly for current user session.
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

// TrusteeParams struct using in HstAccessControl.GetTrustee
type TRParams struct {
	// NID id of trustee\role
	NID int64 `json:"nId,omitempty"`

	// PFieldsToReturn array of strings with attribute names
	PFieldsToReturn []string `json:"pFieldsToReturn"`
}

// Trustee struct using in HstAccessControl.GetTrustee
type Trustee struct {
	Trustee *TrusteeValue `json:"PxgRetVal,omitempty"`
}

// GetTrustee Get trustee data.
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

// GetVisualViewForAccessRights Returns descriptions of visual view for access rights in KSC.
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

// IsTaskTypeReadonly Determines read only attribute by product's task type.
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

// ModifyScObjectAcl Modify ACL for the specified object. Method updates only Accounts, permissions and roles which presented in pAclParams.
// To delete Ace from Acl, it must be added to 'delete' list.
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

// SetScObjectAcl Sets ACL for the specified object.
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

// SetScVServerAcl Set ACL for virtual server.
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

// UpdateRole Update user role.
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
