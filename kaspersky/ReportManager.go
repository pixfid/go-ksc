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

//	ReportManager Class Reference
//
//	Reports managing. More...
//
//	Allows to enumerate, create, execute and delete reports.
//
//	List of all members.
type ReportManager service

//	Enumerate report types supported by administration server.
//
//	Enumerates all existing reports.
//
//	Returns:
//	- (array) collection of report type descriptions. Each entry is params object,
//	containing following report type attributes (See List of report type attributes):
//		- RPT_TYPE
//		- RPT_DN
func (rm *ReportManager) EnumReportTypes(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", rm.client.Server+"/api/v1.0/ReportManager.EnumReportTypes", nil)
	if err != nil {
		return nil, err
	}

	raw, err := rm.client.Do(ctx, request, nil)
	return raw, err
}

//	Enumerate existing reports.
//
//	Enumerates all existing reports.
//
//	Returns:
//	- (array) collection of reports. Each entry is params object,
//	containing following attributes (See List of report attributes):
//	- RPT_ID
//	- RPT_DN
//	- RPT_TYPE
//	- RPT_CREATED
//	- RPT_MODIFIED
//	- RPT_GROUP_ID
//	- RPT_EXTRA_DATA
func (rm *ReportManager) EnumReports(ctx context.Context) ([]byte, error) {
	request, err := http.NewRequest("POST", rm.client.Server+"/api/v1.0/ReportManager.EnumReports", nil)
	if err != nil {
		return nil, err
	}

	raw, err := rm.client.Do(ctx, request, nil)
	return raw, err
}

//	Enumerate available dashboards.
//
//	Returns:
//	- Collection of integer dashboards IDs.
func (rm *ReportManager) GetAvailableDashboards(ctx context.Context) (*PxgValArrayOfInt, []byte, error) {
	request, err := http.NewRequest("POST", rm.client.Server+"/api/v1.0/ReportManager.GetAvailableDashboards", nil)
	if err != nil {
		return nil, nil, err
	}

	reportsArray := new(PxgValArrayOfInt)
	raw, err := rm.client.Do(ctx, request, &reportsArray)
	return reportsArray, raw, err
}

//	Return XSLT transform for report type.
//
//	Returns XSLT transform as a string for specified report type.
//
//	Parameters:
//	- lReportType	(int) report type, see List of report types
//	- lXmlTargetType	(int) XML target type XML Target Type enum
//
//	Returns:
//	- (string) XSLT transform in the form of a string
func (rm *ReportManager) GetConstantOutputForReportType(ctx context.Context, lReportType, lXmlTargetType int64) (*PxgValStr,
	[]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lReportType": %d, "lXmlTargetType": %d}`, lReportType, lXmlTargetType))
	request, err := http.NewRequest("POST", rm.client.Server+"/api/v1.0/ReportManager.GetConstantOutputForReportType", bytes.NewBuffer(postData))
	if err != nil {
		return nil, nil, err
	}

	pxgValStr := new(PxgValStr)
	raw, err := rm.client.Do(ctx, request, &pxgValStr)
	return pxgValStr, raw, err
}

