package main

import (
	constanta "project-golang/constanta"
	"project-golang/operator"
	"project-golang/selection"
	typedata "project-golang/typeData"
	variabel "project-golang/variabelType"
)

func main() {
	variabel.VariabelType()
	variabel.VariabelNotType()
	variabel.MultipleVariabel()

	typedata.Numerik()
	typedata.Boolean()
	typedata.String()

	constanta.Const()

	operator.Aritmatik()
	operator.Perbandinga()
	operator.Logic()

	selection.Selection()
	selection.Temporary()
	selection.Switch()
	selection.SwitchIf()
	selection.Fallthrough()
	selection.Bersarang()
}
