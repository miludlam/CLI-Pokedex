package main

import (
	"fmt"
	"os"
)

func commandExit(c *config) error {
	_ = c // config is unneeded for this function, so discard
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
