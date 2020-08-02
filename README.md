# Golang DDD+HEX demo blog api

## Dependencies:
- [pgx](https://github.com/jackc/pgx)

## Status:
	in progress

## Run:
```sh
go build cmd/boilerplate/main.go
```

## Goals: good structure goal

- consistent.
- easy to test.
- easy to understand.
- easy to change.

## Project Specs

- users can add a post.
- users can add a comment for the post.
- users can vote for the post.
- users can list all posts with votes.
- users can list all reviews for the post.

## Example

```sh
# add new post
curl -X POST "http://localhost:8000/api/post" -H "accept: application/json" -H "Content-Type: application/json" -d '{"content": "hello cruel world"}'

# get all posts
curl -X GET "http://localhost:8000/api/posts" -H "accept: application/json"
```

![](https://visitor-badge.laobi.icu/badge?page_id=Zulbukharov.golang-ddd-hex)
