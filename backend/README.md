# Backend technical tests

At Adfinitas, we love Go. It's very resilient and safe. APIs written in Go are easily secured. We use it with Google Cloud Run and GC App Engine.

## Tests

For that exercise, we will tests your knowledge to write efficient go code, to refractor an existing codebase and read docs.

Be careful of everything : we let bugs, memory leaks or bad codes in the code. Pay attention to every potential cases. Some changes (not written here) may be expected to be done for maintainability and readability.

## Exercises

We advice you to start from the exercise 1 to n as they are sorted by simplicity.

### 1 - Finish the implementation of the displayAllRoutes endpoint

The `/displayAllRoutes` route must return a JSON with all the routes of the API, find a way on how to do that.

### 2 - Return the result of the Fibonacci sequence

The `/computeFibonacci` route is not returning the result of the computation, find a way to return it **from the goroutine**.

### 3 - Create a CRUD of `contacts`

You must create a full CRUD endpoints for the `contacts` in the `/contacts` path.

A contact is defined as follow:
```json
{
    "name": "string",
    "email": "string",
    "createdAt": "datetime",
    "updatedAt": "datetime"
}
```

## Haven't you miss anything ?

In all development there is one important thing not done here, find what and try to do it for that project.
