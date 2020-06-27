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

//	PackagesApi Operating with packages.
type PackagesApi service

type EULAIDParams struct {
	VecEULAIDs []int64 `json:"vecEulaIDs"`
}

// AcceptEulasAccepts given EULAs.
func (pa *PackagesApi) AcceptEulas(ctx context.Context, params EULAIDParams) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.AcceptEulas",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

// AddExtendedSign Add extended certificate/sign with authenticated attributes to executable file.
//
// Remarks:
// If extended sign already exists - one more sign will be added
func (pa *PackagesApi) AddExtendedSign(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.AddExtendedSign", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

// AddExtendedSignAsync Add extended certificate/sign with authenticated attributes to executable file (asynchronously).
func (pa *PackagesApi) AddExtendedSignAsync(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.AddExtendedSignAsync", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

// AllowSharedPrerequisitesInstallation Allow installation of the shared prerequisites.
func (pa *PackagesApi) AllowSharedPrerequisitesInstallation(ctx context.Context, nPackageId int64, bAllow bool) ([]byte,
	error) {
	postData := []byte(fmt.Sprintf(`{"nPackageId": %d, "bAllow": %v}`, nPackageId, bAllow))
	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.AllowSharedPrerequisitesInstallation", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

// CancelCreateExecutablePkg Cancel an asynchronous call to PackagesApi.CreateExecutablePkgAsync.
func (pa *PackagesApi) CancelCreateExecutablePkg(ctx context.Context, wstrRequestId string) ([]byte,
	error) {
	postData := []byte(fmt.Sprintf(`{"wstrRequestId": "%s"}`, wstrRequestId))
	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.CancelCreateExecutablePkg", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

// CancelGetExecutablePkgFile Cancel an asynchronous call to PackagesApi.GetExecutablePkgFileAsync.
func (pa *PackagesApi) CancelGetExecutablePkgFile(ctx context.Context, wstrRequestId string) ([]byte,
	error) {
	postData := []byte(fmt.Sprintf(`{"wstrRequestId": "%s"}`, wstrRequestId))
	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.CancelGetExecutablePkgFile", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

// CancelRecordNewPackage Cancel an asynchronous call to PackagesApi.RecordVapmPackageAsync or PackagesApi.RecordVapmPackageAsync.
func (pa *PackagesApi) CancelRecordNewPackage(ctx context.Context, wstrRequestId string) ([]byte,
	error) {
	postData := []byte(fmt.Sprintf(`{"wstrRequestId": "%s"}`, wstrRequestId))
	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.CancelRecordNewPackage", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

// CancelUpdateBasesInPackages Cancel an asynchronous call to PackagesApi.UpdateBasesInPackagesAsync.
func (pa *PackagesApi) CancelUpdateBasesInPackages(ctx context.Context, wstrRequestId string) ([]byte,
	error) {
	postData := []byte(fmt.Sprintf(`{"wstrRequestId": "%s"}`, wstrRequestId))
	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.CancelUpdateBasesInPackages", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

// CreateExecutablePkgAsync Create a standalone package (asynchronously).
//
// Remarks:
// Check the operation state by calling AsyncActionStateChecker.CheckActionState periodically until it's finalized
// or cancel it by calling PackagesApi.CancelGetExecutablePkgFile.
func (pa *PackagesApi) CreateExecutablePkgAsync(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.CreateExecutablePkgAsync", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

// DeleteExecutablePkg Delete standalone package.
func (pa *PackagesApi) DeleteExecutablePkg(ctx context.Context, nPackageId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nPackageId": %d}`, nPackageId))

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.DeleteExecutablePkg", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

type EULA struct {
	WstrEULAText string `json:"wstrEulaText,omitempty"`
}

// GetEulaText Requests EULA text.
func (pa *PackagesApi) GetEulaText(ctx context.Context, nEulaId int64) (*EULA, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"nEulaId": %d}`, nEulaId))

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.GetEulaText", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	eula := new(EULA)
	raw, err := pa.client.Do(ctx, request, &eula)
	return eula, raw, err
}

// GetExecutablePackages Get standalone packages.
func (pa *PackagesApi) GetExecutablePackages(ctx context.Context, nPackageId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nPackageId": %d}`, nPackageId))

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.GetExecutablePackages", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

// GetExecutablePkgFileAsync Get standalone package file attributes (asynchronously).
//
// Remarks:
// Check the operation state by calling AsyncActionStateChecker.CheckActionState periodically until it's finalized
// or cancel it by calling PackagesApi.CancelGetExecutablePkgFile.
//
// To download it, client should send an HTTP GET-request to the URL of format as follows: "http://host:port" + KLPKG_EP_DOWNLOAD_PATH
//
// If the action failed then call to AsyncActionStateChecker.CheckActionState returns error in pStateData.
func (pa *PackagesApi) GetExecutablePkgFileAsync(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.GetExecutablePkgFileAsync", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

// GetIncompatibleAppsInfo Get incompatible apps info.
func (pa *PackagesApi) GetIncompatibleAppsInfo(ctx context.Context, nPackageId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nPackageId": %d}`, nPackageId))

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.GetIncompatibleAppsInfo", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

// GetIntranetFolderForNewPackage Get intranet folder for a new package.
func (pa *PackagesApi) GetIntranetFolderForNewPackage(ctx context.Context, wstrProductName,
	wstrProductVersion string) (*PxgValStr, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"wstrProductName": "%s", "wstrProductVersion": "%s"}`, wstrProductName, wstrProductVersion))

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.GetIntranetFolderForNewPackage", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := pa.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

// GetIntranetFolderForPackage Get intranet folder for particular package.
func (pa *PackagesApi) GetIntranetFolderForPackage(ctx context.Context, nPackageId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nPackageId": %d}`, nPackageId))

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.GetIntranetFolderForPackage", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

// GetKpdProfileString Read kpd profile string.
func (pa *PackagesApi) GetKpdProfileString(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.GetKpdProfileString", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

// GetLicenseKey Get license key.
func (pa *PackagesApi) GetLicenseKey(ctx context.Context, nPackageId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nPackageId": %d}`, nPackageId))

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.GetLicenseKey", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

// GetLoginScript Get text of the login script.
func (pa *PackagesApi) GetLoginScript(ctx context.Context, nPackageId int64, wstrTaskId string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nPackageId": %d, "wstrTaskId": "%s"}`, nPackageId, wstrTaskId))

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.GetLoginScript", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

// GetMoveRuleInfo Get information about the move-rule used by the standalone-package.
func (pa *PackagesApi) GetMoveRuleInfo(ctx context.Context, nRuleId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nRuleId": %d}`, nRuleId))

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.GetMoveRuleInfo", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

// GetPackageInfo Get package info.
func (pa *PackagesApi) GetPackageInfo(ctx context.Context, nPackageId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nPackageId": %d}`, nPackageId))

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.GetPackageInfo", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

// GetPackageInfo2 Get package info.
func (pa *PackagesApi) GetPackageInfo2(ctx context.Context, nPackageId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nPackageId": %d}`, nPackageId))

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.GetPackageInfo2", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

// GetPackageInfoFromArchive Get information from archive (zip, cab, tar, tar.gz) with package data
//
// First scenario: Upload archive with kpd-file to FileTransfer -> call PackagesApi.GetPackageInfoFromArchive -> call PackagesApi.RecordNewPackage3.
//
// Second scenario: Upload archive with executable file to FileTransfer -> call PackagesApi.GetPackageInfoFromArchive -> call PackagesApi.RecordNewPackage3.
func (pa *PackagesApi) GetPackageInfoFromArchive(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.GetPackageInfoFromArchive", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

// GetPackagePlugin Get package plugin
func (pa *PackagesApi) GetPackagePlugin(ctx context.Context, nPackageId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nPackageId": %d}`, nPackageId))

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.GetPackagePlugin", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

type Packages struct {
	Packages []Package `json:"PxgRetVal"`
}

type Package struct {
	Type  string        `json:"type"`
	Value PackageStruct `json:"value"`
}

type PackageStruct struct {
	KlpkgNpiCreationTime        KlpkgNpiTime       `json:"KLPKG_NPI_CREATION_TIME"`
	KlpkgNpiModifTime           KlpkgNpiTime       `json:"KLPKG_NPI_MODIF_TIME"`
	KlpkgNpiName                string             `json:"KLPKG_NPI_NAME"`
	KlpkgNpiPackagePath         string             `json:"KLPKG_NPI_PACKAGE_PATH"`
	KlpkgNpiPkgid               int64              `json:"KLPKG_NPI_PKGID"`
	KlpkgNpiProductDisplName    string             `json:"KLPKG_NPI_PRODUCT_DISPL_NAME"`
	KlpkgNpiProductDisplVersion string             `json:"KLPKG_NPI_PRODUCT_DISPL_VERSION"`
	KlpkgNpiProductName         string             `json:"KLPKG_NPI_PRODUCT_NAME"`
	KlpkgNpiProductVersion      string             `json:"KLPKG_NPI_PRODUCT_VERSION"`
	KlpkgNpiSize                KlpkgNpiSize       `json:"KLPKG_NPI_SIZE"`
	KlpkgNpiSsDescr             string             `json:"KLPKG_NPI_SS_DESCR"`
	KlpkgNpiAVBasesUpdateTime   *KlpkgNpiTime      `json:"KLPKG_NPI_AV_BASES_UPDATE_TIME,omitempty"`
	KlpkgNpiAVBasesUpdSupported bool               `json:"KLPKG_NPI_AV_BASES_UPD_SUPPORTED,omitempty"`
	KlpkgNpiExtraData           *KlpkgNpiExtraData `json:"KLPKG_NPI_EXTRA_DATA,omitempty"`
}

type KlpkgNpiTime struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type KlpkgNpiSize struct {
	Type  string `json:"type"`
	Value int64  `json:"value"`
}

type KlpkgNpiExtraData struct {
	KlpggVapmDistribGlbid *KlpkgNpiSize `json:"KLPGG_VAPM_DISTRIB_GLBID,omitempty"`
	KlpkgEULAUid          *Klpkg        `json:"KLPKG_EULA_UID,omitempty"`
	KlpkgFormat           int64         `json:"KLPKG_FORMAT,omitempty"`
	KlpkgIsMSI            bool          `json:"KLPKG_IS_MSI,omitempty"`
	KlpkgLangTag          string        `json:"KLPKG_LANG_TAG,omitempty"`
	KlpkgParentID         int64         `json:"KLPKG_PARENT_ID,omitempty"`
	KlpkgPkgMan           int64         `json:"KLPKG_PKG_MAN,omitempty"`
	KlpkgPlatform         int64         `json:"KLPKG_PLATFORM,omitempty"`
	KlpkgPrdType          int64         `json:"KLPKG_PRD_TYPE,omitempty"`
	KlpkgType             int64         `json:"KLPKG_TYPE,omitempty"`
	BPkgPrereqAllowed     bool          `json:"bPkgPrereqAllowed,omitempty"`
	NPatchGlbID           *KlpkgNpiSize `json:"nPatchGlbId,omitempty"`
	NPatchLcid            int64         `json:"nPatchLcid,omitempty"`
}

type Klpkg struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

// GetPackages Get packages information.
func (pa *PackagesApi) GetPackages(ctx context.Context) (*Packages, []byte, error) {
	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.GetPackages", nil)
	if err != nil {
		return nil, nil, err
	}

	packages := new(Packages)
	raw, err := pa.client.Do(ctx, request, packages)
	return packages, raw, err
}

// GetPackages2 Get packages.
func (pa *PackagesApi) GetPackages2(ctx context.Context) (*Packages, []byte, error) {
	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.GetPackages2", nil)
	if err != nil {
		return nil, nil, err
	}

	packages := new(Packages)
	raw, err := pa.client.Do(ctx, request, packages)
	return packages, raw, err
}

type RebootOptionsEx struct {
	OptionsEx *OptionsEx `json:"PxgRetVal,omitempty"`
}

type OptionsEx struct {
	KlpkgRoptsAskForRebootPeriodMin int64  `json:"KLPKG_ROPTS_ASK_FOR_REBOOT_PERIOD_MIN,omitempty"`
	KlpkgRoptsAskRebootMsgtext      string `json:"KLPKG_ROPTS_ASK_REBOOT_MSGTEXT,omitempty"`
	KlpkgRoptsForceAppsClosed       bool   `json:"KLPKG_ROPTS_FORCE_APPS_CLOSED,omitempty"`
	KlpkgRoptsForceRebootTimeMin    int64  `json:"KLPKG_ROPTS_FORCE_REBOOT_TIME_MIN,omitempty"`
}

// GetRebootOptionsEx Get reboot options.
func (pa *PackagesApi) GetRebootOptionsEx(ctx context.Context, nPackageId int64) (*RebootOptionsEx, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"nPackageId": %d}`, nPackageId))

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.GetRebootOptionsEx", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	rebootOptionsEx := new(RebootOptionsEx)
	raw, err := pa.client.Do(ctx, request, &rebootOptionsEx)
	return rebootOptionsEx, raw, err
}

type UserEULAS struct {
	UserEULA []UserEULA `json:"PxgRetVal"`
}

type UserEULA struct {
	Type          string         `json:"type,omitempty"`
	UserEULAValue *UserEULAValue `json:"value,omitempty"`
}

type UserEULAValue struct {
	KlpkgEULATextParams *KlpkgEULA `json:"KLPKG_EULA_TEXT_PARAMS,omitempty"`
	KlpkgEULAUid        *KlpkgEULA `json:"KLPKG_EULA_UID,omitempty"`
	KlpkgLangTag        string     `json:"KLPKG_LANG_TAG,omitempty"`
	BEULAAccepted       bool       `json:"bEulaAccepted,omitempty"`
	NAgreementType      int64      `json:"nAgreementType,omitempty"`
	NEULADBID           int64      `json:"nEulaDbId,omitempty"`
	NLCID               int64      `json:"nLCID,omitempty"`
}

type KlpkgEULA struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

// GetUserAgreements Request user agreements related to user packages, registered on current VS.
func (pa *PackagesApi) GetUserAgreements(ctx context.Context) (*UserEULAS, []byte, error) {
	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.GetUserAgreements", nil)
	if err != nil {
		return nil, nil, err
	}

	userEULAS := new(UserEULAS)
	raw, err := pa.client.Do(ctx, request, &userEULAS)
	return userEULAS, raw, err
}

// IsPackagePublished Check whether the package is published on KSC web server.
func (pa *PackagesApi) IsPackagePublished(ctx context.Context, nPkgExecId int64) (*PxgValBool, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"nPkgExecId": %d}`, nPkgExecId))

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.IsPackagePublished", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValBool := new(PxgValBool)
	raw, err := pa.client.Do(ctx, request, &pxgValBool)
	return pxgValBool, raw, err
}

// PrePublishMobilePackage Prepare server-side data for mobile package publication on KSC web server.
func (pa *PackagesApi) PrePublishMobilePackage(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.PrePublishMobilePackage", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

// PublishMobileManifest Publish mobile manifest on KSC web server.
func (pa *PackagesApi) PublishMobileManifest(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.PublishMobileManifest", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

// PublishMobilePackage Publish mobile package on KSC web server.
func (pa *PackagesApi) PublishMobilePackage(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.PublishMobilePackage", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

// PublishStandalonePackage Publish/Unpublish a standalone package on KSC web server.
//
// Note:
// You can publish an already published package and vice versa
func (pa *PackagesApi) PublishStandalonePackage(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.PublishStandalonePackage", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

// ReadKpdFile Read kpd file.
func (pa *PackagesApi) ReadKpdFile(ctx context.Context, nPackageId int64) (*PxgValStr, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"nPackageId": %d}`, nPackageId))

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.ReadKpdFile", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := pa.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

// ReadPkgCfgFile Read package configuration file.
func (pa *PackagesApi) ReadPkgCfgFile(ctx context.Context, nPackageId int64, wstrFileName string) (*PxgValStr, []byte,
	error) {
	postData := []byte(fmt.Sprintf(`{"nPackageId": %d, "wstrFileName": "%s"}`, nPackageId, wstrFileName))

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.ReadPkgCfgFile", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := pa.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

// NewPackage struct using in PackagesApi.RecordNewPackage and PackagesApi.RecordNewPackage2
type NewPackage struct {
	WstrPackageName         string `json:"wstrPackageName,omitempty"`
	WstrFileID              string `json:"wstrFileId,omitempty"`
	WstrFolder              string `json:"wstrFolder,omitempty"`
	WstrProductName         string `json:"wstrProductName,omitempty"`
	WstrProductVersion      string `json:"wstrProductVersion,omitempty"`
	WstrProductDisplName    string `json:"wstrProductDisplName,omitempty"`
	WstrProductDisplVersion string `json:"wstrProductDisplVersion,omitempty"`
}

// RecordNewPackage Creates a package with the default settings based on the product, overwritten in the folder,
// the path to which was obtained by calling the PackagesApi.GetIntranetFolderForNewPackage.
func (pa *PackagesApi) RecordNewPackage(ctx context.Context, params NewPackage) (*PxgValStr, []byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.RecordNewPackage",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := pa.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

// RecordNewPackage2 Creates a package with the default settings based on the product, transferred using FileTransfer.
func (pa *PackagesApi) RecordNewPackage2(ctx context.Context, params *NewPackage) (*PxgValStr, []byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.RecordNewPackage2",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := pa.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

// RecordNewPackage3 Create a package using data from archive (zip, cab, tar, tar.gz) or executable file.
//
// Use PackagesApi.GetPackageInfoFromArchive to get package information from archive.
//
// First scenario: Upload archive with kpd-file to FileTransfer -> call PackagesApi.GetPackageInfoFromArchive -> call PackagesApi.RecordNewPackage3.
//
// Second scenario: Upload archive with executable file to FileTransfer -> call PackagesApi.GetPackageInfoFromArchive -> call PackagesApi.RecordNewPackage3.
//
// Third scenario: Upload executable file to FileTransfer -> call PackagesApi.RecordNewPackage3.
func (pa *PackagesApi) RecordNewPackage3(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.RecordNewPackage3", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

// RecordNewPackage3Async Creates a package (asynchronously) with the default settings based on the product, overwritten in the folder,
// the path to which was obtained by calling the PackagesApi.GetIntranetFolderForNewPackage.
//
// Remarks:
// Check the operation state by calling AsyncActionStateChecker.CheckActionState periodically until it's finalized or
// cancel it by calling PackagesApi.CancelRecordNewPackage.
//
// If the operation succeeds then AsyncActionStateChecker.CheckActionState returns these attributes in pStateData container:
// KLPKG_NPI_PKGID - (int64) ID of the executable package.
//
// If the action failed then call to AsyncActionStateChecker.CheckActionState returns error in pStateData.
func (pa *PackagesApi) RecordNewPackage3Async(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.RecordNewPackage3Async", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

// RecordNewPackageAsync Creates a package (asynchronously) with the default settings based on the product,
// overwritten in the folder, the path to which was obtained by calling the PackagesApi.GetIntranetFolderForNewPackage.
//
// Remarks:
// Check the operation state by calling AsyncActionStateChecker.CheckActionState periodically until it's finalized or
// cancel it by calling PackagesApi.CancelRecordNewPackage.
//
// If the operation succedes then AsyncActionStateChecker.CheckActionState returns these attributes in pStateData container:
// KLPKG_NPI_PKGID - (int) ID of the executable package.
//
// If the action failed then call to AsyncActionStateChecker.CheckActionState returns error in pStateData.
func (pa *PackagesApi) RecordNewPackageAsync(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.RecordNewPackageAsync", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

// RecordVapmPackageAsync Create a package using VAPM product information.
//
// Remarks:
// Check the operation state by calling AsyncActionStateChecker.CheckActionState periodically until it's finalized or cancel
// it by calling PackagesApi.CancelRecordNewPackage.
//
// If the operation succedes then AsyncActionStateChecker.CheckActionState returns Package ID in pStateData container
// as KLPKG_NPI_PKGID (paramInt) attribute.
//
// Otherwise, a call to AsyncActionStateChecker.CheckActionState returns error in pStateData.
func (pa *PackagesApi) RecordVapmPackageAsync(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.RecordVapmPackageAsync", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

// RemovePackage Remove a package.
func (pa *PackagesApi) RemovePackage(ctx context.Context, nPackageId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nPackageId": %d}`, nPackageId))

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.RemovePackage", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

// RemovePackageResult struct using in PackagesApi.RemovePackage2
type RemovePackageResult struct {
	BResult bool        `json:"bResult,omitempty"`
	PTasks  interface{} `json:"pTasks"`
}

// RemovePackage2 Remove a package and get the list of dependent tasks.
func (pa *PackagesApi) RemovePackage2(ctx context.Context, nPackageId int64) (*RemovePackageResult, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"nPackageId": %d}`, nPackageId))

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.RemovePackage2", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	removePackageResult := new(RemovePackageResult)
	raw, err := pa.client.Do(ctx, request, &removePackageResult)
	return removePackageResult, raw, err
}

// RenamePackage Rename package.
func (pa *PackagesApi) RenamePackage(ctx context.Context, nPackageId int64, wstrNewPackageName string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nPackageId": %d, "wstrNewPackageName": "%s"}`, nPackageId, wstrNewPackageName))

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.RenamePackage", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

// ResetDefaultServerSpecificSettings Reset server-specific settings for package.
func (pa *PackagesApi) ResetDefaultServerSpecificSettings(ctx context.Context, nPackageId int64) (*PxgValBool, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"nPackageId": %d}`, nPackageId))

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.ResetDefaultServerSpecificSettings", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValBool := new(PxgValBool)
	raw, err := pa.client.Do(ctx, request, &pxgValBool)
	return pxgValBool, raw, err
}

// ResolvePackageLcid Resolve LCID of a package.
func (pa *PackagesApi) ResolvePackageLcid(ctx context.Context, nPackageId, nLcid int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nPackageId": %d, "nLcid": %d}`, nPackageId, nLcid))

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.ResolvePackageLcid", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

// RetranslateToVServerAsync Retranslate package to a Virtual Server (asynchronously).
//
// Remarks:
// Check the operation state by calling AsyncActionStateChecker.CheckActionState periodically until it's finalized.
//
// If the operation succeeds then AsyncActionStateChecker.CheckActionState does not return any attributes in
// pStateData container. If the action failed then call to AsyncActionStateChecker.CheckActionState returns error in pStateData.
func (pa *PackagesApi) RetranslateToVServerAsync(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.RetranslateToVServerAsync", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

// SetLicenseKey Set license key.
func (pa *PackagesApi) SetLicenseKey(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.SetLicenseKey", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

// SetRemoveIncompatibleApps Set incompatible apps info.
func (pa *PackagesApi) SetRemoveIncompatibleApps(ctx context.Context, nPackageId int64, bRemoveIncompatibleApps bool) (*PxgValBool, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"nPackageId": %d, "bRemoveIncompatibleApps": %v}`, nPackageId, bRemoveIncompatibleApps))

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.SetRemoveIncompatibleApps", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValBool := new(PxgValBool)
	raw, err := pa.client.Do(ctx, request, &pxgValBool)
	return pxgValBool, raw, err
}

// SSGetNames Get settings storage section names.
func (pa *PackagesApi) SSGetNames(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.SS_GetNames", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

// SSRead Read settings storage data.
func (pa *PackagesApi) SSRead(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.SS_Read", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

// SSSectionOperation Perform operation on a settings storage section.
func (pa *PackagesApi) SSSectionOperation(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.SS_SectionOperation", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

// SSWrite Write settings storage data.
func (pa *PackagesApi) SSWrite(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.SS_Write", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

// UnpublishMobilePackage Unpublish a previously published mobile package on KSC web server.
func (pa *PackagesApi) UnpublishMobilePackage(ctx context.Context, wstrProfileId string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"wstrProfileId": "%s"}`, wstrProfileId))

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.UnpublishMobilePackage", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

// UpdateBasesInPackagesAsync Get standalone package file attributes (asynchronously).
//
// Remarks:
// Check the operation state by calling AsyncActionStateChecker.CheckActionState periodically until it's finalized or
// cancel it by calling PackagesApi.CancelGetExecutablePkgFile.
//
// To download it, client should send an HTTP GET-request to the URL of format as follows:
// "http://host:port" + KLPKG_EP_DOWNLOAD_PATH If the action failed then call to AsyncActionStateChecker.CheckActionState returns error in pStateData.
func (pa *PackagesApi) UpdateBasesInPackagesAsync(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.UpdateBasesInPackagesAsync", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

// WriteKpdProfileString Write kpd profile string.
func (pa *PackagesApi) WriteKpdProfileString(ctx context.Context, nPackageId int64, wstrSection, wstrKey, wstrValue string) ([]byte,
	error) {
	postData := []byte(fmt.Sprintf(`{"nPackageId": %d, "wstrSection": "%s", "wstrKey": "%s", "wstrValue": "%s"}`,
		nPackageId, wstrSection, wstrKey, wstrValue))

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.WriteKpdProfileString", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

// PkgCFGFileParams struct using in PackagesApi.WritePkgCfgFile
type PkgCFGFileParams struct {
	NPackageID   int64  `json:"nPackageId,omitempty"`
	WstrFileName string `json:"wstrFileName,omitempty"`
	PData        string `json:"pData,omitempty"`
}

// WritePkgCfgFile Write package configuration file.
func (pa *PackagesApi) WritePkgCfgFile(ctx context.Context, params PkgCFGFileParams) ([]byte,
	error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.WritePkgCfgFile", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}
