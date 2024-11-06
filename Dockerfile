# Usa una imagen base de Go
FROM golang:1.23.2 as builder

# Establece el directorio de trabajo
WORKDIR /app

# Copia los archivos de módulos
COPY go.mod go.sum ./
RUN go mod download

# Copia todo el código de la aplicación
COPY . .

# Compila el proyecto
RUN CGO_ENABLED=0 GOOS=linux go build -o grpc_server ./main/server/server.go

# Etapa final
FROM debian:bullseye-slim

# Establece el directorio de trabajo para la imagen final
WORKDIR /root/

# Copia el binario compilado desde la etapa de construcción
COPY --from=builder /app/grpc_server .

# Expone el puerto que usa el servidor gRPC
EXPOSE 50051

# Comando por defecto al ejecutar el contenedor
CMD ["./grpc_server"]
