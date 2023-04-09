package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const inputFileName = "input.txt"

type lineType int

const (
	Empty lineType = iota
	Crate
	Index
	Move
)

type crate struct {
	Name     byte
	Previous *crate
	Next     *crate
}

func (c *crate) insert(b byte) *crate {
	if c == nil {
		return &crate{Name: b}
	}
	if c.Name == 0 {
		c.Name = b
		c.Previous = nil
		c.Next = nil
		return c
	} else {
		new := &crate{Name: b, Previous: c.Previous, Next: c}
		c.Previous = new
		return new
	}
}

func (c *crate) tail() *crate {
	return c.tailN(1)
	/* Previous working script.
	tail := c
	for {
		if tail.Next == nil {
			break
		} else {
			tail = tail.Next
		}
	}
	return tail */
}

// tail(1) returns the last item, tail(2) returns from 2nd from last, etc.
func (c *crate) tailN(offset int) *crate {
	tail := c
	for {
		if tail.Next == nil {
			break
		} else {
			tail = tail.Next
		}
	}
	for m := 1; m < offset; m++ {
		if tail.Previous != nil {
			tail = tail.Previous
		} else {
			return nil
		}
	}
	return tail
}

// returns head, err
func (c *crate) appendTail(new *crate) (*crate, error) {
	if new == nil {
		return c, fmt.Errorf("unable to append nil")
	}
	if c == nil {
		//new.Next = nil
		new.Previous = nil
		return new, nil
	}
	tail := c.tail()
	new.Previous = tail
	tail.Next = new
	return c, nil
}

func (c *crate) removeTail() (*crate, *crate, error) {
	return c.removeTailN(1)
	/* if c == nil {
		return nil, nil, fmt.Errorf("cannot remove tail from nil")
	}
	tail := c.tail()
	//fmt.Printf("%c\n", tail.Name)
	if tail == c {
		return nil, tail, nil
	} else {
		//fmt.Println(tail.Previous.Name)
		tail.Previous.Next = nil
		tail.Previous = nil
		return c, tail, nil
	} */
}

// returns head, tail, err. Head is the new head in case the head is the one removed (nil).
func (c *crate) removeTailN(offset int) (*crate, *crate, error) {
	if c == nil {
		return nil, nil, fmt.Errorf("cannot remove tail from nil")
	}
	tail := c.tailN(offset)
	if tail == nil {
		return nil, nil, fmt.Errorf("tailN returned nil")
	}
	if tail == c {
		return nil, tail, nil
	} else {
		tail.Previous.Next = nil
		tail.Previous = nil
		return c, tail, nil
	}
}

func (stack *crate) print() string {
	print := ""
	for c := stack; c != nil; c = c.Next {
		print = fmt.Sprintf("%s%c", print, c.Name)
	}
	return print
}

type crates []*crate

func (stacks crates) print() {
	for s := 0; s < len(stacks); s++ {
		crate := stacks[s].print()
		fmt.Printf("%d: %s\n", s+1, crate)
	}

}

func (stacks crates) copyCrates() crates {
	new := make(crates, len(stacks))
	for i := 0; i < len(stacks); i++ {
		name := stacks[i].Name
		new[i] = &crate{Name: name}

		// Loop into the next crates
		source := stacks[i]
		dest := new[i]
		for {
			if source.Next == nil {
				break
			} else {
				n := &crate{Name: source.Next.Name, Previous: dest}
				dest.Next = n
				dest = n
				source = source.Next
			}
		}
	}
	return new
}

func readInputFile() []string {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Println("Error opening file!!!")
		panic(err)
	}
	defer file.Close()

	data := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
		panic(err)
	}
	return data
}

func getLineType(str string) lineType {
	if len(str) == 0 {
		return Empty
	} else if strings.Contains(str, "move") {
		return Move
	} else if strings.Contains(str, "[") {
		return Crate
	} else if strings.ContainsAny(str, "0123456789") {
		return Index
	} else {
		return Empty
	}
}

func main() {
	var numStacks int
	cratesLines := make([]string, 0)
	moveLines := make([]string, 0)

	input := readInputFile()
	for _, in := range input {
		lt := getLineType(in)
		if lt == Crate {
			ns := len(in)/4 + 1
			if ns > numStacks {
				numStacks = ns
			}
			cratesLines = append(cratesLines, in)
		} else if lt == Move {
			moveLines = append(moveLines, in)
		} else {
			continue
		}
	}

	// Populate the stacks
	stacks := make(crates, numStacks)
	for _, c := range cratesLines {
		fmt.Printf("%s", c)
		var j int
		for i := 0; i < len(c); i += 4 {
			name := c[i+1]
			fmt.Printf("%c", name)
			if name != ' ' {
				stacks[j] = stacks[j].insert(name)
			}
			j++
		}
		fmt.Println()
	}

	// Just making sure our stacks are correct
	stacks.print()
	stacks2 := stacks.copyCrates()
	stacks2.print()

	// Process the moves
	moves := make([][]int, len(moveLines))
	for j, m := range moveLines {
		moveStr := make([]string, 3)
		move1 := strings.Split(m[5:], " from ")
		moveStr[0] = move1[0]
		move2 := strings.Split(move1[1], " to ")
		moveStr[1] = move2[0]
		moveStr[2] = move2[1]
		move := make([]int, 3)
		for i := 0; i < 3; i++ {
			conv, err := strconv.Atoi(moveStr[i])
			if err != nil {
				log.Fatalf("error converting %s to int", moveStr[i])
			}
			move[i] = conv
		}
		moves[j] = move
		//fmt.Println(move)

		// Part 1 moves each item one by one
		for i := 0; i < move[0]; i++ {
			from := stacks[move[1]-1]
			to := stacks[move[2]-1]
			head, tail, _ := from.removeTail()
			stacks[move[1]-1] = head
			head, _ = to.appendTail(tail)
			stacks[move[2]-1] = head
		}

	}
	fmt.Println("Final Stacks:")
	stacks.print()

	// Par 2 moves each items as a stack
	fmt.Println("Now processing stacks2 for part 2")
	stacks2.print()
	for _, move := range moves {
		fmt.Println(move)
		from2 := stacks2[move[1]-1]
		to2 := stacks2[move[2]-1]
		head2, tail2, err := from2.removeTailN(move[0])
		if err != nil {
			log.Println(err)
			fmt.Println(move)
			stacks2.print()
			os.Exit(1)
		}
		stacks2[move[1]-1] = head2
		head2, err = to2.appendTail(tail2)
		if err != nil {
			log.Println(err)
			fmt.Println(move)
			stacks2.print()
			os.Exit(1)
		}
		stacks2[move[2]-1] = head2
		stacks2.print()
	}
	fmt.Println("Final stacks2:")
	stacks2.print()
}
