package options

import "testing"
import "github.com/spf13/pflag"

var svrOpt *ServerOption

func init() {
	svrOpt = NewServerOption()
}

func TestNewServerOption(t *testing.T) {
	svrOpt = NewServerOption()
	if svrOpt == nil {
		t.Error("NewServerOption failed.")
	}
}

func TestServerOption_AddFlags(t *testing.T) {
	svrOpt.AddFlags(pflag.CommandLine)
}
