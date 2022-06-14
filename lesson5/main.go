// Одне яблуко коштує 5.99 грн. Ціна однієї груші - 7 грн.
// Ми маємо 23 грн.
// 1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?
// 2. Скільки груш ми можемо купити?
// 3. Скільки яблук ми можемо купити?
// 4. Чи ми можемо купити 2 груші та 2 яблука?
//
// Задача:
// Опишіть вирішення всіх пунктів задачі використовуючи необхідні змінні чи/та константи.
// Під опишіть, я маю на увазі наступне:
// Я маю збілдити ваш код і запустити програму. В результаті цього, я маю побачити роздруковані // відповіді на поставлені вище запитання. Перед відповідю треба роздрукувати саме питання.
// Правильно обирайте типи даних та назви змінних чи констант.
// Публікація:
// Створити папку в своєму репозиторії в github та завантажити туди main.go файл, в котрому буде зроблено дане завдання.

package main

import "fmt"

var (
	thePriceOfAnApple float64 = 5.99
	thePriceOfaPear   float64 = 7
	myWholeMoney      float64 = 23
	iNeedMoney        float64
)

func main() {
	iNeedMoney := thePriceOfAnApple*9 + thePriceOfaPear*8
	fmt.Println("1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?")
	fmt.Println("   Потрібно витратити", iNeedMoney, "грн., щоб купити 9 яблук та 8 груш")
	fmt.Println("")
	fmt.Println("2. Скільки груш ми можемо купити?")
	fmt.Println("   Ми можемо купити", int64(myWholeMoney/thePriceOfaPear), "груші за 23 грн.")
	fmt.Println("")
	fmt.Println("3. Скільки яблук ми можемо купити?")
	fmt.Println("   Ми можемо купити", int64(myWholeMoney/thePriceOfAnApple), "яблука за 23 грн.")
	fmt.Println("")
	fmt.Println("4. Чи ми можемо купити 2 груші та 2 яблука?")
	costOfTwoApplesAndTwoPear := thePriceOfAnApple*2 + thePriceOfaPear*2
	fmt.Println("   Ми", costOfTwoApplesAndTwoPear < myWholeMoney, "купити 2 груші та 2 яблука")

}
