package neron

import (
	"fmt"
	"math"
)

// init struct dataset
var Neron neron_cash

type neron_cash struct {
	NODE [][]node
}
type node struct {
	W     []float64
	H     float64
	delta float64
}

/*Создание массива NODE
 */
func (nerv *neron_cash) NewInitNode(input []int) ([]*float64, []*float64) {
	//Создаем двухмерный массив где учитываем размер входных данных
	nerv.NODE = make([][]node, len(input))

	inp := make([]*float64, input[0])
	out := make([]*float64, input[len(input)-1])
	for x := range nerv.NODE {
		//проходим по первому массиву и создаем второй ряд массива и присваем значения из входных данных
		nerv.NODE[x] = make([]node, input[x])
		for y := range nerv.NODE[x] {
			switch x {
			case 0:
				//учитываем что на первом массиве нам ненужны синапсы поэтому присвоим nil
				nerv.NODE[x][y].W = nil
				inp[y] = &nerv.NODE[x][y].H
			case len(input) - 1:
				/*input[x-1] так как при прохождение Х буде Х>=1 , тем саммы учитываем
				предыдущий слой и создаем синапсы равными количеству предыдущего слоя неронов*/
				nerv.NODE[x][y].W = make([]float64, input[x-1])
				out[y] = &nerv.NODE[x][y].H
			default:
				nerv.NODE[x][y].W = make([]float64, input[x-1])
			}
		}
	}
	return inp, out
}

/*Зависит -> NewInitNode
Рандомизатор синапсов
Принимает функцию
*/
func (nerv *neron_cash) Random(number func() float64) {
	//создаем сыллку на NODE не учитывая первый слой и при этом не выделяеться память
	coppy := nerv.NODE[1:]
	for x := range coppy {
		for y := range coppy[x] {
			for w := range coppy[x][y].W {
				coppy[x][y].W[w] = number()
			}
		}
	}
}

/*Зависит -> NewInitNode -> Random
Производит расчет неросети
*/
func (nerv *neron_cash) Calculation() {

	coppy := nerv.NODE[1:]
	for x := range coppy {
		for y := range coppy[x] {
			coppy[x][y].H = 0
			for w := range coppy[x][y].W {
				coppy[x][y].H += coppy[x][y].W[w] * nerv.NODE[x][w].H
			}
			coppy[x][y].H = sigma(coppy[x][y].H)
		}
	}
}

//сигмоида
func sigma(x float64) float64 {
	return math.Tanh(x) //1.0 / (1.0 + math.Exp(-x))
}

//производная от сигмоиды
func dsigma(x float64) float64 {
	return 1.0 - math.Pow(x, 2) //x * (1.0 - x)
}

/*обучение неросети
реализован по градиентному спуску
*/
func (nerv *neron_cash) Training(ExitError []float64, n float64) {
	if len(ExitError) != len(nerv.NODE[len(nerv.NODE)-1]) {
		panic(fmt.Errorf("ExitError"))
	}
	coppy := nerv.NODE[1:]
	i := len(coppy) - 1

	//расчет ошибки d выходного слоя сети.
	for y := range ExitError {
		coppy[i][y].delta = ExitError[y] - coppy[i][y].H
	}
	//распространение сигнала ошибки d по всем узлам
	for x := range coppy {
		for y := range coppy[i-x] {
			for w := range coppy[i-x][y].W {
				nerv.NODE[i-x][w].delta = coppy[i-x][y].W[w] * coppy[i-x][y].delta
			}
		}
	}
	//корректировка весовых коэффициентов
	for x := range coppy {
		for y := range coppy[x] {
			for w := range coppy[x][y].W {
				coppy[x][y].W[w] += nerv.NODE[x][w].H * coppy[x][y].delta * n * dsigma(coppy[x][y].H)
			}
		}
	}

}

//очищает неросеть от вычисления
func (nerv *neron_cash) Clear() {
	coppy := nerv.NODE
	for x := range coppy {
		for y := range coppy[x] {
			coppy[x][y].H = 0
			coppy[x][y].delta = 0
		}
	}
}
