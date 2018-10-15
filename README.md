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

## use mysql through api

`curl -i http://localhost:8000/text` get all mysql text slices

`curl -d -X POST http://localhost:8000/text/schlubbel` add schlubbel as new text item to mysql

# commandline application

for bash autocompletition

bash

`cd gh && go install && echo "PROG=gh source $GOPATH/src/github.com/urfave/cli/autocomplete/bash_autocomplete" >> ~/.bashrc && source ~/.bashrc`

zsh

`cd gh && go install && echo "PROG=gh source $GOPATH/src/github.com/urfave/cli/autocomplete/zsh_autocomplete" >> ~/.zshrc && source ~/.zshrc`

## usage

`gh get http://localhost:8000/time`
