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

var stream_create = []prompt.Suggest{
	{Text: "-path", Description: "-path Stream Path"},
	{Text: "-ttl", Description: "[ -ttl Time to live in seconds. default:604800 ]"},
	{Text: "-autocreate", Description: "[ -autocreate Auto create topics. default:true ]"},
	{Text: "-defaultpartitions", Description: "[ -defaultpartitions Default partitions per topic. default:1 ]"},
	{Text: "-compression", Description: "[ -compression off|lz4|lzf|zlib. default:inherit from parent directory ]"},
	{Text: "-produceperm", Description: "[ -produceperm Producer access control expression. default u:creator ]"},
	{Text: "-consumeperm", Description: "[ -consumeperm Consumer access control expression. default u:creator ]"},
	{Text: "-topicperm", Description: "[ -topicperm Topic CRUD access control expression. default u:creator ]"},
	{Text: "-copyperm", Description: "[ -copyperm Stream copy access control expression. default u:creator ]"},
	{Text: "-adminperm", Description: "[ -adminperm Stream administration access control expression. default u:creator ]"},
	{Text: "-copymetafrom", Description: "[ -copymetafrom Stream to copy attributes from. default:none ]"},
	{Text: "-ischangelog", Description: "[ -ischangelog Stream to store changelog. default:false ]"},
	{Text: "-defaulttimestamptype", Description: "[ -defaulttimestamptype Timestamp type: createtime | logappendtime. default: createtime ]"},
	{Text: "-pidexpirysecs", Description: "[ -pidexpirysecs Producer ID expiry in secs. default: 604800 ]"},
}

var stream_edit = []prompt.Suggest{
	{Text: "-path", Description: "-path Stream Path"},
	{Text: "-ttl", Description: "[ -ttl Time to live in seconds ]"},
	{Text: "-autocreate", Description: "[ -autocreate Auto create topics ]"},
	{Text: "-defaultpartitions", Description: "[ -defaultpartitions Default partitions per topic ]"},
	{Text: "-compression", Description: "[ -compression off|lz4|lzf|zlib ]"},
	{Text: "-produceperm", Description: "[ -produceperm Producer access control expression. default u:creator ]"},
	{Text: "-consumeperm", Description: "[ -consumeperm Consumer access control expression. default u:creator ]"},
	{Text: "-topicperm", Description: "[ -topicperm Topic CRUD access control expression. default u:creator ]"},
	{Text: "-copyperm", Description: "[ -copyperm Stream copy access control expression. default u:creator ]"},
	{Text: "-adminperm", Description: "[ -adminperm Stream administration access control expression. default u:creator ]"},
	{Text: "-defaulttimestamptype", Description: "[ -defaulttimestamptype timestamp type: createtime | logappendtime. default: createtime ]"},
	{Text: "-compact", Description: "[ -compact Set log compaction for stream. default: false ]"},
	{Text: "-mincompactionlag", Description: "[ -mincompactionlag Set time in millisecond a message should remain uncompacted for. default: 0 ]"},
	{Text: "-deleteretention", Description: "[ -deleteretention Set the time in millisecond for which delete records are retained. default: 86400000 ]"},
	{Text: "-pidexpirysecs", Description: "[ -pidexpirysecs Producer ID expiry in secs. default: 604800 ]"},
	{Text: "-force", Description: "[ -force When used with -compact, forces enabling log compaction on a stream Parameter takes no value  ]"},
}

var stream_info = []prompt.Suggest{
	{Text: "-path", Description: "-path Stream Path"},
}

var stream_delete = []prompt.Suggest{
	{Text: "-path", Description: "-path Stream Path"},
}

var stream_purge = []prompt.Suggest{
	{Text: "-path", Description: "-path Stream Path"},
}

// {Text: "", Description: ""},
var stream_topic = []prompt.Suggest{
	{Text: "create", Description: "create"},
	{Text: "edit", Description: "edit"},
	{Text: "delete", Description: "delete"},
	{Text: "info", Description: "info"},
	{Text: "list", Description: "list"},
}

var stream_topic_create = []prompt.Suggest{
	{Text: "-path", Description: "-path Stream Path"},
	{Text: "-topic", Description: "-topic Topic Name"},
	{Text: "-partitions", Description: "[ -partitions Number of partitions. default: attribute defaultpartitions on the stream ]"},
	{Text: "-timestamptype", Description: "[ -timestamptype Timestamp type: createtime | logappendtime default: createtime ]"},
}
var stream_topic_edit = []prompt.Suggest{
	{Text: "-path", Description: "-path Stream Path"},
	{Text: "-topic", Description: "-topic Topic Name"},
	{Text: "-partitions", Description: " -partitions Number of partitions ]"},
	{Text: "-timestamptype", Description: "[ -timestamptype Timestamp type: createtime | logappendtime default: createtime ]"},
}
var stream_topic_delete = []prompt.Suggest{
	{Text: "-path", Description: "-path Stream Path"},
	{Text: "-topic", Description: "-topic Topic Name"},
}
var stream_topic_info = []prompt.Suggest{
	{Text: "-path", Description: "-path Stream Path"},
	{Text: "-topic", Description: "-topic Topic Name"},
}
var stream_topic_list = []prompt.Suggest{
	{Text: "-path", Description: "-path Stream Path"},
}

var stream_cursor = []prompt.Suggest{
	{Text: "delete", Description: "delete"},
	{Text: "list", Description: "list"},
}

