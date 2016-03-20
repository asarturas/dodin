package dodin

import (
	"testing"
)

func Test_config_parser_parses_dodin_ini_file(t *testing.T) {
	iniFile := `[test]
members=master[12]
`;
	config := ParseConfig(iniFile)
	if len(config.Groups()) != 1 {
		t.Errorf("Expected to have %d groups, but got %d instead", 1, len(config.Groups()))
	}
	firstGroup := config.Groups()[0]
	if firstGroup.Name() != "test" {
		t.Error("Expected first group to be a name %s, but got %s instead", "test", firstGroup.Name())
	}
	if firstGroup.memberNamePattern != "master[12]" {
		t.Error("Expected first pattern to be %s, but got %s instead", "master[12]", firstGroup.memberNamePattern)
	}
}
