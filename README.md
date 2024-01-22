# League Backend Challenge

In main.go you will find a basic web server written in GoLang.

Given an uploaded csv file
```
1,2,3
4,5,6
7,8,9
```

1. Echo (given)
    - Return the matrix as a string in matrix format.
    
    ```
    // Expected output
    1,2,3
    4,5,6
    7,8,9
    ``` 
2. Invert
    - Return the matrix as a string in matrix format where the columns and rows are inverted
    ```
    // Expected output
    1,4,7
    2,5,8
    3,6,9
    ``` 
3. Flatten
    - Return the matrix as a 1 line string, with values separated by commas.
    ```
    // Expected output
    1,2,3,4,5,6,7,8,9
    ``` 
4. Sum
    - Return the sum of the integers in the matrix
    ```
    // Expected output
    45
    ``` 
5. Multiply
    - Return the product of the integers in the matrix
    ```
    // Expected output
    362880
    ``` 

The input file to these functions is a matrix, of any dimension where the number of rows are equal to the number of columns (square). Each value is an integer, and there is no header row. matrix.csv is example valid input.  

Run web server
```
go run .
```

Send request
```
curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"
```

## What was implemented

- The solution runs

  To run the web server:
  ```
  go run .
  ``` 

- The code is tested

  Tests are included in matrix folder, to run the test from the command line:
  ```
  make test
  ```  
- The code is robust and handles invalid input and provides helpful error messages
  
  The matrix implementation is generic and allow to specify the number type
  ```
  type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
  }
  ```
  In the current implementation the matrix values types are float64 to handle more than just integers.

  If the matrix values conversion will fail and error message will be returned.

## Code Structure

#### main.go

The server is initialized a separate goroutine so that it won't block and make possible to graceful shutdown.

#### Routes

Routes are defined in routes/routes.go

In addition to required API endpoints some helpful middleware were implemented.
- Global panic recovery
- Random Request ID generation
- Log each request details
- Cache control - make sure that the requested data is not going to be cached by the browser

#### handlers

The API endpoints are implemented in handlers/handlers.go

#### matrix

The matrix manipulations are implemented in matrix/matrix.go and tested in matrix/matrix_test.go
