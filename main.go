package main

import (
	"fmt"
)

// หา speedup S ด้วย Amdahl's Law
func amdahlsLaw(P float64, N int) float64 {
	// Speedup = 1 / ((1 - P) + (P / N))
	// ของ Amdahl มองที่งานทั้งหมดเป็นก้อน คือ 1  ตอนแรกปกติงาน 1
	// พอมี parallelเข้ามา ทำให้งานส่วนPลดงานลงNเท่า
	speedup := 1.0 / ((1.0 - P) + (P / float64(N)))
	return speedup
}

// หา speedup S ด้วย Gustafson's Law
func gustafsonsLaw(P float64, N int) float64 {
	// Speedup = (P*N) + (1-P)
	// S คือการเพิ่มประสิทธิภาพ (Speedup)
	// P Parallelizable fraction คือส่วนของโปรแกรมที่สามารถทำงานขนานได้
	//	 ค่า P ระหว่าง 0 ถึง 1
	// N คือจำนวน processor
	// ส่วน(P*N)ที่โปรแกรมได้ประสิทธิภาพเต็มที่  ในขณะที่ 1-P คือไม่สามารถทำงานแบบparallel
	S := P*float64(N) + (1 - P)

	return S
}

func main() {
	// e.g.
	P := 0.9 // 90% of the program can be parallelized
	// 90%นี้ของโปรแกรมจึงรับ speedup เต็มที่เลย
	N := 8 // Number of processors

	speedupA := amdahlsLaw(P, N)
	speedupG := gustafsonsLaw(P, N) //คำนวณความเร็วที่เพิ่มขึ้นหลัง parallel

	fmt.Printf("The speedup S using Amdahl's Law is: %.2f\n", speedupA)
	fmt.Printf("The speedup S using Gustafson's Law is: %.2f\n", speedupG)
}

//สรุปคือ
//Amdahl มองเป็นอัตราส่วนบนล่าง ระหว่างงานตอนแรก100% หารกับ งานหลังparallelที่ลดลง
//Gustafson แบ่งแยกงาน 2 ส่วน คือ ส่วนของโปรแกรมที่ speedupได้ กับ ส่วนของโปรแกรมที่ speedup ไม่ได้
