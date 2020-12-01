package main

import (
	"reflect"
	"testing"
)

func TestGradingStudents(t *testing.T) {

	t.Run("Base case", func(t *testing.T) {
		grades := []int32{75, 32, 43, 59, 35, 95, 43, 68}

		got := gradingStudents(grades)
		want := []int32{75, 32, 45, 60, 35, 95, 45, 70}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d given", got, want)
		}
	})
}
