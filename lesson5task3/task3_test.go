package lesson5task3

import "testing"

func BenchmarkWriteRead(b *testing.B) {
	writeRead(10, 100000)
}

func BenchmarkWriteReadRW(b *testing.B) {
	writeReadRW(10, 100000)
}
