# Databases Workshop

## Prerequisites
1. Docker
2. go v1.17+
3. make

### For Windows Users
Install the following:
1. Install WSL2 driver: https://docs.docker.com/desktop/install/windows-install/
2. Docker (Docker Desktop; Docker Engine - Community 20.10.6)
3. GnuWin32 (contains `make` utility)
4. PostgreSQL 14 (client only) 

> **_NOTE_** Please note if you are unable to run make on your local machine, consider downloading the VM from: https://ctipub-my.sharepoint.com/:u:/g/personal/andrei_albisoru_stud_acs_upb_ro/Eb7uqByEaalNrzALgdeTjpkBhwXFSkbIc20TZyaZscKuyA?e=DxPa8H

## Tasks

### Task 1: Environment setup
1. To build and run the postgres docker image, run `make all`
2. To cleanup the docker container and image created at `1.` `make clean`
3. Download GO modules using `go mod download`
4. Run `go build` in the root directory `lab08`
5. ensure there are no compilation errors and you are able to run the `lab08` binary successfully

### Task 2.a.: Run CRUD commands from psql

#### Use the `psql` tool to perform DB operations
Connect to the Postgres DB: 
```
psql -h 0.0.0.0 -p 5432 -U upb
```
Perform the following `select` operation:
```
upb=> select * from actors                                                                                                                                                                                    ;
                  id                  |        name        |                                  quote
--------------------------------------+--------------------+--------------------------------------------------------------------------
 001                                  | Keanu Reeves       | Sometimes life imitates art
 002                                  | Lawrence Fishburne | People think that I'm haughty and stuck up, but really I'm just very shy
 003                                  | Leonardo DiCaprio  | Only you and you alone can change your situation
(3 rows)

upb=>
```
##### Read by ID
Select a given actor with the provided ID:
```
select * from actors where id='001';
```
##### Update by ID
For the above actor perform the update operation and check for changes:
```
update actors set quote='I have no idea what I am doing' where id='001';
UPDATE 1
upb=> select * from actors where id='001';
                  id                  |     name     |             quote
--------------------------------------+--------------+--------------------------------
 001                                  | Keanu Reeves | I have no idea what I am doing
(1 row)

upb=>
```
##### Create a new record
Create a new actor and check that it got succesfully created in the DB:
```
insert into actors(id, name, quote) values('0010', 'Daniel Craig', 'The name is Bond. James Bond.');
INSERT 0 1
upb=> select * from actors where id='0010';
                  id                  |     name     |             quote
--------------------------------------+--------------+-------------------------------
 0010                                 | Daniel Craig | The name is Bond. James Bond.
(1 row)

upb=>
```
##### Delete the previous record
Delete the previously created actor and check again the DB:
```
upb=> DELETE from actors where id='0010';
DELETE 1
upb=> select * from actors where id='0010';
 id | name | quote
----+------+-------
(0 rows)

upb=>
```

#### Now, we will perform the same operations in GO
First of all, we are using a `driver` to connect to the local Postgres DB. In order to do so, we have a `connection string` that contains all the necesary information. So let's inspect the function call `sql.CreatePostgresConnection` from `main.go`

Let's analyze the movies interface found in `domain/types.go`:
```
type Movies interface {
	CreateMovie(ctx context.Context, movie Movie) error

	GetMoviesByID(ctx context.Context, ids []string) ([]Movie, error)
	ListMovies(ctx context.Context, limit, offset int) ([]Movie, error)

	DeleteMovie(ctx context.Context, id string) error
}
```
now
complete each `TODO` specified in the `main.go` file:
- create 10 favourite movies. what happens if we specify an empty array of actor ids?
- check if they appear in the DB
- remove the previously 10 created movies
- check if the DB does not contain them

### Task 2.b.: Implement the actors repository
The `Actors` interface has the same operations as the `Movies` repositories, handling the same `CRD` operations.
Implement the following methods:
- `CreateActor(ctx context.Context, actor domain.Actor) error`: creates a new actor
- `GetActorsByID(ctx context.Context, ids []string) ([]domain.Actor, error) `: retrieves an actor by it's ID
- `DeleteActor(ctx context.Context, id string) error `: removes an actor from the DB
- `ListActors(ctx context.Context, limit, offset int) ([]domain.Actor, error)`: retrieves a list of actors based on the `limit` and `offset` parameters
> **_NOTE_** You should keep in mind that list operations does a full table scan, so that is why we are using `limit` and `offset` parameters to control the number of iterated records. Imagine what would happen if we would have 3M+ records in the DB, trying to fetch them in a single request.

