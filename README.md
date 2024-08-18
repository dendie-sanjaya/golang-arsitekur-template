# Den Golang Aristkur Template 

A golang template architecture using fiber framework version 2, the goal of this architecture is to make it easy to create rest api based services using the golang programming language as easy as using php, nodejs or python

## Structur Directory


- **db/migrations**: Direktori ini berisi file-file migrasi SQL untuk memperbarui skema database.
  - `xxxx_add_age_to_users.up.sql`: Skrip untuk menambahkan kolom `Age` ke tabel `users`.
  - `xxxx_add_age_to_users.down.sql`: Skrip untuk menghapus kolom `Age` dari tabel `users`.
  - `xxxx_remove_email_from_users.up.sql`: Skrip untuk menghapus kolom `Email` dari tabel `users`.
  - `xxxx_remove_email_from_users.down.sql`: Skrip untuk menambahkan kembali kolom `Email` ke tabel `users`.

- **db/seed.go**: File ini berisi fungsi untuk mengisi database dengan data awal (seed data).

- **entity/user.go**: File ini berisi definisi entitas `User`.

- **main.go**: File utama untuk menjalankan aplikasi.

- **README.md**: File ini berisi dokumentasi proyek.

- 

- Config/config.go     : Place for all global varible configuration 
- Controller/*.go : Place for all logic process
- Entity/*.entiy    : Place for all model like strukur data
- Migrate/db.go   : Place for auto migration create database and table by reference model  in entity directory
- Migrate/seed.go   : Place for auto populate date di database   

## Installation

Steps to install this framework:

1. Clone the repository:
   git clone https://github.com/dendie-sanjaya/golang-arsitekur-template

##  Contact

If you have question, you can contact this email   
Email: dendie.sanjaya@gmail.com


License

This project is licensed under the MIT License.
