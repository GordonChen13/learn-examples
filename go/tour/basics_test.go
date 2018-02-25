package tour

import "testing"

func TestPackages(t *testing.T) {
	Packages()
}

func BenchmarkPackages(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Packages()
	}
}

func TestImport(t *testing.T) {
	Import()
}
