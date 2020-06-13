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

//	PackagesApi Class Reference
//
//	Operating with packages...
//
//	List of all members.
type PackagesApi service

type EULAIDParams struct {
	VecEULAIDs []int64 `json:"vecEulaIDs"`
}

//	Accepts given EULAs.
//
//	Parameters:
//	- vecEulaIDs	(array) set of EULA IDs to accept, each item is paramInt
//
//	Exceptions:
//	Throws	exception in case of error.
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

//TODO PackagesApi::AddExtendedSign
//TODO PackagesApi::AddExtendedSignAsync

//	Allow installation of the shared prerequisites.
//
//	Parameters:
//	- nPackageId	(int64) Package ID.
//	- bAllow	(bool) true to allow the installation, false otherwise.
//
//	Exceptions:
//	Throws	exception in case of error.
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

//	Cancel an asynchronous call to PackagesApi.CreateExecutablePkgAsync.
//
//	Parameters:
//	- wstrRequestId	(string) Request ID.
//
//	Exceptions:
//	Throws	exception in case of error.
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

//	Cancel an asynchronous call to PackagesApi.GetExecutablePkgFileAsync.
//
//	Parameters:
//	- wstrRequestId	(string) Request ID.
//
//	Exceptions:
//	Throws	exception in case of error.
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

//	Cancel an asynchronous call to PackagesApi.RecordVapmPackageAsync or PackagesApi.RecordVapmPackageAsync.
//
//	Parameters:
//	- szwRequestId	(string) Request ID of the recording process.
//
//	Exceptions:
//	Throws	exception in case of error.
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

//	Cancel an asynchronous call to PackagesApi.UpdateBasesInPackagesAsync.
//
//	Parameters:
//	- wstrRequestId	(string) Request ID.
//
//	Exceptions:
//	Throws	exception in case of error.
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

//TODO PackagesApi::CreateExecutablePkgAsync

//	Delete standalone package.
//
//	Parameters:
//	- nPackageId	(int64) Executable package ID.
//
//	Exceptions:
//	Throws	exception in case of error.
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

//	Requests EULA text.
//
//	Parameters:
//	- nEulaId	(int64) EULA id
//
//	Return:
//	- wstrEulaText	(string) EULA text
//
//	Exceptions:
//	Throws	exception in case of error.
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

