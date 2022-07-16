package main

import (
	_ "swagger/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"swagger/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
