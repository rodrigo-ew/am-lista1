package main


import (
	"strconv"
    "fmt"

	"aprendizado-maquina/utils"
)

func main() {
		data := utils.ReadFile("iris.csv")
		list := extractData(data)

		// Calculando variância
		avg := utils.AvgFloat(list)
		variance := utils.UnbiasedVarianceFloat(list, avg)

		fmt.Println("A variância calculada para o 3º parâmetro é: " + strconv.FormatFloat(variance, 'f', 2, 64))

}

func extractData(data [][]string) []float64 {
	var list []float64
    for _, line := range data {
		var n float64
		for j, field := range line {
			if j == 2 {
				n,_ = strconv.ParseFloat(field, 64)
			} 
		}
		list = append(list, n)
	
    }
    return list
}