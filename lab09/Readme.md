# Databases Redis Workshop

## Prerequisites
1. Docker
2. go v1.17+
3. Redis (Client only) https://redis.io/docs/getting-started/installation/install-redis-on-windows/

## Tasks

### Task 1: Environment setup

1. To build and run the redis docker image, run `docker-compose up -d`
2. To cleanup the docker container and image created at `1.` use `docker-compose down` and remove the data directory.
3. Download GO modules using `go mod download`
4. Run `go build` in the root directory `lab09`
5. ensure there are no compilation errors and you are able to run the `lab09` binary successfully

### Task 2: Create and retrieve a key

Create an integer key, increment it atomically. Check if the incrementation was done right.

> Hint: INCR <br>
  More info on redis commands https://redis.io/commands/

Create a string key and check that you can add at least 100 characters. Delete this key after using it.

Add a timeout to a key, and check that it dissapears after 1 minute.

### Task 3: Lists and sets

Select a new database and check that none of the other keys exist there.

Create a list of strings and a list with the first letter in each of the previous lists.

Use a way to update and remove the values in both lists atomically.

Create a set from the second list (with the starting letters) and use it to retrieve all the words in the first list starting with the letter in the set.

### Bonus task
Check out how to implement a pubsub using redis:

https://github.com/gomodule/redigo/blob/master/redis/pubsub_example_test.go



