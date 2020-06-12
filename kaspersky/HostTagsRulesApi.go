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

//	HostTagsRulesApi Class Reference
//
//	Interface allows to acquire and manage host automatic tagging rules
//
//	Detailed Description
//
//	Interface allows to acquire and manage host automatic tagging rules.
//	Administration server contains global list of the rules that may automatically set tags for computers.
//	Every rule is identified by szwTagValue that the rule will be set: szwTagValue is a non-empty string,
//	up to 256 unicode characters.
//	Application of the rules happens on by HostTagsRulesApi.ExecuteRule call Periodically.
//	By default every 2 hours, by notification about changing important settings that can change rule query output.
//
//	Application of the rule is set the rule szwTagValue to the hosts falling under a rule scope and reset szwTagValue
//	for other hosts if the tag has been established by the rule earlier
//
//	Public Member Functions:
type HostTagsRulesApi service

//HostTagsRulesParams struct
type HostTagsRulesParams struct {
	PFields2ReturnArray []string `json:"pFields2ReturnArray"`
}

//	Enumerates all rules.
//
//	Returns specified attributes of all rules.
//
//	Parameters:
//	- pFields2ReturnArray	(array) string array with names of requested rule attribute names that need to return see List of host automatic tagging rule attributes.
//{
//	"pFields2ReturnArray" : ["KLHST_HTR_DN",
//		"KLHST_HTR_Enabled",
//		"KLHST_HTR_TagValue",
//		"KLHST_HTR_Custom",
//		"KLHST_HTR_Query"]
//	}
//
//	Returns:
//	- ppRules (params) contains following attributes:
//	"KLHST_HTR_Rules" - host automatic tagging rules (paramArray|paramParams)
//	list of attributes that are specified in pFields2ReturnArray
func (htra *HostTagsRulesApi) GetRules(ctx context.Context, params HostTagsRulesParams) ([]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", htra.client.Server+"/api/v1.0/HostTagsRulesApi.GetRules", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := htra.client.Do(ctx, request, nil)
	return raw, err
}

//	Acquire attributes of specified rule.
//
//	Returns attributes of specified rule.
//
//	Parameters:
//	- szwTagValue	(string). tag of the rule. rule id
//
//	Returns:
//	- (params) object containing attributes of specified rule, see List of host automatic tagging rule attributes.
func (htra *HostTagsRulesApi) GetRule(ctx context.Context, szwTagValue string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"szwTagValue": "%s"}`, szwTagValue))
	request, err := http.NewRequest("POST", htra.client.Server+"/api/v1.0/HostTagsRulesApi.GetRule", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := htra.client.Do(ctx, request, nil)
	return raw, err
}

//	Initiate application of the rule.
//
//	Method initiates application of specified rule.
//
//	It is also allowed to execute disabled rule ( when "KLHST_HTR_Enabled" attribute is false ).
//	Such run is reset rule tag for all hosts where it was been previously set by the rule.
//
//	After returning from this method it is needed to wait while AsyncActionStateChecker.CheckActionState will return bFinalized or call HostTagsRulesApi.CancelAsyncAction with wstrActionGuid
//
//	Parameters:
//	- szwTagValue	(string). tag of the rule. rule id
//	- [out]	wstrActionGuid	(string) id of asynchronous operation, to get status use AsyncActionStateChecker.
//	CheckActionState
func (htra *HostTagsRulesApi) ExecuteRule(ctx context.Context, szwTagValue string) (*WActionGUID, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"szwTagValue": "%s"}`, szwTagValue))
	request, err := http.NewRequest("POST", htra.client.Server+"/api/v1.0/HostTagsRulesApi.ExecuteRule", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	wActionGUID := new(WActionGUID)
	raw, err := htra.client.Do(ctx, request, &wActionGUID)
	return wActionGUID, raw, err
}

