package main

import (
	"log"
	"net/http"
	_ "net/http/pprof" // Обов'язковий імпорт для запуску профілювальника

	"lab3-detector/internal/processor"
)

func main() {
	// Запускаємо pprof сервер у фоновому потоці
	go func() {
		log.Println("Pprof server started on :6060")
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	log.Println("Image Metadata Processor started...")
	// Запускаємо пул воркерів, які генерують навантаження і помилки
	processor.RunWorkerPool(5)
}
