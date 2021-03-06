package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"vless/core"
	"vless/core/common/cmdarg"
	"vless/core/platform"
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
	version := core.VersionStatement()
	for _, s := range version {
		fmt.Println(s)
	}
}

func main() {
	flag.Parse()
	printVersion()

	if *version {
		return
	}
	server, err := startV2Ray()

	if err != nil {
		fmt.Println(err)
		os.Exit(23)
	}
	if *test {
		fmt.Println("Configguration OK.")
		os.Exit(0)
	}

	if err := server.Start(); err != nil {
		fmt.Println("Failed to start", err)
		os.Exit(-1)
	}
	defer server.Close()
	runtime.GC()

	{
		osSignals := make(chan os.Signal, 1)
		signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)
		<-osSignals
	}
}

func GetConfigFormat() string {
	switch strings.ToLower(*format) {
	case "pb", "protobuf":
		return "protobuf"
	default:
		return "json"
	}

}

func startV2Ray() (core.Server, error) {
	configFiles := getConfigFilePath()

	config, err := core.LoadConfig(GetConfigFormat(), configFiles[0], configFiles)
	if err != nil {
		return nil, newError("failed to read config files: [", configFiles.String(), "]").Base(err)
	}

	server, err := core.New(config)
	if err != nil {
		return nil, newError("failed to create server").Base(err)
	}
	return server, nil

	return ser

}

func fileExists(file string) bool {
	info, err := os.Stat(file)
	return err == nil && !info.IsDir()

}

func getConfigFilePath() cmdarg.Arg {
	if dirExists(configDir) {
		log.Println("Using confdir from arg:", configDir)
		readConfDir(configDir)
	} else if envConfDir := platform.GetConfDirPath(); dirExists(envConfDir) {
		log.Println("Using confdir from env:", envConfDir)
		readConfDir(envConfDir)
	}
	if len(configFiles) > 0 {
		return configFiles
	}
	if workingDir, err := os.Getwd(); err == nil {
		configFile := filepath.Join(workingDir, "config.json")
		if fileExists(configFile) {
			log.Println("Using default config: ", configFile)
			return cmdarg.Arg{configFile}
		}

	}

	if configFiles := platform.GetConfigurationPath(); fileExists(configFiles) {
		log.Println("Using config from env: ", configFiles)
		return cmdarg.Arg{configFiles}
	}
	log.Println("Using config form STDIN")
	return cmdarg.Arg{"stdin:"}

}

func readConfDir(dirPath string) {
	confs, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatalln(err)
	}
	for _, f := range confs {
		if strings.HasPrefix(f.Name(), ".json") {
			_ = configFiles.Set(path.Join(dirPath, f.Name()))
		}
	}

}

func dirExists(file string) bool {
	if file == "" {
		return false
	}
	info, err := os.Stat(file)
	return err == nil && info.IsDir()
}
