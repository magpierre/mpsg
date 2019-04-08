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

var volume_audit = []prompt.Suggest{
	{Text: "-cluster", Description: "[ -cluster cluster_name ]"},
	{Text: "-name", Description: "-name volumeName "},
	{Text: "-enabled", Description: "[ -enabled <true|false> ]"},
	{Text: "-forceenable", Description: "[ -forceenable <true|false> ]"},
	{Text: "-coalesce", Description: "[ -coalesce interval in mins ]"},
	{Text: "-dataauditops", Description: "[ -dataauditops data audit operations ]"},
}
var volume_create = []prompt.Suggest{
	{Text: "-cluster", Description: "[ -cluster cluster_name ]"},
	{Text: "-name", Description: "-name volumeName "},
	{Text: "-path", Description: "[ -path mountdir ]"},
	{Text: "-tenantuser", Description: "[ -tenantuser tenantUser ]"},
	{Text: "-createparent", Description: "[ -createparent createparent. default: false ]"},
	{Text: "-mount", Description: "[ -mount mount. default: true ]"},
	{Text: "-rootdirperms", Description: "[ -rootdirperms rootdirperms ]"},
	{Text: "-rootdiruser", Description: "[ -rootdiruser rootdiruser ]"},
	{Text: "-rootdirgroup", Description: "[ -rootdirgroup rootdirgroup ]"},
	{Text: "-rereplicationtimeoutsec", Description: "[ -rereplicationtimeoutsec rereplicationtimeoutsec ]"},
	{Text: "-criticalrereplicationtimeoutsec", Description: "[ -criticalrereplicationtimeoutsec criticalrereplicationtimeoutsec ]"},
	{Text: "-localvolumehost", Description: "[ -localvolumehost localvolumehost ]"},
	{Text: "-localvolumeport", Description: "[ -localvolumeport localvolumeport. default: 5660 ]"},
	{Text: "-replication", Description: "[ -replication replication ]"},
	{Text: "-minreplication", Description: "[ -minreplication minreplication ]"},
	{Text: "-nsreplication", Description: "[ -nsreplication nsreplication ]"},
	{Text: "-nsminreplication", Description: "[ -nsminreplication nsminreplication ]"},
	{Text: "-enforceminreplicationforio", Description: "[ -enforceminreplicationforio stall io when replication falls below min replication, default:false ]"},
	{Text: "-replicationtype", Description: "[ -replicationtype low_latency or high_throughput. default: high_throughput ]"},
	{Text: "-user", Description: "[ -user space separated list of user:permissions,perimssions,.. to be set ]"},
	{Text: "-group", Description: "[ -group space separated list of user:permissions,perimssions,.. to be set ]"},
	{Text: "-aetype", Description: "[ -aetype 0(user) or 1(group) default: 0 ]"},
	{Text: "-ae", Description: "[ -ae ae ]"},
	{Text: "-quota", Description: "[ -quota quota ]"},
	{Text: "-advisoryquota", Description: "[ -advisoryquota advisoryquota ]"},
	{Text: "-topology", Description: "[ -topology topology ]"},
	{Text: "-readonly", Description: "[ -readonly readonly ]"},
	{Text: "-mirrorthrottle", Description: "[ -mirrorthrottle mirrorthrottle ]"},
	{Text: "-type", Description: "[ -type type of volume: rw or mirror ]"},
	{Text: "-source", Description: "[ -source source ]"},
	{Text: "-schedule", Description: "[ -schedule schedule ID ]"},
	{Text: "-mirrorschedule", Description: "[ -mirrorschedule mirror schedule ID ]"},
	{Text: "-maxinodesalarmthreshold", Description: "[ -maxinodesalarmthreshold maxinodesalarmthreshold ]"},
	{Text: "-maxnssizembalarmthreshold", Description: "[ -maxnssizembalarmthreshold maxnssizembalarmthreshold ]"},
	{Text: "-dbrepllagsecalarmthresh", Description: "[ -dbrepllagsecalarmthresh dbrepllagsecalarmthresh ]"},
	{Text: "-dbindexlagsecalarmthresh", Description: "[ -dbindexlagsecalarmthresh dbindexlagsecalarmthresh ]"},
	{Text: "-auditenabled", Description: "[ -auditenabled <true|false> ]"},
	{Text: "-forceauditenable", Description: "[ -forceauditenable <true|false> ]"},
	{Text: "-coalesce", Description: "[ -coalesce interval in mins ]"},
	{Text: "-dataauditops", Description: "[ -dataauditops data audit operations ]"},
	{Text: "-wiresecurityenabled", Description: "[ -wiresecurityenabled <true|false> ]"},
	{Text: "-skipwiresecurityfortierinternalops", Description: "[ -skipwiresecurityfortierinternalops Skip Wire level security for backend volumes <true|false> ]"},
	{Text: "-allowgrant", Description: "[ -allowgrant allowgrant ]"},
	{Text: "-inherit", Description: "[ -inherit volume to copy properties from: defaults to parent volume ]"},
	{Text: "-allowinherit", Description: "[ -allowinherit allowinherit ]"},
	{Text: "-skipinherit", Description: "[ -skipinherit volume properties not to inherit ]"},
	{Text: "-readAce", Description: "[ -readAce <Acess Control Expression> ]"},
	{Text: "-writeAce", Description: "[ -writeAce <Acess Control Expression> ]"},
	{Text: "-metricsenabled", Description: "[ -metricsenabled <true|false> ]"},
	{Text: "-containerallocationfactor", Description: "[ -containerallocationfactor <+ve integer, default:0> ]"},
	{Text: "-namecontainerdatathreshold", Description: "[ -namecontainerdatathreshold name container data size in MB ]"},
	{Text: "-dare", Description: "[ -dare Enable Data at rest encryption for volume. <true|false> ]"},
	{Text: "-tieringenable ", Description: "[ -tieringenable <true|false> ]"},
	{Text: "-ecenable", Description: "[ -ecenable <true|false> ]"},
	{Text: "-tiername", Description: "[ -tiername tier name ]"},
	{Text: "-tieringrule", Description: "[ -tieringrule tiering rule ]"},
	{Text: "-offloadschedule", Description: "[ -offloadschedule tiering schedule ID ]"},
	{Text: "-tierkey", Description: "[ -tierkey tier encryption key ]"},
	{Text: "-tierencryption", Description: "[ -tierencryption <true|false> ]"},
	{Text: "-recallexpirytime", Description: "[ -recallexpirytime recallExpiryTime in days < Low Boundary: 1. Upper Boundary: 7500> ]"},
	{Text: "-ecscheme", Description: "[ -ecscheme ec scheme ]"},
	{Text: "-ectopology", Description: "[ -ectopology ec volume topology ]"},
	{Text: "-autooffloadthresholdgb", Description: "[ -autooffloadthresholdgb <ec auto offload size threshold in GB> ]"},
	{Text: "-compactionschedule", Description: "[ -compactionschedule compaction schedule Id ]"},
	{Text: "-compactionoverheadthreshold", Description: "[ -compactionoverheadthreshold compaction overhead ]"},
}
var volume_unmount = []prompt.Suggest{
	{Text: "-cluster", Description: "[ -cluster cluster_name ]"},
	{Text: "-name", Description: "-name name "},
	{Text: "-force", Description: "[ -force force. default: false ]"},
}
var volume_mount = []prompt.Suggest{
	{Text: "-cluster", Description: "[ -cluster cluster_name ]"},
	{Text: "-name", Description: "-name name "},
	{Text: "-path", Description: "[ -path path ]"},
	{Text: "-createparent", Description: "[ -createparent createparent. default: false ]"},
}
var volume_showmounts = []prompt.Suggest{
	{Text: "-cluster", Description: "[ -cluster cluster_name ]"},
	{Text: "-name", Description: "-name name "},
}
var volume_offload = []prompt.Suggest{
	{Text: "-cluster", Description: "[ -cluster cluster_name ]"},
	{Text: "-name", Description: "-name name "},
	{Text: "-ignorerule", Description: "[ -ignorerule <true|false>. default: false ]"},
}
var volume_recall = []prompt.Suggest{
	{Text: "-cluster", Description: "[ -cluster cluster_name ]"},
	{Text: "-name", Description: "-name name "},
}
var volume_compact = []prompt.Suggest{
	{Text: "-cluster", Description: "[ -cluster cluster_name ]"},
	{Text: "-name", Description: "-name name"},
	{Text: "-forcerecallexpiry", Description: "[ -forcerecallexpiry <true|false>. default: false ]"},
}
var volume_tierjobabort = []prompt.Suggest{
	{Text: "-cluster", Description: "[ -cluster cluster_name ]"},
	{Text: "-name", Description: "-name name"},
}
var volume_tierjobstatus = []prompt.Suggest{
	{Text: "-cluster", Description: "[ -cluster cluster_name ]"},
	{Text: "-name", Description: "-name name"},
	{Text: "-verbose", Description: "[ -verbose <true|false>. default: false ]"},
}
var volume_tierstats = []prompt.Suggest{
	{Text: "-cluster", Description: "[ -cluster cluster_name ]"},
	{Text: "-name", Description: "-name name"},
}
var volume_remove = []prompt.Suggest{
	{Text: " -cluster", Description: "[ -cluster cluster_name ]"},
	{Text: "-name", Description: "[ -name volumeName ]"},
	{Text: "-force", Description: "[ -force force. default: false ]"},
	{Text: "-filter", Description: "[ -filter remove volumes that match the filter. default: none ]"},
}
var volume_move = []prompt.Suggest{
	{Text: "-cluster", Description: "[ -cluster cluster_name ]"},
	{Text: "-name", Description: "-name volumeName>"},
	{Text: "-topology", Description: "[ -topology topology ]"},
	{Text: "-ectopology", Description: "[ -ectopology topology for ec-store volume ]"},
}
var volume_rename = []prompt.Suggest{
	{Text: "-cluster", Description: "[ -cluster cluster_name ]"},
	{Text: "-name", Description: "-name volumeName"},
	{Text: "-newname", Description: "-newname newVolumeName"},
}
var volume_upgradeformat = []prompt.Suggest{
	{Text: "-cluster", Description: "[ -cluster cluster_name ]"},
	{Text: "-name", Description: "-name volumeNames"},
}
var volume_modify = []prompt.Suggest{
	{Text: "-cluster", Description: "[ -cluster cluster_name ]"},
	{Text: "-name", Description: "-name volumeName"},
	{Text: "-source", Description: "[ -source source ]"},
	{Text: "-replication", Description: "[ -replication replication ]"},
	{Text: "-minreplication", Description: "[ -minreplication minreplication ]"},
	{Text: "-nsreplication", Description: "[ -nsreplication nsreplication ]"},
	{Text: "-nsminreplication", Description: "[ -nsminreplication nsminreplication ]"},
	{Text: "-enforceminreplicationforio", Description: "[ -enforceminreplicationforio enforceminreplicationforio: stall io when replication falls below min replication, default:false ]"},
	{Text: "-containerallocationfactor", Description: "[ -containerallocationfactor <+ve integer, default:0> ]"},
	{Text: "-user", Description: "[ -user space separated list of user:permissions,perimssions,.. to be set ]"},
	{Text: "-group", Description: "[ -group space separated list of user:permissions,perimssions,.. to be set ]"},
	{Text: "-aetype", Description: "[ -aetype aetype ]"},
	{Text: "-ae", Description: "[ -ae ae ]"},
	{Text: "-quota", Description: "[ -quota quota ]"},
	{Text: "-advisoryquota", Description: "[ -advisoryquota advisoryquota ]"},
	{Text: "-readonly", Description: "[ -readonly readonly ]"},
	{Text: "-mirrorthrottle", Description: "[ -mirrorthrottle mirrorthrottle ]"},
	{Text: "-schedule", Description: "[ -schedule schedule ID ]"},
	{Text: "-mirrorschedule", Description: "[ -mirrorschedule mirror schedule ID ]"},
	{Text: "-maxinodesalarmthreshold", Description: "[ -maxinodesalarmthreshold maxinodesalarmthreshold ]"},
	{Text: "-maxnssizembalarmthreshold", Description: "[ -maxnssizembalarmthreshold maxnssizembalarmthreshold ]"},
	{Text: "-dbrepllagsecalarmthresh", Description: "[ -dbrepllagsecalarmthresh dbrepllagsecalarmthresh ]"},
	{Text: "-dbindexlagsecalarmthresh", Description: "[ -dbindexlagsecalarmthresh dbindexlagsecalarmthresh ]"},
	{Text: "-type", Description: "[ -type type of volume: rw or mirror ]"},
	{Text: "-auditenabled", Description: "[ -auditenabled <true|false> ]"},
	{Text: "-forceauditenable", Description: "[ -forceauditenable <true|false> ]"},
	{Text: "-coalesce", Description: "[ -coalesce interval in mins ]"},
	{Text: "-dataauditops", Description: "[ -dataauditops data audit operations ]"},
	{Text: "-wiresecurityenabled", Description: "[ -wiresecurityenabled <true|false> ]"},
	{Text: "-skipwiresecurityfortierinternalops", Description: "[ -skipwiresecurityfortierinternalops Skip Wire level security for backend volumes <true|false> ]"},
	{Text: "-allowgrant", Description: "[ -allowgrant let child volume inherit volume properties <true|false> ]"},
	{Text: "-rereplicationtimeoutsec", Description: "[ -rereplicationtimeoutsec rereplicationtimeoutsec ]"},
	{Text: "-criticalrereplicationtimeoutsec", Description: "[ -criticalrereplicationtimeoutsec criticalrereplicationtimeoutsec ]"},
	{Text: "-readAce", Description: "[ -readAce <Acess Control Expression> ]"},
	{Text: "-writeAce", Description: "[ -writeAce <Acess Control Expression> ]"},
	{Text: "-metricsenabled", Description: "[ -metricsenabled <true|false> ]"},
	{Text: "-namecontainerdatathreshold", Description: "[ -namecontainerdatathreshold name container data size in MB ]"},
	{Text: "-tiername", Description: "[ -tiername tier name ]"},
	{Text: "-ecenable", Description: "[ -ecenable <true|false> ]"},
	{Text: "-tieringrule", Description: "[ -tieringrule tiering rule ]"},
	{Text: "-offloadschedule", Description: "[ -offloadschedule tiering schedule ID ]"},
	{Text: "-tierkey", Description: "[ -tierkey tier encryption key ]"},
	{Text: "-tierencryption", Description: "[ -tierencryption <true|false> ]"},
	{Text: "-recallexpirytime", Description: "[ -recallexpirytime recallExpiryTime in days < Low Boundary: 1. Upper Boundary: 7500> ]"},
	{Text: "-compactionschedule", Description: "[ -compactionschedule compaction schedule Id ]"},
	{Text: "-compactionoverheadthreshold", Description: "[ -compactionoverheadthreshold compaction overhead ]"},
	{Text: "-ecscheme", Description: "[ -ecscheme ec scheme ]"},
	{Text: "-ectopology", Description: "[ -ectopology ec volume topology ]"},
	{Text: "-autooffloadthresholdgb", Description: "[ -autooffloadthresholdgb <ec auto offload size threshold in GB> ]"},
}
var volume_info = []prompt.Suggest{
	{Text: "-cluster", Description: "[ -cluster cluster_name ]"},
	{Text: "-output", Description: "[ -output verbose. default: verbose ]"},
	{Text: "-path", Description: "[ -path mountdir ]"},
	{Text: "-name", Description: "[ -name volumeName ]"},
	{Text: "-columns", Description: "[ -columns comma separated list of column names. default: all ]"},
}
var volume_list = []prompt.Suggest{
	{Text: "-cluster", Description: "[ -cluster cluster_name ]"},
	{Text: "-output", Description: "[ -output verbose. default: verbose ]"},
	{Text: "-start", Description: "[ -start start. default: 0 ]"},
	{Text: "-limit", Description: "[ -limit limit. default: 2147483647 ]"},
	{Text: "-filter", Description: "[ -filter none. default: none ]"},
	{Text: "-nodes", Description: "[ -nodes nodes ]"},
	{Text: "-columns", Description: "[ -columns comma separated list of column names. default: all ]"},
	{Text: "-alarmedvolumes", Description: "[ -alarmedvolumes alarmsonly. default: false ]"},
	{Text: "-sortby", Description: "[ -sortby <volumeowner|volumenumreplicas|volumeminreplicas|volumerackpath|volumemountdir|volumename|volumequota|volumeused|volumequotaadvisory|volumeaename|volumeaetype|volumeschedule|volumetype|volumemirrortype|volumemirrorpercentcomplete|volumesnapshotcount|volumeid|volumenamecontainersize|volumelocalpath|volumesnapshotused|volumetotalused|volumelogicalused|volumecontainercount|volumemirrorschedule|volumeaccesstime|volumenamespacecontainernumreplicas|volumenamespacecontainerminreplicas|volumerereplicationtimeoutsec|volumecriticalrereplicationtimeoutsec|volumecreatetime|volumedareenabled>, column names of supported fields. ]"},
	{Text: "-sortorder", Description: "[ -sortorder <asc|desc> ]"},
}
var volume_fixmountpath = []prompt.Suggest{
	{Text: "-cluster", Description: "[ -cluster cluster_name ]"},
	{Text: "-name", Description: "-name volumeName"},
}
var volume_link = []prompt.Suggest{
	{Text: "create", Description: "create"},
	{Text: "remove", Description: "remove"},
}
var volume_link_create = []prompt.Suggest{
	{Text: "-cluster", Description: "[ -cluster cluster_name ]"},
	{Text: "-volume", Description: "-volume volume"},
	{Text: "-type", Description: "-type type <writeable|mirror>"},
	{Text: "-path", Description: "-path path"},
}
var volume_link_remove = []prompt.Suggest{
	{Text: "-cluster", Description: "[ -cluster cluster_name ]"},
	{Text: "-path", Description: "-path vollink"},
}

