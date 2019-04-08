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

// {Text:"", Description:"

var table_create = []prompt.Suggest{
	{Text: "-path", Description: "-path path"},
	{Text: "-copymetafrom", Description: "[ -copymetafrom SrcTablePath ]"},
	{Text: "-copymetatype", Description: "[ -copymetatype all|cfs|aces|splits|attrs ]"},
	{Text: "-regionsizemb", Description: "[ -regionsizemb Region Size in MB ]"},
	{Text: "-autosplit", Description: "[ -autosplit Auto Split table ]"},
	{Text: "-bulkload", Description: "[ -bulkload Bulk load ]"},
	{Text: "-audit", Description: "[ -audit Enable Audit ]"},
	{Text: "-tabletype", Description: "[ -tabletype Table Type - json or binary. default: binary ]"},
	{Text: "-packperm", Description: "[ -packperm Pack Permission settings ]"},
	{Text: "-bulkloadperm", Description: "[ -bulkloadperm Bulk load Permission settings ]"},
	{Text: "-splitmergeperm", Description: "[ -splitmergeperm Split and Merge Permission settings ]"},
	{Text: "-createrenamefamilyperm", Description: "[ -createrenamefamilyperm Add/Rename Family Permission settings ]"},
	{Text: "-deletefamilyperm", Description: "[ -deletefamilyperm Delete Family Permission settings ]"},
	{Text: "-adminaccessperm", Description: "[ -adminaccessperm Ace Admin Permission settings ]"},
	{Text: "-replperm", Description: "[ -replperm Replication Admin Permission settings ]"},
	{Text: "-indexperm", Description: "[ -indexperm Secondary Index Admin Permission settings ]"},
	{Text: "-defaultversionperm", Description: "[ -defaultversionperm CF Versions Default Permission for binary tabletype ]"},
	{Text: "-defaultcompressionperm", Description: "[ -defaultcompressionperm CF Compression Default Permission ]"},
	{Text: "-defaultmemoryperm", Description: "[ -defaultmemoryperm CF Memory Default Permission ]"},
	{Text: "-defaultreadperm", Description: "[ -defaultreadperm CF Read Default Permission ]"},
	{Text: "-defaultwriteperm", Description: "[ -defaultwriteperm CF Write Default Permission ]"},
	{Text: "-defaulttraverseperm", Description: "[ -defaulttraverseperm CF Traverse Default Permission for json tabletype ]"},
	{Text: "-defaultappendperm", Description: "[ -defaultappendperm CF Append Default Permission for binary tabletype ]"},
	{Text: "-metricsinterval", Description: "[ -metricsinterval Metrics collection interval, in seconds ]"},
}

var table_edit = []prompt.Suggest{
	{Text: "-path", Description: "-path path"},
	{Text: "-autosplit", Description: "[ -autosplit Auto Split table ]"},
	{Text: "-regionsizemb", Description: "[ -regionsizemb Region Size in MB ]"},
	{Text: "-bulkload", Description: "[ -bulkload Bulk load ]"},
	{Text: "-audit", Description: "[ -audit Enable Audit ]"},
	{Text: "-deletettl", Description: "[ -deletettl delete TTL in secs ]"},
	{Text: "-packperm", Description: "[ -packperm Pack Permission settings ]"},
	{Text: "-bulkloadperm", Description: "[ -bulkloadperm Bulk load Permission settings ]"},
	{Text: "-splitmergeperm", Description: "[ -splitmergeperm Split and Merge Permission settings ]"},
	{Text: "-createrenamefamilyperm", Description: "[ -createrenamefamilyperm Add/Rename Family Permission settings ]"},
	{Text: "-deletefamilyperm", Description: "[ -deletefamilyperm Delete Family Permission settings ]"},
	{Text: "-adminaccessperm", Description: "[ -adminaccessperm Ace Admin Permission settings ]"},
	{Text: "-replperm", Description: "[ -replperm Replication Admin Permission settings ]"},
	{Text: " -indexperm", Description: "[ -indexperm Secondary Index Admin Permission settings ]"},
	{Text: "-defaultversionperm", Description: "[ -defaultversionperm CF Versions Default Permission for binary tabletype ]"},
	{Text: "-defaultcompressionperm", Description: "[ -defaultcompressionperm CF Compression Default Permission ]"},
	{Text: "-defaultmemoryperm", Description: "[ -defaultmemoryperm CF Memory Default Permission ]"},
	{Text: "-defaultreadperm", Description: "[ -defaultreadperm CF Read Default Permission ]"},
	{Text: "-defaultwriteperm", Description: "[ -defaultwriteperm CF Write Default Permission ]"},
	{Text: "-defaulttraverseperm", Description: "[ -defaulttraverseperm CF Traverse Default Permission for json tabletype ]"},
	{Text: "-defaultappendperm", Description: "[ -defaultappendperm CF Append Default Permission for binary tabletype ]"},
	{Text: "-metricsinterval", Description: "[ -metricsinterval Metrics collection interval, in seconds ]"},
}

