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
	"context"
	"net/http"
)

type ListTags service

//TODO Call ListTags.GetAllTags for the instance '' (listener '') does not exist (any more?)
func (lt *ListTags) GetAllTags(ctx context.Context) (*PxgRetError, []byte, error) {
	request, err := http.NewRequest("POST", lt.client.Server+"/api/v1.0/ListTags.GetAllTags", nil)
	pxgValTODO := new(PxgRetError)
	raw, err := lt.client.Do(ctx, request, &pxgValTODO)
	//TODO Found Correct Response format
	return pxgValTODO, raw, err
}
