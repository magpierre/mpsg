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

package main

import (
	"fmt"
	"os"

	prompt "github.com/c-bata/go-prompt"
	"github.com/magpierre/mpsg/mapr"
)

func main() {
	version := "0.3"
	fmt.Printf("Magnus Pierre's Statement Generator for MapRCLI (version: %s)\n", version)
	fmt.Println("Please use `exit` or `Ctrl-D` to exit this program.")
	c, err := mapr.NewCompleter()
	if err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}
	p := prompt.New(mapr.Executor, c.Complete, prompt.OptionTitle("Statement generator"),
		prompt.OptionPrefix(">>> "),
		prompt.OptionInputTextColor(prompt.DarkGreen),
	)

	p.Run()

}
