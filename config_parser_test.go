package dodin

import (
	"testing"
)

func Test_config_parser_parses_dodin_ini_file(t *testing.T) {
	iniFile := `[test]
members=master[12]
[test:vars]
variable_name="variable_value"
`
	config := ParseConfig(iniFile)
	if len(config.Groups()) != 1 {
		t.Errorf("Expected to have %d groups, but got %d instead", 1, len(config.Groups()))
	}
	firstGroup := config.Groups()[0]
	if firstGroup.Name() != "test" {
		t.Errorf("Expected first group to be a name 'test', but got %s instead", firstGroup.Name())
	}
	if firstGroup.memberNamePattern != "master[12]" {
		t.Errorf("Expected first pattern to be 'master[12], but got %s instead", firstGroup.memberNamePattern)
	}
	if firstGroup.variables["variable_name"] != "variable_value" {
		t.Error("Unexpected group variables", firstGroup.variables)
	}
}
