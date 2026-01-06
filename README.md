![Go](https://img.shields.io/badge/Language-Go-00ADD8?logo=go)

## API de Banco

### Contexto
+ API de banco com usuários(empresa/indivíduo), saques e depósitos
+ Não contempla concorrência

### Tecnologias usadas
+ Go Language
+ PostgreSQL
+ Chi framework v5
+ Pacote errors
+ Pacote google/uuid
+ Pacote gorilla/csrf
+ Pacote json
+ Pacote http
+ Pacote log
+ Pacote pgx/v5 
+ SQLc
+ Tern
+ Docker

### Como rodar o programa
```bash
docker-compose up -d
go mod tidy
go run ./cmd/api
```
- OBS: Requer Go e Docker instalado em seu computador

### Como parar o container Docker e excluir os dados

```bash
docker-compose down -v
```