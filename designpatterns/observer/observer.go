package main

import "fmt"

type Subject interface {
	register(observer Observer)
	deregister(observer Observer)
	notifyAll()
}

type Item struct {
	ObserverList []Observer
	name         string
	inStock      bool
}

func newItem(name string) *Item {
	return &Item{
		name: name,
	}
}

func (i *Item) updateAvailability() {
	fmt.Printf("Item %s is now available\n", i.name)
	i.inStock = true
	i.notifyAll()
}

func (i *Item) register(o Observer) {
	fmt.Printf("Registering customer: %s\n", o.getID())
	i.ObserverList = append(i.ObserverList, o)
}
func (i *Item) deregister(o Observer) {
	fmt.Printf("Deregistering customer: %s\n", o.getID())
	i.ObserverList = removeFromSlice(i.ObserverList, o)
}

func (i *Item) notifyAll() {
	for _, observer := range i.ObserverList {
		observer.update(i.name)
	}
}

func removeFromSlice(observerList []Observer, obserToRemove Observer) []Observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if obserToRemove.getID() == observer.getID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}

// observer
type Observer interface {
	update(string)
	getID() string
}

// concrete observer

type Customer struct {
	id string
}

func (c *Customer) update(itemName string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.id, itemName)
}

func (c *Customer) getID() string {
	return c.id
}

// cliente

func main() {
	shirtItem := newItem("Nike Shirt")

	observerFirst := &Customer{id: "f3rcho@f3rcho.com"}
	observerSecond := &Customer{id: "adan@adan.com"}

	shirtItem.register(observerFirst)
	shirtItem.register(observerSecond)

	shirtItem.updateAvailability()

	fmt.Println()

	shirtItem.deregister(observerFirst)
	shirtItem.deregister(observerSecond)
	fmt.Println("Process done!")

}
