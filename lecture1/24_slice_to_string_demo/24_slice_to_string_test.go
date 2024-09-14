package _4_slice_to_string_demo

import (
	"github.com/stretchr/testify/require"
	"math/rand/v2"
	"strings"
	"testing"
	"unsafe"
)

// go test -bench=. -benchmem -cpuprofile=cpu.prof -memprofile=mem.prof
func BenchmarkMatch(b *testing.B) {
	b.Run("fast match", func(b *testing.B) {
		b.ReportAllocs()

		b.StopTimer()
		data := getData()
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			fastMatch(data)
		}
	})

	b.Run("slow match", func(b *testing.B) {
		b.ReportAllocs()

		b.StopTimer()
		data := getData()
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			slowMatch(data)
		}
	})
}

func getData() []string {
	data := make([]string, 0, 1<<21)

	for i := 0; i < cap(data); i++ {
		// see
		// https://github.com/golang/go/blob/master/src/runtime/string.go#L176
		data = append(data, strings.Repeat("1", 33<<6+1<<8+rand.N(2)))
	}

	return data
}

func TestMatch(t *testing.T) {
	data := getData()
	require.Equal(t, fastMatch(data), slowMatch(data))
}

func fastMatch(exampleData []string) [][]byte {
	result := make([][]byte, 0, len(exampleData))

	for i := 0; i < len(exampleData); i++ {
		data := unsafe.Slice(unsafe.StringData(exampleData[i]), len(exampleData[i]))

		if externalMatchFunction(data) {
			result = append(result, data)
		}
	}

	return result
}

func slowMatch(exampleData []string) [][]byte {
	result := make([][]byte, 0, len(exampleData))

	for i := 0; i < len(exampleData); i++ {
		data := []byte(exampleData[i])

		if externalMatchFunction(data) {
			result = append(result, data)
		}
	}

	return result
}

func externalMatchFunction(data []byte) bool {
	return len(data)%2 == 0
}
