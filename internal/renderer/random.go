package renderer

import "math/rand/v2"

const seed = 100000

func randomIntBetween(min, max int) int {
  mn := min * seed
  mx := max * seed
  return int((rand.IntN(mx - mn) + mn )/ seed)
}
