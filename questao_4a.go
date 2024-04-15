package main

import(
	"fmt"
	"math"
	"sort"
	"strings"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"

	"aprendizado-maquina/utils"
)

type Dataset struct {
	ErrorRate		float64
	AttrNumber		int
	AttrList		[]string
}

type CVRating struct {
	Attr 	string
	CV		float64
}

func main() {

	//generateDatasets()
	plotGraph()

}


func plotGraph() {

	filename := "cv_ordered/glass"
	//filename := "cv_ordered/ionosphere"

	data := utils.ReadFile(filename + "_cv.csv")

	var list []Dataset

	for _,line := range data {
		errorRate,_ := strconv.ParseFloat(line[0],64)
		attrNumber,_ := strconv.ParseInt(line[1], 10, 32)
		attrList := strings.Split(line[2], " ")

		list = append(list, Dataset{errorRate, int(attrNumber), attrList})
	}

	xy := map[int]float64{} 
	
	for _,item := range list {
		xy[item.AttrNumber] = item.ErrorRate
	}

	pts := make(plotter.XYs, len(xy))

	for i := range pts {
		pts[i].X = float64(i+1)
		pts[i].Y = xy[i+1]

	}

	p := plot.New()

	p.Title.Text = "Taxa de Erro vs CV"
	p.X.Label.Text = "quantidade de atributos considerados"
	p.Y.Label.Text = "Taxa de Erro"

	err := plotutil.AddLinePoints(p,pts)
	if err != nil {
		panic(err)
	}

	if err := p.Save(8*vg.Inch, 8*vg.Inch, "histogramas/glass.png"); err != nil {
		panic(err)
	}
	
}


func generateDatasets() {
	// 1 parte : criar os datasets com conjuntos de dados
	// organizados por ordem crescente de CV

	data := utils.ReadFile("ionosphere.csv")
	//data := utils.ReadFile("glass.csv")

	// Calculando CV para cada atributo
	ratingList := getCVRatingListSorted(data)

	for _,item := range ratingList {
		fmt.Println(item)
	}

	datasets := []Dataset{}
	for i,value := range ratingList {

		var dataset Dataset

		if(i == 0) {
			dataset = Dataset{0.0, i+1, []string{value.Attr}}
		} else {
			dataset = Dataset{0.0, i+1, append(datasets[i-1].AttrList, value.Attr)}
		}

		datasets = append(datasets, dataset)
	}

	for _, item := range datasets {
		fmt.Println(item)
	}
}

func getCVRatingListSorted(data [][]string) []CVRating {
	ratingList := []CVRating{}
	for i, _ := range data[0] {
		if(i == len(data[0]) - 1) {
			ratingList = append(ratingList, CVRating{"class", 0.0})
		} else {
			ratingList = append(ratingList, CVRating{"a"+strconv.Itoa(i), 0.0})
		}
	}

	result := []CVRating{}
	for i,item := range ratingList {
		valueList := []float64{}
		
		for _, line := range data {
			v,_ := strconv.ParseFloat(line[i], 64)
			valueList = append(valueList,v)
		}

		avg := utils.AvgFloat(valueList)
		variance := utils.UnbiasedVarianceFloat(valueList, avg)
		cv := utils.AbsFloat(math.Sqrt(variance) / avg)

		if(math.IsNaN(cv)) {
			continue
		}

		result = append(result, CVRating{item.Attr, cv})
	}
	
	sort.Slice(result, func(i, j int) bool {
		return result[i].CV < result[j].CV
	})

	return result
}