var stream_cursor_delete = []prompt.Suggest{
	{Text: "-path", Description: "-path Stream Path"},
	{Text: "-consumergroup", Description: "[ -consumergroup Consumer Group ID ]"},
	{Text: "-topic", Description: "[ -topic Topic Name ]"},
	{Text: "-partition", Description: "[ -partition Partition ID ]"},
}

var stream_cursor_list = []prompt.Suggest{
	{Text: "-path", Description: "-path Stream Path"},
	{Text: "-consumergroup", Description: "[ -consumergroup Consumer Group ID ]"},
	{Text: "-topic", Description: "[ -topic Topic Name ]"},
	{Text: "-partition", Description: "[ -partition Partition ID ]"},
}

var stream_assign = []prompt.Suggest{
	{Text: "-path", Description: "-path Stream Path"},
	{Text: "-consumergroup", Description: "[ -consumergroup Consumer Group ID ]"},
	{Text: "-topic", Description: "[ -topic Topic Name ]"},
	{Text: "-partition", Description: "[ -partition Partition ID ]"},
	{Text: "-detail", Description: "[ -detail Detail Parameter takes no value  ]"},
}

var stream_replica = []prompt.Suggest{
	{Text: "add", Description: "add"},
	{Text: "edit", Description: "edit"},
	{Text: "list", Description: "list"},
	{Text: "remove", Description: "remove"},
	{Text: "pause", Description: "pause"},
	{Text: "resume", Description: "resume"},
	{Text: "autosetup", Description: "autosetup"},
}

var stream_replica_add = []prompt.Suggest{
	{Text: "-path", Description: "-path stream path"},
	{Text: "-replica", Description: "-replica remote stream path"},
	{Text: "-paused", Description: "[ -paused start replication in paused state. default: false ]"},
	{Text: "-throttle", Description: "[ -throttle throttle replication operations under load. default: false ]"},
	{Text: "-networkencryption", Description: "[ -networkencryption enable on-wire encryption. default: false ]"},
	{Text: "-synchronous", Description: "[ -synchronous replicate to remote stream before acknowledging producers. default: false "},
	{Text: "-networkcompression", Description: "[ -networkcompression on-wire compression type: off|lz4|lzf|zlib default: compression setting on stream "},
}
var stream_replica_edit = []prompt.Suggest{
	{Text: "-path", Description: "-path stream path"},
	{Text: "-replica", Description: "-replica remote stream path"},
	{Text: "-newreplica", Description: "[ -newreplica renamed stream path ]"},
	{Text: "-throttle", Description: "[ -throttle throttle replication operations under load ]"},
	{Text: "-networkencryption", Description: "[ -networkencryption enable on-wire encryption ]"},
	{Text: "-synchronous", Description: "[ -synchronous replicate to remote stream before acknowledging producers ]"},
	{Text: "-networkcompression", Description: "[ -networkcompression on-wire compression type: off|lz4|lzf|zlib ]"},
}
var stream_replica_list = []prompt.Suggest{
	{Text: "-path", Description: "-path stream path"},
	{Text: "-refreshnow", Description: "[ -refreshnow refreshnow. default: false ]"},
}
var stream_replica_remove = []prompt.Suggest{
	{Text: "-path", Description: "-path stream path"},
	{Text: "-replica", Description: "-replica remote stream path"},
}
var stream_replica_pause = []prompt.Suggest{
	{Text: "-path", Description: "-path stream path"},
	{Text: "-replica", Description: "-replica replica stream path"},
}
var stream_replica_resume = []prompt.Suggest{
	{Text: "-path", Description: "-path stream path"},
	{Text: "-replica", Description: "-replica stream table path"},
}
var stream_replica_autosetup = []prompt.Suggest{
	{Text: "-path", Description: "-path stream path"},
	{Text: "-replica", Description: " -replica remote stream path"},
	{Text: "-synchronous", Description: "[ -synchronous replicate to remote stream before acknowledging producers. default: false ]"},
	{Text: "-multimaster", Description: "[ -multimaster set up bi-directional replication. default: false ]"},
	{Text: "-throttle", Description: "[ -throttle throttle replication operations under load. default: false ]"},
	{Text: "-networkencryption", Description: "[ -networkencryption enable on-wire encryption. default: false ]"},
	{Text: "-networkcompression", Description: "[ -networkcompression on-wire compression type: off|lz4|lzf|zlib default: compression setting on stream ]"},
	{Text: "-directcopy", Description: "[ -directcopy enable directcopy. default: true ]"},
	{Text: "-useexistingreplica", Description: "[ -useexistingreplica use existing replica table if present. default: false ]"},
}
var stream_upstream = []prompt.Suggest{
	{Text: "add", Description: "add"},
	{Text: "list", Description: "list"},
	{Text: "remove", Description: "remove"},
}
var stream_upstream_add = []prompt.Suggest{
	{Text: "-path", Description: "-path stream path"},
	{Text: "-upstream", Description: "-upstream upstream stream path"},
}
var stream_upstream_list = []prompt.Suggest{
	{Text: "-path", Description: "-path stream path"},
}
var stream_upstream_remove = []prompt.Suggest{
	{Text: "-path", Description: "-path stream path"},
	{Text: "-upstream", Description: "-upstream upstream stream path"},
}
var stream_listrecent = []prompt.Suggest{}
