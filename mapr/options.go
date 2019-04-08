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

import prompt "github.com/c-bata/go-prompt"

var commands = []prompt.Suggest{
	{Text: "acerole", Description: "acerole [validate] -role [role_name]"},
	{Text: "acl", Description: "acl [show|set|edit] -type [cluster|volume] -name <volume name>"},
	{Text: "alarm", Description: "alarm [list|raise|clear|clearall|mute|unmute|config|names|group]"},
	{Text: "audit", Description: "audit [data|cluster|info]"},
	{Text: "blacklist", Description: "blacklist [user|listusers]"},
	{Text: "cluster", Description: "cluster [mapreduce [get|set] | gateway | feature [enable|list] | info getminmax]"},
	{Text: "config", Description: "config [load|save]"},
	{Text: "dashboard", Description: "dashboard [info] -[version|multi_cluster_info]"},
	{Text: "debugdb", Description: "debugdb cdscan|bmap|checkTablet|rawScan|multiOp"},
	{Text: "dialhome", Description: "dialhome [metrics|enable|status|ackdial|lastdialed]"},
	{Text: "disk", Description: "disk [list|listall|add|remove]"},
	{Text: "dump", Description: "dump [containerinfo|containers|volumeinfo|volumenodes|replicationmanagerinfo|replicationmanagerqueueinfo|rereplicationinfo|rereplicationmetrics|balancerinfo|balancermetrics|rolebalancerinfo|rolebalancermetrics|assignvouchers|fileserverworkinfo|zkinfo|supportdump|cldbnodes]snapshotsizeupdateinfo"},
	{Text: "entity", Description: "entity [list|info|modify|remove]"},
	{Text: "fid", Description: "fid [dump|stat]"},
	{Text: "file", Description: "file [offload|recall|tierjobabort|tierjobstatus|tierstatus]"},
	{Text: "job", Description: "job [kill|changepriority|linklogs|status]"},
	{Text: "license", Description: "license [add|remove|list|apps|showid|addcrl|listcrl]"},
	{Text: "nagios", Description: "nagios [generate]"},
	{Text: "nfs4mgmt", Description: "nfs4mgmt [add-export|remove-export|list-exports]"},
	{Text: "nfsmgmt", Description: "nfsmgmt [refreshexports|refreshgidcache]"},
	{Text: "node", Description: "node [list|move|modify|allow-into-cluster|services|topo|remove|heatmap|listcldbs|listzookeepers|cldbmaster|canremovesp|maintenance|metrics|show|listcldbzks|refill]"},
	{Text: "rlimit", Description: "rlimit [set|get] -resource <resource>"},
	{Text: "schedule", Description: "schedule [create|modify|list|remove]"},
	{Text: "security", Description: "security [genkey|genticket|getmapruserticket]"},
	{Text: "service", Description: "service [list|info]"},
	{Text: "setloglevel", Description: "setloglevel [cldb|fileserver|nfs|jobtracker|tasktracker|hbmaster|hbregionserver]"},
	{Text: "stream", Description: "stream [create|edit|info|delete|topic|cursor|replica|upstream|listrecent]"},
	{Text: "table", Description: "table [create|edit|delete|listrecent|info|cf|region|replica|upstream|index|changelog]"},
	{Text: "task", Description: "task [killattempt|failattempt]"},
	{Text: "tier", Description: "tier [create|modify|remove|info|list|rule]"},
	{Text: "trace", Description: "trace [dump|setmode|setlevel|resize|reset|info|print] -host x.x.x.x -port <port>"},
	{Text: "urls", Description: "getting service webserver link"},
	{Text: "virtualip", Description: "virtualip [add|remove|edit|move|list]"},
	{Text: "volume", Description: "volume [create|modify|remove|mount|unmount|list|snapshot|mirror|dump|rename|upgradeformat|move|link|info|showmounts|fixmountpath|offload|recall|compact|tierjobabort|tierjobstatus|tierstats]"},
}

var acerole_validate = []prompt.Suggest{
	{Text: "-role", Description: "-role role which need to be validated "},
}

var acl = []prompt.Suggest{
	{Text: "show", Description: "show"},
	{Text: "userperms", Description: "userperms"},
	{Text: "set", Description: "set"},
	{Text: "edit", Description: "ecit"},
}

var acl_show = []prompt.Suggest{
	{Text: "-type", Description: "-type object type [cluster|volume]"},
	{Text: "-name", Description: "-name name"},
	{Text: "-cluster", Description: "-cluster cluster name"},
	{Text: "-user", Description: "-user userName whose ACL is queried"},
	{Text: "-group", Description: "-group groupName whose ACL is queried "},
	{Text: "-output", Description: "-output output format short|long|terse (default short). default: short"},
	{Text: "-perm", Description: "-perm list of available permissions Parameter takes no value"},
}

