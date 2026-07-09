# Laravel-Inspired Custom Go HTTP Router

A lightweight, high-performance HTTP multiplexer and routing engine built from scratch in Go. This project explores the core mechanics of the HTTP request lifecycle, modular software architecture, and concurrent middleware execution without relying on heavy third-party frameworks.

⚠️ **Project Status: In Progress** *This framework is currently under active development. I am actively refining the middleware pipeline, adding robust HTTP method support, and optimizing performance.*

---

## 🚀 Why I Built This

As a backend engineer with experience in Laravel, I wanted to bridge the gap between Laravel's incredible **Developer Experience (DX)** and Go's raw **network performance and concurrency model**. 

Instead of just blindly using third-party Go frameworks like Gin or Fiber, I wanted to build the engine myself. This project demonstrates a deep understanding of standard library fundamentals, memory management, and how web servers handle data under the hood.

## ✨ Key Features (So Far)

* **$O(1)$ Route Lookup:** Uses a strict map-based routing table (`"METHOD/path"`). Whether the application has 5 or 5,000 routes, the router locates the correct handler instantly without linear array scanning.
* **Laravel-Inspired Fluent Interface:** Features method chaining for elegant, highly readable route and middleware registration (e.g., `Router.GET(...).Middleware = []func...`).
* **Decoupled Architecture (Clean Code):** Complete separation of concerns. The Router handles multiplexing, Middlewares intercept requests, and Controllers manage the business logic—communicating purely through standard interfaces (`http.ResponseWriter` and `*http.Request`).
* **Built for Concurrency:** Designed to safely leverage Go's lightweight Goroutines, allowing the server to handle thousands of concurrent connections with minimal memory overhead.

## 🛠️ Current Project Structure

```text
.
├── main.go                     # Server entry point & initialization
├── router/
│   └── router.go               # Core routing engine & ServeHTTP implementation
│   └── routes/
│       └── web.go              # Route registration and middleware mapping
├── http/
│   ├── controllers/
│   │   └── UserController.go   # Sample controller handling JSON payloads
│   └── Middleware/
│       ├── logMiddleware.go   # Request logging middleware
│       └── printMiddleware.go  # Sample interceptor middleware