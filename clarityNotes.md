Before we move on, I see we used this before:
switch r.Method {
	case http.MethodGet:
		getTasks(w, r)
	case http.MethodPost:
		createTask(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

now it is:
switch r.Method {
	case "GET":
		GetTasks(w, r)
	case "POST":
		CreateTask(w, r)
	case "PUT":
		UpdateTask(w, r)
	case "DELETE":
		DeleteTask(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}



Also, do these represent similar as NodeJS's request and response?
w http.ResponseWriter, r *http.Request



========================


Great observations! Let's break down your two questions:

### 1. **Switch on `r.Method`: String Literals vs Constants**
You noticed that in the first case, we used `http.MethodGet`, and in the second case, we used `"GET"`, `"POST"`, etc. These are functionally the same, but there's a key difference:

- **Using string literals** like `"GET"`, `"POST"`, etc., is valid and works perfectly fine. It's simply using the string representation of HTTP methods directly.
  
- **Using constants like `http.MethodGet`** is generally preferred because:
  - It reduces the risk of typos (e.g., accidentally typing `"GETT"` or `"get"` would lead to issues).
  - It makes the code more readable and consistent.
  - It leverages Go's built-in constants from the `net/http` package, which are more robust and less error-prone.
  
In Go, the `http.MethodGet`, `http.MethodPost`, etc., are predefined constants:
```go
const (
    MethodGet     = "GET"
    MethodHead    = "HEAD"
    MethodPost    = "POST"
    MethodPut     = "PUT"
    MethodPatch   = "PATCH"
    MethodDelete  = "DELETE"
    MethodConnect = "CONNECT"
    MethodOptions = "OPTIONS"
    MethodTrace   = "TRACE"
)
```

While both approaches will work, using `http.MethodGet`, `http.MethodPost`, etc., is considered better practice.

#### Updated Example Using Constants:
```go
switch r.Method {
	case http.MethodGet:
		GetTasks(w, r)
	case http.MethodPost:
		CreateTask(w, r)
	case http.MethodPut:
		UpdateTask(w, r)
	case http.MethodDelete:
		DeleteTask(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
```

This approach makes your code more consistent and reduces potential errors due to typos.

### 2. **Understanding `http.ResponseWriter` and `*http.Request`**

Yes, in Go, the `w http.ResponseWriter, r *http.Request` is quite similar to how you handle requests and responses in Node.js, with a few differences in how they're implemented.

- **`http.ResponseWriter (w)`**: This is an interface in Go that is used to send responses back to the client. It's like the `response` object (`res`) in Node.js (`res.send()`, `res.status()`, etc.). You use `http.ResponseWriter` to write the response, set the HTTP status code, headers, and content.

  Example:
  ```go
  w.WriteHeader(http.StatusOK) // Set status code
  w.Header().Set("Content-Type", "application/json") // Set headers
  w.Write([]byte("Hello World")) // Write response body
  ```

  This is similar to:
  ```javascript
  res.status(200).send("Hello World");
  ```

- **`*http.Request (r)`**: This is a pointer to an `http.Request` object that represents the incoming request from the client. It's like the `request` object (`req`) in Node.js (`req.body`, `req.query`, `req.params`, etc.). You use `*http.Request` to read data from the request such as query parameters, headers, body, and more.

  Example:
  ```go
  r.Method // "GET", "POST", etc.
  r.URL.Query().Get("id") // Get query parameter
  ```

  This is similar to:
  ```javascript
  req.method // "GET", "POST", etc.
  req.query.id // Get query parameter
  ```

### Key Differences between Go and Node.js:
- In **Go**, you explicitly pass the response writer and request as parameters (`w http.ResponseWriter, r *http.Request`). Go uses the concept of passing by value or reference to functions, rather than having objects as methods on the `req`/`res` like in Node.js.
  
- In **Node.js**, `req` and `res` are typically available as objects directly within the route handler, without needing to pass them explicitly.

### Summary:
- **Switch on HTTP method**: It's better to use `http.MethodGet`, `http.MethodPost`, etc., for readability and preventing errors.
- **`http.ResponseWriter` and `*http.Request`** in Go are similar to Node.js's `req` and `res`, but Go handles them explicitly as function parameters.
