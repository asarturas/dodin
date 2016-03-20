package dodin

import (
	"gopkg.in/ini.v1"
	"strings"
)

func ParseConfig(configFileContents string) Config {
	config := Config{
		groups: make([]GroupConfig, 0),
	}
	iniConfig := ini.Empty()
	iniConfig.Append([]byte(configFileContents))
	for _, groupIni := range iniConfig.Sections() {
		parseGroup(&config, groupIni)
	}
	return config
}

func parseGroup(config *Config, groupIni *ini.Section) {
	if strings.HasSuffix(groupIni.Name(), ":vars") {
		parseVariables(config, groupIni)
	} else {
		parseHosts(config, groupIni)
	}
}

func parseVariables(config *Config, groupIni *ini.Section) {
	groupName := strings.Replace(groupIni.Name(), ":vars", "", 1)
	vars := map[string]string{}
	for _, key := range groupIni.Keys() {
		vars[key.Name()] = key.Value()
	}
	for i, group := range config.groups {
		if group.name == groupName {
			config.groups[i].variables = vars
		}
	}
}

func parseHosts(config *Config, groupIni *ini.Section) {
	if groupIni.HasKey("members") {
		memberNamePattern, _ := groupIni.GetKey("members")
		groupName := groupIni.Name()
		memberNamePatternValue := memberNamePattern.Value()
		config.groups = append(
			config.groups,
			NewGroupConfig(
				groupName,
				memberNamePatternValue,
				map[string]string{},
			),
		)
	}
}
