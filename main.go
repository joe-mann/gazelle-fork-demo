package main

import (
	"fmt"

	"github.com/pelletier/go-toml"
)

func main() {
	config, _ := toml.Load(`
[postgres]
user = "pelletier"
password = "mypassword"`)
	user := config.Get("postgres.user").(string)
	fmt.Printf("Hello, %s!\n", user)
}
