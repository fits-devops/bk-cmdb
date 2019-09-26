package main

import (
	"context"
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/pflag"

	"configcenter/src/common"
	"configcenter/src/common/blog"
	"configcenter/src/common/types"
	"configcenter/src/common/util"
	"configcenter/src/scene_server/admin_server/app"
	"configcenter/src/scene_server/admin_server/app/options"
	"configcenter/src/scene_server/admin_server/command"
)

func main() {
	common.SetIdentification(types.CC_MODULE_MIGRATE)

	runtime.GOMAXPROCS(runtime.NumCPU())

	blog.InitLogs()
	defer blog.CloseLogs()

	op := options.NewServerOption()
	op.AddFlags(pflag.CommandLine)

	if err := command.Parse(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	util.InitFlags()

	if err := app.Run(context.Background(), op); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		blog.Errorf("process stopped by %v", err)
		blog.CloseLogs()
		os.Exit(1)
	}
}
