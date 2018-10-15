// command line interface
// go install and then root /urfave/cli/autocomplete/.. in .zshrc or .bashrc
package main

import (
  "fmt"
  "log"
  "os"
  "github.com/urfave/cli"
)

func main() {
    apiUrls := []string{"/time", "/time/set/hh:mm/hh:mm", "/text", "/text/x" }

  	app := cli.NewApp()
  	app.EnableBashCompletion = true
  	app.Commands = []cli.Command{
	{
      Name:  "post",
      Aliases: []string{"c"},
      Usage: "send a http post request",
      Action: func(c *cli.Context) error {
		/*err := */sendPost(c.Args().Get(0)) 
		return nil
      },
      BashComplete: func(c *cli.Context) {
        if c.NArg() > 0 {
          return
        }
        for _, t := range apiUrls {
          fmt.Println(t)
        }
      },
	},
	{
      Name:  "get",
      Aliases: []string{"c"},
      Usage: "get and print a http get request",
      Action: func(c *cli.Context) error {
		r := getGet(c.Args().Get(0)) 
		fmt.Println(r)
		return nil
      },
      BashComplete: func(c *cli.Context) {
        if c.NArg() > 0 {
          return
        }
        for _, t := range apiUrls {
          fmt.Println(t)
        }
      },
    },
  }

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}