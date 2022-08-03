Observation run the service stack using `docker-compose up`.
In order to generate traffic for `api-1` and `api-2` use the following:

```bash
# Method can be POST, GET, PUT, PATCH
curl -X <method> http://localhost:18081/helloworld # For api-1
curl -X <method> http://localhost:18082/helloworld # For api-2
```