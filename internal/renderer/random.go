package renderer

import "math/rand/v2"

func randomIntBetween(min, max int) int {
  return min + rand.IntN(max-min)
}