var table_delete = []prompt.Suggest{
	{Text: "path", Description: "-path path"},
}

var table_info = []prompt.Suggest{
	{Text: "-path", Description: "-path table path"},
	{Text: "-index", Description: "[ -index index name ]"},
}
var table_cf = []prompt.Suggest{
	{Text: "create", Description: "create"},
	{Text: "edit", Description: "edit"},
	{Text: "delete", Description: "delete"},
	{Text: "list", Description: "list"},
	{Text: "colperm", Description: "colperm"},
}

// {Text:"", Description:"

var table_cf_create = []prompt.Suggest{
	{Text: "-path", Description: "-path Table path"},
	{Text: "-cfname", Description: "-cfname Column family name"},
	{Text: "-minversions", Description: "[ -minversions Min versions to keep. default: 0 ]"},
	{Text: "-maxversions", Description: "[ -maxversions Max versions to keep. default: 1 ]"},
	{Text: "-ttl", Description: "[ -ttl Time to live. Enter 0 for forever. Otherwise enter time in seconds. default: 0 ]"},
	{Text: "-inmemory", Description: "[ -inmemory In-memory. default: false ]"},
	{Text: "-compression", Description: "[ -compression off|lzf|lz4|zlib. default: table's compression setting is applied. ]"},
	{Text: "-versionperm", Description: "[ -versionperm Version Permissions for binary tabletype ]"},
	{Text: "-compressionperm", Description: "[ -compressionperm Compression Permissions ]"},
	{Text: "-memoryperm", Description: "[ -memoryperm Memory Permissions ]"},
	{Text: "-readperm", Description: "[ -readperm Read Permissions ]"},
	{Text: "-writeperm", Description: "[ -writeperm Write Permissions ]"},
	{Text: "-appendperm", Description: "[ -appendperm Append Permissions for binary tabletype ]"},
	{Text: "-traverseperm", Description: "[ -traverseperm Traverse Permissions for json tabletype ]"},
	{Text: "-jsonpath", Description: "[ -jsonpath Json Family Path - needed for JSON column family, like a.b.c ]"},
	{Text: "-force", Description: "[ -force Force create non-default column family for json tabletype. default: false ]"},
}
var table_cf_edit = []prompt.Suggest{
	{Text: "-path", Description: "-path Table path"},
	{Text: "-cfname", Description: "-cfname Column family name"},
	{Text: "-newcfname", Description: "[ -newcfname New column family name ]"},
	{Text: "-minversions", Description: "[ -minversions Min versions to keep ]"},
	{Text: "-maxversions", Description: "[ -maxversions Max versions to keep ]"},
	{Text: "-ttl", Description: "[ -ttl Time to live (in seconds) ]"},
	{Text: "-inmemory", Description: "[ -inmemory In-memory ]"},
	{Text: "-compression", Description: "[ -compression off|lzf|lz4|zlib ]"},
	{Text: "-versionperm", Description: "[ -versionperm Version Permissions for binary tabletype ]"},
	{Text: "-compressionperm", Description: "[ -compressionperm Compression Permissions ]"},
	{Text: "-memoryperm", Description: "[ -memoryperm Memory Permissions ]"},
	{Text: "-readperm", Description: "[ -readperm Read Permissions ]"},
	{Text: "-writeperm", Description: "[ -writeperm Write Permissions ]"},
	{Text: "-appendperm", Description: "[ -appendperm Append Permissions for binary tabletype ]"},
	{Text: "-traverseperm", Description: "[ -traverseperm Traverse Permissions for json tabletype ]"},
}

