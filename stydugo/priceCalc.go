package stydugo

import "fmt"

func PriceCalc() {
	fmt.Println("Калькулятор стоимости товара")
	fmt.Println("Введите стоимость товара:")
	var cost int
	fmt.Scan(&cost)

	fmt.Println("Введите стоимость доставки:")
	var deliveryCost int
	fmt.Scan(&deliveryCost)

	fmt.Println("Введите скидку:")
	var discount int
	fmt.Scan(&discount)

	sum := (cost + deliveryCost) - discount

	fmt.Println("Итого:", sum, "рублей")
}
