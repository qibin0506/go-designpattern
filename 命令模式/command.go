package main 
import "fmt"

func main() {
    var tv TV
    openCommand := OpenCommand{tv}
    invoker := Invoker{openCommand}
    invoker.Do()

    invoker.SetCommand(CloseCommand{tv})
    invoker.Do()

    var openClose *OpenCloseCommand = NewOpenCloseCommand()
    openClose.AddCommand(OpenCommand{tv})
    openClose.AddCommand(CloseCommand{tv})
    invoker.SetCommand(openClose)
    invoker.Do()
}

func NewOpenCloseCommand() *OpenCloseCommand {
    var openClose *OpenCloseCommand = &OpenCloseCommand{}
    openClose.cmds = make([]Command, 2)
    return openClose
}

type OpenCloseCommand struct {
    index int
    cmds  []Command
}

func (p *OpenCloseCommand) AddCommand(cmd Command) {
    p.cmds[p.index] = cmd
    p.index++
}

func (p OpenCloseCommand) Press() {
    for _, item := range p.cmds {
        item.Press()
    }
}

// receiver
type TV struct{}

func (p TV) Open() {
    fmt.Println("play...")
}

func (p TV) Close() {
    fmt.Println("stop...")
}

// command
type Command interface {
    Press()
}

type OpenCommand struct {
    tv TV
}

func (p OpenCommand) Press() {
    p.tv.Open()
}

type CloseCommand struct {
    tv TV
}

func (p CloseCommand) Press() {
    p.tv.Close()
}

// sender
type Invoker struct {
    cmd Command
}

func (p *Invoker) SetCommand(cmd Command) {
    p.cmd = cmd
}

func (p Invoker) Do() {
    p.cmd.Press()
}