var table_cf_delete = []prompt.Suggest{
	{Text: "-path", Description: "-path Table path"},
	{Text: "-cfname", Description: "-cfname Column family name "},
}

var table_cf_list = []prompt.Suggest{
	{Text: "-path", Description: "-path Table path"},
	{Text: "-cfname", Description: "[-cfname Column family name ]"},
}

var table_cf_colperm = []prompt.Suggest{
	{Text: "get", Description: "get"},
	{Text: "set", Description: "set"},
	{Text: "delete", Description: "delete"},
}

var table_cf_colperm_get = []prompt.Suggest{
	{Text: "-path", Description: "-path Table path"},
	{Text: "-cfname", Description: "-cfname Column family name"},
	{Text: "-name", Description: "[ -name Column name ]"},
}

var table_cf_colperm_set = []prompt.Suggest{
	{Text: "-path", Description: "-path Table path"},
	{Text: "-cfname", Description: "-cfname Column family name"},
	{Text: "-name", Description: "-name Column name"},
	{Text: "-readperm", Description: "[ -readperm Read column permission settings ]"},
	{Text: "-writeperm", Description: "[ -writeperm Write column permission settings ]"},
	{Text: "-appendperm", Description: "[ -appendperm Append column permission settings ]"},
	{Text: "-traverseperm", Description: "[ -traverseperm Traverse column permission settings ]"},
}

var table_cf_colperm_delete = []prompt.Suggest{
	{Text: "-path", Description: "-path Table path"},
	{Text: "-cfname", Description: "-cfname Column family name"},
	{Text: "-name", Description: "-name Column name "},
}

var table_region = []prompt.Suggest{
	{Text: "split", Description: "split"},
	{Text: "pack", Description: "pack"},
	{Text: "merge", Description: "merge"},
	{Text: "list", Description: "list"},
	{Text: "stat", Description: "stat"},
}

var table_region_split = []prompt.Suggest{
	{Text: "-path", Description: "-path Table path"},
	{Text: "-fid", Description: "fid"},
}
var table_region_pack = []prompt.Suggest{
	{Text: "-path", Description: "-path table path"},
	{Text: "-fid", Description: "-fid fid|all"},
	{Text: "-nthreads", Description: "[ -nthreads nthreads. default: 16 ]"},
}

// {Text:"", Description:"
var table_region_merge = []prompt.Suggest{
	{Text: "-path", Description: "-path Table path"},
	{Text: "-fid", Description: "fid"},
}
var table_region_list = []prompt.Suggest{
	{Text: "-path", Description: "-path table path"},
	{Text: "-start", Description: "[ -start start. default: 0 ]"},
	{Text: "-limit", Description: "[ -limit limit. default: 2147483647 ]"},
	{Text: "-index", Description: "[ -index index name ]"},
}
var table_region_stat = []prompt.Suggest{
	{Text: "-path", Description: "-path table path"},
	{Text: "-index", Description: "[ -index index name ]"},
}

