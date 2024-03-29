package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"

	ttt "github.com/lelledaniele/gotictactoe"
)

func main() {
	a := os.Args

	if len(a) != 2 {
		fmt.Println("You must specify an argument for the battle's field square side")
		os.Exit(1)
	}

	n64, e := strconv.ParseInt(a[1], 10, 0)

	if e != nil {
		fmt.Println("The argument is not an integer")
		os.Exit(1)
	}

	n := int(n64)
	g := ttt.NewGame(n)
	p := g.GetPlayers()

	e = clearTerminal()

	if e != nil {
		fmt.Printf("\nTerminal clear not supported\n")
	}

	for {
		for i := range p {
			_ = clearTerminal() // Not important for the game

			printBattleField(g.GetBattleField())

			if wp, w := g.GetWinner(); w {
				fmt.Printf("'%v' is the winner\n", string(wp.GetSymbol()))
				os.Exit(1)
			}
			ec := g.GetCellsWithEmptyValue()

			if len(ec) == 0 {
				fmt.Printf("Draft!\n")
				os.Exit(1)
			}

			fmt.Printf("Player '%v', please select a choice\n", string(p[i].GetSymbol()))

			for j, c := range ec {
				fmt.Printf("%d) %v\n", j, c)
			}

			reader := bufio.NewReader(os.Stdin)
			b, e := reader.ReadByte()

			if e != nil {
				fmt.Println(e)
				os.Exit(1)
			}

			j, e := strconv.Atoi(string(b))

			if e != nil || j >= len(ec) {
				fmt.Printf("'%v' is an invalid choice\n", string(b))
				os.Exit(1)
			}

			e = g.AddTurn(ec[j], p[i].GetSymbol())

			if e != nil {
				fmt.Printf("'Impossible add game '%v' turn for '%v' player\n", ec[j], string(p[i].GetSymbol()))
				os.Exit(1)
			}
		}
	}
}

func printBattleField(bt [][]byte) {
	fmt.Printf("\n")

	for _, r := range bt {
		for _, rc := range r {
			if rc == 0 {
				fmt.Printf("-")
			} else {
				fmt.Printf("%v", string(rc))
			}

			fmt.Printf(" ")
		}

		fmt.Printf("\n")
	}

	fmt.Printf("\n")
}

func clearTerminal() error {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
