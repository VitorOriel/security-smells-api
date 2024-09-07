# Smelly Kube API

Smelly Kube API was created to identify security vulnerabilities in Kubernetes manifests. It uses a REST API to receive the manifest files from the client, and returns a response with the found vulnerabilities. The API documentation can be seen in **API** section.

## Clients

At the moment, Smelly Kube have the following clients:
- [Visual Studio Code extension](https://github.com/VitorOriel/smelly-kube-vscode-plugin)

## Running Smelly Kube
You can either run the server directly through golang or via Docker.

**Run using Docker**:
To run the application using Docker, just run `docker compose up server`

**Run directly using go**:
1. `cd security-smells-api/`
2. `go mod tidy`
3. `go run main.go`

The server uses the localhost with port `3000`.

## Operating System

The server was developed under Ubuntu 22.04.4 LTS

## Technologies

- Golang v1.22
- Docker v24.0.5

## Dependencies

The dependencies can be seen in [go.mod](security-smells-api/go.mod) file

## API
### POST /api/v1/smelly
#### Request
**Headers**
| Key        | Value         |
|------------|---------------|
| Content-Type | application/json |

**Body**
```json
{
  "fileName": "manifest.yaml",
  "yamlToValidate": "<file content>",
}
```

#### Response 200 - OK
```json
{
  "meta": {
    "totalOfSmells": 1,
    "decodedWorkloads": {
      "ReplicaSets": 0,
      "Deployments": 1,
      "Pods": 0,
      "Jobs": 0,
      "CronJobs": 0,
      "StatefulSets": 0,
      "DaemonSets": 0,
    }
  },
  "data": {
    "ReplicaSets": [],
    "Deployments": [{
      "namespace": "default",
      "workload_kind": "Deployment",
      "workload_label_name": "example",
      "workload_position": 0,
      "rule": "K8S_SEC_RUNASUSER_UNSET",
      "message": "RunAsUser not set",
      "suggestion": "Set RunAsUser",
    }],
    "Pods": [],
    "Jobs": [],
    "CronJobs": [],
    "StatefulSets": [],
    "DaemonSets": [],
  },
}
```

#### Response 400 - Bad Request
```json
{
  "message": "Empty file"
}
```
