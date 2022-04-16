package state

import "testing"

func TestDesign(t *testing.T) {
    t.Run("test for state", func(t *testing.T) {
        assert := func(got, want string) {
            if got != want {
                t.Error("state test fail")
            }
        }

        account := NewAccount(1)

        got := account.SendMsg()
        want := "send msg success"
        assert(got, want)

        got = account.SendMsg()
        want = "balance not enough, send msg fail"
        assert(got, want)
    })
}
