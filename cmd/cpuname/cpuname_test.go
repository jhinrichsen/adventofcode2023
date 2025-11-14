package main

import "testing"

// BenchmarkDetectCPU is a minimal benchmark used solely for CPU detection.
// This ensures the cpuname tool can always detect CPU info without depending
// on any specific day's implementation.
func BenchmarkDetectCPU(b *testing.B) {
	for range b.N {
		_ = 42
	}
}
