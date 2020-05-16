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
	"net/http"
)

//	GroupTaskControlApi Class Reference
//
//	Interface to perform some management actions over group tasks..

//	List of all members.
type GroupTaskControlApi service

//	Completes import of task.
//
//	Performs actual import of task, which was initiated earlier by GroupTaskControlApi.ImportTask method. This method works with task data, which was already processed and analyzed. At time of calling this method, one should have appropriate rights for task creation. In case of group task, there should be write access to that group. If imported task is task for specified hosts, there should be read access to that hosts. If task is query-based task, there should be read access to related query. For additional info on task import options, see pExtraData parameter of GroupTaskControlApi.ImportTask method.
//
//	Parameters:
//	- wstrId	(string) Identifier of blob with task data. Use GroupTaskControlApi.ImportTask method to get its value
//	- bCommit	(bool) Whether to commit or not. If bCommit is true,
//	method will perform its work and complete import operation by creating new task and returning its identifier. If bCommit is false, method will cleanup old data and import will be cancelled.
//
//	Returns:
//	- (string) Unique identifier of new task if bCommit was set to true, and empty string otherwise
//
//	Exceptions:
//	- KLERR.Error*	Method must throw KLERR.Error* exception with code KLSTD.STDE_NOACCESS. For possible reasons,
//	see detailed method description above
//	- KLERR.Error*	Method must throw KLERR.Error* exception with code KLSTD.STDE_NOTFOUND in case object with
//	identifier wstrId not found: either identifier is not correct (check that it identifier string, which was returned by GroupTaskControlApi.ImportTask method), or object lifetime reached limit
//	- KLERR.Error*	Method must throw KLERR.Error* exception with code KLSTD.STDE_NOMEMORY in case of insufficient
//	memory for storing required task data
func (gta *GroupTaskControlApi) CommitImportedTask(ctx context.Context, wstrId string, bCommit bool) (*TaskDescribe,
	[]byte,
	error) {
	postData := []byte(fmt.Sprintf(`{"wstrId": "%s", "bCommit": %v}`, wstrId, bCommit))
	request, err := http.NewRequest("POST", gta.client.Server+"/api/v1.0/GroupTaskControlApi.CommitImportedTask",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	taskDescribe := new(TaskDescribe)
	raw, err := gta.client.Do(ctx, request, &taskDescribe)
	return taskDescribe, raw, err
}

//TasksIDSParams struct
type TasksIDSParams struct {
	PTasksIDS []string `json:"pTasksIds"`
}

//	Request statistics of the given tasks.
//
//	Actual statistics for the tasks will be reported by appropriate "KLEVP_EventGroupTaskStats" events publications,
//	event parameters:
//
//	- KLTSK_GRP_TSK_ID, paramString - task db id as a string
//	- KLTSK_GRP_TSK_STATS_PARAMS, paramParams - the task's statistics,
//	see List of task statistics attributes
//
//	Parameters:
//	- pTasksIds	[in] (array) - array of the tasks identifiers, each item is paramString
//
//	Note:
//	- to get task ids you can use Tasks.GetAllTasksOfHost
func (gta *GroupTaskControlApi) RequestStatistics(ctx context.Context, params TasksIDSParams) ([]byte, error) {
	postData, err := json.Marshal(&params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", gta.client.Server+"/api/v1.0/GroupTaskControlApi.RequestStatistics", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := gta.client.Do(ctx, request, nil)
	return raw, err
}

//	Exports task.
//
//	Gets specific task by its identifier and save data to memory chunk. Chunk can be later saved to file or sent over network
//
//	Possible errors:
//
//	- Task with specified id does not exist
//	- Read access to task with specified id is denied
//	- Not enough memory to store result chunk
//
//	Parameters:
//		wstrTaskId	(string) Task identifier
//
//	Returns:
//	- (binary) Pointer to memory chunk with exported task data
func (gta *GroupTaskControlApi) ExportTask(ctx context.Context, wstrTaskId string) (*PxgValStr, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"wstrTaskId": "%s"}`, wstrTaskId))
	request, err := http.NewRequest("POST", gta.client.Server+"/api/v1.0/GroupTaskControlApi.ExportTask",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := gta.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

//TaskDescribe struct
type TaskDescribe struct {
	TaskValue *TaskValue `json:"PxgRetVal,omitempty"`
}

type TaskValue struct {
	EventType                    *string                    `json:"EVENT_TYPE,omitempty"`
	FilterEventsComponentName    *string                    `json:"FILTER_EVENTS_COMPONENT_NAME,omitempty"`
	FilterEventsInstanceID       *string                    `json:"FILTER_EVENTS_INSTANCE_ID,omitempty"`
	FilterEventsProductName      *string                    `json:"FILTER_EVENTS_PRODUCT_NAME,omitempty"`
	FilterEventsVersion          *string                    `json:"FILTER_EVENTS_VERSION,omitempty"`
	TaskidComponentName          *string                    `json:"TASKID_COMPONENT_NAME,omitempty"`
	TaskidInstanceID             *string                    `json:"TASKID_INSTANCE_ID,omitempty"`
	TaskidProductName            *string                    `json:"TASKID_PRODUCT_NAME,omitempty"`
	TaskidVersion                *string                    `json:"TASKID_VERSION,omitempty"`
	TaskschEwDay                 *int64                     `json:"TASKSCH_EW_DAY,omitempty"`
	TaskschEwHours               *int64                     `json:"TASKSCH_EW_HOURS,omitempty"`
	TaskschEwMins                *int64                     `json:"TASKSCH_EW_MINS,omitempty"`
	TaskschEwSecs                *int64                     `json:"TASKSCH_EW_SECS,omitempty"`
	TaskschFirstExecutionTime    *TaskschFirstExecutionTime `json:"TASKSCH_FIRST_EXECUTION_TIME,omitempty"`
	TaskschFirstExecutionTimeSEC *int64                     `json:"TASKSCH_FIRST_EXECUTION_TIME_SEC,omitempty"`
	TaskschLifetime              *TaskschFirstExecutionTime `json:"TASKSCH_LIFETIME,omitempty"`
	TaskschMSPeriod              *int64                     `json:"TASKSCH_MS_PERIOD,omitempty"`
	TaskschRunMissedFlag         *bool                      `json:"TASKSCH_RUN_MISSED_FLAG,omitempty"`
	TaskschType                  *int64                     `json:"TASKSCH_TYPE,omitempty"`
	TaskAdditionalParams         *TaskAdditionalParams      `json:"TASK_ADDITIONAL_PARAMS,omitempty"`
	TaskClassID                  *int64                     `json:"TASK_CLASS_ID,omitempty"`
	TaskDelAfterRunFlag          *bool                      `json:"TASK_DEL_AFTER_RUN_FLAG,omitempty"`
	TaskInfoParams               *TaskInfoParams            `json:"TASK_INFO_PARAMS,omitempty"`
	TaskLastExecTime             *TaskschFirstExecutionTime `json:"TASK_LAST_EXEC_TIME,omitempty"`
	TaskLastExecTimeSEC          *int64                     `json:"TASK_LAST_EXEC_TIME_SEC,omitempty"`
	TaskMaxExecTime              *int64                     `json:"TASK_MAX_EXEC_TIME,omitempty"`
	TaskName                     *string                    `json:"TASK_NAME,omitempty"`
	TaskPrepStart                *int64                     `json:"TASK_PREP_START,omitempty"`
	TaskPriority                 *int64                     `json:"TASK_PRIORITY,omitempty"`
	TaskStartDelta               *int64                     `json:"TASK_START_DELTA,omitempty"`
	TaskUniqueID                 *string                    `json:"TASK_UNIQUE_ID,omitempty"`
}

type TaskAdditionalParams struct {
	Type  *string                    `json:"type,omitempty"`
	Value *TASKADDITIONALPARAMSValue `json:"value,omitempty"`
}

type TASKADDITIONALPARAMSValue struct {
	KlnagTskVapmsearchPaths []string `json:"KLNAG_TSK_VAPMSEARCH_PATHS"`
	KlprtsTaskStorageID     *string  `json:"klprts-TaskStorageId,omitempty"`
	NSource                 *int64   `json:"nSource,omitempty"`
}

type TaskInfoParams struct {
	Type  *string              `json:"type,omitempty"`
	Value *TASKINFOPARAMSValue `json:"value,omitempty"`
}

type TASKINFOPARAMSValue struct {
	DisplayName                   *string                    `json:"DisplayName,omitempty"`
	KlevpNotificationDescrID      *string                    `json:"KLEVP_NOTIFICATION_DESCR_ID,omitempty"`
	KlhstWksCtype                 *int64                     `json:"KLHST_WKS_CTYPE,omitempty"`
	KLPRSSEVPNotifications        *KLPRSSEVPNotifications    `json:"KLPRSS_EVPNotifications,omitempty"`
	KlsrvPrtsTaskEnabledFlag      *bool                      `json:"KLSRV_PRTS_TASK_ENABLED_FLAG,omitempty"`
	KltskAllowAutoRandomization   *bool                      `json:"KLTSK_ALLOW_AUTO_RANDOMIZATION,omitempty"`
	NhTaskCreatedByQsw            *bool                      `json:"NH_TASK_CREATED_BY_QSW,omitempty"`
	PrtsExceptGroupids            []interface{}              `json:"PRTS_EXCEPT_GROUPIDS"`
	PrtsTaskCreationDate          *TaskschFirstExecutionTime `json:"PRTS_TASK_CREATION_DATE,omitempty"`
	PrtsTaskEnabled               *bool                      `json:"PRTS_TASK_ENABLED,omitempty"`
	PrtsTaskGroupid               *int64                     `json:"PRTS_TASK_GROUPID,omitempty"`
	PrtsTaskGroupname             *string                    `json:"PRTS_TASK_GROUPNAME,omitempty"`
	KlprtsDontApplyToSlaveServers *bool                      `json:"klprts-DontApplyToSlaveServers,omitempty"`
	KlprtsTaskScheduleSubtype     *int64                     `json:"klprts-TaskScheduleSubtype,omitempty"`
}

type KLPRSSEVPNotifications struct {
	Type  *string                      `json:"type,omitempty"`
	Value *KLPRSSEVPNotificationsValue `json:"value,omitempty"`
}

type KLPRSSEVPNotificationsValue struct {
	Err []Err `json:"ERR"`
	Inf []Err `json:"INF"`
	Wrn []Err `json:"WRN"`
}

type Err struct {
	Type  *string   `json:"type,omitempty"`
	Value *ERRValue `json:"value,omitempty"`
}

type ERRValue struct {
	KlevpNdDaysToStoreEvent  *int64             `json:"KLEVP_ND_DAYS_TO_STORE_EVENT,omitempty"`
	KlevpNdEvetnType         *string            `json:"KLEVP_ND_EVETN_TYPE,omitempty"`
	KlevpNdStoreAtClientLog  *bool              `json:"KLEVP_ND_STORE_AT_CLIENT_LOG,omitempty"`
	KlevpNdStoreAtClientPres *bool              `json:"KLEVP_ND_STORE_AT_CLIENT_PRES,omitempty"`
	KlevpNdStoreAtServerLog  *bool              `json:"KLEVP_ND_STORE_AT_SERVER_LOG,omitempty"`
	KlevpNdBodyFilter        *KlevpNdBodyFilter `json:"KLEVP_ND_BODY_FILTER,omitempty"`
}

type KlevpNdBodyFilter struct {
	Type  *string                 `json:"type,omitempty"`
	Value *KLEVPNDBODYFILTERValue `json:"value,omitempty"`
}

type KLEVPNDBODYFILTERValue struct {
	KLPRCINewState *int64 `json:"KLPRCI_newState,omitempty"`
}

type TaskschFirstExecutionTime struct {
	Type  *string `json:"type,omitempty"`
	Value *string `json:"value,omitempty"`
}

//	GetTaskByRevision - get the task data by revision.
//
//	Returns all task data for a group/set task with a given object identity and revision.
//
//	Parameters:
//	- nObjId	(int64) Task identity.
//	- nRevision	(int64) Task revision id, zero value means 'current task'.
//
//	If Administration Server version is less than "SC 10 SP2 MR1" then nRevision must be zero.
//
//	Returns:
//	- (params) describing the task, see Task settings format
func (gta *GroupTaskControlApi) GetTaskByRevision(ctx context.Context, nObjId, nRevision int64) (*TaskDescribe, []byte,
	error) {
	postData := []byte(fmt.Sprintf(`{"nObjId": %d, "nRevision": %d}`, nObjId, nRevision))
	request, err := http.NewRequest("POST", gta.client.Server+"/api/v1.0/GroupTaskControlApi.GetTaskByRevision",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	taskDescribe := new(TaskDescribe)
	raw, err := gta.client.Do(ctx, request, &taskDescribe)
	return taskDescribe, raw, err
}

//	Restore task from revision.
//
//	Rolls back the group/set task specified by nObjId to the revision nRevision.
//
//	Parameters:
//	- nObjId	(int64) Task identity.
//	- nRevision	(int64) Task revision id, value cannot be zero.
//
//	Note:
//	ExtAud interface allow you to get a revision of an object and update description.
func (gta *GroupTaskControlApi) RestoreTaskFromRevision(ctx context.Context, nObjId, nRevision int64) (*TaskDescribe, []byte,
	error) {
	postData := []byte(fmt.Sprintf(`{"nObjId": %d, "nRevision": %d}`, nObjId, nRevision))
	request, err := http.NewRequest("POST", gta.client.Server+"/api/v1.0/GroupTaskControlApi.RestoreTaskFromRevision",
		bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	taskDescribe := new(TaskDescribe)
	raw, err := gta.client.Do(ctx, request, &taskDescribe)
	return taskDescribe, raw, err
}

//TODO ImportTask
//TODO ResetTasksIteratorForCluster
