package main

type Storage struct {
	blocks []Block
}

type Block struct {
	EventsBlock
	FramesBlock
}

type EventsBlock struct{}

type FramesBlock struct{}

func main() {

}
