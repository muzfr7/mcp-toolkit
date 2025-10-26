# MCP Toolkit

A boilerplate for building Model Context Protocol (MCP) servers in Go, developed using the [official Go SDK](https://github.com/modelcontextprotocol/go-sdk) for MCP. This project provides a foundation for creating MCP implementations with reusable components and example tools.

**🚧 Work in Progress** - This project is actively being developed and improved.

## Features

- **Calculator Tool**: Supports add, subtract, multiply, and divide operations
- **Health Check Endpoints**: Kubernetes-ready liveness and readiness probes
- **Structured JSON Logging**: With server name and environment metadata
- **Environment Configuration**: Configurable via environment variables
- **Clean Architecture**: Following Go best practices

## Quick Start

### 1. Setup Environment

Copy the environment template and customize:

```bash
cp .env.dist .env
# Edit .env with your preferred values
```

### 2. Run the Server

```bash
# Using default configuration
go run ./cmd/mcp-server

# Or with custom environment variables
SERVER_NAME=mcp-server PORT=9090 go run ./cmd/mcp-server
```

### 3. Build for Production

```bash
# Build the server
go build -o bin/mcp-server ./cmd/mcp-server
./bin/mcp-server
```

## Usage

### MCP Server

Once running, the server will be available at port 8080 (or your configured port):
- **MCP Endpoint**: `http://localhost:8080/sse`
- **Liveness Check**: `http://localhost:8080/live`
- **Readiness Check**: `http://localhost:8080/ready`

### Calculator Tool

The calculator tool accepts:

```json
{
  "x": 10.5,
  "y": 3.2,
  "operation": "add"
}
```

Supported operations: `add`, `subtract`, `multiply`, `divide`

### Health Checks

The server provides Kubernetes-compatible health check endpoints:

#### Liveness Probe
- **Endpoint**: `/live`
- **Purpose**: Checks if the application is running and not deadlocked
- **Response**: JSON with status, timestamp, service name, and version

#### Readiness Probe
- **Endpoint**: `/ready`  
- **Purpose**: Checks if the application is ready to serve traffic
- **Response**: JSON with status, timestamp, service name, and version

#### Example Response
```json
{
  "status": "healthy",
  "timestamp": "2025-10-27T05:57:11.899468+01:00",
  "service": "mcp-server",
  "version": "0.0.1"
}
```

## Development

### Adding New Tools

1. Create a new package under `internal/tools/`
2. Implement the tool interface and handlers
3. Register the tool in `internal/mcpserver/server.go`

## License
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

This project is distributed under the MIT License. See `LICENSE.md` for more information.
