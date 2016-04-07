package manager
import (
    "sync"
    "fmt"
)

var m *Manager
var once sync.Once

func GetInstance() *Manager {
    once.Do(func() {
        m = &Manager {}
    })
    return m
}

type Manager struct {}

func (p Manager) Manage() {
    fmt.Println("manage...")
}