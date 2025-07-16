func canVisitAllRooms(rooms [][]int) bool {
    opened := map[int]bool{0: true}
    queued := []int{0}

    for len(queued) > 0 {
        key := queued[0]
        queued = queued[1:]

        for _, key := range rooms[key] {
            if !opened[key] {
                opened[key] = true
                queued = append(queued, key)
            }
        }
    }

    return len(opened) == len(rooms)
}
