# Kanban Backend - Go

Backend REST API para aplicação Mini Kanban desenvolvido em Go 1.25.

## Arquitetura

O projeto segue uma arquitetura em camadas:

- **models/** - Entidades de domínio e DTOs
- **repository/** - Camada de persistência (in-memory)
- **service/** - Lógica de negócio e validações
- **handlers/** - Camada HTTP (controllers)

## Endpoints

### Tasks

- `GET /tasks` - Lista todas as tarefas
- `GET /tasks/{id}` - Busca tarefa por ID
- `POST /tasks` - Cria nova tarefa
- `PUT /tasks/{id}` - Atualiza tarefa
- `DELETE /tasks/{id}` - Remove tarefa

### Exemplo de Request

**Criar tarefa:**
```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title":"Minha tarefa","description":"Descrição opcional"}'
```

**Atualizar status:**
```bash
curl -X PUT http://localhost:8080/tasks/{id} \
  -H "Content-Type: application/json" \
  -d '{"status":"in_progress"}'
```

## Como Rodar

### Localmente

```bash
# Instalar dependências
go mod download

# Rodar aplicação
make run

# Ou diretamente
go run main.go
```

### Docker

```bash
# Build da imagem
make docker-build

# Rodar container
make docker-run
```

## Testes

```bash
# Rodar todos os testes
make test

# Com coverage
go test -v -cover ./...
```

## Linter

```bash
# Instalar golangci-lint
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Rodar linter
make lint
```

## Decisões Técnicas

- **In-memory storage**: Persistência em memória com sync.RWMutex para thread-safety
- **Stdlib HTTP**: Uso da biblioteca padrão sem frameworks externos para simplicidade
- **UUID**: Geração de IDs únicos com google/uuid
- **CORS**: Middleware configurado para permitir acesso do frontend
- **Validações**: Título obrigatório, status validado (todo/in_progress/done)

## Limitações

- Dados não persistem após restart (in-memory)
- Sem autenticação/autorização
- Sem paginação na listagem
- Sem logging estruturado

## Melhorias Futuras

- Adicionar persistência em banco de dados (PostgreSQL)
- Implementar logging estruturado (zerolog/zap)
- Adicionar métricas e observabilidade
- Implementar paginação e filtros
- Adicionar autenticação JWT
