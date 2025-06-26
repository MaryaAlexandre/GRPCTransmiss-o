package main

import (
    "bufio"
    "context"
    "fmt"
    "io"
    "log"
    "os"
    "time"

    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"

    pb "GRPCTransmiss-o/proto"
)

func main() {
    // Conexão gRPC
    conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatalf("Falha ao conectar: %v", err)
    }
    defer conn.Close()

    client := pb.NewChatServiceClient(conn)
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    stream, err := client.ChatStream(ctx)
    if err != nil {
        log.Fatalf("Erro ao abrir stream: %v", err)
    }

    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Digite seu nome: ")
    user, _ := reader.ReadString('\n')
    user = user[:len(user)-1] 

    fmt.Println("Digite mensagens (ENTER para enviar). Ctrl+C para sair.")

    // Goroutine para receber mensagens
    go func() {
        for {
            msg, err := stream.Recv()
            if err == io.EOF {
                return
            }
            if err != nil {
                log.Printf("Stream encerrada: %v", err)
                return
            }
    
            if msg.User == user {
                continue
            }
           
            fmt.Printf("\n[%s]: %s\n> ", msg.User, msg.Message)
        }
    }()

 
    for {
        fmt.Print("> ")
        text, err := reader.ReadString('\n')
        if err != nil {
            log.Fatalf("Erro na leitura: %v", err)
        }
        text = text[:len(text)-1]
        if text == "" {
            continue
        }

        msg := &pb.ChatMessage{
            User:      user,
            Message:   text,
            Timestamp: time.Now().Unix(),
        }
        if err := stream.Send(msg); err != nil {
            log.Fatalf("Erro ao enviar mensagem: %v", err)
        }
    }
}
