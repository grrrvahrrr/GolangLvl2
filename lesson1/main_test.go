package main

import "testing"

func BenchmarkCreateOneMillionFiles(b *testing.B) {
	createOneMillionFiles()
}
