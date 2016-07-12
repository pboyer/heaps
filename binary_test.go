package heap

import (
    "testing"
    "math/rand"
)

func TestBinary(t *testing.T) {
    for r := 0; r < 20; r++ {
        c := 1000
        h := makeBinaryHeap()

        for i := 0; i < c; i++ {
            h.insert(rand.Int())
        }

        last := 0
        for i := 0; i < c/2; i++ {
            v, err := h.dequeue()

            if err != nil {
                t.Fatalf("unexpected error %v", err)
            }

            if i == 0 {
                last = v
                continue
            }

            if v < last {
                t.Fatalf("violation of heap property",)
            }

            last = v
        }
    }
}
