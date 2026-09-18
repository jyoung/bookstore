# Bookstore Project

This repository contains the `books-api` Go service and a PostgreSQL database setup managed using Podman Compose.

---

## Prerequisites

- **Podman** and `podman compose` installed.
- Ensure the environment configuration file exists. Copy the example file if needed:
  ```bash
  cp books-api/.env.example books-api/.env
  ```
  *(Make sure `DB_USER`, `DB_PASSWORD`, and `DB_NAME` match your desired local credentials).*

---

## Running the Containers

### Development Environment

In development mode, Podman Compose automatically loads `compose.yaml` and `compose.override.yaml`. This builds the `dev` stage from `books-api/Dockerfile`, launching the application under the [Delve](https://github.com/go-delve/delve) debugger (`dlv`) on port `40000` alongside HTTP on port `8080`.

- **Start (build and run in foreground):**
  ```bash
  podman compose --env-file ./books-api/.env up --build
  ```
- **Start in background (detached mode):**
  ```bash
  podman compose --env-file ./books-api/.env up --build -d
  ```

---

### Production Environment

In production mode, specify `-f compose.yaml` to bypass `compose.override.yaml`. This builds the `prod` multi-stage target (minimal Alpine image running the compiled binary directly without Delve) and exposes HTTP on port `8080`.

- **Start (build and run in foreground):**
  ```bash
  podman compose -f compose.yaml --env-file ./books-api/.env up --build
  ```
- **Start in background (detached mode):**
  ```bash
  podman compose -f compose.yaml --env-file ./books-api/.env up --build -d
  ```

---

## Stopping and Deleting Instances

### 1. Stop Containers (keep container instances and data)
- **Development:**
  ```bash
  podman compose stop
  ```
- **Production:**
  ```bash
  podman compose -f compose.yaml stop
  ```

### 2. Stop and Remove Containers & Networks
- **Development:**
  ```bash
  podman compose down
  ```
- **Production:**
  ```bash
  podman compose -f compose.yaml down
  ```

### 3. Stop and Remove Containers, Networks, and Database Volumes (Full Reset)
To completely delete the database volume (`db_data`) and start fresh:
- **Development:**
  ```bash
  podman compose down -v
  ```
- **Production:**
  ```bash
  podman compose -f compose.yaml down -v
  ```

*(Optional) To also remove the built container images, add `--rmi all` to `down`.*

---

## Debugging in GoLand (Attach Remote Debugger)

When running the **Development** environment, the Go application runs inside the container managed by Delve headless server listening on port `40000`.

To connect and debug from GoLand:

1. **Start the Development containers:**
   ```bash
   podman compose --env-file ./books-api/.env up --build
   ```
2. **Create a Run/Debug Configuration in GoLand:**
   - In the top menu, go to **Run** > **Edit Configurations...** (or click the configuration selector dropdown and choose *Edit Configurations*).
   - Click the **`+`** (Add New Configuration) button in the top left.
   - Select **Go Remote**.
   - Configure the parameters:
     - **Name:** `Remote Debug (Podman)`
     - **Host:** `localhost` (or `127.0.0.1`)
     - **Port:** `40000`
   - Click **Apply** and **OK**.
3. **Attach the Debugger:**
   - Select your new `Remote Debug (Podman)` configuration.
   - Click the **Debug** icon (bug button) or press `Shift + F9`.
   - GoLand will connect to Delve inside the Podman container.
4. **Set Breakpoints and Test:**
   - Set a breakpoint in any handler or repository (e.g., in `books-api/src/main.go` or `books-api/src/authors/repository/repository.go`).
   - Trigger an HTTP request in your browser or terminal:
     ```bash
     curl http://localhost:8080/v1/authors
     ```
   - GoLand will pause execution at your breakpoint, allowing you to inspect variables, step through code, and evaluate expressions.
