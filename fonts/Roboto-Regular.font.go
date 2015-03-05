package fonts //change this
import (
	"github.com/yaslama/gopdf"
)
type Roboto-Regular struct {
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
func (me * Roboto-Regular) Init(){
	me.cw = make(gopdf.FontCw)
	me.cw[gopdf.Chr(0)]=248
	me.cw[gopdf.Chr(1)]=248
	me.cw[gopdf.Chr(2)]=248
	me.cw[gopdf.Chr(3)]=248
	me.cw[gopdf.Chr(4)]=248
	me.cw[gopdf.Chr(5)]=248
	me.cw[gopdf.Chr(6)]=248
	me.cw[gopdf.Chr(7)]=248
	me.cw[gopdf.Chr(8)]=248
	me.cw[gopdf.Chr(9)]=248
	me.cw[gopdf.Chr(10)]=248
	me.cw[gopdf.Chr(11)]=248
	me.cw[gopdf.Chr(12)]=248
	me.cw[gopdf.Chr(13)]=248
	me.cw[gopdf.Chr(14)]=248
	me.cw[gopdf.Chr(15)]=248
	me.cw[gopdf.Chr(16)]=248
	me.cw[gopdf.Chr(17)]=248
	me.cw[gopdf.Chr(18)]=248
	me.cw[gopdf.Chr(19)]=248
	me.cw[gopdf.Chr(20)]=248
	me.cw[gopdf.Chr(21)]=248
	me.cw[gopdf.Chr(22)]=248
	me.cw[gopdf.Chr(23)]=248
	me.cw[gopdf.Chr(24)]=248
	me.cw[gopdf.Chr(25)]=248
	me.cw[gopdf.Chr(26)]=248
	me.cw[gopdf.Chr(27)]=248
	me.cw[gopdf.Chr(28)]=248
	me.cw[gopdf.Chr(29)]=248
	me.cw[gopdf.Chr(30)]=248
	me.cw[gopdf.Chr(31)]=248
	me.cw[gopdf.ToByte(" ")]=248
	me.cw[gopdf.ToByte("!")]=257
	me.cw[gopdf.ToByte("\"")]=320
	me.cw[gopdf.ToByte("#")]=623
	me.cw[gopdf.ToByte("$")]=562
	me.cw[gopdf.ToByte("%")]=732
	me.cw[gopdf.ToByte("&")]=622
	me.cw[gopdf.ToByte("'")]=174
	me.cw[gopdf.ToByte("(")]=329
	me.cw[gopdf.ToByte(")")]=333
	me.cw[gopdf.ToByte("*")]=431
	me.cw[gopdf.ToByte("+")]=567
	me.cw[gopdf.ToByte(",")]=196
	me.cw[gopdf.ToByte("-")]=276
	me.cw[gopdf.ToByte(".")]=263
	me.cw[gopdf.ToByte("/")]=412
	me.cw[gopdf.ToByte("0")]=562
	me.cw[gopdf.ToByte("1")]=562
	me.cw[gopdf.ToByte("2")]=562
	me.cw[gopdf.ToByte("3")]=562
	me.cw[gopdf.ToByte("4")]=562
	me.cw[gopdf.ToByte("5")]=562
	me.cw[gopdf.ToByte("6")]=562
	me.cw[gopdf.ToByte("7")]=562
	me.cw[gopdf.ToByte("8")]=562
	me.cw[gopdf.ToByte("9")]=562
	me.cw[gopdf.ToByte(":")]=246
	me.cw[gopdf.ToByte(";")]=250
	me.cw[gopdf.ToByte("<")]=508
	me.cw[gopdf.ToByte("=")]=563
	me.cw[gopdf.ToByte(">")]=522
	me.cw[gopdf.ToByte("?")]=472
	me.cw[gopdf.ToByte("@")]=898
	me.cw[gopdf.ToByte("A")]=661
	me.cw[gopdf.ToByte("B")]=636
	me.cw[gopdf.ToByte("C")]=640
	me.cw[gopdf.ToByte("D")]=676
	me.cw[gopdf.ToByte("E")]=583
	me.cw[gopdf.ToByte("F")]=582
	me.cw[gopdf.ToByte("G")]=687
	me.cw[gopdf.ToByte("H")]=712
	me.cw[gopdf.ToByte("I")]=281
	me.cw[gopdf.ToByte("J")]=552
	me.cw[gopdf.ToByte("K")]=643
	me.cw[gopdf.ToByte("L")]=547
	me.cw[gopdf.ToByte("M")]=876
	me.cw[gopdf.ToByte("N")]=713
	me.cw[gopdf.ToByte("O")]=695
	me.cw[gopdf.ToByte("P")]=637
	me.cw[gopdf.ToByte("Q")]=695
	me.cw[gopdf.ToByte("R")]=662
	me.cw[gopdf.ToByte("S")]=611
	me.cw[gopdf.ToByte("T")]=597
	me.cw[gopdf.ToByte("U")]=676
	me.cw[gopdf.ToByte("V")]=645
	me.cw[gopdf.ToByte("W")]=881
	me.cw[gopdf.ToByte("X")]=629
	me.cw[gopdf.ToByte("Y")]=618
	me.cw[gopdf.ToByte("Z")]=598
	me.cw[gopdf.ToByte("[")]=265
	me.cw[gopdf.ToByte("\\")]=410
	me.cw[gopdf.ToByte("]")]=265
	me.cw[gopdf.ToByte("^")]=418
	me.cw[gopdf.ToByte("_")]=451
	me.cw[gopdf.ToByte("`")]=309
	me.cw[gopdf.ToByte("a")]=548
	me.cw[gopdf.ToByte("b")]=566
	me.cw[gopdf.ToByte("c")]=529
	me.cw[gopdf.ToByte("d")]=566
	me.cw[gopdf.ToByte("e")]=527
	me.cw[gopdf.ToByte("f")]=343
	me.cw[gopdf.ToByte("g")]=566
	me.cw[gopdf.ToByte("h")]=566
	me.cw[gopdf.ToByte("i")]=248
	me.cw[gopdf.ToByte("j")]=255
	me.cw[gopdf.ToByte("k")]=509
	me.cw[gopdf.ToByte("l")]=248
	me.cw[gopdf.ToByte("m")]=876
	me.cw[gopdf.ToByte("n")]=566
	me.cw[gopdf.ToByte("o")]=566
	me.cw[gopdf.ToByte("p")]=566
	me.cw[gopdf.ToByte("q")]=566
	me.cw[gopdf.ToByte("r")]=349
	me.cw[gopdf.ToByte("s")]=521
	me.cw[gopdf.ToByte("t")]=318
	me.cw[gopdf.ToByte("u")]=566
	me.cw[gopdf.ToByte("v")]=501
	me.cw[gopdf.ToByte("w")]=757
	me.cw[gopdf.ToByte("x")]=501
	me.cw[gopdf.ToByte("y")]=501
	me.cw[gopdf.ToByte("z")]=501
	me.cw[gopdf.ToByte("{")]=338
	me.cw[gopdf.ToByte("|")]=244
	me.cw[gopdf.ToByte("}")]=338
	me.cw[gopdf.ToByte("~")]=680
	me.cw[gopdf.Chr(127)]=248
	me.cw[gopdf.Chr(128)]=530
	me.cw[gopdf.Chr(129)]=248
	me.cw[gopdf.Chr(130)]=199
	me.cw[gopdf.Chr(131)]=340
	me.cw[gopdf.Chr(132)]=344
	me.cw[gopdf.Chr(133)]=669
	me.cw[gopdf.Chr(134)]=551
	me.cw[gopdf.Chr(135)]=570
	me.cw[gopdf.Chr(136)]=471
	me.cw[gopdf.Chr(137)]=958
	me.cw[gopdf.Chr(138)]=611
	me.cw[gopdf.Chr(139)]=300
	me.cw[gopdf.Chr(140)]=954
	me.cw[gopdf.Chr(141)]=248
	me.cw[gopdf.Chr(142)]=598
	me.cw[gopdf.Chr(143)]=248
	me.cw[gopdf.Chr(144)]=248
	me.cw[gopdf.Chr(145)]=200
	me.cw[gopdf.Chr(146)]=200
	me.cw[gopdf.Chr(147)]=354
	me.cw[gopdf.Chr(148)]=357
	me.cw[gopdf.Chr(149)]=337
	me.cw[gopdf.Chr(150)]=691
	me.cw[gopdf.Chr(151)]=811
	me.cw[gopdf.Chr(152)]=472
	me.cw[gopdf.Chr(153)]=625
	me.cw[gopdf.Chr(154)]=521
	me.cw[gopdf.Chr(155)]=300
	me.cw[gopdf.Chr(156)]=908
	me.cw[gopdf.Chr(157)]=248
	me.cw[gopdf.Chr(158)]=501
	me.cw[gopdf.Chr(159)]=618
	me.cw[gopdf.Chr(160)]=248
	me.cw[gopdf.Chr(161)]=244
	me.cw[gopdf.Chr(162)]=547
	me.cw[gopdf.Chr(163)]=581
	me.cw[gopdf.Chr(164)]=713
	me.cw[gopdf.Chr(165)]=605
	me.cw[gopdf.Chr(166)]=240
	me.cw[gopdf.Chr(167)]=613
	me.cw[gopdf.Chr(168)]=494
	me.cw[gopdf.Chr(169)]=786
	me.cw[gopdf.Chr(170)]=447
	me.cw[gopdf.Chr(171)]=469
	me.cw[gopdf.Chr(172)]=554
	me.cw[gopdf.Chr(173)]=276
	me.cw[gopdf.Chr(174)]=786
	me.cw[gopdf.Chr(175)]=458
	me.cw[gopdf.Chr(176)]=374
	me.cw[gopdf.Chr(177)]=535
	me.cw[gopdf.Chr(178)]=421
	me.cw[gopdf.Chr(179)]=426
	me.cw[gopdf.Chr(180)]=313
	me.cw[gopdf.Chr(181)]=566
	me.cw[gopdf.Chr(182)]=489
	me.cw[gopdf.Chr(183)]=261
	me.cw[gopdf.Chr(184)]=248
	me.cw[gopdf.Chr(185)]=269
	me.cw[gopdf.Chr(186)]=455
	me.cw[gopdf.Chr(187)]=469
	me.cw[gopdf.Chr(188)]=776
	me.cw[gopdf.Chr(189)]=823
	me.cw[gopdf.Chr(190)]=865
	me.cw[gopdf.Chr(191)]=491
	me.cw[gopdf.Chr(192)]=661
	me.cw[gopdf.Chr(193)]=661
	me.cw[gopdf.Chr(194)]=661
	me.cw[gopdf.Chr(195)]=661
	me.cw[gopdf.Chr(196)]=661
	me.cw[gopdf.Chr(197)]=661
	me.cw[gopdf.Chr(198)]=935
	me.cw[gopdf.Chr(199)]=640
	me.cw[gopdf.Chr(200)]=583
	me.cw[gopdf.Chr(201)]=583
	me.cw[gopdf.Chr(202)]=583
	me.cw[gopdf.Chr(203)]=583
	me.cw[gopdf.Chr(204)]=281
	me.cw[gopdf.Chr(205)]=281
	me.cw[gopdf.Chr(206)]=281
	me.cw[gopdf.Chr(207)]=281
	me.cw[gopdf.Chr(208)]=691
	me.cw[gopdf.Chr(209)]=713
	me.cw[gopdf.Chr(210)]=695
	me.cw[gopdf.Chr(211)]=695
	me.cw[gopdf.Chr(212)]=695
	me.cw[gopdf.Chr(213)]=695
	me.cw[gopdf.Chr(214)]=695
	me.cw[gopdf.Chr(215)]=533
	me.cw[gopdf.Chr(216)]=681
	me.cw[gopdf.Chr(217)]=676
	me.cw[gopdf.Chr(218)]=676
	me.cw[gopdf.Chr(219)]=676
	me.cw[gopdf.Chr(220)]=676
	me.cw[gopdf.Chr(221)]=618
	me.cw[gopdf.Chr(222)]=591
	me.cw[gopdf.Chr(223)]=595
	me.cw[gopdf.Chr(224)]=548
	me.cw[gopdf.Chr(225)]=548
	me.cw[gopdf.Chr(226)]=548
	me.cw[gopdf.Chr(227)]=548
	me.cw[gopdf.Chr(228)]=548
	me.cw[gopdf.Chr(229)]=548
	me.cw[gopdf.Chr(230)]=844
	me.cw[gopdf.Chr(231)]=529
	me.cw[gopdf.Chr(232)]=527
	me.cw[gopdf.Chr(233)]=527
	me.cw[gopdf.Chr(234)]=527
	me.cw[gopdf.Chr(235)]=527
	me.cw[gopdf.Chr(236)]=247
	me.cw[gopdf.Chr(237)]=247
	me.cw[gopdf.Chr(238)]=247
	me.cw[gopdf.Chr(239)]=247
	me.cw[gopdf.Chr(240)]=586
	me.cw[gopdf.Chr(241)]=566
	me.cw[gopdf.Chr(242)]=566
	me.cw[gopdf.Chr(243)]=566
	me.cw[gopdf.Chr(244)]=566
	me.cw[gopdf.Chr(245)]=566
	me.cw[gopdf.Chr(246)]=566
	me.cw[gopdf.Chr(247)]=571
	me.cw[gopdf.Chr(248)]=567
	me.cw[gopdf.Chr(249)]=566
	me.cw[gopdf.Chr(250)]=566
	me.cw[gopdf.Chr(251)]=566
	me.cw[gopdf.Chr(252)]=566
	me.cw[gopdf.Chr(253)]=501
	me.cw[gopdf.Chr(254)]=576
	me.cw[gopdf.Chr(255)]=501
	me.up = -73
	me.ut = 49
	me.fonttype = "TrueType"
	me.name = "Roboto-Regular"
	me.enc = "cp1252"
	me.desc = make([]gopdf.FontDescItem,8)
	me.desc[0] =  gopdf.FontDescItem{ Key:"Ascent",Val : "750" }
	me.desc[1] =  gopdf.FontDescItem{ Key: "Descent", Val : "-250" }
	me.desc[2] =  gopdf.FontDescItem{ Key:"CapHeight", Val :  "711" }
	me.desc[3] =  gopdf.FontDescItem{ Key: "Flags", Val :  "32" }
	me.desc[4] =  gopdf.FontDescItem{ Key:"FontBBox", Val :  "[-681 -271 1182 1048]" }
	me.desc[5] =  gopdf.FontDescItem{ Key:"ItalicAngle", Val :  "0" }
	me.desc[6] =  gopdf.FontDescItem{ Key:"StemV", Val :  "70" }
 	me.desc[7] =  gopdf.FontDescItem{ Key:"MissingWidth", Val :  "248" } 
 }
func (me * Roboto-Regular)GetType() string{
	return me.fonttype
}
func (me * Roboto-Regular)GetName() string{
	return me.name
}	
func (me * Roboto-Regular)GetDesc() []gopdf.FontDescItem{
	return me.desc
}
func (me * Roboto-Regular)GetUp() int{
	return me.up
}
func (me * Roboto-Regular)GetUt()  int{
	return me.ut
}
func (me * Roboto-Regular)GetCw() gopdf.FontCw{
	return me.cw
}
func (me * Roboto-Regular)GetEnc() string{
	return me.enc
}
func (me * Roboto-Regular)GetDiff() string {
	return me.diff
}
func (me * Roboto-Regular) GetOriginalsize() int{
	return 98764
}
func (me * Roboto-Regular)  SetFamily(family string){
	me.family = family
}
func (me * Roboto-Regular) 	GetFamily() string{
	return me.family
}
