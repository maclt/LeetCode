func nearestExit(maze [][]byte, entrance []int) int {
    height := len(maze)
    width := len(maze[0])

    stepped := make([][]bool, height)
    for i := range stepped {
        stepped[i] = make([]bool, width)
    }
    stepped[entrance[0]][entrance[1]] = true
    
    queued := make([][]int, 0)
    queued = append(queued, entrance)
    distance := 0

    isInRange := func(p []int) bool {
        return 0 <= p[0] &&
            height > p[0] &&
            0 <= p[1] &&
            width > p[1]
    }

    isEmpty := func(p []int) bool {
        return maze[p[0]][p[1]] == '.'
    }

    isExit := func(p []int) bool{
        return p[0] == 0 ||
            p[0] == height-1 ||
            p[1] == 0 ||
            p[1] == width-1
    }

    for len(queued) > 0 {
        for _, position := range queued {
            neighbors := getNeighbors(position)
            queued = queued[1:]

            for _, n := range neighbors {
                if !isInRange(n) {
                    continue
                }

                if stepped[n[0]][n[1]] {
                    continue
                }

                if !isEmpty(n) {
                    continue
                }

                if isExit(n) {
                    return distance + 1
                }

                stepped[n[0]][n[1]] = true
                queued = append(queued, n)
            }
        }
        distance = distance + 1
    }

    return -1
}


func getNeighbors(p []int) [][]int {
    var neighbors = make([][]int, 0, 4)
    neighbors = append(neighbors, []int{p[0]+1, p[1]})
    neighbors = append(neighbors, []int{p[0]-1, p[1]})
    neighbors = append(neighbors, []int{p[0], p[1]+1})
    neighbors = append(neighbors, []int{p[0], p[1]-1})

    return neighbors
}