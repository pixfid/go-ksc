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

//	Tasks Class Reference
//
//	Group tasks.
//
//	Allows to acquire task attributes, enumerate, control and delete tasks.
//
//	List of all members.
type Tasks service

//	Get all group and global tasks of specified host.
//
//	Parameters:
//	- ctx	(context.Context) context.
//	- strDomainName	(string) domain name.
//	- strHostName	(string) host name.
//
//	Returns:
//	- (array) array of string with task ids
func (ts *Tasks) GetAllTasksOfHost(ctx context.Context, strDomainName, strHostName string) (*PxgValArrayOfString,
	[]byte, error) {
	postData := []byte(fmt.Sprintf(`{"strDomainName": "%s","strHostName": "%s"}`, strDomainName, strHostName))
	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/Tasks.GetAllTasksOfHost", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValArrayOfString := new(PxgValArrayOfString)
	raw, err := ts.client.Do(ctx, request, &pxgValArrayOfString)
	return pxgValArrayOfString, raw, err
}

//	TaskData struct
type TaskData struct {
	PxgRetVal Task `json:"PxgRetVal"`
}

type Task struct {
	DisplayName          string               `json:"DisplayName"`
	PrtsTaskCreationDate PrtsTaskCreationDate `json:"PRTS_TASK_CREATION_DATE"`
	TaskidProductName    string               `json:"TASKID_PRODUCT_NAME"`
	TaskidVersion        string               `json:"TASKID_VERSION"`
	TaskName             string               `json:"TASK_NAME"`
	TaskUniqueID         string               `json:"TASK_UNIQUE_ID"`
}

type PrtsTaskCreationDate struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

//	Acquire attributes of specified task.
//
//	Returns attributes of specified task.
//
//	Parameters:
//	- ctx	(context.Context) context.
//	- strTask	(string) task id.
//
//	Returns:
//	- (params) object containing task attributes (see List of task attributes).
func (ts *Tasks) GetTask(ctx context.Context, strTask string) (*TaskData, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"strTask": "%s"}`, strTask))
	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/Tasks.GetTask", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	taskData := new(TaskData)
	raw, err := ts.client.Do(ctx, request, &taskData)
	return taskData, raw, err
}

//	Acquire task settings.
//
//	Returns task settings.
//
//	Parameters:
//	- ctx	(context.Context) context.
//	- strTask	(string) task id.
//
//	Returns:
//	- (params) task settings, see Task settings format
func (ts *Tasks) GetTaskData(ctx context.Context, strTask string, tsk interface{}) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`
	{
	"strTask": "%s"
	}`, strTask))
	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/Tasks.GetTaskData", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ts.client.Do(ctx, request, &tsk)
	return raw, err
}

//	Return the group id for the group task.
//
//	Parameters:
//	- ctx	(context.Context) context.
//	- strTaskId	(string) task id
//
//	Returns:
//	- (int64) group id
func (ts *Tasks) GetTaskGroup(ctx context.Context, strTaskId string) (*PxgValInt, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"strTaskId": "%s"}`, strTaskId))
	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/Tasks.GetTaskGroup", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValInt := new(PxgValInt)
	raw, err := ts.client.Do(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

//TaskStatistics struct
type TaskStatistics struct {
	TaskStatistic TaskStatistic `json:"PxgRetVal"`
}

type TaskStatistic struct {
	The1                 int64 `json:"1"`
	The2                 int64 `json:"2"`
	The4                 int64 `json:"4"`
	The8                 int64 `json:"8"`
	The16                int64 `json:"16"`
	The32                int64 `json:"32"`
	The64                int64 `json:"64"`
	GnrlCompletedPercent int64 `json:"GNRL_COMPLETED_PERCENT"`
	KltskNeedRbtCnt      int64 `json:"KLTSK_NEED_RBT_CNT"`
}

//	Acquire statistics of the specified task.
//
//	Returns statistics of the specified task.
//
//	Parameters:
//	- ctx	(context.Context) context.
//	- strTask	(string) task id.
//
//	Returns:
//	- (TaskStatistics struct) object containing task statistics, see List of task statistics attributes.
func (ts *Tasks) GetTaskStatistics(ctx context.Context, strTask string) (*TaskStatistics, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"strTask": "%s"}`, strTask))
	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/Tasks.GetTaskStatistics", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	taskStatistics := new(TaskStatistics)
	raw, err := ts.client.Do(ctx, request, &taskStatistics)
	return taskStatistics, raw, err
}

//	Suspend execution of the specified task.
//
//	Suspends execution of the specified task.
//
//	Parameters:
//	- ctx	(context.Context) context.
//	- strTask	(string) task id.
func (ts *Tasks) SuspendTask(ctx context.Context, strTask string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"strTask": "%s"}`, strTask))
	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/Tasks.SuspendTask", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ts.client.Do(ctx, request, nil)
	return raw, err
}

