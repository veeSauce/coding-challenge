package stack

// This is not my code, It is from an open source repo and I modified it to support my application
// Original repo is >> github.com/golang-collections/collections

type (
	Stack struct {
		top *node
		length int
	}
	node struct {
		value map[string]string
		prev *node
	}
)
// Create a new stack
func New() *Stack {
	return &Stack{nil,0}
}
// Return the number of items in the stack
func (this *Stack) Len() int {
	return this.length
}
// View the top item on the stack
func (this *Stack) Peek() map[string]string {
	if this.length == 0 {
		return nil
	}
	return this.top.value
}
// Pop the top item of the stack and return it
func (this *Stack) Pop() map[string]string {
	if this.length == 0 {
		return nil
	}

	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}
// Push a value onto the top of the stack
func (this *Stack) Push(value map[string]string) {
	n := &node{value,this.top}
	this.top = n
	this.length++
}
