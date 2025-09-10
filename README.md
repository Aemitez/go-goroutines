# 🚀 ทำความรู้จักกับ Goroutine ในภาษา Go

## ✨ Goroutine คืออะไร?
- Goroutine คือ **lightweight thread** ของภาษา Go  
- ใช้สำหรับรันงานหลายอย่างพร้อมกัน (**Concurrency**)  
- สร้างได้เป็นแสน–ล้าน goroutines โดยใช้ memory และ resource น้อย  

---

## ⚡ ทำไมต้องใช้ Goroutine?
- เขียนง่าย → แค่เติม `go` หน้า function
- เบากว่า Thread ปกติ
- มี Go runtime scheduler จัดการการทำงานอัตโนมัติ
- เหมาะกับงาน **I/O-bound** (เช่น call API, DB, Kafka)  
- สามารถใช้กับ **Parallel processing** เมื่อมีหลาย CPU cores  

---

## 📌 ตัวอย่างโค้ด 1 : No Goroutine (Sequential) vs  Goroutine
```go

func main() {

	// TODO example :No Goroutine (Sequential) vs  Goroutine
	exampleCompare("Normal")

	go exampleCompare("Power GO")

	time.Sleep(time.Millisecond)
}
func exampleCompare(state string) {
	start := time.Now()

	// ทำงานปกติ (รอ 2 วิ)
	fmt.Printf("%s: Start working...\n", state)
	time.Sleep(2 * time.Nanosecond)
	fmt.Printf("%s: Work done!\n", state)

	elapsed := time.Since(start)
	fmt.Printf("%s:Finished in %s\n", state, elapsed)
	fmt.Println("------------------------")
}
```

## 📌 ตัวอย่างโค้ด 2 : Goroutine with chanel
```go

func main() {

	// TODO example :2 Goroutine with chanel
	ch := make(chan int)

	go producer(ch) //  1 Goroutine
	go consumer(ch) //  2 Goroutine

	time.Sleep(time.Second)
	fmt.Println("done")
}
func producer(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Producing: %d\n", i)
		ch <- i
		time.Sleep(100 * time.Millisecond)
	}
	close(ch) // Close the channel when done
}

func consumer(ch <-chan int) {
	for num := range ch {
		fmt.Printf("Consuming: %d\n", num)
	}
}
```
