package main

import (
	"fmt"
	"math"
	"strconv"

    "aprendizado-maquina/utils"
)

func main() {
	
	data := []int{1,3,2,3,2,2,0,1,0,0,3,0,2,3,2,2,3,3,0,3,2,0}

	// obtendo os diferentes K
	kSturges := getKSturges(data)
	kPow2 := getKPow2(len(data))
	kSqrt := getKSqrt(data)

	fmt.Printf("k obtidos para o conjuntos de dados: \n\nk Sturges: %d ; k Potência de 2: %d ; k Sqrt: %d\n\n",kSturges, kPow2, kSqrt)

	// plottando histogramas para diferentes K
	fmt.Println("Plotando histogramas...")
	utils.PlotHistogram("Histograma - K Sturges", "sturges", data, kSturges)
	utils.PlotHistogram("Histograma - K Potência de 2", "pow2", data, kPow2)
	utils.PlotHistogram("Histograma - K raiz Quadrada", "sqrt", data, kSqrt)

	//média
	media := utils.Avg(data)
	fmt.Println("A média é: " + strconv.FormatFloat(media, 'f', 2, 64))
	
	// Variância e desvio padrão
	variance := utils.UnbiasedVariance(data,media)
	fmt.Println("A variância é: " + strconv.FormatFloat(variance, 'f', 2, 64))
	fmt.Println("Desvio padrão: " + strconv.FormatFloat(math.Sqrt(variance), 'f', 2, 64))
	fmt.Println("O desvio absoluto médio é: " + strconv.FormatFloat(desvioAbsMedio(data),'f',2,64))
}

func getKPow2(length int) int {
	result := 0.0
	for i := 0.0 ; int(math.Pow(2 ,i) + 0.5) <= length; i = i + 1.0 {
		result = i
	}
	return int(result + 0.5)
}

func getKSturges(data []int) int {
	return int(1 + 3.3 * math.Log10(float64(len(data))) + 0.5)
}
func getKSqrt(data []int) int {
	return int(math.Sqrt(float64(len(data))) + 0.5)
}

func desvioAbsMedio(data []int) float64 {
	avg := utils.Avg(data)

	sum := 0.0
	for i := 0 ; i < len(data) ; i++ {
		sum = sum + utils.AbsFloat(float64(data[i]) - avg)
	}

	return sum / float64(len(data))
}

