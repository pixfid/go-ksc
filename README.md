# go-ksc #
go-ksc is a Go client library for accessing the KSC (Kaspersky Security Center) Open API.

## Usage ##

```go
package main

import (
    "github.com/pixfid/go-ksc/kaspersky"
)

```

Construct a new KSC client, then use the various services on the client to
access different parts of the KSC (Kaspersky Security Center) Open API. For example:

```go

package main

import (
	"context"
	"fmt"
	"github.com/pixfid/go-ksc/kaspersky"
)

func main() {        
        ctx := context.Background()
    	cfg := kaspersky.Config {
    		Username: "login",
    		Password: "password",
    		Server: fmt.Sprintf(`https://%s:%s`, "ip", "port"),
    	}
    
    	client := kaspersky.New(cfg)
    	client.KSCAuth(ctx)

        //Get List of Windows domain in the network.
        raw,_ := client.HostGroup.GetDomains(context.Background())
        println(string(raw))
}
```

As example find online hosts:
```go
func Online(ctx context.Context, client *kaspersky.Client) *FullHostsInfo {
	hField := config.Config.HParams
	chunks := &FullHostsInfo{}
	hostsParam := kaspersky.HGParams{
		WstrFilter: `
		(&
			(KLHST_WKS_GROUPID_GP <> 4)
			(KLHST_WKS_STATUS&1<>0)
		)`,
		VecFieldsToReturn: hField,
		PParams: kaspersky.PParams{
			KlsrvhSlaveRecDepth:    0,
			KlgrpFindFromCurVsOnly: true,
		},
		LMaxLifeTime: 100,
	}

	accessor, _, _ := client.HostGroup.FindHosts(ctx, hostsParam)
	count, _, _ := client.ChunkAccessor.GetItemsCount(ctx, accessor.StrAccessor)
	_, _ = client.ChunkAccessor.GetItemsChunk(ctx, kaspersky.ItemsChunkParams{
		StrAccessor: accessor.StrAccessor,
		NStart:      0,
		NCount:      count.Int,
	}, chunks)

	client.ChunkAccessor.Release(ctx, accessor.StrAccessor)
	return chunks
}
```


### List of currently, supported method and classes:
___
* [x] **interface AKPatches** - Interface to manage system of autoupdating by patch.exe patches
    * [x]   - func (akp *AKPatches) ApprovePatch(ctx context.Context, params interface{}) ([]byte, error)
    * [x]   - func (akp *AKPatches) ForbidPatch(ctx context.Context, params interface{}) ([]byte, error)
* [x] **interface AdHosts** - Scanned active directory OU structure
    * [x]   - func (ah *AdHosts) FindAdGroups(ctx context.Context, params FindAdGroupsParams) (*ADHostIterator, []byte, error)
    * [x]   - func (ah *AdHosts) GetChildComputer(ctx context.Context, params ChildComputerParams) (*AdHstIDParent, []byte, error)
    * [x]   - func (ah *AdHosts) GetChildComputers(ctx context.Context, params ChildComputersParams) (*PxgValStr, []byte, error)
    * [x]   - func (ah *AdHosts) GetChildOUs(ctx context.Context, params ChildOUParams) (*PxgValStr, []byte, error)
    * [x]   - func (ah *AdHosts) GetOU(ctx context.Context, params OUAttributesParams) (*OUAttributes, []byte, error)
    * [x]   - func (ah *AdHosts) UpdateOU(ctx context.Context, idOU int, params interface{}) ([]byte, error)

* [x] **interface AdfsSso** - Interface for working with ADFS SSO
    * [x]   - func (as *AdfsSso) GetSettings(ctx context.Context) ([]byte, error)
    * [x]   - func (as *AdfsSso) SetSettings(ctx context.Context, params interface{}) ([]byte, error)

* [x] **interface AdmServerSettings** - AdmServerSettings interface
    * [x]   - func (as *AdmServerSettings) ChangeSharedFolder(ctx context.Context, wstrNetworkPath string) ([]byte, error)
    * [x]   - func (as *AdmServerSettings) GetSharedFolder(ctx context.Context) (*PxgValStr, []byte, error)

* [x] **interface AppCtrlApi** - Interface to get info about execution files
    * [x]   - func (ac *AppCtrlApi) GetExeFileInfo(ctx context.Context, params ExeFileInfoParams) ([]byte, error)

* [x] **interface AsyncActionStateChecker** - Interface to monitor state of async action
    * [x]   - func (ac *AsyncActionStateChecker) CheckActionState(ctx context.Context, wstrActionGuid string) (*ActionStateResult, []byte, error)

* [x] **interface CertPoolCtrl** - Interface to manage the pool of certificates used by the Kaspersky Security Center Server
    * [x]   - func (cp *CertPoolCtrl) GetCertificateInfo(ctx context.Context, nVServerId, nFunction int64) ([]byte, error)

* [x] **interface CertPoolCtrl2** - 2nd interface to manage the pool of certificates used by the Kaspersky Security Center Server
    * [x]   - func (cp *CertPoolCtrl2) GetCertificateInfoDetails(ctx context.Context, nVServerId, nFunction int64) ([]byte, error)

* [x] **interface CgwHelper** - CgwHelper (Connection Gateway) helper proxy class
    * [x]   - func (cp *CgwHelper) GetNagentLocation(ctx context.Context, wsHostName string) ([]byte, error)
    * [x]   - func (cp *CgwHelper) GetSlaveServerLocation(ctx context.Context, nSlaveServerId int64) ([]byte, error)

* [x] **interface ChunkAccessor** - Working with host result-set
    * [x]   - func (ca *ChunkAccessor) GetItemsChunk(ctx context.Context, params ItemsChunkParams, result interface{}) ([]byte, error)
    * [x]   - func (ca *ChunkAccessor) GetItemsCount(ctx context.Context, accessor string) (*PxgValInt, []byte, error)
    * [x]   - func (ca *ChunkAccessor) Release(ctx context.Context, accessor string) bool

* [x] **interface ConEvents** - Interface to server events
    * [x]   - func (ce *ConEvents) Retrieve(ctx context.Context) ([]byte, error)
    * [x]   - func (ce *ConEvents) Subscribe(ctx context.Context, v interface{}) ([]byte, error)
    * [x]   - func (ce *ConEvents) UnSubscribe(ctx context.Context, nSubsId int64) ([]byte, error)

* [x] **interface DataProtectionApi** - Allows to protect sensitive data in policies, tasks, and/or on specified host
    * [x]   - func (dp *DataProtectionApi) CheckPasswordSplPpc(ctx context.Context, szwPassword string) (*PxgValBool, []byte, error)
    * [x]   - func (dp *DataProtectionApi) ProtectUtf16StringForHost(ctx context.Context, szwHostId, szwPlainText string) ([]byte, error)
    * [x]   - func (dp *DataProtectionApi) ProtectUtf16StringGlobally(ctx context.Context, szwPlainText string) ([]byte, error)
    * [x]   - func (dp *DataProtectionApi) ProtectUtf8StringForHost(ctx context.Context, szwHostId, szwPlainText string) ([]byte, error)
    * [x]   - func (dp *DataProtectionApi) ProtectUtf8StringGlobally(ctx context.Context, szwPlainText string) ([]byte, error)

* [x] **interface DatabaseInfo** - Database processing
    * [x]   - func (di *DatabaseInfo) CheckBackupPath(ctx context.Context, szwPath string) (*PxgValBool, []byte, error)
    * [x]   - func (di *DatabaseInfo) CheckBackupPath2(ctx context.Context, szwWinPath, szwLinuxPath string) (*PxgValBool, []byte, error)
    * [x]   - func (di *DatabaseInfo) GetDBDataSize(ctx context.Context) (*PxgValInt, []byte, error)
    * [x]   - func (di *DatabaseInfo) GetDBEventsCount(ctx context.Context) (*PxgValInt, []byte, error)
    * [x]   - func (di *DatabaseInfo) GetDBSize(ctx context.Context) (*PxgValInt, []byte, error)
    * [x]   - func (di *DatabaseInfo) IsCloudSQL(ctx context.Context, nCloudType int64) (*PxgValBool, []byte, error)
    * [x]   - func (di *DatabaseInfo) IsLinuxSQL(ctx context.Context) (*PxgValBool, []byte, error)

* [x] **interface DpeKeyService** - Interface for working with encrypted devices
    * [x]   - func (di *DpeKeyService) GetDeviceKeys3(ctx context.Context, wstrDeviceId string) ([]byte, error)

* [x] **interface EventNotificationProperties** - Notification properties
    * [x]   - func (enp *EventNotificationProperties) GetDefaultSettings(ctx context.Context) ([]byte, error)
    * [x]   - func (enp *EventNotificationProperties) GetNotificationLimits(ctx context.Context) ([]byte, error)
    * [x]   - func (enp *EventNotificationProperties) TestNotification(ctx context.Context, eType int, pSettings interface{}) ([]byte, error)

* [x] **interface EventNotificationsApi** - Publish event
    * [x]   - func (ts *EventNotificationsApi) PublishEvent(ctx context.Context, params interface{}) ([]byte, error)

* [x] **interface EventProcessing** - Interface implements the functionality for viewing and deleting events
    * [x]   - func (ts *EventProcessing) CancelDelete(ctx context.Context, params interface{}) ([]byte, error)
    * [x]   - func (ts *EventProcessing) GetRecordCount(ctx context.Context, strIteratorId string) (*PxgValInt, []byte, error)
    * [x]   - func (ts *EventProcessing) GetRecordRange(ctx context.Context, strIteratorId string, nStart, nEnd int64) ([]byte, error)
    * [x]   - func (ts *EventProcessing) InitiateDelete(ctx context.Context, params interface{}) ([]byte, error)
    * [x]   - func (ts *EventProcessing) ReleaseIterator(ctx context.Context, strIteratorId string) (*PxgValInt, []byte, error)

* [x] **interface EventProcessingFactory** - Interface to create event processing iterators
    * [x]   - func (epf *EventProcessingFactory) CreateEventProcessing(ctx context.Context, params interface{}) (*StrIteratorId, []byte, error)
    * [x]   - func (epf *EventProcessingFactory) CreateEventProcessing2(ctx context.Context, params interface{}) (*StrIteratorId, []byte, error)
    * [x]   - func (epf *EventProcessingFactory) CreateEventProcessingForHost(ctx context.Context, params interface{}) (*StrIteratorId, []byte, error)
    * [x]   - func (epf *EventProcessingFactory) CreateEventProcessingForHost2(ctx context.Context, params interface{}) (*StrIteratorId, []byte, error)

* [x] **interface FileCategorizer2** - Interface for working with FileCategorizer subsystem
    * [x]   - func (fc *FileCategorizer2) CancelFileMetadataOperations(ctx context.Context) (*PxgValInt, []byte, error)
    * [x]   - func (fc *FileCategorizer2) CancelFileUpload(ctx context.Context) (*PxgValInt, []byte, error)
    * [x]   - func (fc *FileCategorizer2) CreateCategory(ctx context.Context, params CategoryParams) (*PxgValStr, []byte, error)
    * [x]   - func (fc *FileCategorizer2) DeleteCategory(ctx context.Context, nCategoryId int64) ([]byte, error)
    * [x]   - func (fc *FileCategorizer2) ForceCategoryUpdate(ctx context.Context, nCategoryId int64) ([]byte, error)
    * [x]   - func (fc *FileCategorizer2) GetCategoriesModificationCounter(ctx context.Context) (*PxgValInt, []byte, error)
    * [x]   - func (fc *FileCategorizer2) GetCategory(ctx context.Context, nCategoryId int64) ([]byte, error)
    * [x]   - func (fc *FileCategorizer2) GetCategoryByUUID(ctx context.Context, pCategoryUUID string) ([]byte, error)
    * [x]   - func (fc *FileCategorizer2) GetFileMetadata(ctx context.Context, ulFlag int64) ([]byte, error)
    * [x]   - func (fc *FileCategorizer2) GetFilesMetadata(ctx context.Context, ulFlag int64) ([]byte, error)
    * [x]   - func (fc *FileCategorizer2) GetFilesMetadataFromMSI(ctx context.Context, ulFlag int64) ([]byte, error)
    * [x]   - func (fc *FileCategorizer2) GetRefPolicies(ctx context.Context, nCatId int64) (*RefPolicies, []byte, error)
    * [x]   - func (fc *FileCategorizer2) GetSerializedCategoryBody(ctx context.Context, nCategoryId int64) ([]byte, error)
    * [x]   - func (fc *FileCategorizer2) GetSerializedCategoryBody2(ctx context.Context, nCategoryId int64) ([]byte, error)
    * [x]   - func (fc *FileCategorizer2) GetSyncId(ctx context.Context) (*PxgValInt, []byte, error)

* [x] **interface GroupSync** - Access to group synchronization objects
    * [x]   - func (gs *GroupSync) GetSyncInfo(ctx context.Context, params GroupSyncInfoParams) (*GroupSyncInfo, []byte, error)

* [x] **interface GroupSyncIterator** - Access to the group synchronization forward iterator for the result-set
    * [x]   - func (ca *GroupSyncIterator) GetNextItems(ctx context.Context, szwIterator string, nCount int64, v interface{}) ([]byte, error)
    * [x]   - func (ca *GroupSyncIterator) ReleaseIterator(ctx context.Context, szwIterator string) bool

* [x] **interface GroupTaskControlApi** - Interface to perform some management actions over group tasks
    * [x]   - func (gta *GroupTaskControlApi) CommitImportedTask(ctx context.Context, wstrId string, bCommit bool) (*TaskDescribe, []byte, error)
    * [x]   - func (gta *GroupTaskControlApi) ExportTask(ctx context.Context, wstrTaskId string) (*PxgValStr, []byte, error)
    * [x]   - func (gta *GroupTaskControlApi) GetTaskByRevision(ctx context.Context, nObjId, nRevision int64) (*TaskDescribe, []byte, error)
    * [x]   - func (gta *GroupTaskControlApi) RequestStatistics(ctx context.Context, params TasksIDSParams) ([]byte, error)
    * [x]   - func (gta *GroupTaskControlApi) RestoreTaskFromRevision(ctx context.Context, nObjId, nRevision int64) (*TaskDescribe, []byte, error)

* [x] **interface HWInvStorage** - Interface for working with Hardware storage subsystem
    * [x]   - func (hw *HWInvStorage) EnumDynColumns(ctx context.Context) ([]byte, error)
    * [x]   - func (hw *HWInvStorage) ExportHWInvStorage2(ctx context.Context, eExportType int) ([]byte, error)
    * [x]   - func (hw *HWInvStorage) ExportHWInvStorageGetData(ctx context.Context, wstrAsyncId string) ([]byte, error)
    * [x]   - func (hw *HWInvStorage) GetHWInvObject(ctx context.Context, nObjId int64) ([]byte, error)
    * [x]   - func (hw *HWInvStorage) GetProcessingRules(ctx context.Context) ([]byte, error)
    * [x]   - func (hw *HWInvStorage) ImportHWInvStorage2(ctx context.Context, eImportType int64) ([]byte, error)

* [x] **interface HostGroup** - Hosts and management groups processing
    * [x]   - func (hg *HostGroup) AddDomain(ctx context.Context, strDomain string, nType int64) ([]byte, error)
    * [x]   - func (hg *HostGroup) AddGroup(ctx context.Context, name string, parentId int) (*PxgValInt, []byte, error)
    * [x]   - func (hg *HostGroup) AddGroupHostsForSync(ctx context.Context, nGroupId int64, strSSType string) (*WActionGUID, []byte, error)
    * [x]   - func (hg *HostGroup) AddHost(ctx context.Context, params interface{}) (*PxgValStr, []byte, error)
    * [x]   - func (hg *HostGroup) AddIncident(ctx context.Context, params AddIncidentsParams) (*PxgValStr, []byte, error)
    * [x]   - func (hg *HostGroup) DelDomain(ctx context.Context, strDomain string) ([]byte, error)
    * [x]   - func (hg *HostGroup) DeleteIncident(ctx context.Context, nId int64) ([]byte, error)
    * [x]   - func (hg *HostGroup) FindGroups(ctx context.Context, params HGParams) (*Accessor, []byte, error)
    * [x]   - func (hg *HostGroup) FindHosts(ctx context.Context, params HGParams) (*Accessor, []byte, error)
    * [x]   - func (hg *HostGroup) FindHostsAsync(ctx context.Context, params HGParams) (*RequestID, []byte, error)
    * [x]   - func (hg *HostGroup) FindHostsAsyncCancel(ctx context.Context, strRequestId string)
    * [x]   - func (hg *HostGroup) FindHostsAsyncGetAccessor(ctx context.Context, strRequestId string) (*AsyncAccessor, []byte, error)
    * [x]   - func (hg *HostGroup) FindIncidents(ctx context.Context, params FindIncidentsParams) (*Accessor, []byte, error)
    * [x]   - func (hg *HostGroup) FindUsers(ctx context.Context, params UHGParams) (*Accessor, []byte, error)
    * [x]   - func (hg *HostGroup) GetAllHostfixes(ctx context.Context) ([]byte, error)
    * [x]   - func (hg *HostGroup) GetComponentsForProductOnHost(ctx context.Context, strHostName, strProductName, strProductVersion string) (*ProductComponents, []byte, error)
    * [x]   - func (hg *HostGroup) GetDomainHosts(ctx context.Context, domain string) ([]byte, error)
    * [x]   - func (hg *HostGroup) GetDomains(ctx context.Context) ([]byte, error)
    * [x]   - func (hg *HostGroup) GetGroupId(ctx context.Context, nParent int64, strName string) (*PxgValInt, []byte, error)
    * [x]   - func (hg *HostGroup) GetGroupInfo(ctx context.Context, nGroupId int64) ([]byte, error)
    * [x]   - func (hg *HostGroup) GetGroupInfoEx(ctx context.Context, params interface{}) ([]byte, error)
    * [x]   - func (hg *HostGroup) GetHostInfo(ctx context.Context, params interface{}) ([]byte, error)
    * [x]   - func (hg *HostGroup) GetHostProducts(ctx context.Context, strHostName string) ([]byte, error)
    * [x]   - func (hg *HostGroup) GetHostTasks(ctx context.Context, hostId string) (*PxgValStr, []byte, error)
    * [x]   - func (hg *HostGroup) GetHostfixesForProductOnHost(ctx context.Context, strHostName, strProductName, strProductVersion string) (*ProductFixes, []byte, error)
    * [x]   - func (hg *HostGroup) GetInstanceStatistics(ctx context.Context, params InstanceStatisticsParams) ([]byte, error)
    * [x]   - func (hg *HostGroup) GetRunTimeInfo(ctx context.Context, params StaticInfoParams) ([]byte, error)
    * [x]   - func (hg *HostGroup) GetStaticInfo(ctx context.Context, params StaticInfoParams) ([]byte, error)
    * [x]   - func (hg *HostGroup) GetSubgroups(ctx context.Context, nGroupId int64, nDepth int64) ([]byte, error)
    * [x]   - func (hg *HostGroup) GroupIdGroups(ctx context.Context) (*PxgValInt, []byte, error)
    * [x]   - func (hg *HostGroup) GroupIdSuper(ctx context.Context) (*PxgValInt, []byte, error)
    * [x]   - func (hg *HostGroup) GroupIdUnassigned(ctx context.Context) (*PxgValInt, []byte, error)
    * [x]   - func (hg *HostGroup) MoveHostsFromGroupToGroup(ctx context.Context, nSrcGroupId int64, nDstGroupId int64) (*WActionGUID, []byte, error)
    * [x]   - func (hg *HostGroup) MoveHostsToGroup(ctx context.Context, params HostsToGroupParams) ([]byte, error)
    * [x]   - func (hg *HostGroup) RemoveGroup(ctx context.Context, nGroup, nFlags int64) (*WActionGUID, []byte, error)
    * [x]   - func (hg *HostGroup) RemoveHost(ctx context.Context, strHostName string)
    * [x]   - func (hg *HostGroup) RemoveHosts(ctx context.Context, params struct{}) ([]byte, error)
    * [x]   - func (hg *HostGroup) ResolveAndMoveToGroup(ctx context.Context, params PInfoRaM) (*KlhstWksResults, []byte, error)
    * [x]   - func (hg *HostGroup) RestartNetworkScanning(ctx context.Context, nType int64) (*PxgRetError, []byte, error)
    * [x]   - func (hg *HostGroup) SS_GetNames(ctx context.Context, params interface{}) ([]byte, error)
    * [x]   - func (hg *HostGroup) SS_Read(ctx context.Context, params interface{}) ([]byte, error)
    * [x]   - func (hg *HostGroup) UpdateGroup(ctx context.Context, params interface{}) (*PxgValStr, []byte, error)
    * [x]   - func (hg *HostGroup) UpdateHost(ctx context.Context, v interface{}) (*Accessor, []byte, error)
    * [x]   - func (hg *HostGroup) ZeroVirusCountForGroup(ctx context.Context, nParent int64) (*WActionGUID, []byte, error)
    * [x]   - func (hg *HostGroup) ZeroVirusCountForHosts(ctx context.Context, params interface{}) (*WActionGUID, []byte, error)

* [x] **interface HostMoveRules** - Modify and acquire move rules to hosts
    * [x]   - func (hs *HostMoveRules) DeleteRule(ctx context.Context, nRule int64) ([]byte, error)
    * [x]   - func (hs *HostMoveRules) GetRule(ctx context.Context, nRule int64) ([]byte, error)
    * [x]   - func (hs *HostMoveRules) GetRules(ctx context.Context) ([]byte, error)

* [x] **interface HostTagsRulesApi** Interface allows to acquire and manage host automatic tagging rules
    * [x]   - func (hs *HostTagsRulesApi) CancelAsyncAction(ctx context.Context, wstrActionGuid string) ([]byte, error)
    * [x]   - func (hs *HostTagsRulesApi) DeleteRule(ctx context.Context, szwTagValue string) ([]byte, error)
    * [x]   - func (hs *HostTagsRulesApi) ExecuteRule(ctx context.Context, szwTagValue string) (*WActionGUID, []byte, error)
    * [x]   - func (hs *HostTagsRulesApi) GetRule(ctx context.Context, szwTagValue string) ([]byte, error)
    * [x]   - func (hs *HostTagsRulesApi) GetRules(ctx context.Context, params HostTagsRulesParams) ([]byte, error)

* [x] **interface HostTasks** - Basic management operations with host tasks
    * [x]   - func (ht *HostTasks) GetNextTask(ctx context.Context, strSrvObjId string) ([]byte, error)
    * [x]   - func (ht *HostTasks) ResetTasksIterator(ctx context.Context, ...) ([]byte, error)

* [x] **interface InvLicenseProducts** - Interface to manage License Management (third party) Functionality
    * [x]   - func (ilp *InvLicenseProducts) GetLicenseProducts(ctx context.Context) ([]byte, error)

* [x] **interface InventoryApi** - Interface for working with Software Inventory subsystem
    * [x]   - func (ia *InventoryApi) DeleteUninstalledApps(ctx context.Context) ([]byte, error)
    * [x]   - func (ia *InventoryApi) GetHostInvPatches(ctx context.Context, szwHostId string, v interface{}) ([]byte, error)
    * [x]   - func (ia *InventoryApi) GetHostInvProducts(ctx context.Context, szwHostId string, v interface{}) ([]byte, error)
    * [x]   - func (ia *InventoryApi) GetInvPatchesList(ctx context.Context, v interface{}) ([]byte, error)
    * [x]   - func (ia *InventoryApi) GetInvProductsList(ctx context.Context, v interface{}) ([]byte, error)
    * [x]   - func (ia *InventoryApi) GetSrvCompetitorIniFileInfoList(ctx context.Context, wstrType string) (*PxgValCIFIL, []byte, error)

* [x] **interface KsnInternal** - Interface for working with KsnProxy subsystem
    * [x]   - func (sd *KsnInternal) CheckKsnConnection(ctx context.Context) (*PxgValBool, []byte, error)
    * [x]   - func (sd *KsnInternal) GetNKsnEula(ctx context.Context, wstrNKsnLoc string) ([]byte, error)
    * [x]   - func (sd *KsnInternal) GetNKsnEulas(ctx context.Context) ([]byte, error)
    * [x]   - func (sd *KsnInternal) GetSettings(ctx context.Context) (*KsnSettings, []byte, error)
    * [x]   - func (sd *KsnInternal) NeedToSendStatistics(ctx context.Context) (*PxgValBool, []byte, error)

* [x] **interface LicenseInfoSync** - Operating with licenses
    * [x]   - func (lis *LicenseInfoSync) IsPCloudKey(ctx context.Context, nProductId int64) ([]byte, error)
    * [x]   - func (lis *LicenseInfoSync) SynchronizeLicInfo2(ctx context.Context) (*PxgValStr, []byte, error)
    * [x]   - func (lis *LicenseInfoSync) TryToUnistallLicense(ctx context.Context, bCurrent bool) ([]byte, error)

* [x] **interface LicenseKeys** - Operating with keys
    * [x]   - func (lk *LicenseKeys) AcquireKeyHosts(ctx context.Context, params AcquireKeyHostsParams) (*HostsKeyIterator, []byte, error)
    * [x]   - func (lk *LicenseKeys) DownloadKeyFiles(ctx context.Context, wstrActivationCode string) bool
    * [x]   - func (lk *LicenseKeys) EnumKeys(ctx context.Context, params EnumKeysParams, v interface{}) ([]byte, error)
    * [x]   - func (lk *LicenseKeys) GetKeyData(ctx context.Context, params KeyDataParams, v interface{}) ([]byte, error)
    * [x]   - func (lk *LicenseKeys) InstallKey(ctx context.Context, pKeyInfo interface{}) bool

* [x] **interface LicensePolicy** - License policy
    * [x]   - func (lp *LicensePolicy) GetFreeLicenseCount(ctx context.Context, nFunctionality int64) ([]byte, error)
    * [x]   - func (lp *LicensePolicy) GetTotalLicenseCount(ctx context.Context, nFunctionality int64) ([]byte, error)
    * [x]   - func (lp *LicensePolicy) IsLimitedMode(ctx context.Context, nFunctionality int64) ([]byte, error)
    * [x]   - func (lp *LicensePolicy) SetLimitedModeTest(ctx context.Context, bLimited bool, eFunctionality int64) ([]byte, error)
    * [x]   - func (lp *LicensePolicy) SetTotalLicenseCountTest(ctx context.Context, eFunctionality, nCount int64) ([]byte, error)
    * [x]   - func (lp *LicensePolicy) SetUsedLicenseCountTest(ctx context.Context, eFunctionality, nCount int64) ([]byte, error)

* [x] **interface Limits** - Interface for working with Limits subsystem
    * [x]   - func (ls *Limits) GetLimits(ctx context.Context, param int64) ([]byte, error)

* [x] **interface ListTags** - Interface allows to acquire and manage tags to various KSC objects
    * [x]   - func (lt *ListTags) AddTag(ctx context.Context) ([]byte, error)
    * [x]   - func (lt *ListTags) DeleteTags2(ctx context.Context) ([]byte, error)
    * [x]   - func (lt *ListTags) GetAllTags(ctx context.Context) ([]byte, error)
    * [x]   - func (lt *ListTags) GetTags(ctx context.Context) ([]byte, error)
    * [x]   - func (lt *ListTags) RenameTag(ctx context.Context) ([]byte, error)
    * [x]   - func (lt *ListTags) SetTags(ctx context.Context) ([]byte, error)

* [x] **interface MigrationData** - Migration of data between KSC On-Premise and KSCHosted
    * [x]   - func (md *MigrationData) AcquireKnownProducts(ctx context.Context) (*KnownProducts, []byte, error)
    * [x]   - func (md *MigrationData) CancelExport(ctx context.Context, wstrActionGuid string) ([]byte, error)

* [x] **interface Multitenancy** - Multitenancy product managing
    * [x]   - func (m *Multitenancy) GetAuthToken(ctx context.Context) ([]byte, error)
    * [x]   - func (m *Multitenancy) GetProducts(ctx context.Context, strProdName, strProdVersion string) ([]byte, error)
    * [x]   - func (m *Multitenancy) GetTenantId(ctx context.Context) ([]byte, error)

* [x] **interface NagCgwHelper** - Nagent CGW (Connection Gateway) API
    * [x]   - func (nc *NagCgwHelper) GetProductComponentLocation(ctx context.Context, szwProduct, szwVersion, szwComponent string) ([]byte, error)

* [x] **interface NagGuiCalls** - Remote host caller
    * [x]   - func (ngc *NagGuiCalls) CallConnectorAsync(ctx context.Context, params interface{}) ([]byte, error)

* [x] **interface NagHstCtl** - Manage nagent on host
    * [x]   - func (nh *NagHstCtl) GetHostRuntimeInfo(ctx context.Context, params interface{}) ([]byte, error)
    * [x]   - func (nh *NagHstCtl) SendProductAction(ctx context.Context, szwProduct, szwVersion string, nProductAction int64) ([]byte, error)
    * [x]   - func (nh *NagHstCtl) SendTaskAction(ctx context.Context, szwProduct, szwVersion, szwTaskStorageId string, ...) ([]byte, error)

* [x] **interface NagRdu** - Remote diagnostics on host
    * [x]   - func (nr *NagRdu) GetCurrentHostState(ctx context.Context) ([]byte, error)
    * [x]   - func (nr *NagRdu) GetUrlToUploadFileToHost(ctx context.Context) ([]byte, error)

* [x] **interface PackagesApi** - Operating with packages
    * [x]   - func (pa *PackagesApi) GetPackages(ctx context.Context) (*ListOfPackages, []byte, error)
    * [x]   - func (pa *PackagesApi) GetPackages2(ctx context.Context) (*ListOfPackages, []byte, error)
    * [x]   - func (pa *PackagesApi) GetUserAgreements(ctx context.Context) (*PxgValStr, []byte, error)

* [x] **interface Policy** - Policies processing
    * [x]   - func (pl *Policy) DeletePolicy(ctx context.Context, nPolicy int64) ([]byte, error)
    * [x]   - func (pl *Policy) ExportPolicy(ctx context.Context, lPolicy int64) (*PxgValStr, []byte, error)
    * [x]   - func (pl *Policy) GetEffectivePoliciesForGroup(ctx context.Context, nGroupId int64) ([]byte, error)
    * [x]   - func (pl *Policy) GetOutbreakPolicies(ctx context.Context) ([]byte, error)
    * [x]   - func (pl *Policy) GetPoliciesForGroup(ctx context.Context, nGroupId int64) ([]byte, error)
    * [x]   - func (pl *Policy) GetPolicyContents(ctx context.Context, nPolicy, nRevisionId, nLifeTime int64) (*PxgValStr, []byte, error)
    * [x]   - func (pl *Policy) GetPolicyData(ctx context.Context, nPolicy int64) (*PxgValPolicy, []byte, error)
    * [x]   - func (pl *Policy) ImportPolicy(ctx context.Context, params PolicyBlob) (*PxgValStr, []byte, error)
    * [x]   - func (pl *Policy) MakePolicyActive(ctx context.Context, nPolicy int64, bActive bool) (*PxgValPolicy, []byte, error)
    * [x]   - func (pl *Policy) MakePolicyRoaming(ctx context.Context, nPolicy int64) (*PxgValPolicy, []byte, error)
    * [x]   - func (pl *Policy) RevertPolicyToRevision(ctx context.Context, nPolicy, nRevisionId int64) (*PxgValPolicy, []byte, error)

* [x] **interface ReportManager** - Reports managing
    * [x]   - func (rm *ReportManager) CancelStatisticsRequest(ctx context.Context, strRequestId string) ([]byte, error)
    * [x]   - func (rm *ReportManager) EnumReportTypes(ctx context.Context) ([]byte, error)
    * [x]   - func (rm *ReportManager) EnumReports(ctx context.Context) ([]byte, error)
    * [x]   - func (rm *ReportManager) ExecuteReportAsync(ctx context.Context, params *interface{}) ([]byte, error)
    * [x]   - func (rm *ReportManager) ExecuteReportAsyncCancel(ctx context.Context, strRequestId string) ([]byte, error)
    * [x]   - func (rm *ReportManager) ExecuteReportAsyncCancelWaitingForSlaves(ctx context.Context, strRequestId string) ([]byte, error)
    * [x]   - func (rm *ReportManager) GetAvailableDashboards(ctx context.Context) (*PxgValArrayOfInt, []byte, error)
    * [x]   - func (rm *ReportManager) GetConstantOutputForReportType(ctx context.Context, lReportType, lXmlTargetType int64) (*PxgValStr, []byte, error)
    * [x]   - func (rm *ReportManager) GetDefaultReportInfo(ctx context.Context, lReportType int64) ([]byte, error)
    * [x]   - func (rm *ReportManager) GetFilterSettings(ctx context.Context, lReportType int64) ([]byte, error)
    * [x]   - func (rm *ReportManager) GetReportCommonData(ctx context.Context, lReportId int64) ([]byte, error)
    * [x]   - func (rm *ReportManager) GetReportIds(ctx context.Context) (*PxgValArrayOfInt, []byte, error)
    * [x]   - func (rm *ReportManager) GetReportInfo(ctx context.Context, lReportId int64) ([]byte, error)
    * [x]   - func (rm *ReportManager) GetReportTypeDetailedInfo(ctx context.Context, lReportType int64) ([]byte, error)
    * [x]   - func (rm *ReportManager) GetStatisticsData(ctx context.Context, strRequestId string) ([]byte, error)
    * [x]   - func (rm *ReportManager) RemoveReport(ctx context.Context, lReportId int64) ([]byte, error)
    * [x]   - func (rm *ReportManager) RequestStatisticsData(ctx context.Context, params interface{}) (*RequestID, []byte, error)

* [x] **interface ScanDiapasons** - Network subnets processing
    * [x]   - func (sd *ScanDiapasons) NotifyDpnsTask(ctx context.Context) ([]byte, error)

* [x] **interface SecurityPolicy** - Allows to manage users and permissions.
    * [x]   - func (sp *SecurityPolicy) AddUser(ctx context.Context, params PUserData) (*PxgValInt, []byte, error)
    * [x]   - func (sp *SecurityPolicy) GetCurrentUserId(ctx context.Context) (*UserInfo, []byte, error)
    * [x]   - func (sp *SecurityPolicy) GetCurrentUserId2(ctx context.Context) (*UserInfoEx, []byte, error)
    * [x]   - func (sp *SecurityPolicy) GetUsers(ctx context.Context, lUserId, lVsId int64) (*UsersInfo, []byte, error)
    * [x]   - func (sp *SecurityPolicy) LoadPerUserData(ctx context.Context) ([]byte, error)
    * [x]   - func (sp *SecurityPolicy) UpdateUser(ctx context.Context, lUserId int, params PUserData) (*PxgValInt, []byte, error)

* [x] **interface SecurityPolicy3** - Allows to manage security groups of internal users. Use srvview SplUserGroupSrvViewName to get information about relationship between users and groups
    * [x]   - func (sp *SecurityPolicy3) AddSecurityGroup(ctx context.Context, params interface{}) ([]byte, error)
    * [x]   - func (sp *SecurityPolicy3) AddUserIntoSecurityGroup(ctx context.Context, lUserId, lGrpId int64) ([]byte, error)
    * [x]   - func (sp *SecurityPolicy3) CloseUserConnections(ctx context.Context, lUserId int64) ([]byte, error)
    * [x]   - func (sp *SecurityPolicy3) DeleteSecurityGroup(ctx context.Context, lGrpId int64) ([]byte, error)
    * [x]   - func (sp *SecurityPolicy3) DeleteUserFromSecurityGroup(ctx context.Context, lUserId, lGrpId int64) ([]byte, error)
    * [x]   - func (sp *SecurityPolicy3) MoveUserIntoOtherSecurityGroup(ctx context.Context, lUserId, lGrpIdFrom, lGrpIdTo int64) ([]byte, error)
    * [x]   - func (sp *SecurityPolicy3) UpdateSecurityGroup(ctx context.Context, params interface{}) ([]byte, error)

* [x] **interface ServerHierarchy** - Server hierarchy management interface
    * [x]   - func (sh *ServerHierarchy) DelServer(ctx context.Context, lServer int64) ([]byte, error)
    * [x]   - func (sh *ServerHierarchy) GetChildServers(ctx context.Context, nGroupId int64) ([]byte, error)
    * [x]   - func (sh *ServerHierarchy) GetServerInfo(ctx context.Context, params interface{}) ([]byte, error)

* [x] **interface ServerTransportSettings** - Server transport settings proxy class
    * [x]   - func (sts *ServerTransportSettings) CheckDefaultCertificateExists(ctx context.Context, szwCertType string) (*PxgValBool, []byte, error)
    * [x]   - func (sts *ServerTransportSettings) GetCurrentConnectionSettings(ctx context.Context, szwCertType string) (*CurrentConnectionSettings, []byte, error)
    * [x]   - func (sts *ServerTransportSettings) GetCustomSrvCertificateInfo(ctx context.Context, szwCertType string) ([]byte, error)
    * [x]   - func (sts *ServerTransportSettings) GetDefaultConnectionSettings(ctx context.Context, szwCertType string) (*CurrentConnectionSettings, []byte, error)
    * [x]   - func (sts *ServerTransportSettings) GetNumberOfManagedDevicesAgentless(ctx context.Context) ([]byte, error)
    * [x]   - func (sts *ServerTransportSettings) GetNumberOfManagedDevicesKSM(ctx context.Context) ([]byte, error)
    * [x]   - func (sts *ServerTransportSettings) IsFeatureActive(ctx context.Context, szwCertType string) (*PxgValBool, []byte, error)
    * [x]   - func (sts *ServerTransportSettings) ResetCstmReserveCertificate(ctx context.Context, szwCertType string) ([]byte, error)
    * [x]   - func (sts *ServerTransportSettings) ResetDefaultReserveCertificate(ctx context.Context, szwCertType string) ([]byte, error)
    * [x]   - func (sts *ServerTransportSettings) SetFeatureActive(ctx context.Context, szwCertType string, bFeatureActive bool) (*PxgValBool, []byte, error)

* [x] **interface Session** - Session management interface
    * [x]   - func (s *Session) CreateToken(ctx context.Context) (*PxgValStr, []byte, error)
    * [x]   - func (s *Session) EndSession(ctx context.Context) ([]byte, error)
    * [x]   - func (s *Session) Ping(ctx context.Context) ([]byte, error)
    * [x]   - func (s *Session) StartSession(ctx context.Context) (*PxgValStr, []byte, error)

* [x] **interface SmsQueue** - Manage SMS message queue
    * [x]   - func (sq *SmsQueue) Cancel(ctx context.Context, params SQCParams) ([]byte, error)
    * [x]   - func (sq *SmsQueue) Clear(ctx context.Context) ([]byte, error)
    * [x]   - func (sq *SmsQueue) Enqueue(ctx context.Context, params SQParams) ([]byte, error)

* [x] **interface SmsSenders** - Configure mobile devices as SMS senders
    * [x]   - func (ss *SmsSenders) HasAllowedSenders(ctx context.Context) (*PxgValBool, []byte, error)

* [x] **interface SrvSsRevision** - Access to virtual server settings storage revisions
    * [x]   - func (ssr *SrvSsRevision) SsRevision_Close(ctx context.Context, szwType string) ([]byte, error)
    * [x]   - func (ssr *SrvSsRevision) SsRevision_Open(ctx context.Context, nVServer, nRevision int64, szwType string) ([]byte, error)

* [x] **interface SrvView** - Interface to get plain-queries from SC-server
    * [x]   - func (sv *SrvView) GetRecordCount(ctx context.Context, wstrIteratorId string) (*PxgValInt, []byte, error)
    * [x]   - func (sv *SrvView) GetRecordRange(ctx context.Context, params *RecordRangeParams) ([]byte, error)
    * [x]   - func (sv *SrvView) ReleaseIterator(ctx context.Context, wstrIteratorId string) ([]byte, error)
    * [x]   - func (sv *SrvView) ResetIterator(ctx context.Context, params *SrvViewParams) (*WstrIteratorID, []byte, error)

* [x] **interface SsContents** - Access to settings storage
    * [x]   - func (sc *SsContents) SS_GetNames(ctx context.Context, wstrID, wstrProduct, wstrVersion string) ([]byte, error)
    * [x]   - func (sc *SsContents) Ss_Apply(ctx context.Context, wstrID string) ([]byte, error)
    * [x]   - func (sc *SsContents) Ss_CreateSection(ctx context.Context, wstrID, wstrProduct, wstrVersion, wstrSection string) ([]byte, error)
    * [x]   - func (sc *SsContents) Ss_DeleteSection(ctx context.Context, wstrID, wstrProduct, wstrVersion, wstrSection string) ([]byte, error)
    * [x]   - func (sc *SsContents) Ss_Read(ctx context.Context, wstrID, wstrProduct, wstrVersion, wstrSection string) ([]byte, error)
    * [x]   - func (sc *SsContents) Ss_Release(ctx context.Context, wstrID string) ([]byte, error)

* [x] **interface Tasks** - Group tasks
    * [x]   - func (ts *Tasks) CancelTask(ctx context.Context, strTask string) ([]byte, error)
    * [x]   - func (ts *Tasks) DeleteTask(ctx context.Context, strTask string) ([]byte, error)
    * [x]   - func (ts *Tasks) GetAllTasksOfHost(ctx context.Context, strDomainName, strHostName string) (*PxgValArrayOfString, []byte, error)
    * [x]   - func (ts *Tasks) GetHostStatusRecordRange(ctx context.Context, strHostIteratorId string, nStart, nEnd int64) ([]byte, error)
    * [x]   - func (ts *Tasks) GetHostStatusRecordsCount(ctx context.Context, strHostIteratorId string) (*PxgValInt, []byte, error)
    * [x]   - func (ts *Tasks) GetNextHostStatus(ctx context.Context, strTaskIteratorId string) ([]byte, error)
    * [x]   - func (ts *Tasks) GetNextTask(ctx context.Context, strTaskIteratorId string) ([]byte, error)
    * [x]   - func (ts *Tasks) GetTask(ctx context.Context, strTask string) (*TaskData, []byte, error)
    * [x]   - func (ts *Tasks) GetTaskData(ctx context.Context, strTask string, tsk interface{}) ([]byte, error)
    * [x]   - func (ts *Tasks) GetTaskGroup(ctx context.Context, strTaskId string) (*PxgValInt, []byte, error)
    * [x]   - func (ts *Tasks) GetTaskHistory(ctx context.Context, params interface{}) (*StrIteratorId, []byte, error)
    * [x]   - func (ts *Tasks) GetTaskStartEvent(ctx context.Context, strTask string) ([]byte, error)
    * [x]   - func (ts *Tasks) GetTaskStatistics(ctx context.Context, strTask string) (*TaskStatistics, []byte, error)
    * [x]   - func (ts *Tasks) ProtectPassword(ctx context.Context, strPassword string) ([]byte, error)
    * [x]   - func (ts *Tasks) ReleaseHostStatusIterator(ctx context.Context, strHostIteratorId string) ([]byte, error)
    * [x]   - func (ts *Tasks) ReleaseTasksIterator(ctx context.Context, strTaskIteratorId string) ([]byte, error)
    * [x]   - func (ts *Tasks) ResetHostIteratorForTaskStatus(ctx context.Context, params HostIteratorForTaskParams) ([]byte, error)
    * [x]   - func (ts *Tasks) ResetHostIteratorForTaskStatusEx(ctx context.Context, params HostIteratorForTaskParamsEx) (*StrHostIteratorId, []byte, error)
    * [x]   - func (ts *Tasks) ResetTasksIterator(ctx context.Context, params TasksIteratorParams) ([]byte, error)
    * [x]   - func (ts *Tasks) ResolveTaskId(ctx context.Context, strPrtsTaskId string) ([]byte, error)
    * [x]   - func (ts *Tasks) ResumeTask(ctx context.Context, strTask string) ([]byte, error)
    * [x]   - func (ts *Tasks) RunTask(ctx context.Context, strTask string) ([]byte, error)
    * [x]   - func (ts *Tasks) SuspendTask(ctx context.Context, strTask string) ([]byte, error)

* [x] **interface TrafficManager** - Traffic manager interface
    * [x]   - func (tm *TrafficManager) AddRestriction(ctx context.Context, params interface{}) (*PxgValInt, []byte, error)
    * [x]   - func (tm *TrafficManager) DeleteRestriction(ctx context.Context, nRestrictionId int64) ([]byte, error)
    * [x]   - func (tm *TrafficManager) GetRestrictions(ctx context.Context) ([]byte, error)
    * [x]   - func (tm *TrafficManager) UpdateRestriction(ctx context.Context, params interface{}) ([]byte, error)

* [x] **interface UaControl** - Update agents and Connection gateways management
    * [x]   - func (uc *UaControl) GetAssignUasAutomatically(ctx context.Context) (*PxgValBool, []byte, error)
    * [x]   - func (uc *UaControl) GetDefaultUpdateAgentRegistrationInfo(ctx context.Context) ([]byte, error)
    * [x]   - func (uc *UaControl) GetUpdateAgentsDisplayInfoForHost(ctx context.Context, wstrHostId string) ([]byte, error)
    * [x]   - func (uc *UaControl) GetUpdateAgentsList(ctx context.Context) ([]byte, error)
    * [x]   - func (uc *UaControl) SetAssignUasAutomatically(ctx context.Context, bEnabled bool) ([]byte, error)
    * [x]   - func (uc *UaControl) UnregisterUpdateAgent(ctx context.Context, wstrUaHostId string) ([]byte, error)

* [x] **interface Updates** - Updates processing
    * [x]   - func (upd *Updates) GetAvailableUpdatesInfo(ctx context.Context, strLocalization string) ([]byte, error)
    * [x]   - func (upd *Updates) GetUpdatesInfo(ctx context.Context) ([]byte, error)
    * [x]   - func (upd *Updates) RemoveUpdates(ctx context.Context) (*RequestID, []byte, error)
    * [x]   - func (upd *Updates) RemoveUpdatesCancel(ctx context.Context, strRequestId string) (*RequestID, []byte, error)

* [x] **interface UserDevicesApi** - Interface to unified mobile device management
    * [x]   - func (uda *UserDevicesApi) DeleteCommand(ctx context.Context, c_wstrCommandGuid string, bForced bool) ([]byte, error)
    * [x]   - func (uda *UserDevicesApi) DeleteDevice(ctx context.Context, lDeviceId int64) ([]byte, error)
    * [x]   - func (uda *UserDevicesApi) DeleteEnrollmentPackage(ctx context.Context, lEnrPkgId int64) ([]byte, error)
    * [x]   - func (uda *UserDevicesApi) GenerateQRCode(ctx context.Context, strInputData string, lQRCodeSize, lImageFormat int64) ([]byte, error)
    * [x]   - func (uda *UserDevicesApi) GetCommands(ctx context.Context, lDeviceId int64) ([]byte, error)
    * [x]   - func (uda *UserDevicesApi) GetCommandsLibrary(ctx context.Context) ([]byte, error)
    * [x]   - func (uda *UserDevicesApi) GetDevice(ctx context.Context, lDeviceId int64) ([]byte, error)
    * [x]   - func (uda *UserDevicesApi) GetEnrollmentPackage(ctx context.Context, llEnrPkgId int64) ([]byte, error)
    * [x]   - func (uda *UserDevicesApi) GetEnrollmentPackageFileData(ctx context.Context, c_wstrPackageId, c_wstrPackageFileType string, ...) ([]byte, error)
    * [x]   - func (uda *UserDevicesApi) GetEnrollmentPackageFileInfo(ctx context.Context, ...) ([]byte, error)
    * [x]   - func (uda *UserDevicesApi) GetJournalCommandResult(ctx context.Context, llJrnlId int64) ([]byte, error)
    * [x]   - func (uda *UserDevicesApi) GetJournalRecords(ctx context.Context, lDeviceId int64) ([]byte, error)
    * [x]   - func (uda *UserDevicesApi) GetJournalRecords2(ctx context.Context, lDeviceId int64) ([]byte, error)
    * [x]   - func (uda *UserDevicesApi) GetLatestDeviceActivityDate(ctx context.Context, lDeviceId int64) ([]byte, error)
    * [x]   - func (uda *UserDevicesApi) GetMobileAgentSettingStorageData(ctx context.Context, lDeviceId int64, c_wstrSectionName string) ([]byte, error)
    * [x]   - func (uda *UserDevicesApi) GetMultitenancyServerSettings(ctx context.Context, c_wstrMtncServerId string) ([]byte, error)
    * [x]   - func (uda *UserDevicesApi) GetMultitenancyServersInfo(ctx context.Context, nProtocolIds int64) ([]byte, error)
    * [x]   - func (uda *UserDevicesApi) GetSafeBrowserAutoinstallFlag(ctx context.Context) (*PxgValBool, []byte, error)
    * [x]   - func (uda *UserDevicesApi) GlueDevices(ctx context.Context, lDevice1Id, lDevice2Id int64) ([]byte, error)
    * [x]   - func (uda *UserDevicesApi) RecallCommand(ctx context.Context, c_wstrCommandGuid string) ([]byte, error)
    * [x]   - func (uda *UserDevicesApi) SetSafeBrowserAutoinstallFlag(ctx context.Context, bInstall bool) ([]byte, error)
    * [x]   - func (uda *UserDevicesApi) SspLoginAllowed(ctx context.Context) ([]byte, error)

* [x] **interface VServers** - Virtual servers processing
    * [x]   - func (vs *VServers) AddVServerInfo(ctx context.Context, strDisplayName string, lParentGroup int64) (*VServer, []byte, error)
    * [x]   - func (vs *VServers) DelVServer(ctx context.Context, lVServer int64) ([]byte, error)
    * [x]   - func (vs *VServers) GetPermissions(ctx context.Context, lVServer int64) ([]byte, error)
    * [x]   - func (vs *VServers) GetVServerInfo(ctx context.Context, lVServer int64) ([]byte, error)
    * [x]   - func (vs *VServers) GetVServers(ctx context.Context, lParentGroup int64) ([]byte, error)
    * [x]   - func (vs *VServers) MoveVServer(ctx context.Context, lVServer int64, lNewParentGroup int64) (*WActionGUID, []byte, error)
    * [x]   - func (vs *VServers) RecallCertAndCloseConnections(ctx context.Context, lVServer int64) ([]byte, error)
    * [x]   - func (vs *VServers) SetPermissions(ctx context.Context, lVServer int64, params interface{}, bProtection bool) ([]byte, error)
    * [x]   - func (vs *VServers) UpdateVServerInfo(ctx context.Context, lVServer int64, params interface{}) ([]byte, error)

* [x] **interface VServers2** - Virtual servers processing
    * [x]   - func (vs *VServers2) GetVServerStatistic(ctx context.Context, lVsId int) (*VServerStatistic, []byte, error)

* [x] **interface WolSender** - Wake-On-LAN signal sender
    * [x]   - func (ws *WolSender) SendWolSignal(ctx context.Context, szwHostId string) error


NOTE: Using the [context](https://godoc.org/context) package, one can easily
pass cancelation signals and deadlines to various services of the client for
handling a request. In case there is no context available, then `context.Background()`
can be used as a starting point.


## License ##

This library is distributed under the  GNU GENERAL PUBLIC LICENSE Version 3 found in the [LICENSE](./LICENSE)
file.