var volume_snapshot = []prompt.Suggest{
	{Text: "create", Description: "create"},
	{Text: "remove", Description: "remove"},
	{Text: "preserve", Description: "preserve"},
	{Text: "list", Description: "list"},
}

var volume_snapshot_create = []prompt.Suggest{
	{Text: "-cluster", Description: "[ -cluster cluster_name ]"},
	{Text: "-retain", Description: "[ -retain <+ve integer>mi|h|d|w|m|y ]"},
	{Text: "-snapshotname", Description: " -snapshotname snapshotName"},
	{Text: "-volume", Description: "-volume volume"},
}
var volume_snapshot_remove = []prompt.Suggest{
	{Text: "-cluster", Description: "[ -cluster cluster name ]"},
	{Text: "-snapshotname", Description: "[ -snapshotname snapshotName ]"},
	{Text: "-volume", Description: "[ -volume volumeName ]"},
	{Text: "-snapshots", Description: "[ -snapshots comma separated IDs of snapshots ]"},
}
var volume_snapshot_preserve = []prompt.Suggest{
	{Text: "-cluster", Description: "[ -cluster cluster name ]"},
	{Text: "-volume", Description: "[ -volume comma separated volume names to preserve snapshots for ]"},
	{Text: "-path", Description: "[ -path comma separated volume pathes to preserve snapshots for ]"},
	{Text: "-filter", Description: "[ -filter filter to preserve snapshots for ]"},
	{Text: "-snapshots", Description: "[ -snapshots comma separated IDs of snapshots to preserve ]"},
}
var volume_snapshot_list = []prompt.Suggest{
	{Text: "-cluster", Description: "[ -cluster cluster name ]"},
	{Text: "-volume", Description: "[ -volume volume ]"},
	{Text: "-path", Description: "[ -path path ]"},
	{Text: "-output", Description: "[ -output verbose. default: verbose ]"},
	{Text: "-start", Description: "[ -start start. default: 0 ]"},
	{Text: "-limit", Description: "[ -limit limit. default: 2147483647 ]"},
	{Text: "-filter", Description: "[ -filter none. default: none ]"},
	{Text: "-columns", Description: "[ -columns none. default: none ]"},
}

