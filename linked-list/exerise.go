package main

import "fmt"

type Address struct {
	name            string
	phone_number    string
	email_address   string
}

type Node struct {
	addr *Address
	next *Node
}

type AddressList struct {
	root *Node
	len  int
}

type AddressListI interface {
	Insert(cur *Node, name, phone_number, email_address string) *Node
	GetAddress(n int) *Node
	Remove(n int, cur *Node) (*Node, error)
	Traverse()
}

var _AddressListI = &AddressList{}

func (a *AddressList) Insert(cur *Node, name, phone_number, email_address string) *Node {
	address := &Address{
		name:name,
		phone_number: phone_number,
		email_address: email_address,
	}

	var node *Node = &Node{addr: address}
	if a.len != 0 {
		node.next = cur.next
		cur.next = node
		cur = cur.next
	} else {
		a.root = node
		cur = a.root
	}

	a.len++
	return cur
}

func (a *AddressList) GetAddress(n int) *Node {
	if n > a.len {
		return nil
	}
	ptr := a.root
	for i := 1; i < n; i++ {
		ptr = ptr.next
	}
	return ptr
}

func (a *AddressList) Remove(n int, cur *Node) (*Node, error) {
	if a.len == 0 {
		return cur, fmt.Errorf("List is empty")
	}

	target := a.GetAddress(n)
	if target == nil {
		return cur, fmt.Errorf("Address not found")
	}

	if n > 1 {
		prev := a.GetAddress(n - 1)
		prev.next = target.next
		if n == a.len {
			cur = prev
		}
	} else {
		a.root = target.next
	}

	a.len--
	return cur, nil
}

func (a *AddressList) Traverse() {
	if a.len == 0 {
		fmt.Println("List is empty")
	}

	fmt.Printf("Size of AddressList: %d\n", a.len)
	ptr := a.root
	for i := 1; i <= a.len; i++ {
		fmt.Printf("%d name: %s, phone_number: %s, email_address: %s\n", i, ptr.addr.name, ptr.addr.phone_number, ptr.addr.email_address)
		ptr = ptr.next
	}
}

func main() {
	var addressList *AddressList = &AddressList{}
	fmt.Println("\n~~Init AddressList~~")
	addressList.Traverse()

	var cur *Node = &Node{}
	cur = addressList.Insert(cur, "Sakura", "090-1234-5678", "sakura@example.com")
	cur = addressList.Insert(cur, "Kaede", "080-1234-5678", "kaede@example.com")
	cur = addressList.Insert(cur, "Yuzu", "070-1234-5678", "yuzu@example.com")
	fmt.Println("\n~~After Insert Addresses~~")
	addressList.Traverse()

	cur, err := addressList.Remove(2, cur)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("\n~~After Remove Address~~")
	addressList.Traverse()
}