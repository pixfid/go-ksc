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
	"encoding/json"
	"fmt"
	"net/http"
)

// ReportManager Reports managing.
//
// Allows to enumerate, create, execute and delete reports.
type ReportManager service

// EnumReportTypes Enumerate report types supported by administration server.
//
// Enumerates all existing types.
func (rm *ReportManager) EnumReportTypes(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", rm.client.Server+"/api/v1.0/ReportManager.EnumReportTypes", nil)
	if err != nil {
		return nil, err
	}

	raw, err := rm.client.Request(ctx, request, nil)
	return raw, err
}

// EnumReports Enumerate existing reports.
//
// Enumerates all existing reports.
func (rm *ReportManager) EnumReports(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", rm.client.Server+"/api/v1.0/ReportManager.EnumReports", nil)
	if err != nil {
		return nil, err
	}

	raw, err := rm.client.Request(ctx, request, nil)
	return raw, err
}

// GetAvailableDashboards Enumerate available dashboards.
func (rm *ReportManager) GetAvailableDashboards(ctx context.Context) (*PxgValArrayOfInt, []byte, error) {
	request, err := http.NewRequest("POST", rm.client.Server+"/api/v1.0/ReportManager.GetAvailableDashboards", nil)
	if err != nil {
		return nil, nil, err
	}

	reportsArray := new(PxgValArrayOfInt)
	raw, err := rm.client.Request(ctx, request, &reportsArray)
	return reportsArray, raw, err
}

