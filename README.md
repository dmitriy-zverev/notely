# Notely - CI/CD Pipeline & Google Cloud Deployment

![CI](https://github.com/dmitriy-zverev/notely/actions/workflows/ci.yml/badge.svg)
![CD](https://github.com/dmitriy-zverev/notely/actions/workflows/cd.yml/badge.svg)

This is a fork of the [Boot.dev Learn CICD starter project](https://github.com/bootdotdev/learn-cicd-starter) where I've implemented a complete CI/CD pipeline and deployment to Google Cloud Platform.

## 🚀 What's Added

This fork extends the original Notely application with:

- **Complete CI/CD Pipeline** using GitHub Actions
- **Google Cloud Run Deployment** with automated container builds
- **Database Migrations** with Goose
- **Security Scanning** with gosec
- **Code Quality Checks** with staticcheck and go fmt
- **Dockerized Application** for cloud deployment

## 🏗️ CI/CD Pipeline

### Continuous Integration (CI)
Triggered on pull requests to `main` branch:

- **Tests Job**: Runs unit tests with coverage and security scanning (gosec)
- **Style Job**: Enforces code formatting (go fmt) and linting (staticcheck)

### Continuous Deployment (CD)
Triggered on pushes to `main` branch:

1. **Build**: Compiles the Go application using production build script
2. **Docker**: Builds and pushes container image to Google Artifact Registry
3. **Database**: Runs migrations using Goose
4. **Deploy**: Deploys to Google Cloud Run in `europe-north1` region

## ☁️ Google Cloud Infrastructure

- **Platform**: Google Cloud Run (serverless containers)
- **Region**: europe-north1
- **Container Registry**: Google Artifact Registry
- **Database**: PostgreSQL (managed via DATABASE_URL secret)
- **Scaling**: Auto-scaling with max 4 instances
- **Access**: Public (--allow-unauthenticated)

## 🛠️ Local Development

### Prerequisites
- Go 1.22+ installed
- Docker (optional, for local containerization)

### Setup

1. Clone the repository:
```bash
git clone https://github.com/dmitriy-zverev/notely.git
cd notely
```

2. Create a `.env` file in the root directory:
```bash
PORT="8080"
# Add DATABASE_URL for database mode (optional for local dev)
```

3. Run the application:
```bash
go build -o notely && ./notely
```

The server will start at `http://localhost:8080`.

### Development Commands

```bash
# Run tests
go test ./... -cover

# Format code
go fmt ./...

# Run security scan
gosec ./...

# Run linter
staticcheck ./...

# Build for production
./scripts/buildprod.sh
```

## 📁 Project Structure

```
.
├── .github/workflows/     # CI/CD pipeline definitions
├── internal/             # Internal Go packages
│   ├── auth/            # Authentication logic
│   └── database/        # Database models and queries
├── scripts/             # Build and migration scripts
├── sql/                 # Database schema and queries
├── static/              # Static web assets
├── Dockerfile           # Container definition
└── main.go             # Application entry point
```

## 🔧 Configuration

The application uses environment variables for configuration:

- `PORT`: Server port (default: 8080)
- `DATABASE_URL`: PostgreSQL (or another database) connection string (for database mode)

## 📝 Original Project

This is based on the "Notely" starter application from the [Boot.dev Learn CICD course](https://boot.dev). The original project provides a foundation for learning CI/CD concepts, which we've extended with a production-ready deployment pipeline.

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Ensure all CI checks pass
5. Submit a pull request

The CI pipeline will automatically run tests and style checks on your pull request.
