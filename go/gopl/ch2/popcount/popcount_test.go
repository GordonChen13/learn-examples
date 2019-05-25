package popcount

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestPopCount(t *testing.T) {
	cases := []struct{
		name string
		in uint64
		want int
	} {
		{
			"input 10000 and return 5",
			10000,
			5,
		},
		{
			"input 100000000 and return 12",
			100000000,
			12,
		},
		{
			"input 1000000000000 and return 13",
			1000000000000,
			13,
		},
	}

	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			get := PopCount(item.in)
			assert.Equal(t, get, item.want)
		})
	}

}

func TestPopCount3(t *testing.T) {
	cases := []struct{
		name string
		in uint64
		want int
	} {
		{
			"input 10000 and return 5",
			10000,
			5,
		},
		{
			"input 100000000 and return 12",
			100000000,
			12,
		},
		{
			"input 1000000000000 and return 13",
			1000000000000,
			13,
		},
	}

	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			get := PopCount3(item.in)
			assert.Equal(t, get, item.want)
		})
	}
}

func TestPopCount4(t *testing.T) {
	cases := []struct{
		name string
		in uint64
		want int
	} {
		{
			"input 10000 and return 5",
			10000,
			5,
		},
		{
			"input 100000000 and return 12",
			100000000,
			12,
		},
		{
			"input 1000000000000 and return 13",
			1000000000000,
			13,
		},
	}

	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			get := PopCount4(item.in)
			assert.Equal(t, get, item.want)
		})
	}
}

func TestPopCount5(t *testing.T) {
	cases := []struct{
		name string
		in uint64
		want int
	} {
		{
			"input 10000 and return 5",
			10000,
			5,
		},
		{
			"input 100000000 and return 12",
			100000000,
			12,
		},
		{
			"input 1000000000000 and return 13",
			1000000000000,
			13,
		},
	}

	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			get := PopCount5(item.in)
			assert.Equal(t, get, item.want)
		})
	}
}

func BenchmarkPopCount10000(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		PopCount(10000)
	}
}

func BenchmarkPopCount100000000(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		PopCount(100000000)
	}
}

func BenchmarkPopCount1000000000000(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		PopCount(1000000000000)
	}
}

func BenchmarkPopCount310000(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		PopCount3(10000)
	}
}

func BenchmarkPopCount3100000000(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		PopCount3(100000000)
	}
}

func BenchmarkPopCount31000000000000(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		PopCount3(1000000000000)
	}
}

func BenchmarkPopCount410000(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		PopCount4(10000)
	}
}

func BenchmarkPopCount4100000000(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		PopCount4(100000000)
	}
}

func BenchmarkPopCount41000000000000(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		PopCount4(1000000000000)
	}
}

func BenchmarkPopCount510000(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		PopCount5(10000)
	}
}

func BenchmarkPopCount5100000000(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		PopCount5(100000000)
	}
}

func BenchmarkPopCount51000000000000(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		PopCount5(1000000000000)
	}
}
