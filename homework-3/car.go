package main

import (
	"fmt"
)

/* 1. Описать несколько структур — любой легковой автомобиль и грузовик.
Структуры должны содержать марку авто, год выпуска, объем багажника/кузова, запущен ли двигатель, открыты ли окна, насколько заполнен объем багажника.
Инициализировать несколько экземпляров структур.
2. Применить к ним различные действия. Вывести значения свойств экземпляров в консоль.
*/


type Car struct {
	Type string
	Model string
	Year  int
	TrunkVolume int
	Seats int
	CarProperty
}

type CarProperty struct {
	StartedEngine bool
	OpenWindows bool
	Trunk int
}

// TruncComplatePersent вернёт заполненность багажника
func (c *Car) TruncComplatePersent() int {
	return 100 * c.Trunk / c.TrunkVolume
}

func main() {
	var cars = []Car{
		Car{
			Type: "Sedan",
			Model: "Toyota Camry",
			Year: 2015,
			TrunkVolume: 300,
			Seats: 5,
			CarProperty: CarProperty{
				StartedEngine: false,
				OpenWindows: false,
				Trunk: 0,
			},
		},
		Car{
			Type: "Minivan",
			Model: "Toyota Ipsum",
			Year: 2002,
			TrunkVolume: 1200,
			Seats: 12,
			CarProperty: CarProperty{
				StartedEngine: false,
				OpenWindows: false,
				Trunk: 0,
			},
		},
	}

	// Загружаем багаэник
	cars[0].Trunk = 200;
	// Запускаем двигатель
	cars[0].StartedEngine = true;
	// Открываем окно
	cars[0].OpenWindows = true;


	var boolName = map[bool]string{
		true: "Да",
		false: "Нет",
	}

	format := `
	Класс автомобиля: %s
	Модель: %s
	Год выпуска: %d
	Объём багажника: %d
	Кол-во сидений: %d
	Двигатель запущен? %v
	Окна открыты? %v
	На сколько заполнен багажник: %d%%
	`

	for _, car := range cars {
		fmt.Printf(format,
			car.Type,
			car.Model,
			car.Year,
			car.TrunkVolume,
			car.Seats,
			boolName[car.StartedEngine],
			boolName[car.OpenWindows],
			car.TruncComplatePersent(),
		)
	}
}