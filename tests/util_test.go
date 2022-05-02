package tests

import (
    "github-actions-example/utils"
    "testing"
)

func TestUtil_Random(t *testing.T) {
    randStr := utils.RandomStr(10)
    if len(randStr) != 10 {
        t.Error("random string number not 10")
    }
}
