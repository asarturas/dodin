package dodin

import (
	"gopkg.in/ini.v1"
)

func ParseConfig(configFileContents string) Config {
	config := Config{
		groups: make([]GroupConfig, 0),
	}
	iniConfig := ini.Empty()
	iniConfig.Append([]byte(configFileContents))
	for _, groupIni := range iniConfig.Sections() {
		if groupIni.HasKey("members") {
			memberNamePattern, _ := groupIni.GetKey("members")
			groupName := groupIni.Name()
			memberNamePatternValue := memberNamePattern.Value()
			config.groups = append(config.groups, NewGroupConfig(groupName, memberNamePatternValue))
		}
	}
	return config
}
