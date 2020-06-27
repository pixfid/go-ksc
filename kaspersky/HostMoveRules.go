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

// HostMoveRules service to Modify and acquire move rules to hosts.
//
// Service allows to acquire and manage host moving rules.
// The rule will move host which fits KLHST_MR_Query to KLHST_MR_Group
type HostMoveRules service

// AddRule Creates new extended host moving rule with specified attributes.
func (hmr *HostMoveRules) AddRule(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", hmr.client.Server+"/api/v1.0/HostMoveRules.AddRule", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hmr.client.Do(ctx, request, nil)
	return raw, err
}

// DeleteRule Removes specified extended host moving rule.
func (hmr *HostMoveRules) DeleteRule(ctx context.Context, nRule int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nRule": %d}`, nRule))
	request, err := http.NewRequest("POST", hmr.client.Server+"/api/v1.0/HostMoveRules.DeleteRule", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hmr.client.Do(ctx, request, nil)
	return raw, err
}

// ExecuteRulesParams struct using in HostMoveRules.ExecuteRulesNow
type ExecuteRulesParams struct {
	// NGroupID group to launch rules for
	NGroupID int64 `json:"nGroupId"`

	// PRules array of rule ids
	PRules []int64 `json:"pRules"`

	// NOptions possible values:
	//  0 - rule is processed for hosts that need it
	//  1 - execute even if rule has been already executed
	NOptions int64 `json:"nOptions"`
}

// ExecuteRulesNow Executes rules for a specific group
func (hmr *HostMoveRules) ExecuteRulesNow(ctx context.Context, params ExecuteRulesParams) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", hmr.client.Server+"/api/v1.0/HostMoveRules.ExecuteRulesNow", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hmr.client.Do(ctx, request, nil)
	return raw, err
}

// HMoveRule struct using in HostMoveRules.GetRule
type HMoveRule struct {
	HMRule *HMRule `json:"PxgRetVal,omitempty"`
}

// HMRule struct
type HMRule struct {
	KLHSTMRAutoDelete bool            `json:"KLHST_MR_AutoDelete,omitempty"`
	KLHSTMRCustom     *KLHSTMRCustom  `json:"KLHST_MR_Custom,omitempty"`
	KlhstMrDN         string          `json:"KLHST_MR_DN,omitempty"`
	KLHSTMREnabled    bool            `json:"KLHST_MR_Enabled,omitempty"`
	KLHSTMRGroup      int64           `json:"KLHST_MR_Group,omitempty"`
	KLHSTMROptions    int64           `json:"KLHST_MR_Options,omitempty"`
	KLHSTMRQuery      string          `json:"KLHST_MR_Query,omitempty"`
	KlhstMrSpecial    *KlhstMrSpecial `json:"KLHST_MR_SPECIAL,omitempty"`
	KLHSTMRType       int64           `json:"KLHST_MR_Type,omitempty"`
}

type KLHSTMRCustom struct {
	Type  string              `json:"type,omitempty"`
	Value *KLHSTMRCustomValue `json:"value,omitempty"`
}

type KLHSTMRCustomValue struct {
	HruleFromUnassigned    bool    `json:"HRULE_FROM_UNASSIGNED,omitempty"`
	HruleIncludeChildOu    bool    `json:"HRULE_INCLUDE_CHILD_OU,omitempty"`
	HruleNagentStatus      int64   `json:"HRULE_NAGENT_STATUS,omitempty"`
	HruleOSVersions        []int64 `json:"HRULE_OS_VERSIONS"`
	HruleQueryPart1        string  `json:"HRULE_QUERY_PART1,omitempty"`
	HruleQueryPart3        string  `json:"HRULE_QUERY_PART3,omitempty"`
	HruleQueryPart4        string  `json:"HRULE_QUERY_PART4,omitempty"`
	HruleUserCERTInstalled int64   `json:"HRULE_USER_CERT_INSTALLED,omitempty"`
	KlhstAdGroup           int64   `json:"KLHST_AD_GROUP,omitempty"`
	KlhstAdOrgunit         int64   `json:"KLHST_AD_ORGUNIT,omitempty"`
	OSBuild                int64   `json:"OsBuild,omitempty"`
	OSBuildCond            int64   `json:"OsBuildCond,omitempty"`
	OSRelease              int64   `json:"OsRelease,omitempty"`
	OSReleaseCond          int64   `json:"OsReleaseCond,omitempty"`
}