var table_replica = []prompt.Suggest{
	{Text: "add", Description: "add"},
	{Text: "edit", Description: "edit"},
	{Text: "list", Description: "list"},
	{Text: "remove", Description: "remove"},
	{Text: "pause", Description: "pause"},
	{Text: "resume", Description: "resume"},
	{Text: "autosetup", Description: "autosetup"},
}
var table_replica_add = []prompt.Suggest{
	{Text: "-path", Description: "-path table path"},
	{Text: "-replica", Description: "-replica replica table path"},
	{Text: "-columns", Description: "[ -columns comma separated list of <family>[:<column>] ]"},
	{Text: "-paused", Description: "[ -paused is replication paused. default: false ]"},
	{Text: "-throttle", Description: "[ -throttle throttle replication ops. default: false ]"},
	{Text: "-networkencryption", Description: "[ -networkencryption enable on-wire encryption. default: false ]"},
	{Text: "-synchronous", Description: "[ -synchronous is synchronous replication. default: false ]"},
	{Text: "-networkcompression", Description: "[ -networkcompression on-wire compression type: off|on|lzf|lz4|zlib. default: on ]"},
}
var table_replica_edit = []prompt.Suggest{
	{Text: "-path", Description: "-path table path"},
	{Text: "-replica", Description: "-replica replica table path"},
	{Text: "-newreplica", Description: "[ -newreplica renamed table path ]"},
	{Text: "-columns", Description: "[ -columns comma separated list of <family>[:<column>] ]"},
	{Text: "-throttle", Description: "[ -throttle throttle replication ops ]"},
	{Text: "-networkencryption", Description: "[ -networkencryption enable on-wire encryption ]"},
	{Text: "-synchronous", Description: "[ -synchronous is synchronous replication ]"},
	{Text: "-networkcompression", Description: "[ -networkcompression on-wire compression type: off|on|lzf|lz4|zlib ]"},
}
var table_replica_list = []prompt.Suggest{
	{Text: "-path", Description: "-path table path"},
	{Text: "-refreshnow", Description: "[ -refreshnow refreshnow. default: false ]"},
}
var table_replica_remove = []prompt.Suggest{
	{Text: "-path", Description: "-path table path"},
	{Text: "-replica", Description: "-replica replica table path"},
}
var table_replica_pause = []prompt.Suggest{
	{Text: "-path", Description: "-path table path"},
	{Text: "-replica", Description: "-replica replica table path"},
}
var table_replica_resume = []prompt.Suggest{
	{Text: "-path", Description: "-path table path"},
	{Text: "-replica", Description: "-replica replica table path"},
}
var table_replica_autosetup = []prompt.Suggest{
	{Text: "-path", Description: "-path table path"},
	{Text: "-replica", Description: "-replica replica table path"},
	{Text: "-columns", Description: "[ -columns comma separated list of <family>[:<column>] ]"},
	{Text: "-synchronous", Description: "[ -synchronous is synchronous replication. default: false ]"},
	{Text: "-multimaster", Description: "[ -multimaster is multi master replication. default: false ]"},
	{Text: "-throttle", Description: "[ -throttle throttle replication ops. default: false ]"},
	{Text: "-networkencryption", Description: "[ -networkencryption enable on-wire encryption. default: false ]"},
	{Text: "-networkcompression", Description: "[ -networkcompression on-wire compression type: off|on|lzf|lz4|zlib. default: on ]"},
	{Text: "-directcopy", Description: "[ -directcopy enable directcopy. default: true ]"},
	{Text: "-useexistingreplica", Description: "[ -useexistingreplica use existing replica table if present. default: false ]"},
}

var table_upstream = []prompt.Suggest{
	{Text: "add", Description: "add"},
	{Text: "list", Description: "list"},
	{Text: "remove", Description: "remove"},
}
var table_upstream_add = []prompt.Suggest{
	{Text: "-path", Description: "-path table path"},
	{Text: "-upstream", Description: "-upstream upstream table path"},
}
var table_upstream_list = []prompt.Suggest{
	{Text: "-path", Description: "-path table path"},
}
var table_upstream_remove = []prompt.Suggest{
	{Text: "-path", Description: "-path table path"},
	{Text: "-upstream", Description: "-upstream upstream table path"},
}

var table_index = []prompt.Suggest{
	{Text: "add", Description: "add"},
	{Text: "remove", Description: "remove"},
	{Text: "list", Description: "list"},
}

