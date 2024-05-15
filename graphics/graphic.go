package graphics

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type Row struct {
	Year    int
	Month   time.Month
	Week    int
	Buyers  int
	Viewers int
	Income  int
}

// Создаем слайс для хранения структур Row
var rows []Row

func init() {
	// Открываем CSV файл
	file, err := os.Open("data_mod_ok.csv")
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	// Создаем Reader для CSV файла
	reader := csv.NewReader(file)

	// Считываем все строки из CSV файла
	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Ошибка чтения CSV файла:", err)
		return
	}

	// Проходим по строкам CSV файла и преобразуем их в структуры Row
	for i, line := range lines {
		if i == 0 {
			continue
		}
		line = strings.Split(line[0], ";")
		// Преобразуем строки в соответствующие типы данных
		year, _ := strconv.Atoi(line[0])
		month, _ := strconv.Atoi(line[1])
		week, _ := strconv.Atoi(line[2])
		buyers, _ := strconv.Atoi(line[3])
		viewers, _ := strconv.Atoi(line[4])
		income, _ := strconv.Atoi(line[5])

		// Создаем новую структуру Row и добавляем ее в слайс rows
		row := Row{
			Year:    year,
			Month:   time.Month(month),
			Week:    week,
			Buyers:  buyers,
			Viewers: viewers,
			Income:  income,
		}
		rows = append(rows, row)
	}
}

func AllGraphics() {
	Buyers()
	BuyersMonth()
	ViewersMonth()
}

func Buyers() {
	// Создаем новый график
	p := plot.New()

	ptsN := make(plotter.XYs, len(rows))
	for i, row := range rows {
		ptsN[i].Y = float64(row.Buyers)
		ptsN[i].X = float64(i)
	}
	// Добавляем точки на график
	line, err := plotter.NewLine(ptsN)
	if err != nil {
		panic(err)
	}
	p.Add(line)

	// Устанавливаем подписи к осям и заголовок
	p.X.Label.Text = "Месяцы"
	p.Y.Label.Text = "Недели"
	p.Title.Text = "Покупатели по неделям"

	if err := p.Save(4*vg.Points(100), 4*vg.Points(100), "buyers_new.png"); err != nil {
		panic(err)
	}
}

func BuyersMonth() {
	// Создаем новый график
	p := plot.New()

	ptsN := make(plotter.XYs, 51)
	count, summ := 0, 0
	for i, row := range rows {
		if i != 0 && i%4 == 0 {
			ptsN[count].Y = float64(summ)
			ptsN[count].X = float64(count)
			summ = 0
			count++
		}
		summ += row.Buyers
	}
	// Добавляем точки на график
	line, err := plotter.NewLine(ptsN)
	if err != nil {
		panic(err)
	}
	p.Add(line)

	// Устанавливаем подписи к осям и заголовок
	p.X.Label.Text = "Месяца"
	p.Y.Label.Text = "Покупатели"
	p.Title.Text = "Покупатели по месяцам"

	if err := p.Save(4*vg.Points(100), 4*vg.Points(100), "buyersMonth_new.png"); err != nil {
		panic(err)
	}
}

func ViewersMonth() {
	// Создаем новый график
	p := plot.New()

	ptsN := make(plotter.XYs, 51)
	count, summ := 0, 0
	for i, row := range rows {
		if i != 0 && i%4 == 0 {
			ptsN[count].Y = float64(summ)
			ptsN[count].X = float64(count)
			summ = 0
			count++
		}
		summ += row.Viewers
	}
	// Добавляем точки на график
	line, err := plotter.NewLine(ptsN)
	if err != nil {
		panic(err)
	}
	p.Add(line)

	// Устанавливаем подписи к осям и заголовок
	p.X.Label.Text = "Месяца"
	p.Y.Label.Text = "Посетители"
	p.Title.Text = "Посетители по месяцам"

	if err := p.Save(4*vg.Points(100), 4*vg.Points(100), "viewersMonth_new.png"); err != nil {
		panic(err)
	}
}
