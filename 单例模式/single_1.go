package manager
import (
    "fmt"
)

var m *Manager

func GetInstance() *Manager {
    if m == nil {
        m = &Manager {}
    }
    return m
}

type Manager struct {}

func (p Manager) Manage() {
    fmt.Println("manage...")
}