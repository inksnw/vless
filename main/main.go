package main

import (
	"flag"
	"fmt"
	"vless/core"
	"vless/core/common/cmdarg"
)

var (
	configFiles cmdarg.Arg
	configDir   string
	version     = flag.Bool("version", false, "Show current version of V2Ray.")
	test        = flag.Bool("test", false, "Test config file only,without launching V2Ray server")
	format      = flag.String("format", "json", "Format of input file.")

	_ = func() error {
		flag.Var(&configFiles, "config", "Config file for V2Ray.Multiple assign is accepted(only json).Latter ones overrides the former ones.")
		flag.Var(&configFiles, "c", "Short alias of -config")
		flag.StringVar(&configDir, "confdir", "", "A dir with multiple json config")
		return nil
	}()
)

func printVersion() {
	version:=core.VersionStatement()
	for _,s:=range version{
		fmt.Println(s)
	}
}

func main() {
	flag.Parse()
	printVersion()
}