//	Resume specified task.
//
//	Resumes specified task
//
//	Parameters:
//	- ctx	(context.Context) context.
//	- strTask	(string) task id.
func (ts *Tasks) ResumeTask(ctx context.Context, strTask string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"strTask": "%s"}`, strTask))
	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/Tasks.ResumeTask", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ts.client.Do(ctx, request, nil)
	return raw, err
}

//	Start specified task.
//
//	Forces starting specified task.
//
//	Parameters:
//	- ctx	(context.Context) context.
//	- strTask	(string) task id.
func (ts *Tasks) RunTask(ctx context.Context, strTask string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"strTask": "%s"}`, strTask))
	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/Tasks.RunTask", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ts.client.Do(ctx, request, nil)
	return raw, err
}

//	Delete the specified task.
//
//	Deletes the specified task.
//
//	Parameters:
//	- ctx	(context.Context) context.
//	- strTask	(string) task id.
func (ts *Tasks) DeleteTask(ctx context.Context, strTask string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"strTask": "%s"}`, strTask))
	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/Tasks.DeleteTask", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ts.client.Do(ctx, request, nil)
	return raw, err
}

//	Cancel execution of the specified task.
//
//	Cancels execution of the specified task.
//
//	Parameters:
//	- strTask	(string) task id.
func (ts *Tasks) CancelTask(ctx context.Context, strTask string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"strTask": "%s"}`, strTask))
	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/Tasks.CancelTask", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ts.client.Do(ctx, request, nil)
	return raw, err
}

type TaskHistoryParams struct {
	StrTask        string          `json:"strTask"`
	PFields2Return []string        `json:"pFields2Return"`
	PSortFields    []FieldsToOrder `json:"pSortFields"`
	StrHostName    string          `json:"strHostName"`
	PFilter        interface{}     `json:"pFilter"`
}

//	Acquire task execution history events.
//
//	Returns task execution history events.
//
//Parameters:
//	- strTask	(string) task id.
//	- pFields2Return	(array) array of task history event attribute names to return See List of event attributes for
//	attribute names.
//	- pSortFields	(array) array of containers each of them containing two attributes:
//		-- "Name" of type String, name of attribute used for sorting. See List of event attributes for attribute names.
//		-- "Asc" of type bool, ascending if true descending otherwise
//	- strHostName	(string) name of the host. Events for specified host will be returned.
//	*Events from all hosts will be returned if "" (empty string) is specified.*
//	- pFilter	(params) object containing values for attributes to filter events.
//	Only events with matching attribute values will be returned.
//	If empty all events for task will be returned. See List of event filter attributes for attribute names.
//
//	Return:
//	 -strIteratorId	(string) result-set ID, identifier of the server-side collection of task history events,
//	 to acquire data use EventProcessing.GetRecordRange,
//	 after iterator MUST be realesed by EventProcessing.ReleaseIterator.
//
//	Example:
//	strIteratorId, _, _ := client.Tasks.GetTaskHistory(ctx, kaspersky.TaskHistoryParams{
//		StrTask:        "195",
//		PFields2Return: []string{
//			"hostdn",
//			"product_name",
//			"product_displ_version",
//			"product_version",
//			"task_display_name",
//			"GNRL_COMPLETED_PERCENT",
//			"event_id",
//			"host_type",
//		},
//		PSortFields:    []kaspersky.PSortFields{
//			{Type: "params", PSortField: kaspersky.PSortField{
//				Name: "event_id",
//				Asc:  true,
//			}},
//		},
//		StrHostName:    "c2b22f83-307c-45aa-8533-5ffffbcc6bf1",
//		PFilter:        nil,
//	})
func (ts *Tasks) GetTaskHistory(ctx context.Context, params interface{}) (*StrIteratorId, []byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/Tasks.GetTaskHistory", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	strIteratorID := new(StrIteratorId)
	raw, err := ts.client.Do(ctx, request, &strIteratorID)
	return strIteratorID, raw, err
}

