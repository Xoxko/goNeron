package neron

import (
	"fmt"
	"math/rand"
	"time"
)

var Neron neron_cash

type neron_cash struct {
	NODE   [][]node
	Output []node
	Input  []float64
	H      []float64
}
type node struct {
	W []float64
}

/*init sistem Neron*/
func (nerv *neron_cash) NewNode(input, output, W_size, W_deep uint32) {
	//Создаем двумерный массив node
	nerv.NODE = make([][]node, W_deep)
	//Создание кол-во входных/выходных сигналов
	nerv.Input = make([]float64, input)
	nerv.Output = make([]node, output)
	nerv.H = make([]float64, W_size)

	for i := range nerv.NODE {
		//Создаем второй слой двумерного массива NODE[]
		nerv.NODE[i] = make([]node, W_size)

		for j := range nerv.NODE[i] {
			//Создаем на первом слое синапсы, равными входных сигналов
			if i == 0 {
				nerv.NODE[i][j].W = make([]float64, len(nerv.Input))
			} else {
				//Проходим по каждой node и создаем массив синапсов
				nerv.NODE[i][j].W = make([]float64, W_size)
			}
		}
	}
	// Выходные сигналы это копия Node так как нам нужны синапсы
	for i := range nerv.Output {
		nerv.Output[i].W = make([]float64, W_size)
	}
}

// @addiction -> NewNode
func (nerv *neron_cash) RandomNode() {
	rand.Seed(time.Now().UnixNano())
	for X := range nerv.NODE {
		for Y := range nerv.NODE[X] {
			for W := range nerv.NODE[X][Y].W {
				nerv.NODE[X][Y].W[W] = (rand.Float64() * 2) - 1
			}
		}
	}
	for X := range nerv.Output {
		for W := range nerv.Output[X].W {
			nerv.Output[X].W[W] = (rand.Float64() * 2) - 1
		}
	}
}

/*
@addiction -> NewNode -> RandomNode
#Обработка данных
Массив данных должны быть равен массиву
входных нейронов, в случае если не равны,
вернет ошибку и нулевые значения в массиве
*/
func (nerv *neron_cash) Process(input []float64) ([]float64, error) {
	output := make([]float64, len(nerv.Output))
	if len(nerv.H) != len(input) {
		return output, fmt.Errorf("error: input data size nerv.H(%v) != input(%v)", len(nerv.H), len(input))
	}

	nerv.Input = input

	for Y := range nerv.NODE[0] {
		for W := range nerv.NODE[0][Y].W {
			nerv.H[Y] = nerv.NODE[0][Y].W[W]*nerv.Input[W] + nerv.H[Y]
		}

	}

	for X := range nerv.NODE {
		for Y := range nerv.NODE[X] {
			for W := range nerv.NODE[X][Y].W {
				nerv.H[Y] = nerv.NODE[X][Y].W[W]*nerv.Input[W] + nerv.H[Y]
			}

		}
	}

	for X := range nerv.Output {
		for W := range nerv.Output[X].W {
			nerv.Output[X].W[W] = (rand.Float64() * 2) - 1
		}
	}

	return output, nil
}
func (n *node) Len() int {
	return len(n.W)
}

// @addiction -> W_point
func (n *node) Summation(W_number int, EN float64) {

}

// @addiction -> Summation
func (n *node) Sigmoid() {

}
