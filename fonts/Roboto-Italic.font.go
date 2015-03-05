package fonts //change this
import (
	"github.com/yaslama/gopdf"
)
type Roboto-Italic struct {
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
func (me * Roboto-Italic) Init(){
	me.cw = make(gopdf.FontCw)
	me.cw[gopdf.Chr(0)]=236
	me.cw[gopdf.Chr(1)]=236
	me.cw[gopdf.Chr(2)]=236
	me.cw[gopdf.Chr(3)]=236
	me.cw[gopdf.Chr(4)]=236
	me.cw[gopdf.Chr(5)]=236
	me.cw[gopdf.Chr(6)]=236
	me.cw[gopdf.Chr(7)]=236
	me.cw[gopdf.Chr(8)]=236
	me.cw[gopdf.Chr(9)]=236
	me.cw[gopdf.Chr(10)]=236
	me.cw[gopdf.Chr(11)]=236
	me.cw[gopdf.Chr(12)]=236
	me.cw[gopdf.Chr(13)]=236
	me.cw[gopdf.Chr(14)]=236
	me.cw[gopdf.Chr(15)]=236
	me.cw[gopdf.Chr(16)]=236
	me.cw[gopdf.Chr(17)]=236
	me.cw[gopdf.Chr(18)]=236
	me.cw[gopdf.Chr(19)]=236
	me.cw[gopdf.Chr(20)]=236
	me.cw[gopdf.Chr(21)]=236
	me.cw[gopdf.Chr(22)]=236
	me.cw[gopdf.Chr(23)]=236
	me.cw[gopdf.Chr(24)]=236
	me.cw[gopdf.Chr(25)]=236
	me.cw[gopdf.Chr(26)]=236
	me.cw[gopdf.Chr(27)]=236
	me.cw[gopdf.Chr(28)]=236
	me.cw[gopdf.Chr(29)]=236
	me.cw[gopdf.Chr(30)]=236
	me.cw[gopdf.Chr(31)]=236
	me.cw[gopdf.ToByte(" ")]=236
	me.cw[gopdf.ToByte("!")]=249
	me.cw[gopdf.ToByte("\"")]=307
	me.cw[gopdf.ToByte("#")]=584
	me.cw[gopdf.ToByte("$")]=528
	me.cw[gopdf.ToByte("%")]=685
	me.cw[gopdf.ToByte("&")]=584
	me.cw[gopdf.ToByte("'")]=167
	me.cw[gopdf.ToByte("(")]=313
	me.cw[gopdf.ToByte(")")]=316
	me.cw[gopdf.ToByte("*")]=406
	me.cw[gopdf.ToByte("+")]=532
	me.cw[gopdf.ToByte(",")]=188
	me.cw[gopdf.ToByte("-")]=260
	me.cw[gopdf.ToByte(".")]=253
	me.cw[gopdf.ToByte("/")]=390
	me.cw[gopdf.ToByte("0")]=528
	me.cw[gopdf.ToByte("1")]=528
	me.cw[gopdf.ToByte("2")]=528
	me.cw[gopdf.ToByte("3")]=528
	me.cw[gopdf.ToByte("4")]=528
	me.cw[gopdf.ToByte("5")]=528
	me.cw[gopdf.ToByte("6")]=528
	me.cw[gopdf.ToByte("7")]=528
	me.cw[gopdf.ToByte("8")]=528
	me.cw[gopdf.ToByte("9")]=528
	me.cw[gopdf.ToByte(":")]=239
	me.cw[gopdf.ToByte(";")]=243
	me.cw[gopdf.ToByte("<")]=477
	me.cw[gopdf.ToByte("=")]=528
	me.cw[gopdf.ToByte(">")]=491
	me.cw[gopdf.ToByte("?")]=446
	me.cw[gopdf.ToByte("@")]=838
	me.cw[gopdf.ToByte("A")]=624
	me.cw[gopdf.ToByte("B")]=599
	me.cw[gopdf.ToByte("C")]=594
	me.cw[gopdf.ToByte("D")]=617
	me.cw[gopdf.ToByte("E")]=548
	me.cw[gopdf.ToByte("F")]=546
	me.cw[gopdf.ToByte("G")]=644
	me.cw[gopdf.ToByte("H")]=668
	me.cw[gopdf.ToByte("I")]=268
	me.cw[gopdf.ToByte("J")]=518
	me.cw[gopdf.ToByte("K")]=562
	me.cw[gopdf.ToByte("L")]=516
	me.cw[gopdf.ToByte("M")]=821
	me.cw[gopdf.ToByte("N")]=668
	me.cw[gopdf.ToByte("O")]=639
	me.cw[gopdf.ToByte("P")]=600
	me.cw[gopdf.ToByte("Q")]=654
	me.cw[gopdf.ToByte("R")]=621
	me.cw[gopdf.ToByte("S")]=569
	me.cw[gopdf.ToByte("T")]=560
	me.cw[gopdf.ToByte("U")]=634
	me.cw[gopdf.ToByte("V")]=608
	me.cw[gopdf.ToByte("W")]=823
	me.cw[gopdf.ToByte("X")]=592
	me.cw[gopdf.ToByte("Y")]=581
	me.cw[gopdf.ToByte("Z")]=524
	me.cw[gopdf.ToByte("[")]=255
	me.cw[gopdf.ToByte("\\")]=389
	me.cw[gopdf.ToByte("]")]=255
	me.cw[gopdf.ToByte("^")]=394
	me.cw[gopdf.ToByte("_")]=427
	me.cw[gopdf.ToByte("`")]=295
	me.cw[gopdf.ToByte("a")]=516
	me.cw[gopdf.ToByte("b")]=533
	me.cw[gopdf.ToByte("c")]=498
	me.cw[gopdf.ToByte("d")]=533
	me.cw[gopdf.ToByte("e")]=497
	me.cw[gopdf.ToByte("f")]=326
	me.cw[gopdf.ToByte("g")]=533
	me.cw[gopdf.ToByte("h")]=533
	me.cw[gopdf.ToByte("i")]=238
	me.cw[gopdf.ToByte("j")]=245
	me.cw[gopdf.ToByte("k")]=481
	me.cw[gopdf.ToByte("l")]=238
	me.cw[gopdf.ToByte("m")]=818
	me.cw[gopdf.ToByte("n")]=533
	me.cw[gopdf.ToByte("o")]=533
	me.cw[gopdf.ToByte("p")]=533
	me.cw[gopdf.ToByte("q")]=533
	me.cw[gopdf.ToByte("r")]=330
	me.cw[gopdf.ToByte("s")]=491
	me.cw[gopdf.ToByte("t")]=301
	me.cw[gopdf.ToByte("u")]=533
	me.cw[gopdf.ToByte("v")]=473
	me.cw[gopdf.ToByte("w")]=709
	me.cw[gopdf.ToByte("x")]=473
	me.cw[gopdf.ToByte("y")]=473
	me.cw[gopdf.ToByte("z")]=473
	me.cw[gopdf.ToByte("{")]=321
	me.cw[gopdf.ToByte("|")]=235
	me.cw[gopdf.ToByte("}")]=321
	me.cw[gopdf.ToByte("~")]=637
	me.cw[gopdf.Chr(127)]=236
	me.cw[gopdf.Chr(128)]=499
	me.cw[gopdf.Chr(129)]=236
	me.cw[gopdf.Chr(130)]=193
	me.cw[gopdf.Chr(131)]=323
	me.cw[gopdf.Chr(132)]=333
	me.cw[gopdf.Chr(133)]=631
	me.cw[gopdf.Chr(134)]=518
	me.cw[gopdf.Chr(135)]=536
	me.cw[gopdf.Chr(136)]=448
	me.cw[gopdf.Chr(137)]=894
	me.cw[gopdf.Chr(138)]=574
	me.cw[gopdf.Chr(139)]=284
	me.cw[gopdf.Chr(140)]=895
	me.cw[gopdf.Chr(141)]=236
	me.cw[gopdf.Chr(142)]=561
	me.cw[gopdf.Chr(143)]=236
	me.cw[gopdf.Chr(144)]=236
	me.cw[gopdf.Chr(145)]=194
	me.cw[gopdf.Chr(146)]=194
	me.cw[gopdf.Chr(147)]=341
	me.cw[gopdf.Chr(148)]=345
	me.cw[gopdf.Chr(149)]=320
	me.cw[gopdf.Chr(150)]=648
	me.cw[gopdf.Chr(151)]=759
	me.cw[gopdf.Chr(152)]=449
	me.cw[gopdf.Chr(153)]=588
	me.cw[gopdf.Chr(154)]=491
	me.cw[gopdf.Chr(155)]=284
	me.cw[gopdf.Chr(156)]=847
	me.cw[gopdf.Chr(157)]=236
	me.cw[gopdf.Chr(158)]=473
	me.cw[gopdf.Chr(159)]=581
	me.cw[gopdf.Chr(160)]=236
	me.cw[gopdf.Chr(161)]=235
	me.cw[gopdf.Chr(162)]=514
	me.cw[gopdf.Chr(163)]=547
	me.cw[gopdf.Chr(164)]=667
	me.cw[gopdf.Chr(165)]=569
	me.cw[gopdf.Chr(166)]=231
	me.cw[gopdf.Chr(167)]=576
	me.cw[gopdf.Chr(168)]=468
	me.cw[gopdf.Chr(169)]=733
	me.cw[gopdf.Chr(170)]=421
	me.cw[gopdf.Chr(171)]=443
	me.cw[gopdf.Chr(172)]=521
	me.cw[gopdf.Chr(173)]=260
	me.cw[gopdf.Chr(174)]=734
	me.cw[gopdf.Chr(175)]=435
	me.cw[gopdf.Chr(176)]=352
	me.cw[gopdf.Chr(177)]=503
	me.cw[gopdf.Chr(178)]=399
	me.cw[gopdf.Chr(179)]=402
	me.cw[gopdf.Chr(180)]=301
	me.cw[gopdf.Chr(181)]=533
	me.cw[gopdf.Chr(182)]=461
	me.cw[gopdf.Chr(183)]=250
	me.cw[gopdf.Chr(184)]=236
	me.cw[gopdf.Chr(185)]=257
	me.cw[gopdf.Chr(186)]=428
	me.cw[gopdf.Chr(187)]=443
	me.cw[gopdf.Chr(188)]=729
	me.cw[gopdf.Chr(189)]=779
	me.cw[gopdf.Chr(190)]=813
	me.cw[gopdf.Chr(191)]=464
	me.cw[gopdf.Chr(192)]=624
	me.cw[gopdf.Chr(193)]=624
	me.cw[gopdf.Chr(194)]=624
	me.cw[gopdf.Chr(195)]=624
	me.cw[gopdf.Chr(196)]=624
	me.cw[gopdf.Chr(197)]=624
	me.cw[gopdf.Chr(198)]=877
	me.cw[gopdf.Chr(199)]=599
	me.cw[gopdf.Chr(200)]=548
	me.cw[gopdf.Chr(201)]=548
	me.cw[gopdf.Chr(202)]=548
	me.cw[gopdf.Chr(203)]=548
	me.cw[gopdf.Chr(204)]=268
	me.cw[gopdf.Chr(205)]=268
	me.cw[gopdf.Chr(206)]=268
	me.cw[gopdf.Chr(207)]=268
	me.cw[gopdf.Chr(208)]=650
	me.cw[gopdf.Chr(209)]=668
	me.cw[gopdf.Chr(210)]=654
	me.cw[gopdf.Chr(211)]=654
	me.cw[gopdf.Chr(212)]=654
	me.cw[gopdf.Chr(213)]=654
	me.cw[gopdf.Chr(214)]=654
	me.cw[gopdf.Chr(215)]=502
	me.cw[gopdf.Chr(216)]=639
	me.cw[gopdf.Chr(217)]=634
	me.cw[gopdf.Chr(218)]=634
	me.cw[gopdf.Chr(219)]=634
	me.cw[gopdf.Chr(220)]=634
	me.cw[gopdf.Chr(221)]=581
	me.cw[gopdf.Chr(222)]=554
	me.cw[gopdf.Chr(223)]=559
	me.cw[gopdf.Chr(224)]=516
	me.cw[gopdf.Chr(225)]=516
	me.cw[gopdf.Chr(226)]=516
	me.cw[gopdf.Chr(227)]=516
	me.cw[gopdf.Chr(228)]=516
	me.cw[gopdf.Chr(229)]=516
	me.cw[gopdf.Chr(230)]=790
	me.cw[gopdf.Chr(231)]=498
	me.cw[gopdf.Chr(232)]=497
	me.cw[gopdf.Chr(233)]=497
	me.cw[gopdf.Chr(234)]=497
	me.cw[gopdf.Chr(235)]=497
	me.cw[gopdf.Chr(236)]=238
	me.cw[gopdf.Chr(237)]=238
	me.cw[gopdf.Chr(238)]=238
	me.cw[gopdf.Chr(239)]=238
	me.cw[gopdf.Chr(240)]=550
	me.cw[gopdf.Chr(241)]=533
	me.cw[gopdf.Chr(242)]=533
	me.cw[gopdf.Chr(243)]=533
	me.cw[gopdf.Chr(244)]=533
	me.cw[gopdf.Chr(245)]=533
	me.cw[gopdf.Chr(246)]=533
	me.cw[gopdf.Chr(247)]=536
	me.cw[gopdf.Chr(248)]=534
	me.cw[gopdf.Chr(249)]=533
	me.cw[gopdf.Chr(250)]=533
	me.cw[gopdf.Chr(251)]=533
	me.cw[gopdf.Chr(252)]=533
	me.cw[gopdf.Chr(253)]=473
	me.cw[gopdf.Chr(254)]=543
	me.cw[gopdf.Chr(255)]=473
	me.up = -73
	me.ut = 49
	me.fonttype = "TrueType"
	me.name = "Roboto-Italic"
	me.enc = "cp1252"
	me.desc = make([]gopdf.FontDescItem,8)
	me.desc[0] =  gopdf.FontDescItem{ Key:"Ascent",Val : "750" }
	me.desc[1] =  gopdf.FontDescItem{ Key: "Descent", Val : "-250" }
	me.desc[2] =  gopdf.FontDescItem{ Key:"CapHeight", Val :  "711" }
	me.desc[3] =  gopdf.FontDescItem{ Key: "Flags", Val :  "32" }
	me.desc[4] =  gopdf.FontDescItem{ Key:"FontBBox", Val :  "[-661 -5517 5164 1048]" }
	me.desc[5] =  gopdf.FontDescItem{ Key:"ItalicAngle", Val :  "0" }
	me.desc[6] =  gopdf.FontDescItem{ Key:"StemV", Val :  "70" }
 	me.desc[7] =  gopdf.FontDescItem{ Key:"MissingWidth", Val :  "236" } 
 }
func (me * Roboto-Italic)GetType() string{
	return me.fonttype
}
func (me * Roboto-Italic)GetName() string{
	return me.name
}	
func (me * Roboto-Italic)GetDesc() []gopdf.FontDescItem{
	return me.desc
}
func (me * Roboto-Italic)GetUp() int{
	return me.up
}
func (me * Roboto-Italic)GetUt()  int{
	return me.ut
}
func (me * Roboto-Italic)GetCw() gopdf.FontCw{
	return me.cw
}
func (me * Roboto-Italic)GetEnc() string{
	return me.enc
}
func (me * Roboto-Italic)GetDiff() string {
	return me.diff
}
func (me * Roboto-Italic) GetOriginalsize() int{
	return 98764
}
func (me * Roboto-Italic)  SetFamily(family string){
	me.family = family
}
func (me * Roboto-Italic) 	GetFamily() string{
	return me.family
}
