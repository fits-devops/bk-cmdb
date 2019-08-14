package main

import (
	"context"
	"fmt"
	"os"
	"runtime"

	"configcenter/src/common"
	"configcenter/src/common/blog"
	"configcenter/src/common/types"
	"configcenter/src/common/util"
	"configcenter/src/source_controller/testservice/app"
	"configcenter/src/source_controller/testservice/app/options"

	"github.com/spf13/pflag"
)

func main() {
	common.SetIdentification(types.CC_MODULE_TESTSERVICE)  //设置服务的全局ID
	runtime.GOMAXPROCS(runtime.NumCPU())

	blog.InitLogs()
	defer blog.CloseLogs()

	op := options.NewServerOption()
	op.AddFlags(pflag.CommandLine)

	util.InitFlags()

	if err := app.Run(context.Background(), op); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		blog.Errorf("process stoped by %v", err)
		blog.CloseLogs()
		os.Exit(1)
	}
}
