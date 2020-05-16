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
)

//	ListTags Class Reference
//
//	Interface allows to acquire and manage tags to various KSC objects..
//
//	To manage tags for the list with name "ListName" it is needed to create this Interface with instance "ListName"
//	. Examples of list names: "HostsTags", "InventoryTags", "UmdmDeviceTags"
//
//	List of all members.
type ListTags service

//TODO Call ListTags.GetAllTags for the instance '' (listener '') does not exist (any more?)
func (lt *ListTags) GetAllTags(ctx context.Context) ([]byte, error) { return nil, nil }

//TODO Call ListTags.AddTag for the instance '' (listener '') does not exist (any more?)
func (lt *ListTags) AddTag(ctx context.Context) ([]byte, error) { return nil, nil }

//TODO Call ListTags.DeleteTags2 for the instance '' (listener '') does not exist (any more?)
func (lt *ListTags) DeleteTags2(ctx context.Context) ([]byte, error) { return nil, nil }

//TODO Call ListTags.GetTags for the instance '' (listener '') does not exist (any more?)
func (lt *ListTags) GetTags(ctx context.Context) ([]byte, error) { return nil, nil }

//TODO Call ListTags.RenameTag for the instance '' (listener '') does not exist (any more?)
func (lt *ListTags) RenameTag(ctx context.Context) ([]byte, error) { return nil, nil }

//TODO Call ListTags.SetTags for the instance '' (listener '') does not exist (any more?)
func (lt *ListTags) SetTags(ctx context.Context) ([]byte, error) { return nil, nil }
