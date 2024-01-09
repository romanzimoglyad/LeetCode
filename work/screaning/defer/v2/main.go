/**
package main

import "fmt"

func deferLogic(i int) (result int) {
	defer func() {
		result = result + 7
	}()
	defer func() {
		result = result * 2

	}()
	result = i * 5
	return 20

}

func main() {
	result := deferLogic(5)
	fmt.Println(result) //вопрос: каков порядок работы и что в итоге?
}

*/

package main

import "fmt"

func deferLogic(i int) (result int) {
	defer func() {
		result = result + 7 //ответ (3)
	}()
	defer func() {
		result = result * 2 //ответ (2)
	}()
	result = i * 5 // ответ дропаем это
	return 20      //ответ (1)
}

func main() {
	result := deferLogic(5)
	fmt.Println(result) //вопрос: каков порядок работы и что в итоге?
	//ответ: порядок в функции описан, ответ 47
}
