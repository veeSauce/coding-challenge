package main

// Simple key value store with nested transactions allowed

import (
	"PA/stack"
	"bufio"
	. "fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

type data struct {
	// Stack stack.Stack
	Command  []string
	KeyStore map[string]string
}

func main() {

	log.SetOutput(os.Stderr)
	Stack := stack.New()

	m := make(map[string]string)

	d := data{
		// Stack
		nil,
		m,
	}

	for {

		reader := bufio.NewReader(os.Stdin)

		Print(">")

		// read in arguments
		text, _ := reader.ReadString('\n')

		// collect arguments in slice
		text = strings.Replace(text, "\n", "", -1)
		command := strings.Fields(text)
		command[0] = strings.ToUpper(command[0])
		d.Command = command

		switch d.Command[0] {

		case "WRITE":

			if ok := d.Check(3); ok {
				d.Write(Stack)

			} else {
				log.Println("write must have a key and value")
			}

		case "READ":

			if ok := d.Check(2); ok {
				d.Read()
			} else {
				log.Println("Read command requires a key")
			}

		case "DELETE":
			if ok := d.Check(2); ok {
				d.Delete()
			} else {
				log.Println("Delete command requires a key")
			}

		case "START":

			// push the stack with updated key value map as a new Map not pointer to old one
			newM := createNewMap()
			for k, v := range d.KeyStore {
				newM[k] = v
			}
			Stack.Push(newM)

		case "ABORT":

			// wipe out the current transaction
			if Stack.Len() == 0 {
				m := make(map[string]string)
				d.KeyStore = m
			} else {
				Stack.Pop()
				peekMap := Stack.Peek()
				d.KeyStore = peekMap
			}

		case "COMMIT":

			// pop stack and push a commit in popped index
			Stack.Pop()
			Stack.Pop()
			newM := createNewMap()
			for k, v := range d.KeyStore {
				newM[k] = v
			}
			Stack.Push(newM)

		case "QUIT":

			log.Println("Exiting...")
			os.Exit(0)

		default:

			// write error to stderr
			log.Println("Invalid command")
		}
	}
}

func (d *data) Write(Stack *stack.Stack) *stack.Stack {

	// first write writes to stack as root store
	if Stack.Len() == 0 && len(d.KeyStore) == 0 {
		d.KeyStore[d.Command[1]] = d.Command[2]

		newMap := make(map[string]string)
		for k, v := range d.KeyStore {
			newMap[k] = v
		}
		Stack.Push(newMap)
		return Stack

	}
	// just a regular write
	d.KeyStore[d.Command[1]] = d.Command[2]
	return Stack
}

func (d *data) Read() {

	if val, ok := d.KeyStore[d.Command[1]]; ok {
		Println(val)
	} else {
		log.Println("Key not found: ", d.Command[1])
	}
}

func (d *data) Check(reqLen int) bool {

	if len(d.Command) == reqLen {
		return true
	} else {
		return false
	}
}

func (d *data) Delete() {

	if _, ok := d.KeyStore[d.Command[1]]; ok {
		delete(d.KeyStore, d.Command[1])
	} else {
		log.Println("Key not found: ", d.Command[1])
	}
}

func createNewMap() map[string]string {
	new := make(map[string]string)
	return new
}
