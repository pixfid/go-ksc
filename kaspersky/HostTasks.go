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
	"fmt"
	"net/http"
)

//	HostTasks Class Reference
//
//	Basic management operations with host tasks..
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
