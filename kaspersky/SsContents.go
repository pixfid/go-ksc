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

//	SsContents Class Reference
//
//	Access to settings storage. More...
//
//	List of all members.
type SsContents service

//TODO Ss_Add

//	Save changes
//
//	Saves changes made by methods:
//
//	Ss_Update,
//	Ss_Add,
//	Ss_Replace,
//	Ss_Clear,
//	Ss_Delete,
//	Ss_CreateSection,
//	Ss_DeleteSection
//
//	Parameters:
//	- wstrID	(string) identifier of opened SsContents
func (sc *SsContents) Ss_Apply(ctx context.Context, wstrID string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`	{ "wstrID": "%s"}`, wstrID))
	request, err := http.NewRequest("POST", sc.client.Server+"/api/v1.0/SsContents.Ss_Apply", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := sc.client.Do(ctx, request, nil)
	return raw, err
}

//TODO Ss_Clear

//	Create section in settings storage.
//
//	Creates empty section in settings storage
//
//	Changes are not saved until method Ss_Apply is called. Unsaved data is not available by methods Ss_Read and SS_GetNames.
//
//	Parameters:
//	- wstrID	(string) identifier of opened SsContents
//	- wstrProduct	(string) product name string, non-empty string, not longer than 31 character,
//	and cannot contain characters /\:*?"<>.
//	- wstrVersion	(string) version string, non-empty string, not longer than 31 character,
//	and cannot contain characters /\:*?"<>.
//	- wstrSection	(string) section name string, non-empty string, not longer than 31 character,
//	and cannot contain characters /\:*?"<>.
func (sc *SsContents) Ss_CreateSection(ctx context.Context, wstrID, wstrProduct, wstrVersion, wstrSection string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`
	{ 
		"wstrID": "%s",
		"wstrProduct": "%s",
		"wstrVersion": "%s",
		"wstrSection": "%s" }`, wstrID, wstrProduct, wstrVersion, wstrSection))
	request, err := http.NewRequest("POST", sc.client.Server+"/api/v1.0/SsContents.Ss_CreateSection", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := sc.client.Do(ctx, request, nil)
	return raw, err
}

//TODO Ss_Delete

func (sc *SsContents) Ss_DeleteSection(ctx context.Context, wstrID, wstrProduct, wstrVersion, wstrSection string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`
	{ 
		"wstrID": "%s",
		"wstrProduct": "%s",
		"wstrVersion": "%s",
		"wstrSection": "%s" }`, wstrID, wstrProduct, wstrVersion, wstrSection))
	request, err := http.NewRequest("POST", sc.client.Server+"/api/v1.0/SsContents.Ss_DeleteSection", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := sc.client.Do(ctx, request, nil)
	return raw, err
}

//	Enumerate contents of settings storage
//
//	Retrieves list of sections.
//
//	Parameters:
//	- wstrID	(string) identifier of opened SsContents
//	- wstrProduct	(string) product name string, string, not longer than 31 character,
//	and cannot contain characters /\:*?"<>. Empty string means that list of products is needed.
//	- wstrVersion	(string) version string, string, not longer than 31 character,
//	and cannot contain characters /\:*?"<>. Empty string means that list of versions is needed.
//
//	Returns:
//	- (array) list of section if both wstrProduct and wstrVersion specified,
//	list of version if only wstrProduct is specified, list of product is wstrProduct is not specified
func (sc *SsContents) SS_GetNames(ctx context.Context, wstrID, wstrProduct, wstrVersion string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{ "wstrID": "%s" , "wstrProduct": "%s" , "nLifeTime": "%s"}`,
		wstrID,
		wstrProduct,
		wstrVersion))
	request, err := http.NewRequest("POST", sc.client.Server+"/api/v1.0/SsContents.SS_GetNames", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := sc.client.Do(ctx, request, nil)
	return raw, err
}

//	Read data from settings storage.
//
//	Reads saved data from the specified section of settings storage.
//
//	Parameters:
//	- wstrID	(string) identifier of opened SsContents. Use PolicyProfiles.GetEffectivePolicyContents to get it
//	- wstrProduct	(string) product name string (see Settings storage section parameters)
//	- wstrVersion	(string) version string (see Settings storage section parameters)
//	- wstrSection	(string) section name string (see Settings storage section parameters)
//
//	Returns:
//	- (params) data from the specified section
//
//	See also:
//	Settings storage section parameters
func (sc *SsContents) Ss_Read(ctx context.Context, wstrID, wstrProduct, wstrVersion, wstrSection string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`
	{ 
		"wstrID": "%s",
		"wstrProduct": "%s",
		"wstrVersion": "%s",
		"wstrSection": "%s" }`, wstrID, wstrProduct, wstrVersion, wstrSection))
	request, err := http.NewRequest("POST", sc.client.Server+"/api/v1.0/SsContents.Ss_Read", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := sc.client.Do(ctx, request, nil)
	return raw, err
}

//	Close opened SsContents
//
//	Closes opened SsContents and releases associated server resources.
//	After calling this method wstrID is not longer valid.
//
//	Parameters:
//	- wstrID	(string) identifier of opened SsContents
func (pl *SsContents) Ss_Release(ctx context.Context, wstrID string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"wstrID": "%s"}`, wstrID))
	request, err := http.NewRequest("POST", pl.client.Server+"/api/v1.0/SsContents.Ss_Release", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pl.client.Do(ctx, request, nil)
	return raw, err
}

//TODO Ss_Replace
//TODO Ss_Update