var volume_mirror = []prompt.Suggest{
	{Text: "start", Description: "start"},
	{Text: "stop", Description: "stop"},
	{Text: "push", Description: "push"},
}
var volume_mirror_start = []prompt.Suggest{
	{Text: "-cluster", Description: "[ -cluster cluster_name ]"},
	{Text: "-name", Description: "-name name"},
	{Text: "-full", Description: "[ -full <true|false>. default: false ]"},
}
var volume_mirror_stop = []prompt.Suggest{
	{Text: "-cluster", Description: "[ -cluster cluster_name ]"},
	{Text: "-name", Description: "-name name"},
}
var volume_mirror_push = []prompt.Suggest{
	{Text: "-cluster", Description: "[ -cluster cluster_name ]"},
	{Text: "-name", Description: "-name name"},
	{Text: "-verbose", Description: "[ -verbose verbose. default: true ]"},
}

var volume_dump = []prompt.Suggest{
	{Text: "create", Description: "create"},
	{Text: "restore", Description: "restore"},
}
var volume_dump_create = []prompt.Suggest{
	{Text: "-cluster", Description: "[ -cluster cluster_name ]"},
	{Text: "-s", Description: "[ -s start volumepoint name ]"},
	{Text: "-e", Description: "[ -e end volumepoint name ]"},
	{Text: "-o", Description: "[ -o (generate dumpfile on stdout) ]"},
	{Text: "-dumpfile", Description: " -dumpfile dumpfilename (not needed if -o is used)"},
	{Text: "-name", Description: " -name volumename "},
}
var volume_dump_restore = []prompt.Suggest{
	{Text: "-cluster", Description: "[ -cluster cluster_name ]"},
	{Text: "-i", Description: "[ -i (read dumpfile from stdin) ]"},
	{Text: "-n", Description: "[ -n (create new volume if it doesn't exist) ]"},
	{Text: "-dumpfile", Description: "-dumpfile dumpfilename (not needed if -i is used)"},
	{Text: "-name", Description: "-name volumename"},
}

