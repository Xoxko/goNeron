package neron

// init struct dataset
var Neron neron_cash

type neron_cash struct {
	NODE [][]node
}
type node struct {
	W []float64
	H float64
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
	coppy := nerv.NODE[1:][:]
	for x := range coppy {
		for y := range coppy[x] {
			for w := range coppy[x][y].W {
				coppy[x][y].W[w] = number()
			}
		}
	}
}

func (nerv *neron_cash) Calculation() {

}
