package tests

import (
    "github-actions-example/utils"
    "os"
    "testing"
)

func TestUtil_Random(t *testing.T) {
    randStr := utils.RandomStr(10)
    if len(randStr) != 10 {
        t.Error("random string number not 10")
    }
}

func TestUtil_LookEnv(t *testing.T) {
    if os.Getenv("GO_USERNAME") != "root" {
        t.Error("GO_USERNAME NOT ROOT")
    }
    if os.Getenv("GO_PASSWORD") != "root" {
        t.Error("GO_PASSWORD NOT ROOT")
    }
    t.Log(os.Getenv("GO_EXPRESSION"))
}
