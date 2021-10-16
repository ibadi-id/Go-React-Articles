
# CRUD Article Menggunakan CoreUI Free React Admin Template v3 + GO Sebagai backend Rest API

CoreUI is meant to be the UX game changer. Pure & transparent code is devoid of redundant components, so the app is light enough to offer ultimate user experience. This means mobile devices also, where the navigation is just as easy and intuitive as on a desktop or laptop. The CoreUI Layout API lets you customize your project for almost any device – be it Mobile, Web or WebApp – CoreUI covers them all!

## Table of Contents

* [Installation](#installation)
* [Database](#database)
* [Penggunaan](#penggunaan)
* [Routes](#routes)


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

Migrasi awal untuk initial data:

1. install driver golang migrate
```
$ go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

2. perintah untuk migration :
```
$ cd backend
$ migrate -database mysql://testuser:tes12345@tcp(localhost:3306)/testdb?query -path migration up
```


### Penggunaan

Untuk backend start server dengan
``` bash
# development server
$ cd backend
$ go run .
```

Coba lihat ke [http://localhost:3333](http://localhost:3333). Ini adalah endpoint untuk server

Untuk frontend start server dengan
``` bash
# development server
$ cd frontend
$ yarn start
```

Coba lihat ke [http://localhost:3000](http://localhost:3000). Ini adalah url untuk frontend

### Build

Run `build` untuk build backend.

```bash
# build for production with minification
$ cd backend
$ go run build
```

Run `build` untuk build frontend. Build file akan tersimpan di `build/` directory.

```bash
# build for production with minification
$ cd frontend
$ yarn build
```

## Routes

<details>
<summary>`/article`</summary>

- [(*Cors).Handler-fm]()
- [RequestID]()
- [Logger]()
- [Recoverer]()
- [URLFormat]()
- [SetContentType.func1]()
- **/article**
	- **/**
		- _GET_
			- [ListArticles]()
		- _POST_
			- [CreateArticle]()

</details>
<details>
<summary>`/article/{articleID}`</summary>

- [(*Cors).Handler-fm]()
- [RequestID]()
- [Logger]()
- [Recoverer]()
- [URLFormat]()
- [SetContentType.func1]()
- **/article**
	- **/{articleID}**
		- [ArticleCtx]()
		- **/**
			- _DELETE_
				- [DeleteArticle]()
			- _GET_
				- [GetArticle]()
			- _PUT_
				- [UpdateArticle]()

</details>
<details>
<summary>`/article/{limit}/{offset}`</summary>

- [(*Cors).Handler-fm]()
- [RequestID]()
- [Logger]()
- [Recoverer]()
- [URLFormat]()
- [SetContentType.func1]()
- **/article**
	- **/{limit}/{offset}**
		- _GET_
			- [ListArticlesWithPage]()

</details>

Total # of routes: 3