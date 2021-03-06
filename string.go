package webscraper

import (
    "strings"
    "sort"
)

// Returns true if two []string contains equal number of strings, and each
// string from 1st slice equals to string from 2nd slice in ascending order
func compareStr(s1, s2 []string) bool {
    if len(s1) != len(s2) {
        return false
    }

    sort.Strings(s1)
    sort.Strings(s2)

    for i, v := range s1 {
        if strings.ToLower(v) != strings.ToLower(s2[i]) {
            return false
        }
    }

    return true
}

// Returns the number of times the same word occurs in the slice
// e.g {"red", "red", "red"} = 2
func hasRepetition(val []string) int {
    count := 0
    for i:=0; i < len(val); i++ {
        for j:=i+1; j < len(val); j++ {
            if val[i] == val[j] {
                count += 1
            }
        }
    }

    if count > 2 {
        if count == len(val) - (len(val) - count) { return count-1 }
    }

    return count
}
