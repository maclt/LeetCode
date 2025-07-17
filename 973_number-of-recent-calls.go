type RecentCounter struct {
    timestamp []int
}

func Constructor() RecentCounter {
    return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
    this.timestamp = append(this.timestamp, t)
    var filtered []int
    for i:=0; i<len(this.timestamp); i++ {
        if this.timestamp[i] >= t-3000 && this.timestamp[i] <= t {
            filtered = append(filtered, this.timestamp[i])
        }
    }
    this.timestamp = filtered
    return len(this.timestamp)
}



/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */