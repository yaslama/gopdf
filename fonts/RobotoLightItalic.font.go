package fonts //change this
import (
	"github.com/yaslama/gopdf"
)
type RobotoLightItalic struct {
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
func (me * RobotoLightItalic) Init(){
	me.cw = make(gopdf.FontCw)
	me.cw[gopdf.Chr(0)]=231
	me.cw[gopdf.Chr(1)]=231
	me.cw[gopdf.Chr(2)]=231
	me.cw[gopdf.Chr(3)]=231
	me.cw[gopdf.Chr(4)]=231
	me.cw[gopdf.Chr(5)]=231
	me.cw[gopdf.Chr(6)]=231
	me.cw[gopdf.Chr(7)]=231
	me.cw[gopdf.Chr(8)]=231
	me.cw[gopdf.Chr(9)]=231
	me.cw[gopdf.Chr(10)]=231
	me.cw[gopdf.Chr(11)]=231
	me.cw[gopdf.Chr(12)]=231
	me.cw[gopdf.Chr(13)]=231
	me.cw[gopdf.Chr(14)]=231
	me.cw[gopdf.Chr(15)]=231
	me.cw[gopdf.Chr(16)]=231
	me.cw[gopdf.Chr(17)]=231
	me.cw[gopdf.Chr(18)]=231
	me.cw[gopdf.Chr(19)]=231
	me.cw[gopdf.Chr(20)]=231
	me.cw[gopdf.Chr(21)]=231
	me.cw[gopdf.Chr(22)]=231
	me.cw[gopdf.Chr(23)]=231
	me.cw[gopdf.Chr(24)]=231
	me.cw[gopdf.Chr(25)]=231
	me.cw[gopdf.Chr(26)]=231
	me.cw[gopdf.Chr(27)]=231
	me.cw[gopdf.Chr(28)]=231
	me.cw[gopdf.Chr(29)]=231
	me.cw[gopdf.Chr(30)]=231
	me.cw[gopdf.Chr(31)]=231
	me.cw[gopdf.ToByte(" ")]=231
	me.cw[gopdf.ToByte("!")]=217
	me.cw[gopdf.ToByte("\"")]=274
	me.cw[gopdf.ToByte("#")]=582
	me.cw[gopdf.ToByte("$")]=521
	me.cw[gopdf.ToByte("%")]=692
	me.cw[gopdf.ToByte("&")]=578
	me.cw[gopdf.ToByte("'")]=163
	me.cw[gopdf.ToByte("(")]=297
	me.cw[gopdf.ToByte(")")]=302
	me.cw[gopdf.ToByte("*")]=400
	me.cw[gopdf.ToByte("+")]=530
	me.cw[gopdf.ToByte(",")]=183
	me.cw[gopdf.ToByte("-")]=271
	me.cw[gopdf.ToByte(".")]=229
	me.cw[gopdf.ToByte("/")]=375
	me.cw[gopdf.ToByte("0")]=521
	me.cw[gopdf.ToByte("1")]=521
	me.cw[gopdf.ToByte("2")]=521
	me.cw[gopdf.ToByte("3")]=521
	me.cw[gopdf.ToByte("4")]=521
	me.cw[gopdf.ToByte("5")]=521
	me.cw[gopdf.ToByte("6")]=521
	me.cw[gopdf.ToByte("7")]=521
	me.cw[gopdf.ToByte("8")]=521
	me.cw[gopdf.ToByte("9")]=521
	me.cw[gopdf.ToByte(":")]=205
	me.cw[gopdf.ToByte(";")]=207
	me.cw[gopdf.ToByte("<")]=480
	me.cw[gopdf.ToByte("=")]=526
	me.cw[gopdf.ToByte(">")]=487
	me.cw[gopdf.ToByte("?")]=428
	me.cw[gopdf.ToByte("@")]=853
	me.cw[gopdf.ToByte("A")]=592
	me.cw[gopdf.ToByte("B")]=583
	me.cw[gopdf.ToByte("C")]=601
	me.cw[gopdf.ToByte("D")]=615
	me.cw[gopdf.ToByte("E")]=541
	me.cw[gopdf.ToByte("F")]=542
	me.cw[gopdf.ToByte("G")]=644
	me.cw[gopdf.ToByte("H")]=663
	me.cw[gopdf.ToByte("I")]=257
	me.cw[gopdf.ToByte("J")]=517
	me.cw[gopdf.ToByte("K")]=550
	me.cw[gopdf.ToByte("L")]=500
	me.cw[gopdf.ToByte("M")]=812
	me.cw[gopdf.ToByte("N")]=666
	me.cw[gopdf.ToByte("O")]=631
	me.cw[gopdf.ToByte("P")]=582
	me.cw[gopdf.ToByte("Q")]=639
	me.cw[gopdf.ToByte("R")]=617
	me.cw[gopdf.ToByte("S")]=564
	me.cw[gopdf.ToByte("T")]=560
	me.cw[gopdf.ToByte("U")]=629
	me.cw[gopdf.ToByte("V")]=584
	me.cw[gopdf.ToByte("W")]=835
	me.cw[gopdf.ToByte("X")]=576
	me.cw[gopdf.ToByte("Y")]=570
	me.cw[gopdf.ToByte("Z")]=515
	me.cw[gopdf.ToByte("[")]=229
	me.cw[gopdf.ToByte("\\")]=373
	me.cw[gopdf.ToByte("]")]=229
	me.cw[gopdf.ToByte("^")]=392
	me.cw[gopdf.ToByte("_")]=408
	me.cw[gopdf.ToByte("`")]=272
	me.cw[gopdf.ToByte("a")]=505
	me.cw[gopdf.ToByte("b")]=523
	me.cw[gopdf.ToByte("c")]=487
	me.cw[gopdf.ToByte("d")]=523
	me.cw[gopdf.ToByte("e")]=485
	me.cw[gopdf.ToByte("f")]=312
	me.cw[gopdf.ToByte("g")]=523
	me.cw[gopdf.ToByte("h")]=523
	me.cw[gopdf.ToByte("i")]=217
	me.cw[gopdf.ToByte("j")]=226
	me.cw[gopdf.ToByte("k")]=463
	me.cw[gopdf.ToByte("l")]=217
	me.cw[gopdf.ToByte("m")]=828
	me.cw[gopdf.ToByte("n")]=523
	me.cw[gopdf.ToByte("o")]=523
	me.cw[gopdf.ToByte("p")]=523
	me.cw[gopdf.ToByte("q")]=523
	me.cw[gopdf.ToByte("r")]=322
	me.cw[gopdf.ToByte("s")]=479
	me.cw[gopdf.ToByte("t")]=300
	me.cw[gopdf.ToByte("u")]=523
	me.cw[gopdf.ToByte("v")]=461
	me.cw[gopdf.ToByte("w")]=709
	me.cw[gopdf.ToByte("x")]=461
	me.cw[gopdf.ToByte("y")]=461
	me.cw[gopdf.ToByte("z")]=461
	me.cw[gopdf.ToByte("{")]=313
	me.cw[gopdf.ToByte("|")]=212
	me.cw[gopdf.ToByte("}")]=313
	me.cw[gopdf.ToByte("~")]=641
	me.cw[gopdf.Chr(127)]=231
	me.cw[gopdf.Chr(128)]=493
	me.cw[gopdf.Chr(129)]=231
	me.cw[gopdf.Chr(130)]=167
	me.cw[gopdf.Chr(131)]=310
	me.cw[gopdf.Chr(132)]=278
	me.cw[gopdf.Chr(133)]=603
	me.cw[gopdf.Chr(134)]=517
	me.cw[gopdf.Chr(135)]=529
	me.cw[gopdf.Chr(136)]=407
	me.cw[gopdf.Chr(137)]=911
	me.cw[gopdf.Chr(138)]=564
	me.cw[gopdf.Chr(139)]=282
	me.cw[gopdf.Chr(140)]=875
	me.cw[gopdf.Chr(141)]=231
	me.cw[gopdf.Chr(142)]=561
	me.cw[gopdf.Chr(143)]=231
	me.cw[gopdf.Chr(144)]=231
	me.cw[gopdf.Chr(145)]=172
	me.cw[gopdf.Chr(146)]=172
	me.cw[gopdf.Chr(147)]=287
	me.cw[gopdf.Chr(148)]=289
	me.cw[gopdf.Chr(149)]=307
	me.cw[gopdf.Chr(150)]=647
	me.cw[gopdf.Chr(151)]=762
	me.cw[gopdf.Chr(152)]=413
	me.cw[gopdf.Chr(153)]=580
	me.cw[gopdf.Chr(154)]=479
	me.cw[gopdf.Chr(155)]=282
	me.cw[gopdf.Chr(156)]=863
	me.cw[gopdf.Chr(157)]=231
	me.cw[gopdf.Chr(158)]=461
	me.cw[gopdf.Chr(159)]=570
	me.cw[gopdf.Chr(160)]=231
	me.cw[gopdf.Chr(161)]=212
	me.cw[gopdf.Chr(162)]=511
	me.cw[gopdf.Chr(163)]=537
	me.cw[gopdf.Chr(164)]=677
	me.cw[gopdf.Chr(165)]=561
	me.cw[gopdf.Chr(166)]=208
	me.cw[gopdf.Chr(167)]=568
	me.cw[gopdf.Chr(168)]=442
	me.cw[gopdf.Chr(169)]=747
	me.cw[gopdf.Chr(170)]=417
	me.cw[gopdf.Chr(171)]=430
	me.cw[gopdf.Chr(172)]=513
	me.cw[gopdf.Chr(173)]=271
	me.cw[gopdf.Chr(174)]=750
	me.cw[gopdf.Chr(175)]=404
	me.cw[gopdf.Chr(176)]=356
	me.cw[gopdf.Chr(177)]=499
	me.cw[gopdf.Chr(178)]=385
	me.cw[gopdf.Chr(179)]=393
	me.cw[gopdf.Chr(180)]=269
	me.cw[gopdf.Chr(181)]=523
	me.cw[gopdf.Chr(182)]=448
	me.cw[gopdf.Chr(183)]=234
	me.cw[gopdf.Chr(184)]=231
	me.cw[gopdf.Chr(185)]=238
	me.cw[gopdf.Chr(186)]=424
	me.cw[gopdf.Chr(187)]=427
	me.cw[gopdf.Chr(188)]=715
	me.cw[gopdf.Chr(189)]=721
	me.cw[gopdf.Chr(190)]=784
	me.cw[gopdf.Chr(191)]=441
	me.cw[gopdf.Chr(192)]=592
	me.cw[gopdf.Chr(193)]=592
	me.cw[gopdf.Chr(194)]=592
	me.cw[gopdf.Chr(195)]=592
	me.cw[gopdf.Chr(196)]=592
	me.cw[gopdf.Chr(197)]=592
	me.cw[gopdf.Chr(198)]=854
	me.cw[gopdf.Chr(199)]=603
	me.cw[gopdf.Chr(200)]=541
	me.cw[gopdf.Chr(201)]=541
	me.cw[gopdf.Chr(202)]=541
	me.cw[gopdf.Chr(203)]=541
	me.cw[gopdf.Chr(204)]=257
	me.cw[gopdf.Chr(205)]=257
	me.cw[gopdf.Chr(206)]=257
	me.cw[gopdf.Chr(207)]=257
	me.cw[gopdf.Chr(208)]=639
	me.cw[gopdf.Chr(209)]=666
	me.cw[gopdf.Chr(210)]=639
	me.cw[gopdf.Chr(211)]=639
	me.cw[gopdf.Chr(212)]=639
	me.cw[gopdf.Chr(213)]=639
	me.cw[gopdf.Chr(214)]=639
	me.cw[gopdf.Chr(215)]=492
	me.cw[gopdf.Chr(216)]=632
	me.cw[gopdf.Chr(217)]=629
	me.cw[gopdf.Chr(218)]=629
	me.cw[gopdf.Chr(219)]=629
	me.cw[gopdf.Chr(220)]=629
	me.cw[gopdf.Chr(221)]=570
	me.cw[gopdf.Chr(222)]=556
	me.cw[gopdf.Chr(223)]=550
	me.cw[gopdf.Chr(224)]=505
	me.cw[gopdf.Chr(225)]=505
	me.cw[gopdf.Chr(226)]=505
	me.cw[gopdf.Chr(227)]=505
	me.cw[gopdf.Chr(228)]=505
	me.cw[gopdf.Chr(229)]=505
	me.cw[gopdf.Chr(230)]=792
	me.cw[gopdf.Chr(231)]=487
	me.cw[gopdf.Chr(232)]=485
	me.cw[gopdf.Chr(233)]=485
	me.cw[gopdf.Chr(234)]=485
	me.cw[gopdf.Chr(235)]=485
	me.cw[gopdf.Chr(236)]=214
	me.cw[gopdf.Chr(237)]=214
	me.cw[gopdf.Chr(238)]=214
	me.cw[gopdf.Chr(239)]=214
	me.cw[gopdf.Chr(240)]=546
	me.cw[gopdf.Chr(241)]=523
	me.cw[gopdf.Chr(242)]=523
	me.cw[gopdf.Chr(243)]=523
	me.cw[gopdf.Chr(244)]=523
	me.cw[gopdf.Chr(245)]=523
	me.cw[gopdf.Chr(246)]=523
	me.cw[gopdf.Chr(247)]=534
	me.cw[gopdf.Chr(248)]=524
	me.cw[gopdf.Chr(249)]=523
	me.cw[gopdf.Chr(250)]=523
	me.cw[gopdf.Chr(251)]=523
	me.cw[gopdf.Chr(252)]=523
	me.cw[gopdf.Chr(253)]=461
	me.cw[gopdf.Chr(254)]=529
	me.cw[gopdf.Chr(255)]=461
	me.up = -73
	me.ut = 49
	me.fonttype = "TrueType"
	me.name = "Roboto-LightItalic"
	me.enc = "cp1252"
	me.desc = make([]gopdf.FontDescItem,8)
	me.desc[0] =  gopdf.FontDescItem{ Key:"Ascent",Val : "750" }
	me.desc[1] =  gopdf.FontDescItem{ Key: "Descent", Val : "-250" }
	me.desc[2] =  gopdf.FontDescItem{ Key:"CapHeight", Val :  "711" }
	me.desc[3] =  gopdf.FontDescItem{ Key: "Flags", Val :  "32" }
	me.desc[4] =  gopdf.FontDescItem{ Key:"FontBBox", Val :  "[-407 -271 1147 1048]" }
	me.desc[5] =  gopdf.FontDescItem{ Key:"ItalicAngle", Val :  "0" }
	me.desc[6] =  gopdf.FontDescItem{ Key:"StemV", Val :  "70" }
 	me.desc[7] =  gopdf.FontDescItem{ Key:"MissingWidth", Val :  "231" } 
 }
func (me * RobotoLightItalic)GetType() string{
	return me.fonttype
}
func (me * RobotoLightItalic)GetName() string{
	return me.name
}	
func (me * RobotoLightItalic)GetDesc() []gopdf.FontDescItem{
	return me.desc
}
func (me * RobotoLightItalic)GetUp() int{
	return me.up
}
func (me * RobotoLightItalic)GetUt()  int{
	return me.ut
}
func (me * RobotoLightItalic)GetCw() gopdf.FontCw{
	return me.cw
}
func (me * RobotoLightItalic)GetEnc() string{
	return me.enc
}
func (me * RobotoLightItalic)GetDiff() string {
	return me.diff
}
func (me * RobotoLightItalic) GetOriginalsize() int{
	return 98764
}
func (me * RobotoLightItalic)  SetFamily(family string){
	me.family = family
}
func (me * RobotoLightItalic) 	GetFamily() string{
	return me.family
}