//	Get standalone packages.
//
//	Parameters:
//	- nPackageId	(int64) Package ID (positive integer or 0 = No package filtration, or -1 = Mobile packages only)
//
//Returns:
//	- (param) Packages info container.
//		|- "KLPKG_evpExecs" - Array of packages (paramArray <paramParams>)
//		|- "KLPKG_evpExecPkgId" - Standalone package ID (paramInt)
//		|- "KLPKG_evpPkgId" - Package ID (paramInt)
//		|- "KLPKG_evpPkgPath" - Path to the file in the shared folder (paramString)
//		|- "KLPKG_evpPkgSize" - Package size (paramInt)
//		|- "KLPKG_evpAddPkgId" - Additional package ID (paramInt)
//		|- "KLPKG_ProdName" - Application ID (paramString)
//		|- "KLPKG_ProdVersion" - Application version (paramString)
//		|- "KLPKG_ProdPkgName" - Product package name (paramString)
//		|- "KLPKG_ProdDisplayVersion" - Application version (paramString)
//		|- "KLPKG_NagentDisplayVersion" - Agent version (paramString)
//		|- "KLPKG_CreationDate" - Creation date and time (paramDateTime)
//		|- "KLPKG_ModificationDate" - Modification date and time (paramDateTime)
//		|- "KLPKG_IsVirtual" - Is virtual package (paramBool)
//		|- "KLPKG_IsPublished" - Is package published (paramBool)
//		|- "KLPKG_NAME" - Package name (paramString)
//		|- "KLPKG_WebURL" - Package published URL (paramString)
//		|- "KLPKG_EP_SHA256" - Package Sha256 in hex format (paramString)
//
//	Exceptions:
//	Throws	exception in case of error.
func (pa *PackagesApi) GetExecutablePackages(ctx context.Context, nPackageId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nPackageId": %d}`, nPackageId))

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.GetExecutablePackages", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

//TODO PackagesApi::GetExecutablePkgFileAsync

//	Get incompatible apps info.
//
//	Parameters:
//	- nPackageId	(int64) Package ID.
//
//	Returns:
//	- (param) Incompatible apps info.
//		|- "KLPKG_GIAI_RESULT" - Is operation completed successfully (paramBool)
//		|- "KLPKG_GIAI_INFO" - Info (param)
//					|- "KLPKG_IncompatibleAppInfoType" - Info type (enum) (paramInt) 0 = Plain text
//		|- "KLPKG_IncompatibleAppInfoPlainText" - Info data (paramString)
//		|- "KLPKG_GIAI_CAN_REMOVE_BY_INSTALLER" - Is delete-incompatible-apps supported
//		|- "KLPKG_GIAI_REMOVE_BY_INSTALLER" - Should incompatible-apps be deleted by installer
//
//	Exceptions:
//	Throws	exception in case of error.
func (pa *PackagesApi) GetIncompatibleAppsInfo(ctx context.Context, nPackageId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nPackageId": %d}`, nPackageId))

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.GetIncompatibleAppsInfo", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

//	Get intranet folder for a new package.
//
//	Parameters:
//	- wstrProductName	(string) Product name.
//	- wstrProductVersion	(string) Product version.
//
//	Returns:
//	- (string) Intranet folder path.
//
//	Exceptions:
//	Throws	exception in case of error.
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

//	Get intranet folder for particular package.
//
//	Parameters:
//	- nPackageId	(int64) Package ID.
//
//	Returns:
//	(string) Intranet folder path.
//
//	Exceptions:
//	Throws	exception in case of error.
func (pa *PackagesApi) GetIntranetFolderForPackage(ctx context.Context, nPackageId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nPackageId": %d}`, nPackageId))

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.GetIntranetFolderForPackage", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

//TODO PackagesApi::GetKpdProfileString

//	Get license key.
//
//	Parameters:
//	- nPackageId	(int64) Package ID.
//
//	Return:
//	- wstrKeyFileName	(string) Key file name.
//	- pMemoryChunk	(binary) Key data.
//
//	Exceptions:
//	Throws	exception in case of error.
func (pa *PackagesApi) GetLicenseKey(ctx context.Context, nPackageId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nPackageId": %d}`, nPackageId))

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.GetLicenseKey", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

//	Get text of the login script.
//
//	Parameters:
//	- nPackageId	(int64) Package ID.
//	- wstrTaskId	(string) Task id.
//
//	Returns:
//	- (string) Script text.
//
//	Exceptions:
//	Throws	exception in case of error.
func (pa *PackagesApi) GetLoginScript(ctx context.Context, nPackageId int64, wstrTaskId string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nPackageId": %d, "wstrTaskId": "%s"}`, nPackageId, wstrTaskId))

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.GetLoginScript", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

//	Get information about the move-rule used by the standalone-package.
//
//	Parameters:
//	- nRuleId	Move-rule ID.
//
//	Returns:
//	- (param) Move-rule information
//		|- "KLPKG_evpExecs" - (paramArray (paramParams)) Array of standalone-packages that use the rule
//			|- "KLPKG_evpExecPkgId" - (paramInt) Standalone-package ID
//			|- "KLPKG_evpPkgId" - (paramInt) Installation package ID
func (pa *PackagesApi) GetMoveRuleInfo(ctx context.Context, nRuleId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nRuleId": %d}`, nRuleId))

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.GetMoveRuleInfo", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

//	Get package info.
//
//	Parameters:
//	- nPackageId	(int64) Package ID.
//
//	Returns:
//	(param) Package info List of package attributes.
//
//	Exceptions:
//	Throws	exception in case of error.
func (pa *PackagesApi) GetPackageInfo(ctx context.Context, nPackageId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nPackageId": %d}`, nPackageId))

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.GetPackageInfo", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

//	Get package info.
//
//	Parameters:
//	- nPackageId	(int64) Package ID.
//
//	Returns:
//	(param) Package info List of package attributes.
//
//	Exceptions:
//	Throws	exception in case of error.
func (pa *PackagesApi) GetPackageInfo2(ctx context.Context, nPackageId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nPackageId": %d}`, nPackageId))

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.GetPackageInfo2", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

//TODO  PackagesApi::GetPackageInfoFromArchive

//	Get package plugin
//
//	Parameters:
//	- nPackageId	(int64) Package ID
//
//	Returns:
//	Information about plugin-file (or empty Params if plugin-file not found)
//	- "KLPKG_PLUGIN_FILENAME" - (paramString) Short filename of package plugin-file
//	- "KLPKG_PLUGIN_FILESIZE" - (paramLong) Size of package plugin-file (in bytes)
//	- "KLPKG_PLUGIN_FILEURL" - (paramString) URL which can be used to download package plugin-file
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

//	Get packages information.
//
//	Returns:
//	- (array) of (param) Each array entry is a paramParams container with attributes described in section List
//	of package attributes
//
//	Exceptions:
//	Throws	exception in case of error.
func (pa *PackagesApi) GetPackages(ctx context.Context) (*Packages, []byte, error) {
	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.GetPackages", nil)
	if err != nil {
		return nil, nil, err
	}

	packages := new(Packages)
	raw, err := pa.client.Do(ctx, request, packages)
	return packages, raw, err
}

//	Get packages.
//
//	Returns:
//	- (array) Packages array. Each item is paramParams with List of package attributes.
//
//	Exceptions:
//	Throws	exception in case of error.
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

//	Get reboot options.
//
//	Parameters:
//	- nPackageId	(int) Package ID.
//
//	Returns:
//	(params) Reboot options.
//	- "KLPKG_ROPTS_ASK_REBOOT_MSGTEXT" - Reboot request text (paramString)
//	- "KLPKG_ROPTS_ASK_FOR_REBOOT_PERIOD_MIN" - Period (in minutes) of display reboot message (paramLong)
//	- "KLPKG_ROPTS_FORCE_REBOOT_TIME_MIN" - Force reboot time (in minutes) (paramLong)
//	- "KLPKG_ROPTS_FORCE_APPS_CLOSED" - Force applications close (paramBool)
//
//	Exceptions:
//	Throws	exception in case of error.
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

//	Request user agreements related to user packages, registered on current VS.
//
//	Returns:
//	(array) of (paramParams), each element describes an agreement and might contain the following attributes:
//	- "nEulaDbId" - Database id of EULA (paramInt).
//	- "nAgreementType" - type of the user license agreement (paramInt, AT_EULA = 0, AT_KSN = 1).
//	- "strEULA" - Agreement text (paramString).
//	- "KLPKG_LANG_TAG" - Agreement language tag (paramString).
//	- "bEulaAccepted" - true if the agreement is accepted for the current virtual server, false if it's not (paramBool).
//	- "nLCID" - Agreement LCID (paramInt).
//
//	Exceptions:
//	Throws	exception in case of error.
func (pa *PackagesApi) GetUserAgreements(ctx context.Context) (*UserEULAS, []byte, error) {
	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.GetUserAgreements", nil)
	if err != nil {
		return nil, nil, err
	}

	userEULAS := new(UserEULAS)
	raw, err := pa.client.Do(ctx, request, &userEULAS)
	return userEULAS, raw, err
}

//	Check whether the package is published on KSC web server.
//
//	Parameters:
//	- nPkgExecId	(int64) Executable ID of the package.
//
//	Returns:
//	- (boolean) true if package is published, false otherwise.
//
//	Exceptions:
//	Throws	exception in case of error.
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

//TODO PackagesApi::PrePublishMobilePackage
//TODO PackagesApi::PublishMobileManifest
//TODO PackagesApi::PublishMobilePackage
//TODO PackagesApi::PublishStandalonePackage

//	Read kpd file.
//
//	Parameters:
//	- nPackageId	(int64) Package ID.
//
//	Returns:
//	- (binary) Read contents.
//
//	Exceptions:
//	Throws	exception in case of error.
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

//	Read package configuration file.
//
//	Parameters:
//	- nPackageId	(int64) Package ID.
//	- wstrFileName	(string) File name.
//
//	Returns:
//	- (binary) Readed contents.
//
//	Exceptions:
//	Throws	exception in case of error.
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

//NewPackage struct using in PackagesApi.RecordNewPackage and PackagesApi.RecordNewPackage2
type NewPackage struct {
	WstrPackageName         string `json:"wstrPackageName,omitempty"`
	WstrFileID              string `json:"wstrFileId,omitempty"`
	WstrFolder              string `json:"wstrFolder,omitempty"`
	WstrProductName         string `json:"wstrProductName,omitempty"`
	WstrProductVersion      string `json:"wstrProductVersion,omitempty"`
	WstrProductDisplName    string `json:"wstrProductDisplName,omitempty"`
	WstrProductDisplVersion string `json:"wstrProductDisplVersion,omitempty"`
}

//	Creates a package with the default settings based on the product,
//	overwritten in the folder, the path to which was obtained by calling the PackagesApi.GetIntranetFolderForNewPackage.
//
//	Parameters:
//	- params (NewPackage)
//	|- wstrPackageName	(string) Package name.
//	|- wstrFolder	(string) Product folder (obtained by calling the PackagesApi::GetIntranetFolderForNewPackage).
//	|- wstrProductName	(string) Product name.
//	|- wstrProductVersion	(string) Product version.
//	|- wstrProductDisplName	(string) Product display name.
//	|- wstrProductDisplVersion	(string) Product display version.
//
//	Returns:
//	- (params) Container with package attributes List of package attributes.
//
//	Exceptions:
//	Throws	exception in case of error.
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

//	Creates a package with the default settings based on the product, transferred using FileTransfer.
//
//	Parameters:
//	- params (NewPackage)
//	|- wstrPackageName	(string) File id.
//	|- wstrFileId	(string) Package name.
//	|- wstrFolder	(string) Product folder (obtained by calling the PackagesApi::GetIntranetFolderForNewPackage).
//	|- wstrProductName	(string) Product name.
//	|- wstrProductVersion	(string) Product version.
//	|- wstrProductDisplName	(string) Product display name.
//	|- wstrProductDisplVersion	(string) Product display version.
//
//	Returns:
//	- (params) Container with package attributes List of package attributes.
//
//	Exceptions:
//	Throws	exception in case of error.
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

//TODO PackagesApi::RecordNewPackage3
//TODO PackagesApi::RecordNewPackage3Async
//TODO PackagesApi::RecordNewPackageAsync
//TODO PackagesApi::RecordVapmPackageAsync

//	Remove a package.
//
//	Parameters:
//	- nPackageId	(int64) Package ID.
//
//	Exceptions:
//	Throws	exception in case of error.
func (pa *PackagesApi) RemovePackage(ctx context.Context, nPackageId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nPackageId": %d}`, nPackageId))

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.RemovePackage", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

//RemovePackageResult struct using in PackagesApi.RemovePackage2
type RemovePackageResult struct {
	BResult bool        `json:"bResult,omitempty"`
	PTasks  interface{} `json:"pTasks"` //TODO?
}

//	Remove a package and get the list of dependent tasks.
//
//	Parameters:
//	- nPackageId	(int64) Package ID.
//
//	Return:
//	- bResult	(bool) Operation success flag.
//	- pTasks	(array) Array of dependent tasks. Each array entry is a paramParams container with attributes
//		|- "KLPKG_TASKINFO_TASK_ID" - Task ID (paramLong)
//		|- "KLPKG_TASKINFO_GROUP_ID" - Task group ID (paramLong)
//		|- "KLPKG_TASKINFO_TASK_DISPLNAME" - Task display name (paramString)
//		|- "KLPKG_TASKINFO_TASK_GROUP_DISPLNAME" - Task group display name (paramString)
//
//	Exceptions:
//	Throws	exception in case of error.
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

//	Rename package.
//
//	Parameters:
//	- nPackageId	(int64) Package ID.
//	- wstrNewPackageName	(string) New name of the package.
//
//	Exceptions:
//	Throws	exception in case of error.
func (pa *PackagesApi) RenamePackage(ctx context.Context, nPackageId int64, wstrNewPackageName string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nPackageId": %d, "wstrNewPackageName": "%s"}`, nPackageId, wstrNewPackageName))

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.RenamePackage", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

//	Reset server-specific settings for package.
//
//	Parameters:
//	- nPackageId	(int64) - package identifier.
//
//	Returns:
//	- (bool) True if reset complete successfully
//
//	Exceptions:
//	Throws	exception in case of error.
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

//	Resolve LCID of a package.
//
//	Parameters:
//	- nPackageId	(int64) Package ID.
//	- nLcid	(int64) LCID.
//
//	Exceptions:
//	Throws	exception in case of error.
func (pa *PackagesApi) ResolvePackageLcid(ctx context.Context, nPackageId, nLcid int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"nPackageId": %d, "nLcid": %d}`, nPackageId, nLcid))

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.ResolvePackageLcid", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

//TODO PackagesApi::RetranslateToVServerAsync
//TODO PackagesApi::SetLicenseKey

//	Set incompatible apps info.
//
//	Parameters:
//	- nPackageId	(int64) Package ID.
//	- bRemoveIncompatibleApps	(bool) Remove incompatible apps flag.
//
//	Returns:
//	- (bool) Operation success flag.
//
//	Exceptions:
//	Throws	exception in case of error.
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

//TODO PackagesApi::SS_GetNames
//TODO PackagesApi::SS_Read
//TODO PackagesApi::SS_SectionOperation
//TODO PackagesApi::SS_Write

//	Unpublish a previously published mobile package on KSC web server.
//
//	Parameters:
//	- wstrProfileId	(string) Profile ID.
//
//	Exceptions:
//	Throws	exception in case of error.
func (pa *PackagesApi) UnpublishMobilePackage(ctx context.Context, wstrProfileId string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"wstrProfileId": "%s"}`, wstrProfileId))

	request, err := http.NewRequest("POST", pa.client.Server+"/api/v1.0/PackagesApi.UnpublishMobilePackage", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := pa.client.Do(ctx, request, nil)
	return raw, err
}

//TODO PackagesApi::UpdateBasesInPackagesAsync

//	Write kpd profile string.
//
//	Parameters:
//	- nPackageId	(int64) Package ID.
//	- wstrSection	(string) Profile storage section.
//	- wstrKey	(string) Profile storage key.
//	- wstrValue	(string) Value to be written.
//
//	Exceptions:
//	Throws	exception in case of error.
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

//PkgCFGFileParams struct using in PackagesApi.WritePkgCfgFile
type PkgCFGFileParams struct {
	NPackageID   int64  `json:"nPackageId,omitempty"`
	WstrFileName string `json:"wstrFileName,omitempty"`
	PData        string `json:"pData,omitempty"`
}

//	Write package configuration file.
//
//	Parameters:
//	- params PkgCFGFileParams
//		|- nPackageId	(int64) Package ID.
//		|- wstrFileName	(string) File name.
//		|- pData	(binary base64 encoded string) Contents to be written.
//
//	Exceptions:
//	Throws	exception in case of error.
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
