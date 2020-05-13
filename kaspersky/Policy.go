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

	"net/http"
)

//	Policy Class Reference
//
//	Policies processing. More...
//
//	Allows to manage policies, change policies states and acquire policies data
//
//	List of all members.
type Policy service

type PxgValPolicy struct {
	PxgRetValPolicy PxgRetValPolicy `json:"PxgRetVal"`
}

type PxgRetValPolicy struct {
	KlpolAcceptParent          bool   `json:"KLPOL_ACCEPT_PARENT"`
	KlpolActive                bool   `json:"KLPOL_ACTIVE"`
	KlpolCreated               Klpol  `json:"KLPOL_CREATED"`
	KlpolDN                    string `json:"KLPOL_DN"`
	KlpolForced                bool   `json:"KLPOL_FORCED"`
	KlpolForceDistrib2Children bool   `json:"KLPOL_FORCE_DISTRIB2CHILDREN"`
	KlpolGroupID               int64  `json:"KLPOL_GROUP_ID"`
	KlpolGroupName             string `json:"KLPOL_GROUP_NAME"`
	KlpolGsynID                int64  `json:"KLPOL_GSYN_ID"`
	KlpolHideOnSlaves          bool   `json:"KLPOL_HIDE_ON_SLAVES"`
	KlpolID                    int64  `json:"KLPOL_ID"`
	KlpolInherited             bool   `json:"KLPOL_INHERITED"`
	KlpolModified              Klpol  `json:"KLPOL_MODIFIED"`
	KlpolProduct               string `json:"KLPOL_PRODUCT"`
	KlpolRoaming               bool   `json:"KLPOL_ROAMING"`
	KlpolVersion               string `json:"KLPOL_VERSION"`
}

type Klpol struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

//TODO AddPolicy
//TODO CopyOrMovePolicy

