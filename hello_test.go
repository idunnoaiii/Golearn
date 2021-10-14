package main

import "testing"

func TestHello(t *testing.T) {
    
    assertCorrectMessage := func(t testing.TB, got, want string) {
        //Helper function help report the line number in function which 
        //we want tracking instead of this helper fucntion
        t.Helper()
        if got != want {
            t.Errorf("got %q want %q", got, want)
        }
    }

    t.Run("saying hello to people", func(t *testing.T){
        got := Hello("Chris", "")
        want := "Hello, Chris"
        assertCorrectMessage(t, got, want)
    })

    t.Run("saying hello world when string is empty", func(t *testing.T){
        got := Hello("", "")
        want := "Hello, world"
        assertCorrectMessage(t, got, want)
    })

    t.Run("in Spanish", func(t *testing.T){
        got := Hello("Elodie", "Spanish") 
        want := "Hola, Elodie"
        assertCorrectMessage(t, got, want)
    })

}