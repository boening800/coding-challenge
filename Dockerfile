# imagen base
FROM golang:1.22

# carpeta dentro del contenedor
WORKDIR /app

# copiar dependencias
COPY go.mod go.sum ./
RUN go mod download

# copiar todo el proyecto
COPY . .

# compilar aplicación
RUN go build -o main .

# puerto de la API
EXPOSE 3000

# ejecutar app
CMD ["./main"]