//	Get task start event.
//
//	Returns event which should run the task.
//
//	Parameters:
//	- strTask	(string) task id.
//
//	Returns:
//	- (params) parameters of the event, see Task event filter attributes:
//
//	+------------------------------+-------------+-----------------------------------------------------------------------------------+
//	|             Name             |    Type     |                                    Description                                    |
//	+------------------------------+-------------+-----------------------------------------------------------------------------------+
//	|                              |             |                                                                                   |
//	| FILTER_EVENTS_PRODUCT_NAME   | string      | Name of a product which publishes an event                                        |
//	|                              |             |                                                                                   |
//	| FILTER_EVENTS_VERSION        | string      | Version of a product which publishes an event                                     |
//	|                              |             |                                                                                   |
//	| FILTER_EVENTS_COMPONENT_NAME | string      | Name of a component which publishes an event                                      |
//	|                              |             |                                                                                   |
//	| FILTER_EVENTS_INSTANCE_ID    | string      | Instance of a component which publishes an event                                  |
//	|                              |             |                                                                                   |
//	| EVENT_TYPE                   | string      | Type of an event. See List of event attributes for event types                    |
//	|                              |             |                                                                                   |
//	| EVENT_BODY_FILTER            | interface{} | Events filter. Here is an filter example for failure of a task:                   |
//	|                              |             | +---EVENT_TYPE = (string)KLPRCI_TaskState                                    |
//	|                              |             |     +---EVENT_BODY_FILTER (paramParams)                                           |
//	|                              |             |     |   +---KLPRCI_TASK_TS_ID = (string)fef4c022-ae55-41a7-afb0-0cc7b3654e70 |
//	|                              |             |     |   +---KLPRCI_newState = (int64)3                                         |
//	|                              |             |                                                                                   |
//	+------------------------------+-------------+-----------------------------------------------------------------------------------+
func (ts *Tasks) GetTaskStartEvent(ctx context.Context, strTask string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"strTask": "%s"}`, strTask))
	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/Tasks.GetTaskStartEvent", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ts.client.Do(ctx, request, nil)
	return raw, err
}

//	Encrypt an account password.
//
//	Encrypts account password to push as "klprts-TaskAccountPassword" attribute. See List of deployment task attributes
//
//	The same as DataProtectionApi.ProtectUtf16StringGlobally
//
//	Parameters:
//	- strPassword	(string) password to protect
//
//	Returns:
//	- (binary) Encrypted password
//
//	See also:
//	DataProtectionApi.ProtectUtf16StringGlobally
func (ts *Tasks) ProtectPassword(ctx context.Context, strPassword string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"strPassword": "%s"}`, strPassword))
	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/Tasks.ProtectPassword", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ts.client.Do(ctx, request, nil)
	return raw, err
}

//TasksIteratorParams struct
type TasksIteratorParams struct {
	NGroupID            int64  `json:"nGroupId"`
	BGroupIDSignificant bool   `json:"bGroupIdSignificant"`
	StrProductName      string `json:"strProductName"`
	StrVersion          string `json:"strVersion"`
	StrComponentName    string `json:"strComponentName"`
	StrInstanceID       string `json:"strInstanceId"`
	StrTaskName         string `json:"strTaskName"`
	BIncludeSupergroups bool   `json:"bIncludeSupergroups"`
}

//	Reset task iterator for a specified filter data.
//
//	If one of the parameters is not specified then the filtration will not be performed by this parameter.
//
//	Parameters:
//	- nGroupId	(int64) group id
//	- bGroupIdSignificant	(bool) set true if group id is significant, if false nGroupId will be skipped
//	- strProductName	(string) product name
//	- strVersion	(string) product version
//	- strComponentName	(string) component name
//	- strInstanceId	(string) instance id
//	- strTaskName	(string) task name
//	- bIncludeSupergroups	(bool) set true if you need to include Super groups.
//
//	The group Super is parent of the group Groups and is intended
//	for assignment of group tasks and policies received from the master server.
//
//	Return:
//	- strTaskIteratorId	(string) iterator id, to get data use Tasks.GetNextTask,
//	after iterator MUST be realesed by Tasks.ReleaseTasksIterator
func (ts *Tasks) ResetTasksIterator(ctx context.Context, params TasksIteratorParams) ([]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/Tasks.ResetTasksIterator", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ts.client.Do(ctx, request, nil)
	return raw, err
}

