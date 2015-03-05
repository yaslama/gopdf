package fonts //change this
import (
	"github.com/yaslama/gopdf"
)
type RobotoCondensedLightItalic struct {
	family string
	fonttype string
	name string
	desc  []gopdf.FontDescItem
	up int
	ut int
	cw gopdf.FontCw
	enc string
	diff string
}
func (me * RobotoCondensedLightItalic) Init(){
	me.cw = make(gopdf.FontCw)
	me.cw[gopdf.Chr(0)]=200
	me.cw[gopdf.Chr(1)]=200
	me.cw[gopdf.Chr(2)]=200
	me.cw[gopdf.Chr(3)]=200
	me.cw[gopdf.Chr(4)]=200
	me.cw[gopdf.Chr(5)]=200
	me.cw[gopdf.Chr(6)]=200
	me.cw[gopdf.Chr(7)]=200
	me.cw[gopdf.Chr(8)]=200
	me.cw[gopdf.Chr(9)]=200
	me.cw[gopdf.Chr(10)]=200
	me.cw[gopdf.Chr(11)]=200
	me.cw[gopdf.Chr(12)]=200
	me.cw[gopdf.Chr(13)]=200
	me.cw[gopdf.Chr(14)]=200
	me.cw[gopdf.Chr(15)]=200
	me.cw[gopdf.Chr(16)]=200
	me.cw[gopdf.Chr(17)]=200
	me.cw[gopdf.Chr(18)]=200
	me.cw[gopdf.Chr(19)]=200
	me.cw[gopdf.Chr(20)]=200
	me.cw[gopdf.Chr(21)]=200
	me.cw[gopdf.Chr(22)]=200
	me.cw[gopdf.Chr(23)]=200
	me.cw[gopdf.Chr(24)]=200
	me.cw[gopdf.Chr(25)]=200
	me.cw[gopdf.Chr(26)]=200
	me.cw[gopdf.Chr(27)]=200
	me.cw[gopdf.Chr(28)]=200
	me.cw[gopdf.Chr(29)]=200
	me.cw[gopdf.Chr(30)]=200
	me.cw[gopdf.Chr(31)]=200
	me.cw[gopdf.ToByte(" ")]=200
	me.cw[gopdf.ToByte("!")]=198
	me.cw[gopdf.ToByte("\"")]=255
	me.cw[gopdf.ToByte("#")]=479
	me.cw[gopdf.ToByte("$")]=433
	me.cw[gopdf.ToByte("%")]=593
	me.cw[gopdf.ToByte("&")]=501
	me.cw[gopdf.ToByte("'")]=152
	me.cw[gopdf.ToByte("(")]=257
	me.cw[gopdf.ToByte(")")]=260
	me.cw[gopdf.ToByte("*")]=352
	me.cw[gopdf.ToByte("+")]=460
	me.cw[gopdf.ToByte(",")]=163
	me.cw[gopdf.ToByte("-")]=227
	me.cw[gopdf.ToByte(".")]=205
	me.cw[gopdf.ToByte("/")]=319
	me.cw[gopdf.ToByte("0")]=433
	me.cw[gopdf.ToByte("1")]=433
	me.cw[gopdf.ToByte("2")]=433
	me.cw[gopdf.ToByte("3")]=433
	me.cw[gopdf.ToByte("4")]=433
	me.cw[gopdf.ToByte("5")]=433
	me.cw[gopdf.ToByte("6")]=433
	me.cw[gopdf.ToByte("7")]=433
	me.cw[gopdf.ToByte("8")]=433
	me.cw[gopdf.ToByte("9")]=433
	me.cw[gopdf.ToByte(":")]=189
	me.cw[gopdf.ToByte(";")]=192
	me.cw[gopdf.ToByte("<")]=396
	me.cw[gopdf.ToByte("=")]=435
	me.cw[gopdf.ToByte(">")]=405
	me.cw[gopdf.ToByte("?")]=362
	me.cw[gopdf.ToByte("@")]=726
	me.cw[gopdf.ToByte("A")]=498
	me.cw[gopdf.ToByte("B")]=485
	me.cw[gopdf.ToByte("C")]=491
	me.cw[gopdf.ToByte("D")]=506
	me.cw[gopdf.ToByte("E")]=448
	me.cw[gopdf.ToByte("F")]=448
	me.cw[gopdf.ToByte("G")]=520
	me.cw[gopdf.ToByte("H")]=545
	me.cw[gopdf.ToByte("I")]=223
	me.cw[gopdf.ToByte("J")]=427
	me.cw[gopdf.ToByte("K")]=457
	me.cw[gopdf.ToByte("L")]=419
	me.cw[gopdf.ToByte("M")]=642
	me.cw[gopdf.ToByte("N")]=546
	me.cw[gopdf.ToByte("O")]=521
	me.cw[gopdf.ToByte("P")]=485
	me.cw[gopdf.ToByte("Q")]=530
	me.cw[gopdf.ToByte("R")]=508
	me.cw[gopdf.ToByte("S")]=467
	me.cw[gopdf.ToByte("T")]=444
	me.cw[gopdf.ToByte("U")]=500
	me.cw[gopdf.ToByte("V")]=489
	me.cw[gopdf.ToByte("W")]=651
	me.cw[gopdf.ToByte("X")]=479
	me.cw[gopdf.ToByte("Y")]=457
	me.cw[gopdf.ToByte("Z")]=428
	me.cw[gopdf.ToByte("[")]=215
	me.cw[gopdf.ToByte("\\")]=317
	me.cw[gopdf.ToByte("]")]=215
	me.cw[gopdf.ToByte("^")]=328
	me.cw[gopdf.ToByte("_")]=346
	me.cw[gopdf.ToByte("`")]=250
	me.cw[gopdf.ToByte("a")]=421
	me.cw[gopdf.ToByte("b")]=436
	me.cw[gopdf.ToByte("c")]=407
	me.cw[gopdf.ToByte("d")]=436
	me.cw[gopdf.ToByte("e")]=406
	me.cw[gopdf.ToByte("f")]=268
	me.cw[gopdf.ToByte("g")]=436
	me.cw[gopdf.ToByte("h")]=436
	me.cw[gopdf.ToByte("i")]=195
	me.cw[gopdf.ToByte("j")]=201
	me.cw[gopdf.ToByte("k")]=390
	me.cw[gopdf.ToByte("l")]=195
	me.cw[gopdf.ToByte("m")]=672
	me.cw[gopdf.ToByte("n")]=436
	me.cw[gopdf.ToByte("o")]=436
	me.cw[gopdf.ToByte("p")]=436
	me.cw[gopdf.ToByte("q")]=436
	me.cw[gopdf.ToByte("r")]=274
	me.cw[gopdf.ToByte("s")]=390
	me.cw[gopdf.ToByte("t")]=254
	me.cw[gopdf.ToByte("u")]=436
	me.cw[gopdf.ToByte("v")]=386
	me.cw[gopdf.ToByte("w")]=580
	me.cw[gopdf.ToByte("x")]=386
	me.cw[gopdf.ToByte("y")]=386
	me.cw[gopdf.ToByte("z")]=373
	me.cw[gopdf.ToByte("{")]=267
	me.cw[gopdf.ToByte("|")]=192
	me.cw[gopdf.ToByte("}")]=267
	me.cw[gopdf.ToByte("~")]=524
	me.cw[gopdf.Chr(127)]=200
	me.cw[gopdf.Chr(128)]=410
	me.cw[gopdf.Chr(129)]=200
	me.cw[gopdf.Chr(130)]=157
	me.cw[gopdf.Chr(131)]=266
	me.cw[gopdf.Chr(132)]=255
	me.cw[gopdf.Chr(133)]=505
	me.cw[gopdf.Chr(134)]=427
	me.cw[gopdf.Chr(135)]=439
	me.cw[gopdf.Chr(136)]=354
	me.cw[gopdf.Chr(137)]=735
	me.cw[gopdf.Chr(138)]=468
	me.cw[gopdf.Chr(139)]=240
	me.cw[gopdf.Chr(140)]=720
	me.cw[gopdf.Chr(141)]=200
	me.cw[gopdf.Chr(142)]=445
	me.cw[gopdf.Chr(143)]=200
	me.cw[gopdf.Chr(144)]=200
	me.cw[gopdf.Chr(145)]=159
	me.cw[gopdf.Chr(146)]=159
	me.cw[gopdf.Chr(147)]=263
	me.cw[gopdf.Chr(148)]=265
	me.cw[gopdf.Chr(149)]=276
	me.cw[gopdf.Chr(150)]=531
	me.cw[gopdf.Chr(151)]=621
	me.cw[gopdf.Chr(152)]=356
	me.cw[gopdf.Chr(153)]=479
	me.cw[gopdf.Chr(154)]=390
	me.cw[gopdf.Chr(155)]=240
	me.cw[gopdf.Chr(156)]=697
	me.cw[gopdf.Chr(157)]=200
	me.cw[gopdf.Chr(158)]=373
	me.cw[gopdf.Chr(159)]=457
	me.cw[gopdf.Chr(160)]=200
	me.cw[gopdf.Chr(161)]=192
	me.cw[gopdf.Chr(162)]=424
	me.cw[gopdf.Chr(163)]=446
	me.cw[gopdf.Chr(164)]=580
	me.cw[gopdf.Chr(165)]=465
	me.cw[gopdf.Chr(166)]=188
	me.cw[gopdf.Chr(167)]=471
	me.cw[gopdf.Chr(168)]=394
	me.cw[gopdf.Chr(169)]=638
	me.cw[gopdf.Chr(170)]=348
	me.cw[gopdf.Chr(171)]=362
	me.cw[gopdf.Chr(172)]=426
	me.cw[gopdf.Chr(173)]=227
	me.cw[gopdf.Chr(174)]=639
	me.cw[gopdf.Chr(175)]=363
	me.cw[gopdf.Chr(176)]=312
	me.cw[gopdf.Chr(177)]=434
	me.cw[gopdf.Chr(178)]=327
	me.cw[gopdf.Chr(179)]=332
	me.cw[gopdf.Chr(180)]=250
	me.cw[gopdf.Chr(181)]=436
	me.cw[gopdf.Chr(182)]=376
	me.cw[gopdf.Chr(183)]=207
	me.cw[gopdf.Chr(184)]=200
	me.cw[gopdf.Chr(185)]=211
	me.cw[gopdf.Chr(186)]=354
	me.cw[gopdf.Chr(187)]=361
	me.cw[gopdf.Chr(188)]=590
	me.cw[gopdf.Chr(189)]=610
	me.cw[gopdf.Chr(190)]=651
	me.cw[gopdf.Chr(191)]=375
	me.cw[gopdf.Chr(192)]=498
	me.cw[gopdf.Chr(193)]=498
	me.cw[gopdf.Chr(194)]=498
	me.cw[gopdf.Chr(195)]=498
	me.cw[gopdf.Chr(196)]=498
	me.cw[gopdf.Chr(197)]=498
	me.cw[gopdf.Chr(198)]=704
	me.cw[gopdf.Chr(199)]=480
	me.cw[gopdf.Chr(200)]=448
	me.cw[gopdf.Chr(201)]=448
	me.cw[gopdf.Chr(202)]=448
	me.cw[gopdf.Chr(203)]=448
	me.cw[gopdf.Chr(204)]=223
	me.cw[gopdf.Chr(205)]=223
	me.cw[gopdf.Chr(206)]=223
	me.cw[gopdf.Chr(207)]=223
	me.cw[gopdf.Chr(208)]=531
	me.cw[gopdf.Chr(209)]=546
	me.cw[gopdf.Chr(210)]=527
	me.cw[gopdf.Chr(211)]=527
	me.cw[gopdf.Chr(212)]=527
	me.cw[gopdf.Chr(213)]=527
	me.cw[gopdf.Chr(214)]=527
	me.cw[gopdf.Chr(215)]=431
	me.cw[gopdf.Chr(216)]=521
	me.cw[gopdf.Chr(217)]=500
	me.cw[gopdf.Chr(218)]=500
	me.cw[gopdf.Chr(219)]=500
	me.cw[gopdf.Chr(220)]=500
	me.cw[gopdf.Chr(221)]=457
	me.cw[gopdf.Chr(222)]=457
	me.cw[gopdf.Chr(223)]=457
	me.cw[gopdf.Chr(224)]=421
	me.cw[gopdf.Chr(225)]=421
	me.cw[gopdf.Chr(226)]=421
	me.cw[gopdf.Chr(227)]=421
	me.cw[gopdf.Chr(228)]=421
	me.cw[gopdf.Chr(229)]=421
	me.cw[gopdf.Chr(230)]=646
	me.cw[gopdf.Chr(231)]=407
	me.cw[gopdf.Chr(232)]=406
	me.cw[gopdf.Chr(233)]=406
	me.cw[gopdf.Chr(234)]=406
	me.cw[gopdf.Chr(235)]=406
	me.cw[gopdf.Chr(236)]=193
	me.cw[gopdf.Chr(237)]=193
	me.cw[gopdf.Chr(238)]=193
	me.cw[gopdf.Chr(239)]=193
	me.cw[gopdf.Chr(240)]=452
	me.cw[gopdf.Chr(241)]=436
	me.cw[gopdf.Chr(242)]=436
	me.cw[gopdf.Chr(243)]=436
	me.cw[gopdf.Chr(244)]=436
	me.cw[gopdf.Chr(245)]=436
	me.cw[gopdf.Chr(246)]=436
	me.cw[gopdf.Chr(247)]=441
	me.cw[gopdf.Chr(248)]=436
	me.cw[gopdf.Chr(249)]=436
	me.cw[gopdf.Chr(250)]=436
	me.cw[gopdf.Chr(251)]=436
	me.cw[gopdf.Chr(252)]=436
	me.cw[gopdf.Chr(253)]=386
	me.cw[gopdf.Chr(254)]=441
	me.cw[gopdf.Chr(255)]=386
	me.up = -73
	me.ut = 49
	me.fonttype = "TrueType"
	me.name = "RobotoCondensed-LightItalic"
	me.enc = "cp1252"
	me.desc = make([]gopdf.FontDescItem,8)
	me.desc[0] =  gopdf.FontDescItem{ Key:"Ascent",Val : "750" }
	me.desc[1] =  gopdf.FontDescItem{ Key: "Descent", Val : "-250" }
	me.desc[2] =  gopdf.FontDescItem{ Key:"CapHeight", Val :  "711" }
	me.desc[3] =  gopdf.FontDescItem{ Key: "Flags", Val :  "32" }
	me.desc[4] =  gopdf.FontDescItem{ Key:"FontBBox", Val :  "[-406 -271 957 1048]" }
	me.desc[5] =  gopdf.FontDescItem{ Key:"ItalicAngle", Val :  "0" }
	me.desc[6] =  gopdf.FontDescItem{ Key:"StemV", Val :  "70" }
 	me.desc[7] =  gopdf.FontDescItem{ Key:"MissingWidth", Val :  "200" } 
 }
func (me * RobotoCondensedLightItalic)GetType() string{
	return me.fonttype
}
func (me * RobotoCondensedLightItalic)GetName() string{
	return me.name
}	
func (me * RobotoCondensedLightItalic)GetDesc() []gopdf.FontDescItem{
	return me.desc
}
func (me * RobotoCondensedLightItalic)GetUp() int{
	return me.up
}
func (me * RobotoCondensedLightItalic)GetUt()  int{
	return me.ut
}
func (me * RobotoCondensedLightItalic)GetCw() gopdf.FontCw{
	return me.cw
}
func (me * RobotoCondensedLightItalic)GetEnc() string{
	return me.enc
}
func (me * RobotoCondensedLightItalic)GetDiff() string {
	return me.diff
}
func (me * RobotoCondensedLightItalic) GetOriginalsize() int{
	return 98764
}
func (me * RobotoCondensedLightItalic)  SetFamily(family string){
	me.family = family
}
func (me * RobotoCondensedLightItalic) 	GetFamily() string{
	return me.family
}
