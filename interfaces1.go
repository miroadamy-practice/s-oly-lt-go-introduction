package main

import "fmt"

type TapePlayer struct {
	Batteries string
}

func (t TapePlayer) Play(song string) {
	fmt.Println("Playing", song)
}
func (t TapePlayer) Stop() {
	fmt.Println("Stopped!")
}

// another type with play/stop
type TapeRecorder struct {
	Microphones int
}
func (t TapeRecorder) Play(song string) {
	fmt.Println("Playing", song)
}
func (t TapeRecorder) Record() {
	fmt.Println("Recording")
}
func (t TapeRecorder) Stop() {
	fmt.Println("Stopped!")
}

// function that ONLY accepts tape player
func playList(device TapePlayer, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func playList2(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

type Player interface {
	// Must have a Play method with
	// a single string parameter
	Play(string)
	// Must have a Stop method with
	// no parameters
	Stop()
}
/*

This Works:

	mixtape := []string{"Jessie's Girl", "Whip It", "9 to 5"}
	var player TapePlayer
	playList(player, mixtape)

Cannot use Recorder - even if has same methods

mixtape := []string{"Jessie's Girl", "Whip It", "9 to 5"}
var recorder TapeRecorder
playList(recorder, mixtape)

prog.go:40:10: cannot use recorder (type TapeRecorder) as type TapePlayer in argument to playList

*/


func main() {
	mixtape := []string{"Jessie's Girl", "Whip It", "9 to 5"}
	var player TapePlayer
	playList(player, mixtape)

	// This works
	var recorder TapeRecorder
	playList2(player, mixtape)
	playList2(recorder, mixtape)

	TryOut(TapeRecorder{})
	TryOut2(TapeRecorder{})
}

// Type assertions

func TryOut(player Player) {
	player.Play("Test Track")
	player.Stop()
	// Player interface doesn't include this method!
	// THIS WILL NOT WORK
	// player.Record()
	// even if we call it TryOut(recorder)
}

// This will work
func TryOut2(player Player) {
	player.Play("Test Track")
	player.Stop()
	// Do a type assertion to get the concrete value back...
	recorder := player.(TapeRecorder)
	// Then you can call Record on that.
	recorder.Record()
}