package heap

import ( "fmt" )

type binaryHeap struct {
    vals []int
}

func makeBinaryHeap() *binaryHeap {
    return &binaryHeap{ 
        vals : make([]int, 0, 10),
    }
}

func left(i int) int {
    return 2 * i + 1
}

func right(i int) int {
    return 2 * i + 2
}

func par(i int) int {
    return (i - 1) / 2
}

func (h *binaryHeap) insert(val int) {
    h.vals = append(h.vals, val)

    i := len(h.vals) - 1

    for {
        if i == 0 {
            return
        }

        parI := par(i)
        if h.vals[parI] > h.vals[i] {
            h.vals[parI], h.vals[i] = h.vals[i], h.vals[parI]
            i = parI
        } else {
            return
        }
    }
}

func (h *binaryHeap) dequeue() (int, error) {
    if len(h.vals) == 0 {
        return -1, fmt.Errorf("Empty heap")
    }

    val := h.vals[0]
    h.heapify()

    return val, nil
}

func (h *binaryHeap) heapify(){
    h.vals[0] = h.vals[len(h.vals) - 1]
    h.vals = h.vals[0:len(h.vals) - 1]

    if len(h.vals) == 0 {
        return
    }

    i := 0

    for {
        leftI := left(i)
        rightI := right(i)
        smallest := i

        // promote the lesser of the two children
        if leftI < len(h.vals) && h.vals[smallest] > h.vals[leftI] { 
            smallest = leftI
        }

        if rightI < len(h.vals) && h.vals[smallest] > h.vals[rightI] {
            smallest = rightI
        }

        if smallest != i {
            h.vals[smallest], h.vals[i] = h.vals[i], h.vals[smallest]
            i = smallest
        } else {
            // this subtree satisfies the heap property, done
            return
        }
    }
}


