package state

import "ondrejmaksi.com/db2project/ui/lib"

type KeyHintsState struct {
	*lib.GenericState[string]
}

func (k *KeyHintsState) SetForUser() {
	k.SetState("Edit: e  |  Delete: d  |  Basket: Enter  |  Create order: o")
}

func (k *KeyHintsState) SetForProduct() {
	k.SetState("Edit: e  |  Delete: d ")
}

func (k *KeyHintsState) SetForBasket() {
	k.SetState("Edit: e  |  Delete: d | Add: a ")
}

func (k *KeyHintsState) SetForAddToBasket() {
	k.SetState("Add: a ")
}

func (k *KeyHintsState) SetForOrder() {
	k.SetState("Pay: p  |  Expedite: e | Cancel: c ")
}

func (k *KeyHintsState) Clear() {
	k.SetState("")
}

func NewKeyHintsState(value string) *KeyHintsState {
	return &KeyHintsState{
		lib.NewGenericState(value),
	}
}
