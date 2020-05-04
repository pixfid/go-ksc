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
	"fmt"
	"log"
	"net/http"
)

//	Tasks Class Reference
//
//	Group tasks.
//
//	Allows to acquire task attributes, enumerate, control and delete tasks.
//
//	List of all members.
type Tasks struct {
	client *Client
}

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
	[]byte,
	error) {
	postData := []byte(fmt.Sprintf(`
	{
	"strDomainName": "%s",
	"strHostName": "%s"
	}`, strDomainName, strHostName))

	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/Tasks.GetAllTasksOfHost", bytes.NewBuffer(postData))

	if err != nil {
		log.Fatal(err.Error())
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
		log.Fatal(err.Error())
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
		log.Fatal(err.Error())
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
func (ts *Tasks) GetTaskGroup(ctx context.Context, strTaskId string, v interface{}) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`
	{
	"strTaskId": "%s"
	}`, strTaskId))

	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/Tasks.GetTaskGroup", bytes.NewBuffer(postData))

	if err != nil {
		log.Fatal(err.Error())
	}

	raw, err := ts.client.Do(ctx, request, &v)
	return raw, err
}

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
//(TaskStatistics) object containing task statistics, see List of task statistics attributes.
func (ts *Tasks) GetTaskStatistics(ctx context.Context, strTask string) (*TaskStatistics, []byte,
	error) {
	postData := []byte(fmt.Sprintf(`
	{
	"strTask": "%s"
	}`, strTask))

	tsks := new(TaskStatistics)
	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/Tasks.GetTaskStatistics", bytes.NewBuffer(postData))
	if err != nil {
		log.Fatal(err.Error())
	}

	raw, err := ts.client.Do(ctx, request, &tsks)
	return tsks, raw, err
}

//	Suspend execution of the specified task.
//
//	Suspends execution of the specified task.
//
//	Parameters:
//	- ctx	(context.Context) context.
//	- strTask	(string) task id.
func (ts *Tasks) SuspendTask(ctx context.Context, strTask string) ([]byte,
	error) {
	postData := []byte(fmt.Sprintf(`
	{
	"strTask": "%s"
	}`, strTask))

	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/Tasks.SuspendTask", bytes.NewBuffer(postData))
	if err != nil {
		log.Fatal(err.Error())
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
func (ts *Tasks) ResumeTask(ctx context.Context, strTask string) ([]byte,
	error) {
	postData := []byte(fmt.Sprintf(`
	{
	"strTask": "%s"
	}`, strTask))

	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/Tasks.ResumeTask", bytes.NewBuffer(postData))
	if err != nil {
		log.Fatal(err.Error())
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
func (ts *Tasks) RunTask(ctx context.Context, strTask string) ([]byte,
	error) {
	postData := []byte(fmt.Sprintf(`
	{
	"strTask": "%s"
	}`, strTask))

	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/Tasks.RunTask", bytes.NewBuffer(postData))
	if err != nil {
		log.Fatal(err.Error())
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
func (ts *Tasks) DeleteTask(ctx context.Context, strTask string) ([]byte,
	error) {
	postData := []byte(fmt.Sprintf(`
	{
	"strTask": "%s"
	}`, strTask))

	request, err := http.NewRequest("POST", ts.client.Server+"/api/v1.0/Tasks.DeleteTask", bytes.NewBuffer(postData))
	if err != nil {
		log.Fatal(err.Error())
	}

	raw, err := ts.client.Do(ctx, request, nil)
	return raw, err
}