//	Obtain default report info for specified report type.
//
//	Returns default report info for the specified report type.
//
//	Parameters:
//	- lReportType	(int64) id of report type, see List of report types
//
//	Returns:
//	- (params) report info, see List of report attributes
func (rm *ReportManager) GetDefaultReportInfo(ctx context.Context, lReportType int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lReportType": %d}`, lReportType))
	request, err := http.NewRequest("POST", rm.client.Server+"/api/v1.0/ReportManager.GetDefaultReportInfo", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := rm.client.Do(ctx, request, nil)
	return raw, err
}

//	Get filter settings.
//
//	Parameters:
//	- lReportType	(int64) report type, see List of report types
//
//	Returns:
//	- (params) filter settings, see List of report filter attributes
func (rm *ReportManager) GetFilterSettings(ctx context.Context, lReportType int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lReportType": %d}`, lReportType))
	request, err := http.NewRequest("POST", rm.client.Server+"/api/v1.0/ReportManager.GetFilterSettings", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := rm.client.Do(ctx, request, nil)
	return raw, err
}

//	Obtain report common data.
//
//	Returns common data for specified report.
//
//	Parameters:
//	- lReportId	(int64) id of report
//
//	Returns:
//	- (params) report common data, see List of report common attributes
func (rm *ReportManager) GetReportCommonData(ctx context.Context, lReportId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lReportId": %d}`, lReportId))
	request, err := http.NewRequest("POST", rm.client.Server+"/api/v1.0/ReportManager.GetReportCommonData", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := rm.client.Do(ctx, request, nil)
	return raw, err
}

//	Enumerate ids of existing reports.
//
//	Returns array of existing report ids.
//
//	Returns:
//	- (array) array of event ids(int)
func (rm *ReportManager) GetReportIds(ctx context.Context) (*PxgValArrayOfInt, []byte, error) {
	request, err := http.NewRequest("POST", rm.client.Server+"/api/v1.0/ReportManager.GetReportIds", nil)
	if err != nil {
		return nil, nil, err
	}

	reportsIdsArray := new(PxgValArrayOfInt)
	raw, err := rm.client.Do(ctx, request, &reportsIdsArray)
	return reportsIdsArray, raw, err
}

//	Obtain report info.
//
//	Returns report info for specified report.
//
//	Parameters:
//	- lReportId	(int64) id of report
//
//	Returns:
//	- (params) report info, see List of report attributes
func (rm *ReportManager) GetReportInfo(ctx context.Context, lReportId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lReportId": %d}`, lReportId))
	request, err := http.NewRequest("POST", rm.client.Server+"/api/v1.0/ReportManager.GetReportInfo", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := rm.client.Do(ctx, request, nil)
	return raw, err
}

//	Obtain report type info.
//
//	Returns report info for specified report type.
//
//	Parameters:
//	- lReportTypeId	(int64) id of report type
//
//	Returns:
//	(params) report info, containing following report type attributes (See List of report type attributes)
//	- RPT_DESCR
//	- RPT_FIELDS
//	- RPT_SUMM_FIELDS
//	- RPT_TOTALS
//	- RPT_ACCEPT_TIME_INTERVAL
//	- RPT_ACCEPT_COMPS_LIST
//	- RPT_EXTRA_DATA
func (rm *ReportManager) GetReportTypeDetailedInfo(ctx context.Context, lReportType int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lReportType": %d}`, lReportType))
	request, err := http.NewRequest("POST", rm.client.Server+"/api/v1.0/ReportManager.GetReportTypeDetailedInfo", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := rm.client.Do(ctx, request, nil)
	return raw, err
}

//	Get result of ReportManager::RequestStatisticsData operation.
//
//	Gets result of asynchronous operation
//	ReportManager::RequestStatisticsData, such as statistics, general statuses and dashboards data.
//
//	Parameters:
//	- strRequestId	(string) identity of asynchronous operation
//
//	Return:
//	- pResultData	(params) result data, see List of statistics result dataset
func (rm *ReportManager) GetStatisticsData(ctx context.Context, strRequestId string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"strRequestId": "%s"}`, strRequestId))
	request, err := http.NewRequest("POST", rm.client.Server+"/api/v1.0/ReportManager.GetStatisticsData", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := rm.client.Do(ctx, request, nil)
	return raw, err
}

//	Remove existing report.
//
//	Removes specified report
//
//	Parameters:
//	- lReportId	(int64) id of report to remove
func (rm *ReportManager) RemoveReport(ctx context.Context, lReportId int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"lReportId": %d}`, lReportId))
	request, err := http.NewRequest("POST", rm.client.Server+"/api/v1.0/ReportManager.RemoveReport", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := rm.client.Do(ctx, request, nil)
	return raw, err
}

//	Request statistics, general statuses and dashboards.
//
//	Asynchronously requests statistics, general statuses and dashboards data.
//
//	Parameters:
//	- pRequestParams	(params) Params with the list of requested parameters,
//	each element is of type Params and optionally contains query parameters, see List of statistics query attributes)
//
//	Return:
//	- strRequestId	(string) identity of asynchronous operation to be used to get the result data by
//	ReportManager.GetStatisticsData or cancel the request by ReportManager.CancelStatisticsRequest.
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
	raw, err := rm.client.Do(ctx, request, &requestID)
	return requestID, raw, err
}

//	Execute report.
//
//	Asynchronously executes specified report, creates resulting data in XML and data to chart.
//	The progress and result of the report generation is reported by the event KLPPT_EventRptExecDone.
//
//	Parameters:
//	- lReportId	(int64) report id
//	- pOptions	(params) options (see description below)
//
//	Return:
//	- strRequestId	(string) identity of asynchronous operation,
//	to get status use AsyncActionStateChecker.CheckActionState, lStateCode "1" means OK and "0" means fail,
//	to get result use AsyncActionStateChecker.CheckActionState,
//	pStateData contains URL-like links to download report files via HTTP GET request (see description below),
//	to cancel operation call ReportManager.ExecuteReportAsyncCancel
func (rm *ReportManager) ExecuteReportAsync(ctx context.Context, params *interface{}) ([]byte, error) {
	postData, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", rm.client.Server+"/api/v1.0/ReportManager.ExecuteReportAsync", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := rm.client.Do(ctx, request, nil)
	return raw, err
}

//	Cancel ReportManager.RequestStatisticsData operation.
//
//	Cancels asynchronous operation ReportManager.RequestStatisticsData.
//
//	Parameters:
//	- strRequestId	(string) identity of asynchronous operation
func (rm *ReportManager) CancelStatisticsRequest(ctx context.Context, strRequestId string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"strRequestId": "%s"}`, strRequestId))
	request, err := http.NewRequest("POST", rm.client.Server+"/api/v1.0/ReportManager.CancelStatisticsRequest", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := rm.client.Do(ctx, request, nil)
	return raw, err
}

//	Cancel ReportManager::ExecuteReportAsync operation.
//
//	Cancels asynchronous operation ReportManager.ExecuteReportAsync.
//
//	Parameters:
//	- strRequestId	(string) identity of asynchronous operation
func (rm *ReportManager) ExecuteReportAsyncCancel(ctx context.Context, strRequestId string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"strRequestId": "%s"}`, strRequestId))
	request, err := http.NewRequest("POST", rm.client.Server+"/api/v1.0/ReportManager.ExecuteReportAsyncCancel", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := rm.client.Do(ctx, request, nil)
	return raw, err
}

//	Cancel waiting for report data from slave servers.
//
//	Cancels waiting for report data from slave servers when generating report.
//	Can be used when report data from current server is already ready
//	(an appropriate report generation progress status event KLPPT_EventRptExecDone received),
//	but the data from slave servers is not, to get the report without data from some slave servers.
//
//	Parameters:
//	- strRequestId	(string) identity of asynchronous operation
func (rm *ReportManager) ExecuteReportAsyncCancelWaitingForSlaves(ctx context.Context, strRequestId string) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{"strRequestId": "%s"}`, strRequestId))
	request, err := http.NewRequest("POST", rm.client.Server+"/api/v1.0/ReportManager.ExecuteReportAsyncCancelWaitingForSlaves", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := rm.client.Do(ctx, request, nil)
	return raw, err
}

//TODO ResetStatisticsData
//TODO UpdateReport
//TODO AddReport
//TODO CreateChartPNG
//TODO ExecuteReportAsyncGetData
