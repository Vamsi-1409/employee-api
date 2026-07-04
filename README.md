# 🚀 Employee API - Cloud Native DevOps Project

A production-inspired **Go REST API** deployed on **Amazon EKS** using **Kubernetes**, **Helm**, **GitHub Actions**, and **Argo CD** following **GitOps** principles.

The project demonstrates a complete CI/CD pipeline that automatically builds, tests, containerizes, publishes, and deploys application changes to Kubernetes.

---

## ✨ Key Features

* RESTful API built with Go
* PostgreSQL database integration
* Dockerized application
* Kubernetes deployment using Helm
* GitHub Actions CI/CD
* Amazon ECR for container registry
* Amazon EKS for Kubernetes orchestration
* Argo CD for GitOps-based deployment
* Automated image versioning using Git commit SHA
* Health checks with Kubernetes probes
* ConfigMaps and Secrets for configuration management
* Ingress for external access

---

# 🏗️ Solution Architecture

```text id="0ow9p7"
                     Developer
                         │
              Feature Branch Development
                         │
                         ▼
                  Pull Request Review
                         │
                         ▼
                    Merge to main
                         │
                         ▼
                GitHub Actions (CI)
         ┌────────────────────────────────┐
         │ • Go Format                    │
         │ • Go Vet                       │
         │ • Unit Tests                   │
         │ • Build Binary                 │
         │ • Build Docker Image           │
         │ • Push Image to Amazon ECR     │
         └────────────────────────────────┘
                         │
                         ▼
               GitHub Actions (CD)
         ┌────────────────────────────────┐
         │ Update Helm Image Tag          │
         │ Commit values.yaml             │
         └────────────────────────────────┘
                         │
                         ▼
                     Git Repository
                         │
                         ▼
                      Argo CD
                         │
                         ▼
                       Helm
                         │
                         ▼
                  Amazon EKS Cluster
                         │
          ┌──────────────┴──────────────┐
          │                             │
          ▼                             ▼
     Employee API                 PostgreSQL
```

---

# ⚙️ Technology Stack

| Category           | Technologies   |
| ------------------ | -------------- |
| Language           | Go             |
| Database           | PostgreSQL     |
| Containerization   | Docker         |
| Orchestration      | Kubernetes     |
| Package Manager    | Helm           |
| CI/CD              | GitHub Actions |
| GitOps             | Argo CD        |
| Cloud              | AWS            |
| Container Registry | Amazon ECR     |
| Kubernetes Service | Amazon EKS     |
| Version Control    | Git & GitHub   |

---

# 📁 Repository Structure

```text id="w4ojoc"
employee-api
│
├── .github/
│   └── workflows/
│       ├── ci.yaml
│       └── cd.yaml
│
├── argocd/
│   └── application.yaml
│
├── helm/
│   ├── templates/
│   ├── Chart.yaml
│   └── values.yaml
│
├── handlers/
├── services/
├── routes/
├── models/
├── database/
├── config/
├── logger/
│
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── main.go
└── README.md
```

---

# 🔄 CI/CD Workflow

## Continuous Integration

Every Pull Request and push to the `main` branch triggers the CI pipeline.

The pipeline performs:

* Repository checkout
* Go environment setup
* Dependency download
* Code formatting validation
* Static code analysis (`go vet`)
* Unit testing
* Application build
* Multi-platform Docker image build
* Push Docker image to Amazon ECR

---

## Continuous Deployment

After a successful CI pipeline:

* The latest Docker image is pushed to Amazon ECR.
* The Helm chart (`values.yaml`) is updated with the new Git commit SHA.
* The update is committed back to the repository.
* Argo CD detects the change.
* Kubernetes is synchronized automatically.
* The application is updated without manual intervention.

---

# 🔁 GitOps Workflow

```text id="ixw6nm"
Code Change
      │
      ▼
Feature Branch
      │
      ▼
Pull Request
      │
      ▼
Merge to Main
      │
      ▼
GitHub Actions (CI)
      │
      ▼
Docker Image → Amazon ECR
      │
      ▼
GitHub Actions (CD)
      │
      ▼
Update Helm values.yaml
      │
      ▼
Git Commit
      │
      ▼
Argo CD Sync
      │
      ▼
Amazon EKS Deployment
```

---

# 🚀 Running the Application

### Clone the Repository

```bash id="g9m4ef"
git clone https://github.com/Vamsi-1409/employee-api.git
cd employee-api
```

### Run Locally

```bash id="9hxfzs"
go mod download
go run main.go
```

---

# 🐳 Docker

Build the image

```bash id="k3zhbj"
docker build -t employee-api .
```

Run the container

```bash id="1jlwmg"
docker run -p 8080:8080 employee-api
```

---

# ☸️ Kubernetes Deployment

Deploy using Helm

```bash id="4mb9gc"
cd helm

helm install employee-api .
```

Upgrade

```bash id="7l0zzy"
helm upgrade employee-api .
```

---

# 📦 Kubernetes Resources

The Helm chart deploys:

* Namespace
* Employee API Deployment
* PostgreSQL Deployment
* Services
* ConfigMap
* Secret
* Ingress

---

# 📈 Project Highlights

* End-to-end DevOps implementation
* GitOps deployment strategy
* Immutable Docker image versioning
* Multi-architecture image builds
* Infrastructure deployed on Amazon EKS
* Automated deployment with Argo CD
* Kubernetes best practices
* Clean repository structure
* Feature branch development workflow
* Pull Request-based delivery process

---

# 🛣️ Future Enhancements

* Terraform for Infrastructure as Code
* Prometheus & Grafana monitoring
* Loki-based centralized logging
* AWS Secrets Manager integration
* PostgreSQL StatefulSet with Persistent Volumes
* Horizontal Pod Autoscaler (HPA)
* Trivy image scanning
* SonarQube code quality analysis
* Slack deployment notifications

---

# 👨‍💻 Author

**Vamsi Krishna**

DevOps Engineer passionate about Cloud, Kubernetes, CI/CD, and GitOps.

GitHub: https://github.com/Vamsi-1409/employee-api

