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

//	HostTagsRulesApi Class Reference
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
func (hs *HostTagsRulesApi) GetRules(ctx context.Context, params HostTagsRulesParams) ([]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", hs.client.Server+"/api/v1.0/HostTagsRulesApi.GetRules", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hs.client.Do(ctx, request, nil)
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
func (hs *HostTagsRulesApi) GetRule(ctx context.Context, szwTagValue string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"szwTagValue": "%s"}`, szwTagValue))
	request, err := http.NewRequest("POST", hs.client.Server+"/api/v1.0/HostTagsRulesApi.GetRule", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hs.client.Do(ctx, request, nil)
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
func (hs *HostTagsRulesApi) ExecuteRule(ctx context.Context, szwTagValue string) (*WActionGUID, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"szwTagValue": "%s"}`, szwTagValue))
	request, err := http.NewRequest("POST", hs.client.Server+"/api/v1.0/HostTagsRulesApi.ExecuteRule", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	wActionGUID := new(WActionGUID)
	raw, err := hs.client.Do(ctx, request, &wActionGUID)
	return wActionGUID, raw, err
}

//	Cancel asynchronous operation.
//
//	This method should be called if there is no wish to wait while AsyncActionStateChecker.CheckActionState will return bFinalized for earlier launched asynchronous operation.
//
//	Parameters:
//	- wstrActionGuid	(string). id of asynchronous operation that has been started earlier
func (hs *HostTagsRulesApi) CancelAsyncAction(ctx context.Context, wstrActionGuid string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"wstrActionGuid": "%s"}`, wstrActionGuid))
	request, err := http.NewRequest("POST", hs.client.Server+"/api/v1.0/HostTagsRulesApi.CancelAsyncAction", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hs.client.Do(ctx, request, nil)
	return raw, err
}

//	Remove host automatic tagging rule.
//
//	Parameters:
//	- szwTagValue	(string). tag of the rule. rule id
func (hs *HostTagsRulesApi) DeleteRule(ctx context.Context, szwTagValue string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"szwTagValue": "%s"}`, szwTagValue))
	request, err := http.NewRequest("POST", hs.client.Server+"/api/v1.0/HostTagsRulesApi.DeleteRule", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hs.client.Do(ctx, request, nil)
	return raw, err
}

//TODO HostTagsRulesApi::UpdateRule
