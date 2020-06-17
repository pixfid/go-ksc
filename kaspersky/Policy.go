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

//	Policy Class Reference
//
//	Policies processing..
//
//	Allows to manage policies, change policies states and acquire policies data
//
//	List of all members.
type Policy service

//PxgValPolicy struct
type PxgValPolicy struct {
	PxgRetValPolicy PxgRetValPolicy `json:"PxgRetVal"`
}

type PxgRetValPolicy struct {
	KlpolAcceptParent          bool     `json:"KLPOL_ACCEPT_PARENT"`
	KlpolActive                bool     `json:"KLPOL_ACTIVE"`
	KlpolCreated               DateTime `json:"KLPOL_CREATED"`
	KlpolDN                    string   `json:"KLPOL_DN"`
	KlpolForced                bool     `json:"KLPOL_FORCED"`
	KlpolForceDistrib2Children bool     `json:"KLPOL_FORCE_DISTRIB2CHILDREN"`
	KlpolGroupID               int64    `json:"KLPOL_GROUP_ID"`
	KlpolGroupName             string   `json:"KLPOL_GROUP_NAME"`
	KlpolGsynID                int64    `json:"KLPOL_GSYN_ID"`
	KlpolHideOnSlaves          bool     `json:"KLPOL_HIDE_ON_SLAVES"`
	KlpolID                    int64    `json:"KLPOL_ID"`
	KlpolInherited             bool     `json:"KLPOL_INHERITED"`
	KlpolModified              DateTime `json:"KLPOL_MODIFIED"`
	KlpolProduct               string   `json:"KLPOL_PRODUCT"`
	KlpolRoaming               bool     `json:"KLPOL_ROAMING"`
	KlpolVersion               string   `json:"KLPOL_VERSION"`
}

// NewPolicy new policy with the specified attributes
type NewPolicy struct {
	PPolicyData PPolicyData `json:"pPolicyData,omitempty"`
}

type PPolicyData struct {
	KlpolDN                    string `json:"KLPOL_DN"`
	KlpolProduct               string `json:"KLPOL_PRODUCT"`
	KlpolVersion               string `json:"KLPOL_VERSION"`
	KlpolGroupID               int64  `json:"KLPOL_GROUP_ID"`
	KlpolAcceptParent          bool   `json:"KLPOL_ACCEPT_PARENT,omitempty"`
	KlpolActive                bool   `json:"KLPOL_ACTIVE,omitempty"`
	KlpolForceDistrib2Children bool   `json:"KLPOL_FORCE_DISTRIB2CHILDREN,omitempty"`
	KlpolHideOnSlaves          bool   `json:"KLPOL_HIDE_ON_SLAVES,omitempty"`
	KlpolRoaming               bool   `json:"KLPOL_ROAMING,omitempty"`
}

