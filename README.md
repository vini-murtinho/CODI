# Sistema Kanban - Mini Board

Sistema completo de gerenciamento de tarefas estilo Kanban com backend em Go e frontend em React com TypeScript.

## 📋 Índice

- [Visão Geral](#visão-geral)
- [Tecnologias](#tecnologias)
- [Estrutura do Projeto](#estrutura-do-projeto)
- [Backend - Go](#backend---go)
- [Frontend - React](#frontend---react)
- [Como Executar](#como-executar)
- [API Endpoints](#api-endpoints)
- [Funcionalidades](#funcionalidades)

## 🎯 Visão Geral

Sistema de quadro Kanban para gerenciamento de tarefas com três colunas:
- **A Fazer** (Todo)
- **Em Progresso** (In Progress)
- **Concluídas** (Done)

O sistema permite criar, editar, deletar e mover tarefas entre colunas através de drag-and-drop.

## 🚀 Tecnologias

### Backend
- **Go 1.25** - Linguagem de programação
- **Stdlib HTTP** - Servidor HTTP nativo
- **UUID** - Geração de IDs únicos

### Frontend
- **React 18** - Biblioteca UI
- **TypeScript** - Tipagem estática
- **Tailwind CSS** - Estilização
- **React DnD** - Drag and drop
- **React DnD HTML5 Backend** - Backend para drag and drop

## 📁 Estrutura do Projeto

```
.
├── frontend/     # Aplicação React
│   ├── src/
│   │   ├── components/              # Componentes React
│   │   │   ├── AddTodoTaskButton.tsx
│   │   │   ├── AddTodoTaskModal.tsx
│   │   │   ├── DraggableTaskCard.tsx
│   │   │   └── TodoColumn.tsx
│   │   ├── helpers/                 # Funções auxiliares e API
│   │   │   └── helpers.ts
│   │   ├── reducers/                # Gerenciamento de estado
│   │   │   └── taskReducer.ts
│   │   ├── screen/                  # Telas principais
│   │   │   └── App.tsx
│   │   └── index.js
│   └── package.json
│
└── backend/                 # API Backend
    ├── handlers/                    # Controladores HTTP
    │   └── task_handler.go
    ├── models/                      # Modelos de dados
    │   └── task.go
    ├── repository/                  # Camada de persistência
    │   ├── task_repository.go
    │   ├── task_repository_test.go
    │   └── mock_repository.go
    ├── service/                     # Lógica de negócio
    │   ├── task_service.go
    │   └── task_service_test.go
    ├── main.go                      # Entry point
    ├── Dockerfile
    ├── Makefile
    └── go.mod
```

## 🔧 Backend - Go

### Arquitetura

O backend segue uma arquitetura em camadas:

1. **Handlers** - Camada HTTP que processa requisições
2. **Service** - Lógica de negócio e validações
3. **Repository** - Persistência de dados (in-memory)
4. **Models** - Estruturas de dados e DTOs

### Características

- ✅ Persistência em memória com thread-safety (sync.RWMutex)
- ✅ Validação de dados (título obrigatório, status válido)
- ✅ Testes unitários com 92.6% de cobertura
- ✅ CORS configurado para o frontend
- ✅ Erros tratados adequadamente
- ✅ Geração de IDs únicos

### Modelos de Dados

```go
type Task struct {
    ID          string `json:"id"`
    Title       string `json:"title"`
    Description string `json:"description,omitempty"`
    Status      Status `json:"status"`
    Completed   bool   `json:"completed"`
}

type Status string
const (
    StatusTodo       Status = "todo"
    StatusInProgress Status = "in_progress"
    StatusDone       Status = "done"
)
```

### Executar Backend

```bash
cd backend

# Instalar dependências
go mod download

# Executar servidor
make run
# ou
go run main.go

# Executar testes
make test

# Build
make build

# Docker
make docker-build
make docker-run
```

O servidor será iniciado em `http://localhost:8080`

## 💻 Frontend - React

### Componentes

#### AddTodoTaskButton
Botão para abrir o modal de criação de tarefas.

#### AddTodoTaskModal
Modal para criar e editar tarefas com campos de título e descrição.

#### DraggableTaskCard
Card de tarefa com suporte a drag-and-drop, exibindo:
- Título e descrição
- Status visual com cores
- Botões de ação (completar, editar, deletar)

#### TodoColumn
Coluna do Kanban que aceita drop de tarefas e organiza cards por status.

### Gerenciamento de Estado

Utiliza `useReducer` para gerenciar o estado das tarefas com as seguintes ações:
- `ADD_TASK` - Adiciona nova tarefa
- `TOGGLE_TASK` - Alterna status de conclusão
- `DELETE_TASK` - Remove tarefa
- `UPDATE_TASK` - Atualiza tarefa existente
- `SET_TASKS` - Define lista completa de tarefas
- `MOVE_TASK` - Move tarefa entre colunas

### API Helper

Arquivo `helpers.ts` contém funções para comunicação com o backend:
- `fetchAllTasks()` - Busca todas as tarefas
- `createTask()` - Cria nova tarefa
- `updateTask()` - Atualiza tarefa
- `deleteTask()` - Remove tarefa
- Funções auxiliares de filtragem por status

### Executar Frontend

```bash
cd frontend

# Instalar dependências
npm install

# Executar em desenvolvimento
npm start

# Build para produção
npm run build
```

A aplicação será aberta em `http://localhost:3000`

## 🌐 API Endpoints

### GET /tasks
Lista todas as tarefas.

**Resposta:**
```json
[
  {
    "id": "uuid-123",
    "title": "Minha Tarefa",
    "description": "Descrição",
    "status": "todo",
    "completed": false
  }
]
```

### GET /tasks/{id}
Busca tarefa específica por ID.

### POST /tasks
Cria nova tarefa.

**Request:**
```json
{
  "title": "Nova Tarefa",
  "description": "Descrição opcional"
}
```

**Resposta:** Tarefa criada com status 201

### PUT /tasks/{id}
Atualiza tarefa existente.

**Request:**
```json
{
  "title": "Título atualizado",
  "description": "Nova descrição",
  "status": "in_progress"
}
```

### DELETE /tasks/{id}
Remove tarefa por ID.

**Resposta:** Status 204 (No Content)

## ✨ Funcionalidades

### Implementadas

- ✅ Criar tarefas com título e descrição
- ✅ Editar tarefas existentes
- ✅ Deletar tarefas
- ✅ Mover tarefas entre colunas via drag-and-drop
- ✅ Marcar tarefas como concluídas
- ✅ Indicadores visuais de status (cores)
- ✅ Contador de tarefas por coluna
- ✅ Validações no frontend e backend
- ✅ Loading state durante carregamento inicial
- ✅ Animações e transições suaves
- ✅ Design responsivo com Tailwind CSS

### Limitações Conhecidas

- ❌ Dados não persistem após reiniciar o servidor (in-memory)
- ❌ Sem autenticação/autorização
- ❌ Sem paginação na listagem
- ❌ Sem filtros ou busca

## 🚀 Como Executar o Sistema Completo

### Opção 1: Execução Manual

**Terminal 1 - Backend:**
```bash
cd backend
go run main.go
```

**Terminal 2 - Frontend:**
```bash
cd frontend
npm install
npm start
```

Acesse: `http://localhost:3000`

### Opção 2: Com Docker (Backend)

```bash
# Backend
cd backend
make docker-build
make docker-run

# Frontend (terminal separado)
cd frontend
npm install
npm start
```

## 🧪 Testes

### Backend
```bash
cd backend

# Executar todos os testes
make test

# Com coverage detalhado
go test -v -cover ./...

# Gerar relatório HTML de coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

**Cobertura atual:** 92.6% no service layer

## 📝 Melhorias Futuras

### Backend
- [ ] Adicionar persistência em banco de dados (PostgreSQL)
- [ ] Implementar logging estruturado (zerolog/zap)
- [ ] Adicionar métricas e observabilidade
- [ ] Implementar paginação e filtros
- [ ] Adicionar autenticação JWT
- [ ] Implementar rate limiting
- [ ] Adicionar validação com biblioteca externa

### Frontend
- [ ] Adicionar testes unitários (Jest/React Testing Library)
- [ ] Implementar testes E2E (Cypress/Playwright)
- [ ] Adicionar filtros e busca de tarefas
- [ ] Implementar ordenação customizável
- [ ] Adicionar modo escuro
- [ ] Suporte a múltiplos boards
- [ ] Adicionar tags/categorias às tarefas
- [ ] Implementar notificações