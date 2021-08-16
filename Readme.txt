Tes - Article - FS - Muhammad Imam Maulana
------------------------------------------

Versi Go = go1.16.6 windows/amd64
Versi React = 17.0.2
Versi Node = 14.17.4
FrameWork Go = Echo
Framework React = CoreUI
Platform API = Postman
Link Frontend = https://github.com/mimamm/tes-sv-imam-FE.git
Link Backend = https://github.com/mimamm/tes-sv-imam-BE.git
postman collection = (Direktori BE)/Article.postman_collection.json
Database = Generate dari DB melalui run Backend dengan request di postman collection
File SQL Database = (Direktori BE)/article SQL (jika database ingin import sql)
Database name = article
Database username = root
Database password = "mymysql" (bisa diganti sesuai password MySQL user)


How to use:
-----------
Lakukan Git Clone pada kedua repositori Git:
- BE: https://github.com/mimamm/tes-sv-imam-BE.git
- FE: https://github.com/mimamm/tes-sv-imam-FE.git

Untuk mengganti password database aplikasi supaya sesuai dengan password MySQL dari user dapat dilakukan pada direktori:
	tes-sv-imam-BE/driver/connect.go 

Run terminal pada folder BE (tes-sv-imam-BE) jalankan "go run main.go"

Import Postman Collection yang berada di direktori BE dengan nama "Article.postman_collection"
	(sudah tersedia body & parameternya)

Jalankan Request "Generate Database & Table" untuk membuat Database & Table
Jalankan Request "Insert Example Dummy Data" untuk memasukkan data dummy



Run Application (Menggunakan UI Frontend)
-----------------------------------------

Run terminal pada folder FE (tes-sv-imam-FE) jalankan "npm install"
setelah selesai, jalankan "npm start"

Aplikasi melalui Frontend siap digunakan dengan menu-menu pada sidebar



Run Application (Backend - Melalui Request Postman)
---------------------------------------------------

1. Membuat Database & Table
postman request: Generate Database & Table
fungsi: membuat database dan tablenya jika belum ada

2. Input Data-data Dummy
postman request: Insert Example Dummy Data
fungsi: memasukkan data-data dummy untuk testing

3. Create Article
postman request: Create Article
fungsi: membuat article baru

4. Menampilkan data menggunakan Pagination
postman request: Read with Pagination
fungsi: get data yang menerapkan pagination

5. Get 1 article by Id
postman request: Get Article by Id
fungsi: mendapatkan 1 article berdasarkan id yang dimasukkan

6. Update Article
postman request: Update Article
fungsi: melakukan update article berdasarkan id

7. Delete Article
postman request: Delete Article
fungsi: melakukan proses delete article (mengubah status menjadi Thrash)
