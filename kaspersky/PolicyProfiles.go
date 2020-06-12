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

//	PolicyProfiles Class Reference
//	Policy profiles processing.
//
//	Provides access to policy profiles.
//
//	To determine which hosts have the specified policy profile active,
//	or which policy profiles are active at the specified host use Srvview KLPOL_PROFILE_HST
//	— active policy profiles at hosts
//
//	List of all members.
type PolicyProfiles service

//	Create a new profile.
//
//	Creates a new profile with the specified name.
//	The method returns identifier of opened SsContents, hat must be filled by the client.
//
//	Actually the profile is saved only after the returned SsContents is filled with one or more sections,
//	and SsContents::Ss_Apply method is called.
//	So, the SsContents::Ss_Apply method may also throw an exception if the profile is not unique.
//
//	Note:
//	Don't forget to fill returned SsContents, and call SsContents::Ss_Apply and SsContents::SS_Release methods
//
//	Parameters:
//	- nPolicy	(int64) policy id
//	- szwName	(string) profile name, a non-empty string, up to 100 unicode characters
//	- pAttrs	(params) profile data, following attributes may be specified in the policy format
//		|- EXPRESSION (required)
//	 	|- KLSSPOL_PRF_ENABLED (optional, false by default)
//	- nLifeTime	(int) timeout in milliseconds to keep this SsContents object alive, zero means 'default value'
//
//	Returns:
//	- (string) identifier of opened SsContents, must be closed with SsContents::SS_Release
func (pp *PolicyProfiles) AddProfile(ctx context.Context, params interface{}) ([]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", pp.client.Server+"/api/v1.0/PolicyProfiles.AddProfile", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pp.client.Do(ctx, request, nil)
	return raw, err
}

