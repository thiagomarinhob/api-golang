# Etapa 1: Use uma imagem base para construir o aplicativo Go
FROM golang:1.23-alpine AS builder

# Defina o diretório de trabalho no container
WORKDIR /app

# Copie os arquivos do projeto para o diretório de trabalho no container
COPY . .

# Baixe as dependências do Go
RUN go mod download

# Compile o aplicativo Go apontando para o diretório onde está o main.go
RUN go build -o main ./cmd/api

# Etapa 2: Use uma imagem menor apenas para rodar o binário
FROM alpine:3.18

# Defina o diretório de trabalho no container
WORKDIR /app

# Copie o binário da etapa de build
COPY --from=builder /app/main .

# Exponha a porta que a aplicação vai usar (altere para a porta que seu app usa)
EXPOSE 8080

# Comando para rodar o binário
CMD ["./main"]
