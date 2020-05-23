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

//	Limits Class Reference
//
//	Interface for working with Limits subsystem..
//
//	This interface allow you to get a limit of specified parameter
//
//	List of all members.
type Limits service

//  Returns a limit of specified parameter.
//
//  Parameters:
//  - param (int64) See Limited parameters.
//
//  Returns:
//  - (int64) Parameter limit. For bool types 1 - true, 0 - false.
//
// Exceptions:
//  - KLSTD.STDE_NOACCESS	- Access to object is denied
//
// Limits params:
//	╔═══════╦════════════════════════════════════════════╦═══════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╗
//	║ Value ║               Mnemonic name                ║                                                                            Description                                                                            ║
//	╠═══════╬════════════════════════════════════════════╬═══════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╣
//	║     0 ║ LP_MaxCountOfVirtualServer                 ║ Max count of virtual servers                                                                                                                                      ║
//	║     1 ║ LP_MaxCountOfHosts                         ║ Max count of hosts                                                                                                                                                ║
//	║     2 ║ LP_MaxCountOfInternalUsers                 ║ Max count of internal users                                                                                                                                       ║
//	║     3 ║ LP_MaxCountOfEvents                        ║ Max count of events in db                                                                                                                                         ║
//	║     4 ║ LP_NagentMoving                            ║ Possibility of moving nagent                                                                                                                                      ║
//	║     5 ║ LP_SoftwareInventory                       ║ Software inventory enabled                                                                                                                                        ║
//	║     6 ║ LP_HardwareInventory                       ║ Hardware inventory enabled                                                                                                                                        ║
//	║     7 ║ LP_UpdateAgent                             ║ Update agents enabled                                                                                                                                             ║
//	║     8 ║ LP_SlaveServer                             ║ Slave server enabled                                                                                                                                              ║
//	║     9 ║ LP_NetworkScanByServer                     ║ Administration Server is able to scan the network                                                                                                                 ║
//	║    10 ║ LP_SslOnly                                 ║ Connect only using SSL                                                                                                                                            ║
//	║    11 ║ LP_ConsoleMustUsePort13291                 ║ Console must use port 13291                                                                                                                                       ║
//	║    12 ║ LP_NagentMustUsePort13000                  ║ Nagent must use port 13000                                                                                                                                        ║
//	║    13 ║ LP_NagentMustUseTwoWayAuth                 ║ Nagent must use two way authentication                                                                                                                            ║
//	║    14 ║ LP_MobileMustUseTwoWayAuthOnPort13292      ║ Mobile must use two way authentication using port 13292                                                                                                           ║
//	║    15 ║ LP_ManualSetFlagKeepConnection             ║ Manual change of flag "Keep connection" enabled                                                                                                                   ║
//	║    16 ║ LP_ManualCreationOfMovingRules             ║ Manual creation of moving rules enabled                                                                                                                           ║
//	║    17 ║ LP_ManualCreationOfGroupsOnVirtServer      ║ Manual creation of group on virtual servers enabled                                                                                                               ║
//	║    18 ║ LP_HostTagRules                            ║ Auto host tag rules enabled                                                                                                                                       ║
//	║    19 ║ LP_BackupAndRestore                        ║ Backup and restore enabled                                                                                                                                        ║
//	║    20 ║ LP_SystemManagement                        ║ System management enabled                                                                                                                                         ║
//	║    21 ║ LP_SM_NAC                                  ║ NAC enabled                                                                                                                                                       ║
//	║    22 ║ LP_SM_PXE                                  ║ PXE enabled                                                                                                                                                       ║
//	║    23 ║ LP_SM_ExtPatches                           ║ ExtPatches installation enabled                                                                                                                                   ║
//	║    24 ║ LP_SM_WSUS                                 ║ WSUS enabled                                                                                                                                                      ║
//	║    25 ║ LP_SM_ThirdPartyLicenseManagement          ║ Third party license management                                                                                                                                    ║
//	║    26 ║ LP_CustomCategories                        ║ Custom categories enabled                                                                                                                                         ║
//	║    27 ║ LP_AppControl                              ║ Application control enabled                                                                                                                                       ║
//	║    28 ║ LP_KsnProxy                                ║ KsnProxy enabled                                                                                                                                                  ║
//	║    29 ║ LP_ActivationProxy                         ║ Activation proxy enabled                                                                                                                                          ║
//	║    30 ║ LP_VS_MaxCountOfInstallationPackages       ║ Max count of installation packages on virtual server                                                                                                              ║
//	║    31 ║ LP_VS_MaxCountOfMovingRules                ║ Max count of moving rules on virtual server                                                                                                                       ║
//	║    32 ║ LP_VS_MaxCountOfTasks                      ║ Max count of tasks on virtual server                                                                                                                              ║
//	║    33 ║ LP_VS_MaxCountOfPolicies                   ║ Max count of policies on virtual server                                                                                                                           ║
//	║    34 ║ LP_VS_MaxCountOfLicenses                   ║ Max count of licenses on virtual server                                                                                                                           ║
//	║    35 ║ LP_VS_MaxCountOfReportInstances            ║ Max count of instances of reports on virtual server                                                                                                               ║
//	║    36 ║ LP_VS_MaxCountOfComputerQueries            ║ Max count of computer queries on virtual server                                                                                                                   ║
//	║    37 ║ LP_VS_MaxCountOfEventQueries               ║ Max count of event queries on virtual server                                                                                                                      ║
//	║    38 ║ LP_VS_MaxCountOfHosts                      ║ Max count of hosts on virtual server                                                                                                                              ║
//	║    39 ║ LP_VS_MaxCountOfInternalSecurityGroups     ║ Max count of internal security groups on virtual server                                                                                                           ║
//	║    40 ║ LP_LicLoadRestrictIosMdm                   ║ Check iOS MDM license restrict                                                                                                                                    ║
//	║    41 ║ LP_VS_LicLoadKeyFile                       ║ Add keys version 1.0 to KSC key storage                                                                                                                           ║
//	║    42 ║ LP_OfflineUpdates                          ║ Offline updates enabled                                                                                                                                           ║
//	║    43 ║ LP_BroadcastDomains                        ║ Broadcast domains detection enabled                                                                                                                               ║
//	║    44 ║ LP_InterUserUniqVsScope                    ║ If true then user must be unique in virtual server, else unique in phys server                                                                                    ║
//	║    45 ║ LP_SysPowerManagement                      ║ Check sleep(hibernation) enabled                                                                                                                                  ║
//	║    46 ║ LP_RestrictRemoteConsole                   ║ Restrict console connection enabled                                                                                                                               ║
//	║    47 ║ LP_MaxCountOfConsoles                      ║ Max count of console connections                                                                                                                                  ║
//	║    48 ║ LP_VS_MaxTotalCountOfConsoles              ║ Max count of console connections to all virtual servers                                                                                                           ║
//	║    49 ║ LP_KLOAPI_MaxSizeOfJsonRequestInBytes      ║ Max size of JSON body in KLOAPI request in bytes                                                                                                                  ║
//	║    50 ║ LP_KLOAPI_MaxSizeOfHttpRequestInBytes      ║ Max size of allocated data per HTTP request parsing                                                                                                               ║
//	║    51 ║ LP_TiedObjLifeTimeSecDefault               ║ Default lifetime in seconds, of tied objects (chunk accessors, settings storages, etc)                                                                            ║
//	║    52 ║ LP_TiedObjLifeTimeSecMax                   ║ Maximum lifetime in seconds, of tied objects (chunk accessors, settings storages, etc)                                                                            ║
//	║    53 ║ LP_InTrashObjectsSupported                 ║ Information about deleted objects is stored                                                                                                                       ║
//	║    54 ║ LP_HrchMustUseTwoWayAuth                   ║ Require client authentication in master - slave connections                                                                                                       ║
//	║    55 ║ LP_FcMaxUploadFileSize                     ║ Maximum size of uploading file to get metadata                                                                                                                    ║
//	║    56 ║ LP_MasterAffectsVsOfSlave                  ║ Groups syncs(tasks, policies, etc) from the master server affects virtual servers of current server(which is slave)                                               ║
//	║    57 ║ LP_MaxLoginQueueSize                       ║ Max length of login queue                                                                                                                                         ║
//	║    58 ║ LP_CategoryFromDir                         ║ Allow to create a category based on hashes of files from directory                                                                                                ║
//	║    59 ║ LP_AssignUaAutomatically                   ║ Automatic assignment of UAs is allowed                                                                                                                            ║
//	║    60 ║ LP_CustomInstallationPackages              ║ If custom installation packages are allowed                                                                                                                       ║
//	║    61 ║ LP_MassEvents                              ║ Mass eventing is allowed                                                                                                                                          ║
//	║    62 ║ LP_LicLoadKeyFilePhysServer                ║ Allow to put license keys activation 1.0 into license storage of the 'physical' server                                                                            ║
//	║    63 ║ LP_NlstMaxListSizeMultiplier               ║ Multiplier to get default max count of network list items per 1000 hosts to keep in the database                                                                  ║
//	║    64 ║ LP_EVP_NotifyByMail                        ║ Notification on events by email available                                                                                                                         ║
//	║    65 ║ LP_EVP_NotifyBySms                         ║ Notification on events by SMS available                                                                                                                           ║
//	║    66 ║ LP_EVP_NotifyByScript                      ║ Notification on events by running executable files or scripts available                                                                                           ║
//	║    67 ║ LP_EVP_NotifyBySnmp                        ║ Notification on events by SNMP available                                                                                                                          ║
//	║    68 ║ LP_EVP_LimitNotificationsByEventsTypes     ║ Support general events notification limits by events types ("KLEVP_TEST_PERIOD_TO_SEND_EVENTS" and "KLEVP_MAX_EVENTS_TO_SEND_PER_PERIOD")                         ║
//	║    69 ║ LP_EVP_AccumulateNotificationsByRecipients ║ Support notification accumulation by notification recipients ("KLEVP_NF_SINGLE_INTERVAL_FOR_RECIPIENT_SEC" and "KLEVP_NF_ACCUMULATIVE_INTERVAL_FOR_RECIPIENT_SEC")║
//	║    70 ║ LP_AuthSessionLifetimeSec                  ║ Authentication session maximum time to live, in seconds. unlim for infinity                                                                                       ║
//	║    71 ║ LP_AuthSessionInactiveLifetimeSec          ║ Authentication session maximum inactive timeout, in seconds. unlim for infinity                                                                                   ║
//	║    72 ║ LP_AdministrationServerPolicyAllowed       ║ if creation of policy for 1093 product is allowed                                                                                                                 ║
//	║    73 ║ LP_MaxCountOfGuiStatistics                 ║ Maximum count of gui-statistics to send at once                                                                                                                   ║
//	║    74 ║ LP_RestrictRemoteOsAuth                    ║ Restrict remote connections for OS users enabled                                                                                                                  ║
//	║    75 ║ LP_SendKsnStatistics                       ║ Send KSN statistics                                                                                                                                               ║
//	╚═══════╩════════════════════════════════════════════╩═══════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╝
func (ls *Limits) GetLimits(ctx context.Context, param int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{ "param": %d }`, param))
	request, err := http.NewRequest("POST", ls.client.Server+"/api/v1.0/Limits.GetLimits", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ls.client.Do(ctx, request, nil)
	return raw, err
}
