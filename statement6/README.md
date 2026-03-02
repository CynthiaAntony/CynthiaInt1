# Solution to statement 7
---

String list management system implemented using grpc
The system is split into two components:
`Server`: Maintains the Trie logic in memory and listens for gRPC calls.
`Client`: A CLI-based remote control that takes user input and communicates with the server via Protocol Buffers.

---

## Pre-requisites
1. Go installed - 1.25.7
2. Protobuf compiler
3. Podman and minikube

---

## Setup
1. Start minikube with Podman driver
```bash
minikube start --driver=podman
```
2. Generate grpc code
```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/trie.proto
```
3. Build and load containers into minikube
```bash
# Build images
podman build -t trie-server:latest -f server.Dockerfile .
podman build -t trie-client:v4 -f client.Dockerfile .

# Save to tar and load into Minikube
podman save -o server.tar localhost/trie-server:latest
podman save -o client.tar localhost/trie-client:v3

minikube image load server.tar
minikube image load client.tar
```
4. Deploy to kubernetes
```bash
kubectl apply -f deploy.yaml
```
5. Attach to client terminal to access the running program
```bash
kubectl attach -it trie-client-pod
```

---

## Example input
```bash
Enter choice: 1. Add
2. Remove
3. Check
4. List
5. Exit: add
Enter word: cat 
Word added
Enter choice: 1. Add
2. Remove
3. Check
4. List
5. Exit: list
Words : [cat cynthia antony]
Enter choice: 1. Add
2. Remove
3. Check
4. List
5. Exit: remove
Enter word: cat
Word removed
Enter choice: 1. Add
2. Remove
3. Check
4. List
5. Exit: list
Words : [cynthia antony]
Enter choice: 1. Add
2. Remove
3. Check
4. List
5. Exit: check
Enter word: cynthia
Word Exists
Enter choice: 1. Add
2. Remove
3. Check
4. List
5. Exit: check
Enter word: ant
Word does not exist
Enter choice: 1. Add
2. Remove
3. Check
4. List
5. Exit: remove
Enter word: ant
Word not found
Enter choice: 1. Add
2. Remove
3. Check
4. List
5. Exit: add
Enter word: cyn
Word added
Enter choice: 1. Add
2. Remove
3. Check
4. List
5. Exit: list
Words : [cyn cynthia antony]
Enter choice: 1. Add
2. Remove
3. Check
4. List
5. Exit: list
Words : [cyn cynthia antony]
Enter choice: 1. Add
2. Remove
3. Check
4. List
5. Exit: exit
Exiting...
```

---