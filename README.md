
# CRUD React + GO Rest API

Ini adalah aplikasi sederhana dengan menerapkan <strong>golang</strong> sebagai backend dan <strong>react</strong> sebagai frontend</p>

Screenshot :

![Alt text](./screenshot/home.png?raw=true "home")
![Alt text](./screenshot/allposts.png?raw=true "allposts")
![Alt text](./screenshot/addpost.png?raw=true "addpost")
![Alt text](./screenshot/preview.png?raw=true "preview")

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