package day16

import "testing"

var testcases = []struct {
	size     int
	orders   []interface{}
	getLast  []int
	expected []interface{}
}{
	{5, []interface{}{5, 4, "abc", 10, "def", "ged"},
		[]int{1, 2, 3, 4, 5},
		[]interface{}{"ged", "def", 10, "abc", 4}},
	{5, []interface{}{"abc"},
		[]int{1, 2, 3, 4, 5},
		[]interface{}{"abc", nil, nil, nil, nil}},
}

func TestOrderLogSlice(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		r := NewOrderLogSlice(tc.size)
		for _, v := range tc.orders {
			r.Record(v)
		}
		for i, v := range tc.getLast {
			if result := r.GetLast(v); result != tc.expected[i] {
				t.Errorf("Expected %v but got %v", tc.expected[i], result)
			}
		}
	}
}

func TestOrderLogRing(t *testing.T) {
	t.Parallel()
	for _, tc := range testcases {
		r := NewOrderLogRing(tc.size)
		for _, v := range tc.orders {
			r.Record(v)
		}
		for i, v := range tc.getLast {
			if result := r.GetLast(v); result != tc.expected[i] {
				t.Errorf("Expected %v but got %v", tc.expected[i], result)
			}
		}
	}
}

func BenchmarkOrderLogSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			r := NewOrderLogSlice(tc.size)
			for _, v := range tc.orders {
				r.Record(v)
			}
			for _, v := range tc.getLast {
				r.GetLast(v)
			}
		}
	}
}

func BenchmarkOrderLogRing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testcases {
			r := NewOrderLogRing(tc.size)
			for _, v := range tc.orders {
				r.Record(v)
			}
			for _, v := range tc.getLast {
				r.GetLast(v)
			}
		}
	}
}
