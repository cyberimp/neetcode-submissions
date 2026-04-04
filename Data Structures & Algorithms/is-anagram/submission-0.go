func isAnagram(s string, t string) bool {
    mapS := map[rune]int{}
    mapT := map[rune]int{}
    var r rune
    for r = 'a'; r <= 'z'; r++{
        mapS[r]++
        mapT[r]++
    }
    for _,r = range s {
        mapS[r]++
    }
    for _,r = range t {
        mapT[r]++
    }
    for r,v := range mapS {
        if mapT[r] != v {return false}
    }
    return true

}
