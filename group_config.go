package dodin

import "regexp"

func NewGroupConfig(name, memberNamePattern string, variables map[string]string) GroupConfig {
	matcher, _ := regexp.Compile(memberNamePattern)
	return GroupConfig{
		name: name,
		memberNamePattern: memberNamePattern,
		matcher: matcher,
		variables: variables,
	}
}

type GroupConfig struct {
	name, memberNamePattern string
	matcher *regexp.Regexp
	variables map[string]string
}

func (config GroupConfig) Name() string {
	return config.name
}

func (config GroupConfig) MatchingMemberName(name string) bool {
	return config.matcher.MatchString(name)
}

func (config GroupConfig) Variables() map[string]string {
	return config.variables
}
