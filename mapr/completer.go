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

import (
	"fmt"
	"strings"

	"github.com/c-bata/go-prompt"
)

func NewCompleter() (*Completer, error) {
	return &Completer{}, nil
}

type Completer struct {
}

func (c *Completer) Complete(d prompt.Document) []prompt.Suggest {
	if d.TextBeforeCursor() == "" {
		return []prompt.Suggest{}
	}
	args := strings.Split(d.TextBeforeCursor(), " ")
	for i := range args {
		if args[i] == "|" {
			return []prompt.Suggest{}
		}
	}
	if len(args) <= 1 {
		return prompt.FilterHasPrefix(commands, args[0], true)
	}

	return argumentCompleter(args)
}

func remainingOptions(args []string, options []prompt.Suggest) []prompt.Suggest {

	if len(args) > 2 {
		currLen := len(args)
		optionList := make([]prompt.Suggest, len(options))
		copy(optionList, options)
		for _, o := range args {
			for i := 0; i < len(optionList); i++ {
				if optionList[i].Text == o {
					newList := make([]prompt.Suggest, len(optionList)-1)
					newList = append(optionList[:i], optionList[i+1:]...)
					optionList = newList
				}
			}
		}

		if len(optionList) > 0 {
			return prompt.FilterHasPrefix(optionList, args[currLen-1], true)
		}
	}
	return []prompt.Suggest{}
}

var selection []prompt.Suggest

