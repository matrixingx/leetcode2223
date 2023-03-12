package test8

import "fmt"


type Relate struct {
    Val int
    Mark bool
}

func main() {
    var r1 = &Relate{
        Val: 1,
    }
    var r2 = &Relate{
        Val: 2,
    }
    var r3 = &Relate{
        Val: 3,
    }
    var r4 = &Relate{
        Val: 4,
    }
    var r5 = &Relate{
        Val: 5,
    }
    var r6 = &Relate{
        Val: 6,
    }
    var r7 = &Relate{
        Val: 7,
    }
    var r8 = &Relate{
        Val: 8,
    }
    var related = make(map[*Relate][]*Relate)
    related[r1] = []*Relate{r2,r3}
    related[r4] = []*Relate{r5,r6}
    related[r7] = []*Relate{r1,r8}

    var res bool
    var dfs func(r *Relate)
    dfs = func (r *Relate) {
		if r.Mark {
            res = true
            return
        }
        r.Mark = true
        if res {
            return
        }
        if relatedArr,ok := related[r];ok {
            for i := range relatedArr {
                dfs(relatedArr[i])
            }
        }
    }
    dfs(r4)
    dfs(r5)
    fmt.Println(res)
}