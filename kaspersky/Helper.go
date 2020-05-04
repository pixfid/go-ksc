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
	"fmt"
	"net/http"
	"time"
)

//	func withContext(ctx context.Context, req *http.Request) *http.Request
func withContext(ctx context.Context, req *http.Request) *http.Request {
	return req.WithContext(ctx)
}

//	PxgValInt struct
type PxgValInt struct {
	Int int64 `json:"PxgRetVal"`
}

//	PxgValStr struct
type PxgValStr struct {
	Str string `json:"PxgRetVal"`
}

//	PxgValBool struct
type PxgValBool struct {
	Bool bool `json:"PxgRetVal"`
}

//	PxgValArrayOfInt struct
type PxgValArrayOfInt struct {
	ArrayInt []int64 `json:"PxgRetVal"`
}

//	PxgValArrayOfString struct
type PxgValArrayOfString struct {
	ArrayString []string `json:"PxgRetVal"`
}

//	RequestID struct
type RequestID struct {
	PxgError     *PxgError `json:"PxgError,omitempty"`
	StrRequestID *string   `json:"strRequestId,omitempty"`
}

//	PxgValCIFIL struct
type PxgValCIFIL struct {
	PxgRetValCIFIL []PxgRetValCIFIL `json:"PxgRetVal"`
}
type PxgRetValCIFIL struct {
	CIFILValue CIFILValue `json:"value"`
}
type CIFILValue struct {
	InifileDetectMsi      string `json:"INIFILE_DETECT_MSI"`
	InifileDetectRegistry string `json:"INIFILE_DETECT_REGISTRY"`
	InifileFileName       string `json:"INIFILE_FILE_NAME"`
	InifileProductName    string `json:"INIFILE_PRODUCT_NAME"`
	InifileType           string `json:"INIFILE_TYPE"`
}

//	PxgRetError struct
type PxgRetError struct {
	Response *http.Response
	PxgError *PxgError `json:"PxgError,omitempty"`
}

type PxgError struct {
	Code    *int64  `json:"code,omitempty"`
	File    *string `json:"file,omitempty"`
	Line    *int64  `json:"line,omitempty"`
	Message *string `json:"message,omitempty"`
	Module  *string `json:"module,omitempty"`
}

func (p PxgError) Error() string {
	return fmt.Sprintf("Module: %s, Code: %d, Message: %s", *p.Module, *p.Code, *p.Message)
}

//	AsyncAccessor struct
type AsyncAccessor struct {
	StrAccessor       string             `json:"strAccessor,omitempty"`
	PFailedSlavesInfo *PFailedSlavesInfo `json:"pFailedSlavesInfo,omitempty"`
	PxgRetVal         int64              `json:"PxgRetVal,omitempty"`
}
type PFailedSlavesInfo struct {
	KlgrpFailedSlavesParams []interface{} `json:"KLGRP_FAILED_SLAVES_PARAMS"`
}

//	Accessor struct
type Accessor struct {
	StrAccessor string `json:"strAccessor"`
	PxgRetVal   int64  `json:"PxgRetVal"`
}

//	ADHostIterator struct
type ADHostIterator struct {
	WstrIterator string `json:"wstrIterator"`
	PxgRetVal    int64  `json:"PxgRetVal"`
}

//WActionGUID struct
type WActionGUID struct {
	WstrActionGUID string `json:"wstrActionGuid"`
}

const (
	RFC3339 = "2006-01-02T15:04:05Z07:00"
	RUS     = "2 Jan 2006 15:04"
)

func ParseTime(dt string) string {
	t, _ := time.Parse(RFC3339, dt)
	return t.Format(RUS)
}

//VecFieldsToOrder struct
type VecFieldsToOrder struct {
	Type  string `json:"type"`
	Value Value  `json:"value"`
}

type Value struct {
	Name string `json:"Name"`
	Asc  bool   `json:"Asc"`
}
