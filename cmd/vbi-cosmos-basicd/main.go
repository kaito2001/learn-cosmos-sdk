package main

import (
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"vbi-cosmos-basic/app"
	"vbi-cosmos-basic/cmd/vbi-cosmos-basicd/cmd"
	prototest "vbi-cosmos-basic/proto-test"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}

	prototest.Execution()

}
