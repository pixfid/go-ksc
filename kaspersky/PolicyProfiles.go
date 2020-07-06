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

// PolicyProfiles service Policy profiles processing.
//
// Provides access to policy profiles.
//
// To determine which hosts have the specified policy profile active, or which policy profiles are active
// at the specified host use Srvview KLPOL_PROFILE_HST active policy profiles at hosts
type PolicyProfiles service

// AddProfile Create a new profile.
//
// Creates a new profile with the specified name.
// The method returns identifier of opened SsContents, hat must be filled by the client.
//
// Actually the profile is saved only after the returned SsContents is filled with one or more sections,
// and SsContents.SsApply method is called.
// So, the SsContents.SsApply method may also throw an exception if the profile is not unique.
//
// Note:
// Don't forget to fill returned SsContents, and call SsContents::Ss_Apply and SsContents::SS_Release methods
func (pp *PolicyProfiles) AddProfile(ctx context.Context, params interface{}) ([]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", pp.client.Server+"/api/v1.0/PolicyProfiles.AddProfile", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pp.client.Do(ctx, request, nil)
	return raw, err
}

// DeleteProfile Deletes the specified profile and all data associated.
func (pp *PolicyProfiles) DeleteProfile(ctx context.Context, nPolicy int64, szwName string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nPolicy": %d, "szwName": "%s"}`, nPolicy, szwName))
	request, err := http.NewRequest("POST", pp.client.Server+"/api/v1.0/PolicyProfiles.DeleteProfile", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pp.client.Do(ctx, request, nil)
	return raw, err
}

// EnumProfiles Acquire profile list.
//
// Returns array of all profiles for the specified policy
func (pp *PolicyProfiles) EnumProfiles(ctx context.Context, nPolicy, nRevision int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nPolicy": %d, "nRevision": %d}`, nPolicy, nRevision))
	request, err := http.NewRequest("POST", pp.client.Server+"/api/v1.0/PolicyProfiles.EnumProfiles", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pp.client.Do(ctx, request, nil)
	return raw, err
}

// ExportProfile Export policy profile to a blob.
func (pp *PolicyProfiles) ExportProfile(ctx context.Context, lPolicy int64, szwName string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lPolicy": %d, "szwName": "%s"}`, lPolicy, szwName))
	request, err := http.NewRequest("POST", pp.client.Server+"/api/v1.0/PolicyProfiles.ExportProfile", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pp.client.Do(ctx, request, nil)
	return raw, err
}

// GetEffectivePolicyContents Acquire effective policy contents for host.
//
// Creates a copy of the settings storage SsContents of the specified policy,
// and applies to it those policy profiles which are active at the specified host.
func (pp *PolicyProfiles) GetEffectivePolicyContents(ctx context.Context, nPolicy, nLifeTime int64,
	szwHostId string) (*PxgValStr, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"nPolicy": %d, "szwHostId": "%s", "nLifeTime": %d}`, nPolicy, szwHostId,
		nLifeTime))
	request, err := http.NewRequest("POST", pp.client.Server+"/api/v1.0/PolicyProfiles.GetEffectivePolicyContents", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := pp.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

// GetPriorities Acquire profile priority array.
//
// Returns array of profile names, the profile with lesser index has greater priority.
func (pp *PolicyProfiles) GetPriorities(ctx context.Context, nPolicy, nRevision int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nPolicy": %d, "nRevision": %d}`, nPolicy, nRevision))
	request, err := http.NewRequest("POST", pp.client.Server+"/api/v1.0/PolicyProfiles.GetPriorities", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pp.client.Do(ctx, request, nil)
	return raw, err
}

// GetProfile Acquire profile attributes.
//
// Returns profile data for the specified policy profile.
func (pp *PolicyProfiles) GetProfile(ctx context.Context, nPolicy, nRevision int64, szwName string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nPolicy": %d, "nRevision": %d, "szwName": "%s"}`, nPolicy, nRevision, szwName))
	request, err := http.NewRequest("POST", pp.client.Server+"/api/v1.0/PolicyProfiles.GetProfile", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pp.client.Do(ctx, request, nil)
	return raw, err
}

// GetProfileSettings Acquire profile contents.
//
// Returns SsContents interface for the profile contents
func (pp *PolicyProfiles) GetProfileSettings(ctx context.Context, nPolicy, nLifeTime, nRevision int64,
	szwName string) (*PxgValStr, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"nPolicy": %d, "nRevision": %d, "szwName": "%s", "nLifeTime": %d}`, nPolicy,
		nRevision,
		szwName,
		nLifeTime))
	request, err := http.NewRequest("POST", pp.client.Server+"/api/v1.0/PolicyProfiles.GetProfileSettings", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := pp.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

// ImportProfile Import policy profile from blob.
func (pp *PolicyProfiles) ImportProfile(ctx context.Context, lPolicy int64, pData string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lPolicy": %d, "pData": "%s"}`, lPolicy, pData))
	request, err := http.NewRequest("POST", pp.client.Server+"/api/v1.0/PolicyProfiles.ImportProfile", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pp.client.Do(ctx, request, nil)
	return raw, err
}

// ProfilesPrioritiesParams struct
type ProfilesPrioritiesParams struct {
	// NPolicy policy id
	NPolicy int64 `json:"nPolicy,omitempty"`
	//  PArrayOfNames each array entry is the profile name
	PArrayOfNames []string `json:"pArrayOfNames"`
}

// PutPriorities Update profile priority array.
func (pp *PolicyProfiles) PutPriorities(ctx context.Context, params ProfilesPrioritiesParams) ([]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", pp.client.Server+"/api/v1.0/PolicyProfiles.PutPriorities", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pp.client.Do(ctx, request, nil)
	return raw, err
}

// RenameProfile Changes the name of the existing profile.
func (pp *PolicyProfiles) RenameProfile(ctx context.Context, nPolicy int64, szwExistingName, szwNewName string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nPolicy": %d, "szwExistingName": "%s", "szwExistingName": "%s"}`, nPolicy,
		szwExistingName, szwNewName))
	request, err := http.NewRequest("POST", pp.client.Server+"/api/v1.0/PolicyProfiles.RenameProfile", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pp.client.Do(ctx, request, nil)
	return raw, err
}

// UpdateProfile Update attributes of an existing profile.
func (pp *PolicyProfiles) UpdateProfile(ctx context.Context, params interface{}) ([]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", pp.client.Server+"/api/v1.0/PolicyProfiles.UpdateProfile", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pp.client.Do(ctx, request, nil)
	return raw, err
}
