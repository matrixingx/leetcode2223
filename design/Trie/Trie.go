package trie


type Trie struct {
	IsEnd bool
	Val byte
	List [26]*Trie
}


func Constructor() Trie {
	return Trie{}
}


func (this *Trie) Insert(word string)  {
	t := this
	for i := range word {
		isEnd := i == len(word)-1
		if !isEnd {
			if t.List[word[i]-'a'] != nil {
				t = t.List[word[i]-'a']
			} else {
				t.List[word[i]-'a'] = &Trie{
					Val: word[i],
				}
				t = t.List[word[i]-'a']
			}
		} else {
			if t.List[word[i]-'a'] != nil {
				t.List[word[i]-'a'].IsEnd = true
			} else {
				t.List[word[i]-'a'] = &Trie{
					Val: word[i],
					IsEnd: true,
				}
			}
		}
	}
}


func (this *Trie) Search(word string) bool {
	t := this
	for i := range word {
		isEnd := i == len(word)-1
		if t.List[word[i]-'a'] == nil {
			return false
		}
		if isEnd && !t.List[word[i]-'a'].IsEnd {
			return false
		}
		t = t.List[word[i]-'a']
	}
	return true
}


func (this *Trie) StartsWith(prefix string) bool {
	t := this
	for i := range prefix {
		if t.List[prefix[i]-'a'] == nil {
			return false
		}
		
		t = t.List[prefix[i]-'a']
	}
	return true
}