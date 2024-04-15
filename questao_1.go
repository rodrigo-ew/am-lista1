package main

import (
	"fmt"
	"strconv"
	"slices"
	"math"

    "aprendizado-maquina/utils"
)

func main() {

	// Taxa de colesterol total para 30 indivíduos
	data := []int {140, 160, 168, 180, 180, 
				   180, 180, 184, 185, 190,
				   190, 192, 192, 196, 200, 
				   200, 200, 205, 205, 208, 
				   214, 214, 220, 220, 225, 
				   230, 240, 260, 280, 315}

	slices.Sort(data)

	// distribuição de frequencia e histograma
	fd := fdSturges(data)
	utils.PlotHistogram("Histograma de distribuição de frequência - taxa de colesterol", "questao1", data, len(fd))

	////////
	// obtendo a média e a mediana
	media := utils.Avg(data)

	central, centralPlus1 := data[len(data)/2], data[len(data)/2 + 1] 
	mediana := float64(central + centralPlus1) / 2.0

	fmt.Println("A média é: " + strconv.FormatFloat(media, 'f', -1, 64))
	fmt.Println("A mediana é: " + strconv.FormatFloat(mediana, 'f', -1, 64))

	////////
	// Variância e desvio padrão

	variance := utils.UnbiasedVariance(data,media)
	fmt.Println("A variância é: " + strconv.FormatFloat(variance, 'f', 2, 64))
	fmt.Println("Desvio padrão: " + strconv.FormatFloat(math.Sqrt(variance), 'f', 2, 64))

	////////

	fmt.Println("\nDados organizados em classes\n")
	fmt.Println("classe|ponto médio|qtd|freq. relativa|freq. acumulada abs|freq. acumulada rel|items")
	for _, value := range fd {
		fmt.Printf("%s|%d|%d|%d%%|%d|%d%%|%v\n", value.Interval, value.MiddlePoint,value.IntervalSize,value.RelFreqPer,value.CumAbsFreq,value.CumRelFreqPer,value.Values)
	}	
}

func fdSturges(data []int)[]utils.FrequencyDistribution {
	// encontrando a quantidade de classes k
	// através da regra de Sturges
	k := int(1 + 3.3 * math.Log10(float64(len(data))) + 0.5)
	
	// determinando amplitude das classes
	h := (data[len(data) - 1] - data[0])/k

	//gerando os intervalos
	fd := utils.GenerateDistributionTable(k, h, data)

	return fd
}


