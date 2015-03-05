package fonts //change this
import (
	"github.com/yaslama/gopdf"
)
type RobotoThinItalic struct {
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
func (me * RobotoThinItalic) Init(){
	me.cw = make(gopdf.FontCw)
	me.cw[gopdf.Chr(0)]=5
	me.cw[gopdf.Chr(1)]=5
	me.cw[gopdf.Chr(2)]=5
	me.cw[gopdf.Chr(3)]=5
	me.cw[gopdf.Chr(4)]=5
	me.cw[gopdf.Chr(5)]=5
	me.cw[gopdf.Chr(6)]=5
	me.cw[gopdf.Chr(7)]=5
	me.cw[gopdf.Chr(8)]=5
	me.cw[gopdf.Chr(9)]=5
	me.cw[gopdf.Chr(10)]=5
	me.cw[gopdf.Chr(11)]=5
	me.cw[gopdf.Chr(12)]=5
	me.cw[gopdf.Chr(13)]=5
	me.cw[gopdf.Chr(14)]=5
	me.cw[gopdf.Chr(15)]=5
	me.cw[gopdf.Chr(16)]=5
	me.cw[gopdf.Chr(17)]=5
	me.cw[gopdf.Chr(18)]=5
	me.cw[gopdf.Chr(19)]=5
	me.cw[gopdf.Chr(20)]=5
	me.cw[gopdf.Chr(21)]=5
	me.cw[gopdf.Chr(22)]=5
	me.cw[gopdf.Chr(23)]=5
	me.cw[gopdf.Chr(24)]=5
	me.cw[gopdf.Chr(25)]=5
	me.cw[gopdf.Chr(26)]=5
	me.cw[gopdf.Chr(27)]=5
	me.cw[gopdf.Chr(28)]=5
	me.cw[gopdf.Chr(29)]=5
	me.cw[gopdf.Chr(30)]=5
	me.cw[gopdf.Chr(31)]=5
	me.cw[gopdf.ToByte(" ")]=228
	me.cw[gopdf.ToByte("!")]=185
	me.cw[gopdf.ToByte("\"")]=242
	me.cw[gopdf.ToByte("#")]=579
	me.cw[gopdf.ToByte("$")]=514
	me.cw[gopdf.ToByte("%")]=698
	me.cw[gopdf.ToByte("&")]=571
	me.cw[gopdf.ToByte("'")]=159
	me.cw[gopdf.ToByte("(")]=281
	me.cw[gopdf.ToByte(")")]=288
	me.cw[gopdf.ToByte("*")]=394
	me.cw[gopdf.ToByte("+")]=527
	me.cw[gopdf.ToByte(",")]=179
	me.cw[gopdf.ToByte("-")]=281
	me.cw[gopdf.ToByte(".")]=204
	me.cw[gopdf.ToByte("/")]=360
	me.cw[gopdf.ToByte("0")]=514
	me.cw[gopdf.ToByte("1")]=514
	me.cw[gopdf.ToByte("2")]=514
	me.cw[gopdf.ToByte("3")]=514
	me.cw[gopdf.ToByte("4")]=514
	me.cw[gopdf.ToByte("5")]=514
	me.cw[gopdf.ToByte("6")]=514
	me.cw[gopdf.ToByte("7")]=514
	me.cw[gopdf.ToByte("8")]=514
	me.cw[gopdf.ToByte("9")]=514
	me.cw[gopdf.ToByte(":")]=170
	me.cw[gopdf.ToByte(";")]=170
	me.cw[gopdf.ToByte("<")]=483
	me.cw[gopdf.ToByte("=")]=523
	me.cw[gopdf.ToByte(">")]=483
	me.cw[gopdf.ToByte("?")]=410
	me.cw[gopdf.ToByte("@")]=869
	me.cw[gopdf.ToByte("A")]=560
	me.cw[gopdf.ToByte("B")]=566
	me.cw[gopdf.ToByte("C")]=607
	me.cw[gopdf.ToByte("D")]=613
	me.cw[gopdf.ToByte("E")]=534
	me.cw[gopdf.ToByte("F")]=537
	me.cw[gopdf.ToByte("G")]=644
	me.cw[gopdf.ToByte("H")]=658
	me.cw[gopdf.ToByte("I")]=248
	me.cw[gopdf.ToByte("J")]=516
	me.cw[gopdf.ToByte("K")]=538
	me.cw[gopdf.ToByte("L")]=485
	me.cw[gopdf.ToByte("M")]=802
	me.cw[gopdf.ToByte("N")]=663
	me.cw[gopdf.ToByte("O")]=625
	me.cw[gopdf.ToByte("P")]=563
	me.cw[gopdf.ToByte("Q")]=625
	me.cw[gopdf.ToByte("R")]=613
	me.cw[gopdf.ToByte("S")]=561
	me.cw[gopdf.ToByte("T")]=561
	me.cw[gopdf.ToByte("U")]=625
	me.cw[gopdf.ToByte("V")]=560
	me.cw[gopdf.ToByte("W")]=847
	me.cw[gopdf.ToByte("X")]=560
	me.cw[gopdf.ToByte("Y")]=560
	me.cw[gopdf.ToByte("Z")]=505
	me.cw[gopdf.ToByte("[")]=204
	me.cw[gopdf.ToByte("\\")]=356
	me.cw[gopdf.ToByte("]")]=204
	me.cw[gopdf.ToByte("^")]=391
	me.cw[gopdf.ToByte("_")]=389
	me.cw[gopdf.ToByte("`")]=249
	me.cw[gopdf.ToByte("a")]=496
	me.cw[gopdf.ToByte("b")]=514
	me.cw[gopdf.ToByte("c")]=477
	me.cw[gopdf.ToByte("d")]=514
	me.cw[gopdf.ToByte("e")]=474
	me.cw[gopdf.ToByte("f")]=298
	me.cw[gopdf.ToByte("g")]=515
	me.cw[gopdf.ToByte("h")]=514
	me.cw[gopdf.ToByte("i")]=196
	me.cw[gopdf.ToByte("j")]=207
	me.cw[gopdf.ToByte("k")]=445
	me.cw[gopdf.ToByte("l")]=196
	me.cw[gopdf.ToByte("m")]=838
	me.cw[gopdf.ToByte("n")]=514
	me.cw[gopdf.ToByte("o")]=514
	me.cw[gopdf.ToByte("p")]=514
	me.cw[gopdf.ToByte("q")]=514
	me.cw[gopdf.ToByte("r")]=313
	me.cw[gopdf.ToByte("s")]=468
	me.cw[gopdf.ToByte("t")]=299
	me.cw[gopdf.ToByte("u")]=514
	me.cw[gopdf.ToByte("v")]=449
	me.cw[gopdf.ToByte("w")]=709
	me.cw[gopdf.ToByte("x")]=449
	me.cw[gopdf.ToByte("y")]=449
	me.cw[gopdf.ToByte("z")]=449
	me.cw[gopdf.ToByte("{")]=305
	me.cw[gopdf.ToByte("|")]=189
	me.cw[gopdf.ToByte("}")]=305
	me.cw[gopdf.ToByte("~")]=646
	me.cw[gopdf.Chr(127)]=5
	me.cw[gopdf.Chr(128)]=486
	me.cw[gopdf.Chr(129)]=5
	me.cw[gopdf.Chr(130)]=141
	me.cw[gopdf.Chr(131)]=297
	me.cw[gopdf.Chr(132)]=223
	me.cw[gopdf.Chr(133)]=575
	me.cw[gopdf.Chr(134)]=516
	me.cw[gopdf.Chr(135)]=523
	me.cw[gopdf.Chr(136)]=366
	me.cw[gopdf.Chr(137)]=929
	me.cw[gopdf.Chr(138)]=555
	me.cw[gopdf.Chr(139)]=279
	me.cw[gopdf.Chr(140)]=855
	me.cw[gopdf.Chr(141)]=5
	me.cw[gopdf.Chr(142)]=561
	me.cw[gopdf.Chr(143)]=5
	me.cw[gopdf.Chr(144)]=5
	me.cw[gopdf.Chr(145)]=150
	me.cw[gopdf.Chr(146)]=150
	me.cw[gopdf.Chr(147)]=232
	me.cw[gopdf.Chr(148)]=233
	me.cw[gopdf.Chr(149)]=293
	me.cw[gopdf.Chr(150)]=646
	me.cw[gopdf.Chr(151)]=765
	me.cw[gopdf.Chr(152)]=376
	me.cw[gopdf.Chr(153)]=571
	me.cw[gopdf.Chr(154)]=468
	me.cw[gopdf.Chr(155)]=279
	me.cw[gopdf.Chr(156)]=878
	me.cw[gopdf.Chr(157)]=5
	me.cw[gopdf.Chr(158)]=449
	me.cw[gopdf.Chr(159)]=560
	me.cw[gopdf.Chr(160)]=228
	me.cw[gopdf.Chr(161)]=189
	me.cw[gopdf.Chr(162)]=509
	me.cw[gopdf.Chr(163)]=527
	me.cw[gopdf.Chr(164)]=687
	me.cw[gopdf.Chr(165)]=553
	me.cw[gopdf.Chr(166)]=186
	me.cw[gopdf.Chr(167)]=560
	me.cw[gopdf.Chr(168)]=417
	me.cw[gopdf.Chr(169)]=761
	me.cw[gopdf.Chr(170)]=413
	me.cw[gopdf.Chr(171)]=416
	me.cw[gopdf.Chr(172)]=505
	me.cw[gopdf.Chr(173)]=281
	me.cw[gopdf.Chr(174)]=765
	me.cw[gopdf.Chr(175)]=373
	me.cw[gopdf.Chr(176)]=360
	me.cw[gopdf.Chr(177)]=494
	me.cw[gopdf.Chr(178)]=372
	me.cw[gopdf.Chr(179)]=382
	me.cw[gopdf.Chr(180)]=237
	me.cw[gopdf.Chr(181)]=514
	me.cw[gopdf.Chr(182)]=435
	me.cw[gopdf.Chr(183)]=219
	me.cw[gopdf.Chr(184)]=228
	me.cw[gopdf.Chr(185)]=219
	me.cw[gopdf.Chr(186)]=420
	me.cw[gopdf.Chr(187)]=412
	me.cw[gopdf.Chr(188)]=701
	me.cw[gopdf.Chr(189)]=662
	me.cw[gopdf.Chr(190)]=755
	me.cw[gopdf.Chr(191)]=419
	me.cw[gopdf.Chr(192)]=560
	me.cw[gopdf.Chr(193)]=560
	me.cw[gopdf.Chr(194)]=560
	me.cw[gopdf.Chr(195)]=560
	me.cw[gopdf.Chr(196)]=560
	me.cw[gopdf.Chr(197)]=560
	me.cw[gopdf.Chr(198)]=830
	me.cw[gopdf.Chr(199)]=607
	me.cw[gopdf.Chr(200)]=534
	me.cw[gopdf.Chr(201)]=534
	me.cw[gopdf.Chr(202)]=534
	me.cw[gopdf.Chr(203)]=534
	me.cw[gopdf.Chr(204)]=248
	me.cw[gopdf.Chr(205)]=248
	me.cw[gopdf.Chr(206)]=248
	me.cw[gopdf.Chr(207)]=248
	me.cw[gopdf.Chr(208)]=628
	me.cw[gopdf.Chr(209)]=663
	me.cw[gopdf.Chr(210)]=625
	me.cw[gopdf.Chr(211)]=625
	me.cw[gopdf.Chr(212)]=625
	me.cw[gopdf.Chr(213)]=625
	me.cw[gopdf.Chr(214)]=625
	me.cw[gopdf.Chr(215)]=482
	me.cw[gopdf.Chr(216)]=625
	me.cw[gopdf.Chr(217)]=625
	me.cw[gopdf.Chr(218)]=625
	me.cw[gopdf.Chr(219)]=625
	me.cw[gopdf.Chr(220)]=625
	me.cw[gopdf.Chr(221)]=560
	me.cw[gopdf.Chr(222)]=559
	me.cw[gopdf.Chr(223)]=542
	me.cw[gopdf.Chr(224)]=496
	me.cw[gopdf.Chr(225)]=496
	me.cw[gopdf.Chr(226)]=496
	me.cw[gopdf.Chr(227)]=496
	me.cw[gopdf.Chr(228)]=496
	me.cw[gopdf.Chr(229)]=496
	me.cw[gopdf.Chr(230)]=793
	me.cw[gopdf.Chr(231)]=477
	me.cw[gopdf.Chr(232)]=474
	me.cw[gopdf.Chr(233)]=474
	me.cw[gopdf.Chr(234)]=474
	me.cw[gopdf.Chr(235)]=474
	me.cw[gopdf.Chr(236)]=189
	me.cw[gopdf.Chr(237)]=189
	me.cw[gopdf.Chr(238)]=189
	me.cw[gopdf.Chr(239)]=189
	me.cw[gopdf.Chr(240)]=542
	me.cw[gopdf.Chr(241)]=514
	me.cw[gopdf.Chr(242)]=514
	me.cw[gopdf.Chr(243)]=514
	me.cw[gopdf.Chr(244)]=514
	me.cw[gopdf.Chr(245)]=514
	me.cw[gopdf.Chr(246)]=514
	me.cw[gopdf.Chr(247)]=531
	me.cw[gopdf.Chr(248)]=514
	me.cw[gopdf.Chr(249)]=514
	me.cw[gopdf.Chr(250)]=514
	me.cw[gopdf.Chr(251)]=514
	me.cw[gopdf.Chr(252)]=514
	me.cw[gopdf.Chr(253)]=449
	me.cw[gopdf.Chr(254)]=514
	me.cw[gopdf.Chr(255)]=449
	me.up = -73
	me.ut = 49
	me.fonttype = "TrueType"
	me.name = "Roboto-ThinItalic"
	me.enc = "cp1252"
	me.desc = make([]gopdf.FontDescItem,8)
	me.desc[0] =  gopdf.FontDescItem{ Key:"Ascent",Val : "750" }
	me.desc[1] =  gopdf.FontDescItem{ Key: "Descent", Val : "-250" }
	me.desc[2] =  gopdf.FontDescItem{ Key:"CapHeight", Val :  "711" }
	me.desc[3] =  gopdf.FontDescItem{ Key: "Flags", Val :  "32" }
	me.desc[4] =  gopdf.FontDescItem{ Key:"FontBBox", Val :  "[-352 -271 1120 1048]" }
	me.desc[5] =  gopdf.FontDescItem{ Key:"ItalicAngle", Val :  "0" }
	me.desc[6] =  gopdf.FontDescItem{ Key:"StemV", Val :  "70" }
 	me.desc[7] =  gopdf.FontDescItem{ Key:"MissingWidth", Val :  "5" } 
 }
func (me * RobotoThinItalic)GetType() string{
	return me.fonttype
}
func (me * RobotoThinItalic)GetName() string{
	return me.name
}	
func (me * RobotoThinItalic)GetDesc() []gopdf.FontDescItem{
	return me.desc
}
func (me * RobotoThinItalic)GetUp() int{
	return me.up
}
func (me * RobotoThinItalic)GetUt()  int{
	return me.ut
}
func (me * RobotoThinItalic)GetCw() gopdf.FontCw{
	return me.cw
}
func (me * RobotoThinItalic)GetEnc() string{
	return me.enc
}
func (me * RobotoThinItalic)GetDiff() string {
	return me.diff
}
func (me * RobotoThinItalic) GetOriginalsize() int{
	return 98764
}
func (me * RobotoThinItalic)  SetFamily(family string){
	me.family = family
}
func (me * RobotoThinItalic) 	GetFamily() string{
	return me.family
}
