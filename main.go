package main

import "os"
import "time"
import "github.com/faiface/beep"
import "github.com/faiface/beep/wav"
import "github.com/faiface/beep/speaker"
import "github.com/nsf/termbox-go"

func main() {

  // Load and decode audio.
  f, _ := os.Open("441414__greek555__sample-128-bpm.wav")
  s, format, _ := wav.Decode(f)

  // Frequencies.
  frequencies := make([]float64, 16)

  // Add equalizer filter.
  s2 := &Equalizer {
    Streamer: s,
    Frequencies: &frequencies,
  }

   // Init termbox.
  err := termbox.Init()
  if err != nil {
    panic(err)
  }

  // Close terminal when the function exists.
  defer termbox.Close()

  // Init speaker.
  speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
  
  // Make done function.
  done := make(chan struct{})

  // Start render loop.
  go render_loop(&frequencies, done)

  // Play sound and call done when finished.
  speaker.Play(beep.Seq(s2, beep.Callback(func() {
    close(done)
  })))

  // Wait until callback.
  <-done

}

func render_loop(frequencies *[]float64, done chan struct{}) {

  // Listen events from terminal.
  event_queue := make(chan termbox.Event)
  go func() {
    for {
      event_queue <- termbox.PollEvent()
    }
  }()

  // Loopz.
  loop:
  for {
    select {
    case ev := <-event_queue:
      if ev.Type == termbox.EventKey && ev.Key == termbox.KeyEsc {
        close(done)
        break loop
      }
    default:
      draw_bars(frequencies);
      time.Sleep(10 * time.Millisecond)
    }
  }

}

//
// Render random characters to screen.
//
func draw_bars(frequencies *[]float64) {

  // Get screen size.
  // w, h := termbox.Size()

  // Clear.
  termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

  for i := 0; i < len(*frequencies); i++ {
    bar_length := int((*frequencies)[i])
    for x := 0; x < bar_length; x++ {
      termbox.SetCell(x, i, '#', termbox.ColorWhite, termbox.ColorBlack)
    }
    //fmt.Println((*frequencies)[i])
  }

  // Flush to screen.
  termbox.Flush()

}


