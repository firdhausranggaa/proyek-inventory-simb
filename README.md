# Enterprise Fullstack Book Inventory System (SIMB)

A robust, decoupled full-stack application for library and inventory management. This project demonstrates production-ready backend architecture paired with a reactive, modern frontend, showcasing industry-standard practices in relational database management, concurrency control, and secure authorization mechanisms.

## 🚀 Key Enterprise Features

*   **Decoupled Architecture:** Clean separation of concerns between the headless Go backend and the Vite-powered Vue 3 frontend.
*   **Concurrency & Data Integrity:** Implemented **Pessimistic Locking** (`FOR UPDATE`) during book borrowing to prevent race conditions, and **Database Transactions** (`TX`) to guarantee absolute stock accuracy.
*   **Resource Optimization:** Configured **Connection Pooling** (Max Open/Idle connections) to handle high traffic and prevent memory leaks.
*   **Graceful Shutdown:** Server safely completes ongoing database transactions before shutting down upon receiving termination signals.
*   **Audit Trails & Soft Delete:** Integrated GORM's Soft Delete (`DeletedAt`) for safe record auditing without permanent data loss.
*   **Relational Mapping:** Automated Foreign Key relations fetching synchronized transaction histories via GORM Preloading.
*   **Secure Authentication:** User registration and login utilizing **Bcrypt** for password hashing and **JWT** for stateless authentication.
*   **Role-Based Access Control (RBAC):** Distinct permission levels separating `Admin` (full CRUD access) and `Member` (read and borrow access).

## 🛠️ Tech Stack

**Backend:**
*   **Language:** Go (Golang)
*   **Framework:** Gin-Gonic (w/ CORS Middleware)
*   **ORM & DB:** GORM, PostgreSQL (Relational constraints & Locking)
*   **Security:** Golang-JWT (v4), X/Crypto (Bcrypt)

**Frontend:**
*   **Framework:** Vue 3 (Composition API)
*   **Build Tool:** Vite
*   **HTTP Client:** Axios (w/ JWT Interceptors)

## ⚙️ Installation & Local Setup

Clone the repository to get started:
```bash
git clone [https://github.com/firdhausranggaa/proyek-inventory-simb.git](https://github.com/firdhausranggaa/proyek-inventory-simb.git)
cd proyek-inventory-simb

```

### 1. Backend Setup (Go)

Navigate to the backend directory:

```bash
cd backend

```

Create a `.env` file in the `backend` root:

```env
POSTGRES_URL="host=localhost user=postgres password=yourpassword dbname=yourdb port=5432 sslmode=disable"
SUPER_USER="admin"
SUPER_PASS="123"
SUPER_SECRET="your-secure-jwt-secret-key"
APP_PORT="8080"

```

Install dependencies and run the server:

```bash
go mod tidy
go run main.go

```

*Note: Upon the first run, the API will automatically migrate tables, establish Foreign Keys, and seed the initial admin credentials and book catalogs.*

### 2. Frontend Setup (Vue)

Open a new terminal window and navigate to the frontend directory:

```bash
cd frontend

```

Create a `.env` file in the `frontend` root:

```env
VITE_API_URL=http://localhost:8080/api

```

Install dependencies and start the Vite development server:

```bash
npm install
npm run dev

```

Access the web application at `http://localhost:5173`.

## 📡 API Endpoints Reference

All protected routes require an `Authorization` header with the format: `Bearer <token>`.

### Authentication (Public)

| Method | Endpoint | Description |
| --- | --- | --- |
| `POST` | `/api/register` | Register a new member account |
| `POST` | `/api/login` | Authenticate and receive a JWT token |

### Book Management (Protected)

| Method | Endpoint | Access Role | Description |
| --- | --- | --- | --- |
| `GET` | `/api/books` | Admin / Member | List all books. Supports `?search=`, `?page=`, `?limit=` |
| `GET` | `/api/books/:id` | Admin / Member | Get details of a specific book |
| `POST` | `/api/books` | **Admin Only** | Add a new book to the inventory |
| `PUT` | `/api/books/:id` | **Admin Only** | Update an existing book's details |
| `DELETE` | `/api/books/:id` | **Admin Only** | Remove a book (Utilizes Audit/Soft Delete) |

### Borrowing System (Protected & Transactional)

| Method | Endpoint | Access Role | Description |
| --- | --- | --- | --- |
| `GET` | `/api/borrowings/me` | Admin / Member | Retrieve current user's active/past borrowing history |
| `POST` | `/api/borrow` | Admin / Member | Borrow a book (Applies Pessimistic Locking) |
| `POST` | `/api/return/:id` | Admin / Member | Return a borrowed book by Transaction ID |
