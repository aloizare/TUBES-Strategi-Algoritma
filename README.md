TUGAS BESAR STRATEGI ALGORITMA 2024
=============
_Analisis Performa Algoritma Greedy dan Dynamic Programming untuk Problem “House Robber”_
### Penjelasan Singkat Tentang Algoritma yang Diimplementasikan
Algoritma Greedy merupakan pendekatan sederhana yang mengambil keputusan berdasarkan pilihan terbaik saat ini tanpa mempertimbangkan konsekuensi jangka panjang. Di sisi lain, algoritma Dynamic Programming menggunakan pendekatan yang lebih efisien dengan memanfaatkan submasalah yang saling terkait, menyimpan informasi yang relevan untuk mempercepat proses pengambilan keputusan.
### Requirement
`Golang` dan `Python3`.
IDE bisa disesuaikan seperti `Visual Studio Code` atau compiler lainnya.
### Implementasi Algoritma
Greedy
```go
func rob(nums []int) int {
	inc := nums[0]
	exc := 0
	for i := 1; i < len(nums); i++ {
		ninc := exc + nums[i]
		nexc := max(inc, exc)
		inc = ninc
		exc = nexc
	}
	return max(inc, exc)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

Dynamic Programming
```go
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	prev2 := 0
	prev1 := 0
	for _, num := range nums {
		tmp := prev1
		prev1 = max(prev2+num, prev1)
		prev2 = tmp
	}
	return prev1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

### Analisis
Pengujian akan dilakukan terhadap parameter (a) running time dan (b) kompleksitas waktu
Skema pengujian dilakukan sebanyak 4 kondisi dengan masing-masing skema menggunakan jumlah rumah yang berbeda mulai dari 5, 100, 500, dan 1000 rumah. Berikut adalah tabel 
#### Running Time
Analisis _running time_ (waktu eksekusi) memberikan informasi tentang efisiensi algoritma dalam menyelesaikan permasalahan “House Robber”. Hasil penelitian menunjukkan bahwa algoritma Dynamic Programming secara konsisten memiliki running time yang lebih rendah dibandingkan dengan algoritma Greedy walaupun dengan selisih yang sangat tipis, hal ini mungkin dikarenakan bahasa Go diketahui memiliki kinerja yang sangat baik, baik dalam hal kecepatan maupun penggunaan memori
Hasil pengujian berupa waktu eksekusi dalam milidetik (ms) sebagai berikut
| Jumlah rumah  | Greedy        | Dynamic Programming |
| ------------- | ------------- | ------------------- |
| 5 | 0.0002 | 0.00019 |
| 100  | 0.00037 | 0.00022 |
| 500  | 0.00087  | 0.00047 |
| 1000  | 0.00143  | 0.00073 |

Dapat juga dilihat pada tabel banyaknya rumah (sumbu X) dengan waktu eksekusi (sumbu Y)
![image](https://github.com/aloizare/TUBES---Strategi-Algoritma/assets/168240032/5e2f359a-d3f4-4bf7-95d6-1dd124a6a6c6)

#### Kompleksitas Waktu
Algoritma Greedy dan Dynamic Programming keduanya melakukan perulangan sebanyak jumlah rumah (n), dengan setiap iterasi melakukan operasi aritmatika sederhana yang memiliki kompleksitas waktu O(1). Oleh karena itu, kompleksitas waktu total untuk kedua algoritma adalah O(n).

Meskipun kompleksitas waktu keduanya sama, dalam implementasi nyata, Dynamic Programming cenderung lebih efisien daripada Greedy.

### Author/Identitas Pembuat
| No. | Nama | NIM |
| --- | ---- | --- |
| 1 | M. Raditya Faturrahman | 1304221050 |
| 2 | Alif Lohen | 1304212117 |
