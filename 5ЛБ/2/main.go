package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"time"
)

func filter(img *image.RGBA) {
	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := img.RGBAAt(x, y)
			// Перевод в оттенки серого
			gray := uint8((uint16(c.R) + uint16(c.G) + uint16(c.B)) / 3)
			img.SetRGBA(x, y, color.RGBA{R: gray, G: gray, B: gray, A: c.A})
		}
	}
}

func main() {
	file, err := os.Open("C:/Users/m9775/OneDrive/Изображения/4.1/foto.png")
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

	// Преобразование в *image.RGBA
	bounds := img.Bounds()
	rgbaImg := image.NewRGBA(bounds)
	draw.Draw(rgbaImg, bounds, img, bounds.Min, draw.Src)

	// Замер времени обработки
	start := time.Now()
	filter(rgbaImg)
	fmt.Println("Время обработки:", time.Since(start))

	output, err := os.Create("foto2.png")
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		return
	}
	defer output.Close()

	err = png.Encode(output, rgbaImg)
	if err != nil {
		fmt.Println("Ошибка сохранения изображения:", err)
	}
}
