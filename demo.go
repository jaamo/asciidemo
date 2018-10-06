// Prototype file :)

package main

import "github.com/nsf/termbox-go"
import "math/rand"
import "time"

/*
func draw() {
  w, h := termbox.Size()
  termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
  for y := 0; y < h; y++ {
    for x := 0; x < w; x++ {
      termbox.SetCell(x, y, ' ', termbox.ColorDefault,
      termbox.Attribute(rand.Int()%8)+1)
    }
  }
  termbox.Flush()
}
*/

//
// Render random characters to screen.
//
func draw_characters() {

  // Get screen size.
  w, h := termbox.Size()

  // Clear.
  termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

  // Available characters.
  available_characters := []rune{'.', ',', '_', '¨'}

  // Render characters.
  for y := 0; y < h; y++ {
    for x := 0; x < w; x++ {
      termbox.SetCell(x, y, available_characters[rand.Intn(len(available_characters))], termbox.ColorWhite, termbox.ColorBlack)
      //termbox.SetCell(x, y, []rune(available_characters[rand.Intn(len(available_characters) - 1)]), termbox.ColorWhite, termbox.ColorBlack)
    }
  }

  // Flush to screen.
  termbox.Flush()

}




// 
// Draw random characters.
// 
func main2() {

  // Init termbox.
  err := termbox.Init()
  if err != nil {
    panic(err)
  }

  // Close terminal when the function exists.
  defer termbox.Close()

  event_queue := make(chan termbox.Event)
  go func() {
    for {
      event_queue <- termbox.PollEvent()
    }
  }()

  loop:
  for {
    select {
    case ev := <-event_queue:
      if ev.Type == termbox.EventKey && ev.Key == termbox.KeyEsc {
        break loop
      }
    default:
      draw_characters()
      time.Sleep(10 * time.Millisecond)
    }
  }
}
