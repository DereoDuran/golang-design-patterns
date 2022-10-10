package main

type Person struct {
	Name    string
	Room    *ChatRoom
	ChatLog []string
}

func (p *Person) Receive(sender, message string) {
	s := sender + ": '" + message + "'"
	println("[", p.Name, "'s chat session] ", s)
	p.ChatLog = append(p.ChatLog, s)
}

func (p *Person) Say(message string) {
	p.Room.Broadcast(p.Name, message)
}

func (p *Person) PrivateMessage(who, message string) {
	p.Room.Message(p.Name, who, message)
}

type ChatRoom struct {
	People []*Person
}

func (c *ChatRoom) Broadcast(source, message string) {
	for _, p := range c.People {
		if p.Name != source {
			p.Receive(source, message)
		}
	}
}

func (c *ChatRoom) Join(p *Person) {
	joinMsg := p.Name + " joins the chat"
	c.Broadcast("room", joinMsg)

	p.Room = c
	c.People = append(c.People, p)
}

func (c *ChatRoom) Message(source, destination, message string) {
	for _, p := range c.People {
		if p.Name == destination {
			p.Receive(source, message)
		}
	}
}

func main() {
	room := &ChatRoom{}

	john := &Person{Name: "John"}
	jane := &Person{Name: "Jane"}

	room.Join(john)
	room.Join(jane)

	john.Say("hi room")
	jane.Say("oh, hey john")

	simon := &Person{Name: "Simon"}
	room.Join(simon)
	simon.Say("hi everyone!")

	jane.PrivateMessage("Simon", "glad you could join us!")

}
