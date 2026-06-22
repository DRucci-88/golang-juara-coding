# Plan Implementasi `mini-hris`

## Tujuan
Membangun backend Mini HRIS menggunakan Go, Gin, dan SQLite untuk menangani data master HR, manajemen karyawan, pencatatan kehadiran, pengajuan cuti, dan proses payroll bulanan dengan relasi database yang benar.

## Teknologi dan Library yang Digunakan

### Teknologi Utama
- `Go`
  - Bahasa pemrograman utama untuk backend API.
- `SQLite`
  - Database lokal ringan untuk penyimpanan data aplikasi.

### Library Go Utama
- `github.com/gin-gonic/gin`
  - Framework HTTP untuk routing, JSON response, binding request, dan validasi input.
- `gorm.io/gorm`
  - ORM untuk query database, relasi antar tabel, dan auto migration.
- `gorm.io/driver/sqlite`
  - Driver SQLite untuk GORM.

### Library/Paket Standard yang Dipakai
- `net/http`
  - Menyediakan konstanta status HTTP seperti `http.StatusOK`, `http.StatusBadRequest`, dan `http.StatusCreated`.
- `time`
  - Membantu parsing dan validasi format tanggal, jam, dan periode payroll.
- `fmt`
  - Digunakan bila perlu untuk formatting pesan atau debug ringan.
- `strings`
  - Membantu pencarian dan normalisasi input filter.
- `strconv`
  - Membantu parsing parameter path atau query menjadi integer.

### Fitur Framework yang Akan Dimanfaatkan
- Binding JSON dari Gin dengan `ShouldBindJSON`.
- Validasi input menggunakan tag `binding`, seperti:
  - `binding:"required"`
  - `binding:"required,email"`
  - `binding:"oneof=ACTIVE SUSPENDED TERMINATED"`
- Preload relasi GORM untuk menampilkan data employee beserta department dan position.
- Auto migration GORM untuk membentuk schema database secara otomatis.

### Dependensi Minimum Saat Inisialisasi Project
- `go get github.com/gin-gonic/gin`
- `go get gorm.io/gorm`
- `go get gorm.io/driver/sqlite`

## Struktur Folder

```text
mini-hris/
+-- config/
|   +-- db.go             # Konfigurasi koneksi dan inisialisasi SQLite
|   +-- seed.go           # Seeder data awal departments, positions, employees
+-- controllers/
|   +-- department.go     # Handler CRUD departemen
|   +-- position.go       # Handler CRUD jabatan
|   +-- employee.go       # Handler manajemen karyawan relasional
|   +-- attendance.go     # Handler pencatatan kehadiran
|   +-- leave.go          # Handler pengajuan dan approval cuti
|   +-- salary.go         # Handler kalkulasi dan list payroll
+-- models/
|   +-- department.go     # Struct entitas department
|   +-- position.go       # Struct entitas position
|   +-- employee.go       # Struct entitas employee + relasi
|   +-- attendance.go     # Struct entitas attendance
|   +-- leave.go          # Struct entitas leave
|   +-- salary.go         # Struct entitas salary
+-- routes/
|   +-- routes.go         # Registrasi seluruh endpoint Gin
+-- go.mod
+-- go.sum
+-- main.go               # Entry point server
```

## Tanggung Jawab Tiap Bagian

### 1. `config/db.go`
- Membuka koneksi SQLite.
- Menyimpan instance database agar bisa dipakai controller.
- Menjalankan auto migration untuk semua tabel:
  - `departments`
  - `positions`
  - `employees`
  - `attendances`
  - `leaves`
  - `salaries`

### 2. `config/seed.go`
- Menyediakan data awal untuk:
  - beberapa `departments`
  - beberapa `positions`
  - beberapa `employees`
- Seeder hanya menambahkan data bila tabel masih kosong.

### 3. `models/`
- Mendefinisikan seluruh struct entitas dan relasi foreign key.
- Menambahkan tag `json`, `gorm`, dan `binding` sesuai kebutuhan.
- Menggunakan relasi pada `Employee` ke `Department` dan `Position`.

### 4. `controllers/`
- Menangani request, validasi input, query database, dan response JSON.
- Memisahkan logika per domain agar setiap file fokus pada satu resource.

### 5. `routes/routes.go`
- Mendaftarkan semua endpoint di bawah prefix `/api`.
- Menghubungkan route ke handler controller yang sesuai.

### 6. `main.go`
- Inisialisasi database.
- Menjalankan migration dan seeding.
- Membuat router Gin.
- Memanggil registrasi route.
- Menjalankan server, misalnya di `:8080`.

## Desain Entitas Database

### 1. `Department`
- `ID`
- `Name`
- `Code`

### 2. `Position`
- `ID`
- `Title`
- `BaseSalary`

### 3. `Employee`
- `ID`
- `NIK`
- `FullName`
- `Email`
- `DepartmentID`
- `PositionID`
- `Status`

### 4. `Attendance`
- `ID`
- `EmployeeID`
- `Date`
- `CheckIn`
- `CheckOut`
- `Status`

### 5. `Leave`
- `ID`
- `EmployeeID`
- `StartDate`
- `EndDate`
- `Reason`
- `Status`

