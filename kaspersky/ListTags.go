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
