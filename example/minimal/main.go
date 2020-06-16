package main

import (
	"github.com/bbiao/nephele/interpret"
	"github.com/bbiao/nephele/process"
	"github.com/bbiao/nephele/storage"

	"github.com/bbiao/nephele/server"
	"github.com/bbiao/nephele/util"

	_ "github.com/bbiao/nephele/command/autoorient"
	_ "github.com/bbiao/nephele/command/crop"
	_ "github.com/bbiao/nephele/command/format"
	_ "github.com/bbiao/nephele/command/quality"
	_ "github.com/bbiao/nephele/command/resize"
	_ "github.com/bbiao/nephele/command/rotate"
	_ "github.com/bbiao/nephele/command/sharpen"
	_ "github.com/bbiao/nephele/command/strip"
	_ "github.com/bbiao/nephele/command/watermark"

	_ "github.com/bbiao/nephele/interpret/neph"
	_ "github.com/bbiao/nephele/storage/neph"
	_ "github.com/bbiao/nephele/server/ping"
)

var Config = struct {
	ServerConfigPath string `toml:"server_config_path"`
	Interpret        *map[string]string
	Process          *map[string]string
	Storage          *map[string]string
}{
	Interpret: &interpret.Config,
	Process:   &process.Config,
	Storage:   &storage.Config,
}

func main() {
	util.FromToml("default.toml", &Config)
	util.FromToml(Config.ServerConfigPath, &server.Config)

	interpret.Init()
	process.Init()
	storage.Init()

	server.Run()
}