//	Release task iterator.
//
//	Parameters:
//	- strTaskIteratorId	(string) iterator id got from Tasks.ResetTasksIterator
func (ts *Tasks) ReleaseTasksIterator(ctx context.Context, strTaskIteratorId string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"strTaskIteratorId": "%s"}`, strTaskIteratorId))
	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/Tasks.ReleaseTasksIterator", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ts.client.Do(ctx, request, nil)
	return raw, err
}

//	Release iterator of specified data.
//
//	Releases iterator of specified data and frees associated memory
//
//	Parameters:
//	- strHostIteratorId	(string) iterator id which got from
//	Tasks.ResetHostIteratorForTaskStatus
//	or
//	Tasks.ResetHostIteratorForTaskStatusEx
func (ts *Tasks) ReleaseHostStatusIterator(ctx context.Context, strHostIteratorId string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"strHostIteratorId": "%s"}`, strHostIteratorId))
	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/Tasks.ReleaseHostStatusIterator", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ts.client.Do(ctx, request, nil)
	return raw, err
}

//HostIteratorForTaskParams struct
type HostIteratorForTaskParams struct {
	StrTask        string   `json:"strTask"`
	NHostStateMask string   `json:"nHostStateMask"`
	PFields2Return []string `json:"pFields2Return"`
	NLifetime      int64    `json:"nLifetime"`
}

//	Make host task states request.
//
//	Makes request of the status of group task that runs on many machines
//
//	Parameters:
//	- strTask	(wstring) task id.
//	- pFields2Return	(array) array of attribute names to return. See Host task state attributes
//	- nHostStateMask	(int) host task state. See Bit masks of host task states
//	- nLifetime	(int) lifetime in seconds
//
//	Return:
//	- strHostIteratorId	(wstring) iterator id, to get requsted data use Tasks.GetNextHostStatus,
//	Tasks.GetHostStatusRecordRange
func (ts *Tasks) ResetHostIteratorForTaskStatus(ctx context.Context, params HostIteratorForTaskParams) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/Tasks.ResetHostIteratorForTaskStatus", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ts.client.Do(ctx, request, nil)
	return raw, err
}

//HostIteratorForTaskParamsEx struct
type HostIteratorForTaskParamsEx struct {
	StrTask        string          `json:"strTask"`
	NHostStateMask int64           `json:"nHostStateMask"`
	PFields2Return []string        `json:"pFields2Return"`
	PFields2Order  []FieldsToOrder `json:"pFields2Order"`
	NLifetime      int64           `json:"nLifetime"`
}

//	Make host task states request.
//
//	Makes request of the status of group task that runs on many machines
//
//	Parameters:
//	- strTask	(string) task id.
//	- pFields2Return	(array) array of attribute names to return. See Host task state attributes
//	- pFields2Order	(array) array of containers each of them containing two attributes:
//		-- "Name" (string) name of attribute used for sorting
//		-- "Asc" (paramBool) ascending if true descending otherwise
//	- nHostStateMask	(int) host task state.
//	Bit masks of host task states:
//
//	+------+------------------------------------------------------------------------------------------------------------------------+
//	| Mask |                                                      Description                                                       |
//	+------+------------------------------------------------------------------------------------------------------------------------+
//	| 0x01 | "Pending" state, which means actual group synchronization associated with the task is not delivered to the target host |
//	| 0x02 | "Running" state                                                                                                        |
//	| 0x04 | "Finished successfully" state                                                                                          |
//	| 0x08 | "Finished with warning" state                                                                                          |
//	| 0x10 | "Failed" state                                                                                                         |
//	| 0x20 | "Scheduled" state, which means task is ready to start (manually or by the schedule) on the target host                 |
//	| 0x40 | "Paused" state                                                                                                         |
//	+------+------------------------------------------------------------------------------------------------------------------------+
//	- nLifetime	(int) lifetime in seconds
//
//	Return:
//	- strHostIteratorId	(string) iterator id,
//	to get requsted data use Tasks.GetNextHostStatus, Tasks.GetHostStatusRecordRange
//
//Example:
//	iteratorForTaskStatusEx, _, _ := client.Tasks.ResetHostIteratorForTaskStatusEx(ctx, kaspersky.HostIteratorForTaskParams{
//		StrTask:        "195",
//		NHostStateMask: 0x01,
//		PFields2Return: []string{"hostname", "state_descr"},
//		PFields2Order:  []kaspersky.FieldsToOrder{
//			{
//				Type:	"params",
//				OrderValue: kaspersky.OrderValue{
//					Name: "hostname",
//					Asc: true},
//			},
//		},
//		NLifetime:      100,
//	})
func (ts *Tasks) ResetHostIteratorForTaskStatusEx(ctx context.Context, params HostIteratorForTaskParamsEx) (*StrHostIteratorId, []byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/Tasks.ResetHostIteratorForTaskStatusEx", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	strHostIteratorId := new(StrHostIteratorId)
	raw, err := ts.client.Do(ctx, request, &strHostIteratorId)
	return strHostIteratorId, raw, err
}

