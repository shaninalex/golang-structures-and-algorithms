package algs

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		target   int
		expected int
	}{
		{"found", []int{1, 3, 5, 7, 9, 11, 13}, 7, 3},
		{"not found", []int{1, 3, 5, 7, 9, 11, 13}, 6, -1},
		{"empty array", []int{}, 5, -1},
		{"single element", []int{5}, 5, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := BinarySearch(test.arr, test.target)
			if actual != test.expected {
				t.Errorf("Expected index %d but got %d", test.expected, actual)
			}
		})
	}
}
