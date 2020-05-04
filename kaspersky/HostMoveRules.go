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
	"log"
	"net/http"
)

type HostMoveRules struct {
	client *Client
}

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
		log.Fatal(err.Error())
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
		log.Fatal(err.Error())
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
		log.Fatal(err.Error())
	}

	raw, err := hs.client.Do(ctx, request, nil)

	return raw, err
}

//TODO HostMoveRules::AddRule
//TODO HostMoveRules::ExecuteRulesNow
//TODO HostMoveRules::SetRulesOrder
//TODO HostMoveRules::UpdateRule