type KlhstMrSpecial struct {
	Type  string               `json:"type,omitempty"`
	Value *KLHSTMRSPECIALValue `json:"value,omitempty"`
}

type KLHSTMRSPECIALValue struct {
	KlhstMrSpecialAd *KlhstMrSpecialAd `json:"KLHST_MR_SPECIAL_AD,omitempty"`
}

type KlhstMrSpecialAd struct {
	Type  string                 `json:"type,omitempty"`
	Value *KLHSTMRSPECIALADValue `json:"value,omitempty"`
}

type KLHSTMRSPECIALADValue struct {
	KlhstMrSpecialAdCreateSubgroups bool  `json:"KLHST_MR_SPECIAL_AD_CREATE_SUBGROUPS,omitempty"`
	KlhstMrSpecialAdDeleteSubgroups bool  `json:"KLHST_MR_SPECIAL_AD_DELETE_SUBGROUPS,omitempty"`
	KlhstMrSpecialAdMoveToSubgroups bool  `json:"KLHST_MR_SPECIAL_AD_MOVE_TO_SUBGROUPS,omitempty"`
	KlhstMrSpecialAdOuid            int64 `json:"KLHST_MR_SPECIAL_AD_OUID,omitempty"`
}

type HMoveRules struct {
	HMRules []HMRules `json:"PxgRetVal,omitempty"`
}

type HMRules struct {
	Type   string  `json:"type,omitempty"`
	HMRule *HMRule `json:"value,omitempty"`
}

// GetRule Acquire attributes of specified rule.
func (hmr *HostMoveRules) GetRule(ctx context.Context, nRule int64) (*HMoveRule, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"nRule": %d}`, nRule))
	request, err := http.NewRequest("POST", hmr.client.Server+"/api/v1.0/HostMoveRules.GetRule", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	hMoveRule := new(HMoveRule)
	raw, err := hmr.client.Do(ctx, request, &hMoveRule)
	return hMoveRule, raw, err
}

// Rules struct
type Rules struct {
	PFields []string `json:"pFields"`
}

// GetRules Enumerate all extended host moving rules. Enumerates all extended host moving rules.
func (hmr *HostMoveRules) GetRules(ctx context.Context, params Rules) (*HMoveRules, []byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", hmr.client.Server+"/api/v1.0/HostMoveRules.GetRules", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	hMoveRules := new(HMoveRules)
	raw, err := hmr.client.Do(ctx, request, &hMoveRules)
	return hMoveRules, raw, err
}

// RulesOrderParams struct using in HostMoveRules.SetRulesOrder
type RulesOrderParams struct {
	// PRules array of rule ids
	PRules []int64 `json:"pRules"`
}

// SetRulesOrder Modifies order of specified rules in the global list. Order of rules not contained in pRules array will be indefinite.
func (hmr *HostMoveRules) SetRulesOrder(ctx context.Context, params RulesOrderParams) (*HMoveRules, []byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", hmr.client.Server+"/api/v1.0/HostMoveRules.SetRulesOrder", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	hMoveRules := new(HMoveRules)
	raw, err := hmr.client.Do(ctx, request, &hMoveRules)
	return hMoveRules, raw, err
}

// UpdateRule Modify attributes of specified rule.
func (hmr *HostMoveRules) UpdateRule(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", hmr.client.Server+"/api/v1.0/HostMoveRules.UpdateRule", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hmr.client.Do(ctx, request, nil)
	return raw, err
}
