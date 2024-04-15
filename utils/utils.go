package utils

import(
	"encoding/csv"
	"strconv"
	"math"
	"log"
    "os"
	
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)


type FrequencyDistribution struct {
	Interval 		string
	MinValue 		int
	MaxValue 		int
	MiddlePoint 	int
	IntervalSize 	int
	RelFreqPer		int
	CumAbsFreq		int
	CumRelFreqPer	int
	Values 			[]int
}

func ReadFile(filename string) [][]string {
	file, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	
		csvReader := csv.NewReader(file)
		data, err := csvReader.ReadAll()
		if err != nil {
			log.Fatal(err)
		}
		return data
}

func GenerateDistributionTable(k int, h int, data []int) []FrequencyDistribution {
	result := make([]FrequencyDistribution, 0)

	// gerando a tabela vazia
	for i, interval := 0, data[0]; i < k; i++ {
		limitMin := interval
		limitMax := 0

		if(i == k - 1){
			limitMax = data[len(data) - 1]
		} else {
			limitMax = interval + h
		}

		intervalName := "[" + strconv.Itoa(limitMin) + ", " + strconv.Itoa(limitMax) + ")"
		item := FrequencyDistribution{intervalName,limitMin,limitMax,0,0,0,0,0, make([]int, 0)}
		result = append(result, item)
		interval = limitMax
	}

	// preenchendo com o arranjo de dados
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(result); j++ {
			if (data[i] >= result[j].MinValue && data[i] < result[j].MaxValue){
				result[j].Values = append(result[j].Values, data[i])
				result[j].IntervalSize++
				break
			} else if (j == len(result) - 1 && data[i] == result[j].MaxValue) {
				result[j].Values = append(result[j].Values, data[i])
				result[j].IntervalSize++
			}
		} 
	}

	for i, cumulative := 0, 0; i < len(result); i++ {
		result[i].MiddlePoint = (result[i].MaxValue + result[i].MinValue) / 2
		result[i].RelFreqPer = int(((float64(result[i].IntervalSize) / float64(len(data))) * 100) + 0.5)
		
		cumulative = cumulative + result[i].IntervalSize
		result[i].CumAbsFreq = cumulative
		result[i].CumRelFreqPer = int(((float64(cumulative) / float64(len(data))) * 100) + 0.5)
	}

	return result
}


func PlotHistogram(title string,filename string, data []int,k int) {
	
	v := make(plotter.Values, len(data))
	for i := range v {
		v[i] = float64(data[i])
	}
	p := plot.New()
	p.Title.Text = title

	h, err := plotter.NewHist(v, k)
	if err != nil {
		panic(err)
	}

	p.Add(h)

	if err := p.Save(4*vg.Inch, 4*vg.Inch,"histogramas/" + filename + ".png"); err != nil {
		panic(err)
	}
}

func Avg(data []int) float64 {
	soma := 0;
	for _, value := range data {
		soma = soma + value
	}

	return float64(soma) / float64(len(data))
}

func AvgFloat(data []float64) float64 {
	soma := 0.0;
	for _, value := range data {
		soma = soma + value
	}

	return soma / float64(len(data))
}

func UnbiasedVariance(data []int, avg float64) float64 {
	sum := 0.0
	for _,value := range data {
		sum = sum + math.Pow(float64(value) - avg, 2)
	}

	return sum / (float64(len(data)) - 1.0)
}

func UnbiasedVarianceFloat(data []float64, avg float64) float64 {
	sum := 0.0
	for _,value := range data {
		sum = sum + math.Pow(value - avg, 2)
	}

	return sum / (float64(len(data)) - 1.0)
}

func AbsFloat(x float64) float64 {
	return absDiffFloat(x, 0.0)
 }

func absDiffFloat(x, y float64) float64 {
	if x < y {
   		return y - x
	}
	return x - y
}