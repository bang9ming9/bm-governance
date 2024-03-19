package main

import (
	"flag"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	compiledPath string = "/home/go/src/github.com/bang9ming9/bm-governance/bin/contracts"
	outPath      string = "/home/go/src/github.com/bang9ming9/bm-governance/abis"
)

var ContractsName = flag.String("contracts", "", "")

func main() {
	flag.Parse()
	for _, name := range strings.Split(*ContractsName, ",") {
		name = strings.TrimSpace(name)
		if name == "" {
			continue
		}

		cPath := filepath.Join(compiledPath, name)
		oPath := filepath.Join(outPath, name)
		err := exec.Command("abigen", "--pkg=abis", fmt.Sprintf("--abi=%s.abi", cPath), fmt.Sprintf("--bin=%s.bin", cPath), fmt.Sprintf("--out=%s.go", oPath), fmt.Sprintf("--type=%s", name)).Run()
		if err != nil {
			fmt.Println(name, err)
		}
	}
}
