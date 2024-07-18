package main

import "fmt"

// Abstract Factory Interface
type GUIFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
}

// Concrete Factory 1
type MacOSFactory struct{}

func (f MacOSFactory) CreateButton() Button {
	return MacOSButton{}
}

func (f MacOSFactory) CreateCheckbox() Checkbox {
	return MacOSCheckbox{}
}

// Concrete Factory 2
type WindowsFactory struct{}

func (f WindowsFactory) CreateButton() Button {
	return WindowsButton{}
}

func (f WindowsFactory) CreateCheckbox() Checkbox {
	return WindowsCheckbox{}
}

// Abstract Product: Button
type Button interface {
	Paint()
}

// Concrete Product: MacOS Button
type MacOSButton struct{}

func (b MacOSButton) Paint() {
	fmt.Println("Painting a MacOS button")
}

// Concrete Product: Windows Button
type WindowsButton struct{}

func (b WindowsButton) Paint() {
	fmt.Println("Painting a Windows button")
}

// Abstract Product: Checkbox
type Checkbox interface {
	Paint()
}

// Concrete Product: MacOS Checkbox
type MacOSCheckbox struct{}

func (c MacOSCheckbox) Paint() {
	fmt.Println("Painting a MacOS checkbox")
}

// Concrete Product: Windows Checkbox
type WindowsCheckbox struct{}

func (c WindowsCheckbox) Paint() {
	fmt.Println("Painting a Windows checkbox")
}

func main() {
	macFactory := MacOSFactory{}
	application(macFactory)

	winFactory := WindowsFactory{}
	application(winFactory)
}

func application(factory GUIFactory) {
	button := factory.CreateButton()
	checkbox := factory.CreateCheckbox()

	button.Paint()
	checkbox.Paint()
}