//	Cancel asynchronous operation.
//
//	This method should be called if there is no wish to wait while AsyncActionStateChecker.CheckActionState will return bFinalized for earlier launched asynchronous operation.
//
//	Parameters:
//	- wstrActionGuid	(string). id of asynchronous operation that has been started earlier
func (htra *HostTagsRulesApi) CancelAsyncAction(ctx context.Context, wstrActionGuid string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"wstrActionGuid": "%s"}`, wstrActionGuid))
	request, err := http.NewRequest("POST", htra.client.Server+"/api/v1.0/HostTagsRulesApi.CancelAsyncAction", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := htra.client.Do(ctx, request, nil)
	return raw, err
}

//	Remove host automatic tagging rule.
//
//	Parameters:
//	- szwTagValue	(string). tag of the rule. rule id
func (htra *HostTagsRulesApi) DeleteRule(ctx context.Context, szwTagValue string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"szwTagValue": "%s"}`, szwTagValue))
	request, err := http.NewRequest("POST", htra.client.Server+"/api/v1.0/HostTagsRulesApi.DeleteRule", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := htra.client.Do(ctx, request, nil)
	return raw, err
}

//UpdateRuleParams struct using in HostTagsRulesApi.UpdateRule
type UpdateRuleParams struct {
	//tag of the rule. rule id
	SzwTagValue string `json:"szwTagValue"`

	//object containing rule attributes, see List of host automatic tagging rule attributes.
	PRuleInfo PRuleInfo `json:"pRuleInfo"`
}

type PRuleInfo struct {
	//Rule display name.
	KlhstHtrDN string `json:"KLHST_HTR_DN"`

	//Whether rule is turned on.
	KLHSTHTREnabled bool `json:"KLHST_HTR_Enabled"`

	//Tag value that will be set by rule. It is rule identifier.
	KLHSTHTRTagValue string `json:"KLHST_HTR_TagValue"`

	//	Host filtering expression (see Search filter syntax).
	//
	//	Don't use filtering by fields other then listed below.
	//
	//	- KLHST_WKS_FROM_UNASSIGNED
	//	- KLHST_NAG_INSTID
	//	- KLHST_WKS_STATUS
	//	- KLHST_WKS_WINHOSTNAME
	//	- KLHST_WKS_WINDOMAIN
	//	- KLHST_WKS_DNSNAME
	//	- KLHST_WKS_DNSDOMAIN
	//	- KLHST_WKS_IP_LONG
	//	- KLHST_WKS_CONNECT_IP_LONG
	//	- KLHST_WKS_CTYPE
	//	- KLHST_WKS_PTYPE
	//	- KLHST_WKS_OS_VER_MAJOR
	//	- KLHST_WKS_OS_VER_MINOR
	//	- KLHST_WKS_OSSP_VER_MAJOR
	//	- KLHST_WKS_CPU_ARCH
	//	- KLDPNS_ID
	//	- KLHST_WKS_GROUPID
	//	- KLHST_AD_ORGUNIT
	//	- KLHST_AD_ORGUNIT_GP
	//	- KLHST_AD_GROUP
	//	- KLHST_MOB_HAS_OWNER_CERT
	//	- HST_VM_TYPE
	//	- HST_VM_VDI
	//	- KLHST_INVENTORY_PRODUCT_NAME
	//	- KLHST_INVENTORY_PRODUCT_DISPLAY_VERSION
	//	- KLHST_INVENTORY_PRODUCT_PUBLISHER
	KLHSTHTRQuery string `json:"KLHST_HTR_Query"`

	//Any data associated with rule. It is not analyzed by the Administration Server
	KLHSTHTRCustom KLHSTHTRCustom `json:"KLHST_HTR_Custom"`
}

type KLHSTHTRCustom struct {
	//type "params"
	Type string `json:"type"`

	//Any data associated with rule. It is not analyzed by the Administration Server
	CustomValue CustomValue `json:"value"`
}

type CustomValue struct{}

//	Adds/Updates host automatic tagging rule.
//
//	Parameters:
//	- szwTagValue	(wstring). tag of the rule. rule id
//	- pRuleInfo	(params) object containing rule attributes, see List of host automatic tagging rule attributes.
//	Following attributes are required:
//	- "KLHST_HTR_DN"
func (htra *HostTagsRulesApi) UpdateRule(ctx context.Context, params UpdateRuleParams) ([]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", htra.client.Server+"/api/v1.0/HostTagsRulesApi.UpdateRule", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := htra.client.Do(ctx, request, nil)
	return raw, err
}
