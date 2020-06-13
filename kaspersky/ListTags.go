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
	"net/http"
)

//ListTags Class Reference
//
//Interface allows to acquire and manage tags to various KSC objects..
//
//To manage tags for the list with name "ListName" it is needed to create this Interface with instance "ListName"
//Examples of list names: "HostsTags", "InventoryTags", "UmdmDeviceTags"
//
//List of all members.
type ListTags service

//GetAllTags
//Get all existing tag values.
//
//Retrieves all known tag values that can be set for a list item
//
//	Parameters:
//	- pParams	(params) object with attributes:
//"KLTAGS_INCLUDE_VS" (bool) optional flag. true(default) - to include tags information from virtual servers, false - to exclude tags information from virtual servers
//
//	Returns:
//	- (array) of paramString - tags.
func (lt *ListTags) GetAllTags(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", lt.client.Server+"/api/v1.0/ListTags.GetAllTags", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := lt.client.Do(ctx, request, nil)
	return raw, err
}

//AddTag
//Add a new tag.
//
//Allows to add a new tag value that can be set for a list item
//
//	Parameters:
//	- szwTagValue	non-empty tag value of type (string)
//	- pParams	reserved. (params)
func (lt *ListTags) AddTag(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", lt.client.Server+"/api/v1.0/ListTags.AddTag", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := lt.client.Do(ctx, request, nil)
	return raw, err
}

//DeleteTags2
//Delete tags.
//
//Allows to delete multiple tag values in one method call
//
//	Parameters:
//	- pTagValue	(array) of paramString tags that will be deleted
//	- pParams	reserved. (params)
func (lt *ListTags) DeleteTags2(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", lt.client.Server+"/api/v1.0/ListTags.DeleteTags2", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := lt.client.Do(ctx, request, nil)
	return raw, err
}

//GetTags
//Get tag values for specified list items.
//
//Retrieves tag values for specified list items
//
//	Parameters:
//	- pIdArray	(array) collection of list item identifiers . Value semantics depends on the list type (
//	for HostsTags it is the host id and so on)
//	- pParams	reserved. (params)
//
//	Returns:
//	- (array) of paramParams objects where each of them has the following structure:
//"KLTAGS_ITEM_ID" - list item identifier. Type depends on the list type ( for InventoryTags and UmdmDeviceTags it is integer, for HostsTags from 10SP2 it is string )
//"KLTAGS_TAGS" - (array) of (string) tags that are associated with list item
func (lt *ListTags) GetTags(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", lt.client.Server+"/api/v1.0/ListTags.GetTags", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := lt.client.Do(ctx, request, nil)
	return raw, err
}

//RenameTag
//Rename a tag.
//
//Allows to rename tag. If szwOldTagValue is assigned to some objects then szwNewTagValue will be also assigned to these objects. So RenameTag is not equivalent to deleting szwOldTagValue and adding szwNewTagValue
//
//	Parameters:
//	- szwOldTagValue	non-empty old tag value of type (string)
//	- szwNewTagValue	non-empty new tag value of type (string)
//	- pParams	reserved. (params)
func (lt *ListTags) RenameTag(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", lt.client.Server+"/api/v1.0/ListTags.RenameTag", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := lt.client.Do(ctx, request, nil)
	return raw, err
}

//SetTags
//Set tag values for specified list items.
//
//Assign tag values for specified list items
//
//	Parameters:
//	- pListItemsTags	(array) collection of paramParams objects where each of them has the following structure:
//"KLTAGS_ITEM_ID" - list item identifier. Type depends on the list type
//"KLTAGS_TAGS" - (array) collection of (params) objects where each of them has the following structure:
//"KLTAGS_VALUE" - (string) value of a tag being set
//"KLTAGS_SET" - (bool) flag. true - to set tag, false - to reset it
//	- pParams	(params) object of the following structure:
//"KLTAGS_FULL_REPLACE" - (bool) optional flag. true - to make a full replacement of previously set up tags, false(default) - to set up this tag and leave the old ones untouched
func (lt *ListTags) SetTags(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", lt.client.Server+"/api/v1.0/ListTags.SetTags", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := lt.client.Do(ctx, request, nil)
	return raw, err
}
