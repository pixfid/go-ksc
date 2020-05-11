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
 * along with this program.If not, see <http://  www.gnu.org/licenses/>.
 * /
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
//	Interface for working with Limits subsystem. More...
//
//	This interface allow you to get a limit of specified parameter
//
//	List of all members.
type Limits service

//  Returns a limit of specified parameter.
//
//  Parameters:
//  - param	[in] (int64) See Limited parameters.
//
//  Returns:
//  - (int64) Parameter limit. For bool types 1 - true, 0 - false.
//
// Exceptions:
//  - KLSTD::STDE_NOACCESS	- Access to object is denied
//
// Limits params:
//  ParamID	Description
//  0	Max count of virtual servers
//  1	Max count of hosts
//  2	Max count of internal users
//  3	Max count of events in db
//  4	Possibility of moving nagent
//  5	Software inventory enabled
//  6	Hardware inventory enabled
//  7	Update agents enabled
//  8	Slave server enabled
//  9	All network scan types enabled
//  10	Connect only using SSL
//  11	Console must use port 13291
//  12	Nagent must use port 13000
//  13	Nagent must use two way authentication
//  14	Mobile must use two way authentication using port 13292
//  15	Manual change of flag "Keep connection" enabled
//  16	Manual creation of moving rules enabled
//  17	Manual creation of group on virtual servers enabled
//  18	Auto host tag rules enabled
//  19	Backup and restore enabled
//  20	System management enabled
//  21	NAC enabled
//  22	PXE enabled
//  23	ExtPatches installation enabled
//  24	WSUS enabled
//  25	Third party license management
//  26	Custom categories enabled
//  27	Application control enabled
//  28	KsnProxy enabled
//  29	Activation proxy enabled
//  30	Max count of installation packages on virtual server
//  31	Max count of moving rules on virtual server
//  32	Max count of tasks on virtual server
//  33	Max count of policies on virtual server
//  34	Max count of licenses on virtual server
//  35	Max count of instances of reports< on virtual server/TD>
//  36	Max count of computer queries on virtual server
//  37	Max count of event queries on virtual server
//  38	Max count of hosts on virtual server
//  39	Max count of internal security groups on virtual server
//  40	Check iOS MDM license restrict
//  41	Add keys version 1.0 to KSC key storage
//  42	Offline updates enabled
//  43	Broadcast domains detection enabled
//  44	If true then user must be unique in virtual server, else unique in phys server
//  45	Check sleep (hibernation) enabled
//  46	Restrict console connection enabled
//  47	Max count of console connections
//  48	Max count of console connections to all virtual servers
//  49	Max size of JSON body in KLOAPI request in bytes
//  50	Max size of allocated data per HTTP request parsing
//  51	Default lifetime in seconds, of tied objects (chunkaccessors, prssp, etc)
//  52	Maximum lifetime in seconds, of tied objects (chunkaccessors, prssp, etc)
//  53	Information about deleted objects is stored
//  54	Require client authentication in master-slave connectinos
//  55	Maximum size of uploading file to get metadata
//  56	Groups syncs (tasks, polocies, etc) from the master server affects virtual servers of current server (which is slave)
func (ls *Limits) GetLimits(ctx context.Context, param int64) ([]byte, error) {
	postData := []byte(fmt.Sprintf(`{ "param": %d }`, param))
	request, err := http.NewRequest("POST", ls.client.Server+"/api/v1.0/Limits.GetLimits", bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}

	raw, err := ls.client.Do(ctx, request, nil)
	return raw, err
}
