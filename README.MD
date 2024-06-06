# [English](https://github.com/linchengzhi/go-clean-backend) | [中文](https://github.com/linchengzhi/go-clean-backend/blob/main/README_ZH.MD)
# Golang Backend Clean Architecture -- Blog Demo

## Introduction
This project is a blog demo developed using Gin, designed to showcase how to implement a clean architecture.

## Clean Architecture
Clean architecture is a software design principle aimed at creating systems that are easy to understand, maintain, and test. Its core focus is on the separation of concerns while ensuring the system remains scalable.

### Key Principles
1. **Framework Independence**: The design does not depend on any specific framework. Changing the framework will not affect the logic. Whether using Gin or Go-zero, swapping frameworks should be straightforward.
2. **Component Independence**: The project uses GORM to interact with MySQL, but it can be easily switched to other databases (e.g., MongoDB, SQLite) without affecting other layers.
3. **Testability**: Business logic in each layer can be tested without relying on the UI, database, or any external components.

### Layered Structure
Clean architecture typically consists of the following four layers:
- Entities
- Usecase
- Controller
- Framework & Driver

In this project, the structure is slightly different and includes:
- Domain
- Repository
- Usecase
- Delivery

### Domain
The Domain layer is similar to the Entities layer. It contains core business entities, data transfer objects, error definitions, common definitions, and constants. Additionally, it includes interface definitions for the Repository and Usecase layers.

### Repository
The Repository layer is responsible for data interactions, including database operations, cache operations, sending notifications, and requesting third-party services. For example, MySQL interactions, gRPC requests to other microservices, and HTTP notifications. Interfaces are defined in the Domain layer to facilitate switching databases (e.g., from MySQL to MongoDB) by implementing the new database's interface without modifying upper layers.

### Usecase
The Usecase layer encapsulates and implements all system use cases. All incoming and outgoing data is processed here before flowing to the Repository layer or returning to the Delivery layer. A use case typically involves an actor or user, an action, and a response. Usecase layer interfaces are generally placed in the Domain layer, but for small projects with minimal changes, they can be placed in this layer.

### Delivery
The Delivery layer, also known as the interface adapter, is mainly used for data transmission and use case triggering, such as HTTP, gRPC, Kafka, and scheduled tasks. For example, in an API call, user-provided parameters are organized into the data format required by the use case, and the data returned by the use case is converted into the response format.

## Directory Structure
```shell
Api        # Corresponds to the Delivery layer, categorized by protocol type; this demo uses HTTP
Build      # Files for building the project, typically includes Dockerfile, Makefile, etc.
Cmd        # Project entry point, contains Main function and initialization code
Config     # Configuration files directory
Docs       # Documentation directory, contains project documentation
Domain     # Contains entity structures, data transfer objects, constants, and error definitions
Infra      # Infrastructure components directory, initialization of components like MySQL, Redis, Log
Repository # Responsible for data interactions, including database operations, cache operations, sending notifications, and requesting third-party services
Usecase    # Usecase layer, can be divided into multiple subdirectories if the logic is complex
```

## Quick Start
1. Modify the MySQL configuration in the configuration file.
2. Run the following command to start the project:
```shell
go run cmd/main.go
```

## License
This project is licensed under the MIT License. For details, please refer to the [LICENSE](./LICENSE) file.