#### Solution
Please try to solve the problem on your own, before checking the solution:
```
package sql

import (
	"context"
	"database/sql"
	"fmt"
	"lab08/domain"
	"strings"

	"github.com/alecthomas/log4go"
	"github.com/lib/pq"
)

const (
	sqlActorsCreateStmt = `INSERT INTO actors (id, name, quote)
					VALUES ($1, $2, $3) 
					RETURNING id, name, quote`
	sqlActorsDeleteStmt = `DELETE FROM actors WHERE id = $1 
					RETURNING id, name, quote`
	sqlActorsGetByIDStmt = `SELECT id, name, quote
					FROM actors 
					WHERE id = ANY ($1)`
	listActorsStmt = `SELECT id, name, quote
				FROM actors 
				ORDER BY id
				LIMIT $1 OFFSET $2`
	insertActorsBulkStmt        = "INSERT INTO actors (id, name, quote) VALUES %s"
	countActorsPartialMatchStmt = "SELECT count(*) from actors where name like"
	countActorsExactMatchStmt   = "SELECT count(*) from actors where name="
)

type ActorRepository struct {
	db     *sql.DB
	logger *log4go.Logger
}

func NewActorRepository(db *sql.DB, logger *log4go.Logger) *ActorRepository {
	mr := ActorRepository{
		db:     db,
		logger: logger,
	}

	return &mr
}

func (mr *ActorRepository) Count(ctx context.Context, name string, useExactMatch bool) (int, error) {
    var stmt string
    if useExactMatch {
        stmt = countActorsExactMatchStmt + "'" + name + "'"
    } else {
        stmt = countActorsPartialMatchStmt + "'%" + name + "%'"
    }
    
    mr.logger.Info("Running count statement: %s", stmt)
    // TODO: implement me
    
    return 0, nil
}

func (mr *ActorRepository) CreateActor(ctx context.Context, actor domain.Actor) error {
	err := mr.db.QueryRowContext(ctx, sqlActorsCreateStmt, actor.ID, actor.Name, actor.Quote)
	if err != nil && err.Err() != nil {
		return err.Err()
	}

	return nil
}

func (mr *ActorRepository) CreateActors(ctx context.Context, actors []domain.Actor) error {
    // TODO: task 5: index bulk data in the actors table	
	return nil
}

func (mr *ActorRepository) GetActorsByID(ctx context.Context, ids []string) ([]domain.Actor, error) {
	rows, err := mr.db.QueryContext(ctx, sqlActorsGetByIDStmt, pq.Array(ids))
	if err != nil {
		return nil, err
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	defer func() {
		closeErr := rows.Close()
		if closeErr != nil {
			_ = mr.logger.Error(err, domain.FailedRowsErrMsg)
		}
	}()

	var actors []domain.Actor
	for rows.Next() {
		actor := domain.Actor{}

		if err := rows.Scan(&actor.ID, &actor.Name, &actor.Quote); err != nil {
			return nil, err
		}

		actors = append(actors, actor)
	}

	return actors, nil
}

func (mr *ActorRepository) DeleteActor(ctx context.Context, id string) error {
	rows, err := mr.db.QueryContext(ctx, sqlActorsDeleteStmt, id)
	if err != nil {
		return err
	}
	if rows.Err() != nil {
		return rows.Err()
	}
	defer func() {
		closeErr := rows.Close()
		if closeErr != nil {
			_ = mr.logger.Error(err, domain.FailedRowsErrMsg)
		}
	}()

	return nil
}

func (mr *ActorRepository) ListActors(ctx context.Context, limit, offset int) ([]domain.Actor, error) {
	rows, err := mr.db.QueryContext(ctx, listActorsStmt, limit, offset)
	if err != nil {
		return nil, err
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	defer func() {
		closeErr := rows.Close()
		if closeErr != nil {
			_ = mr.logger.Error(err, domain.FailedRowsErrMsg)
		}
	}()

	var actors []domain.Actor
	for rows.Next() {
		var actor domain.Actor
		if err := rows.Scan(&actor.ID, &actor.Name, &actor.Quote); err != nil {
			return nil, err
		}
		actors = append(actors, actor)
	}

	return actors, nil
}
```

