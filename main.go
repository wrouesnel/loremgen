// Test file for the lorem ipsum generator

package main

import (
	"time"
	lorem "github.com/drhodes/golorem"
	cli "gopkg.in/urfave/cli.v1"
	"fmt"
	"os"
	"math/rand"
)

func main() {
	app := cli.NewApp()

	app.Name = "Lorem Ipsum Generator"

	app.Flags = []cli.Flag{
		cli.Float64Flag{Name: "rate", Usage: "Lines of lorem ipsum per second", Value: 1,},
		cli.IntFlag{Name: "min-words", Usage: "Minimum number of words per line", Value: 10,},
		cli.IntFlag{Name: "max-words", Usage: "Maximum number of words per line", Value: 10,},
		cli.Int64Flag{Name: "rand-seed", Usage: "Specify the random seed of the generator.", Value: 0,},

	}

	app.Action = func(c *cli.Context) error {
		rand.Seed(c.Int64("rand-seed"))

		interval := float64(1) / c.Float64("rate")

		ticker := time.Tick(time.Duration(float64(time.Second) * interval))

		for _ = range ticker {
			fmt.Println(lorem.Sentence(c.Int("min-words"), c.Int("max-words")))
		}

		return nil
	}

	app.Run(os.Args)
}