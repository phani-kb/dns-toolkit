package cmd

import (
	"github.com/phani-kb/dns-toolkit/internal/config"
	"github.com/phani-kb/multilog"
)

var Logger *multilog.Logger
var AppConfig *config.AppConfig
var SourcesConfigs []config.SourcesConfig
