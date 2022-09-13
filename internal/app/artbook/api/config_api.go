package api

import (
	"github.com/maldan/gam-app-artbook/internal/app/artbook/core"
	"github.com/maldan/go-cmhp/cmhp_file"
)

type ConfigApi struct {
}

// Save config
func (r ConfigApi) PostIndex(args core.Config) {
	cmhp_file.Write(core.DataDir+"/config.json", &args)
}

// Get config
func (r ConfigApi) GetIndex() core.Config {
	config := core.Config{}
	cmhp_file.ReadJSON(core.DataDir+"/config.json", &config)
	return config
}