### Task 3: Join Operations on DBs
Keypoints to conisder:
- what happens when listing all the movies from the DB?
- what happens if we pull the actor IDs, when pulling the data from the actors table?
- how can we do better? enter join tables

#### `psql` commands
create a new table:
```
upb=> create table movies_with_actors(id SERIAL PRIMARY KEY, movie_id char(36), actor_id char(36));
CREATE TABLE
upb=> select * from movies_with_actors;
 id | movie_id | actor_id
----+----------+----------
(0 rows)
```
insert rows in the table matching the entries from the `movies` table:
```
insert into movies_with_actors(movie_id, actor_id) values('001', '001');
INSERT 0 1
upb=> select * from movies_with_actors;
 id |               movie_id               |               actor_id
----+--------------------------------------+--------------------------------------
  1 | 001                                  | 001
(1 row)
```
perform the inner join:
```
upb=> SELECT m.id, m.name, m.description, a.id, a.name, a.quote from MOVIES AS m INNER JOIN movies_with_actors AS mwa on m.id=mwa.movie_id INNER JOIN actors AS a ON mwa.actor_id=a.id ORDER BY mwa.id;
                  id                  |   name    |    description     |                  id                  |        name        |                                  quote
--------------------------------------+-----------+--------------------+--------------------------------------+--------------------+--------------------------------------------------------------------------
 001                                  | Matrix    | Best movie         | 001                                  | Keanu Reeves       | I have no idea what I am doing
 001                                  | Matrix    | Best movie         | 002                                  | Lawrence Fishburne | People think that I'm haughty and stuck up, but really I'm just very shy
```

#### `GO` implementation
drop the previously created table, as we will create it from the `dbmigrator/movies_with_actors_repository.go`:
```
drop table 

```
as previously 
observed with `psql` we have manually inserted some rows in the DB, so we would need to the migration. 
For this task, we need to implement the following interface:
```
type ExtendedMovies interface {
	MigrateMovies(ctx context.Context) error
	ListAllMovies(ctx context.Context, limit, offset int) ([]MovieExt, error)
}
```
#### Task 3.a.: 
For `MigrateMovies()` the steps are:
- create the table
- run a full table scan over the `movies` table
- for each batch, do a bulk insert in the `movies_with_actors` table
- check that the table `movies_with_actors` is properly populated
- try to create with `psql` an entry that has an invalid/missing actor ID from the `actors` table
- now, fill out the concrete implementation in the `dbmigrator/movies_with_actors_repository.go`. The `MoviesWithActorsRepository` implements the `ExtendedMovies` interface

#### Solution
Please try to solve the problem on your own, before checking the solution:
```$xslt
func (mwa *MoviesWithActorsRepository) MigrateMovies(ctx context.Context) error {
	_, err := mwa.db.ExecContext(ctx, createTableStmt)
	if err != nil {
		return err
	}

	for offset := 0; ; offset += mwa.batchSize {
		movies, err := mwa.movies.ListMovies(ctx, mwa.batchSize, offset)
		if err != nil {
			return err
		}
		if len(movies) == 0 {
			return nil
		}

		var valueArgs []string
		for _, movie := range movies {
			for _, actorID := range movie.Actors {
				valueArgs = append(valueArgs, fmt.Sprintf("('%s', '%s')", movie.ID, actorID))
			}
		}

		stmt := fmt.Sprintf(insertBulkStmt, strings.Join(valueArgs, ","))
		_, err = mwa.db.ExecContext(ctx, stmt)
		if err != nil {
			return err
		}
	}
}
```

#### Task 3.b.:
For `ListAllMovies()` the steps are:
- run the select statement with the inner joins. why do we need limit and offset?
- use an intermediary struct to read all the rows
- after reading all the rows, emit the final array with objects containing a full joined data

