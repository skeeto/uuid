package bench

import (
	"testing"

	google "github.com/google/uuid"
	gofrs "github.com/gofrs/uuid"
	self "nullprogram.com/x/uuid"
)

func BenchmarkSelf(b *testing.B) {
	g := self.NewGen()
	for i := 0; i < b.N; i++ {
		g.NewV4()
	}
}

func BenchmarkGofrs(b *testing.B) {
	g := gofrs.NewGen()
	for i := 0; i < b.N; i++ {
		_, err := g.NewV4()
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkGoogle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		google.New()
	}
}
