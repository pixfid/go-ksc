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
Examples:

Find online hosts:

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

Get installed products on host by HostId:

```go
products, err := client.HostGroup.GetHostProducts(ctx, "8910f900-3807-4b97-8a97-d49e73ec5ab1")
```

Response:

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

Get Lists tasks on Host: 

```go
tastList, raw, err := client.Tasks.GetAllTasksOfHost(ctx, "", "c2b22f83-307c-45aa-8533-5ffffbcc6bf1")
```

Response:

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
Find srvview data by filter string. A removable device's collection. 

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

Response:

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


NOTE: Using the [context](https://godoc.org/context) package, one can easily
pass cancelation signals and deadlines to various services of the client for
handling a request. In case there is no context available, then `context.Background()`
can be used as a starting point.


## License ##

This library is distributed under the  MIT LICENSE found in the [LICENSE](./LICENSE)
file.
