# Go Product API Server
This server provides an API for a product test data.

Valid routes:

/products  
/product/{id}  

### Prequisites

A running mogodb server on localhost.

### Running the server
Load sample data:

```
mongoimport -c products -d products --file sample/products.json --jsonArray
```

To run the API server:

```
cd server
go run main.go
```

