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
	"fmt"
	"net/http"
)

//	HostMoveRules Class Reference
//
//	Modify and acquire move rules to hosts..
//
//	Interface allows to acquire and manage host moving rules.
//	The rule will move host which fits KLHST_MR_Query to KLHST_MR_Group
//
//	List of all members.
type HostMoveRules service

//Enumerate all extended host moving rules.
//
//Enumerates all extended host moving rules.
//
// Parameters:
//	- pFields	(array) array containing names of attributes to acquire
//
// Returns:
//	- (array) array, each element is (params) object containing attributes of rule
//
//See List of extended host moving rule attributes
func (hs *HostMoveRules) GetRules(ctx context.Context) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`
	{
	"pFields": ["KLHST_MR_ID",
		"KLHST_MR_Group",
		"KLHST_MR_DN",
		"KLHST_MR_Enabled",
		"KLHST_MR_AutoDelete",
		"KLHST_MR_Options",
		"KLHST_MR_Type",
		"KLHST_MR_Query",
		"KLHST_MR_Custom",
		"KLHST_MR_SPECIAL"],
	"pFileds2Return" : ["KLHST_MR_ID",
		"KLHST_MR_Group",
		"KLHST_MR_DN",
		"KLHST_MR_Enabled",
		"KLHST_MR_AutoDelete",
		"KLHST_MR_Options",
		"KLHST_MR_Type",
		"KLHST_MR_Query",
		"KLHST_MR_Custom",
		"KLHST_MR_SPECIAL"]
	}`))
	request, err := http.NewRequest("POST", hs.client.Server+"/api/v1.0/HostMoveRules.GetRules", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hs.client.Do(ctx, request, nil)
	return raw, err
}

//Acquire attributes of specified rule.
//
//Returns attributes of specified rule.
//
//	Parameters:
//	- nRule	(int64) id of rule
//
//	Returns:
//	- (params) object containing attributes of specified rule, see List of extended host moving rule attributes
func (hs *HostMoveRules) GetRule(ctx context.Context, nRule int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nRule": %d}`, nRule))
	request, err := http.NewRequest("POST", hs.client.Server+"/api/v1.0/HostMoveRules.GetRule", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hs.client.Do(ctx, request, nil)
	return raw, err
}

//Remove extended host moving rule.
//
//Removes specified extended host moving rule.
//
//	Parameters:
//	- nRule	(int64) id of rule to remove
func (hs *HostMoveRules) DeleteRule(ctx context.Context, nRule int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nRule": %d}`, nRule))
	request, err := http.NewRequest("POST", hs.client.Server+"/api/v1.0/HostMoveRules.DeleteRule", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := hs.client.Do(ctx, request, nil)
	return raw, err
}

//TODO HostMoveRules.AddRule
//TODO HostMoveRules.ExecuteRulesNow
//TODO HostMoveRules.SetRulesOrder
//TODO HostMoveRules.UpdateRule
