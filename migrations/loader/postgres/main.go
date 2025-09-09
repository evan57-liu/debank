package main

import (
	"fmt"
	"io"
	"os"

	"ariga.io/atlas-provider-gorm/gormschema"

	"github.com/coin50etf/coin-market/internal/model"
)

func main() {
	stmts, err := gormschema.New("postgres").Load(
		&model.ProtocolPosition{},
		&model.ProtocolMapping{},
		&model.UserToken{},
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}
	if _, err = io.WriteString(os.Stdout, stmts); err != nil {
		fmt.Fprintf(os.Stderr, "failed to write gorm schema: %v\n", err)
		os.Exit(1)
	}
}
