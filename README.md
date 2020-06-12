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
	c, _, _ := client.SrvView.GetRecordCount(ctx, *iterator.WstrIteratorID)

	recordRange := &kaspersky.RecordRangeParams{
		WstrIteratorID: *iterator.WstrIteratorID, //iterator string
		NStart:         0, //start number
		NEnd:           *c.Int, //count
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

##TODO##
| Status   | Implement in go-ksc            | KSC Classes                 | Description
|----------|--------------------------------|-----------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------|
&#9745;    | AdfsSso.go                     | AdfsSso                     | Interface for working with ADFS SSO                                                                                                                     |
&#9745;    | AdHosts.go                     | AdHosts                     | Scanned active directory OU structure                                                                                                                   |
&#9745;    | AdmServerSettings.go           | AdmServerSettings           | AdmServerSettings interface                                                                                                                             |
&#9745;    | AdSecManager.go                | AdSecManager                | Adaptive Security managing                                                                                                                              |
&#9745;    | AKPatches.go                   | AKPatches                   | Interface to manage system of autoupdating by patch.exe patches                                                                                         |
&#9745;    | AppCtrlApi.go                  | AppCtrlApi                  | Interface to get info about execution files                                                                                                             |
&#9745;    | AsyncActionStateChecker.go     | AsyncActionStateChecker     | Interface to monitor state of async action                                                                                                              |
&#9745;    | CertPoolCtrl.go                | CertPoolCtrl                | Interface to manage the pool of certificates used by the Kaspersky Security Center Server                                                               |
&#9745;    | CertPoolCtrl2.go               | CertPoolCtrl2               | 2nd interface to manage the pool of certificates used by the Kaspersky Security Center Server                                                           |
&#9745;    | CgwHelper.go                   | CgwHelper                   | CgwHelper (Connection Gateway) helper proxy class                                                                                                       |
&#9745;    | ChunkAccessor.go               | ChunkAccessor               | Working with host result-set                                                                                                                            |
&#9745;    | CloudAccess.go                 | CloudAccess                 | Interface to check access of public clouds                                                                                                              |
&#9745;    | ConEvents.go                   | ConEvents                   | Interface to server events                                                                                                                              |
&#9745;    | DatabaseInfo.go                | DatabaseInfo                | Database processing                                                                                                                                     |
&#9745;    | DataProtectionApi.go           | DataProtectionApi           | Allows to protect sensitive data in policies, tasks, and/or on specified host                                                                           |
&#9745;    | DpeKeyService.go               | DpeKeyService               | Interface for working with encrypted devices                                                                                                            |
&#9745;    | EventNotificationProperties.go | EventNotificationProperties | Notification properties                                                                                                                                 |
&#9745;    | EventNotificationsApi.go       | EventNotificationsApi       | Publish event                                                                                                                                           |
&#9745;    | EventProcessing.go             | EventProcessing             | Interface implements the functionality for viewing and deleting events                                                                                  |
&#9745;    | EventProcessingFactory.go      | EventProcessingFactory      | Interface to create event processing iterators                                                                                                          |
&#9745;    | ExtAud.go                      | ExtAud                      | Interface for working with ExtAudit subsystem                                                                                                           |
&#9745;    | FileCategorizer2.go            | FileCategorizer2            | Interface for working with FileCategorizer subsystem                                                                                                    |
&#9745;    | FilesAcceptor.go               | FilesAcceptor               | Upload files to server                                                                                                                                  |
&#9745;    | GatewayConnection.go           | GatewayConnection           | Interface for creating gateway connections                                                                                                              |
&#9745;    | GroupSync.go                   | GroupSync                   | Access to group synchronization objects                                                                                                                 |
&#9745;    | GroupSyncIterator.go           | GroupSyncIterator           | Access to the group synchronization forward iterator for the result-set                                                                                 |
&#9745;    | GroupTaskControlApi.go         | GroupTaskControlApi         | Interface to perform some management actions over group tasks                                                                                           |
&#9745;    | HostGroup.go                   | HostGroup                   | Hosts and management groups processing                                                                                                                  |
&#9745;    | HostMoveRules.go               | HostMoveRules               | Modify and acquire move rules to hosts                                                                                                                  |
&#9745;    | HostTagsApi.go                 | HostTagsApi                 | Interface allows to acquire and manage tags for hosts. It is additional interface for common ListTags                                                   |
&#9745;    | HostTagsRulesApi.go            | HostTagsRulesApi            | Interface allows to acquire and manage host automatic tagging rules                                                                                     |
&#9745;    | HostTasks.go                   | HostTasks                   | Basic management operations with host tasks                                                                                                             |
&#9745;    | HstAccessControl.go            | HstAccessControl            | Security policy Allows to specify permissions for administrration groups and non-group objects                                                          |
&#9745;    | HWInvStorage.go                | HWInvStorage                | Interface for working with Hardware storage subsystem                                                                                                   |
&#9745;    | InventoryApi.go                | InventoryApi                | Interface for working with Software Inventory subsystem                                                                                                 |
&#9745;    | InvLicenseProducts.go          | InvLicenseProducts          | Interface to manage License Management (third party) Functionality                                                                                      |
&#9745;    | IWebSrvSettings.go             | IWebSrvSettings             | Web server settings proxy class                                                                                                                         |
&#9744;    |                                | IWebUsersSrv                | Send an email to multiple recipients.                                                                                                                   |
&#9744;    |                                | IWebUsersSrv2               | Operating with emails from GUI                                                                                                                          |
&#9744;    |                                | KeyService                  | Interface for working with KeyService subsystem                                                                                                         |
&#9744;    |                                | KeyService2                 | Additional interface for working with KeyService subsystem                                                                                              |
&#9745;    | KillChain.go                   | KillChain                   | KillChain info from host                                                                                                                                |
&#9745;    | KLEVerControl.go               | KLEVerControl               | Controls the possibility to download and automatically create installation packages                                                                     |
&#9745;    | KsnInternal.go                 | KsnInternal                 | Interface for working with KsnProxy subsystem                                                                                                           |
&#9745;    | LicenseInfoSync.go             | LicenseInfoSync             | Operating with licenses                                                                                                                                 |
&#9745;    | LicenseKeys.go                 | LicenseKeys                 | Operating with keys                                                                                                                                     |
&#9745;    | LicensePolicy.go               | LicensePolicy               | License policy                                                                                                                                          |
&#9745;    | Limits.go                      | Limits                      | Interface for working with Limits subsystem                                                                                                             |
&#9745;    | ListTags.go                    | ListTags                    | Interface allows to acquire and manage tags to various KSC objects                                                                                      |
&#9745;    | MigrationData.go               | MigrationData               | Migration of data between KSC On-Premise and KSCHosted                                                                                                  |
&#9745;    | Multitenancy.go                | Multitenancy                | Multitenancy product managing                                                                                                                           |
&#9745;    | NagCgwHelper.go                | NagCgwHelper                | Nagent CGW (Connection Gateway) API                                                                                                                     |
&#9745;    | NagGuiCalls.go                 | NagGuiCalls                 | Remote host caller                                                                                                                                      |
&#9745;    | NagHstCtl.go                   | NagHstCtl                   | Manage nagent on host                                                                                                                                   |
&#9744;    |                                | NagNetworkListApi           | Nagent OpenAPI to work with network lists                                                                                                               |
&#9745;    | NagRdu.go                      | NagRdu                      | Remote diagnostics on host                                                                                                                              |
&#9745;    | NagRemoteScreen.go             | NagRemoteScreen             | Interface for remote screen session management                                                                                                          |
&#9745;    | NlaDefinedNetworks.go          | NlaDefinedNetworks          | Network location awareness (NLA) defined networks. Used as a scope for Update agents. Each NLA-defined network is defined by list of NLA locations      |
&#9745;    | OsVersion.go                   | OsVersion                   | Operating systems dictionary access                                                                                                                     |
&#9745;    | PackagesApi.go                 | PackagesApi                 | Operating with packages                                                                                                                                 |
&#9744;    |                                | PatchParameters             | Patch parameters processing                                                                                                                             |
&#9744;    |                                | PLCDevApi                   | Interface allows to acquire and manage PLC devices registry                                                                                             |
&#9745;    | Policy.go                      | Policy                      | Policies processing                                                                                                                                     |
&#9744;    |                                | PolicyProfiles              | Policy profiles processing                                                                                                                              |
&#9745;    | QBTNetworkListApi.go           | QBTNetworkListApi           | Interface to working with Quarantine, Backup and TIF network lists                                                                                      |
&#9745;    | QueriesStorage.go              | QueriesStorage              | QueriesStorage interface                                                                                                                                |
&#9745;    | ReportManager.go               | ReportManager               | Reports managing                                                                                                                                        |
&#9745;    | RetrFiles.go                   | RetrFiles                   | Class provides means to get retranslated files info                                                                                                     |
&#9745;    | ScanDiapasons.go               | ScanDiapasons               | Network subnets processing                                                                                                                              |
&#9745;    | SecurityPolicy.go              | SecurityPolicy              | Allows to manage users and permissions.                                                                                                                 |
&#9745;    | SecurityPolicy3.go             | SecurityPolicy3             | Allows to manage security groups of internal users. Use srvview SplUserGroupSrvViewName to get information about relationship between users and groups  |
&#9745;    | ServerHierarchy.go             | ServerHierarchy             | Server hierarchy management interface                                                                                                                   |
&#9745;    | ServerTransportSettings.go     | ServerTransportSettings     | Server transport settings proxy class                                                                                                                   |
&#9745;    | Session.go                     | Session                     | Session management interface                                                                                                                            |
&#9745;    | SmsQueue.go                    | SmsQueue                    | Manage SMS message queue                                                                                                                                |
&#9745;    | SmsSenders.go                  | SmsSenders                  | Configure mobile devices as SMS senders                                                                                                                 |
&#9745;    | SrvCloud.go                    | SrvCloud                    | Interface to acquire info about public clouds                                                                                                           |
&#9745;    | SrvSsRevision.go               | SrvSsRevision               | Access to virtual server settings storage revisions                                                                                                     |
&#9745;    | SrvView.go                     | SrvView                     | Interface to get plain-queries from SC-server                                                                                                           |
&#9745;    | SsContents.go                  | SsContents                  | Access to settings storage                                                                                                                              |
&#9745;    | SubnetMasks.go                 | SubnetMasks                 | Subnets provider                                                                                                                                        |
&#9745;    | Tasks.go                       | Tasks                       | Group tasks                                                                                                                                             |
&#9745;    | TrafficManager.go              | TrafficManager              | Traffic manager interface                                                                                                                               |
&#9745;    | UaControl.go                   | UaControl                   | Update agents and Connection gateways management                                                                                                        |
&#9745;    | Updates.go                     | Updates                     | Updates processing                                                                                                                                      |
&#9744;    |                                | UpdComps                    | Class provides means to manage updatable components (bases)                                                                                             |
&#9745;    | UserDevicesApi.go              | UserDevicesApi              | Interface to unified mobile device management                                                                                                           |
&#9745;    | VapmControlApi.go              | VapmControlApi              | VAPM                                                                                                                                                    |
&#9745;    | VServers.go                    | VServers                    | Virtual servers processing                                                                                                                              |
&#9745;    | VServers2.go                   | VServers2                   | Virtual servers processing                                                                                                                              |
&#9745;    | WolSender.go                   | WolSender                   | Wake-On-LAN signal sender                                                                                                                               |

## NOTE
Using the [context](https://godoc.org/context) package, one can easily
pass cancelation signals and deadlines to various services of the client for
handling a request. In case there is no context available, then `context.Background()`
can be used as a starting point.


## License

This library is distributed under the  MIT LICENSE found in the [LICENSE](./LICENSE)
file.
