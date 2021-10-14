package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Sprintln("Type convert")
	var a int = 88
	fmt.Printf("%v,%T\n", a, a)
	//từ int =>float64
	var b float64 = float64(a)
	fmt.Printf("%v,%T\n", b, b)
	//từ flaot64 => int
	var c int = int(b)
	fmt.Printf("%v,%T\n", c, c)
	//từ int =>string
	var d string = string(c)
	fmt.Printf("%v,%T\n", d, d)
	//từ String => int
	k, _ := strconv.Atoi("-55")
	fmt.Printf("%v,%T\n", k, k)
	//từ int => string
	s := strconv.Itoa(42)
	fmt.Printf("%v,%T\n", s, s)

	/* hệ nhị phân có các phép toán &,|,^,~,<<,>>
	1/AND (&)
	0 and 0 = 0
	1 and 1 = 1
	0 and 1 = 0
	1 and 0 = 0
	Như vậy chỉ khi nào 2 bit đều là 1 thì kết quả trả về mới là 1, các trường hợp còn lại đều là 0.
	*/
	and := 6 & 7
	fmt.Printf("%v,%T\n", and, and)
	/*
		2/OR (&)
		0 or 0 = 0
		1 or 1 = 1
		0 or 1 = 1
		1 or 0 = 1
		Như vậy chỉ cần 1 trong 2 bit là 1 thì kết quả trả về sẽ là 1.
	*/
	or := 6 | 9
	fmt.Printf("%v,%T\n", or, or)
	/*
		3/ XOR (^)
		Giả sử ta có 2 bit 0 và 1 thì:
		0 xor 0 = 0
		1 xor 1 = 0
		0 xor 1 = 1
		1 xor 0 = 1
		Như vậy nếu 2 bit khác nhau sẽ cho ra kết quả 1 và ngược lại, 2 bit giống nhau sẽ cho ra kết quả 0
	*/
	xor := 2 ^ 3
	fmt.Printf("%v,%T\n", xor, xor)
	/*
		4/ NOT (~)
		Phép toán not thay đổi bit 0 thành bit 1 và ngược lại, bit 1 thành bit 0.
		Tức là:
		not 0 = 1
		not 1 = 0
		Phép toán này còn được gọi là phép đảo bit.
	*/
	/* địch chuyển << >>
	Trong dịch chuyển số học, các bit được dịch chuyển ra khỏi đầu hoặc đuôi sẽ bị loại bỏ.
	Trong phép dịch chuyển số học về bên trái, các số 0 được dịch chuyển vào bên phải;
	trong phép dịch chuyển số học bên phải, bit thể hiện dấu được thêm vào bên trái, do đó dấu của số được giữ nguyên.
	*/
	var t, i uint
	t, i = 1, 1

	for i = 1; i < 10; i++ {
		fmt.Printf("%d << %d = %d \n", t, i, t<<i)
	}

	fmt.Println()

	t = 512
	for i = 1; i < 10; i++ {
		fmt.Printf("%d >> %d = %d \n", t, i, t>>i)
	}
}
