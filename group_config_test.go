package dodin

import (
	"testing"
)

func Test_group_config_matches_machine_name_with_member_name_pattern(t *testing.T) {
	config := NewGroupConfig("test", "member[12]")
	if config.MatchingMemberName("member1") == false {
		t.Errorf("Expected %s to be matching %s", "member1", config.memberNamePattern)
	}
	if config.MatchingMemberName("member3") == true {
		t.Errorf("Expected %s not to be matching %s", "member3", config.memberNamePattern)
	}
}
