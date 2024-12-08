# Tokye Auth Server

The **Tokye Auth Server** is a robust and lightweight authentication server designed for the [Tokye](https://github.com/BenDerFarmer/tokye-js) project.
It provides a seamless and secure user authentication experience, leveraging modern standards such as passkeys and passwordless authentication.

This service utilizes **JSON Web Tokens (JWT)** to ensure secure and easy-to-verify authentication, making it compatible across various backend systems.

With a focus on simplicity, the Tokye Auth Server offers a user-friendly interface for authentication while maintaining high security standards. By integrating passwordless authentication, it eliminates the need for traditional passwords, enhancing both security and user experience.

## Features

- **Passwordless Authentication:** Simplifies user interaction by replacing traditional passwords with **passkeys**.
- **JWT-based Authentication:** Generates and verifies **JSON Web Tokens**, ensuring security and compatibility with various backends.
- **Scalable and Efficient:** Built using modern, lightweight technologies designed for performance and scalability.
- **WebAuthn Integration:** Uses **simplewebauthn** to support WebAuthn standards for secure and seamless user authentication.

## Tech Stack

The Tokye Auth Server is powered by a modern and efficient tech stack:

- **PostgreSQL**: A reliable and powerful relational database for storing user data and authentication metadata.
- **Redis**: High-performance in-memory data storage, used for caching and session management.
- **Go**: A fast, efficient, and scalable programming language that powers the core of the server.
- **FileBased**: A file-based router for the Echo web framework, making it easy to manage and modify large routing configurations
- **JWT**: JSON Web Tokens for easy-to-verify, stateless authentication across backend systems.
- **simplewebauthn**: A library for implementing WebAuthn standards, enabling strong passwordless authentication.

## Getting Started

To get started with the Tokye Auth Server, follow the steps below:

1. Clone the repository:

   ```bash
   git clone https://github.com/BenDerFarmer/tokye-auth-server.git
   cd tokye-auth-server
   ```

2. Configure your environment:

   - Set up your **PostgreSQL** and **Redis** instances.
   - Update the configuration file with your database and Redis connection details.

3. Build and run the server:

   ```bash
   go build
   ./tokye-auth-server
   ```

4. Integrate it with the [Tokye Client](https://github.com/BenDerFarmer/tokye-js) or your existing applications.

## Why Tokye Auth Server?

With the increasing need for secure and user-friendly authentication solutions, Tokye Auth Server bridges the gap by offering:

- Simplified authentication with **passwordless login**.
- High security through the use of WebAuthn and **modern cryptographic standards**.
- A developer-friendly approach with well-documented APIs and easy integration options.

Whether you're building a new project or enhancing an existing one, the Tokye Auth Server makes secure authentication effortless.

## Environment Variables

The Tokye Auth Server relies on several environment variables for configuration. These variables allow you to customize the server's behavior and integrate it with your infrastructure. Below is a breakdown of each environment variable and its purpose:

### 1. **Secret Configuration**

- **`REFRESH_TOKEN_SECRET`** _(required)_:  
  A secret key used to sign and verify refresh tokens. Ensure this is a strong, randomly generated string to maintain token security.

- **`SQL_DSN`** _(required)_:  
  The Data Source Name (DSN) for connecting to your PostgreSQL database. Format example: `host=127.0.0.1 user=gorm password=mysecretpassword dbname=gorm port=5432 sslmode=disable TimeZone=Europe/Berlin`.

### 2. **SMTP Configuration**

These variables configure the SMTP server for sending emails, such as for account verification or passwordless login links.

- **`SMTP_DOMAIN`** _(required)_:  
  The domain name of your SMTP service provider.

- **`SMTP_PORT`** _(default: `465`)_:  
  The port used to connect to the SMTP server. Typically `465` for SSL or `587` for STARTTLS.

- **`SMTP_USERNAME`** _(required)_:  
  The username or email address used to authenticate with the SMTP server.

- **`SMTP_PASSWORD`** _(required)_:  
  The password for your SMTP account.

- **`SMTP_FROM_MAIL`** _(default: `SMTP_USERNAME`)_:  
  The email address used in the "From" field of outgoing emails. If not provided, it defaults to the SMTP username.

### 3. **Redis Configuration**

These variables configure the Redis instance for caching and session management.

- **`REDIS_ADDRESS`** _(default: `localhost:6379`)_:  
  The address and port of your Redis instance.

- **`REDIS_USERNAME`** _(optional)_:  
  The username for authenticating with Redis (if applicable).

- **`REDIS_PASSWORD`** _(optional)_:  
  The password for authenticating with Redis.

- **`REDIS_DATABASE`** _(default: `0`)_:  
  The Redis database index to use.

### 4. **Passkey Configuration**

These variables configure settings for passkey-based authentication.

- **`PASSKEY_DISPLAYNAME`** _(default: `Auth`)_:  
  The display name of the relying party (RP), shown to users during passkey authentication.

- **`PASSKEY_RPID`** _(required)_:  
  The Relying Party Identifier (RPID) for passkey operations. Typically, this is your application's domain (e.g., `example.com`).

- **`PASSKEY_ORIGINS`** _(default: empty)_:  
  A semicolon-separated list of allowed origins for passkey operations (e.g., `https://example.com;https://app.example.com`).

### 5. **CORS Configuration**

- **`CORS_ORIGINS`** _(default: empty)_:  
  A semicolon-separated list of allowed origins for Cross-Origin Resource Sharing (CORS). Example: `http://localhost:3000;https://example.com`.

### 6. **Miscellaneous**

- **`DEBUG_MODE`** _(default: `false`)_:  
  Enables debug logging when set to `true`. Useful during development to identify issues.

- **`PORT`** _(default: `3000`)_:  
  The port on which the server will run. Example: `8080` for production.

### Example `.env` File

```plaintext
REFRESH_TOKEN_SECRET=vFNjHjFccr%fUN6A%Uwb4TM5UA!mSD^NuW3*nEWk*EM^WQjK9K
SQL_DSN="host=127.0.0.1 user=gorm password=mysecretpassword dbname=gorm port=5432 sslmode=disable TimeZone=Europe/Berlin"
SMTP_DOMAIN=smtp.example.com
SMTP_PORT=465
SMTP_USERNAME=auth@example.com
SMTP_PASSWORD=securepassword
SMTP_FROM_MAIL=auth@example.com
REDIS_ADDRESS=localhost:6379
REDIS_USERNAME=
REDIS_PASSWORD=
REDIS_DATABASE=0
PASSKEY_DISPLAYNAME=Tokye Auth
PASSKEY_RPID=example.com
PASSKEY_ORIGINS=https://example.com;https://auth.example.com
CORS_ORIGINS=http://localhost:3000;https://example.com
DEBUG_MODE=true
PORT=3000
```
