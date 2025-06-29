# Projeto Chat gRPC - Python (servidor) + Go (cliente)

## Descrição

Projeto simples de chat em tempo real usando gRPC com streaming bidirecional.
- Servidor em Python
- Cliente em Go
- Comunicação entre duas linguagens
- Permite múltiplos clientes simultâneos

## Estrutura do projeto

- proto/chat.proto: arquivo de definição gRPC/protobuf
- server-python/server.py: servidor gRPC em Python
- client-go/main.go: cliente gRPC em Go

## Como rodar no Codespaces

1. Instale dependências Python e Go:

```bash
sudo apt update
sudo apt install -y python3-pip golang-go
pip3 install grpcio grpcio-tools
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
export PATH=$PATH:$(go env GOPATH)/bin
```

2. Gere o código gRPC a partir do proto:

```bash
# Python
python3 -m grpc_tools.protoc -I=proto --python_out=proto --grpc_python_out=proto proto/chat.proto

# Go (execute dentro da pasta client-go)
cd client-go
protoc --go_out=. --go-grpc_out=. ../proto/chat.proto
cd ..
```

3. Rode o servidor (em uma aba do terminal):

```bash
python3 server-python/server.py
```

4. Rode o cliente (em outra aba):

```bash
cd client-go
go run main.go
```

## Apresentação (Canva):
(https://www.canva.com/design/DAGrUtWCW3U/zR5V4suyo6qB2QD7h0llyA/view?utm_content=DAGrUtWCW3U&utm_campaign=designshare&utm_medium=link2&utm_source=uniquelinks&utlId=hcb3e4a86e1)

## Dontpad para teste colaborativo:
  [https://dontpad.com/Chat-Real](https://dontpad.com/Chat-Real)
