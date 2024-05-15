package main

//https://vertabelo.com/blog/automobile-repair-shop-data-model/
import (
	"encoding/csv"
	"fmt"
	"graph/graphics"
	"image/color"
	"math"
	"math/rand"
	"os"
	"time"

	"gonum.org/v1/gonum/stat/distuv"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type Row struct {
	year int
	time.Month
	week       int
	buyers     int
	weekbuyers int
	viewers    int
	income     int
	conversion float32
}

func main() {
	graphics.AllGraphics()
	//file := createCsv()
	//defer file.Close()
	//plt := base1()
	//rows := generateData(plt)
	//writeCsv(file, rows)
	//random()
}

func createCsv() *os.File {
	file, err := os.Create("data.csv")
	if err != nil {
		panic(err)
	}
	return file
}

func writeCsv(file *os.File, rows *[]Row) {
	// Создаем писатель CSV
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Записываем заголовки
	headers := []string{"Year", "Month", "Week", "Buyers", "Viewers", "Income", "Conversion"}
	if err := writer.Write(headers); err != nil {
		panic(err)
	}

	for _, row := range *rows {
		record := []string{
			fmt.Sprintf("%d", row.year),
			row.Month.String(),
			fmt.Sprintf("%d", row.week),
			fmt.Sprintf("%d", row.buyers),
			fmt.Sprintf("%d", row.viewers),
			fmt.Sprintf("%d", row.income),
			fmt.Sprintf("%.2f", row.conversion),
		}
		if err := writer.Write(record); err != nil {
			panic(err)
		}
	}
}

func generateData(pts plotter.XYs) *[]Row {
	minBuyers, maxBuyers := float64(80), float64(189)
	arr := []Row{}
	date := time.Date(2019, time.April, 5, 0, 0, 0, 0, time.UTC)
	//currentMonth := date.Month()
	//currentByuers := 0
	prev := Row{
		buyers: int(minBuyers),
	}
	// Создаем нормальное распределение с заданным средним и стандартным отклонением
	mean := 5.0
	stdDev := 2.0
	normalDist := distuv.Normal{
		Mu:    mean,
		Sigma: stdDev,
	}

	meanIncome := 2800.0
	stdDevIncome := 1000.0
	rightShift := 1000.0
	normalDistIncome := distuv.Normal{
		Mu:    meanIncome + rightShift,
		Sigma: stdDevIncome,
	}

	for i := 1; i < 209; i++ {
		num := i
		for num > 57 {
			num = num - 57
			if num < 0 {
				num *= -1
			}
		}
		var conv float32 = 3
		if i >= 168 {
			conv = 2
		}
		row := Row{
			year:       date.AddDate(0, 0, 7*i).Year(),
			week:       i,
			Month:      date.AddDate(0, 0, 7*i).Month(),
			conversion: conv + rand.Float32(),
		}
		scale := pts[num].Y + 1
		extra := float64(1)
		if i > 168 && i < 185 {
			extra = 1.0006
			scale -= 0.02
		} else if i > 184 && i < 200 {
			//extra = 0.8 + 0.02*float64(200-i)
			//scale -= 0.04
			extra = 0.991
			scale -= 0.05
			//} else if i > 200 {
			//	extra = 0.991
			//	scale -= 0.08
		} else {
			extra = 1.00009
			scale -= 0.03822
		}
		//newValue := math.Log(float64(prev.buyers)) + float64(prev.buyers)
		if float64(prev.buyers)*(scale)*extra > minBuyers {
			if float64(prev.buyers)*(scale)*extra < maxBuyers && (float64(prev.buyers)*(scale)*extra < 160 || i > 168) {
				row.buyers = int(float64(prev.buyers) * (scale) * extra)
			} else {
				if i < 168 {
					var buf = float64(161)
					border := float64(160)
					if row.Month == time.November || row.Month == time.October {
						border = 170
						buf = 160 + float64(rand.Int31n(9))
					}
					for buf > border {
						var s = normalDist.Rand()
						if s > 6 {
							buf += s - 5
							break
						}
						buf -= s
					}
					row.buyers = int(buf)
					//row.buyers = 160 - int(maxBuyers) + int(rand.Int31n(10))
				} else {
					row.buyers = int(maxBuyers) - int(rand.Int31n(5))
				}
			}
		} else {
			if i < 168 {
				row.buyers = int(minBuyers) + int(rand.Int31n(10))
			} else {
				row.buyers = int(minBuyers) + int(rand.Int31n(10))
			}
		}
		row.viewers = int(float32(row.buyers) / row.conversion * 100)
		if i > 180 && i < 200 {
			row.income = row.buyers * (int(normalDistIncome.Rand()) + 500)
		} else {
			row.income = row.buyers * int(normalDistIncome.Rand())
		}
		//if {}

		arr = append(arr, row)
		prev = row
	}

	// Создаем новый график
	p := plot.New()

	ptsN := make(plotter.XYs, len(arr))
	for i, row := range arr {
		ptsN[i].Y = float64(row.buyers)
		ptsN[i].X = float64(i)
	}
	// Добавляем точки на график
	line, err := plotter.NewLine(ptsN)
	if err != nil {
		panic(err)
	}
	p.Add(line)

	if err := p.Save(4*vg.Points(100), 4*vg.Points(100), "buyers.png"); err != nil {
		panic(err)
	}

	// Создаем новый график
	pV := plot.New()

	ptsNV := make(plotter.XYs, len(arr))
	for i, row := range arr {
		ptsNV[i].Y = float64(row.viewers)
		ptsNV[i].X = float64(i)
	}
	// Добавляем точки на график
	line, err = plotter.NewLine(ptsNV)
	if err != nil {
		panic(err)
	}
	pV.Add(line)

	if err := pV.Save(4*vg.Inch, 4*vg.Inch, "viewvers.png"); err != nil {
		panic(err)
	}

	// Создаем новый график
	pC := plot.New()

	ptsNC := make(plotter.XYs, len(arr))
	for i, row := range arr {
		ptsNC[i].Y = float64(row.conversion)
		ptsNC[i].X = float64(i)
	}
	// Добавляем точки на график
	line, err = plotter.NewLine(ptsNC)
	if err != nil {
		panic(err)
	}
	pC.Add(line)

	if err := pC.Save(4*vg.Inch, 4*vg.Inch, "conversion.png"); err != nil {
		panic(err)
	}

	// Создаем новый график
	pI := plot.New()

	ptsNI := make(plotter.XYs, len(arr))
	for i, row := range arr {
		ptsNI[i].Y = float64(row.income)
		ptsNI[i].X = float64(i)
	}
	// Добавляем точки на график
	line, err = plotter.NewLine(ptsNI)
	if err != nil {
		panic(err)
	}
	pI.Add(line)

	if err := pI.Save(4*vg.Inch, 4*vg.Inch, "income.png"); err != nil {
		panic(err)
	}

	// Создаем новый график
	pCHK := plot.New()

	ptsCHK := make(plotter.XYs, len(arr))
	for i, row := range arr {
		ptsCHK[i].Y = float64(row.income / row.buyers)
		ptsCHK[i].X = float64(i)
	}
	// Добавляем точки на график
	line, err = plotter.NewLine(ptsCHK)
	if err != nil {
		panic(err)
	}
	pCHK.Add(line)

	if err := pCHK.Save(4*vg.Inch, 4*vg.Inch, "sr_check.png"); err != nil {
		panic(err)
	}

	// Создаем новый график
	pMB := plot.New()

	ptsMB := make(plotter.XYs, 51)
	//currentMonth := time.April
	monthBuyers, count := 0, 0
	for i, row := range arr {
		if i != 0 && i%4 == 0 {
			ptsMB[count].Y = float64(monthBuyers)
			ptsMB[count].X = float64(count)
			monthBuyers = row.buyers
			//currentMonth = row.Month
			count++
		} else {
			monthBuyers += row.buyers
		}

	}
	// Добавляем точки на график
	line, err = plotter.NewLine(ptsMB)
	if err != nil {
		panic(err)
	}
	pMB.Add(line)

	if err := pMB.Save(4*vg.Inch, 4*vg.Inch, "monthBuyers.png"); err != nil {
		panic(err)
	}
	return &arr
}

func base() plotter.XYs {
	// Создаем новый график
	p := plot.New()

	// Создаем точки для графика
	n := 100
	pts := make(plotter.XYs, n)
	for i := range pts {
		x := float64(i) * 2 * math.Pi / float64(n)
		y := math.Sin(x)
		pts[i].X = x
		pts[i].Y = y
	}

	// Добавляем точки на график
	line, err := plotter.NewLine(pts)
	if err != nil {
		panic(err)
	}
	p.Add(line)

	// Сохраняем график в файл
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "sin_wave.png"); err != nil {
		panic(err)
	}
	return pts
}
func base1() plotter.XYs {
	// Создаем новый график
	p := plot.New()

	// Создаем точки для графика
	n := 300
	pts := make(plotter.XYs, n)
	scatterData := make(plotter.XYs, 58)
	var sins int32 = 5
	for i := range pts {
		x := float64(i) * 2 * math.Pi / float64(n)
		y := math.Sin(x)
		for i := 0; i < int(sins); i++ {
			y *= math.Sin(x+float64(rand.Int31n(sins)+1)) * math.Cos(x+float64(rand.Int31n(sins)+1))
		}
		pts[i].X = x
		if y < 0.1 {
			y += 0.05
		} else if y > 0.3 {
			y -= 0.25
		}
		pts[i].Y = y
	}

	// Создаем нормальное распределение с заданным средним и стандартным отклонением
	mean := 50.0
	stdDev := 30.0
	normalDist := distuv.Normal{
		Mu:    mean,
		Sigma: stdDev,
	}

	for i := 0; i < 58; i++ {
		rn := normalDist.Rand()
		scatterData[i].X = pts[int(rn)].X
		scatterData[i].Y = pts[int(rn)].Y
	}
	// Добавляем точки на график
	line, err := plotter.NewLine(pts)
	if err != nil {
		panic(err)
	}
	p.Add(line)

	// Создаем график точек
	scatter, err := plotter.NewScatter(scatterData)
	if err != nil {
		panic(err)
	}
	// Устанавливаем цвет точек
	scatter.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255} // Красный цвет
	p.Add(scatter)

	// Сохраняем график в файл
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "sin_product.png"); err != nil {
		panic(err)
	}
	return scatterData
}

func random() {
	// Создаем новый график
	p := plot.New()

	// Создаем случайные точки для графика
	n := 100
	pts := make(plotter.XYs, n)
	for i := range pts {
		x := rand.Float64() * 10 // Генерируем случайные значения x
		y := rand.Float64() * 10 // Генерируем случайные значения y
		pts[i].X = x
		pts[i].Y = y
	}

	// Добавляем точки на график
	scatter, err := plotter.NewScatter(pts)
	if err != nil {
		panic(err)
	}
	p.Add(scatter)

	// Сохраняем график в файл
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "random_plot.png"); err != nil {
		panic(err)
	}
}
