# go-ksc #
go-ksc is a Go client library for accessing the KSC (Kaspersky) Open API.

## Usage ##

```go
import (
    "github.com/pixfid/go-ksc/kaspersky"
)

```

Construct a new KSC client, then use the various services on the client to
access different parts of the KSC (Kaspersky) Open API. For example:

```go
        ctx := context.Background()
    	cfg := kaspersky.Config {
    		Username: "login",
    		Password: "password",
    		Server: fmt.Sprintf(`https://%s:%s`, "ip", "port"),
    	}
    
    	client := kaspersky.New(cfg)
    	client.KSCAuth(ctx)

        //Get List of Windows domain in the network.
        groups_domains, err := client.HostGroup.GetDomains(context.Background())
```

As example find hosts by params:
```go
	hosts := data.FullHostsInfo{}

	accessor, _ := client.HostGroup.FindHosts(ctx, data.HSPParam{
		WstrFilter:        `(&(KLHST_WKS_GROUPID <> 4)(KLHST_WKS_FQDN = "*t457-zt*"))`,
		VecFieldsToReturn: []string{"KLHST_WKS_DN", "KLHST_WKS_GROUPID", "KLHST_WKS_OS_NAME"},
		PParams:           data.PParams{
			KlgrpFindFromCurVsOnly: true,
			KlsrvhSlaveRecDepth: 0,
		},
		LMaxLifeTime:      100,
	})

	count, _ := client.ChunkAccessor.GetItemsCount(ctx, accessor.StrAccessor)
	_ = client.ChunkAccessor.GetItemsChunk(ctx, accessor.StrAccessor, 0, count.Int, &hosts)
	client.ChunkAccessor.Release(ctx, accessor.StrAccessor)
```


NOTE: Using the [context](https://godoc.org/context) package, one can easily
pass cancelation signals and deadlines to various services of the client for
handling a request. In case there is no context available, then `context.Background()`
can be used as a starting point.


## License ##

This library is distributed under the  GNU GENERAL PUBLIC LICENSE Version 3 found in the [LICENSE](./LICENSE)
file.