/*
 Copyright 2020 Qiniu Cloud (qiniu.com)

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
	"strings"

	"github.com/qiniu/goc/cmd"
	_ "github.com/qiniu/goc/statik"
)

func main() {
	// 解析go flag
	// 判断build模式
	if len(os.Args) > 1 && os.Args[1] == "build" {
		var args []string
		var args1 []string
		args = append(args, os.Args[:2]...)

		var outIndex = -1
		for i, arg := range os.Args[2:] {
			if strings.Contains(arg, " ") {
				arg = fmt.Sprintf("\"%s\"", arg)
			}

			if arg == "-o" {
				outIndex = i + 1
				args = append(args, arg)
				continue
			}

			if outIndex == i {
				args = append(args, arg)
				continue
			}

			args1 = append(args1, arg)
		}

		args = append(args, "--buildflags")
		args = append(args, fmt.Sprintf(`%s`, strings.Join(args1[:len(args1)-1], " ")))
		if args1 != nil && len(args1) > 0 {
			args = append(args, args1[len(args1)-1])
		}
		os.Args = args
	}

	cmd.Execute()
}
