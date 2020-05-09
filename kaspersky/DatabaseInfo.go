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

type DatabaseInfo service

//Get database's files size.
//
//Returns size of files of database
//
//Returns:
//	- (data.PxgValInt) size of files of database.
func (di *DatabaseInfo) GetDBSize(ctx context.Context) (*PxgValInt, []byte, error) {
	request, err := http.NewRequest("POST", di.client.Server+"/api/v1.0/DatabaseInfo.GetDBSize", nil)
	pxgValInt := new(PxgValInt)
	raw, err := di.client.Do(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

//Get database's data size.
//
//Returns size of data of database
//
//Returns:
//	- (data.PxgValInt) size of data of database.
func (di *DatabaseInfo) GetDBDataSize(ctx context.Context) (*PxgValInt, []byte, error) {
	request, err := http.NewRequest("POST", di.client.Server+"/api/v1.0/DatabaseInfo.GetDBDataSize", nil)
	pxgValInt := new(PxgValInt)
	raw, err := di.client.Do(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

//Get database's events count.
//
//Returns count of events of database
//
//Returns:
//	- (data.PxgValInt) count of events of database
func (di *DatabaseInfo) GetDBEventsCount(ctx context.Context) (*PxgValInt, []byte, error) {
	request, err := http.NewRequest("POST", di.client.Server+"/api/v1.0/DatabaseInfo.GetDBEventsCount", nil)
	pxgValInt := new(PxgValInt)
	raw, err := di.client.Do(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

//Check is current SQL server in cloud (Amazon RDS or Azure SQL)
//
//Parameters:
//	- nCloudType	(int64) Cloud type (KLCLOUD::CloudType)
//Returns:
//	- (data.PxgValBool) true if there is SQL database of this cloud type
func (di *DatabaseInfo) IsCloudSQL(ctx context.Context, nCloudType int64) (*PxgValBool, []byte, error) {
	postData := []byte(fmt.Sprintf(`
	{
		"nCloudType": %d
	}`, nCloudType))

	request, err := http.NewRequest("POST", di.client.Server+"/api/v1.0/DatabaseInfo.IsCloudSQL", bytes.NewBuffer(postData))
	pxgValBool := new(PxgValBool)
	raw, err := di.client.Do(ctx, request, &pxgValBool)
	return pxgValBool, raw, err
}

//TODO Check the server administration and SQL-server permissions to read and write files along path.
//
//Parameters:
//	- szwPath	(string) full-path to checkable directory
//Exceptions:
//	- Throw	exception if there are no any permissions
func (di *DatabaseInfo) CheckBackupPath(ctx context.Context, szwPath string) (*PxgValBool, []byte, error) {
	postData := []byte(fmt.Sprintf(`
	{
		"szwPath": "%s"
	}`, szwPath))

	request, err := http.NewRequest("POST", di.client.Server+"/api/v1.0/DatabaseInfo.CheckBackupPath", bytes.NewBuffer(postData))
	pxgValBool := new(PxgValBool)
	raw, err := di.client.Do(ctx, request, &pxgValBool)
	return pxgValBool, raw, err
}
