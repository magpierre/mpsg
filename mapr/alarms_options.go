/*
   Copyright 2019 MapR Technologies

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package mapr

import "github.com/c-bata/go-prompt"

var alarm_list = []prompt.Suggest{
	{Text: "-cluster", Description: "[ -cluster cluster_name ]"},
	{Text: "-summary", Description: "[ -summary summary. default: false ]"},
	{Text: "-alarm", Description: "[ -alarm alarm name ]"},
	{Text: "-type", Description: "[ -type type (CLUSTER OR NODE OR VOLUME OR AE) ]"},
	{Text: "-entity", Description: "[ -entity entity (hostname OR volume name OR Ae name) ]"},
	{Text: "-start", Description: "[ -start start. default: 0 ]"},
	{Text: "-limit", Description: "[ -limit limit. default: 2147483647 ]"},
	{Text: "-output", Description: "[ -output output. default: verbose ]"},
	{Text: "-sortby", Description: "[ -sortby <alarmname|alarmdescription|alarmtype|alarmstate|alarmraised|alarmcleared|alarmentity|alarmmutetime|alarmmuteupto|alarmmuteduration|alarmgroups> ]"},
	{Text: "-history", Description: "[ -history list cleared up alarms only Parameter takes no value  ]"},
	{Text: "-from", Description: "[ -from alarms raised after time(in millis) ]"},
	{Text: "-till", Description: "[ -till alarms raised before time(in millis) ]"},
	{Text: "-muted", Description: "[ -muted list alarms configured to be mute Parameter takes no value  ]"},
	{Text: "-all", Description: "[ -all list raised and muted alarms Parameter takes no value  ]"},
	{Text: "-sortorder", Description: "[ -sortorder <asc|desc> ]"},
}

// {Text:"", Description:"

var alarm_raise = []prompt.Suggest{
	{Text: "-cluster", Description: "[ -cluster cluster_name ]"},
	{Text: "-alarm", Description: " -alarm alarm"},
	{Text: "-entity", Description: "[ -entity entity (hostname OR volume name OR Ae name) ]"},
	{Text: "-description", Description: "[ -description brief description ]"},
}

var alarm_clear = []prompt.Suggest{
	{Text: "-cluster", Description: "[ -cluster cluster_name ]"},
	{Text: "-alarm", Description: "-alarm alarm "},
	{Text: "-entity", Description: "[ -entity entity (hostname OR volume name OR Ae name) ]"},
}
var alarm_clearmulti = []prompt.Suggest{
	{Text: "-cluster", Description: "[ -cluster cluster_name ]"},
	{Text: "-alarm", Description: "-alarm alarm[:entity][:aetype] <comma seperated alarms>"},
}
var alarm_mute = []prompt.Suggest{
	{Text: "-cluster", Description: "[ -cluster cluster_name ]"},
	{Text: "-alarm", Description: "-alarm alarm[:entity][:aetype][:mute_duration] <comma seperated alarms> "},
	{Text: "-muteminutes", Description: "[ -muteminutes duration (in minutes) to mute the alarms ]"},
}
var alarm_unmute = []prompt.Suggest{
	{Text: "-cluster", Description: "[ -cluster cluster_name ]"},
	{Text: "-alarm", Description: "-alarm [alarm[:entity[:aetype]]]+ "},
}
var alarm_clearall = []prompt.Suggest{
	{Text: "-cluster", Description: "[ -cluster cluster_name ]"},
}
var alarm_config = []prompt.Suggest{
	{Text: "load", Description: "load"},
	{Text: "save", Description: "save"},
}

var alarm_config_load = []prompt.Suggest{
	{Text: "-cluster", Description: "[ -cluster cluster_name ]"},
	{Text: "-output", Description: "[ -output output. default: verbose ]"},
}

var alarm_config_save = []prompt.Suggest{
	{Text: "-cluster", Description: "[ -cluster cluster_name ]"},
	{Text: "-values", Description: "-values values. default:"},
}

var alarm_names = []prompt.Suggest{}

var alarm_groups = []prompt.Suggest{
	{Text: "listGroup", Description: "listGroup"},
	{Text: "addEmails", Description: "addEmails"},
	{Text: "deleteEmails", Description: "deleteEmails"},
	{Text: "addAlarms", Description: "addAlarms"},
	{Text: "deleteAlarms", Description: "deleteAlarms"},
}

var alarm_groups_listGroup = []prompt.Suggest{
	{Text: "-cluster", Description: "[ -cluster cluster_name ]"},
	{Text: "-start", Description: "[ -start start. default: 0 ]"},
	{Text: "-limit", Description: "[ -limit limit. default: 2147483647 ]"},
	{Text: "-output", Description: "[ -output output. default: verbose ]"},
}

// {Text:"", Description:"

var alarm_groups_addEmails = []prompt.Suggest{
	{Text: "-cluster", Description: "[ -cluster cluster_name ]"},
	{Text: "-groupname", Description: "-groupname group name"},
	{Text: "-emails", Description: "-emails email addresses"},
}

var alarm_groups_deleteEmails = []prompt.Suggest{
	{Text: "-cluster", Description: "[ -cluster cluster_name ]"},
	{Text: "-groupname", Description: "-groupname group name"},
	{Text: "-emails", Description: "-emails email addresses"},
}

var alarm_groups_addAlarms = []prompt.Suggest{
	{Text: "-cluster", Description: "[ -cluster cluster_name ]"},
	{Text: "-groupname", Description: "-groupname group name"},
	{Text: "-alarms", Description: "-alarms alarm names"},
}

var alarm_groups_deleteAlarms = []prompt.Suggest{
	{Text: "-cluster", Description: "[ -cluster cluster_name ]"},
	{Text: "-groupname", Description: "-groupname group name"},
	{Text: "-alarms", Description: "-alarms alarm names"},
}
