# Fishing Comp App

## Migration

```bash
createdb fishingcomp
migrate -path postgres/migrations -database "postgres://<username>@localhost:5432/fishingcomp?sslmode=disable" up
```