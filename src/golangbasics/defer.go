package golangbasics

import "fmt"

/*
[ENG] Explanation
Advantages:
1 - Regardless of the type of transactions we have opened
We can add the closing task with defer right to the bottom line.
Thus, it becomes difficult for us to forget.

2 - Defer functions are used at runtime called runtime panic.
It will even work with errors.

3 - Defer is particularly useful for resource release.

Functions marked as go defer (including anonymous functions)
it will definitely run after the main function has finished running.
Why do we need the defer method? If you ask, very different answers can be given.
For example, an opened database link, an open file we are working on, etc.
To turn it off. The most important feature I should mention about Defer is
and it will run even on errors that cause the function to stop.
*/

/*
[TR] Explanation
Avantajları:
1 - Türü ne olursa olsun açılışını yaptığımız işlemlerine
hemen alt satırına defer ile kapatma görevini ekleyebiliriz.
Böylece unutmamız zorlaşır.

2 - Defer fonksiyonları runtime panic adını verdiğimiz çalışma anında ortaya çıkan
hatalarda dahi çalışacaktır.

3 - Defer özellikle kaynakların serbest bırakılması için çok işe yarar bir özelliktir.

Go defer olarak işaretlenmiş fonksiyonları (anonim fonksiyonlar da dahildir)
ana fonksiyonun çalışması sona erdikten sonra mutlaka çalıştırır.
Neden defer metoduna ihtiyaç duyarız? diye soracak olursanız çok farklı cevaplar verilebilir.
Örneğin açılmış bir veritabanı bağlantısını, üzerinde çalıştığımız açık bir dosyayı vb.
Kapatmak için. Defer ile ilgili belirtmem gereken en önemli özelliği çalışma anında ortaya çıkan
ve fonksiyonun durmasına neden olan hatalarda bile çalışacak olmasıdır.

*/

func Defer() {
	BasicDefer()
	fmt.Println()
	MultipleDefer()
	fmt.Println()

	// ANONYMUS DEFER FUNCTION

	defer func(age int) {
		fmt.Println("I am", age, "years old.")
	}(18)
}

// Basic Defer

func BasicDefer() {
	defer sayGoodbye()
	fmt.Println("My name is Baris")
}

func sayGoodbye() {
	fmt.Println("Goodbye!")
}

// Multiple Defer
// First In, Last Out Rule

func MultipleDefer() {
	defer sayTitle()
	defer sayFarewell()
	fmt.Println("My name is Baris")
}

func sayFarewell() {
	fmt.Println("See ya' later!")
}

func sayTitle() {
	fmt.Println("I'm a Nuclear Energy Engineer!")
}
