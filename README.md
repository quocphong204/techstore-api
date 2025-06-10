# TechStore API

**TechStore API** là backend RESTful API quản lý bán linh kiện máy tính, xây dựng bằng Golang (Gin-Gonic, GORM), xác thực JWT, caching Redis, queue RabbitMQ, logging, đóng gói Docker, tài liệu hóa bằng Swaggo.

## Tính năng
- Quản lý user, phân quyền Admin/User
- Đăng ký, đăng nhập, xác thực JWT
- CRUD sản phẩm, danh mục, đơn hàng, giỏ hàng, đánh giá
- Caching Redis (sản phẩm, danh mục, token)
- Queue RabbitMQ cho đơn hàng
- Gửi email xác nhận đơn
- Log (logrus/zerolog), hỗ trợ logging tập trung
- Swagger UI (Swaggo)
- Docker Compose full stack

## Công nghệ sử dụng
- Go (Gin, GORM)
- PostgreSQL
- Redis
- RabbitMQ
- JWT
- Swaggo (Swagger UI)
- Docker, Docker Compose

---

## 1. Hướng dẫn cài đặt

### Clone & cấu hình
```sh
git clone https://github.com/quocphong204/techstore-api.git
cd techstore-api
cp .env.example .env
Chạy với Docker
sh
Sao chép
docker-compose up --build
(Mặc định sẽ lên luôn PostgreSQL, Redis, RabbitMQ...)

Chạy local không dùng Docker
Cài sẵn PostgreSQL, Redis, RabbitMQ

Sửa file .env cho đúng cấu hình

Chạy:

sh
Sao chép
go run main.go migrate
go run main.go seed
go run main.go
2. Swagger API docs
Truy cập: http://localhost:8080/swagger/index.html

3. Cấu trúc dự án
arduino
Sao chép
techstore-api/
├── cmd/                # Điểm vào app
├── config/
├── controllers/
├── services/
├── models/
├── repository/
├── middlewares/
├── utils/
├── routes/
├── docs/
├── Dockerfile
├── docker-compose.yml
└── README.md
4. API chính
POST /api/v1/register – Đăng ký

POST /api/v1/login – Đăng nhập (JWT)

GET /api/v1/products – Danh sách sản phẩm

POST /api/v1/products – Thêm (admin)

POST /api/v1/orders – Đặt hàng

(Xem chi tiết trên Swagger)

5. Seed dữ liệu mẫu
sh
Sao chép
go run main.go seed
Admin mặc định:

Email: admin@techstore.com

Password: admin123

6. Một số lệnh Docker
sh
Sao chép
docker-compose up -d
docker-compose down
docker-compose logs -f
7. Đóng góp
Fork, tạo branch mới

Commit rõ ràng

Pull request

8. License
MIT License

Dev chính: Liêu Quốc Phong
Email: [phong150718@gmail.com]
