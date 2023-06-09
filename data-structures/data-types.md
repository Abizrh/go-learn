Tipe data numerik non-desimal atau non floating point di Go ada beberapa jenis. Secara umum ada 2 tipe data kategori ini yang perlu diketahui.

- uint, tipe data untuk bilangan cacah (bilangan positif).
- int, tipe data untuk bilangan bulat (bilangan negatif dan positif).

Kedua tipe data di atas kemudian dibagi lagi menjadi beberapa jenis, dengan pembagian berdasarkan lebar cakupan nilainya, detailnya bisa dilihat di tabel berikut.


Tipe data	Cakupan bilangan
uint8	0 ↔ 255
uint16	0 ↔ 65535
uint32	0 ↔ 4294967295
uint64	0 ↔ 18446744073709551615
uint	sama dengan uint32 atau uint64 (tergantung nilai)
byte	sama dengan uint8
int8	-128 ↔ 127
int16	-32768 ↔ 32767
int32	-2147483648 ↔ 2147483647
int64	-9223372036854775808 ↔ 9223372036854775807
int	sama dengan int32 atau int64 (tergantung nilai)
rune	sama dengan int32 

Dianjurkan untuk tidak sembarangan dalam menentukan tipe data variabel, sebisa mungkin tipe yang dipilih harus disesuaikan dengan nilainya, karena efeknya adalah ke alokasi memori variabel. Pemilihan tipe data yang tepat akan membuat pemakaian memori lebih optimal, tidak berlebihan.