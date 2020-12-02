package core

import (
	"runtime"
	"vless/core/common/serial"
)

var (
	version  = "4.33.0"
	build    = "Custom"
	codename = "V2Fly, a community-driven edition of V2Ray."
	intro    = "A unified platform for anti-censorship."
)

func Version() string {
	return version

}

func VersionStatement() []string {
	return []string{
		serial.Contact("V2Ray", Version(), " (", codename, ") ", build, " (", runtime.Version(), " ", runtime.GOOS, "/", runtime.GOARCH, ")"),
		intro,
	}

}
