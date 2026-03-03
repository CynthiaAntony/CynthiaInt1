# Container Reference

---

## What is a Container?

A container is an isolated environment that has its own processes, networks, and mounts, but shares the same OS kernel with other containers on the host. It packages an application together with all its dependencies so that it runs consistently in any environment. Containers behave like independent entities (isolated processes, networks, mounts), but are far lighter than VMs since they don't carry a full OS.

---

## Image vs Container

| Concept | Description |
| --- | --- |
| `Image` | A read-only blueprint of a container. Built from a Containerfile or pulled from a registry (e.g. Docker Hub). |
| `Container` | A live, running instance of an image. |

The image defines the structure, the container is the thing actually doing work.

---

## Containerflow

Every container follows a 3-step lifecycle:

1. `Write` — Author a Containerfile that describes your environment, dependencies, and app setup.
2. `Build` — The container engine reads the Containerfile and compiles it into an image.
3. `Run` — The image is instantiated into a live, running container.

---

## Containerfile

A Containerfile is a plain text document containing layered instructions that tell the container engine how to build an image. Each instruction adds a new layer on top of the last.

### Commands

| Command | Description |
| --- | --- |
| `FROM` | Sets the base image to build on top of (like FROM golang:1.21). Every Containerfile must start with this. |
| `WORKDIR` | Sets the working directory inside the container. All subsequent commands run from here. |
| `COPY` | Copies files or folders from your local machine into the container's filesystem. |
| `RUN` | Executes shell commands during the build phase and used for installing packages, compiling code, etc. |
| `ENV` | Sets environment variables inside the container. |
| `EXPOSE` | Documents which port the container listens on. Informational — does not actually publish the port. |
| `CMD` | The default command to run when the container starts. Can be overridden at runtime. |
| `ENTRYPOINT` | Configures the container to behave like an executable. Less easily overridden than CMD. |

### Example Containerfile

```dockerfile
FROM golang:1.21

WORKDIR /app

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
```

---

## Podman

### What is Podman?

Podman is a daemonless container engine used for building, managing, and running containers. Unlike Docker, it requires no background daemon process running with root privileges, making it more secure by default. It is fully compatible with Docker images and commands.

### Flag Reference

| Flag | Meaning |
| --- | --- |
| `-t` | Tag — gives the image a name |
| `-d` | Detached — runs the container in the background |
| `--name` | Assigns a human-readable name to the container |
| `-p <host>:<container>` | Maps a host port to a container port |
| `-f` | Follow — streams logs in real time |

---

### Commands

```bash
#Build
podman build -t my-app .

#Run (detached, named, with port mapping)
podman run -d --name app-instance -p 8080:8080 my-app

#Inspect
podman images #to list images
podman ps #to list running containers
podman logs -f <name> #to stream logs

#Cleanup
podman stop <name> #to stop running a container
podman rm <name> #to remove a container
```