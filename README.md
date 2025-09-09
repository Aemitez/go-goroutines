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

## 📌 ตัวอย่างโค้ด

### แบบปกติ (Sequential)
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	fmt.Println("Start working...")
	time.Sleep(2 * time.Second)
	fmt.Println("Work done!")

	fmt.Printf("Finished in %s\n", time.Since(start))
}
