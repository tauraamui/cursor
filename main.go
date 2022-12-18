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

type Frame struct {
	content []byte
}

func main() {

}
