# Fullstack Book Inventory System (SIMB)

A robust, decoupled full-stack application for library and inventory management. This project demonstrates end-to-end architecture, integrating a secure RESTful API built with Go (Golang) and a reactive, modern frontend powered by Vue 3 & Vite.

## 🚀 Key Features

*   **Decoupled Architecture:** Clean separation of concerns between the headless Go backend and the Vite-powered Vue frontend.
*   **Secure Authentication:** User registration and login utilizing **Bcrypt** for password hashing and **JWT** (JSON Web Tokens) for stateless authentication[cite: 27].
*   **Role-Based Access Control (RBAC):** Distinct permission levels separating `Admin` (full CRUD access) and `Member` (read and borrow access)[cite: 27].
*   **Database Transactions & Soft Delete:** Safe borrowing/returning mechanisms using GORM's `TX` functions to ensure stock data consistency[cite: 27]. Integrated **Soft Delete** for safe record auditing without permanent data loss.
*   **Dynamic Queries & Reactive UI:** Implemented mathematical pagination and `ILIKE` search filtering on the backend[cite: 27], consumed by a real-time reactive Vue dashboard.
*   **Personalized Dashboard:** Dedicated borrowing history tracker (`/borrowings/me`) for authenticated users.

## 🛠️ Tech Stack

**Backend:**
*   Go / Gin-Gonic[cite: 27]
*   GORM & PostgreSQL[cite: 27]
*   Golang-JWT (v4) & X/Crypto (Bcrypt)[cite: 27]

**Frontend:**
*   Vue 3 (Composition API)
*   Vite
*   Axios

## ⚙️ Prerequisites

*   Go (1.16+) and Node.js installed on your local machine.
*   PostgreSQL server running locally or remotely.

## 📦 Installation & Setup

Clone the repository to get started:
```bash
git clone [https://github.com/firdhausranggaa/proyek-inventory.git](https://github.com/firdhausranggaa/proyek-inventory.git)
cd proyek-inventory

```

### 1. Backend Setup (Go)

Navigate to the backend directory and configure the environment:

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

*Note: Upon the first run, the API will automatically migrate tables and seed the initial admin credentials and book catalogs.*

### 2. Frontend Setup (Vue 3)

Open a new terminal window, navigate to the frontend directory:

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
| `POST` | `/api/register` | Register a new member account

 |
| `POST` | `/api/login` | Authenticate and receive a JWT token

 |

### Book Management (Protected)

| Method | Endpoint | Access Role | Description |
| --- | --- | --- | --- |
| `GET` | `/api/books` | Admin / Member | List all books. Supports `?search=`<br> |
| `GET` | `/api/books/:id` | Admin / Member | Get details of a specific book

 |
| `POST` | `/api/books` | **Admin Only** | Add a new book to the inventory

 |
| `PUT` | `/api/books/:id` | **Admin Only** | Update an existing book's details

 |
| `DELETE` | `/api/books/:id` | **Admin Only** | Remove a book (Utilizes Soft Delete) |

### Borrowing System (Protected)

| Method | Endpoint | Access Role | Description |
| --- | --- | --- | --- |
| `GET` | `/api/borrowings/me` | Admin / Member | Retrieve current user's borrowing history |
| `POST` | `/api/borrow` | Admin / Member | Borrow a book (requires `book_id` in JSON)

 |
| `POST` | `/api/return/:id` | Admin / Member | Return a borrowed book by Transaction ID |