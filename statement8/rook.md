# Rook Cheatsheet

---

## What is Rook?
Rook basically brings Ceph storage into a Kubernetes cluster. It is open-source and Ceph storage with an operator and CRDs (custom resource definition). It automates deployment, updates and configuration. It allows apps to use storage like any other kubernetes storage.

---

## Rook setup:

1. Create a virtual machine (qemu is used to run a stable a linux vm) with 8GB RAM, 4 CPUs, and an extra disk (since rook needs a raw disk) for storage. Container runtime is **containerd** since it is faster and uses less RAM comapared to Docker
```bash
minikube start -p rook-ceph --driver=qemu --network=builtin --extra-disks=1 --container-runtime=containerd --memory=8192 --cpus=4
```

2. Move to examples folder
```bash
cd ~/rook/deploy/examples
```

3. Rook orchestrator - **crds.yaml** installs custom resource definitions, **common.yaml** creates ceph namespace and the RBAC rules, **csi-driver.yaml** deploys the csi-operator and which in turn manages the csi-driver, and **operator.yaml** installs ceph
```bash
kubectl create -f crds.yaml -f common.yaml -f csi-operator.yaml -f operator.yaml
```

4. Run this to watch the pod (when the pods are running, run the next command)
```bash
kubectl -n rook-lab get pod -w
```

5. Tell the operator to build the storage
```bash
kubectl create -f cluster-test.yaml
```
Post this, run command 4 again and wait till it shows `rook-ceph-osd-0`, then follow with allocating a PVC or anything else

---