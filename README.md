# GO-JTI
![App](/public/images/screenshot/app.png) 

go-jti adalah hasil Teknikal Tes untuk Programmer Backend Developer Golang

IMPLEMENTASI
------------

- [x] Gmail (Sign In With Google)
- [x] Metode ENCRYPT/DECRYPT
- [ ] WEB SOCKET
- [ ] POSTGRE SQL 

INTRODUCTION
------------
* [LINK GITHUB](https://www.github.com/absormu/go-jti)
* DOCUMENTATION API postman collection (folder /docs/go-jti.postman_collection.json)
* REQUEST & RESPONSE API (folder /docs/Input.md & Output.md)
* DATABASE : create database db_gojti, (folder /scripts/db_gojti.sql)

INSTALLATION, CONFIGURATION AND RUNING
------------

 * Import database from folder /scripts/db_gojti.sql
 * Config from /pkg/configuration/config.go
 * go build
 * ./go-jti or ./go-jti.exe
 * Import postman collection FROM /docs/go-jti.postman_collection
  
 TESTING
------------
* http://localhost:9670/ & sign google account


## DESIGN FLOW
![Design Flow](/public/images/Design%20Flow.png) 

## TUGAS 
* Data Base Menggunakan MYSQL Gunakan POSTGRE Menjadi Nilai Tambah.
* Buat Form Input.
    * Penginputan Data Bisa Menggunakan Data PLAIN.
    Menggunakan Metode ENCRYPT/DECRYPT Menjadi Nilai Tambah.
    Tombol Save.
    Hanya menyimpan satu input dari isian.
    Tombol Auto.
    Saat Auto ditekan, maka akan auto generate nomor hp (25 generate) dan auto input.
* Buat Form Output.
    * Untuk Form Output, Data Di Bagi 2 :
Data Yang Habis Di Bagi 2.
Data Yang Tidak Habis Di Bagi 2.
  * Untuk Menampilkan Update Data Di Form Output bisa Menggunakan Metode REFRESH TIMER
Menggunakan WEB SOCKET Menjadi Nilai Tambah.
  * Untuk Edit Dan Delete Bisa Edit No Hp Dan Bisa Hapus. Tambahan Di Form Output Untuk Edit Dan Delete Silahkan Berkreasi Sendiri.
* Buat Restfull API Untuk mengakomodir data dari Form Input & Output.
* Sebelum Masuk Form Input Atau Form Output Ada Form Login Menggunakan Gmail (Sign In With Google).
* Enkripsi Menggunakan openssl, aes-256-cbc. Dan Memakai IV Dan KEY Sendiri.
* Penulisan Code / Script Dibuat Secara Rapi Dan Terstruktur.
* Fungsi Global, Usahakan Buat Class Sendiri ( OOP ).
* Khusus Framework Yang Membutuhkan Pihak Ke 3 Seperti LARAVEL/SYMPHONY Automatic Installation Harus Bisa Berjalan ( Seperti COMPOSSER / LOADER ) Kebutuhan Environment, Config, ID, dll harus sudah disediakan (Khusus DBname config, kami sesuaikan). Cara Running (CMD apa saja yang dibutuhkan harus disertakan lengkap). Team penguji, hanya akan menjalankan sesuai arahan Info tertulis di Github.
* Waktu Pengerjaan 5 hari untuk Fresh Graduate, 2 hari untuk expert.
Kurang Dari Waktu Tersebut Menjadi Nilai Tambah.

## SCREENSHOT
### Menu Awal
Menu             |  Awal
:-------------------------:|:-------------------------:
![](/public/images/screenshot/1.png)  |  ![](/public/images/screenshot/2.png)
![](/public/images/screenshot/3.png)  |  ![](/public/images/screenshot/4.png)

### Menu Input
Menu             |  Awal
:-------------------------:|:-------------------------:
![](/public/images/screenshot/5a.png)  |  ![](/public/images/screenshot/5b.png)
![](/public/images/screenshot/5c.png)  |  ![](/public/images/screenshot/7.png)

### Menu Output
![Design Flow](/public/images/screenshot/6.png) 

### Author
```
Muhamad Ulil Absor
absoralvord07@gmail.com
https://www.linkedin.com/in/absormu
```
