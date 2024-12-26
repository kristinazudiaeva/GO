package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"sync"
	"time"
)

func filterParallel(img draw.Image, y int, wg *sync.WaitGroup) {
	defer wg.Done()

	bounds := img.Bounds()
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		// Преобразуем пиксель в color.RGBA
		c := img.At(x, y)
		r, g, b, a := c.RGBA() // Получаем значения каналов RGBA
		// Преобразуем в серый цвет
		gray := uint16((r + g + b) / 3)
		// Устанавливаем серый пиксель
		img.Set(x, y, color.RGBA64{R: gray, G: gray, B: gray, A: uint16(a)})
	}
}

func main() {
	file, err := os.Open("C:/Users/m9775/OneDrive/Изображения/4.2/foto.png")
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Ошибка декодирования изображения:", err)
		return
	}

	// Преобразование в draw.Image
	drawImg, ok := img.(draw.Image)
	if !ok {
		fmt.Println("Не удалось преобразовать изображение для редактирования")
		return
	}

	// Параллельная обработка
	start := time.Now()
	var wg sync.WaitGroup
	bounds := drawImg.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		wg.Add(1)
		go filterParallel(drawImg, y, &wg)
	}

	wg.Wait()
	fmt.Println("Время параллельной обработки:", time.Since(start))

	output, err := os.Create("foto2.png")
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		return
	}
	defer output.Close()

	err = png.Encode(output, drawImg)
	if err != nil {
		fmt.Println("Ошибка сохранения изображения:", err)
	}
}
