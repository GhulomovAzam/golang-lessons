package main

func main() {
/*Итак, у каждой переменной есть:
1. Имя
2. Значение
3. Тип
4. Время жизни*/
	var (
		priceApple float64 = 12.4
		massApple float64 = 1.0
		result = priceApple * massApple
		)
		println(result)
/*Целые числа
● int - 4 или 8 байт
● int8 – 1 байт (-128 до 127)
● int16 – 2 байта (-32768 до 32767)
● int32 – 4 байта (-2147483648 до 2147483647)
● int64 – 8 байт (-9223372036854775808 до 9223372036854775807)*/
		var (
			pricePotato int8 = 10_1
			massPotato int8 = 1_1 
			)
			println(pricePotato % massPotato)
			println("Hellow World!" )
			println(`Привееет 
			
			
			
			
			аааааааааааааааааа`)
}