var acl_userperms = []prompt.Suggest{
	{Text: "-type", Description: "-type object type [cluster|volume]"},
	{Text: "-name", Description: "-name name"},
	{Text: "-cluster", Description: "-cluster cluster name"},
	{Text: "-user", Description: "-user userName whose permissions are queried"},
}

var acl_set = []prompt.Suggest{
	{Text: "-type", Description: "-type object type [cluster|volume]"},
	{Text: "-name", Description: "-name name"},
	{Text: "-cluster", Description: "-cluster cluster name"},
	{Text: "-user", Description: "-user space separated list of user:permissions,perimssions,.. to be set "},
	{Text: "-group", Description: "-group space separated list of group:permissions,permissions,... to be set "},
}

var acl_edit = []prompt.Suggest{
	{Text: "-type", Description: "-type object type [cluster|volume]"},
	{Text: "-name", Description: "-name name"},
	{Text: "-cluster", Description: "-cluster cluster name"},
	{Text: "-user", Description: "-user space separated list of user:permissions,perimssions,.. to be changed "},
	{Text: "-group", Description: "-group space separated list of group:permissions,permissions,... to be changed "},
}

var alarm = []prompt.Suggest{
	{Text: "list", Description: "list"},
	{Text: "raise", Description: "raise"},
	{Text: "clear", Description: "clear"},
	{Text: "clearmulti", Description: "clearmulti"},
	{Text: "mute", Description: "mute"},
	{Text: "unmute", Description: "unmute"},
	{Text: "clearall", Description: "clearall"},
	{Text: "config", Description: "config"},
	{Text: "names", Description: "names"},
	{Text: "groups", Description: "groups"},
}

var dashboard = []prompt.Suggest{
	{Text: "info", Description: "info"},
}

// {Text:"", Description:"
var dashboard_info = []prompt.Suggest{
	{Text: "-cluster", Description: "[ -cluster cluster name ]"},
	{Text: "-zkconnect", Description: "[ -zkconnect ZooKeeper Connect String: 'host:port,host:port,host:port,...' ]"},
	{Text: "-version", Description: "[ -version true|false. default: false ]"},
	{Text: "-multi_cluster_info", Description: "[ -multi_cluster_info true|false. default: false ]"},
}

var disk = []prompt.Suggest{
	{Text: "listall", Description: "listall"},
	{Text: "list", Description: "list"},
	{Text: "add", Description: "add"},
	{Text: "remove", Description: "remove"},
}

var disk_listall = []prompt.Suggest{
	{Text: "-cluster", Description: "cluster name"},
	{Text: "-start", Description: "-start index of the first node (starting from 0). default: 0"},
	{Text: "-limit", Description: "-limit number of nodes to query. default: 2147483647"},
	{Text: "-output", Description: "-output <terse|verbose>. default: verbose"},
}

var disk_list = []prompt.Suggest{
	{Text: "-host", Description: "name/ip"},
	{Text: "-output", Description: "-output <terse|verbose>. default: verbose"},
	{Text: "-startdisk", Description: "-startdisk index of the first node (starting from 0). default: 0"},
	{Text: "-limitdisk", Description: "-limit number of nodes to query. default: 2147483647"},
	{Text: "-sortby", Description: "-sortby <hostname|diskname|mount|vendor|modelnum|serialnum|firmwareversion|totalspace|usedspace|availablespace|fstype|powerstatus|status|errormsg|storagepoolid|failuretime>"},
	{Text: "-sortorder", Description: "-sortorder <asc|desc>"},
}

var disk_add = []prompt.Suggest{
	{Text: "-host", Description: "name/ip"},
	{Text: "-disks", Description: "-disks comma-separated list of disks"},
	{Text: "-stripeWidth", Description: "-stripeWidth stripe-width"},
}

var disk_remove = []prompt.Suggest{
	{Text: "-host", Description: "name/ip"},
	{Text: "-disks", Description: "-disks comma-separated list of disks"},
	{Text: "-force", Description: "-force <true|false OR 1|0>. Need this parameter to actually remove the disk, otherwise this command behaves like a test remove. default: false"},
	{Text: "-cluster", Description: "cluster name"},
}

var entity = []prompt.Suggest{
	{Text: "list", Description: "list"},
	{Text: "info", Description: "info"},
	{Text: "modify", Description: "modify"},
	{Text: "remove", Description: "remove"},
}

