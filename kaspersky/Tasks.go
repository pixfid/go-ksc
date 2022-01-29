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

// Tasks Group tasks service allows to acquire task attributes, enumerate, control and delete tasks.
type Tasks service

// GetAllTasksOfHost Get all group and global tasks of specified host.
func (ts *Tasks) GetAllTasksOfHost(ctx context.Context, strDomainName, strHostName string) (*PxgValArrayOfString,
	[]byte, error) {
	postData := []byte(fmt.Sprintf(`{"strDomainName": "%s","strHostName": "%s"}`, strDomainName, strHostName))
	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/Tasks.GetAllTasksOfHost", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValArrayOfString := new(PxgValArrayOfString)
	raw, err := ts.client.Request(ctx, request, &pxgValArrayOfString)
	return pxgValArrayOfString, raw, err
}

// TaskData struct
type TaskData struct {
	PxgRetVal Task `json:"PxgRetVal"`
}

// Task struct
type Task struct {
	DisplayName          string               `json:"DisplayName"`
	PrtsTaskCreationDate PrtsTaskCreationDate `json:"PRTS_TASK_CREATION_DATE"`
	TaskidProductName    string               `json:"TASKID_PRODUCT_NAME"`
	TaskidVersion        string               `json:"TASKID_VERSION"`
	TaskName             string               `json:"TASK_NAME"`
	TaskUniqueID         string               `json:"TASK_UNIQUE_ID"`
}

// PrtsTaskCreationDate struct
type PrtsTaskCreationDate struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// GetTask Acquire attributes of specified task.
func (ts *Tasks) GetTask(ctx context.Context, strTask string) (*TaskData, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"strTask": "%s"}`, strTask))
	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/Tasks.GetTask", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	taskData := new(TaskData)
	raw, err := ts.client.Request(ctx, request, &taskData)
	return taskData, raw, err
}

// GetTaskData Acquire task settings.
func (ts *Tasks) GetTaskData(ctx context.Context, strTask string, tsk interface{}) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`
	{
	"strTask": "%s"
	}`, strTask))
	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/Tasks.GetTaskData", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ts.client.Request(ctx, request, &tsk)
	return raw, err
}

// GetTaskGroup Return the group id for the group task.
func (ts *Tasks) GetTaskGroup(ctx context.Context, strTaskId string) (*PxgValInt, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"strTaskId": "%s"}`, strTaskId))
	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/Tasks.GetTaskGroup", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValInt := new(PxgValInt)
	raw, err := ts.client.Request(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

//TaskStatistics struct
type TaskStatistics struct {
	TaskStatistic TaskStatistic `json:"PxgRetVal"`
}

// TaskStatistic struct
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

// GetTaskStatistics Acquire statistics of the specified task.
func (ts *Tasks) GetTaskStatistics(ctx context.Context, strTask string) (*TaskStatistics, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"strTask": "%s"}`, strTask))
	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/Tasks.GetTaskStatistics", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	taskStatistics := new(TaskStatistics)
	raw, err := ts.client.Request(ctx, request, &taskStatistics)
	return taskStatistics, raw, err
}

// SuspendTask Suspend execution of the specified task.
func (ts *Tasks) SuspendTask(ctx context.Context, strTask string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"strTask": "%s"}`, strTask))
	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/Tasks.SuspendTask", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ts.client.Request(ctx, request, nil)
	return raw, err
}

// ResumeTask Resumes specified task
func (ts *Tasks) ResumeTask(ctx context.Context, strTask string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"strTask": "%s"}`, strTask))
	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/Tasks.ResumeTask", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ts.client.Request(ctx, request, nil)
	return raw, err
}

// RunTask Start specified task.
func (ts *Tasks) RunTask(ctx context.Context, strTask string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"strTask": "%s"}`, strTask))
	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/Tasks.RunTask", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ts.client.Request(ctx, request, nil)
	return raw, err
}

// DeleteTask Deletes the specified task.
func (ts *Tasks) DeleteTask(ctx context.Context, strTask string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"strTask": "%s"}`, strTask))
	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/Tasks.DeleteTask", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ts.client.Request(ctx, request, nil)
	return raw, err
}

// CancelTask Cancels execution of the specified task.
func (ts *Tasks) CancelTask(ctx context.Context, strTask string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"strTask": "%s"}`, strTask))
	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/Tasks.CancelTask", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ts.client.Request(ctx, request, nil)
	return raw, err
}

// TaskHistoryParams struct
type TaskHistoryParams struct {
	StrTask        string          `json:"strTask"`
	PFields2Return []string        `json:"pFields2Return"`
	PSortFields    []FieldsToOrder `json:"pSortFields"`
	StrHostName    string          `json:"strHostName"`
	PFilter        interface{}     `json:"pFilter"`
}