func argumentCompleter(args []string) []prompt.Suggest {
	if len(args) <= 1 {
		return prompt.FilterHasPrefix(commands, args[0], true)
	}

	first := args[0]
	switch first {
	case "acerole":
		second := args[1]
		if len(args) == 2 {
			return prompt.FilterHasPrefix(acerole_validate, second, true)
		}

	case "acl":
		second := args[1]
		if len(args) == 2 {
			return prompt.FilterHasPrefix(acl, second, true)
		}
		third := args[2]
		//var selection []prompt.Suggest
		if len(args) == 3 {
			switch second {
			case "show":
				selection = acl_show
			case "userperms":
				selection = acl_userperms
			case "set":
				selection = acl_set

			case "edit":
				selection = acl_edit
			default:
				selection = []prompt.Suggest{}
			}
			return prompt.FilterHasPrefix(selection, third, true)
		}

		return remainingOptions(args, selection)

	case "alarm":
		second := args[1]
		if len(args) == 2 {
			return prompt.FilterHasPrefix(alarm, second, true)
		}
		third := args[2]
		if len(args) == 3 {
			switch second {
			case "list":
				selection = alarm_list
			case "raise":
				selection = alarm_raise
			case "clear":
				selection = alarm_clear
			case "clearmulti":
				selection = alarm_clearmulti
			case "mute":
				selection = alarm_mute
			case "unmute":
				selection = alarm_unmute
			case "clearall":
				selection = alarm_clearall
			case "config":
				selection = alarm_config
			case "names":
				selection = alarm_names
			case "groups":
				selection = alarm_groups
			default:
				selection = []prompt.Suggest{}
			}
			return prompt.FilterHasPrefix(selection, third, true)
		}

		switch third {
		case "load":
			selection = alarm_config_load
		case "save":
			selection = alarm_config_save
		case "listGroup":
			selection = alarm_groups_listGroup
		case "addEmails":
			selection = alarm_groups_addEmails
		case "deleteEmails":
			selection = alarm_groups_deleteEmails
		case "addAlarms":
			selection = alarm_groups_addAlarms
		case "deleteAlarms":
			selection = alarm_groups_deleteAlarms
		}
		return remainingOptions(args, selection)

	case "dashboard":
		second := args[1]
		if len(args) == 2 {
			return prompt.FilterHasPrefix(dashboard, second, true)
		}

		return remainingOptions(args, dashboard_info)

	case "disk":
		second := args[1]
		if len(args) == 2 {
			return prompt.FilterHasPrefix(disk, second, true)
		}

		third := args[2]
		if len(args) == 3 {
			switch second {
			case "listall":
				selection = disk_listall
			case "list":
				selection = disk_list
			case "add":
				selection = disk_add
			case "remove":
				selection = disk_remove
			default:
				selection = []prompt.Suggest{}
			}
			return prompt.FilterHasPrefix(selection, third, true)
		}
		return remainingOptions(args, selection)

	case "entity":
		second := args[1]
		if len(args) == 2 {
			return prompt.FilterHasPrefix(entity, second, true)
		}
		third := args[2]
		if len(args) == 3 {
			switch second {
			case "list":
				selection = entity_list
			case "info":
				selection = entity_info
			case "modify":
				selection = entity_modify
			case "remove":
				selection = entity_remove
			default:
				selection = []prompt.Suggest{}
			}
			return prompt.FilterHasPrefix(selection, third, true)
		}
		return remainingOptions(args, selection)

	case "stream":
		second := args[1]
		if len(args) == 2 {
			return prompt.FilterHasPrefix(stream, second, true)
		}
		third := args[2]
		if len(args) == 3 {
			switch second {
			case "create":
				selection = stream_create
			case "edit":
				selection = stream_edit
			case "delete":
				selection = stream_delete
			case "info":
				selection = stream_info
			case "purge":
				selection = stream_purge
			case "topic":
				selection = stream_topic
			case "cursor":
				selection = stream_cursor
			case "assign":
				selection = stream_assign
			case "replica":
				selection = stream_replica
			case "upstream":
				selection = stream_upstream
			case "listrecent":
				selection = stream_listrecent
			}
			return prompt.FilterHasPrefix(selection, third, true)
		}

		fourth := args[3]
		if len(args) == 4 {
			switch third {
			case "add":
				switch second {
				case "replica":
					selection = stream_replica_add
				case "upstream":
					selection = stream_upstream_add
				}
			case "edit":
				switch second {
				case "topic":
					selection = stream_topic_edit
				case "replica":
					selection = stream_replica_edit
				}
			case "create":
				selection = stream_topic_create
			case "delete":
				switch second {
				case "topic":
					selection = stream_topic_delete
				case "cursor":
					selection = stream_topic_delete
				}
			case "info":
				selection = stream_topic_info

			case "list":
				switch second {
				case "topic":
					selection = stream_topic_list
				case "replica":
					selection = stream_replica_list
				case "upstream":
					selection = stream_upstream_list
				case "cursor":
					selection = stream_cursor_list
				}

			case "pause":
				switch second {
				case "replica":
					selection = stream_replica_pause
				}

			case "remove":
				switch second {
				case "replica":
					selection = stream_replica_remove
				case "upstream":
					selection = stream_upstream_remove
				}
			case "resume":
				switch second {
				case "replica":
					selection = stream_replica_resume
				}

			default:
				return remainingOptions(args, selection)
			}
			return prompt.FilterHasPrefix(selection, fourth, true)
		}

		return remainingOptions(args, selection)

	case "table":
		second := args[1]
		if len(args) == 2 {
			return prompt.FilterHasPrefix(table, second, true)
		}
		third := args[2]
		if len(args) == 3 {
			switch second {
			case "create":
				selection = table_create
			case "edit":
				selection = table_edit
			case "delete":
				selection = table_delete
			case "info":
				selection = table_info
			case "cf":
				selection = table_cf
			case "region":
				selection = table_region
			case "replica":
				selection = table_replica
			case "upstream":
				selection = table_upstream
			case "index":
				selection = table_index
			case "changelog":
				selection = table_changelog
			}
			return prompt.FilterHasPrefix(selection, third, true)
		}

		fourth := args[3]
		if len(args) == 4 {
			switch third {
			case "add":
				switch second {
				case "index":
					selection = table_index_add
				case "replica":
					selection = table_replica_add
				case "upstream":
					selection = table_upstream_add
				case "changelog":
					selection = table_changelog_add
				}
			case "edit":
				switch second {
				case "cf":
					selection = table_cf_edit
				case "replica":
					selection = table_replica_edit
				case "changelog":
					selection = table_changelog_edit
				}
			case "create":
				selection = table_cf_create
			case "delete":
				selection = table_cf_delete
			case "info":
				selection = table_changelog_info

			case "list":
				switch second {
				case "cf":
					selection = table_cf_list
				case "replica":
					selection = table_replica_list
				case "upstream":
					selection = table_upstream_list
				case "region":
					selection = table_region_list
				case "index":
					selection = table_index_list
				case "changelog":
					selection = table_changelog_list
				}
			case "colperm":
				selection = table_cf_colperm
			case "split":
				selection = table_region_split
			case "pack":
				selection = table_region_pack
			case "pause":
				switch second {
				case "changelog":
					selection = table_changelog_pause
				case "replica":
					selection = table_replica_pause
				}

			case "merge":
				selection = table_region_merge
			case "remove":
				switch second {
				case "replica":
					selection = table_replica_remove
				case "upstream":
					selection = table_upstream_remove
				case "index":
					selection = table_index_remove
				case "changelog":
					selection = table_changelog_remove
				}
			case "resume":
				switch second {
				case "replica":
					selection = table_replica_resume
				case "changelog":
					selection = table_changelog_resume
				}

			case "stat":
				selection = table_region_stat
			default:
				return remainingOptions(args, selection)
			}
			return prompt.FilterHasPrefix(selection, fourth, true)
		}
		fifth := args[4]
		if len(args) == 5 {
			switch fourth {
			case "get":
				selection = table_cf_colperm_get
			case "set":
				selection = table_cf_colperm_set
			case "delete":
				selection = table_cf_colperm_delete
			default:
				return remainingOptions(args, selection)
			}
			return prompt.FilterHasPrefix(selection, fifth, true)
		}
		return remainingOptions(args, selection)

	case "volume":
		second := args[1]
		if len(args) == 2 {
			return prompt.FilterHasPrefix(volume, second, true)
		}

		third := args[2]
		if len(args) == 3 {
			switch second {
			case "audit":
				selection = volume_audit
			case "create":
				selection = volume_create
			case "unmount":
				selection = volume_unmount
			case "mount":
				selection = volume_mount
			case "showmounts":
				selection = volume_showmounts
			case "offload":
				selection = volume_offload
			case "recall":
				selection = volume_recall
			case "compact":
				selection = volume_compact
			case "tierjobabort":
				selection = volume_tierjobabort
			case "tierjobstatus":
				selection = volume_tierjobstatus
			case "tierstats":
				selection = volume_tierstats
			case "remove":
				selection = volume_remove
			case "move":
				selection = volume_move
			case "rename":
				selection = volume_rename
			case "upgradeformat":
				selection = volume_upgradeformat
			case "modify":
				selection = volume_modify
			case "info":
				selection = volume_info
			case "list":
				selection = volume_list
			case "fixmountpath":
				selection = volume_fixmountpath
			case "link":
				selection = volume_link
			case "snapshot":
				selection = volume_snapshot
			case "mirror":
				selection = volume_mirror
			case "dump":
				selection = volume_dump
			case "container":
				selection = volume_container
			case "balancecontainers":
				selection = volume_balancecontainers
			case "balancinginfo":
				selection = volume_balancinginfo
			}
			return prompt.FilterHasPrefix(selection, third, true)
		}

		fourth := args[3]
		if len(args) == 4 {
			switch third {
			case "create":
				switch second {
				case "dump":
					selection = volume_dump_create
				case "snapshot":
					selection = volume_snapshot_create
				}

			case "remove":
				selection = volume_snapshot_remove
			case "restore":
				selection = volume_dump_restore
			case "preserve":
				selection = volume_snapshot_preserve
			case "list":
				selection = volume_snapshot_list
			case "start":
				selection = volume_mirror_start
			case "stop":
				selection = volume_mirror_stop
			case "push":
				selection = volume_mirror_push
			case "move":
				selection = volume_container_move
			case "switchmaster":
				selection = volume_container_switchmaster
			case "switchtailwithupstream":
				selection = volume_container_switchtailwithupstream
			}
			return prompt.FilterHasPrefix(selection, fourth, true)
		}

		return remainingOptions(args, selection)

	default:
		return []prompt.Suggest{}

	}
	return []prompt.Suggest{}
}

func Executor(in string) {

	fmt.Println("Command generated:")
	fmt.Println("maprcli", in)

}
