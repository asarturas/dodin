package dodin

import "regexp"

func NewGroupConfig(name, memberNamePattern string) GroupConfig {
	matcher, _ := regexp.Compile(memberNamePattern)
	return GroupConfig{
		name: name,
		memberNamePattern: memberNamePattern,
		matcher: matcher,
	}
}

type GroupConfig struct {
	name, memberNamePattern string
	matcher *regexp.Regexp
}

func (this GroupConfig) Name() string {
	return this.name
}

func (this GroupConfig) MatchingMemberName(name string) bool {
	return this.matcher.MatchString(name)
}