//	Delete profile.
//
//	Deletes the specified profile and all data associated.
//
//	See also:
//	Policy profile data
//
//	Parameters:
//	- nPolicy	(int64) policy id
//	- szwName	(string) profile name
//
//	Exceptions:
//	Throws	exception if KLSSPOL_PRF_PROTECTED is set to true
func (pp *PolicyProfiles) DeleteProfile(ctx context.Context, nPolicy int64, szwName string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nPolicy": %d, "szwName": "%s"}`, nPolicy, szwName))
	request, err := http.NewRequest("POST", pp.client.Server+"/api/v1.0/PolicyProfiles.DeleteProfile", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pp.client.Do(ctx, request, nil)
	return raw, err
}

//	Acquire profile list.
//
//	Returns array of all profiles for the specified policy
//
//	Parameters:
//	- nPolicy	(int) policy id
//	- nRevision	(int) policy revision id, zero means 'current policy'
//
//	Returns:
//	- (params) container, each entry has Params type, the name is a profile name (see KLSSPOL_PRF_NAME),
//	and the contents is profile data with following attributes:
//	|- KLSSPOL_PRF_ENABLED
//	|- KLSSPOL_PRF_PROTECTED
//	|- EXPRESSION
func (pp *PolicyProfiles) EnumProfiles(ctx context.Context, nPolicy, nRevision int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nPolicy": %d, "nRevision": %d}`, nPolicy, nRevision))
	request, err := http.NewRequest("POST", pp.client.Server+"/api/v1.0/PolicyProfiles.EnumProfiles", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pp.client.Do(ctx, request, nil)
	return raw, err
}

//	Export policy profile to a blob.
//
//	Exports policy profile into single chunk.
//
//	Parameters:
//	- lPolicy	(int64) policy id
//	- szwName	(string) profile name, a non-empty string, up to 100 unicode characters
//
//	Returns:
//	- blob with exported policy profile
func (pp *PolicyProfiles) ExportProfile(ctx context.Context, lPolicy int64, szwName string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lPolicy": %d, "szwName": "%s"}`, lPolicy, szwName))
	request, err := http.NewRequest("POST", pp.client.Server+"/api/v1.0/PolicyProfiles.ExportProfile", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pp.client.Do(ctx, request, nil)
	return raw, err
}

//	Acquire effective policy contents for host.
//
//	Creates a copy of the settings storage SsContents of the specified policy,
//	and applies to it those policy profiles which are active at the specified host.
//
//	Parameters:
//	- nPolicy	(int64) policy id
//	- szwHostId	(string) host name (see KLHST_WKS_HOSTNAME)
//	- nLifeTime	(int) timeout in milliseconds to keep this SsContents object alive, zero means 'default value'
//
//	Returns:
//	(string) identifier of opened SsContents, must be closed with SsContents::SS_Release
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

//	Acquire profile priority array.
//
//	Returns array of profile names, the profile with lesser index has greater priority.
//
//	Parameters:
//	- nPolicy	(int64) policy id
//	- nRevision	(int64) policy revision id, zero means 'current policy'
//
//	Returns:
//	- (array) array of paramString, each array entry is the profile name
func (pp *PolicyProfiles) GetPriorities(ctx context.Context, nPolicy, nRevision int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nPolicy": %d, "nRevision": %d}`, nPolicy, nRevision))
	request, err := http.NewRequest("POST", pp.client.Server+"/api/v1.0/PolicyProfiles.GetPriorities", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pp.client.Do(ctx, request, nil)
	return raw, err
}

//	Acquire profile attributes.
//
//	Returns profile data for the specified policy profile.
//
//	Parameters:
//	- nPolicy	(int) policy id
//	- nRevision	(int) policy revision id, zero means 'current policy'
//	- szwName	(wstring) profile name
//
//	Returns:
//	- (params) profile data
func (pp *PolicyProfiles) GetProfile(ctx context.Context, nPolicy, nRevision int64, szwName string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nPolicy": %d, "nRevision": %d, "szwName": "%s"}`, nPolicy, nRevision, szwName))
	request, err := http.NewRequest("POST", pp.client.Server+"/api/v1.0/PolicyProfiles.GetProfile", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pp.client.Do(ctx, request, nil)
	return raw, err
}

//	Acquire profile contents.
//
//	Returns SsContents interface for the profile contents
//
//	Parameters:
//	- nPolicy	(int64) policy id
//	- nRevision	(int64) policy revision id, zero means 'current policy'
//	- szwName	(string) profile name, a non-empty string, up to 100 unicode characters
//	- nLifeTime	(int64) timeout in milliseconds to keep this SsContents object alive, zero means 'default value'
//
//	Returns:
//	(string) identifier of opened SsContents, must be closed with SsContents::SS_Release
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

//	Import policy profile from blob.
//
//	Exports policy profile from a single chunk.
//
//	Parameters:
//	- lPolicy	(int64) policy id
//	- pData		(binary base64 string) policy exported by PolicyProfiles.ExportProfile, cannot be NULL
//
//	Returns:
//	- profile name
func (pp *PolicyProfiles) ImportProfile(ctx context.Context, lPolicy int64, pData string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lPolicy": %d, "pData": "%s"}`, lPolicy, pData))
	request, err := http.NewRequest("POST", pp.client.Server+"/api/v1.0/PolicyProfiles.ImportProfile", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pp.client.Do(ctx, request, nil)
	return raw, err
}

//	Update profile priority array.
//
//	Updates the array of profile names, the profile with lesser index has greater priority.
//	Inexisting profiles are ignored, unmentioned profiles are appended.
//	Profile with KLSSPOL_PRF_PROTECTED set to true cannot be reordered.
//
//	Parameters:
//	- nPolicy	(int64) policy id
//	- pArrayOfNames	(array) of paramString, each array entry is the profile name
func (pp *PolicyProfiles) PutPriorities(ctx context.Context, params interface{}) ([]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", pp.client.Server+"/api/v1.0/PolicyProfiles.PutPriorities", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pp.client.Do(ctx, request, nil)
	return raw, err
}

//	Rename profile.
//
//	Changes the name of the existing profile.
//
//	Parameters:
//	- nPolicy			(int64) policy id
//	- szwExistingName	(string) existing profile name
//	- szwNewName		(string) new profile name, a non-empty string, up to 100 unicode characters
//
//	Exceptions:
//	Throws	exception if KLSSPOL_PRF_PROTECTED is set to true or profile with szwNewName does not exist
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

//	Update attributes of an existing profile.
//
//	Changes attributes of an existing profile.
//
//	Parameters:
//	- nPolicy	(int64) policy id
//	- szwName	(string) profile name, a non-empty string, up to 100 unicode characters
//	- pAttrsToUpdate	(params) profile data, following attributes may be specified:
//
//	EXPRESSION
//	KLSSPOL_PRF_ENABLED
func (pp *PolicyProfiles) UpdateProfile(ctx context.Context, params interface{}) ([]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", pp.client.Server+"/api/v1.0/PolicyProfiles.UpdateProfile", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pp.client.Do(ctx, request, nil)
	return raw, err
}
