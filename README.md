# Go Calculator with Multi-Stage Build and Distroless Image👩‍💻

This repository demonstrates how to containerize a simple Go calculator application using a multi-stage Docker build and a distroless base image.
This approach results in a significantly smaller and more secure container image.

## Introduction

This project explores optimizing Docker images for size and security. 
By using a multi-stage build and a distroless image, the final image size is reduced dramatically, from ~300MB to ~3.8MB! 
This is achieved by separating the build environment from the runtime environment and including only the necessary application binary in the final image.

## What is a Multi-Stage Build?

A multi-stage build in Docker allows you to use multiple `FROM` statements in a single Dockerfile.
Each `FROM` represents a stage of the build process. This is useful for creating leaner images by separating the build and runtime environments. 
You can use heavy tools and dependencies in one stage to compile or build your application, and then copy only the necessary artifacts to a clean, minimal image in the next stage.

## What is a Distroless Image?

A distroless image contains only the application and its runtime dependencies, with no package manager, shell, or other unnecessary software. 
It is designed to run applications in production environments where security and efficiency are priorities. 
Distroless images have minimal attack surfaces because they don’t contain unnecessary utilities or libraries, reducing potential vulnerabilities.

## Build the Docker image:
docker build -t go-calculator .

## Run the container:
docker run --rm go-calculator 20 - 5  # Example: 20 - 5

## Why static:nonroot?
static: Contains only statically linked libraries, reducing dependencies.
nonroot: Runs the application as a non-root user, improving security.

## Key Takeaways
-Multi-stage builds drastically reduce image size by separating build and runtime environments.
-Distroless images improve security (no shell, fewer vulnerabilities).
-While other lightweight images like Alpine Linux exist, distroless offers a smaller attack surface.
=Understanding docker run flags helps in debugging and container lifecycle management.

## Final Image Size!

[Screenshot 2025-01-29 144930](https://github.com/user-attachments/assets/10fd16df-c3c7-42bb-816e-bd4e3504a1a7)

The final image size is approximately 3.8MB.