// GetTaskHistory Acquire task execution history events.
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
	raw, err := ts.client.Request(ctx, request, &strIteratorID)
	return strIteratorID, raw, err
}

// GetTaskStartEvent Returns event which should run the task.
func (ts *Tasks) GetTaskStartEvent(ctx context.Context, strTask string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"strTask": "%s"}`, strTask))
	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/Tasks.GetTaskStartEvent", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ts.client.Request(ctx, request, nil)
	return raw, err
}

// ProtectPassword Encrypt an account password.
func (ts *Tasks) ProtectPassword(ctx context.Context, strPassword string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"strPassword": "%s"}`, strPassword))
	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/Tasks.ProtectPassword", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ts.client.Request(ctx, request, nil)
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

// ResetTasksIterator Reset task iterator for a specified filter data.
//
// If one of the parameters is not specified then the filtration will not be performed by this parameter.
//
// The group Super is parent of the group Groups and is intended for assignment of group tasks and policies received
// from the master server.
func (ts *Tasks) ResetTasksIterator(ctx context.Context, params TasksIteratorParams) ([]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/Tasks.ResetTasksIterator", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ts.client.Request(ctx, request, nil)
	return raw, err
}

// ReleaseTasksIterator Release task iterator.
func (ts *Tasks) ReleaseTasksIterator(ctx context.Context, strTaskIteratorId string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"strTaskIteratorId": "%s"}`, strTaskIteratorId))
	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/Tasks.ReleaseTasksIterator", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ts.client.Request(ctx, request, nil)
	return raw, err
}

// ReleaseHostStatusIterator Releases iterator of specified data and frees associated memory
func (ts *Tasks) ReleaseHostStatusIterator(ctx context.Context, strHostIteratorId string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"strHostIteratorId": "%s"}`, strHostIteratorId))
	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/Tasks.ReleaseHostStatusIterator", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ts.client.Request(ctx, request, nil)
	return raw, err
}

//HostIteratorForTaskParams struct
type HostIteratorForTaskParams struct {
	StrTask        string   `json:"strTask"`
	NHostStateMask string   `json:"nHostStateMask"`
	PFields2Return []string `json:"pFields2Return"`
	NLifetime      int64    `json:"nLifetime"`
}

// ResetHostIteratorForTaskStatus Make host task states request.
// Makes request of the status of group task that runs on many machines
func (ts *Tasks) ResetHostIteratorForTaskStatus(ctx context.Context, params HostIteratorForTaskParams) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/Tasks.ResetHostIteratorForTaskStatus", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ts.client.Request(ctx, request, nil)
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

// ResetHostIteratorForTaskStatusEx Make host task states request.
//
// Example:
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
	raw, err := ts.client.Request(ctx, request, &strHostIteratorId)
	return strHostIteratorId, raw, err
}

// GetHostStatusRecordsCount Get records count of result of operation: ResetHostIteratorForTaskStatus or ResetHostIteratorForTaskStatusEx.
func (ts *Tasks) GetHostStatusRecordsCount(ctx context.Context, strHostIteratorId string) (*PxgValInt, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"strHostIteratorId": "%s"}`, strHostIteratorId))
	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/Tasks.GetHostStatusRecordsCount", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValInt := new(PxgValInt)
	raw, err := ts.client.Request(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

// GetHostStatusRecordRange Gets result of operation Tasks.ResetHostIteratorForTaskStatus or Tasks.ResetHostIteratorForTaskStatusEx
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

	raw, err := ts.client.Request(ctx, request, nil)
	return raw, err
}

// ResolveTaskId Get task id by PRTS task id.
func (ts *Tasks) ResolveTaskId(ctx context.Context, strPrtsTaskId string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"strPrtsTaskId": "%s"}`, strPrtsTaskId))
	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/Tasks.ResolveTaskId", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ts.client.Request(ctx, request, nil)
	return raw, err
}

// GetNextTask Sequentially get task data. Gets result of operation Tasks.ResetTasksIterator
func (ts *Tasks) GetNextTask(ctx context.Context, strTaskIteratorId string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"strTaskIteratorId": "%s"}`, strTaskIteratorId))
	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/Tasks.GetNextTask", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ts.client.Request(ctx, request, nil)
	return raw, err
}

// GetNextHostStatus Sequentially gets result of operation Tasks.ResetHostIteratorForTaskStatus or Tasks.ResetHostIteratorForTaskStatusEx
func (ts *Tasks) GetNextHostStatus(ctx context.Context, strTaskIteratorId string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"strTaskIteratorId": "%s"}`, strTaskIteratorId))
	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/Tasks.GetNextHostStatus", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ts.client.Request(ctx, request, nil)
	return raw, err
}

// AddTask Creates new task.
func (ts *Tasks) AddTask(ctx context.Context, params interface{}) (*PxgValInt, []byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/Tasks.AddTask", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValInt := new(PxgValInt)
	raw, err := ts.client.Request(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}