var table_index_add = []prompt.Suggest{
	{Text: "-path", Description: "-path table path"},
	{Text: "-index", Description: "-index index name"},
	{Text: "-indexedfields", Description: "-indexedfields indexed field name, <fieldpath1: 1/-1 | ASC/DESC | asc/desc, $CAST(fieldpath2@INT)>"},
	{Text: "-includedfields", Description: "[ -includedfields included field name, <fieldpath3, fieldpath4> ]"},
	{Text: "-hashed", Description: "[ -hashed hashed index enabled. default: false ]"},
	{Text: "-numhashpartitions", Description: "[ -numhashpartitions number of hash index partitions when hashed index enabled ]"},
}
var table_index_remove = []prompt.Suggest{
	{Text: "-path", Description: "-path table path"},
	{Text: "-index", Description: "-index index name"},
}

var table_index_list = []prompt.Suggest{
	{Text: "-path", Description: "-path table path"},
	{Text: "-refreshnow", Description: "[ -refreshnow refreshnow. default: false ]"},
	{Text: "-indexname", Description: "[ -indexname index name ]"},
}

var table_changelog = []prompt.Suggest{
	{Text: "add", Description: "add"},
	{Text: "edit", Description: "edit"},
	{Text: "list", Description: "list"},
	{Text: "remove", Description: "remove"},
	{Text: "pause", Description: "pause"},
	{Text: "resume", Description: "resume"},
	{Text: "info", Description: "info"},
}

var table_changelog_add = []prompt.Suggest{
	{Text: "-path", Description: "-path table path"},
	{Text: "-changelog", Description: "-changelog remote changelog path"},
	{Text: "-columns", Description: "[ -columns comma separated list of <family>[:<column>] ]"},
	{Text: "-useexistingtopic", Description: "[ -useexistingtopic allow publishing to an existing topic. default: false ]"},
	{Text: "-propagateexistingdata", Description: "[ -propagateexistingdata publish existing data to the change stream otherwise only new changes will be propagated. default: true ]"},
	{Text: "-paused", Description: "[ -paused is replication paused. default: false ]"},
	{Text: "-throttle", Description: "[ -throttle throttle propagate operations under load. default: false ]"},
	{Text: "-synchronous", Description: "[ -synchronous propagate to remote changelog before acknowledging producers. default: false ]"},
	{Text: "-networkencryption", Description: "[ -networkencryption enable on-wire encryption. default: false ]"},
	{Text: "-networkcompression", Description: "[ -networkcompression on-wire compression type: off|lz4|lzf|zlib default: compression setting on changelog ]"},
}

//{Text:"", Description:"
var table_changelog_edit = []prompt.Suggest{
	{Text: "-path", Description: "-path changelog path"},
	{Text: "-changelog", Description: "-changelog remote changelog path"},
	{Text: "-throttle", Description: "[ -throttle throttle publish operations under load ]"},
	{Text: "-networkencryption", Description: "[ -networkencryption enable on-wire encryption ]"},
	{Text: "-synchronous", Description: "[ -synchronous publish to remote changelog before acknowledging producers ]"},
	{Text: "-networkcompression", Description: "[ -networkcompression on-wire compression type: off|lz4|lzf|zlib ]"},
}
var table_changelog_list = []prompt.Suggest{
	{Text: "-path", Description: "-path table path"},
	{Text: "-refreshnow", Description: "[ -refreshnow refreshnow. default: false ]"},
}
var table_changelog_remove = []prompt.Suggest{
	{Text: "-path", Description: "-path changelog path"},
	{Text: "-changelog", Description: "-changelog remote changelog path"},
}
var table_changelog_pause = []prompt.Suggest{
	{Text: "-path", Description: "-path changelog path"},
	{Text: "-changelog", Description: "-changelog remote changelog path"},
}
var table_changelog_resume = []prompt.Suggest{
	{Text: "-path", Description: "-path changelog path"},
	{Text: "-changelog", Description: "-changelog remote changelog path"},
}
var table_changelog_info = []prompt.Suggest{
	{Text: "-changelog", Description: "-changelog remote changelog path"},
}