### 6. `Salary`
- `ID`
- `EmployeeID`
- `Period`
- `BasicSalary`
- `Allowance`
- `Deductions`
- `NetSalary`

## Aturan Relasi
- Satu `Department` bisa memiliki banyak `Employee`.
- Satu `Position` bisa memiliki banyak `Employee`.
- Satu `Employee` bisa memiliki banyak `Attendance`.
- Satu `Employee` bisa memiliki banyak `Leave`.
- Satu `Employee` bisa memiliki banyak `Salary`.

## Endpoint Wajib

### A. Master Data

#### Departments
- `POST /api/departments`
- `GET /api/departments`
- `PUT /api/departments/:id`
- `DELETE /api/departments/:id`

#### Positions
- `POST /api/positions`
- `GET /api/positions`
- `PUT /api/positions/:id`
- `DELETE /api/positions/:id`

### B. Manajemen Karyawan
- `POST /api/employees`
  - Validasi `DepartmentID` dan `PositionID` harus ada.
- `GET /api/employees`
  - Menampilkan detail employee beserta department dan position.
  - Mendukung query:
    - `?search=nama`
    - `?department_id=1`
    - `?position_id=2`
    - `?status=ACTIVE`
- `PUT /api/employees/:id`

### C. Kehadiran dan Cuti
- `POST /api/attendances`
- `POST /api/leaves`
  - Status awal wajib `PENDING`.
- `PATCH /api/leaves/:id/approve`
  - Mengubah status menjadi `APPROVED` atau `REJECTED`.

### D. Payroll
- `POST /api/salaries/calculate`
  - Menghitung dan menyimpan payroll periode tertentu.
- `GET /api/salaries/period/:period`
  - Menampilkan slip gaji semua karyawan untuk periode tertentu.

## Validasi Minimum

### Department
- `Name` wajib diisi.
- `Code` wajib diisi dan unik.

### Position
- `Title` wajib diisi.
- `BaseSalary` wajib diisi dan lebih dari 0.

### Employee
- `NIK` wajib diisi dan unik.
- `FullName` wajib diisi.
- `Email` wajib diisi dan harus format email valid.
- `DepartmentID` wajib diisi.
- `PositionID` wajib diisi.
- `Status` hanya boleh:
  - `ACTIVE`
  - `SUSPENDED`
  - `TERMINATED`

### Attendance
- `EmployeeID` wajib diisi.
- `Date` wajib format `YYYY-MM-DD`.
- `CheckIn` wajib format `HH:MM`.
- `CheckOut` wajib format `HH:MM`.
- `Status` hanya boleh:
  - `PRESENT`
  - `LATE`
  - `ABSENT`

### Leave
- `EmployeeID` wajib diisi.
- `StartDate` wajib diisi.
- `EndDate` wajib diisi.
- `Reason` wajib diisi.
- `Status` hanya boleh:
  - `PENDING`
  - `APPROVED`
  - `REJECTED`

### Salary Calculation Request
- Period payroll wajib dikirim, format `YYYY-MM`.
- Employee target bisa dihitung per seluruh karyawan aktif atau berdasarkan parameter yang ditentukan saat implementasi.

## Aturan Logika Payroll
- `BasicSalary` diambil dari `Position.BaseSalary`.
- `Allowance` dihitung dari kehadiran tepat waktu.
  - Contoh aturan dari soal: bonus `50000` per hari hadir tepat waktu.
- `Deductions` dihitung dari keterlambatan dan absen.
  - Contoh aturan dari soal:
    - `20000` per hari `LATE`
    - `100000` per hari `ABSENT`
- `NetSalary` dihitung dengan rumus:
  - `BasicSalary + Allowance - Deductions`
- Cuti dengan status `APPROVED` ikut diperhitungkan saat proses payroll sesuai ketentuan bisnis yang diterapkan.

## Strategi Implementasi
1. Inisialisasi project Go dan pasang dependency `gin`, `gorm`, dan driver `sqlite`.
2. Buat seluruh model beserta relasi dan enum status yang dipakai.
3. Implementasikan `config/db.go` untuk koneksi database dan auto migration.
4. Implementasikan `config/seed.go` untuk data awal.
5. Buat controller master data: departments dan positions.
6. Buat controller employee dengan validasi foreign key dan query filter.
7. Buat controller attendance dan leave.
8. Buat controller salary untuk kalkulasi payroll dan list payroll per periode.
9. Registrasikan seluruh route di `routes/routes.go`.
10. Jalankan pengujian endpoint dengan Postman, Insomnia, atau Thunder Client.

## Output Akhir yang Diharapkan
- Struktur project rapi sesuai ketentuan tugas.
- Database SQLite terbentuk otomatis.
- Minimal 6 tabel relasional berhasil dibuat.
- Seeder berjalan dan API bisa langsung diuji.
- Seluruh endpoint wajib tersedia dan mengembalikan JSON.
- Validasi input mengembalikan `400 Bad Request` bila request tidak valid.
- Data employee bisa ditampilkan lengkap dengan department dan position.
- Payroll periode bulanan dapat dihitung dan disimpan ke tabel `salaries`.
