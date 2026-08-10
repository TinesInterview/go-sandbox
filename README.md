# Tines

This is our sandbox repository, containing a sample React application connected to a Go backend, to ensure candidates can run our coding challenge solution without problems.

## Getting Started

First, install all the required dependencies. You'll need Go 1.24 or newer (https://go.dev/dl/) and Node.js 20 or newer (https://nodejs.org/).

```bash
t/setup
```

Once that's complete, you can run both frontend and backend servers using

```bash
t/start
```

Open [http://localhost:3000](http://localhost:3000) with your browser to see the result.

VSCode users: when you open the project, accept the prompt to install the recommended Go extension for full IntelliSense. The Go code lives in `backend/`.

## Running Tests

To run the backend tests:

```bash
t/test
```
