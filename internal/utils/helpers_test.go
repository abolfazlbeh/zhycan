package utils

import "testing"

func TestArrayContains(t *testing.T) {
	// test string array
	arr := []string{"a", "b", "c"}

	mustBeFounded := "a"
	expectedValue := true
	actualValue := ArrayContains(&arr, mustBeFounded)
	if actualValue != expectedValue {
		t.Errorf("Find the %v in %v: Expected %v, but got %v",
			mustBeFounded, arr, expectedValue, actualValue)
	}
}

func TestArrayContainsInt(t *testing.T) {
	arr := []int{1, 2, 3}

	mustBeFounded := 2
	expectedValue := true
	actualValue := ArrayContains(&arr, mustBeFounded)
	if actualValue != expectedValue {
		t.Errorf("Find the %v in %v: Expected %v, but got %v",
			mustBeFounded, arr, expectedValue, actualValue)
	}

}

func TestArrayContainsWithPtr(t *testing.T) {
	arr := &[]int{1, 2, 3}

	mustBeFounded := 2
	expectedValue := true
	actualValue := ArrayContains(arr, mustBeFounded)
	if actualValue != expectedValue {
		t.Errorf("Find the %v in %v: Expected %v, but got %v",
			mustBeFounded, arr, expectedValue, actualValue)
	}

}

func TestArrayContainsDifferentTypes(t *testing.T) {
	TestArrayContains(t)
	TestArrayContainsInt(t)
	TestArrayContainsWithPtr(t)
}
