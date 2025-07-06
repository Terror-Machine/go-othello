# 🧩 go-Othello

**go-Othello** adalah permainan **Othello/Reversi berbasis terminal (CLI)** dengan papan visual yang disimpan sebagai **gambar PNG**. Anda bermain sebagai **Hitam (⚫️)** melawan bot sederhana yang akan otomatis melakukan gerakan.  

> Dibuat oleh: © 2025 HyHy  
> Bahasa: Go (Golang)  
> Output: PNG (menggunakan [fogleman/gg](https://github.com/fogleman/gg)) 

---

## 🎯 Fitur

✅ Bermain Othello langsung dari terminal  
✅ Papan otomatis diperbarui dalam file `othello.png`  
✅ Highlight **langkah valid** dengan tanda lingkaran  
✅ Bot pintar memilih gerakan dengan flips terbanyak  
✅ Hitung skor otomatis saat game berakhir  

---

## 📥 Cara Instalasi

### 1️⃣ Clone Repository
```bash
git clone https://github.com/Terror-Machine/go-othello
cd go-othello
````

### 2️⃣ Install Dependensi

Pastikan Go (Golang) sudah terpasang. Jalankan:

```bash
go mod tidy
```

💡 **Cek Go:**

```bash
go version
```

Kalau belum ada, download di [https://go.dev/dl](https://go.dev/dl).

## 🚀 Cara Menjalankan

Jalankan program dengan:

```bash
go run main.go
```

Setelah itu, Anda akan melihat instruksi seperti ini:

```
Selamat datang di Game Othello CLI!
Papan Othello telah dibuat dan disimpan sebagai 'othello.png'.
Anda bermain sebagai Hitam (⚫️). Giliran Anda untuk bergerak.

Perintah:
  d3        -> Meletakkan keping di D3.
  pass      -> Melewati giliran (jika tidak ada langkah valid).
  new       -> Memulai game baru.
  exit      -> Keluar dari permainan.
```

Setiap gerakan akan memperbarui file **`othello.png`**.

---

## 🎮 Cara Bermain

🟤 Anda bermain sebagai **Hitam (⚫️)**
⚪ Bot bermain sebagai **Putih (⚪️)**

* Masukkan langkah seperti `d3` untuk meletakkan keping di kolom D baris 3.
* Papan akan diperbarui di `othello.png` setelah setiap gerakan.
* Langkah valid Anda ditandai dengan lingkaran putih kecil.
* Bot akan bergerak otomatis setelah giliran Anda.

---

## 📜 Perintah CLI

| Perintah | Fungsi                                        |
| -------- | --------------------------------------------- |
| `d3`     | Meletakkan keping di kolom D baris 3          |
| `pass`   | Melewati giliran jika tidak ada langkah valid |
| `new`    | Memulai permainan baru                        |
| `exit`   | Keluar dari permainan                         |

---

## 🖼️ Contoh Papan

Contoh papan awal:

| ![image](https://github.com/user-attachments/assets/6a28334b-108d-4fee-897c-d027205f1485) |
| ----------------------------------------------------------------------------------------- |

---

## 💡 Tips

* Buka **`othello.png`** di image viewer yang mendukung auto-refresh supaya lebih praktis.
* Gunakan perintah berikut (Linux/macOS):

  ```bash
  watch -n 1 feh othello.png
  ```

---

## 📄 Lisensi

MIT License © 2025 HyHy

---

## 👨‍💻 Kontribusi

Pull request, issue, dan ide pengembangan sangat disambut!
