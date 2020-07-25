package main

import (
	"fmt"
	"os"
	"runtime"
	"testing"

	"github.com/joho/godotenv"
)

func BenchmarkConvert_Serial(b *testing.B) {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	root := os.Getenv("MKWEBPROOT")
	exts := []string{".jpg", ".png"}
	fmt.Println(fmt.Sprintf("Serial: 1"))

	b.ResetTimer()
	convert(root, exts, 1)
}

func BenchmarkConvert_ParallelNumCPUDivTwo(b *testing.B) {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	root := os.Getenv("MKWEBPROOT")

	exts := []string{".jpg", ".png"}
	parallel_count := runtime.NumCPU() / 2
	fmt.Println(fmt.Sprintf("ParallelNumCPUDivTwo: %d", parallel_count))

	b.ResetTimer()
	convert(root, exts, parallel_count)
}

func BenchmarkConvert_ParallelNumCPUMinusOne(b *testing.B) {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	root := os.Getenv("MKWEBPROOT")

	exts := []string{".jpg", ".png"}
	parallel_count := runtime.NumCPU()*1 - 1
	fmt.Println(fmt.Sprintf("ParallelNumCPUMinusOne: %d", parallel_count))

	b.ResetTimer()
	convert(root, exts, parallel_count)
}

func BenchmarkConvert_ParallelNumCPU(b *testing.B) {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	root := os.Getenv("MKWEBPROOT")

	exts := []string{".jpg", ".png"}
	parallel_count := runtime.NumCPU() * 1
	fmt.Println(fmt.Sprintf("ParallelNumCPU: %d", parallel_count))

	b.ResetTimer()
	convert(root, exts, parallel_count)
}
func BenchmarkConvert_ParallelNumCPUTwice(b *testing.B) {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	root := os.Getenv("MKWEBPROOT")

	exts := []string{".jpg", ".png"}
	parallel_count := runtime.NumCPU() * 2
	fmt.Println(fmt.Sprintf("ParallelNumCPUTwice: %d", parallel_count))

	b.ResetTimer()
	convert(root, exts, parallel_count)
}

func BenchmarkConvert_ParallelNumCPUByFour(b *testing.B) {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	root := os.Getenv("MKWEBPROOT")

	exts := []string{".jpg", ".png"}
	parallel_count := runtime.NumCPU() * 4
	fmt.Println(fmt.Sprintf("ParallelNumCPUByFour: %d", parallel_count))

	b.ResetTimer()
	convert(root, exts, parallel_count)
}
