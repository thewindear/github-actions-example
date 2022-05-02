package utils

import (
    "math/rand"
    "strings"
    "time"
)

func RandomStr(length int) string {
    var chr = "-_1234567890qazxswedcvfrtgbnhyujmkiolpQAZXSWEDCVFRTGBNHYUJMKIOLP"
    chrList := strings.Split(chr, "")
    chrListLen := len(chrList)
    var randomStr = make([]string, length)
    rand.Seed(time.Now().Unix())
    for i := 0; i < length; i++ {
        randInt := rand.Intn(chrListLen)
        randomStr = append(randomStr, chrList[randInt])
    }
    return strings.Join(randomStr, "")
}
