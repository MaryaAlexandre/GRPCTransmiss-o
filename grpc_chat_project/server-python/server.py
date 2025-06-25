import grpc
from concurrent import futures
import time
import threading
import proto.chat_pb2 as chat_pb2
import proto.chat_pb2_grpc as chat_pb2_grpc

# Lista global de streams dos clientes conectados
connected_clients = []

class ChatService(chat_pb2_grpc.ChatServiceServicer):
    def ChatStream(self, request_iterator, context):
        # Adiciona o stream do cliente na lista global
        q = queue.Queue()
        connected_clients.append(q)

        def send_messages():
            while True:
                message = q.get()
                if message is None:
                    break
                yield message

        def recv_messages():
            try:
                for new_message in request_iterator:
                    # Enviar mensagem para todos os clientes
                    for client_q in connected_clients:
                        client_q.put(new_message)
            except grpc.RpcError:
                pass
            finally:
                connected_clients.remove(q)
                q.put(None)

        recv_thread = threading.Thread(target=recv_messages)
        recv_thread.start()

        # Gerador que envia mensagens para o cliente
        while True:
            message = q.get()
            if message is None:
                break
            yield message

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    chat_pb2_grpc.add_ChatServiceServicer_to_server(ChatService(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    print("Servidor gRPC rodando na porta 50051...")
    try:
        while True:
            time.sleep(86400)
    except KeyboardInterrupt:
        server.stop(0)

if __name__ == "__main__":
    import queue
    serve()