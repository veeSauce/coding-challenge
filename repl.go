package main

import (
	"bufio"
	."fmt"
	"github.com/golang-collections/collections/stack"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

type data struct {
	Stack    *stack.Stack
	Command  []string
	KeyStore map[string]string
}

func main() {

	log.SetOutput(os.Stderr)

	s := stack.New()
	m := make(map[string]string)

	d := &data{
		s,
		nil,
		m,
	}

	for {
		d.DoStuff()
	}
}

func (d *data) DoStuff() {

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
			d.Write()

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


		// push the stack with updated key value pair transactions
		d.Stack.Push(d.KeyStore)
		Println("These are the values in keystore")
		//d.KeyStore = d.Stack.Peek().(map[string]string)
		Println("The length of the stack is: ",d.Stack.Len())


	case "ABORT":
		Println("length of stack is: ", d.Stack.Len())

		// wipe out the current transaction
		// if stack is 0 clear the map, if its 1 pop the stack then clear the map
		//if d.Stack.Len() == 1 {
		//	d.Stack.Pop()
		//	m := make(map[string]string)
		//	d.KeyStore = m
		//
		//} else
		if d.Stack.Len() == 0 {
			m := make(map[string]string)
			d.KeyStore = m
		} else {
			d.Stack.Pop()

			d.KeyStore = d.Stack.Peek().(map[string]string)
			Print(d.KeyStore)
			Println("after aborting length of stack is: ", d.Stack.Len())
		}

	case "COMMIT":

		// pop stack and push a commit in popped index
		d.Stack.Pop()
		d.Stack.Pop()
		d.Stack.Push(d.KeyStore)
		//d.KeyStore = d.Stack.Peek().(map[string]string)



	case "QUIT":

		log.Println("Exiting...")
		os.Exit(0)

	default:

		// write error to stderr
		log.Println("Invalid command")

	}

}

func (d *data) Write() {

	// first write writes to stack as root store
	if d.Stack.Len() == 0 && len(d.KeyStore) == 0{
		d.KeyStore[d.Command[1]] = d.Command[2]
		d.Stack.Push(d.KeyStore)


	}
		// just a regular write
		d.KeyStore[d.Command[1]] = d.Command[2]
	}


func (d *data) Read(){


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

	log.Println("delete called")

	if _, ok := d.KeyStore[d.Command[1]]; ok {
		Println("deleting: ", d.Command[1])
		delete(d.KeyStore, d.Command[1])
	} else {
		log.Println("Key not found: ", d.Command[1])
	}
}
