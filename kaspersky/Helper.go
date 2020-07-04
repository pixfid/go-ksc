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
	"encoding/json"
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

//	pPngData struct
type PPngData struct {
	PngData string `json:"pPngData"`
}

//	PxgValArrayOfInt struct
type PxgValArrayOfInt struct {
	Array []int64 `json:"PxgRetVal"`
}

//	PxgValArrayOfString struct
type PxgValArrayOfString struct {
	Array []string `json:"PxgRetVal"`
}

//	RequestID struct
type RequestID struct {
	StrRequestID string `json:"strRequestId,omitempty"`
}

//	PxgValCIFIL struct
type PxgValCIFIL struct {
	PxgRetValCIFIL []PxgRetValCIFIL `json:"PxgRetVal"`
}
type PxgRetValCIFIL struct {
	CIFILValue *CIFILValue `json:"value"`
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
	Error *Error `json:"PxgError,omitempty"`
}

type Error struct {
	Code    *int64  `json:"code,omitempty"`
	File    *string `json:"file,omitempty"`
	Line    *int64  `json:"line,omitempty"`
	Message *string `json:"message,omitempty"`
	Module  *string `json:"module,omitempty"`
}

func (e Error) Error() string {
	return fmt.Sprintf(`Code: %d, File: %s, Line: %d, Module: %s, Message: %s`, *e.Code, *e.File, *e.Line, *e.Module, *e.Message)
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

type PFindParams struct {
	StrFilter       string          `json:"strFilter,omitempty"`
	PFieldsToReturn []string        `json:"pFieldsToReturn"`
	PFieldsToOrder  []FieldsToOrder `json:"pFieldsToOrder,omitempty"`
	PParams         PParams         `json:"pParams,omitempty"`
	LMaxLifeTime    int64           `json:"lMaxLifeTime,omitempty"`
}

//FieldsToOrder struct
type FieldsToOrder struct {
	Type       string     `json:"type"`
	OrderValue OrderValue `json:"value"`
}

type OrderValue struct {
	Name string `json:"Name"`
	Asc  bool   `json:"Asc"`
}

type WstrIteratorID struct {
	WstrIteratorID string `json:"wstrIteratorId"`
}

type StrIteratorId struct {
	StrIteratorID string `json:"strIteratorId"`
}

type StrHostIteratorId struct {
	StrHostIteratorId string `json:"strHostIteratorId"`
}

type PFilter struct {
	KlevpRfc2254Filter         string `json:"KLEVP_RFC2254_FILTER"`
	KlevpEventRiseTimeLastDays int64  `json:"KLEVP_EVENT_RISE_TIME_LAST_DAYS"`
}

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Int64 is a helper routine that allocates a new int64 value
// to store v and returns a pointer to it.
func Int64(v int64) *int64 { return &v }

// Int64 is a helper routine that allocates a new int64 value
// to store v and returns a pointer to it.
func Float64(v float64) *float64 { return &v }

// Uint64 is a helper routine that allocates a new Uint64 value
// to store v and returns a pointer to it
func Uint64(v uint64) *uint64 { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }

type Null struct{}

//ToJson is a helper routine that marshal (interface{}) object to string
func ToJson(a interface{}) string {
	jsn, err := json.Marshal(a)
	if err == nil {
		return string(jsn)
	}
	return ""
}

func jsonEscape(i string) string {
	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	s := string(b)
	return s[1 : len(s)-1]
}

type DateTime struct {
	Type  *string `json:"type"`
	Value *string `json:"value"`
}

type Size struct {
	Type  *string `json:"type,omitempty"`
	Value *int64  `json:"value,omitempty"`
}

type CGMobileAuthCERT struct {
	Type  *string `json:"type,omitempty"`
	Value *string `json:"value,omitempty"`
}

type Long struct {
	Type  *string `json:"type,omitempty"`
	Value *int64  `json:"value,omitempty"`
}

type AsyncID struct {
	WstrAsyncID string `json:"wstrAsyncId,omitempty"`
}
