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

//	ExtAud Class Reference
//
//	Interface for working with ExtAudit subsystem. More...
//
//	This interface allow you to get a revision of an object and update description.
//
//	List of all members.
type ExtAud service

//
//	Get object revision.
//	╔════╦══════════════════════╗
//	║ ID ║     Description      ║
//	╠════╬══════════════════════╣
//	║  0 ║ None                 ║
//	║  1 ║ Policy               ║
//	║  2 ║ Task                 ║
//	║  3 ║ Package              ║
//	║  4 ║ Server               ║
//	║  5 ║ Virtual server       ║
//	║  6 ║ User                 ║
//	║  7 ║ Security Group       ║
//	║  8 ║ Administration Group ║
//	╚════╩══════════════════════╝
func (ea *ExtAud) GetRevision(ctx context.Context, nObjId, nObjType, nObjRevision int64, out interface{}) ([]byte,
	error) {
	postData := []byte(fmt.Sprintf(`{"nObjId": %d,"nObjType": %d,"nObjRevision": %d}`, nObjId, nObjType, nObjRevision))

	request, err := http.NewRequest("POST", ea.client.Server+"/api/v1.0/ExtAud.GetRevision", bytes.NewBuffer(postData))
	raw, err := ea.client.Do(ctx, request, &out)
	return raw, err
}

//
//	Update revision description.
//	╔════╦══════════════════════╗
//	║ ID ║     Description      ║
//	╠════╬══════════════════════╣
//	║  0 ║ None                 ║
//	║  1 ║ Policy               ║
//	║  2 ║ Task                 ║
//	║  3 ║ Package              ║
//	║  4 ║ Server               ║
//	║  5 ║ Virtual server       ║
//	║  6 ║ User                 ║
//	║  7 ║ Security Group       ║
//	║  8 ║ Administration Group ║
//	╚════╩══════════════════════╝
func (ea *ExtAud) UpdateRevisionDesc(ctx context.Context, nObjId, nObjType, nObjRevision int64, wstrNewDescription string) ([]byte,
	error) {
	postData := []byte(fmt.Sprintf(`{"nObjId": %d,"nObjType": %d,"nObjRevision": %d, "wstrNewDescription": "%s"}`,
		nObjId,
		nObjType,
		nObjRevision, wstrNewDescription))

	request, err := http.NewRequest("POST", ea.client.Server+"/api/v1.0/ExtAud.UpdateRevisionDesc", bytes.NewBuffer(postData))
	raw, err := ea.client.Do(ctx, request, nil)
	return raw, err
}

//	Final delete for deleted objects.
//
//	Parameters:
//	arrObjects	[interface{}] (array) Array of pairs ObjId-ObjType.
//	Max size of array is 100 elements.
func (ea *ExtAud) FinalDelete(ctx context.Context, params interface{}) ([]byte, error) {
	postData, _ := json.Marshal(params)
	request, err := http.NewRequest("POST", ea.client.Server+"/api/v1.0/ExtAud.FinalDelete", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ea.client.Do(ctx, request, nil)
	return raw, err
}