// AddPolicy Create new policy.
// Creates a new policy with the specified attributes
func (pl *Policy) AddPolicy(ctx context.Context, params NewPolicy) (*PxgValInt, error) {
	postData, err := json.Marshal(&params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", pl.client.Server+"/api/v1.0/Policy.AddPolicy", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	pxgValInt := new(PxgValInt)
	_, err = pl.client.Do(ctx, request, &pxgValInt)
	return pxgValInt, err
}

// MovePolicyParams struct
type MovePolicyParams struct {
	NPolicy       int64           `json:"nPolicy"`
	NNewGroupID   int64           `json:"nNewGroupId"`
	BDeleteOrigin bool            `json:"bDeleteOrigin"`
	PExtraData    PExtraDataValue `json:"pExtraData"`
}

type PExtraDataClass struct {
	Type            string          `json:"type"`
	PExtraDataValue PExtraDataValue `json:"value"`
}

type PExtraDataValue struct {
	KlpolDN     string `json:"KLPOL_DN"`
	KlpolActive bool   `json:"KLPOL_ACTIVE"`
}

// CopyOrMovePolicy Copy or move policy.
// Copies the specified policy and optionally deletes it. Returns id of the new policy
func (pl *Policy) CopyOrMovePolicy(ctx context.Context, params MovePolicyParams) (*PxgValInt, error) {
	postData, err := json.Marshal(&params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", pl.client.Server+"/api/v1.0/Policy.CopyOrMovePolicy", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	pxgValInt := new(PxgValInt)
	_, err = pl.client.Do(ctx, request, &pxgValInt)
	return pxgValInt, err
}

// DeletePolicy Delete policy.
// Makes the the specified policy inactive, and then deletes it.
func (pl *Policy) DeletePolicy(ctx context.Context, nPolicy int64) error {
	postData := []byte(fmt.Sprintf(`{"nPolicy": %d}`, nPolicy))
	request, err := http.NewRequest("POST", pl.client.Server+"/api/v1.0/Policy.DeletePolicy", bytes.NewBuffer(postData))
	if err != nil {
		return err
	}

	_, err = pl.client.Do(ctx, request, nil)
	return err
}

// PolicyList struct
type PolicyList struct {
	PList []PList `json:"PxgRetVal"`
}

type PList struct {
	Type       *string     `json:"type,omitempty"`
	PListValue *PListValue `json:"value,omitempty"`
}

type PListValue struct {
	KlpolAcceptParent          *bool     `json:"KLPOL_ACCEPT_PARENT,omitempty"`
	KlpolActive                *bool     `json:"KLPOL_ACTIVE,omitempty"`
	KlpolCreated               *DateTime `json:"KLPOL_CREATED,omitempty"`
	KlpolDN                    *string   `json:"KLPOL_DN,omitempty"`
	KlpolForced                *bool     `json:"KLPOL_FORCED,omitempty"`
	KlpolForceDistrib2Children *bool     `json:"KLPOL_FORCE_DISTRIB2CHILDREN,omitempty"`
	KlpolGroupID               *int64    `json:"KLPOL_GROUP_ID,omitempty"`
	KlpolGroupName             *string   `json:"KLPOL_GROUP_NAME,omitempty"`
	KlpolGsynID                *int64    `json:"KLPOL_GSYN_ID,omitempty"`
	KlpolHideOnSlaves          *bool     `json:"KLPOL_HIDE_ON_SLAVES,omitempty"`
	KlpolID                    *int64    `json:"KLPOL_ID,omitempty"`
	KlpolInherited             *bool     `json:"KLPOL_INHERITED,omitempty"`
	KlpolModified              *DateTime `json:"KLPOL_MODIFIED,omitempty"`
	KlpolProduct               *string   `json:"KLPOL_PRODUCT,omitempty"`
	KlpolProfilesNum           *int64    `json:"KLPOL_PROFILES_NUM,omitempty"`
	KlpolRoaming               *bool     `json:"KLPOL_ROAMING,omitempty"`
	KlpolVersion               *string   `json:"KLPOL_VERSION,omitempty"`
}

// GetEffectivePoliciesForGroup
// Obtain policies that affect the specified group.
// Returns active and roaming policies that affect specified group
func (pl *Policy) GetEffectivePoliciesForGroup(ctx context.Context, nGroupId int64) (*PolicyList, error) {
	postData := []byte(fmt.Sprintf(`{"nGroupId": %d}`, nGroupId))
	request, err := http.NewRequest("POST", pl.client.Server+"/api/v1.0/Policy.GetEffectivePoliciesForGroup", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	policyList := new(PolicyList)
	_, err = pl.client.Do(ctx, request, &policyList)
	return policyList, err
}

type OutbreakPolicies struct {
	PxgRetVal *OPData `json:"PxgRetVal,omitempty"`
}

// GetOutbreakPolicies Acquire array of outbreak policies.
// Returns the array of outbreak policies
func (pl *Policy) GetOutbreakPolicies(ctx context.Context) (*OutbreakPolicies, error) {
	request, err := http.NewRequest("POST", pl.client.Server+"/api/v1.0/Policy.GetOutbreakPolicies", nil)
	if err != nil {
		return nil, err
	}

	outbreakPolicies := new(OutbreakPolicies)
	_, err = pl.client.Do(ctx, request, &outbreakPolicies)
	return outbreakPolicies, err
}

// GetPoliciesForGroup Obtain policies for specified group.
// Returns policies located in specified group.
func (pl *Policy) GetPoliciesForGroup(ctx context.Context, nGroupId int64) (*PolicyList, error) {
	postData := []byte(fmt.Sprintf(`{"nGroupId": %d}`, nGroupId))
	request, err := http.NewRequest("POST", pl.client.Server+"/api/v1.0/Policy.GetPoliciesForGroup", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	policyList := new(PolicyList)
	_, err = pl.client.Do(ctx, request, &policyList)
	return policyList, err
}

// GetPolicyContents Acquire policy contents.
// Opens settings storage SsContents of the specified policy.
// The settings storage contains both predefined and product-specific sections.
func (pl *Policy) GetPolicyContents(ctx context.Context, nPolicy, nRevisionId, nLifeTime int64) (*PxgValStr, error) {
	postData := []byte(fmt.Sprintf(`{ "nPolicy": %d , "nRevisionId": %d , "nLifeTime": %d }`, nPolicy, nRevisionId,
		nLifeTime))
	request, err := http.NewRequest("POST", pl.client.Server+"/api/v1.0/Policy.GetPolicyContents", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	pxgValStr := new(PxgValStr)
	_, err = pl.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, err
}

// GetPolicyData Obtain policy data.
// Returns data for specified policy
func (pl *Policy) GetPolicyData(ctx context.Context, nPolicy int64) (*PxgValPolicy, error) {
	postData := []byte(fmt.Sprintf(`{"nPolicy": %d}`, nPolicy))
	request, err := http.NewRequest("POST", pl.client.Server+"/api/v1.0/Policy.GetPolicyData", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	pxgValPolicy := new(PxgValPolicy)
	_, err = pl.client.Do(ctx, request, &pxgValPolicy)
	return pxgValPolicy, err
}

// MakePolicyActive Make policy active or inactive.
// Make the specified policy active or inactive
func (pl *Policy) MakePolicyActive(ctx context.Context, nPolicy int64, bActive bool) (*PxgValBool, error) {
	postData := []byte(fmt.Sprintf(`{"nPolicy": %d, "bActive": %v}`, nPolicy, bActive))
	request, err := http.NewRequest("POST", pl.client.Server+"/api/v1.0/Policy.MakePolicyActive", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	pxgValBool := new(PxgValBool)
	_, err = pl.client.Do(ctx, request, &pxgValBool)
	return pxgValBool, err
}

// MakePolicyRoaming Make policy roaming.
// Make the specified policy roaming
func (pl *Policy) MakePolicyRoaming(ctx context.Context, nPolicy int64) (*PxgValBool, error) {
	postData := []byte(fmt.Sprintf(`{"nPolicy": %d}`, nPolicy))
	request, err := http.NewRequest("POST", pl.client.Server+"/api/v1.0/Policy.MakePolicyRoaming", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	pxgValBool := new(PxgValBool)
	_, err = pl.client.Do(ctx, request, &pxgValBool)
	return pxgValBool, err
}

// RevertPolicyToRevision Revert policy to its older version.
// Replaces the specified policy nPolicy by its revision (older version) nRevisionId
func (pl *Policy) RevertPolicyToRevision(ctx context.Context, nPolicy, nRevisionId int64) error {
	postData := []byte(fmt.Sprintf(`{"nPolicy": %d, "nRevisionId": %d}`, nPolicy, nRevisionId))
	request, err := http.NewRequest("POST", pl.client.Server+"/api/v1.0/Policy.RevertPolicyToRevision", bytes.NewBuffer(postData))
	if err != nil {
		return err
	}

	_, err = pl.client.Do(ctx, request, nil)
	return err
}

// OutbreakPoliciesParams struct
type OutbreakPoliciesParams struct {
	OPData OPData `json:"pData,omitempty"`
}

type OPData struct {
	KlpolPolOutbreak []KlpolPolOutbreak `json:"KLPOL_POL_OUTBREAK"`
}

type KlpolPolOutbreak struct {
	Type  string      `json:"type"`
	Value PolOutbreak `json:"value"`
}

type PolOutbreak struct {
	KlpolID           int64 `json:"KLPOL_ID,omitempty"`
	KlpolOutbreakMask int64 `json:"KLPOL_OUTBREAK_MASK,omitempty"`
}

// SetOutbreakPolicies Specify array of outbreak policies.
// Sets the array of outbreak policies
func (pl *Policy) SetOutbreakPolicies(ctx context.Context, params OutbreakPoliciesParams) error {
	postData, err := json.Marshal(&params)
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", pl.client.Server+"/api/v1.0/Policy.SetOutbreakPolicies", bytes.NewBuffer(postData))
	if err != nil {
		return err
	}

	_, err = pl.client.Do(ctx, request, nil)
	return err
}

//PolicyDataUpdateParams struct using in Policy.UpdatePolicyData
type PolicyDataUpdateParams struct {
	// identifier of the policy to update
	NPolicy int64 `json:"nPolicy"`

	// contains attributes to modify
	PPolicyData PPolicyData `json:"pPolicyData"`
}

// UpdatePolicyData Update policy attributes.
// Updates specified policy attributes.
func (pl *Policy) UpdatePolicyData(ctx context.Context, params PolicyDataUpdateParams) ([]byte, error) {
	postData, err := json.Marshal(&params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", pl.client.Server+"/api/v1.0/Policy.UpdatePolicyData", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pl.client.Do(ctx, request, nil)
	return raw, err
}

// ExportPolicy Export policy to a blob.
// Exports policy into single chunk.
func (pl *Policy) ExportPolicy(ctx context.Context, lPolicy int64) (*PxgValStr, error) {
	postData := []byte(fmt.Sprintf(`{"lPolicy": %d}`, lPolicy))
	request, err := http.NewRequest("POST", pl.client.Server+"/api/v1.0/Policy.ExportPolicy", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	pxgValStr := new(PxgValStr)
	_, err = pl.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, err
}

//PolicyBlob struct
type PolicyBlob struct {
	LGroup int64 `json:"lGroup,omitempty"`
	//	PData raw []byte data of policy base64 encoded
	PData string `json:"pData,omitempty"`
}

// ImportPolicy Import policy from blob.
// Exports policy from a single chunk.
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
