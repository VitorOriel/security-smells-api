# Kube Smelly API

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
      "rule": "K8S_SEC_RUNASUSER",
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

#### Response 401 - Bad Request
```json
{
  "message": "Empty file"
}
```