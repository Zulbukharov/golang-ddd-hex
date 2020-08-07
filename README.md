# Golang DDD+HEX demo blog api

- tactical design
    - entities
        - an object defined primarily by its identity
    - value objects
        - immutable
    - repositories
        - how will communicate with storage
    - services
        - corresponds application use cases
        - example:
        ` oauth -> user policy -> adding post`

context: blogpost <br>
ubiquitous language: author, post, vote, comments, storage <br>
entities: author, post, votes, comments <br>
value objects: it can be part of entity object (author, voter, commenter) <br>
aggregates: PostAuthor, PostVotes, PostComments <br>
- service:
    - stateless operations 
        - post adder / listing,
        - vote increment / decrement / count
        - comment adder / listing
        - author register / login 
- events: can affect to the system (errors, logs)
- repository: facade over backend

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

http://localhost:8000/swagger/index.html

```sh
# add new post
curl -X POST "http://localhost:8000/api/post" -H "accept: application/json" -H "Content-Type: application/json" -d '{"content": "hello cruel world"}'

# get all posts
curl -X GET "http://localhost:8000/api/posts" -H "accept: application/json"

# migrate
migrate -source file://$PWD/db/migrations -database "postgres://adm:1234@localhost:5432/alem?sslmode=disable" up

swag init -g cmd/server/main.go --parseDependency --parseInternal -o ./api
```

## Links
https://khalilstemmler.com/wiki/conways-law/ <br>
https://medium.com/@hatajoe/clean-architecture-in-go-4030f11ec1b1 <br>
https://khalilstemmler.com/articles/software-design-architecture/organizing-app-logic/ <br>
http://olivergierke.de/2020/03/Implementing-DDD-Building-Blocks-in-Java/ <br>
https://archfirst.org/domain-driven-design-8-conclusion/ <br>
https://herbertograca.com/2017/09/14/ports-adapters-architecture/ <br>
https://blog.fedecarg.com/2009/03/11/domain-driven-design-and-mvc-architectures/#:~:text=According%20to%20Eric%20Evans%2C%20Domain,a%20technology%20or%20a%20methodology.&text=Domain%2Ddriven%20design%20separates%20the,to%20retrieve%20and%20store%20data. <br>
http://www.newtonmeters.com/blog/domain-driven-design-in-go/ <br>
https://vaadin.com/learn/tutorials/ddd/ddd_and_hexagonal <br>
https://www.vinaysahni.com/best-practices-for-a-pragmatic-restful-api <br>

![](https://visitor-badge.laobi.icu/badge?page_id=Zulbukharov.golang-ddd-hex)
