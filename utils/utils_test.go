package utils

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
    testCases := []struct {
        x, y, expected int
    }{
        {2, 3, 5},
        {-2, -2, -4},
        {2, 2, 4},
        {2, 0, 2},
        {0, 0, 0},
        {2, -2, 0},
    }

    for _, tc := range testCases {
        t.Run(fmt.Sprintf("Add(%d, %d)", tc.x, tc.y), func(t *testing.T)) {
            result := Add(tc.x, tc.y)
            if result != tc.expected {
                t.Errorf("Add(%d, %d) = %d; want %d", tc.x, tc.y, result, tc.expected)
            }
        }
    }
}
