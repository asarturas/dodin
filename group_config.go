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

func (this GroupConfig) Name() string {
	return this.name
}

func (this GroupConfig) MatchingMemberName(name string) bool {
	return this.matcher.MatchString(name)
}

func (this GroupConfig) Variables() map[string]string {
	return this.variables
}
