package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"

	"github.com/admin100/util/console"
	"github.com/eiannone/keyboard"
)

var (
	site string
)


func main() {
	clear()
	console.SetConsoleTitle("Ip Pinger")
	fmt.Println(".___         __________.__                             ")
	fmt.Println("|   |_____   \\______   \\__| ____    ____   ___________ ")
	fmt.Println("|   \\____ \\   |     ___/  |/    \\  / ___\\_/ __ \\_  __ \\")
	fmt.Println("|   |  |_> >  |    |   |  |   |  \\/ /_/  >  ___/|  | \\/")
	fmt.Println("|___|   __/   |____|   |__|___|  /\\___  / \\___  >__|   ")
	fmt.Println("    |__|                       \\//_____/      \\/       ")
	fmt.Print("\n\n")

	fmt.Print("Coloque o link do site: ")
	fmt.Scanln(&site)

	ips, _ := net.LookupIP(site)
    for _, ip := range ips {
        if ipv4 := ip.To4(); ipv4 != nil {
            fmt.Println("IPv4: ", ipv4)
        }
    }

	fmt.Print("\n\n")
	fmt.Print("Pressione R para reiniciar o programa ou qualquer tecla para sair: ")
	
	char, _, err := keyboard.GetSingleKey()
	if err != nil {
		fmt.Println(err)
	}

	if char == 114 || char == 82 {
		main()
	} else {
		os.Exit(0)
	}
}

func clear() {
	c := exec.Command("cmd", "/c", "cls")
	c.Stdout = os.Stdout
	c.Run()
}