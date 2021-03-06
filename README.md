#### Неросеть на базе стандартных библиотек

Мы будем использовать сигмоиду гиперболический тангенc и его диапазон от [-1;1]
![Tanh](/image/functions.png)

```go 
import ("math") 

func main () {
    // Вот такой вид имеет вызов из стандартной 
    // библиотеки гиперболический тангенc в go
    y := math.Tanh(x)
}
```
## Создаем неросеть

```go 
import (
    "math"
    nerv "github.com/Xoxko/goNeron/neron"
    ) 
github.com/Xoxko/goNeron/neron
func main () {
    neron := nerv.Neron
}
```
Для того чтобы создать слои нужно вписать в слайс массив чиссел
sub := []int{2, 10, 20, 1}

Где {2, 10, 20, 1}
* 2 -> это количество входных неронов.
* 10 -> это количество  неронов первого внутренего слоя.
* 20 -> это количество  неронов второго внутренего слоя.
* 1 -> это количество выходных неронов.

```go 
func main () {
    //Создаем неросеть
    neron := nerv.Neron
    //создаем слои
    sub := []int{2, 10, 20, 1}
    //Функция создающая слои неронов
    input, output := neron.NewInitNode(sub)
    //Отдаем функцию для рандомизирования синапсов
    rand.Seed(time.Now().UnixMicro())
    neron.Random(func() float64 { return 1.0 - (rand.Float64() * 2) })
    //Неросеть готова для обучения
}
```
для того чтобы обучить

```go 
func main () {
    neron := nerv.Neron
    sub := []int{2, 10, 20, 1}
    input, output := neron.NewInitNode(sub)
    rand.Seed(time.Now().UnixMicro())
    neron.Random(func() float64 { return 1.0 - (rand.Float64() * 2) })
    //Неросеть готова для обучения
    for i := 0; i < 100000; i++ {
        // Передаем значения на вход
        *input[0] = //float64
        *input[1] = //float64
        //Вычисляем
        neron.Calculation()
        //Считываем значения для проверки, если нужно
        out = *output[0]
        //Отправляем правильный ответ в виде массива в слайсе
        s[0] = //Правильный ответ
        //если s это просто массив то достаточно поставить двоеточие s[:],
        //и он бедет передаваться в виде слайса
        neron.Training(s[:], t)
        //очищаем весса дельты и вычислений сигмоиды иначе все будет друг с другом складываться
        neron.Clear()
    }
}
```


Если хотите изменить неросеть на Логистическую сигмоиду в файле neron/cash.go

``` go
//сигмоида
func sigma(x float64) float64 {
	return math.Tanh(x) //1.0 / (1.0 + math.Exp(-x))
}

//производная от сигмоиды
func dsigma(x float64) float64 {
	return 1.0 - math.Pow(x, 2) //x * (1.0 - x)
}
```
Замените на 

``` go
//сигмоида
func sigma(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-x))
}

//производная от сигмоиды
func dsigma(x float64) float64 {
	return x * (1.0 - x)
}
```