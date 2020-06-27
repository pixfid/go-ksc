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

// DatabaseInfo service to Database processing. Allow to get information from a Database.
type DatabaseInfo service

// GetDBSize Get database's files size.
func (di *DatabaseInfo) GetDBSize(ctx context.Context) (*PxgValInt, []byte, error) {
	request, err := http.NewRequest("POST", di.client.Server+"/api/v1.0/DatabaseInfo.GetDBSize", nil)
	if err != nil {
		return nil, nil, err
	}

	pxgValInt := new(PxgValInt)
	raw, err := di.client.Do(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

// GetDBDataSize Get database's data size.
func (di *DatabaseInfo) GetDBDataSize(ctx context.Context) (*PxgValInt, []byte, error) {
	request, err := http.NewRequest("POST", di.client.Server+"/api/v1.0/DatabaseInfo.GetDBDataSize", nil)
	if err != nil {
		return nil, nil, err
	}

	pxgValInt := new(PxgValInt)
	raw, err := di.client.Do(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

// GetDBEventsCount Get database's events count.
func (di *DatabaseInfo) GetDBEventsCount(ctx context.Context) (*PxgValInt, []byte, error) {
	request, err := http.NewRequest("POST", di.client.Server+"/api/v1.0/DatabaseInfo.GetDBEventsCount", nil)
	if err != nil {
		return nil, nil, err
	}

	pxgValInt := new(PxgValInt)
	raw, err := di.client.Do(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

// IsCloudSQL Check is current SQL server in cloud (Amazon RDS or Azure SQL)
func (di *DatabaseInfo) IsCloudSQL(ctx context.Context, nCloudType int64) (*PxgValBool, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"nCloudType": %d}`, nCloudType))
	request, err := http.NewRequest("POST", di.client.Server+"/api/v1.0/DatabaseInfo.IsCloudSQL", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValBool := new(PxgValBool)
	raw, err := di.client.Do(ctx, request, &pxgValBool)
	return pxgValBool, raw, err
}

// CheckBackupPath Check the server administration and SQL-server permissions to read and write files along path.
func (di *DatabaseInfo) CheckBackupPath(ctx context.Context, szwPath string) (*PxgValBool, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"szwPath": "%s"}`, szwPath))
	request, err := http.NewRequest("POST", di.client.Server+"/api/v1.0/DatabaseInfo.CheckBackupPath", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValBool := new(PxgValBool)
	raw, err := di.client.Do(ctx, request, &pxgValBool)
	return pxgValBool, raw, err
}

// CheckBackupPath2 Check the server administration and SQL-server permissions to read and write files along path.
func (di *DatabaseInfo) CheckBackupPath2(ctx context.Context, szwWinPath, szwLinuxPath string) (*PxgValBool, []byte, error) {
	postData := []byte(fmt.Sprintf(`{"szwWinPath": "%s",	"szwLinuxPath": "%s"}`, szwWinPath, szwLinuxPath))
	request, err := http.NewRequest("POST", di.client.Server+"/api/v1.0/DatabaseInfo.CheckBackupPath2", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValBool := new(PxgValBool)
	raw, err := di.client.Do(ctx, request, &pxgValBool)
	return pxgValBool, raw, err
}

// IsLinuxSQL Check is current SQL server in on Linux.
func (di *DatabaseInfo) IsLinuxSQL(ctx context.Context) (*PxgValBool, []byte, error) {
	request, err := http.NewRequest("POST", di.client.Server+"/api/v1.0/DatabaseInfo.IsLinuxSQL", nil)
	if err != nil {
		return nil, nil, err
	}

	pxgValBool := new(PxgValBool)
	raw, err := di.client.Do(ctx, request, &pxgValBool)
	return pxgValBool, raw, err
}
