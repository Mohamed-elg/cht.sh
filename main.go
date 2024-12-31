package main

import (
	"fmt"

	"github.com/Mohamed-elg/cht.sh/internal/chtsh"
	"github.com/Mohamed-elg/cht.sh/internal/cli"
)

func main() {
	req := cli.GetArgs()
	res := chtsh.GetChtsh(req.Tech, req.AdditionalArgs)
	fmt.Println(res)
}
