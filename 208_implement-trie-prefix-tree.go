type Trie struct {
    isEnd bool
    children map[rune]*Trie
}


func Constructor() Trie {
    return Trie{isEnd: false, children: make(map[rune]*Trie)}
}


func (this *Trie) Insert(word string)  {
    node := this;

    for _, c := range word {
        if node.children[c] == nil {
          tmp := Constructor()
          node.children[c] = &tmp
        }

        node = node.children[c]
    }

    node.isEnd = true
}


func (this *Trie) Search(word string) bool {
    node := this;

    for _, c := range word {
        if node.children[c] == nil {
          return false
        }
        node = node.children[c]
    }

    return node.isEnd;
}


func (this *Trie) StartsWith(prefix string) bool {
    node := this;

    for _, c := range prefix {
        if node.children[c] == nil {
          return false
        }
        node = node.children[c]
    }

    return true;
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
