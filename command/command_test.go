package command

import (
	"fmt"
	"testing"
)

func TestCommand(t *testing.T) {
	remoteControl := NewRemoteControl()
	livingRoomLight := Light{name: "Living Room"}
	kitchenLight := Light{name: "Kitchen"}

	livingRoomLightOn := LightOnCommand{light: livingRoomLight}
	livingRoomLightOff := LightOffCommand{light: livingRoomLight}
	kitchenLightOn := LightOnCommand{light: kitchenLight}
	kitchenLightOff := LightOffCommand{light: kitchenLight}

	onMacroCommand := []Command{
		livingRoomLightOn,
		kitchenLightOn,
	}
	offMacroCommand := []Command{
		livingRoomLightOff,
		kitchenLightOff,
	}
	remoteControl.setCommand(0, livingRoomLightOn, livingRoomLightOff)
	remoteControl.setCommand(1, kitchenLightOn, kitchenLightOff)
	remoteControl.setCommand(2, MacroCommand{commands: onMacroCommand}, MacroCommand{commands: offMacroCommand})

	fmt.Println(`--- Pushing LivingRoomLight On---`)
	remoteControl.onButtonWasPushed(0)
	fmt.Println(`--- Pushing LivingRoomLight Off---`)
	remoteControl.offButtonWasPushed(0)
	fmt.Println()
	fmt.Println(`--- Pushing kitchenLight On---`)
	remoteControl.onButtonWasPushed(1)
	fmt.Println(`--- Pushing kitchenLight Off---`)
	remoteControl.offButtonWasPushed(1)
	fmt.Println()
	fmt.Println(`--- Pushing Macro On---`)
	remoteControl.onButtonWasPushed(2)
	fmt.Println(`--- Pushing Macro Off---`)
	remoteControl.offButtonWasPushed(2)
}
