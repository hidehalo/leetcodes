/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */
 
func merge(intervals []Interval) []Interval {
    length := len(intervals)
    seq := Bystart(intervals)
    sort.Sort(seq)
    for i:=0;i<length-1; {
        c,n := intervals[i],intervals[i+1]
        if n.Start <= c.End {
            if n.End > c.End {
                intervals[i].End = n.End
            } 
            s1,s2 := intervals[:i+1],intervals[i+2:]
            intervals = append(s1,s2...)
            length--
            continue
        }
        i++
    }
    
    return intervals
}

type Bystart []Interval

func (intervals Bystart) Len() int{
    return len(intervals)
}

func (intervals Bystart) Less(i,j int) bool {
    return intervals[i].Start < intervals[j].Start
}

func (intervals Bystart) Swap(i, j int) {
    intervals[i],intervals[j] = intervals[j],intervals[i]
}