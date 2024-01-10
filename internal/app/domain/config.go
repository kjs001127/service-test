package domain

import (
	"context"
)

type ConfigSvc struct {
}

type ConfigValue struct {
	Key   string
	Value string
}

func (s *ConfigSvc) ValidateConfigs(ctx context.Context, install InstallInfo, input []*ConfigValue) error {
	panic("")
}

func (s *ConfigSvc) DefaultConfigOf(ctx context.Context, install InstallInfo) []*ConfigValue {
	panic("")
}
