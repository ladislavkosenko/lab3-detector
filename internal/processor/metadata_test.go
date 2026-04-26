package processor

import (
	"testing"
)

// Тест ПОВІЛЬНОЇ версії (де регулярка компілюється всередині)
func BenchmarkProcessImageSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		processImageSlow(1)
	}
}

// Тест ШВИДКОЇ версії (де регулярка скомпільована глобально)
func BenchmarkProcessImageFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		processImage(1) // викликаємо нашу оптимізовану функцію
	}
}