//	Get records count of result of operation
//	ResetHostIteratorForTaskStatus or ResetHostIteratorForTaskStatusEx.
//
//	Parameters:
//	- strHostIteratorId	(string) iterator id which got from
//	Tasks.ResetHostIteratorForTaskStatus or Tasks.ResetHostIteratorForTaskStatusEx
//
//	Returns:
//	- (int64) returns records count
func (ts *Tasks) GetHostStatusRecordsCount(ctx context.Context, strHostIteratorId string) (*PxgValInt, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"strHostIteratorId": "%s"}`, strHostIteratorId))
	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/Tasks.GetHostStatusRecordsCount", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValInt := new(PxgValInt)
	raw, err := ts.client.Do(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

//	Get result of operation ResetHostIteratorForTaskStatus or ResetHostIteratorForTaskStatusEx.
//
//	Gets result of operation Tasks.ResetHostIteratorForTaskStatus or Tasks.ResetHostIteratorForTaskStatusEx
//
//	Parameters:
//	- strHostIteratorId	(string) iterator id which got from
//	Tasks.ResetHostIteratorForTaskStatus or Tasks.ResetHostIteratorForTaskStatusEx
//	- nStart	(int) zero-based start position.
//	- nEnd	(int) zero-based finish position.
//
//	Returns:
//	- (int64) actual number of elements contained in the record set
//	- pParHostStatus	(params) container that has requested elements in the array with name "statuses",
//	each item of array contains attributes from Host task state attributes
//
//	Example:
//{
//  "pParHostStatus" : {
//    "statuses" : [
//      {
//        "type" : "params",
//        "value" : {
//          "hostdn" : "HostDisplayName",
//          "hostname" : "53bf5bda-d728-4888-b002-67e63b6e4c63"
//        }
//      }
//    ]
//  },
//  "PxgRetVal" : 1
//}
func (ts *Tasks) GetHostStatusRecordRange(ctx context.Context, strHostIteratorId string, nStart, nEnd int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"strHostIteratorId": "%s", "nStart": %d, "nEnd" : %d}`, strHostIteratorId,
		nStart, nEnd))
	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/Tasks.GetHostStatusRecordRange", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ts.client.Do(ctx, request, nil)
	return raw, err
}

//TODO ResolveTaskId

//	Get task id by PRTS task id.
//
//	Parameters:
//	- strPrtsTaskId	(wstring) PRTS task id
//
//	Returns:
//	- (string) task id
func (ts *Tasks) ResolveTaskId(ctx context.Context, strPrtsTaskId string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"strPrtsTaskId": "%s"}`, strPrtsTaskId))
	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/Tasks.ResolveTaskId", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ts.client.Do(ctx, request, nil)
	return raw, err
}

//TODO - GetNextTask

//	Sequentially get task data.
//
//	Gets result of operation Tasks.ResetTasksIterator
//
//	Parameters:
//	- strTaskIteratorId	(string) iterator id got from Tasks.ResetTasksIterator
//
//	Return:
//	- pTaskData	(params) task data with attributes from Task settings format.
//	Call this method while pTaskData is not empty.
func (ts *Tasks) GetNextTask(ctx context.Context, strTaskIteratorId string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"strTaskIteratorId": "%s"}`, strTaskIteratorId))
	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/Tasks.GetNextTask", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ts.client.Do(ctx, request, nil)
	return raw, err
}

//TODO - GetNextHostStatus

//	Sequentially get result of operation ResetHostIteratorForTaskStatus or ResetHostIteratorForTaskStatusEx.
//
//	Sequentially gets result of operation Tasks.ResetHostIteratorForTaskStatus or Tasks.ResetHostIteratorForTaskStatusEx
//
//	Parameters:
//	- strHostIteratorId	(string) iterator id which got from
//	Tasks.ResetHostIteratorForTaskStatus or Tasks.ResetHostIteratorForTaskStatusEx
//	- nCount	(int) requested number of records
//
//	- [out]	nActual	(int) actual count of received records
//	- [out]	pHostStatus	(params) container that has requested elements in the array with name "statuses",
//	each item of array contains attributes from Host task state attributes
//
//	Returns:
//	- (bool) return false if the iterator reached end of the record list, in this case nActual contains zero
func (ts *Tasks) GetNextHostStatus(ctx context.Context, strTaskIteratorId string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"strTaskIteratorId": "%s"}`, strTaskIteratorId))
	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/Tasks.GetNextHostStatus", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ts.client.Do(ctx, request, nil)
	return raw, err
}

//TODO func (ts *Tasks) AddTask(ctx context.Context, params interface{}) ([]byte, error)
