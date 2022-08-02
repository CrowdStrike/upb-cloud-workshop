### Task 0

Lets start with our setup. Follow the guide mentioned [here](https://docs.docker.com/get-docker/) in order to install it. This should also install the `docker-compose` plugin that we are going to use later. An easier alternative is to create a docker account and use [docker playground](https://labs.play-with-docker.com/) which comes along with all the toolchain required for this lab, only be mindful that your session is limited to 4 hours.

### Task 1

In `tasks` folder you are given an app that you have to monitor using prometheus metrics, dockerize it and create some plots. Try adopting an incremental solving approach:

1. Using [this](https://pkg.go.dev/github.com/prometheus/client_golang/prometheus) package and the information provided [here](https://github.com/prometheus/client_golang/blob/main/prometheus/examples_test.go). Create the metrics mentioned with `TODOs` register them and them create the prometheus handler.

2. Using the dockerfile from the demo try creating a docker image for this app.
For testing try running a container using `docker run` and exposing the port used by your app.

3. Create a docker-compose file similar with what you've seen in the demo today and add the following microservices: `prometheus`, `grafana` and `go-app`.

> **_HINT_** Try reusing the `docker-compose` from the demo. You can eliminate the cAdvisor part.

> **_NOTE_** You will have to modify the prometheus config file for jobs: `./monitoring/prometheus/prometheus.yml` and add a new job for you exported metrics.

4. Try creating some plots in grafana with the metrics that you are exporting.

5. Extra - Using one of the following metrics exported by prometheus itself create some grafana panels:

```
go_memstats_alloc_bytes_total
promhttp_metric_handler_requests_total
```