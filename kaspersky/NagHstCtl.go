/*
 *
 * 	Copyright (C) 2020  <Semchenko Aleksandr>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful, but
 * WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.See the GNU
 * General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.If not, see <http://www.gnu.org/licenses/>.
 * /
 */

package kaspersky

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type NagHstCtl struct {
	client *Client
}

//	Acquire runtime host information
//
//	Parameters:
//	- pFilter	filter, may contain following sub-containers:
//	- klhst-rt-TskInfo "Information about running tasks" is required
//	- "klhst-rt-PrdNamesArray" name of interesting products,
//	if this variable is absent info about all products will be returned
//	- "klhst-ProductVersion" this parameter is needed in the resulting container
//	- taskType this parameter is needed in the resulting container
//	- taskState this parameter is needed in the resulting container
//	- taskStorageId this parameter is needed in the resulting container
//	- klhst-rt-InstInfo "Information about running PRCI component instances" is required instStatistics
//	- klhst-ComponentName
//	- klhst-InstanceId
//	- klhst-rt-PrdNamesArray
//	- KLHST_APP_INFO "Information about KL products" is required
//	Example for filter is below:
//
//                    +--- (PARAMS_T)
//                        +---klhst-rt-TskInfo (PARAMS_T)
//                        |    +---klhst-ProductVersion = (STRING_T)""
//                        |    +---taskState = (INT_T)0
//                        |    +---taskStorageId = (STRING_T)""
//                        +---klhst-rt-InstInfo (PARAMS_T)
//                        |    +---instStatistics (PARAMS_T)
//                        |    +---klhst-ComponentName = (STRING_T)""
//                        |    +---klhst-InstanceId = (STRING_T)""
//                        |    +---klhst-rt-PrdNamesArray (ARRAY_T)
//                        |       +---0 = (STRING_T)"1103"
//                        +---KLHST_APP_INFO (PARAMS_T)
//	Returns:
//	runtime host information, may contain following sub-containers:
//	- klhst-rt-TskInfo "Information about running tasks"
//	- klhst-rt-InstInfo "Information about running PRCI component instances"
//	- KLHST_APP_INFO "Information about KL products"
func (nh *NagHstCtl) GetHostRuntimeInfo(ctx context.Context, params interface{}) ([]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", nh.client.Server+"/api/v1.0/NagHstCtl.GetHostRuntimeInfo", bytes.NewBuffer(postData))
	if err != nil {
		log.Fatal(err.Error())
	}

	raw, err := nh.client.Do(ctx, request, nil)
	return raw, err
}

//	Initiate changing state of tasks at host
//
//	The method sends to the specified product task one of such commands as 'start', 'stop', 'suspend', 'resume'.
//
//	Parameters:
//	- szwProduct (string) Product name *
//	- szwVersion (string) Product version *
//	+-----------------------------------------------------------------+--------------------------+-----------------------------+
//	|                             Product                             | "KLHST_WKS_PRODUCT_NAME" | "KLHST_WKS_PRODUCT_VERSION" |
//	+-----------------------------------------------------------------+--------------------------+-----------------------------+
//	| Kaspersky Administration Server                                 | "1093"                   | "1.0.0.0"                   |
//	| Kaspersky Network Agent                                         | "1103"                   | "1.0.0.0"                   |
//	| Kaspersky Anti-Virus 5.0 for Windows Workstation                | "Workstation"            | "5.0.0.0"                   |
//	| Kaspersky Anti-Virus 5.0 for Windows File Server                | "Fileserver"             | "5.0.0.0"                   |
//	| Kaspersky Anti-Virus 6.0 for Windows Workstation                | "KAVWKS6"                | "6.0.0.0"                   |
//	| Kaspersky Anti-Virus 6.0 for Windows File Server                | "KAVFS6"                 | "6.0.0.0"                   |
//	| Kaspersky Anti-Virus 6.0 for Windows Workstation MP1            | "KAVWKS6"                | "6.0.4.0"                   |
//	| Kaspersky Anti-Virus 6.0 for Windows File Server MP1            | "KAVFS6"                 | "6.0.4.0"                   |
//	| Kaspersky Anti-Virus 6.0 for Windows Servers Enterprise Edition | "KAVFSEE"                | "6.0.0.0"                   |
//	| Kaspersky Anti-Virus 8.0 for Windows Servers Enterprise Edition | "KAVFSEE"                | "8.0.0.0"                   |
//	| Kaspersky Anti-Virus 6.0 for Linux File Server                  | "KLIN"                   | "6.0.0.0"                   |
//	| Kaspersky Enterprise Security 8.0                               | "KES"                    | "8.1.0.0"                   |
//	| Kaspersky Enterprise Security 10.0                              | "KES"                    | "10.1.0.0"                  |
//	| Kaspersky Enterprise Security 10.0 Maintenance Release 1        | "KES"                    | "10.2.1.0"                  |
//	| Kaspersky Enterprise Security 10.0 Service Pack 1               | "KES"                    | "10.2.2.0"                  |
//	| Kaspersky Security for Virtualization Light Agent 3.0           | "KSVLA"                  | "3.0.0.0"                   |
//	| Kaspersky Endpoint Security 10 MR1                              | "KES"                    | "10.2.1.0"                  |
//	| Kaspersky Endpoint Security 10 SP1                              | "KES"                    | "10.2.2.0"                  |
//	| Kaspersky Endpoint Security 10 SP1 MR1                          | "KES"                    | "10.2.4.0"                  |
//	| Kaspersky Endpoint Security 11                                  | "KES"                    | "11.0.0.0"                  |
//	+-----------------------------------------------------------------+--------------------------+-----------------------------+
//
//	- szwTaskStorageId (string) "Tasks storage identifier of the task" (see HostTasks)
//	- nTaskAction (int64) one of following values:
//
//	+----+---------------------+------------------+
//	| Id |        Name         |   Description    |
//	+----+---------------------+------------------+
//	|  5 | (TSK_ACTION_START)  | start the task   |
//	|  0 | (TSK_ACTION_STOP)   | stop the task    |
//	|  1 | (TSK_ACTION_SUSPEND | suspend the task |
//	|  2 | (TSK_ACTION_RESUME) | resume the task  |
//	+----+---------------------+------------------+
func (nh *NagHstCtl) SendTaskAction(ctx context.Context, szwProduct, szwVersion, szwTaskStorageId string,
	nTaskAction int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`
	{
		"szwProduct": "%s",
		"szwVersion": "%s",
		"szwTaskStorageId": "%s",
        "parentId": %d
    
	}`, szwProduct, szwVersion, szwTaskStorageId, nTaskAction))

	request, err := http.NewRequest("POST", nh.client.Server+"/api/v1.0/NagHstCtl.SendTaskAction", bytes.NewBuffer(postData))
	if err != nil {
		log.Fatal(err.Error())
	}

	raw, err := nh.client.Do(ctx, request, nil)
	return raw, err
}

