# api 

testing on cli

## time
`curl -i http://localhost:8000/time`

`curl -i http://localhost:8000/time/percent`

`curl -d -X POST http://localhost:8000/time/set/12:00/23:50`

## fs

`curl -i http://localhost:8000/read`

`curl -d -X POST http://localhost:8000/write/test`
