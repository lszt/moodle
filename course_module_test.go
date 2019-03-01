package moodle

import (
	"fmt"
	"testing"
)

func TestCourseModule(t *testing.T) {

	api := NewMoodleApi(requireEnv("MOODLE_URL", t), requireEnv("MOODLE_KEY", t))

	s, err := api.GetCourses(" LIBRARY")
	if err != nil {
		t.Errorf("GetCourses() failed: %v", err)
	}
	if s == nil {
		t.Errorf("GetCourses() should return items")
	}
	if len(*s) < 1 {
		t.Errorf("Expecting search results")
	}

	r, _ := api.GetCourseRoles(36)
	if r == nil {
		t.Errorf("Course does not exist")
	}
	if len(*r) < 1 {
		t.Errorf("No roles in course")
	}

	cm, err := api.GetCourseModule(1155)
	if err != nil {
		t.Errorf("Module fetch failed: %v", err)
	}
	if cm == nil {
		t.Errorf("Module does not exist")
		//t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}

	fmt.Printf("Found availability %v\n", cm.Availability)
	//t.Errorf("%v", cm)
}
