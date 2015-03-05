package fonts //change this
import (
	"github.com/yaslama/gopdf"
)
type Roboto-Light struct {
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
func (me * Roboto-Light) Init(){
	me.cw = make(gopdf.FontCw)
	me.cw[gopdf.Chr(0)]=243
	me.cw[gopdf.Chr(1)]=243
	me.cw[gopdf.Chr(2)]=243
	me.cw[gopdf.Chr(3)]=243
	me.cw[gopdf.Chr(4)]=243
	me.cw[gopdf.Chr(5)]=243
	me.cw[gopdf.Chr(6)]=243
	me.cw[gopdf.Chr(7)]=243
	me.cw[gopdf.Chr(8)]=243
	me.cw[gopdf.Chr(9)]=243
	me.cw[gopdf.Chr(10)]=243
	me.cw[gopdf.Chr(11)]=243
	me.cw[gopdf.Chr(12)]=243
	me.cw[gopdf.Chr(13)]=243
	me.cw[gopdf.Chr(14)]=243
	me.cw[gopdf.Chr(15)]=243
	me.cw[gopdf.Chr(16)]=243
	me.cw[gopdf.Chr(17)]=243
	me.cw[gopdf.Chr(18)]=243
	me.cw[gopdf.Chr(19)]=243
	me.cw[gopdf.Chr(20)]=243
	me.cw[gopdf.Chr(21)]=243
	me.cw[gopdf.Chr(22)]=243
	me.cw[gopdf.Chr(23)]=243
	me.cw[gopdf.Chr(24)]=243
	me.cw[gopdf.Chr(25)]=243
	me.cw[gopdf.Chr(26)]=243
	me.cw[gopdf.Chr(27)]=243
	me.cw[gopdf.Chr(28)]=243
	me.cw[gopdf.Chr(29)]=243
	me.cw[gopdf.Chr(30)]=243
	me.cw[gopdf.Chr(31)]=243
	me.cw[gopdf.ToByte(" ")]=243
	me.cw[gopdf.ToByte("!")]=226
	me.cw[gopdf.ToByte("\"")]=287
	me.cw[gopdf.ToByte("#")]=620
	me.cw[gopdf.ToByte("$")]=554
	me.cw[gopdf.ToByte("%")]=739
	me.cw[gopdf.ToByte("&")]=615
	me.cw[gopdf.ToByte("'")]=170
	me.cw[gopdf.ToByte("(")]=313
	me.cw[gopdf.ToByte(")")]=318
	me.cw[gopdf.ToByte("*")]=424
	me.cw[gopdf.ToByte("+")]=564
	me.cw[gopdf.ToByte(",")]=191
	me.cw[gopdf.ToByte("-")]=286
	me.cw[gopdf.ToByte(".")]=239
	me.cw[gopdf.ToByte("/")]=397
	me.cw[gopdf.ToByte("0")]=554
	me.cw[gopdf.ToByte("1")]=554
	me.cw[gopdf.ToByte("2")]=554
	me.cw[gopdf.ToByte("3")]=554
	me.cw[gopdf.ToByte("4")]=554
	me.cw[gopdf.ToByte("5")]=554
	me.cw[gopdf.ToByte("6")]=554
	me.cw[gopdf.ToByte("7")]=554
	me.cw[gopdf.ToByte("8")]=554
	me.cw[gopdf.ToByte("9")]=554
	me.cw[gopdf.ToByte(":")]=212
	me.cw[gopdf.ToByte(";")]=214
	me.cw[gopdf.ToByte("<")]=511
	me.cw[gopdf.ToByte("=")]=560
	me.cw[gopdf.ToByte(">")]=518
	me.cw[gopdf.ToByte("?")]=454
	me.cw[gopdf.ToByte("@")]=913
	me.cw[gopdf.ToByte("A")]=629
	me.cw[gopdf.ToByte("B")]=620
	me.cw[gopdf.ToByte("C")]=644
	me.cw[gopdf.ToByte("D")]=665
	me.cw[gopdf.ToByte("E")]=576
	me.cw[gopdf.ToByte("F")]=577
	me.cw[gopdf.ToByte("G")]=687
	me.cw[gopdf.ToByte("H")]=707
	me.cw[gopdf.ToByte("I")]=271
	me.cw[gopdf.ToByte("J")]=550
	me.cw[gopdf.ToByte("K")]=639
	me.cw[gopdf.ToByte("L")]=531
	me.cw[gopdf.ToByte("M")]=867
	me.cw[gopdf.ToByte("N")]=710
	me.cw[gopdf.ToByte("O")]=681
	me.cw[gopdf.ToByte("P")]=619
	me.cw[gopdf.ToByte("Q")]=681
	me.cw[gopdf.ToByte("R")]=658
	me.cw[gopdf.ToByte("S")]=601
	me.cw[gopdf.ToByte("T")]=597
	me.cw[gopdf.ToByte("U")]=671
	me.cw[gopdf.ToByte("V")]=621
	me.cw[gopdf.ToByte("W")]=894
	me.cw[gopdf.ToByte("X")]=613
	me.cw[gopdf.ToByte("Y")]=607
	me.cw[gopdf.ToByte("Z")]=598
	me.cw[gopdf.ToByte("[")]=240
	me.cw[gopdf.ToByte("\\")]=394
	me.cw[gopdf.ToByte("]")]=240
	me.cw[gopdf.ToByte("^")]=416
	me.cw[gopdf.ToByte("_")]=432
	me.cw[gopdf.ToByte("`")]=286
	me.cw[gopdf.ToByte("a")]=538
	me.cw[gopdf.ToByte("b")]=557
	me.cw[gopdf.ToByte("c")]=518
	me.cw[gopdf.ToByte("d")]=557
	me.cw[gopdf.ToByte("e")]=515
	me.cw[gopdf.ToByte("f")]=329
	me.cw[gopdf.ToByte("g")]=557
	me.cw[gopdf.ToByte("h")]=557
	me.cw[gopdf.ToByte("i")]=227
	me.cw[gopdf.ToByte("j")]=236
	me.cw[gopdf.ToByte("k")]=491
	me.cw[gopdf.ToByte("l")]=227
	me.cw[gopdf.ToByte("m")]=886
	me.cw[gopdf.ToByte("n")]=557
	me.cw[gopdf.ToByte("o")]=557
	me.cw[gopdf.ToByte("p")]=557
	me.cw[gopdf.ToByte("q")]=557
	me.cw[gopdf.ToByte("r")]=340
	me.cw[gopdf.ToByte("s")]=509
	me.cw[gopdf.ToByte("t")]=317
	me.cw[gopdf.ToByte("u")]=557
	me.cw[gopdf.ToByte("v")]=489
	me.cw[gopdf.ToByte("w")]=757
	me.cw[gopdf.ToByte("x")]=489
	me.cw[gopdf.ToByte("y")]=489
	me.cw[gopdf.ToByte("z")]=489
	me.cw[gopdf.ToByte("{")]=330
	me.cw[gopdf.ToByte("|")]=221
	me.cw[gopdf.ToByte("}")]=330
	me.cw[gopdf.ToByte("~")]=685
	me.cw[gopdf.Chr(127)]=243
	me.cw[gopdf.Chr(128)]=524
	me.cw[gopdf.Chr(129)]=243
	me.cw[gopdf.Chr(130)]=172
	me.cw[gopdf.Chr(131)]=327
	me.cw[gopdf.Chr(132)]=290
	me.cw[gopdf.Chr(133)]=641
	me.cw[gopdf.Chr(134)]=550
	me.cw[gopdf.Chr(135)]=563
	me.cw[gopdf.Chr(136)]=430
	me.cw[gopdf.Chr(137)]=976
	me.cw[gopdf.Chr(138)]=601
	me.cw[gopdf.Chr(139)]=297
	me.cw[gopdf.Chr(140)]=934
	me.cw[gopdf.Chr(141)]=243
	me.cw[gopdf.Chr(142)]=598
	me.cw[gopdf.Chr(143)]=243
	me.cw[gopdf.Chr(144)]=243
	me.cw[gopdf.Chr(145)]=178
	me.cw[gopdf.Chr(146)]=178
	me.cw[gopdf.Chr(147)]=299
	me.cw[gopdf.Chr(148)]=301
	me.cw[gopdf.Chr(149)]=323
	me.cw[gopdf.Chr(150)]=690
	me.cw[gopdf.Chr(151)]=814
	me.cw[gopdf.Chr(152)]=436
	me.cw[gopdf.Chr(153)]=617
	me.cw[gopdf.Chr(154)]=509
	me.cw[gopdf.Chr(155)]=297
	me.cw[gopdf.Chr(156)]=924
	me.cw[gopdf.Chr(157)]=243
	me.cw[gopdf.Chr(158)]=489
	me.cw[gopdf.Chr(159)]=607
	me.cw[gopdf.Chr(160)]=243
	me.cw[gopdf.Chr(161)]=221
	me.cw[gopdf.Chr(162)]=544
	me.cw[gopdf.Chr(163)]=571
	me.cw[gopdf.Chr(164)]=723
	me.cw[gopdf.Chr(165)]=597
	me.cw[gopdf.Chr(166)]=217
	me.cw[gopdf.Chr(167)]=605
	me.cw[gopdf.Chr(168)]=468
	me.cw[gopdf.Chr(169)]=799
	me.cw[gopdf.Chr(170)]=442
	me.cw[gopdf.Chr(171)]=456
	me.cw[gopdf.Chr(172)]=545
	me.cw[gopdf.Chr(173)]=286
	me.cw[gopdf.Chr(174)]=802
	me.cw[gopdf.Chr(175)]=427
	me.cw[gopdf.Chr(176)]=378
	me.cw[gopdf.Chr(177)]=530
	me.cw[gopdf.Chr(178)]=408
	me.cw[gopdf.Chr(179)]=416
	me.cw[gopdf.Chr(180)]=281
	me.cw[gopdf.Chr(181)]=557
	me.cw[gopdf.Chr(182)]=475
	me.cw[gopdf.Chr(183)]=246
	me.cw[gopdf.Chr(184)]=243
	me.cw[gopdf.Chr(185)]=250
	me.cw[gopdf.Chr(186)]=450
	me.cw[gopdf.Chr(187)]=453
	me.cw[gopdf.Chr(188)]=762
	me.cw[gopdf.Chr(189)]=765
	me.cw[gopdf.Chr(190)]=836
	me.cw[gopdf.Chr(191)]=468
	me.cw[gopdf.Chr(192)]=629
	me.cw[gopdf.Chr(193)]=629
	me.cw[gopdf.Chr(194)]=629
	me.cw[gopdf.Chr(195)]=629
	me.cw[gopdf.Chr(196)]=629
	me.cw[gopdf.Chr(197)]=629
	me.cw[gopdf.Chr(198)]=911
	me.cw[gopdf.Chr(199)]=644
	me.cw[gopdf.Chr(200)]=576
	me.cw[gopdf.Chr(201)]=576
	me.cw[gopdf.Chr(202)]=576
	me.cw[gopdf.Chr(203)]=576
	me.cw[gopdf.Chr(204)]=271
	me.cw[gopdf.Chr(205)]=271
	me.cw[gopdf.Chr(206)]=271
	me.cw[gopdf.Chr(207)]=271
	me.cw[gopdf.Chr(208)]=680
	me.cw[gopdf.Chr(209)]=710
	me.cw[gopdf.Chr(210)]=681
	me.cw[gopdf.Chr(211)]=681
	me.cw[gopdf.Chr(212)]=681
	me.cw[gopdf.Chr(213)]=681
	me.cw[gopdf.Chr(214)]=681
	me.cw[gopdf.Chr(215)]=523
	me.cw[gopdf.Chr(216)]=673
	me.cw[gopdf.Chr(217)]=671
	me.cw[gopdf.Chr(218)]=671
	me.cw[gopdf.Chr(219)]=671
	me.cw[gopdf.Chr(220)]=671
	me.cw[gopdf.Chr(221)]=607
	me.cw[gopdf.Chr(222)]=593
	me.cw[gopdf.Chr(223)]=586
	me.cw[gopdf.Chr(224)]=538
	me.cw[gopdf.Chr(225)]=538
	me.cw[gopdf.Chr(226)]=538
	me.cw[gopdf.Chr(227)]=538
	me.cw[gopdf.Chr(228)]=538
	me.cw[gopdf.Chr(229)]=538
	me.cw[gopdf.Chr(230)]=846
	me.cw[gopdf.Chr(231)]=518
	me.cw[gopdf.Chr(232)]=515
	me.cw[gopdf.Chr(233)]=515
	me.cw[gopdf.Chr(234)]=515
	me.cw[gopdf.Chr(235)]=515
	me.cw[gopdf.Chr(236)]=223
	me.cw[gopdf.Chr(237)]=223
	me.cw[gopdf.Chr(238)]=223
	me.cw[gopdf.Chr(239)]=223
	me.cw[gopdf.Chr(240)]=582
	me.cw[gopdf.Chr(241)]=557
	me.cw[gopdf.Chr(242)]=557
	me.cw[gopdf.Chr(243)]=557
	me.cw[gopdf.Chr(244)]=557
	me.cw[gopdf.Chr(245)]=557
	me.cw[gopdf.Chr(246)]=557
	me.cw[gopdf.Chr(247)]=568
	me.cw[gopdf.Chr(248)]=557
	me.cw[gopdf.Chr(249)]=557
	me.cw[gopdf.Chr(250)]=557
	me.cw[gopdf.Chr(251)]=557
	me.cw[gopdf.Chr(252)]=557
	me.cw[gopdf.Chr(253)]=489
	me.cw[gopdf.Chr(254)]=562
	me.cw[gopdf.Chr(255)]=489
	me.up = -73
	me.ut = 49
	me.fonttype = "TrueType"
	me.name = "Roboto-Light"
	me.enc = "cp1252"
	me.desc = make([]gopdf.FontDescItem,8)
	me.desc[0] =  gopdf.FontDescItem{ Key:"Ascent",Val : "750" }
	me.desc[1] =  gopdf.FontDescItem{ Key: "Descent", Val : "-250" }
	me.desc[2] =  gopdf.FontDescItem{ Key:"CapHeight", Val :  "711" }
	me.desc[3] =  gopdf.FontDescItem{ Key: "Flags", Val :  "32" }
	me.desc[4] =  gopdf.FontDescItem{ Key:"FontBBox", Val :  "[-477 -271 1154 1048]" }
	me.desc[5] =  gopdf.FontDescItem{ Key:"ItalicAngle", Val :  "0" }
	me.desc[6] =  gopdf.FontDescItem{ Key:"StemV", Val :  "70" }
 	me.desc[7] =  gopdf.FontDescItem{ Key:"MissingWidth", Val :  "243" } 
 }
func (me * Roboto-Light)GetType() string{
	return me.fonttype
}
func (me * Roboto-Light)GetName() string{
	return me.name
}	
func (me * Roboto-Light)GetDesc() []gopdf.FontDescItem{
	return me.desc
}
func (me * Roboto-Light)GetUp() int{
	return me.up
}
func (me * Roboto-Light)GetUt()  int{
	return me.ut
}
func (me * Roboto-Light)GetCw() gopdf.FontCw{
	return me.cw
}
func (me * Roboto-Light)GetEnc() string{
	return me.enc
}
func (me * Roboto-Light)GetDiff() string {
	return me.diff
}
func (me * Roboto-Light) GetOriginalsize() int{
	return 98764
}
func (me * Roboto-Light)  SetFamily(family string){
	me.family = family
}
func (me * Roboto-Light) 	GetFamily() string{
	return me.family
}
