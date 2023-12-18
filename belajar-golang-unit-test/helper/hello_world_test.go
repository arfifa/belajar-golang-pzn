package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Eko",
			request: "Eko",
		},
		{
			name:    "Kurniawan",
			request: "Kurniawan",
		},
		{
			name:    "EkoKurniawanKhannedy",
			request: "Eko Kurniawan Khannedy",
		},
		{
			name:    "Budi",
			request: "Budi Nugraha",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Eko",func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Eko")
		}
	})

	b.Run("Kurniawan",func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Kurniawan")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		HelloWorld("Eko")
	}
}

func BenchmarkHelloWorldKurniawan(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		HelloWorld("Kurniawan")
	}
}

func TestTableHelloWorld(t *testing.T)  {
	tests := []struct {
		name string
		request string
		expected string
	}{
		{
			name: "Arfifa",
			request: "Arfifa",
			expected: "Hello Arfifa",
		},
		{
			name: "Rahman",
			request: "Rahman",
			expected: "Hello Rahman",
		},
		{
			name: "Blackot",
			request: "Blackot",
			expected: "Hello Blackot",
		},
		{
			name: "Eko",
			request: "Eko",
			expected: "Hello Eko",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Arfifa", func(t *testing.T) {
		result := HelloWorld("Arfifa")
		require.Equal(t, "Hello Arfifa", result, "Result must be 'Hello Arfifa'")
	})

	t.Run("Rahman", func(t *testing.T) {
		result := HelloWorld("Rahman")
		require.Equal(t, "Hello Rahman", result, "Result must be 'Hello Rahman'")
	})
}

func TestMain(m *testing.M)  {
	// before
	fmt.Println("BEFORE UNIT TEST")
	
	m.Run()

	// after
	fmt.Println("AFTER UNIT TEST")
}

func TestSkip(t *testing.T)  {
	if runtime.GOOS == "linux" {
		t.Skip("Can not run on Linux")
	}

	result := HelloWorld("Arfifa")
	require.Equal(t, "Hello Arfifa", result, "Result must be 'Hello Arfifa'")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Arfifa")
	require.Equal(t, "Hello Arfifa", result, "Result must be 'Hello Arfifa'")
	fmt.Println("TestHelloWorld with Require Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Arfifa")
	assert.Equal(t, "Hello Arfifa", result, "Result must be 'Hello Arfifa'")
	fmt.Println("TestHelloWorld with Assert Done")
}

func TestHelloWorld(t *testing.T)  {
	result := HelloWorld("Eko")

	// fmt.Println(t)

	if result != "Hello Eko" {
		// error
		t.Error("Result is not 'Hello Eko'")
		// "Result is not 'Hello Eko'"
		fmt.Println("code program dijalankan")
	}
}

func TestHelloArfifa(t *testing.T)  {
	result := HelloWorld("Arfifa")

	// fmt.Println(t)

	if result != "Hello Arfifa" {
		// error
		t.Fatal("Result is not 'Hello Arfifa'")
		// panic("Result is not 'Hello Eko'")

		fmt.Println("Code program dijalankan")
	}
}