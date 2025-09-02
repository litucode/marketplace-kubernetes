# Kubernetes Global Cluster Control API Endpoints

This API allows you to register, manage, and query Kubernetes clusters around the world by assigning human-friendly names and using a database for cluster configuration management.

---

## Cluster Management

### Register a New Cluster

**POST** `/api/clusters`

Registers a new Kubernetes cluster with a human-friendly name and configuration.

**Request Body Example:**

```json
{
  "name": "tokyo-prod",
  "kubeconfig": "<kubeconfig-content>",
  "location": "Tokyo, Japan"
}
```

**Response:**

```json
{
  "id": 1,
  "name": "tokyo-prod",
  "location": "Tokyo, Japan",
  "createdAt": "2025-09-02T04:30:50Z"
}
```

---

### List All Clusters

**GET** `/api/clusters`

Returns a list of all registered clusters.

**Response:**

```json
[
  {
    "id": 1,
    "name": "tokyo-prod",
    "location": "Tokyo, Japan"
  },
  {
    "id": 2,
    "name": "berlin-dev",
    "location": "Berlin, Germany"
  }
]
```

---

### Get Cluster Info

**GET** `/api/clusters/{name}`

Returns detailed metadata about a specific cluster by its human-friendly name.

**Response:**

```json
{
  "id": 1,
  "name": "tokyo-prod",
  "location": "Tokyo, Japan",
  "createdAt": "2025-09-02T04:30:50Z"
}
```

---

## Kubernetes Resource Queries

All resource query endpoints use the cluster's human-friendly name to select the target cluster.

---

### List Namespaces

**GET** `/api/clusters/{name}/namespaces`

Lists all namespaces in the specified cluster.

**Response:**

```json
{
  "namespaces": ["default", "kube-system", "production"]
}
```

---

### List Pods in a Namespace

**GET** `/api/clusters/{name}/namespaces/{namespace}/pods`

Lists all pods in the given namespace for the specified cluster.

**Response:**

```json
{
  "pods": [
    { "name": "nginx-12345", "status": "Running", "node": "node-1" },
    { "name": "redis-67890", "status": "Pending", "node": null }
  ]
}
```

---

### Get Pod Details

**GET** `/api/clusters/{name}/namespaces/{namespace}/pods/{podName}`

Returns detailed information about a specific pod.

**Response:**

```json
{
  "name": "nginx-12345",
  "status": "Running",
  "node": "node-1",
  "containers": [{ "name": "nginx", "image": "nginx:1.21" }],
  "labels": {
    "app": "nginx"
  }
}
```

---

### List Deployments

**GET** `/api/clusters/{name}/namespaces/{namespace}/deployments`

Lists all deployments in a given namespace for the specified cluster.

**Response:**

```json
{
  "deployments": [
    { "name": "web-app", "replicas": 3, "availableReplicas": 3 },
    { "name": "redis", "replicas": 1, "availableReplicas": 1 }
  ]
}
```

---

### List Services

**GET** `/api/clusters/{name}/namespaces/{namespace}/services`

Lists all services in a given namespace for the specified cluster.

**Response:**

```json
{
  "services": [
    { "name": "web-app", "type": "ClusterIP", "clusterIP": "10.0.0.1" },
    { "name": "redis", "type": "NodePort", "nodePort": 30007 }
  ]
}
```

---

## Notes

- All endpoints return errors in the following format:

  ```json
  {
    "error": "Error message describing the issue."
  }
  ```

- Authentication and security are recommended for production use.
- The `kubeconfig` field may be stored securely (e.g., encrypted) in production.
