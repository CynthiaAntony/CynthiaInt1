# Solution to statement 6
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
podman build -t trie-server:v1 -f server.Dockerfile .
podman build -t trie-client:v1 -f client.Dockerfile .

# Save to tar and load into Minikube
podman save -o server.tar localhost/trie-server:v1
podman save -o client.tar localhost/trie-client:v1

minikube image load server.tar
minikube image load client.tar
```
4. Deploy to kubernetes
```bash
kubectl apply -f k8s/
```
5. Attach to client terminal to access the running program
```bash
kubectl exec -it trie-client-pod -- ./client_app
```

---

## Example input - statement 6
```bash
Connected to server
---String List Manager---
Enter choice: 1. Add
2. Remove
3. Check
4. List
5. Exit: add
Enter word: hi
Word added
Enter choice: 1. Add
2. Remove
3. Check
4. List
5. Exit: add
Enter word: hello
Word added
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
5. Exit: remove
Enter word: cyn
Word removed
Enter choice: 1. Add
2. Remove
3. Check
4. List
5. Exit: add
Enter word: cynthia
Word added
Enter choice: 1. Add
2. Remove
3. Check
4. List
5. Exit: list
Words : [cynthia hi hello]
Enter choice: 1. Add
2. Remove
3. Check
4. List
5. Exit: check
Enter word: hey
Word does not exist
Enter choice: 1. Add
2. Remove
3. Check
4. List
5. Exit: check
Enter word: hi
Word Exists
Enter choice: 1. Add
2. Remove
3. Check
4. List
5. Exit: exit
Exiting...
```

---

## Example - statement 9
```bash
Connected to server
---String List Manager---
Enter choice: 1. Add
2. Remove
3. Check
4. List
5. Exit: list
Words : []
Enter choice: 1. Add
2. Remove
3. Check
4. List
5. Exit: add
Enter word: hi
Word added
Enter choice: 1. Add
2. Remove
3. Check
4. List
5. Exit: add
Enter word: hello
Word added
Enter choice: 1. Add
2. Remove
3. Check
4. List
5. Exit: list
Words : [hi hello]
Enter choice: 1. Add
2. Remove
3. Check
4. List
5. Exit: check
Enter word: hi
Word Exists
Enter choice: 1. Add
2. Remove
3. Check
4. List
5. Exit: remove
Enter word: hi
Word removed
Enter choice: 1. Add
2. Remove
3. Check
4. List
5. Exit: check
Enter word: hi
Word does not exist
Enter choice: 1. Add
2. Remove
3. Check
4. List
5. Exit: list
Words : [hello]
Enter choice: 1. Add
2. Remove
3. Check
4. List
5. Exit: exit
Exiting...
```

Run this command again: 
```bash
kubectl exec -it trie-client-pod -- ./client_app
```

And you get:
```bash
Connected to server
---String List Manager---
Enter choice: 1. Add
2. Remove
3. Check
4. List
5. Exit: list
Words : [hello] #Persists
Enter choice: 1. Add
2. Remove
3. Check
4. List
5. Exit: exit
Exiting...
```