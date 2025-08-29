package main

import (
	"git.sunturtle.xyz/zephyr/errors-my-beloved/distant/brain"
	"git.sunturtle.xyz/zephyr/errors-my-beloved/distant/brain/sqlbrain"
)

func main() {
	brain.Think(new(sqlbrain.Brain), "")
}
