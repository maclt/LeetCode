func predictPartyVictory(senate string) string {
    count := len(senate)
    rs := make([]int, 0)
    ds := make([]int, 0)

    for i:=0; i<count; i++ {
        if senate[i] == 'R' {
            rs = append(rs, i)
        } else if senate[i] == 'D' {
            ds = append(ds, i)
        }
    }

    for len(rs) > 0 && len(ds) > 0 {
        if rs[0] < ds[0] {
            rs = append(rs, count)
        } else {
            ds = append(ds, count)
        }
        rs = rs[1:]
        ds = ds[1:]
        count = count + 1
    }

    if len(ds) == 0 {
        return "Radiant"
    } else {
        return "Dire"
    }

}
