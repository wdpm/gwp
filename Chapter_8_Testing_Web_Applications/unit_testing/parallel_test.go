package main

import (
	"testing"
	"time"
)

func TestParallel_1(t *testing.T) {
	t.Parallel()
	time.Sleep(1 * time.Second)
}

func TestParallel_2(t *testing.T) {
	t.Parallel()
	time.Sleep(2 * time.Second)
}

func TestParallel_3(t *testing.T) {
	t.Parallel()
	time.Sleep(3 * time.Second)
}

// === RUN   TestParallel_2
// === PAUSE TestParallel_2
// === RUN   TestParallel_3
// === PAUSE TestParallel_3
// === CONT  TestParallel_1
// === CONT  TestParallel_3
// === CONT  TestParallel_2
// --- PASS: TestParallel_1 (1.00s)
// --- PASS: TestParallel_2 (2.01s)
// --- PASS: TestParallel_3 (3.02s)
