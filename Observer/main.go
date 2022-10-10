package main

type Observable struct {
	observers []Observer
}

func (o *Observable) Subscribe(observer Observer) {
	o.observers = append(o.observers, observer)
}

func (o *Observable) Unsubscribe(observer Observer) {
	for i, obs := range o.observers {
		if obs == observer {
			o.observers = append(o.observers[:i], o.observers[i+1:]...)
		}
	}
}

type Observer interface {
	Notify(PropertyChange)
}

func (o *Observable) Fire(p PropertyChange) {
	for _, observer := range o.observers {
		observer.Notify(p)
	}
}

type Person struct {
	Observable
	Name string
	age  int
}

type PropertyChange struct {
	property string
	value    int
}

func (p *Person) SetAge(age int) {
	if age == p.age {
		return
	}
	p.age = age
	p.Fire(PropertyChange{"age", age})
}

func (p *Person) Age() int {
	return p.age
}

type PersonObserver struct {
	Name string
}

func (p PersonObserver) Notify(pc PropertyChange) {
	println(p.Name, "says:", "Person's", pc.property, "has changed to", pc.value)
}

func main() {
	p := Person{}
	p.Subscribe(PersonObserver{"John"})
	p.Subscribe(PersonObserver{"Jane"})
	p.SetAge(20)
	p.SetAge(30)
	p.SetAge(30)

}
