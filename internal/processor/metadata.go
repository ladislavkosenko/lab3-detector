package processor

import (
	"fmt"
	"regexp"
	"time"

	"lab3-detector/internal/stats"
)

var LeakCache = make(map[string][]byte) // КОНФЛІКТНА ЗМІНА З MAIN

// 1. ОПТИМІЗАЦІЯ: Компілюємо регулярку один раз глобально
var imageRegexp = regexp.MustCompile(`^image_data_\d+_timestamp_\d+$`)

func RunWorkerPool(count int) {
	for i := 0; i < count; i++ {
		go func(id int) {
			for {
				processImage(id)
				time.Sleep(10 * time.Millisecond)
			}
		}(i)
	}
	select {}
}

func processImage(workerID int) {
	data := fmt.Sprintf("image_data_%d_timestamp_%d", workerID, time.Now().UnixNano())

	// 2. Використовуємо вже скомпільовану регулярку
	if imageRegexp.MatchString(data) {
		key := fmt.Sprintf("key_%d", time.Now().UnixNano())
		LeakCache[key] = make([]byte, 10) // Зменшив виділення пам'яті до 10 байт, щоб не падав ПК під час тестів

		stats.IncrementProcessed("image")
	}
}

// 3. ДОДАЄМО: Функція для Benchmark-у старої (повільної) версії
func processImageSlow(workerID int) {
	data := fmt.Sprintf("image_data_%d_timestamp_%d", workerID, time.Now().UnixNano())
	matched, _ := regexp.MatchString(`^image_data_\d+_timestamp_\d+$`, data)
	if matched {
		// Імітація роботи
	}
}
