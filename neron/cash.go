package neron

import (
	"math"
	"math/rand"
	"time"
)

var Neron neron_cash

type neron_cash struct {
	NODE   [][]node
	Output []node
	Input  []float64
}
type node struct {
	W []float64
	H float64
}

/*init sistem Neron*/
func (nerv *neron_cash) NewNode(input, output, W_size, W_deep uint32) {
	//Создаем двумерный массив node
	nerv.NODE = make([][]node, W_deep)
	//Создание кол-во входных/выходных сигналов
	nerv.Input = make([]float64, input)
	nerv.Output = make([]node, output)

	for X := range nerv.NODE {
		//Создаем второй слой двумерного массива NODE[]
		nerv.NODE[X] = make([]node, W_size)

		for Y := range nerv.NODE[X] {
			//Создаем на первом слое синапсы, равными входных сигналов
			if X == 0 {
				nerv.NODE[X][Y].W = make([]float64, len(nerv.Input))
			} else {
				//Проходим по каждой node и создаем массив синапсов
				nerv.NODE[X][Y].W = make([]float64, W_size)
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
*/
func (nerv *neron_cash) Process() []float64 {
	output := make([]float64, len(nerv.Output))
	//расчет первых слоев
	for Y := range nerv.NODE[0] {
		for W := range nerv.NODE[0][Y].W {
			nerv.NODE[0][Y].H += nerv.NODE[0][Y].W[W] * nerv.Input[W]
		}
		nerv.NODE[0][Y].Sigmoid()
	}
	//nerv.Sigmoid()
	for X := 1; X < len(nerv.NODE); X++ {
		for Y := range nerv.NODE[X] {
			for W := range nerv.NODE[X][Y].W {
				nerv.NODE[X][Y].H += nerv.NODE[X][Y].W[W] * nerv.NODE[X-1][W].H
			}
			nerv.NODE[X][Y].Sigmoid()
		}
	}
	//Прохождение последнего слоя
	for X := range nerv.Output {
		for W := range nerv.Output[X].W {
			nerv.Output[X].H += nerv.Output[X].W[W] * nerv.NODE[len(nerv.NODE)-1][W].H
		}
		nerv.Output[X].Sigmoid()
	}

	//передача аргументов на выход
	for Y := range nerv.Output {
		output[Y] = nerv.Output[Y].H
	}

	return output
}

func (nerv *node) Sigmoid() {
	nerv.H = math.Tanh(nerv.H)
}
