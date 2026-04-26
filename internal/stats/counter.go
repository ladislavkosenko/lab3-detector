package stats

import "sync"

// Спільна мапа
var GlobalStats = make(map[string]int)

// Створюємо м'ютекс (замок) для захисту нашої мапи
var mutex sync.Mutex

func IncrementProcessed(imageType string) {
	mutex.Lock() // Закриваємо двері (інші потоки чекають)
	GlobalStats[imageType]++
	mutex.Unlock() // Відкриваємо двері (наступний потік може зайти)
}
