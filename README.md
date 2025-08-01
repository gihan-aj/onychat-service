# OnyxChat Service 💬

A simple, real-time TCP chat server built with Go. This project is a learning exercise to explore Go's powerful concurrency features (goroutines and channels) for handling networked applications.

## ✨ Features

* **Multi-Client Support:** Handles multiple concurrent client connections.
* **Real-time Broadcasting:** Messages sent by one client are instantly broadcast to all other connected clients.
* **Connection Status:** Announces when a user joins or leaves the chat.

## 🛠️ Tech Stack & Core Concepts

* **Language:** [Go](https://go.dev/)
* **Protocol:** TCP
* **Core Go Concepts:**
    * `net` package for TCP networking.
    * **Goroutines** for handling each client connection concurrently.
    * **Channels** for safe communication between goroutines and the central broadcaster.
    * **Select** statement for managing multiple channel operations.

## 🚀 Getting Started

### Prerequisites

* Go installed on your machine.
* A TCP client like `telnet`.

### How to Run

1.  **Clone the repository:**
    ```sh
    git clone [https://github.com/gihan-aj/onychat-service.git](https://github.com/gihan-aj/onychat-service.git)
    cd onychat-service
    ```

2.  **Run the server:**
    ```sh
    go run main.go
    ```
    The server will start listening on `localhost:8080`.

3.  **Connect and Chat:**
    * Open multiple terminal windows.
    * In each terminal, connect to the server using `telnet`:
        ```sh
        telnet localhost 8080
        ```
    * Start sending messages!

## 🗺️ Future Roadmap

* Transition from TCP to **WebSockets** for browser-based clients.
* Implement user authentication by integrating with a **SSO/user management service**.
* Allow users to set custom **usernames**.
* Persist chat history to a **database**.