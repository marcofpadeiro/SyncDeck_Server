package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	config, err := ReadConfig()
	if err != nil {
		log.Panic("Error reading config")
	}

	if len(os.Args) < 2 {
		help()
		return
	}

	commands := map[string]func(Config, []string){
		"subscribe":  HandleSubscribe,
		"del":        HandleDel,
		"rm":         HandleRemove,
		"add-remote": HandleAddRemote,
		"list":       HandleList,
		"fetch":      HandleFetch,
		"upload":     HandleUpload,
		"refresh":    HandleRefresh,
	}

	cmd, ok := commands[os.Args[1]]
	if !ok {
		help()
		return
	}

	cmd(config, os.Args[2:])
}

func help() {
	fmt.Println("Usage:")
	fmt.Println("    ./syncdeck <command> [arguments]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("    help                             print this menu")
	fmt.Println()
	fmt.Println("Unit Managment:")
	fmt.Println("    subscribe <unit> <path>          subscribe to unit")
	fmt.Println("    add-remote <unit> <path>         add remote unit")
	fmt.Println("    del <unit>                       remove local unit")
	fmt.Println("    rm  <unit>                       remove local unit and its files (DANGEROUS!)")
	fmt.Println("    list                             list local units")
	fmt.Println("    list-remote                      list remote units")
	fmt.Println("    fetch <unit>                     fetch unit")
	fmt.Println("    upload <unit>                    upload into server")
	fmt.Println("    refresh                          refresh all units versions")
}
