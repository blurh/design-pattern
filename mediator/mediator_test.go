package mediator

import (
    "testing"
)

func TestDesign(t *testing.T) {
    t.Run("test for mediator", func(t *testing.T) {
        accTextField := NewTextField()
        passwdTextField := NewTextField()
        repeatTextField := NewTextField()
        accTextField.Input("hello")
        passwdTextField.Input("world")
        repeatTextField.Input("repeat")

        s := NewSelection()
        s.Checked("login")

        dialog := NewDialog()
        dialog.Account = accTextField
        dialog.Passwd = passwdTextField
        dialog.Type = s

        loginGot := dialog.Submit()
        loginWant := "http://xxx?acc=hello&passwd=world"
        if loginGot != loginWant {
            t.Error("mediator test fail")
        }

        s.Checked("register")
        dialog.RepeatPassword = repeatTextField
        registerGot := dialog.Submit()
        registerWant := "http://xxx?acc=hello&passwd=world&repeat=repeat"
        if registerGot != registerWant {
            t.Error("mediator test fail")
        }
    })
}