// {Text:"", Description:"

var entity_list = []prompt.Suggest{
	{Text: "-output", Description: "[ -output <terse|verbose>. default: verbose ]"},
	{Text: "-start", Description: "[ -start start. default: 0 ]"},
	{Text: "-limit", Description: "[ -limit limit. default: 2147483647 ]"},
	{Text: "-filter", Description: "[ -filter none. default: none ]"},
	{Text: "-columns", Description: "[ -columns none. default: none ]"},
	{Text: "-alarmedentities", Description: "[ -alarmedentities true|false. default: false ]"},
	{Text: "-cluster", Description: "[ -cluster cluster name ]"},
	{Text: "-sortby", Description: "[ -sortby <entityname|entitytype|entityquota|entityadvisoryquota|entitydiskusage|entityvolumecount> ]"},
}

// {Text:"", Description:"
var entity_info = []prompt.Suggest{
	{Text: "-type", Description: "-type type. default: false"},
	{Text: "-name", Description: "-name name"},
	{Text: "-output", Description: "[ -output <terse|verbose>. default: verbose ]"},
	{Text: "-cluster", Description: "[ -cluster cluster name ]"},
}

var entity_modify = []prompt.Suggest{
	{Text: "-type", Description: "-type type"},
	{Text: "-name", Description: "[ -name entityname ]"},
	{Text: "-entities", Description: "[ -entities entities (0:<user1>,0:<user2>,1:<group1>,1:<group2>..) ]"},
	{Text: "-email", Description: "[ -email email ]"},
	{Text: "-quota", Description: "[ -quota quota ]"},
	{Text: "-advisoryquota", Description: "[ -advisoryquota advisory quota ]"},
	{Text: "-cluster", Description: "[ -cluster cluster name ]"},
}

var entity_remove = []prompt.Suggest{
	{Text: "-name", Description: "[ -name entityname ]"},
	{Text: "-type", Description: "-type type"},
	{Text: "-entities", Description: "[ -entities entities (0:<user1>,0:<user2>,1:<group1>,1:<group2>..) ]"},
}

var table = []prompt.Suggest{
	{Text: "create", Description: "create"},
	{Text: "edit", Description: "edit"},
	{Text: "delete", Description: "delete"},
	{Text: "info", Description: "info"},
	{Text: "cf", Description: "column family"},
	{Text: "region", Description: "region"},
	{Text: "replica", Description: "replica"},
	{Text: "upstream", Description: "upstream"},
	{Text: "index", Description: "index"},
	{Text: "changelog", Description: "changelog"},
}

var stream = []prompt.Suggest{
	{Text: "create", Description: "create"},
	{Text: "edit", Description: "edit"},
	{Text: "info", Description: "info"},
	{Text: "delete", Description: "delete"},
	{Text: "purge", Description: "purge"},
	{Text: "topic", Description: "topic"},
	{Text: "cursor", Description: "cursor"},
	{Text: "assign", Description: "assign"},
	{Text: "replica", Description: "replica"},
	{Text: "upstream", Description: "upstream"},
	{Text: "listrecent", Description: "listrecent"},
}

var volume = []prompt.Suggest{
	{Text: "audit", Description: "audit"},
	{Text: "create", Description: "create"},
	{Text: "unmount", Description: "unmount"},
	{Text: "mount", Description: "mount"},
	{Text: "showmounts", Description: "showmounts"},
	{Text: "offload", Description: "offload"},
	{Text: "recall", Description: "recall"},
	{Text: "compact", Description: "compact"},
	{Text: "tierjobabort", Description: "tierjobabort"},
	{Text: "tierjobstatus", Description: "tierjobstatus"},
	{Text: "tierstats", Description: "tierstats"},
	{Text: "remove", Description: "remove"},
	{Text: "move", Description: "move"},
	{Text: "rename", Description: "rename"},
	{Text: "upgradeformat", Description: "upgradeformat"},
	{Text: "modify", Description: "modify"},
	{Text: "info", Description: "info"},
	{Text: "list", Description: "list"},
	{Text: "fixmountpath", Description: "fixmountpath"},
	{Text: "link", Description: "link"},
	{Text: "snapshot", Description: "snapshot"},
	{Text: "mirror", Description: "mirror"},
	{Text: "dump", Description: "dump"},
	{Text: "container", Description: "container"},
	{Text: "balancecontainers", Description: "balancecontainers"},
	{Text: "balancinginfo", Description: "balancinginfo"},
}
