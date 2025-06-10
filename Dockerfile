# 1. Base image
FROM golang:1.23.0

# 2. Cài đặt các tiện ích cần thiết
RUN apt-get update && apt-get install -y curl git

# 3. Cài air để hot reload
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b /usr/local/bin

# 4. Set working directory
WORKDIR /app

# 5. Copy go.mod, go.sum và tải dependencies
COPY go.mod go.sum ./
RUN go mod download

# 6. Copy toàn bộ source code
COPY . .

# 7. Expose port
EXPOSE 8080

# 8. Chạy bằng air
CMD ["air"]
