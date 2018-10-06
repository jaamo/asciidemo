package main
import (
  "math/cmplx"
  "github.com/mjibson/go-dsp/fft"
  "github.com/faiface/beep"
)

// Extract audio frequencies from stream.
type Equalizer struct {
  Streamer beep.Streamer
  Frequencies *[]float64
}

func (v *Equalizer) Stream(samples [][2]float64) (n int, ok bool) {

  n, ok = v.Streamer.Stream(samples)

  // Use only left channel.
  channelSamples := make([]float64, len(samples));
  for i := range samples[:n] {
    channelSamples[i] = samples[i][0]
  }

  // Fourier transformation.
  X := fft.FFTReal(channelSamples)

  // Extract frequencies.
  for i := 0; i < 16; i++ {

    // Get magnitude. Ignore angle.
    (*v.Frequencies)[i], _ = cmplx.Polar(X[i])
    
  } 

  // Pass without modifying.  
  return n, ok

}

// Err propagates the wrapped Streamer's errors.
func (v *Equalizer) Err() error {
  return v.Streamer.Err()
}


