# Kubernetes Cheatsheet


## What is kubernetes?
`Kubernetes` is a container orchestration platform used to automate deployment, scaling, management, networking, and the self-healing of containers. It is mainly used to manage multiple docker (or equivalent) containers in multiple nodes (or systems). the goal here is to make cotainerized applications highly available, scalable and have rolling updates.

---

## Kubernetes architecture
A Kubernetes cluster consists of:
`Control Plane / Master Node` - Is responsible for managing the cluster. Consists of **API Server**, **Scheduler**, **Controller Manager**, and **etcd**.

`Worker Node` - This is where containers actually run. Consists of **kubelet**, **kube proxy**, and **Container runtime**

---

## Node
A `Node` is machine in a cluster (physical or VM). A master node manages the cluster and a worker node runs application workloads.

To check nodes:
```bash
kubectl get nodes
```

To get detailed info about a node:
```bash 
kubectl describe node <node-name>
```

---

## Pod
A `Pod` is the samllest deployable unit in kubernetes. Usually consists of one main container and a few helper containers (if necessary). Containers within a pod share storage, network, and IP address. Best practice: `One container per pod`.

To create a pod using a yaml file:
```bash
kubectl apply -f pod.yaml
```

To list pods:
```bash
kubectl get pods
```

Detailed info about a pod:
```bash
kubectl describe pod <pod-name>
```
To delete pod:
```bash
kubectl delete pod <pod-name>
```

To view logs of a pod:
```bash
kubectl logs <pod-name>
```

---

# ReplicaSet
A `ReplicaSet` ensures that a certain number of identical pods are always running. Suppose the replica size is 3, if a pod dies, then the ReplicaSet creates a new one to balance out the number. The replicas are not ordered.

To list ReplicaSets(rs):
```bash
kubectl get rs
```

To apply rs:
```bash
kubectl apply -f rs.yaml
```

To delete rs:
```bash
kubectl delete rs <rs-name>
```

---

## Deployment
A `Deployment` manages ReplicaSets and pods. It provides **rolling updates**, **rollbacks**, **scaling**, and **version control**. It is the recommended way to deploy stateless applications. Unlike a StatefulSet, it does not have ordered replicas.

To apply deployment:
```bash
kubectl apply -f deployment.yaml
```

To list deployments:
```bash
kubectl get deployments
```

To scale a deployment without changing the yaml file:
```bash
kubectl scale deployment <deployment-name> --replicas=5
```

Rolling update:
```bash
kubectl set image deployment/<deployment-name> <container-name>=<image-name>
```

To check rollout status:
```bash
kubectl rollout status deployment/<deployment-name>
```

To rollback:
```bash
kubectl rollback deployment/<deployment-name>
```

---

## Service
A `Service` exposes a pod to other pods, external users or the Internet.
Different types of Services:
1. `ClusterIP` - Internal communication only (default)
2. `NodePort` - Exposes a service on a port of each node
3. `LoadBalancer` - Creates a load balancer based on cloud
4. `ExternalName` - Maps to an external DNS name

To list services:
```bash
kubectl get svc
```

To apply service:
```bash
kubectl apply -f svc.yaml
```

To describe a service:
```bash
kubectl describe svc <svc-name>
```

---

## DaemonSet
A `DaemonSet` ensures that a pod runs on every node. Used in **log collectors**, **monitoing agents**, and **network plugins**

To list DaemonSets:
```bash
kubectl get daemonsets
```

To apply daemonset:
```bash
kubectl apply -f daemonset.yaml
```

---

## Namespace
`Namespaces` help you to segregate and isolate resources in a cluster.

To list all the ns:
```bash
kubectl get ns
```

To create ns:
```bash
kubectl create ns <name>
```

To list all pods in ns:
```bash
kubectl get pods -n <namespace>
```

---