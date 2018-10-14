# api 

testing on cli

## start
`go run *.go` to start api server

`cd client && npm i && npm run dev` to install and start vue dev server

## time api
`curl -i http://localhost:8000/time`

`curl -i http://localhost:8000/time/percent`

`curl -d -X POST http://localhost:8000/time/set/12:00/23:50`

## fs api

`curl -i http://localhost:8000/read`

`curl -d -X POST http://localhost:8000/write/test`