var volume_container = []prompt.Suggest{
	{Text: "move", Description: "move"},
	{Text: "switchmaster", Description: "switchmaster"},
	{Text: "switchtailwithupstream", Description: "switchtailwithupstream"},
}
var volume_container_move = []prompt.Suggest{
	{Text: "-cluster", Description: "[ -cluster cluster_name ]"},
	{Text: "-cid", Description: "-cid cid"},
	{Text: "-fromfileserverid", Description: "-fromfileserverid fromfileserverid"},
	{Text: "-tofileserverid", Description: "[ -tofileserverid tofileserverid ]"},
}
var volume_container_switchmaster = []prompt.Suggest{
	{Text: "-cluster", Description: "[ -cluster cluster_name ]"},
	{Text: "-cid", Description: "-cid cid"},
}
var volume_container_switchtailwithupstream = []prompt.Suggest{
	{Text: "-cluster", Description: "[ -cluster cluster_name ]"},
	{Text: "-cid", Description: "-cid cid"},
}

var volume_balancecontainers = []prompt.Suggest{
	{Text: "-cluster", Description: "[ -cluster cluster_name ]"},
	{Text: "-name", Description: "-name volumeName"},
	{Text: "-cancel", Description: "[ -cancel <true|false>. default: false ]"},
}
var volume_balancinginfo = []prompt.Suggest{
	{Text: "-cluster", Description: "[ -cluster cluster_name ]"},
	{Text: "-name", Description: "[ -name volumeName ]"},
}
