package processor

import (
	"fmt"
	"regexp"
	"time"

	"lab3-detector/internal/stats"

	"go.uber.org/zap"
)

var LeakCache = make(map[string][]byte) // КОНФЛІКТНА ЗМІНА З MAIN
var imageRegexp = regexp.MustCompile(`^image_data_\d+_timestamp_\d+$`)
var logger, _ = zap.NewProduction()

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

	if imageRegexp.MatchString(data) {
		key := fmt.Sprintf("key_%d", time.Now().UnixNano())
		LeakCache[key] = make([]byte, 10)

		stats.IncrementProcessed("image")

		logger.Info("Image successfully processed",
			zap.Int("workerID", workerID),
			zap.String("key", key),
		)
	}
}

// Функція для Benchmark-у (якщо вона була)
func processImageSlow(workerID int) {
	data := fmt.Sprintf("image_data_%d_timestamp_%d", workerID, time.Now().UnixNano())
	matched, _ := regexp.MatchString(`^image_data_\d+_timestamp_\d+$`, data)
	if matched {
		// Імітація роботи
	}
}
