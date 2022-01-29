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

// ListTags Service allows to acquire and manage tags to various KSC objects.
//
// Homogeneous objects, i.e. objects of the same type, can be considered as a list.
// To work with such list, it is convenient to be able to associate string values called tags with list items.
// The tags for each list are completely independent of each other.
//
// Each instance of the ListTags interface is associated with a specific list.
// It is needed to specify ListTagID as Instance (see Request Format of the request)
// The set of possible ListTagIDs is described in The set of lists that support tags
type ListTags service

// GetAllTags Retrieves all known tag values that can be set for a list item
func (lt *ListTags) GetAllTags(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", lt.client.Server+"/api/v1.0/ListTags.GetAllTags", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := lt.client.Request(ctx, request, nil)
	return raw, err
}

// NewTagParams struct
type NewTagParams struct {
	SzwTagValue string `json:"szwTagValue,omitempty"`
	PParams     Null   `json:"pParams,omitempty"`
}

// AddTag Allows to add a new tag value that can be set for a list item
func (lt *ListTags) AddTag(ctx context.Context, params NewTagParams) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", lt.client.Server+"/api/v1.0/ListTags.AddTag", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := lt.client.Request(ctx, request, nil)
	return raw, err
}

// DeleteTags2 Allows to delete multiple tag values in one method call
func (lt *ListTags) DeleteTags2(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", lt.client.Server+"/api/v1.0/ListTags.DeleteTags2", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := lt.client.Request(ctx, request, nil)
	return raw, err
}

// GetTags Retrieves tag values for specified list items
func (lt *ListTags) GetTags(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", lt.client.Server+"/api/v1.0/ListTags.GetTags", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := lt.client.Request(ctx, request, nil)
	return raw, err
}

// RenameTag Allows to rename tag.
// If szwOldTagValue is assigned to some objects then szwNewTagValue will be also assigned to these objects.
// So RenameTag is not equivalent to deleting szwOldTagValue and adding szwNewTagValue
func (lt *ListTags) RenameTag(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", lt.client.Server+"/api/v1.0/ListTags.RenameTag", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := lt.client.Request(ctx, request, nil)
	return raw, err
}

// SetTags Assign tag values for specified list items
func (lt *ListTags) SetTags(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", lt.client.Server+"/api/v1.0/ListTags.SetTags", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := lt.client.Request(ctx, request, nil)
	return raw, err
}
