package day132

import "testing"

func TestHitCounter(t *testing.T) {
	t.Parallel()
	hc := NewHitCounter()
	for i := uint64(0); i < uint64(1000); i++ {
		hc.Record(i)
	}
	if total := hc.Total(); total != 1000 {
		t.Errorf("Expected 1000 got %v", total)
	}
	if size := hc.Range(500, 10000); size != 500 {
		t.Errorf("Expected 500 got %v", size)
	}
}

func BenchmarkHitCounter(b *testing.B) {
	hc := NewHitCounter()
	for i := uint64(0); i < uint64(100000); i++ {
		hc.Record(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		hc.Total()
		hc.Range(631, 632)
	}
}
