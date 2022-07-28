# Workshopd exercises
1. Create a Rest API which does a simple echo for all the methods (GET POST PUT PATCH DELETE).
The endpoints should return a body with the following format and a 200 OK error code:

```
<HTTP METHOD>: request body
```    

If the request has no body the endpoints should return a 404 error code.

2. Query params vs Path params vs Header params 

Create a route "/tools/repeater" that uses two query params, an intger called "limit" and a string. For this route the service should respond with the string repeated until limit characters are generated.

Create another route "/tools/fibo/{n}" which should return the n-th fibonacci number in the body. <br>
Add a query parameter to this route, it should be a boolean named "cache". The default value should be false if the parameter is missing. If the parameter is set the fibonacci value should come from an internal cache instead of being computed on demand. You should save all new values in the cache regardless of the setting of the cache parameter.

Add a new custom header "USER-IDENTIFIER". If the header is not set, generate a UUID and add it as a header on the response. If the header is set, you should retain the user's preferences. Create an endpoint for a user to opt-out from the caching of their data.

3. Full api (all the methods mentioned at 1.) which holds an in-memory "DataBase" with books.

A book has the following properities: <br>
title, author, publication year, number of downloads, tag list

Each book should have an ID computed by the hashing of the title, author and publication year. Usual algorithms are sha256 or md5. <br>
Hint: there are packages for each of those



