package command

import (
	"fmt"
)

// 示例场景：使用遥控器控制卧室和厨房灯的开和关，支持同时开同时关，同时支持返回撤销操作

// 灯
type Light struct {
	name string
}

func (l Light) On() {
	fmt.Println(l.name + " is on")
}

func (l Light) Off() {
	fmt.Println(l.name + " is off")
}

// 命令
type Command interface {
	execute()
	undo()
}

type NoCommand struct{}

func (l NoCommand) execute() {}

func (l NoCommand) undo() {}

type LightOnCommand struct {
	light Light
}

func (l LightOnCommand) execute() {
	l.light.On()
}

func (l LightOnCommand) undo() {
	l.light.Off()
}

type LightOffCommand struct {
	light Light
}

func (l LightOffCommand) execute() {
	l.light.Off()
}

func (l LightOffCommand) undo() {
	l.light.On()
}

type MacroCommand struct {
	commands []Command
}

func (l MacroCommand) execute() {
	for i := 0; i < len(l.commands); i++ {
		l.commands[i].execute()
	}
}

func (l MacroCommand) undo() {
	for i := 0; i < len(l.commands); i++ {
		l.commands[i].execute()
	}
}

var _ Command = (*NoCommand)(nil)
var _ Command = (*LightOnCommand)(nil)
var _ Command = (*LightOffCommand)(nil)
var _ Command = (*MacroCommand)(nil)

// 遥控器
type RemoteControl struct {
	onCommands, offCommands []Command
	undoCommand             Command
}

func NewRemoteControl() *RemoteControl {
	slots := 3
	onCommands := make([]Command, slots)
	offCommands := make([]Command, slots)

	noCommand := &NoCommand{}
	for i := 0; i < slots; i++ {
		onCommands[i] = noCommand
		offCommands[i] = noCommand
	}
	return &RemoteControl{onCommands: onCommands, offCommands: offCommands, undoCommand: noCommand}
}

func (rc *RemoteControl) setCommand(slot int, onCommand, offCommand Command) {
	rc.onCommands[slot] = onCommand
	rc.offCommands[slot] = offCommand
}

func (rc *RemoteControl) onButtonWasPushed(slot int) {
	rc.onCommands[slot].execute()
	rc.undoCommand = rc.onCommands[slot]
}

func (rc *RemoteControl) offButtonWasPushed(slot int) {
	rc.offCommands[slot].execute()
	rc.undoCommand = rc.offCommands[slot]
}

func (rc *RemoteControl) undoButtonWasPushed() {
	rc.undoCommand.undo()
}