//	Delete policy.
//
//	Makes the the specified policy inactive, and then deletes it.
//
//	Parameters:
//	- nPolicy	(int64) identifier of the policy to delete
func (pl *Policy) DeletePolicy(ctx context.Context, nPolicy int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nPolicy": %d}`, nPolicy))
	request, err := http.NewRequest("POST", pl.client.Server+"/api/v1.0/Policy.DeletePolicy", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pl.client.Do(ctx, request, nil)
	return raw, err
}

//	Obtain policies that affect the specified group.
//
//	Returns active and roaming policies that affect specified group
//
//	Parameters:
//	- nGroupId	(int64) group id
//
//	Returns:
//	- (array) array, each element is paramParams which contains policy attributes
//
//	See also:
//	List of policy attributes
func (pl *Policy) GetEffectivePoliciesForGroup(ctx context.Context, nGroupId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nGroupId": %d}`, nGroupId))
	request, err := http.NewRequest("POST", pl.client.Server+"/api/v1.0/Policy.GetEffectivePoliciesForGroup", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pl.client.Do(ctx, request, nil)
	return raw, err
}

//
//	Acquire array of outbreak policies.
//
//	Returns the array of outbreak policies
//
//	Returns:
//	- (params) Container with attribute KLPOL_POL_OUTBREAK of type paramArray,
//	each entry contains two attributes KLPOL_ID and KLPOL_OUTBREAK_MASK
//
//	See also:
//	Policy outbreak attributes
func (pl *Policy) GetOutbreakPolicies(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", pl.client.Server+"/api/v1.0/Policy.GetOutbreakPolicies", nil)
	if err != nil {
		return nil, err
	}

	raw, err := pl.client.Do(ctx, request, nil)
	return raw, err
}

//	Obtain policies for specified group.
//
//	Returns policies located in specified group.
//
//	Parameters:
//	- nGroupId	(int64) value -1 means "all groups"
//
//	Returns:
//	- (array) array, each element is paramParams which contains policy attributes
//
//	See also:
//	List of policy attributes
func (pl *Policy) GetPoliciesForGroup(ctx context.Context, nGroupId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nGroupId": %d}`, nGroupId))
	request, err := http.NewRequest("POST", pl.client.Server+"/api/v1.0/Policy.GetPoliciesForGroup", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pl.client.Do(ctx, request, nil)
	return raw, err
}

//	Acquire policy contents.
//
//	Opens settings storage SsContents of the specified policy. The settings storage contains both predefined and product-specific sections.
//
//	Parameters:
//	- nPolicy	(int64) policy id
//	- nRevisionId	(int64) policy revision id, 0 means 'current policy'
//	- nLifeTime	(int64) timeout in milliseconds to keep this SsContents object alive, zero means 'default value'
//
//	Returns:
//	- (string) identifier of opened SsContents, must be closed with SsContents.SS_Release
func (pl *Policy) GetPolicyContents(ctx context.Context, nPolicy, nRevisionId, nLifeTime int64) (*PxgValStr, []byte,
	error) {
	postData := []byte(fmt.Sprintf(`{ "nPolicy": %d , "nRevisionId": %d , "nLifeTime": %d }`, nPolicy, nRevisionId,
		nLifeTime))
	request, err := http.NewRequest("POST", pl.client.Server+"/api/v1.0/Policy.GetPolicyContents", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := pl.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

//	Obtain policy data.
//
//	Returns data for specified policy
//
//	Parameters:
//	- nPolicy	(int64) policy id
//
//	Returns:
//	- (params) container with policy attributes ( See List of policy attributes)
func (pl *Policy) GetPolicyData(ctx context.Context, nPolicy int64) (*PxgValPolicy, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"nPolicy": %d}`, nPolicy))
	request, err := http.NewRequest("POST", pl.client.Server+"/api/v1.0/Policy.GetPolicyData", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValPolicy := new(PxgValPolicy)
	raw, err := pl.client.Do(ctx, request, &pxgValPolicy)
	return pxgValPolicy, raw, err
}

//	Make policy active or inactive.
//
//	Make the specified policy active or inactive
//
//	Parameters:
//	- nPolicy	(int64) policy id to activate
//	- bActive	(bool) true to make the specified inactive or roaming policy active (
//	returns false if the policy is already active),
//	false to make the specified active or roaming policy inactive
//	(returns false if the policy is already inactive)
//
//	Returns:
//	- ( bool) true if policy was successfully set to active or false otherwise
//
//	See also:
//	List of policy attributes
func (pl *Policy) MakePolicyActive(ctx context.Context, nPolicy int64, bActive bool) (*PxgValPolicy, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"nPolicy": %d, "bActive": %v}`, nPolicy, bActive))
	request, err := http.NewRequest("POST", pl.client.Server+"/api/v1.0/Policy.MakePolicyActive", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValPolicy := new(PxgValPolicy)
	raw, err := pl.client.Do(ctx, request, &pxgValPolicy)
	return pxgValPolicy, raw, err
}

//	Make policy roaming.
//
//	Make the specified policy roaming
//
//	Parameters:
//	- nPolicy	(int64) active or inactive policy id
//
//	Returns:
//	- (bool) false if the policy is already roaming
//
//	See also:
//	KLPOL_ROAMING
//	List of policy attributes
func (pl *Policy) MakePolicyRoaming(ctx context.Context, nPolicy int64) (*PxgValPolicy, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"nPolicy": %d}`, nPolicy))
	request, err := http.NewRequest("POST", pl.client.Server+"/api/v1.0/Policy.MakePolicyRoaming", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValPolicy := new(PxgValPolicy)
	raw, err := pl.client.Do(ctx, request, &pxgValPolicy)
	return pxgValPolicy, raw, err
}

//	Revert policy to its older version.
//
//	Replaces the specified policy nPolicy by its revision (older version) nRevisionId
//
//	Parameters:
//	- nPolicy	(int64) id of policy to revert
//	- nRevisionId	(int64) id of policy revision
func (pl *Policy) RevertPolicyToRevision(ctx context.Context, nPolicy, nRevisionId int64) (*PxgValPolicy, []byte,
	error) {
	postData := []byte(fmt.Sprintf(`{"nPolicy": %d, "nRevisionId": %d}`, nPolicy, nRevisionId))
	request, err := http.NewRequest("POST", pl.client.Server+"/api/v1.0/Policy.RevertPolicyToRevision", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValPolicy := new(PxgValPolicy)
	raw, err := pl.client.Do(ctx, request, &pxgValPolicy)
	return pxgValPolicy, raw, err
}

//TODO SetOutbreakPolicies
//TODO UpdatePolicyData

//	Export policy to a blob.
//
//	Exports policy into single chunk.
//
//	Parameters:
//	- lPolicy	(int64) policy id
//
//	Returns:
//	- blob with exported policy
//
//	See also:
//	Policy.ImportPolicy
func (pl *Policy) ExportPolicy(ctx context.Context, lPolicy int64) (*PxgValStr, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"lPolicy": %d}`, lPolicy))
	request, err := http.NewRequest("POST", pl.client.Server+"/api/v1.0/Policy.ExportPolicy", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := pl.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

//PolicyBlob struct
type PolicyBlob struct {
	LGroup *int64 `json:"lGroup,omitempty"`
	//	PData raw []byte data of policy base64 encoded
	PData *string `json:"pData,omitempty"`
}

//	Import policy from blob.
//
//	Exports policy from a single chunk.
//
//	Parameters:
//	- params	(PolicyBlob)
//
//	Returns:
//	- policy id
//
//	See also:
//	- Policy.ExportPolicy
func (pl *Policy) ImportPolicy(ctx context.Context, params PolicyBlob) (*PxgValStr, []byte, error) {
	postData, err := json.Marshal(&params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", pl.client.Server+"/api/v1.0/Policy.ImportPolicy", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := pl.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}