// GetConstantOutputForReportType Return XSLT transform for report type.
//
// Returns XSLT transform as a string for specified report type.
func (rm *ReportManager) GetConstantOutputForReportType(ctx context.Context, lReportType, lXmlTargetType int64) (*PxgValStr,
	[]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lReportType": %d, "lXmlTargetType": %d}`, lReportType, lXmlTargetType))
	request, err := http.NewRequest("POST", rm.client.Server+"/api/v1.0/ReportManager.GetConstantOutputForReportType", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := rm.client.Request(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

// GetDefaultReportInfo Obtain default report info for specified report type.
//
// Returns default report info for the specified report type.
func (rm *ReportManager) GetDefaultReportInfo(ctx context.Context, lReportType int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lReportType": %d}`, lReportType))
	request, err := http.NewRequest("POST", rm.client.Server+"/api/v1.0/ReportManager.GetDefaultReportInfo", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := rm.client.Request(ctx, request, nil)
	return raw, err
}

// GetFilterSettings Get filter settings.
func (rm *ReportManager) GetFilterSettings(ctx context.Context, lReportType int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lReportType": %d}`, lReportType))
	request, err := http.NewRequest("POST", rm.client.Server+"/api/v1.0/ReportManager.GetFilterSettings", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := rm.client.Request(ctx, request, nil)
	return raw, err
}

// GetReportCommonData Obtain report common data.
//
// Returns common data for specified report.
func (rm *ReportManager) GetReportCommonData(ctx context.Context, lReportId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lReportId": %d}`, lReportId))
	request, err := http.NewRequest("POST", rm.client.Server+"/api/v1.0/ReportManager.GetReportCommonData", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := rm.client.Request(ctx, request, nil)
	return raw, err
}

// GetReportIds Enumerate ids of existing reports.
//
// Returns array of existing report ids.
func (rm *ReportManager) GetReportIds(ctx context.Context) (*PxgValArrayOfInt, []byte, error) {
	request, err := http.NewRequest("POST", rm.client.Server+"/api/v1.0/ReportManager.GetReportIds", nil)
	if err != nil {
		return nil, nil, err
	}

	reportsIdsArray := new(PxgValArrayOfInt)
	raw, err := rm.client.Request(ctx, request, &reportsIdsArray)
	return reportsIdsArray, raw, err
}

// GetReportInfo Obtain report info.
//
// Returns report info for specified report.
func (rm *ReportManager) GetReportInfo(ctx context.Context, lReportId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lReportId": %d}`, lReportId))
	request, err := http.NewRequest("POST", rm.client.Server+"/api/v1.0/ReportManager.GetReportInfo", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := rm.client.Request(ctx, request, nil)
	return raw, err
}

// GetReportTypeDetailedInfo Obtain report type info.
//
// Returns report info for specified report type.
func (rm *ReportManager) GetReportTypeDetailedInfo(ctx context.Context, lReportType int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lReportType": %d}`, lReportType))
	request, err := http.NewRequest("POST", rm.client.Server+"/api/v1.0/ReportManager.GetReportTypeDetailedInfo", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := rm.client.Request(ctx, request, nil)
	return raw, err
}

// GetStatisticsData Gets result of asynchronous operation ReportManager.RequestStatisticsData,
// such as statistics, general statuses and dashboards data.
func (rm *ReportManager) GetStatisticsData(ctx context.Context, strRequestId string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"strRequestId": "%s"}`, strRequestId))
	request, err := http.NewRequest("POST", rm.client.Server+"/api/v1.0/ReportManager.GetStatisticsData", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := rm.client.Request(ctx, request, nil)
	return raw, err
}

// RemoveReport Removes specified report
func (rm *ReportManager) RemoveReport(ctx context.Context, lReportId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lReportId": %d}`, lReportId))
	request, err := http.NewRequest("POST", rm.client.Server+"/api/v1.0/ReportManager.RemoveReport", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := rm.client.Request(ctx, request, nil)
	return raw, err
}

// RequestStatisticsData Request statistics, general statuses and dashboards.
//
// Asynchronously requests statistics, general statuses and dashboards data.
func (rm *ReportManager) RequestStatisticsData(ctx context.Context, params interface{}) (*RequestID, []byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", rm.client.Server+"/api/v1.0/ReportManager.RequestStatisticsData", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	requestID := new(RequestID)
	raw, err := rm.client.Request(ctx, request, &requestID)
	return requestID, raw, err
}

// ExecuteReportParams struct
type ExecuteReportParams struct {
	//report id
	LReportID int64 `json:"lReportId,omitempty"`

	//options (see description below)
	POptions *RPOptions `json:"pOptions,omitempty"`
}

type RPOptions struct {
	//Locale identifier
	RptLOCLocale int64 `json:"RPT_LOC_LOCALE,omitempty"`

	//If the flag is set, values of datetime fields will be in UTC,
	//YYYY-MM-DDTHH:mm:ss format. Otherwise, time zone will be taken into account,
	//long date format.
	KlrptUseUTC bool `json:"KLRPT_USE_UTC,omitempty"`

	//Description of report output format
	KlrptOutputFormat *KlrptOutputFormat `json:"KLRPT_OUTPUT_FORMAT,omitempty"`
}

type KlrptOutputFormat struct {
	Type  string                  `json:"type,omitempty"`
	Value *KlrptOutputFormatValue `json:"value,omitempty"`
}

type KlrptOutputFormatValue struct {
	//Maximum number of records in details table
	KlrptMaxRecordsDetails int64 `json:"KLRPT_MAX_RECORDS_DETAILS,omitempty"`

	//Report target format, see Types of report target format
	//	╔═══════╦══════════╦═════════════╗
	//	║ Value ║  Alias   ║ Description ║
	//	╠═══════╬══════════╬═════════════╣
	//	║     0 ║ RTT_XML  ║ XML         ║
	//	║     1 ║ RTT_CSV  ║ CSV         ║
	//	║     2 ║ RTT_JSON ║ JSON        ║
	//	╚═══════╩══════════╩═════════════╝
	KlrptTargetType int64 `json:"KLRPT_TARGET_TYPE,omitempty"`

	//Report target XML format, see Types of report XML target format
	//	╔═══════╦═════════════╦════════════════════╗
	//	║ Value ║    Alias    ║    Description     ║
	//	╠═══════╬═════════════╬════════════════════╣
	//	║    -1 ║ RTT_UNKNOWN ║ Unknown or not set ║
	//	║     0 ║ RTT_HTML    ║ HTML               ║
	//	║     1 ║ RTT_XLS     ║ XLS                ║
	//	║     2 ║ RTT_PDF     ║ PDF                ║
	//	╚═══════╩═════════════╩════════════════════╝
	KlrptXMLTargetType int64 `json:"KLRPT_XML_TARGET_TYPE,omitempty"`

	//PDF report document orientation

	KlrptPDFLandscape bool `json:"KLRPT_PDF_LANDSCAPE,omitempty"`

	//Page size for PDF report document, see Sizes of report PDF document
	//	╔═══════╦═════════════╦═════════════════════╗
	//	║ Value ║    Alias    ║     Description     ║
	//	╠═══════╬═════════════╬═════════════════════╣
	//	║     0 ║ Custom      ║ User-defined format ║
	//	║     1 ║ Letter      ║ Letter format       ║
	//	║     2 ║ Note        ║ Note format         ║
	//	║     3 ║ Legal       ║ Legal format        ║
	//	║     4 ║ A0          ║ A0 format           ║
	//	║     5 ║ A1          ║ A1 format           ║
	//	║     6 ║ A2          ║ A2 format           ║
	//	║     7 ║ A3          ║ A3 format           ║
	//	║     8 ║ A4          ║ A4 format           ║
	//	║     9 ║ A5          ║ A5 format           ║
	//	║    10 ║ A6          ║ A6 format           ║
	//	║    11 ║ A7          ║ A7 format           ║
	//	║    12 ║ A8          ║ A8 format           ║
	//	║    13 ║ A9          ║ A9 format           ║
	//	║    14 ║ A10         ║ A10 format          ║
	//	║    15 ║ B0          ║ B0 format           ║
	//	║    16 ║ B1          ║ B1 format           ║
	//	║    17 ║ B2          ║ B2 format           ║
	//	║    18 ║ B3          ║ B3 format           ║
	//	║    19 ║ B4          ║ B4 format           ║
	//	║    20 ║ B5          ║ B5 format           ║
	//	║    21 ║ ArchE       ║ ArchE format        ║
	//	║    22 ║ ArchD       ║ ArchD format        ║
	//	║    23 ║ ArchC       ║ ArchC format        ║
	//	║    24 ║ ArchB       ║ ArchB format        ║
	//	║    25 ║ ArchA       ║ ArchA format        ║
	//	║    26 ║ Flsa        ║ Flsa format         ║
	//	║    27 ║ HalfLetter  ║ HalfLetter format   ║
	//	║    28 ║ Letter11x17 ║ 11x17 format        ║
	//	║    29 ║ Ledger      ║ Ledger format       ║
	//	╚═══════╩═════════════╩═════════════════════╝
	KlrptPDFPageSize int64 `json:"KLRPT_PDF_PAGE_SIZE,omitempty"`
}

// ExecuteReportAsync Execute report
//
// Asynchronously executes specified report, creates resulting data in XML and data to chart.
// The progress and result of the report generation is reported by the event KLPPT_EventRptExecDone.
func (rm *ReportManager) ExecuteReportAsync(ctx context.Context, params ExecuteReportParams) (*RequestID, []byte,
	error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", rm.client.Server+"/api/v1.0/ReportManager.ExecuteReportAsync", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	requestID := new(RequestID)
	raw, err := rm.client.Request(ctx, request, &requestID)
	return requestID, raw, err
}

// CancelStatisticsRequest Cancels asynchronous operation ReportManager.RequestStatisticsData.
func (rm *ReportManager) CancelStatisticsRequest(ctx context.Context, strRequestId string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"strRequestId": "%s"}`, strRequestId))
	request, err := http.NewRequest("POST", rm.client.Server+"/api/v1.0/ReportManager.CancelStatisticsRequest", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := rm.client.Request(ctx, request, nil)
	return raw, err
}

// ExecuteReportAsyncCancel Cancels asynchronous operation ReportManager.ExecuteReportAsync.
func (rm *ReportManager) ExecuteReportAsyncCancel(ctx context.Context, strRequestId string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"strRequestId": "%s"}`, strRequestId))
	request, err := http.NewRequest("POST", rm.client.Server+"/api/v1.0/ReportManager.ExecuteReportAsyncCancel", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := rm.client.Request(ctx, request, nil)
	return raw, err
}

// ReportData struct
type ReportData struct {
	PXMLData      string      `json:"pXmlData,omitempty"`
	NDataSizeREST int64       `json:"nDataSizeRest,omitempty"`
	PChartData    *PChartData `json:"pChartData,omitempty"`
}

type PChartData struct {
	KlrptChartData         []KlrptChartDatum `json:"KLRPT_CHART_DATA"`
	KlrptChartDataDesc     string            `json:"KLRPT_CHART_DATA_DESC,omitempty"`
	KlrptChartLgndDesc     string            `json:"KLRPT_CHART_LGND_DESC,omitempty"`
	KlrptChartSeries       []string          `json:"KLRPT_CHART_SERIES"`
	KlrptChartSeriesColors []int64           `json:"KLRPT_CHART_SERIES_COLORS"`
}

type KlrptChartDatum struct {
	Type  string `json:"type,omitempty"`
	Value *Value `json:"value,omitempty"`
}

type Value struct {
	Data []int64 `json:"data"`
}

// ExecuteReportAsyncGetData Get result of ReportManager.ExecuteReportAsync operation.
//
// Gets result of asynchronous operation ReportManager::ExecuteReportAsync.
// If result is not ready pXmlData will be empty.
func (rm *ReportManager) ExecuteReportAsyncGetData(ctx context.Context, strRequestId string,
	nChunkSize int64) (*ReportData, []byte,
	error) {
	postData := []byte(fmt.Sprintf(`{"strRequestId" : "%s", "nChunkSize": %d}`, strRequestId, nChunkSize))
	request, err := http.NewRequest("POST", rm.client.Server+"/api/v1.0/ReportManager.ExecuteReportAsyncGetData", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	reportData := new(ReportData)
	raw, err := rm.client.Request(ctx, request, &reportData)
	return reportData, raw, err
}

// ExecuteReportAsyncCancelWaitingForSlaves Cancel waiting for report data from slave servers.
//
// Cancels waiting for report data from slave servers when generating report.
// Can be used when report data from current server is already ready (an appropriate report generation progress status
// event KLPPT_EventRptExecDone received), but the data from slave servers is not,
// to get the report without data from some slave servers.
func (rm *ReportManager) ExecuteReportAsyncCancelWaitingForSlaves(ctx context.Context, strRequestId string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"strRequestId": "%s"}`, strRequestId))
	request, err := http.NewRequest("POST", rm.client.Server+"/api/v1.0/ReportManager.ExecuteReportAsyncCancelWaitingForSlaves", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := rm.client.Request(ctx, request, nil)
	return raw, err
}

// ChartDataParams struct
type ChartDataParams struct {
	PChartData *PChartData `json:"pChartData,omitempty"`
	CDPOptions *CDPOptions `json:"pOptions,omitempty"`
}

type CDPOptions struct {
	RptChartWidth  int64 `json:"RPT_CHART_WIDTH,omitempty"`
	RptChartHeight int64 `json:"RPT_CHART_HEIGHT,omitempty"`
}

// CreateChartPNG Creates image in PNG format with chart data
func (rm *ReportManager) CreateChartPNG(ctx context.Context, params ChartDataParams) (*PPngData, []byte,
	error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", rm.client.Server+"/api/v1.0/ReportManager.CreateChartPNG", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pPngData := new(PPngData)
	raw, err := rm.client.Request(ctx, request, &pPngData)
	return pPngData, raw, err
}

// ResetStatisticsData Force reset of statistics data.
//
// Force resets statistics data, for example, resets the status of "Virus attack" or "Failed to perform the administration server task"
// after acquaintance with the detailed information.
func (rm *ReportManager) ResetStatisticsData(ctx context.Context, params interface{}) (*RequestID, []byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", rm.client.Server+"/api/v1.0/ReportManager.ResetStatisticsData", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	requestID := new(RequestID)
	raw, err := rm.client.Request(ctx, request, &requestID)
	return requestID, raw, err
}

// AddReport Create new report.
func (rm *ReportManager) AddReport(ctx context.Context, params interface{}) (*PxgValInt, []byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("POST", rm.client.Server+"/api/v1.0/ReportManager.AddReport", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValInt := new(PxgValInt)
	raw, err := rm.client.Request(ctx, request, &pxgValInt)
	return pxgValInt, raw, err
}

// UpdateReport Updates info for existing report
func (rm *ReportManager) UpdateReport(ctx context.Context, params interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", rm.client.Server+"/api/v1.0/ReportManager.UpdateReport", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := rm.client.Request(ctx, request, nil)
	return raw, err
}
