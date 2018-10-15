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
	tasks := []string{"cook", "clean", "laundry", "eat", "sleep", "code"}
	apiUrls := []string{"/time", "time/set/{10:30/18:00}", }

  	app := cli.NewApp()
  	app.EnableBashCompletion = true
  	app.Commands = []cli.Command{
    {
      Name:  "complete",
      Aliases: []string{"c"},
      Usage: "complete a task on the list",
      Action: func(c *cli.Context) error {
		 fmt.Println("completed task: ", c.Args().First())
         return nil
      },
      BashComplete: func(c *cli.Context) {
        // This will complete if no args are passed
        if c.NArg() > 0 {
          return
        }
        for _, t := range tasks {
          fmt.Println(t)
        }
      },
	},
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