##### Solution:
```
func (mwa *MoviesWithActorsRepository) ListAllMovies(ctx context.Context, limit, offset int) ([]domain.MovieExt, error) {
	rows, err := mwa.db.QueryContext(ctx, selectInnerJoinStmt, limit, offset)
	if err != nil {
		return nil, err
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	defer func() {
		closeErr := rows.Close()
		if closeErr != nil {
			_ = mwa.logger.Error(err, domain.FailedRowsErrMsg)
		}
	}()

	actorsWithMovies := make(map[string]map[string]movieWithActorRow)
	for rows.Next() {
		var tmpRow movieWithActorRow
		if err := rows.Scan(&tmpRow.movieID, &tmpRow.movieName, &tmpRow.movieDescription, &tmpRow.actorID, &tmpRow.actorName, &tmpRow.actorQuote); err != nil {
			return nil, err
		}

		actors, found := actorsWithMovies[tmpRow.movieID]
		if found {
			_, found = actors[tmpRow.actorID]
			if !found {
				actors[tmpRow.actorID] = tmpRow
			}
		} else {
			actorsWithMovies[tmpRow.movieID] = make(map[string]movieWithActorRow)
			actorsWithMovies[tmpRow.movieID][tmpRow.actorID] = tmpRow
		}
	}

	var movies []domain.MovieExt
	for _, actorsMap := range actorsWithMovies {
		var m domain.MovieExt
		for _, row := range actorsMap {
			a := domain.Actor{
				Name:  row.actorName,
				Quote: row.actorQuote,
			}
			m.Name = row.movieName
			m.Description = row.movieDescription
			m.Actors = append(m.Actors, a)
		}

		movies = append(movies, m)
	}

	return movies, nil
}
```

#### Task 4.a.: Generate bulk random data
Implement `Generate(ctx context.Context, numRecords int) error` in the `RandomActorGenerator`.

*Hints*:
- start multiple goroutines (one for each worker assigned)
- split the number of generated requests evenly per each worker
- use a fanout approach and call the `createNewActor()` to randomly generate actor data
- when the `batch` number of actors is generated, do a bulk index operation in the actors table
- replicate same behavior on bulk indexing 

#### Task 4.b.: Index bulk data in `actors` table
Implement the `CreateActors(ctx context.Context, actors []domain.Actor) error` in the `ActorRepository`.

> **_NOTE_** Based on the number of records that you want to generate, this may take some time, especially if leaving the input of 3M records. Also, be aware on the number of workers assigned to do the work, as the Postgres server will drop connections, based on the number of configured pool connections, as some connections slots are reserved for admin purposes. This can cause some workers to fail execution, leading to and incomplete number of generated records.

#### Task 5: Performance profiling
We will run 2 types of queries:
- exact matches, where we will be looking up the `actors` table for all actors that have the given name
- partial matches, where we will be looking for only a substring of the actor name
We will then benchmark, the overall latency of these operations

##### Task 5.a.:
Implement the method `Count(ctx context.Context) (domain.GeneratorCounter, error)` of the `RandomActorGenerator`. <br>
*Hints*:
- Randomly generate a first name of the actor, and the last name.
- Use `Actors` interface method `Count()` for partial and exact matches.
- Store the overall latency for each count operation performed on the target repository:
```
t := time.Now()
fn()
latency = time.Now().Sub(t)
fmt.Println(latency)
```

##### Task 5.b.:
Implement the method `Count(ctx context.Context, name string, useExactMatch bool) (int, error) ` for the `ActorRepository`.<br>
*Hints*:
- execute the statement as we did for other operations
- note that the query returns a row, and we need to parse an integer (the count value).

Now we can see different latency times for several queries ran against the DB. Can we run these queries in parallel?

#### Task 6: Improving performance
We can improve search run time, by using a database index. In order to do so, we need to recreate the docker container.<br>
`psql` does not allow creating an index over an existing table, only if we have admin privileges.<br> 
> **_NOTE_** Indexes can be created over an existing table, but the downside is that writes are disallowed during this operation. We can allow writes, but the server will have to perform another full table scan to index the new data.<br> 
Reads however are allowed. <b>Thus</b>, it is best recommened to create an index, on table creation.<br>
Writes, on the other hand, are a bit slower.

*Hints*:
- recreate the docker image: run `make clean` from another terminal in the `lab08` directory
- add an index to the `actors` table in the `schema.sql` file
- recreate the docker image: `make all` in the `lab08` directory
- run the full execution of the lab

> **_NOTE_** Notice the improvement on the read side for exact matches. Any ideas why the partial match did not improve? <br>
Can we also see if the write time for writing in the DB increased now?