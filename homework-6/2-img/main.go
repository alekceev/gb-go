package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

// 2. Дополните пример из раздела «Пакет img» изображением горизонтальных и вертикальных линий.
// Воспользуйтесь статьей «Работа с изображениями».

func main() {
	gray := color.RGBA{38, 37, 45, 255}

	// Квадрат
	rectImg := image.NewRGBA(image.Rect(0, 0, 200, 200))

	// Заливаем
	draw.Draw(rectImg, rectImg.Bounds(), &image.Uniform{gray}, image.ZP, draw.Src)

	red := color.RGBA{255, 0, 0, 255}

	// Рамка
	for i := 0; i < 200; i++ {
		rectImg.Set(i, 0, red)
		rectImg.Set(i, 199, red)
		rectImg.Set(0, i, red)
		rectImg.Set(199, i, red)
	}

	file, err := os.Create("rectangle.png")
	if err != nil {
		log.Fatalf("Failed create file: %s", err)
	}
	defer file.Close()
	png.Encode(file, rectImg)
}
