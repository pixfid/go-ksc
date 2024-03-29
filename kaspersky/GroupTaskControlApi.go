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

// GroupTaskControlApi service to perform some management actions over group tasks.
type GroupTaskControlApi service

// CommitImportedTask Completes import of task. Performs actual import of task, which was initiated earlier by
// GroupTaskControlApi.ImportTask method.
//
// This method works with task data, which was already processed and analyzed.
// At time of calling this method, one should have appropriate rights for task creation.
// In case of group task, there should be write access to that group.
// If imported task is task for specified hosts, there should be read access to that hosts.
// If task is query-based task, there should be read access to related query.
// For additional info on task import options, see pExtraData parameter of GroupTaskControlApi.ImportTask method.
func (gtca *GroupTaskControlApi) CommitImportedTask(ctx context.Context, wstrId string, bCommit bool) (*TaskDescribe,
	[]byte,
	error) {
	postData := []byte(fmt.Sprintf(`{"wstrId": "%s", "bCommit": %v}`, wstrId, bCommit))
	request, err := http.NewRequest("POST", gtca.client.Server+"/api/v1.0/GroupTaskControlApi.CommitImportedTask",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	taskDescribe := new(TaskDescribe)
	raw, err := gtca.client.Request(ctx, request, &taskDescribe)
	return taskDescribe, raw, err
}

// TasksIDSParams struct
type TasksIDSParams struct {
	PTasksIDS []string `json:"pTasksIds"`
}

// RequestStatistics of the given tasks.
// Actual statistics for the tasks will be reported by appropriate "KLEVP_EventGroupTaskStats" events publications.
func (gtca *GroupTaskControlApi) RequestStatistics(ctx context.Context, params TasksIDSParams) ([]byte, error) {
	postData, err := json.Marshal(&params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", gtca.client.Server+"/api/v1.0/GroupTaskControlApi.RequestStatistics", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := gtca.client.Request(ctx, request, nil)
	return raw, err
}

// ExportTask Gets specific task by its identifier and save data to memory chunk.
// Chunk can be later saved to file or sent over network
func (gtca *GroupTaskControlApi) ExportTask(ctx context.Context, wstrTaskId string) (*PxgValStr, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"wstrTaskId": "%s"}`, wstrTaskId))
	request, err := http.NewRequest("POST", gtca.client.Server+"/api/v1.0/GroupTaskControlApi.ExportTask",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := gtca.client.Request(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

// TaskDescribe struct
type TaskDescribe struct {
	TaskValue *TaskValue `json:"PxgRetVal,omitempty"`
}

// TaskValue struct
type TaskValue struct {
	EventType                    string                     `json:"EVENT_TYPE,omitempty"`
	FilterEventsComponentName    string                     `json:"FILTER_EVENTS_COMPONENT_NAME,omitempty"`
	FilterEventsInstanceID       string                     `json:"FILTER_EVENTS_INSTANCE_ID,omitempty"`
	FilterEventsProductName      string                     `json:"FILTER_EVENTS_PRODUCT_NAME,omitempty"`
	FilterEventsVersion          string                     `json:"FILTER_EVENTS_VERSION,omitempty"`
	TaskidComponentName          string                     `json:"TASKID_COMPONENT_NAME,omitempty"`
	TaskidInstanceID             string                     `json:"TASKID_INSTANCE_ID,omitempty"`
	TaskidProductName            string                     `json:"TASKID_PRODUCT_NAME,omitempty"`
	TaskidVersion                string                     `json:"TASKID_VERSION,omitempty"`
	TaskschEwDay                 int64                      `json:"TASKSCH_EW_DAY,omitempty"`
	TaskschEwHours               int64                      `json:"TASKSCH_EW_HOURS,omitempty"`
	TaskschEwMins                int64                      `json:"TASKSCH_EW_MINS,omitempty"`
	TaskschEwSecs                int64                      `json:"TASKSCH_EW_SECS,omitempty"`
	TaskschFirstExecutionTime    *TaskschFirstExecutionTime `json:"TASKSCH_FIRST_EXECUTION_TIME,omitempty"`
	TaskschFirstExecutionTimeSEC int64                      `json:"TASKSCH_FIRST_EXECUTION_TIME_SEC,omitempty"`
	TaskschLifetime              *TaskschFirstExecutionTime `json:"TASKSCH_LIFETIME,omitempty"`
	TaskschMSPeriod              int64                      `json:"TASKSCH_MS_PERIOD,omitempty"`
	TaskschRunMissedFlag         bool                       `json:"TASKSCH_RUN_MISSED_FLAG,omitempty"`
	TaskschType                  int64                      `json:"TASKSCH_TYPE,omitempty"`
	TaskAdditionalParams         *TaskAdditionalParams      `json:"TASK_ADDITIONAL_PARAMS,omitempty"`
	TaskClassID                  int64                      `json:"TASK_CLASS_ID,omitempty"`
	TaskDelAfterRunFlag          bool                       `json:"TASK_DEL_AFTER_RUN_FLAG,omitempty"`
	TaskInfoParams               *TaskInfoParams            `json:"TASK_INFO_PARAMS,omitempty"`
	TaskLastExecTime             *TaskschFirstExecutionTime `json:"TASK_LAST_EXEC_TIME,omitempty"`
	TaskLastExecTimeSEC          int64                      `json:"TASK_LAST_EXEC_TIME_SEC,omitempty"`
	TaskMaxExecTime              int64                      `json:"TASK_MAX_EXEC_TIME,omitempty"`
	TaskName                     string                     `json:"TASK_NAME,omitempty"`
	TaskPrepStart                int64                      `json:"TASK_PREP_START,omitempty"`
	TaskPriority                 int64                      `json:"TASK_PRIORITY,omitempty"`
	TaskStartDelta               int64                      `json:"TASK_START_DELTA,omitempty"`
	TaskUniqueID                 string                     `json:"TASK_UNIQUE_ID,omitempty"`
}

type TaskAdditionalParams struct {
	Type  string                     `json:"type,omitempty"`
	Value *TASKADDITIONALPARAMSValue `json:"value,omitempty"`
}

type TASKADDITIONALPARAMSValue struct {
	KlnagTskVapmsearchPaths []string `json:"KLNAG_TSK_VAPMSEARCH_PATHS"`
	KlprtsTaskStorageID     string   `json:"klprts-TaskStorageId,omitempty"`
	NSource                 int64    `json:"nSource,omitempty"`
}

type TaskInfoParams struct {
	Type  string               `json:"type,omitempty"`
	Value *TASKINFOPARAMSValue `json:"value,omitempty"`
}

type TASKINFOPARAMSValue struct {
	DisplayName                   string                     `json:"DisplayName,omitempty"`
	KlevpNotificationDescrID      string                     `json:"KLEVP_NOTIFICATION_DESCR_ID,omitempty"`
	KlhstWksCtype                 int64                      `json:"KLHST_WKS_CTYPE,omitempty"`
	KLPRSSEVPNotifications        *KLPRSSEVPNotifications    `json:"KLPRSS_EVPNotifications,omitempty"`
	KlsrvPrtsTaskEnabledFlag      bool                       `json:"KLSRV_PRTS_TASK_ENABLED_FLAG,omitempty"`
	KltskAllowAutoRandomization   bool                       `json:"KLTSK_ALLOW_AUTO_RANDOMIZATION,omitempty"`
	NhTaskCreatedByQsw            bool                       `json:"NH_TASK_CREATED_BY_QSW,omitempty"`
	PrtsExceptGroupids            []interface{}              `json:"PRTS_EXCEPT_GROUPIDS"`
	PrtsTaskCreationDate          *TaskschFirstExecutionTime `json:"PRTS_TASK_CREATION_DATE,omitempty"`
	PrtsTaskEnabled               bool                       `json:"PRTS_TASK_ENABLED,omitempty"`
	PrtsTaskGroupid               int64                      `json:"PRTS_TASK_GROUPID,omitempty"`
	PrtsTaskGroupname             string                     `json:"PRTS_TASK_GROUPNAME,omitempty"`
	KlprtsDontApplyToSlaveServers bool                       `json:"klprts-DontApplyToSlaveServers,omitempty"`
	KlprtsTaskScheduleSubtype     int64                      `json:"klprts-TaskScheduleSubtype,omitempty"`
}

type KLPRSSEVPNotifications struct {
	Type  string                       `json:"type,omitempty"`
	Value *KLPRSSEVPNotificationsValue `json:"value,omitempty"`
}

type KLPRSSEVPNotificationsValue struct {
	Err []Err `json:"ERR"`
	Inf []Err `json:"INF"`
	Wrn []Err `json:"WRN"`
}

type Err struct {
	Type  string    `json:"type,omitempty"`
	Value *ERRValue `json:"value,omitempty"`
}

type ERRValue struct {
	KlevpNdDaysToStoreEvent  int64              `json:"KLEVP_ND_DAYS_TO_STORE_EVENT,omitempty"`
	KlevpNdEvetnType         string             `json:"KLEVP_ND_EVETN_TYPE,omitempty"`
	KlevpNdStoreAtClientLog  bool               `json:"KLEVP_ND_STORE_AT_CLIENT_LOG,omitempty"`
	KlevpNdStoreAtClientPres bool               `json:"KLEVP_ND_STORE_AT_CLIENT_PRES,omitempty"`
	KlevpNdStoreAtServerLog  bool               `json:"KLEVP_ND_STORE_AT_SERVER_LOG,omitempty"`
	KlevpNdBodyFilter        *KlevpNdBodyFilter `json:"KLEVP_ND_BODY_FILTER,omitempty"`
}

type KlevpNdBodyFilter struct {
	Type  string                  `json:"type,omitempty"`
	Value *KLEVPNDBODYFILTERValue `json:"value,omitempty"`
}

type KLEVPNDBODYFILTERValue struct {
	KLPRCINewState int64 `json:"KLPRCI_newState,omitempty"`
}

type TaskschFirstExecutionTime struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

// GetTaskByRevision get the task data by revision.
// Returns all task data for a group/set task with a given object identity and revision.
//
// If Administration Server version is less than "SC 10 SP2 MR1" then nRevision must be zero.
func (gtca *GroupTaskControlApi) GetTaskByRevision(ctx context.Context, nObjId, nRevision int64) (*TaskDescribe, []byte,
	error) {
	postData := []byte(fmt.Sprintf(`{"nObjId": %d, "nRevision": %d}`, nObjId, nRevision))
	request, err := http.NewRequest("POST", gtca.client.Server+"/api/v1.0/GroupTaskControlApi.GetTaskByRevision",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	taskDescribe := new(TaskDescribe)
	raw, err := gtca.client.Request(ctx, request, &taskDescribe)
	return taskDescribe, raw, err
}

// RestoreTaskFromRevision Restore task from revision. Rolls back the group/set task specified by nObjId to the revision nRevision.
func (gtca *GroupTaskControlApi) RestoreTaskFromRevision(ctx context.Context, nObjId, nRevision int64) (*TaskDescribe, []byte,
	error) {
	postData := []byte(fmt.Sprintf(`{"nObjId": %d, "nRevision": %d}`, nObjId, nRevision))
	request, err := http.NewRequest("POST", gtca.client.Server+"/api/v1.0/GroupTaskControlApi.RestoreTaskFromRevision",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	taskDescribe := new(TaskDescribe)
	raw, err := gtca.client.Request(ctx, request, &taskDescribe)
	return taskDescribe, raw, err
}

// ImportTask Prepares task import.
//
// Prepares task import operation. This method does not perform actual import of task: instead, it prepares import operation
// by saving task blob and associating unique identifier with it, which should be later used in GroupTaskControlApi.CommitImportedTask method.
//
// Some tasks have import security restrictions, thus import of task can be allowed or denied.
//
// To determine these restrictions, one should analyze info, returned via output parameter pCommitInfo
// (see detailed parameter descriptions in parameter section), and pass analyze result in bCommit parameter of GroupTaskControlApi.CommitImportedTask.
func (gtca *GroupTaskControlApi) ImportTask(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(&params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", gtca.client.Server+"/api/v1.0/GroupTaskControlApi.ImportTask", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := gtca.client.Request(ctx, request, nil)
	return raw, err
}

// ResetIterForClusterParams struct
type ResetIterForClusterParams struct {
	SzwClusterID     string `json:"szwClusterId,omitempty"`
	SzwProductName   string `json:"szwProductName,omitempty"`
	SzwVersion       string `json:"szwVersion,omitempty"`
	SzwComponentName string `json:"szwComponentName,omitempty"`
	SzwInstanceID    string `json:"szwInstanceId,omitempty"`
	SzwTaskName      string `json:"szwTaskName,omitempty"`
}

// ResetTasksIteratorForCluster Reset task iterator for a cluster.
func (gtca *GroupTaskControlApi) ResetTasksIteratorForCluster(ctx context.Context, params ResetIterForClusterParams) ([]byte, error) {
	postData, err := json.Marshal(&params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", gtca.client.Server+"/api/v1.0/GroupTaskControlApi.ResetTasksIteratorForCluster", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := gtca.client.Request(ctx, request, nil)
	return raw, err
}
