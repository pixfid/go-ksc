# go-ksc #
go-ksc is a Go client library for accessing the KSC (Kaspersky Security Center) Open API.

## Badges
[![Go Report Card](https://goreportcard.com/badge/github.com/pixfid/go-ksc)](https://goreportcard.com/report/github.com/pixfid/go-ksc)
[![License](https://img.shields.io/github/license/pixfid/go-ksc?style=flat-square)](/LICENSE)
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

    	    //user login name
    		UserName: "login",

    		//password
    		Password: "password",

            //VServerName: "virtual_server_name", for login on virtual server.
    		Server: fmt.Sprintf(`https://%s:%s`, "ip", "port"),

    		//true using XKscSession tokens (false on default, session token expired time 3 minutes)
    		XKscSession: false,
    	}

        //Construct a new KSC client
    	client := kaspersky.New(cfg)

    	//Auth on KSC server
    	client.KSCAuth(ctx)

        //Get List of Windows domain in the network.
        raw,_ := client.HostGroup.GetDomains(context.Background())
        println(string(raw))
}
```
# Examples:

###### Find online hosts:

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

###### Get installed products on host by HostId:

```go
products, err := client.HostGroup.GetHostProducts(ctx, "8910f900-3807-4b97-8a97-d49e73ec5ab1")
```

###### Response:

```json
{
  "PxgRetVal" : {
    "1103" : {
      "type" : "params",
      "value" : {
        "1.0.0.0" : {
          "type" : "params",
          "value" : {
            "BaseRecords" : 0,
            "LastUpdateTime" : {
              "type" : "datetime",
              "value" : "2019-10-26T23:17:09Z"
            }
          }
        }
      }
    },
    "kesl" : {
      "type" : "params",
      "value" : {
        "10.1.1.0" : {
          "type" : "params",
          "value" : {
            "BaseDate" : {
              "type" : "datetime",
              "value" : "2020-05-07T23:18:00Z"
            },
            "BaseInstallDate" : {
              "type" : "datetime",
              "value" : "2020-05-08T08:59:53Z"
            },
            "BaseRecords" : 14791566
          }
        }
      }
    }
  }
}
```

###### Get Lists tasks on Host:

```go
tastList, raw, err := client.Tasks.GetAllTasksOfHost(ctx, "", "c2b22f83-307c-45aa-8533-5ffffbcc6bf1")
```

###### Response:

```json
{
  "PxgRetVal" : [
    "101",
    "119",
    "120",
    "121",
    "122",
    "123",
    "149",
    "164",
    "192"
  ]
}
```

###### Find srvview data by filter string. A removable device's collection.

```go
	srvVParams := &kaspersky.SrvViewParams{
		WstrViewName:      "HWInvStorageSrvViewName", //Hardware inventory storage view
		WstrFilter:        `(&(Type = 4))`, //Type = 4 (Removable devices)
		VecFieldsToReturn: []string{"Id", "Name", "SerialNumber"}, //Return Fields
		VecFieldsToOrder: []kaspersky.FieldsToOrder{ // Sort by Id field (Ascending)
			{
				Type: "params",
				OrderValue: kaspersky.OrderValue{
					Name: "Id",
					Asc:  true,
				},
			},
		},
		PParams: &kaspersky.ESrvViewParams{
			TopN: 100, //First 100 records
		},
		LifetimeSEC: 1300, //set lifetime in seconds
	}

	iterator, _, _ := client.SrvView.ResetIterator(ctx, srvVParams)
	c, _, _ := client.SrvView.GetRecordCount(ctx, iterator.WstrIteratorID)

	recordRange := &kaspersky.RecordRangeParams{
		WstrIteratorID: iterator.WstrIteratorID, //iterator string
		NStart:         0, //start number
		NEnd:           c.Int, //count
	}

	raw, _ := client.SrvView.GetRecordRange(ctx, recordRange) //[]byte json data example below
	_, _ = client.SrvView.ReleaseIterator(ctx, *iterator.WstrIteratorID) //release iterator set on server
```

###### Response:

```json
{
  "pRecords" : {
    "KLCSP_ITERATOR_ARRAY" : [
      {
        "type" : "params",
        "value" : {
          "Id" : 38236,
          "Name" : "USB DISK 2.0 USB Device",
          "SerialNumber" : "USBSTOR\\DISK&VEN_&PROD_USB_DISK_2.0&REV_PMAP\\9000694205A94058&0"
        }
      },
      {
        "type" : "params",
        "value" : {
          "Id" : 38237,
          "Name" : "Kingston DT Rubber 3.0 USB Device",
          "SerialNumber" : "USBSTOR\\DISK&VEN_KINGSTON&PROD_DT_RUBBER_3.0&REV_PMAP\\001A92053B6ABD7131341955&0"
        }
      }
    ]
  }
}
```

#### Table of Classes:

| Status  | Implement in go-ksc            | KSC Classes                 | Description
|---------|--------------------------------|-----------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|
|    ✔    | AdfsSso.go                     | AdfsSso                     | Interface for working with ADFS SSO                                                                                                                     |
|    ✔    | AdHosts.go                     | AdHosts                     | Scanned active directory OU structure                                                                                                                   |
|    ✔    | AdmServerSettings.go           | AdmServerSettings           | AdmServerSettings interface                                                                                                                             |
|    ✔    | AdSecManager.go                | AdSecManager                | Adaptive Security managing                                                                                                                              |
|    ✔    | AKPatches.go                   | AKPatches                   | Interface to manage system of autoupdating by patch.exe patches                                                                                         |
|    ✔    | AppCtrlApi.go                  | AppCtrlApi                  | Interface to get info about execution files                                                                                                             |
|    ✔    | AsyncActionStateChecker.go     | AsyncActionStateChecker     | Interface to monitor state of async action                                                                                                              |
|    ✔    | CertPoolCtrl.go                | CertPoolCtrl                | Interface to manage the pool of certificates used by the Kaspersky Security Center Server                                                               |
|    ✔    | CertPoolCtrl2.go               | CertPoolCtrl2               | 2nd interface to manage the pool of certificates used by the Kaspersky Security Center Server                                                           |
|    ✔    | CgwHelper.go                   | CgwHelper                   | CgwHelper (Connection Gateway) helper proxy class                                                                                                       |
|    ✔    | ChunkAccessor.go               | ChunkAccessor               | Working with host result-set                                                                                                                            |
|    ✔    | CloudAccess.go                 | CloudAccess                 | Interface to check access of public clouds                                                                                                              |
|    ✔    | ConEvents.go                   | ConEvents                   | Interface to server events                                                                                                                              |
|    ✔    | DatabaseInfo.go                | DatabaseInfo                | Database processing                                                                                                                                     |
|    ✔    | DataProtectionApi.go           | DataProtectionApi           | Allows to protect sensitive data in policies, tasks, and/or on specified host                                                                           |
|    ✔    | DpeKeyService.go               | DpeKeyService               | Interface for working with encrypted devices                                                                                                            |
|    ✔    | EventNotificationProperties.go | EventNotificationProperties | Notification properties                                                                                                                                 |
|    ✔    | EventNotificationsApi.go       | EventNotificationsApi       | Publish event                                                                                                                                           |
|    ✔    | EventProcessing.go             | EventProcessing             | Interface implements the functionality for viewing and deleting events                                                                                  |
|    ✔    | EventProcessingFactory.go      | EventProcessingFactory      | Interface to create event processing iterators                                                                                                          |
|    ✔    | ExtAud.go                      | ExtAud                      | Interface for working with ExtAudit subsystem                                                                                                           |
|    ✔    | FileCategorizer2.go            | FileCategorizer2            | Interface for working with FileCategorizer subsystem                                                                                                    |
|    ✔    | FilesAcceptor.go               | FilesAcceptor               | Upload files to server                                                                                                                                  |
|    ✔    | GatewayConnection.go           | GatewayConnection           | Interface for creating gateway connections                                                                                                              |
|    ✔    | GroupSync.go                   | GroupSync                   | Access to group synchronization objects                                                                                                                 |
|    ✔    | GroupSyncIterator.go           | GroupSyncIterator           | Access to the group synchronization forward iterator for the result-set                                                                                 |
|    ✔    | GroupTaskControlApi.go         | GroupTaskControlApi         | Interface to perform some management actions over group tasks                                                                                           |
|    ✔    | HostGroup.go                   | HostGroup                   | Hosts and management groups processing                                                                                                                  |
|    ✔    | HostMoveRules.go               | HostMoveRules               | Modify and acquire move rules to hosts                                                                                                                  |
|    ✔    | HostTagsApi.go                 | HostTagsApi                 | Interface allows to acquire and manage tags for hosts. It is additional interface for common ListTags                                                   |
|    ✔    | HostTagsRulesApi.go            | HostTagsRulesApi            | Interface allows to acquire and manage host automatic tagging rules                                                                                     |
|    ✔    | HostTasks.go                   | HostTasks                   | Basic management operations with host tasks                                                                                                             |
|    ✔    | HstAccessControl.go            | HstAccessControl            | Security policy Allows to specify permissions for administrration groups and non-group objects                                                          |
|    ✔    | HWInvStorage.go                | HWInvStorage                | Interface for working with Hardware storage subsystem                                                                                                   |
|    ✔    | InventoryApi.go                | InventoryApi                | Interface for working with Software Inventory subsystem                                                                                                 |
|    ✔    | InvLicenseProducts.go          | InvLicenseProducts          | Interface to manage License Management (third party) Functionality                                                                                      |
|    ✔    | IWebSrvSettings.go             | IWebSrvSettings             | Web server settings proxy class                                                                                                                         |
|    ✔    | IWebUsersSrv.go                | IWebUsersSrv                | Send an email to multiple recipients.                                                                                                                   |
|    ✔    | IWebUsersSrv2.go               | IWebUsersSrv2               | Operating with emails from GUI                                                                                                                          |
|    ✔    | KeyService.go                  | KeyService                  | Interface for working with KeyService subsystem                                                                                                         |
|    ✔    | KeyService2.go                 | KeyService2                 | Additional interface for working with KeyService subsystem                                                                                              |
|    ✔    | KillChain.go                   | KillChain                   | KillChain info from host                                                                                                                                |
|    ✔    | KLEVerControl.go               | KLEVerControl               | Controls the possibility to download and automatically create installation packages                                                                     |
|    ✔    | KsnInternal.go                 | KsnInternal                 | Interface for working with KsnProxy subsystem                                                                                                           |
|    ✔    | LicenseInfoSync.go             | LicenseInfoSync             | Operating with licenses                                                                                                                                 |
|    ✔    | LicenseKeys.go                 | LicenseKeys                 | Operating with keys                                                                                                                                     |
|    ✔    | LicensePolicy.go               | LicensePolicy               | License policy                                                                                                                                          |
|    ✔    | Limits.go                      | Limits                      | Interface for working with Limits subsystem                                                                                                             |
|    ✔    | ListTags.go                    | ListTags                    | Interface allows to acquire and manage tags to various KSC objects                                                                                      |
|    ✔    | MigrationData.go               | MigrationData               | Migration of data between KSC On-Premise and KSCHosted                                                                                                  |
|    ✔    | Multitenancy.go                | Multitenancy                | Multitenancy product managing                                                                                                                           |
|    ✔    | NagCgwHelper.go                | NagCgwHelper                | Nagent CGW (Connection Gateway) API                                                                                                                     |
|    ✔    | NagGuiCalls.go                 | NagGuiCalls                 | Remote host caller                                                                                                                                      |
|    ✔    | NagHstCtl.go                   | NagHstCtl                   | Manage nagent on host                                                                                                                                   |
|    ✔    | NagNetworkListApi.go           | NagNetworkListApi           | Nagent OpenAPI to work with network lists                                                                                                               |
|    ✔    | NagRdu.go                      | NagRdu                      | Remote diagnostics on host                                                                                                                              |
|    ✔    | NagRemoteScreen.go             | NagRemoteScreen             | Interface for remote screen session management                                                                                                          |
|    ✔    | NlaDefinedNetworks.go          | NlaDefinedNetworks          | Network location awareness (NLA) defined networks. Used as a scope for Update agents. Each NLA-defined network is defined by list of NLA locations      |
|    ✔    | OsVersion.go                   | OsVersion                   | Operating systems dictionary access                                                                                                                     |
|    ✔    | PackagesApi.go                 | PackagesApi                 | Operating with packages                                                                                                                                 |
|    ✔    | PatchParameters.go             | PatchParameters             | Patch parameters processing                                                                                                                             |
|    ✔    | PLCDevApi.go                   | PLCDevApi                   | Interface allows to acquire and manage PLC devices registry                                                                                             |
|    ✔    | Policy.go                      | Policy                      | Policies processing                                                                                                                                     |
|    ✔    | PolicyProfiles.go              | PolicyProfiles              | Policy profiles processing                                                                                                                              |
|    ✔    | QBTNetworkListApi.go           | QBTNetworkListApi           | Interface to working with Quarantine, Backup and TIF network lists                                                                                      |
|    ✔    | QueriesStorage.go              | QueriesStorage              | QueriesStorage interface                                                                                                                                |
|    ✔    | ReportManager.go               | ReportManager               | Reports managing                                                                                                                                        |
|    ✔    | RetrFiles.go                   | RetrFiles                   | Class provides means to get retranslated files info                                                                                                     |
|    ✔    | ScanDiapasons.go               | ScanDiapasons               | Network subnets processing                                                                                                                              |
|    ✔    | SecurityPolicy.go              | SecurityPolicy              | Allows to manage users and permissions.                                                                                                                 |
|    ✔    | SecurityPolicy3.go             | SecurityPolicy3             | Allows to manage security groups of internal users. Use srvview SplUserGroupSrvViewName to get information about relationship between users and groups  |
|    ✔    | ServerHierarchy.go             | ServerHierarchy             | Server hierarchy management interface                                                                                                                   |
|    ✔    | ServerTransportSettings.go     | ServerTransportSettings     | Server transport settings proxy class                                                                                                                   |
|    ✔    | Session.go                     | Session                     | Session management interface                                                                                                                            |
|    ✔    | SmsQueue.go                    | SmsQueue                    | Manage SMS message queue                                                                                                                                |
|    ✔    | SmsSenders.go                  | SmsSenders                  | Configure mobile devices as SMS senders                                                                                                                 |
|    ✔    | SrvCloud.go                    | SrvCloud                    | Interface to acquire info about public clouds                                                                                                           |
|    ✔    | SrvSsRevision.go               | SrvSsRevision               | Access to virtual server settings storage revisions                                                                                                     |
|    ✔    | SrvView.go                     | SrvView                     | Interface to get plain-queries from SC-server                                                                                                           |
|    ✔    | SsContents.go                  | SsContents                  | Access to settings storage                                                                                                                              |
|    ✔    | SubnetMasks.go                 | SubnetMasks                 | Subnets provider                                                                                                                                        |
|    ✔    | Tasks.go                       | Tasks                       | Group tasks                                                                                                                                             |
|    ✔    | TrafficManager.go              | TrafficManager              | Traffic manager interface                                                                                                                               |
|    ✔    | UaControl.go                   | UaControl                   | Update agents and Connection gateways management                                                                                                        |
|    ✔    | Updates.go                     | Updates                     | Updates processing                                                                                                                                      |
|    ✔    | UpdComps.go                    | UpdComps                    | Class provides means to manage updatable components (bases)                                                                                             |
|    ✔    | UserDevicesApi.go              | UserDevicesApi              | Interface to unified mobile device management                                                                                                           |
|    ✔    | VapmControlApi.go              | VapmControlApi              | VAPM                                                                                                                                                    |
|    ✔    | VServers.go                    | VServers                    | Virtual servers processing                                                                                                                              |
|    ✔    | VServers2.go                   | VServers2                   | Virtual servers processing                                                                                                                              |
|    ✔    | WolSender.go                   | WolSender                   | Wake-On-LAN signal sender                                                                                                                               |

#### TODO
* [x] Implement all classes
* [x] Implements all Methods
* [ ] Write Tests
* [ ] Examples
* [ ] Write Documentation

#### NOTE
Using the [context](https://godoc.org/context) package, one can easily
pass cancellation signals and deadlines to various services of the client for
handling a request. In case there is no context available, then `context.Background()`
can be used as a starting point.


#### License

This library is distributed under the  MIT LICENSE found in the [LICENSE](./LICENSE)
file.
