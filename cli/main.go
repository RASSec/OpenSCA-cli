/*
 * @Descripation: 引擎入口
 * @Date: 2021-11-03 11:12:19
 */
package main

import (
	"analyzer/engine"
	"flag"
	"fmt"
	"os"
	"util/args"
	"util/logs"
	"util/model"
)

func main() {
	args.Parse()
	if len(args.Filepath) > 0 {
		output(engine.NewEngine().ParseFile(args.Filepath))
	} else {
		flag.PrintDefaults()
	}
}

// output 输出结果
func output(depRoot *model.DepTree, err error) {
	// 整理错误信息
	errInfo := ""
	if err != nil {
		errInfo = err.Error()
	}
	// 记录依赖
	logs.Debug("\n" + depRoot.String())
	// 输出结果
	if args.Out != "" {
		// 保存到json
		if f, err := os.Create(args.Out); err != nil {
			logs.Error(err)
		} else {
			defer f.Close()
			if size, err := f.Write(depRoot.Json(errInfo)); err != nil {
				logs.Error(err)
			} else {
				logs.Info(fmt.Sprintf("size: %d, output: %s", size, args.Out))
			}
		}
	} else {
		fmt.Println(string(depRoot.Json(errInfo)))
	}
}
