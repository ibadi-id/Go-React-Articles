
# CRUD Article Menggunakan CoreUI Free React Admin Template v3 + GO Sebagai backend Rest API

Ini adalah aplikasi sederhana dengan menerapkan <strong>golang</strong> sebagai backend dan <strong>react</strong> sebagai frontend</p>

## Table of Contents

* [Installation](#installation)
* [Database](#database)
* [Penggunaan](#penggunaan)
* [Rest API](#rest-api)


## Installation

### Clone repo

``` bash
# clone the repo
$ git clone https://github.com/ibadi-id/Go-React-Articles.git

# go into app's directory
$ cd Go-React-Articles

# install depedency backend
$ cd backend
$ go mod tidy

# install depedency backend
$ cd frontend
$ yarn

```

## Database

Buat dummy database mysql sebagai berikut:

```
nama database : testdb
user: testuser
password: test12345
server: localhost
port: 3306

```


### Penggunaan

Untuk backend start server dengan
``` bash
# development server
$ go run .
```

Coba lihat ke [http://localhost:3333](http://localhost:3333). Ini adalah endpoint untuk server

Untuk frontend start server dengan
``` bash
# development server
$ yarn start
```

Coba lihat ke [http://localhost:3000](http://localhost:3000). Ini adalah url untuk frontend

### Build

Run `build` untuk build backend.

```bash
# build for production with minification
$ go run build
```

Run `build` untuk build frontend. Build file akan tersimpan di `build/` directory.

```bash
# build for production with minification
$ yarn build
```

