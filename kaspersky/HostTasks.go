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
	"net/http"
)

//	HostTasks Class Reference
//
//	Basic management operations with host tasks. More...
//
//	Interface allows to acquire and manage tasks for hosts: add, update, remove, enumerate and perform other actions.
//
//	List of all members.
type HostTasks service

func (ht *HostTasks) GetNextTask(ctx context.Context, strSrvObjId string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`
	{
	"strSrvObjId": "%s"
	}`, strSrvObjId))
	request, err := http.NewRequest("POST", ht.client.Server+"/api/v1.0/HostTasks.GetNextTask", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ht.client.Do(ctx, request, nil)
	return raw, err
}

//Reset task iterator for a specified filter data.
//
//If one of the parameters is not specified then the filtration will not be performed by this parameter.
//
//	Parameters:
//	- strSrvObjId	(string) server object ID that got from HostGroup.GetHostTasks
//	- strProductName	(string) product name
//	- strVersion	(string) product version
//	- strComponentName	(string) component name
//	- strInstanceId	(string) instance id
//	- strTaskName	(string) task name
func (ht *HostTasks) ResetTasksIterator(ctx context.Context, strSrvObjId, strProductName, strVersion,
	strComponentName, strInstanceId, strTaskName string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{
		"strSrvObjId": "%s",
		"strProductName": "%s",
		"strVersion": "%s",
		"strComponentName":	"%s",
		"strInstanceId": "%s",
		"strTaskName": "%s"
	}`, strSrvObjId, strProductName, strVersion, strComponentName, strInstanceId, strTaskName))
	request, err := http.NewRequest("POST", ht.client.Server+"/api/v1.0/HostTasks.ResetTasksIterator", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ht.client.Do(ctx, request, nil)
	return raw, err
}
