### Tasks
1. `GET, POST, PATCH, DELETE` in memory in api-service using a single id/product as input
2. `GET, POST, PATCH, DELETE` in memory in api-service using a batch of ids/products as input
3. `GET, POST, PATCH, DELETE` over http (so passing the request to storage-server) in api-service using a single id/product as input. Storage server should save the data in an database (postgres or any other you’d like to)
4. `GET, POST, PATCH, DELETE` over http (so passing the request to storage-server) in api-service using a batch ids/products as input. Storage server should save the data in an database (postgres or any other you’d like to)
5. Async endpoints for `POST, PATCH, DELETE` using a redis queue

## Points system:
1. Functionality - the amount of correctly working endpoints
2. Efficiency - no goroutines for batch processing/unbounded fan out/bounded fan out
3. Coding style (var names, code comments - especially for public methods, fields, types, ease of reading - no `if if if if`)
4. Unit tests

## After you're done you can give us some feedback here:
https://forms.gle/DuKUXQyE1u6MVa1V8
