Tahapan untuk migrasi data awal artikel:

1. install driver golang migrate
```
go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

2. perintah untuk migration :
```
migrate -database mysql://testuser:tes12345@tcp(localhost:3306)/testdb?query -path migration up
```