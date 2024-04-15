package main

import(
	"fmt"
	"math"
	"sort"
	"strconv"
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

	generateDatasets()
	//plot()

}

/*
func plot() {

	filename := "glass"
	filename := "ionosphere"

	data := utils.ReadFile(filename + "_cv.csv")
	
}
*/

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

