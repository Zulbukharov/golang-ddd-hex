# Golang DDD+HEX demo blog api

- tactical design
    entities
        an object defined primarily by its identity
    value objects
        immutable
    repositories
        how will communicate with storage
    services
        corresponds application use cases
        example:
        oauth -> session -> user policy -> adding post

context: blogpost
ubiquitous language: author, post, vote, comments, storage
entities: author, post, votes, comments
value objects: it can be part of entity object (author, voter, commenter)
aggregates: PostAuthor, PostVotes, PostComments
service: 
    stateless operations 
        post adder / listing,
        vote increment / decrement / count
        comment adder / listing
        author register / login 
events: can affect to the system (errors, logs)
repository: facade over backend

## Goals: good structure goal

- consistent.
- easy to test.
- easy to understand.
- easy to change.

## Project Specs

- author can add a post.
- author can add a comment for the post.
- author can vote for the post.
- author can list all posts with votes.
- author can list all reviews for the post.


## Dependencies:
- [pgx](https://github.com/jackc/pgx)

## Status:
	in progress

## Run:
```sh
go build cmd/boilerplate/main.go
```

## Example


```sh
# add new post
curl -X POST "http://localhost:8000/api/post" -H "accept: application/json" -H "Content-Type: application/json" -d '{"content": "hello cruel world"}'

# get all posts
curl -X GET "http://localhost:8000/api/posts" -H "accept: application/json"

# migrate
migrate -source file://$PWD/db/migrations -database "postgres://adm:1234@localhost:5432/alem?sslmode=disable" up
```

![](https://visitor-badge.laobi.icu/badge?page_id=Zulbukharov.golang-ddd-hex)