//	Initiate changing state of products at host
//
//	The method sends to the specified product 'start' or 'stop' command.
//
//	Parameters:
//	- szwProduct	Product name
//	- szwVersion	Product version
//	+-----------------------------------------------------------------+--------------------------+-----------------------------+
//	|                             Product                             | "KLHST_WKS_PRODUCT_NAME" | "KLHST_WKS_PRODUCT_VERSION" |
//	+-----------------------------------------------------------------+--------------------------+-----------------------------+
//	| Kaspersky Administration Server                                 | "1093"                   | "1.0.0.0"                   |
//	| Kaspersky Network Agent                                         | "1103"                   | "1.0.0.0"                   |
//	| Kaspersky Anti-Virus 5.0 for Windows Workstation                | "Workstation"            | "5.0.0.0"                   |
//	| Kaspersky Anti-Virus 5.0 for Windows File Server                | "Fileserver"             | "5.0.0.0"                   |
//	| Kaspersky Anti-Virus 6.0 for Windows Workstation                | "KAVWKS6"                | "6.0.0.0"                   |
//	| Kaspersky Anti-Virus 6.0 for Windows File Server                | "KAVFS6"                 | "6.0.0.0"                   |
//	| Kaspersky Anti-Virus 6.0 for Windows Workstation MP1            | "KAVWKS6"                | "6.0.4.0"                   |
//	| Kaspersky Anti-Virus 6.0 for Windows File Server MP1            | "KAVFS6"                 | "6.0.4.0"                   |
//	| Kaspersky Anti-Virus 6.0 for Windows Servers Enterprise Edition | "KAVFSEE"                | "6.0.0.0"                   |
//	| Kaspersky Anti-Virus 8.0 for Windows Servers Enterprise Edition | "KAVFSEE"                | "8.0.0.0"                   |
//	| Kaspersky Anti-Virus 6.0 for Linux File Server                  | "KLIN"                   | "6.0.0.0"                   |
//	| Kaspersky Enterprise Security 8.0                               | "KES"                    | "8.1.0.0"                   |
//	| Kaspersky Enterprise Security 10.0                              | "KES"                    | "10.1.0.0"                  |
//	| Kaspersky Enterprise Security 10.0 Maintenance Release 1        | "KES"                    | "10.2.1.0"                  |
//	| Kaspersky Enterprise Security 10.0 Service Pack 1               | "KES"                    | "10.2.2.0"                  |
//	| Kaspersky Security for Virtualization Light Agent 3.0           | "KSVLA"                  | "3.0.0.0"                   |
//	| Kaspersky Endpoint Security 10 MR1                              | "KES"                    | "10.2.1.0"                  |
//	| Kaspersky Endpoint Security 10 SP1                              | "KES"                    | "10.2.2.0"                  |
//	| Kaspersky Endpoint Security 10 SP1 MR1                          | "KES"                    | "10.2.4.0"                  |
//	| Kaspersky Endpoint Security 11                                  | "KES"                    | "11.0.0.0"                  |
//	+-----------------------------------------------------------------+--------------------------+-----------------------------+
//
//	- nProductAction	one of following values:
//	+----+--------------------+--------------------------------------------+
//	| Id |        Name        |                Description                 |
//	+----+--------------------+--------------------------------------------+
//	|  0 | (APP_ACTION_START) | initiate starting of the specified product |
//	|  1 | (APP_ACTION_STOP)  | initiate stopping of the specified product |
//	+----+--------------------+--------------------------------------------+
func (nh *NagHstCtl) SendProductAction(ctx context.Context, szwProduct, szwVersion string,
	nProductAction int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`
	{
		"szwProduct": "%s",
		"szwVersion": "%s",
		"szwTaskStorageId": %d    
	}`, szwProduct, szwVersion, nProductAction))

	request, err := http.NewRequest("POST", nh.client.Server+"/api/v1.0/NagHstCtl.SendProductAction", bytes.NewBuffer(postData))
	if err != nil {
		log.Fatal(err.Error())
	}

	raw, err := nh.client.Do(ctx, request, nil)
	return raw, err
}
