# Employee API - Go + Kubernetes + DevOps Project

A cloud-native Employee Management REST API built with Go and deployed using modern DevOps practices.

## Project Overview

This project demonstrates an end-to-end DevOps workflow, starting from application development to Kubernetes deployment and CI/CD automation.

The application exposes REST APIs to manage employee records and uses PostgreSQL as the backend database.

---

## Tech Stack

### Backend
- Go (Golang)
- Gorilla Mux
- PostgreSQL

### Containerization
- Docker
- Docker Compose

### Orchestration
- Kubernetes
- Helm

### DevOps (In Progress)
- GitHub Actions
- Amazon ECR
- Amazon EKS
- Argo CD
- Prometheus
- Grafana

---

## Project Structure

```text
employee-api/
│
├── config/
├── database/
├── handlers/
├── logger/
├── models/
├── routes/
├── services/
├── helm/
├── k8s/
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
└── main.go
```

---

## Features

- REST API
- CRUD Operations
- PostgreSQL Integration
- Configuration using Environment Variables
- Health Check Endpoint
- Dockerized Application
- Kubernetes Deployment
- Helm Chart
- ConfigMaps and Secrets
- Liveness & Readiness Probes

---

## API Endpoints

| Method | Endpoint | Description |
|---------|----------|-------------|
| GET | /health | Health Check |
| GET | /employees | Get All Employees |
| GET | /employees/{id} | Get Employee by ID |
| POST | /employees | Create Employee |
| PUT | /employees/{id} | Update Employee |
| DELETE | /employees/{id} | Delete Employee |

---

## Run Locally

### Clone Repository

```bash
git clone https://github.com/Vamsi-1409/employee-api.git
cd employee-api
```

### Start PostgreSQL

```bash
docker compose up -d postgres
```

### Run Application

```bash
go run .
```

The application will start on:

```
http://localhost:8080
```

---

## Build Docker Image

```bash
docker build -t employee-api .
```

Run the container:

```bash
docker run -p 8080:8080 employee-api
```

---

## Deploy to Kubernetes

Apply Kubernetes manifests:

```bash
kubectl apply -f k8s/
```

Or deploy using Helm:

```bash
helm install employee-api helm/go-web-app
```

---

## CI/CD Roadmap

- ✅ Go REST API
- ✅ Docker
- ✅ Docker Compose
- ✅ Kubernetes
- ✅ Helm
- 🚧 GitHub Actions
- 🚧 Amazon ECR
- 🚧 Amazon EKS
- 🚧 Argo CD
- 🚧 Prometheus
- 🚧 Grafana

---

## Author

**Vamsi Krishna**

GitHub: https://github.com/Vamsi-1409
