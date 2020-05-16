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

//	HstAccessControl Class Reference
//	Security policy Allows to specify permissions
//	for administrration groups and non-group objects.
//
//	List of all members.
type HstAccessControl service

//TODO AccessCheckToAdmGroup
//TODO AddRole
//TODO DeleteRole
//TODO DeleteScObjectAcl
//TODO DeleteScVServerAcl
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
func (ce *HstAccessControl) GetMappingFuncAreaToPolicies(ctx context.Context, szwProduct, szwVersion string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"szwProduct": "%s","szwVersion": "%s"}`, szwProduct, szwVersion))
	request, err := http.NewRequest("POST", ce.client.Server+"/api/v1.0/HstAccessControl.GetMappingFuncAreaToPolicies",
		bytes.NewBuffer(postData))

	if err != nil {
		return nil, err
	}

	raw, err := ce.client.Do(ctx, request, nil)
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
func (ce *HstAccessControl) GetMappingFuncAreaToReports(ctx context.Context, szwProduct, szwVersion string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"szwProduct": "%s","szwVersion": "%s"}`, szwProduct, szwVersion))
	request, err := http.NewRequest("POST", ce.client.Server+"/api/v1.0/HstAccessControl.GetMappingFuncAreaToReports",
		bytes.NewBuffer(postData))

	if err != nil {
		return nil, err
	}

	raw, err := ce.client.Do(ctx, request, nil)
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
func (ce *HstAccessControl) GetMappingFuncAreaToSettings(ctx context.Context, szwProduct, szwVersion string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"szwProduct": "%s","szwVersion": "%s"}`, szwProduct, szwVersion))
	request, err := http.NewRequest("POST", ce.client.Server+"/api/v1.0/HstAccessControl.GetMappingFuncAreaToSettings",
		bytes.NewBuffer(postData))

	if err != nil {
		return nil, err
	}

	raw, err := ce.client.Do(ctx, request, nil)
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
func (ce *HstAccessControl) GetMappingFuncAreaToTasks(ctx context.Context, szwProduct, szwVersion string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"szwProduct": "%s","szwVersion": "%s"}`, szwProduct, szwVersion))
	request, err := http.NewRequest("POST", ce.client.Server+"/api/v1.0/HstAccessControl.GetMappingFuncAreaToTasks",
		bytes.NewBuffer(postData))

	if err != nil {
		return nil, err
	}

	raw, err := ce.client.Do(ctx, request, nil)
	return raw, err
}

//TODO GetPolicyReadonlyNodes
//TODO GetRole
//TODO GetScObjectAcl
//TODO GetScVServerAcl
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
func (ce *HstAccessControl) GetVisualViewForAccessRights(ctx context.Context, wstrLangCode string, nObjId,
	nObjType int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"wstrLangCode": "%s","nObjId": %d,"nObjType": %d}`, wstrLangCode, nObjId, nObjType))
	request, err := http.NewRequest("POST", ce.client.Server+"/api/v1.0/HstAccessControl.GetVisualViewForAccessRights",
		bytes.NewBuffer(postData))

	if err != nil {
		return nil, err
	}

	raw, err := ce.client.Do(ctx, request, nil)
	return raw, err
}

//TODO IsTaskTypeReadonly
//TODO ModifyScObjectAcl
//TODO SetScObjectAcl
//TODO SetScVServerAcl
//TODO UpdateRole
