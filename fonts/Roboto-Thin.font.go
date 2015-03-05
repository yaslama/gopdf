package fonts //change this
import (
	"github.com/signintech/gopdf"
)
type Roboto-Thin struct {
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
func (me * Roboto-Thin) Init(){
	me.cw = make(gopdf.FontCw)
	me.cw[gopdf.Chr(0)]=239
	me.cw[gopdf.Chr(1)]=239
	me.cw[gopdf.Chr(2)]=239
	me.cw[gopdf.Chr(3)]=239
	me.cw[gopdf.Chr(4)]=239
	me.cw[gopdf.Chr(5)]=239
	me.cw[gopdf.Chr(6)]=239
	me.cw[gopdf.Chr(7)]=239
	me.cw[gopdf.Chr(8)]=239
	me.cw[gopdf.Chr(9)]=239
	me.cw[gopdf.Chr(10)]=239
	me.cw[gopdf.Chr(11)]=239
	me.cw[gopdf.Chr(12)]=239
	me.cw[gopdf.Chr(13)]=239
	me.cw[gopdf.Chr(14)]=239
	me.cw[gopdf.Chr(15)]=239
	me.cw[gopdf.Chr(16)]=239
	me.cw[gopdf.Chr(17)]=239
	me.cw[gopdf.Chr(18)]=239
	me.cw[gopdf.Chr(19)]=239
	me.cw[gopdf.Chr(20)]=239
	me.cw[gopdf.Chr(21)]=239
	me.cw[gopdf.Chr(22)]=239
	me.cw[gopdf.Chr(23)]=239
	me.cw[gopdf.Chr(24)]=239
	me.cw[gopdf.Chr(25)]=239
	me.cw[gopdf.Chr(26)]=239
	me.cw[gopdf.Chr(27)]=239
	me.cw[gopdf.Chr(28)]=239
	me.cw[gopdf.Chr(29)]=239
	me.cw[gopdf.Chr(30)]=239
	me.cw[gopdf.Chr(31)]=239
	me.cw[gopdf.ToByte(" ")]=239
	me.cw[gopdf.ToByte("!")]=194
	me.cw[gopdf.ToByte("\"")]=254
	me.cw[gopdf.ToByte("#")]=618
	me.cw[gopdf.ToByte("$")]=547
	me.cw[gopdf.ToByte("%")]=746
	me.cw[gopdf.ToByte("&")]=609
	me.cw[gopdf.ToByte("'")]=166
	me.cw[gopdf.ToByte("(")]=296
	me.cw[gopdf.ToByte(")")]=304
	me.cw[gopdf.ToByte("*")]=418
	me.cw[gopdf.ToByte("+")]=562
	me.cw[gopdf.ToByte(",")]=187
	me.cw[gopdf.ToByte("-")]=296
	me.cw[gopdf.ToByte(".")]=214
	me.cw[gopdf.ToByte("/")]=382
	me.cw[gopdf.ToByte("0")]=547
	me.cw[gopdf.ToByte("1")]=547
	me.cw[gopdf.ToByte("2")]=547
	me.cw[gopdf.ToByte("3")]=547
	me.cw[gopdf.ToByte("4")]=547
	me.cw[gopdf.ToByte("5")]=547
	me.cw[gopdf.ToByte("6")]=547
	me.cw[gopdf.ToByte("7")]=547
	me.cw[gopdf.ToByte("8")]=547
	me.cw[gopdf.ToByte("9")]=547
	me.cw[gopdf.ToByte(":")]=178
	me.cw[gopdf.ToByte(";")]=178
	me.cw[gopdf.ToByte("<")]=514
	me.cw[gopdf.ToByte("=")]=558
	me.cw[gopdf.ToByte(">")]=514
	me.cw[gopdf.ToByte("?")]=436
	me.cw[gopdf.ToByte("@")]=929
	me.cw[gopdf.ToByte("A")]=597
	me.cw[gopdf.ToByte("B")]=603
	me.cw[gopdf.ToByte("C")]=648
	me.cw[gopdf.ToByte("D")]=654
	me.cw[gopdf.ToByte("E")]=569
	me.cw[gopdf.ToByte("F")]=572
	me.cw[gopdf.ToByte("G")]=687
	me.cw[gopdf.ToByte("H")]=702
	me.cw[gopdf.ToByte("I")]=261
	me.cw[gopdf.ToByte("J")]=549
	me.cw[gopdf.ToByte("K")]=635
	me.cw[gopdf.ToByte("L")]=516
	me.cw[gopdf.ToByte("M")]=857
	me.cw[gopdf.ToByte("N")]=708
	me.cw[gopdf.ToByte("O")]=666
	me.cw[gopdf.ToByte("P")]=601
	me.cw[gopdf.ToByte("Q")]=666
	me.cw[gopdf.ToByte("R")]=654
	me.cw[gopdf.ToByte("S")]=591
	me.cw[gopdf.ToByte("T")]=598
	me.cw[gopdf.ToByte("U")]=667
	me.cw[gopdf.ToByte("V")]=597
	me.cw[gopdf.ToByte("W")]=906
	me.cw[gopdf.ToByte("X")]=597
	me.cw[gopdf.ToByte("Y")]=597
	me.cw[gopdf.ToByte("Z")]=598
	me.cw[gopdf.ToByte("[")]=214
	me.cw[gopdf.ToByte("\\")]=378
	me.cw[gopdf.ToByte("]")]=214
	me.cw[gopdf.ToByte("^")]=415
	me.cw[gopdf.ToByte("_")]=413
	me.cw[gopdf.ToByte("`")]=262
	me.cw[gopdf.ToByte("a")]=528
	me.cw[gopdf.ToByte("b")]=547
	me.cw[gopdf.ToByte("c")]=507
	me.cw[gopdf.ToByte("d")]=547
	me.cw[gopdf.ToByte("e")]=504
	me.cw[gopdf.ToByte("f")]=315
	me.cw[gopdf.ToByte("g")]=548
	me.cw[gopdf.ToByte("h")]=547
	me.cw[gopdf.ToByte("i")]=206
	me.cw[gopdf.ToByte("j")]=217
	me.cw[gopdf.ToByte("k")]=473
	me.cw[gopdf.ToByte("l")]=206
	me.cw[gopdf.ToByte("m")]=896
	me.cw[gopdf.ToByte("n")]=547
	me.cw[gopdf.ToByte("o")]=547
	me.cw[gopdf.ToByte("p")]=547
	me.cw[gopdf.ToByte("q")]=547
	me.cw[gopdf.ToByte("r")]=332
	me.cw[gopdf.ToByte("s")]=498
	me.cw[gopdf.ToByte("t")]=316
	me.cw[gopdf.ToByte("u")]=547
	me.cw[gopdf.ToByte("v")]=478
	me.cw[gopdf.ToByte("w")]=757
	me.cw[gopdf.ToByte("x")]=478
	me.cw[gopdf.ToByte("y")]=478
	me.cw[gopdf.ToByte("z")]=478
	me.cw[gopdf.ToByte("{")]=322
	me.cw[gopdf.ToByte("|")]=198
	me.cw[gopdf.ToByte("}")]=322
	me.cw[gopdf.ToByte("~")]=689
	me.cw[gopdf.Chr(127)]=239
	me.cw[gopdf.Chr(128)]=518
	me.cw[gopdf.Chr(129)]=239
	me.cw[gopdf.Chr(130)]=146
	me.cw[gopdf.Chr(131)]=314
	me.cw[gopdf.Chr(132)]=235
	me.cw[gopdf.Chr(133)]=613
	me.cw[gopdf.Chr(134)]=550
	me.cw[gopdf.Chr(135)]=557
	me.cw[gopdf.Chr(136)]=389
	me.cw[gopdf.Chr(137)]=993
	me.cw[gopdf.Chr(138)]=591
	me.cw[gopdf.Chr(139)]=295
	me.cw[gopdf.Chr(140)]=915
	me.cw[gopdf.Chr(141)]=239
	me.cw[gopdf.Chr(142)]=598
	me.cw[gopdf.Chr(143)]=239
	me.cw[gopdf.Chr(144)]=239
	me.cw[gopdf.Chr(145)]=156
	me.cw[gopdf.Chr(146)]=156
	me.cw[gopdf.Chr(147)]=244
	me.cw[gopdf.Chr(148)]=245
	me.cw[gopdf.Chr(149)]=310
	me.cw[gopdf.Chr(150)]=689
	me.cw[gopdf.Chr(151)]=817
	me.cw[gopdf.Chr(152)]=399
	me.cw[gopdf.Chr(153)]=609
	me.cw[gopdf.Chr(154)]=498
	me.cw[gopdf.Chr(155)]=295
	me.cw[gopdf.Chr(156)]=939
	me.cw[gopdf.Chr(157)]=239
	me.cw[gopdf.Chr(158)]=478
	me.cw[gopdf.Chr(159)]=597
	me.cw[gopdf.Chr(160)]=239
	me.cw[gopdf.Chr(161)]=198
	me.cw[gopdf.Chr(162)]=542
	me.cw[gopdf.Chr(163)]=562
	me.cw[gopdf.Chr(164)]=733
	me.cw[gopdf.Chr(165)]=589
	me.cw[gopdf.Chr(166)]=194
	me.cw[gopdf.Chr(167)]=597
	me.cw[gopdf.Chr(168)]=442
	me.cw[gopdf.Chr(169)]=813
	me.cw[gopdf.Chr(170)]=438
	me.cw[gopdf.Chr(171)]=442
	me.cw[gopdf.Chr(172)]=538
	me.cw[gopdf.Chr(173)]=296
	me.cw[gopdf.Chr(174)]=817
	me.cw[gopdf.Chr(175)]=396
	me.cw[gopdf.Chr(176)]=382
	me.cw[gopdf.Chr(177)]=526
	me.cw[gopdf.Chr(178)]=394
	me.cw[gopdf.Chr(179)]=406
	me.cw[gopdf.Chr(180)]=250
	me.cw[gopdf.Chr(181)]=547
	me.cw[gopdf.Chr(182)]=462
	me.cw[gopdf.Chr(183)]=230
	me.cw[gopdf.Chr(184)]=239
	me.cw[gopdf.Chr(185)]=230
	me.cw[gopdf.Chr(186)]=446
	me.cw[gopdf.Chr(187)]=438
	me.cw[gopdf.Chr(188)]=749
	me.cw[gopdf.Chr(189)]=707
	me.cw[gopdf.Chr(190)]=807
	me.cw[gopdf.Chr(191)]=445
	me.cw[gopdf.Chr(192)]=597
	me.cw[gopdf.Chr(193)]=597
	me.cw[gopdf.Chr(194)]=597
	me.cw[gopdf.Chr(195)]=597
	me.cw[gopdf.Chr(196)]=597
	me.cw[gopdf.Chr(197)]=597
	me.cw[gopdf.Chr(198)]=887
	me.cw[gopdf.Chr(199)]=648
	me.cw[gopdf.Chr(200)]=569
	me.cw[gopdf.Chr(201)]=569
	me.cw[gopdf.Chr(202)]=569
	me.cw[gopdf.Chr(203)]=569
	me.cw[gopdf.Chr(204)]=261
	me.cw[gopdf.Chr(205)]=261
	me.cw[gopdf.Chr(206)]=261
	me.cw[gopdf.Chr(207)]=261
	me.cw[gopdf.Chr(208)]=669
	me.cw[gopdf.Chr(209)]=708
	me.cw[gopdf.Chr(210)]=666
	me.cw[gopdf.Chr(211)]=666
	me.cw[gopdf.Chr(212)]=666
	me.cw[gopdf.Chr(213)]=666
	me.cw[gopdf.Chr(214)]=666
	me.cw[gopdf.Chr(215)]=514
	me.cw[gopdf.Chr(216)]=666
	me.cw[gopdf.Chr(217)]=667
	me.cw[gopdf.Chr(218)]=667
	me.cw[gopdf.Chr(219)]=667
	me.cw[gopdf.Chr(220)]=667
	me.cw[gopdf.Chr(221)]=597
	me.cw[gopdf.Chr(222)]=595
	me.cw[gopdf.Chr(223)]=578
	me.cw[gopdf.Chr(224)]=528
	me.cw[gopdf.Chr(225)]=528
	me.cw[gopdf.Chr(226)]=528
	me.cw[gopdf.Chr(227)]=528
	me.cw[gopdf.Chr(228)]=528
	me.cw[gopdf.Chr(229)]=528
	me.cw[gopdf.Chr(230)]=847
	me.cw[gopdf.Chr(231)]=507
	me.cw[gopdf.Chr(232)]=504
	me.cw[gopdf.Chr(233)]=504
	me.cw[gopdf.Chr(234)]=504
	me.cw[gopdf.Chr(235)]=504
	me.cw[gopdf.Chr(236)]=198
	me.cw[gopdf.Chr(237)]=198
	me.cw[gopdf.Chr(238)]=198
	me.cw[gopdf.Chr(239)]=198
	me.cw[gopdf.Chr(240)]=578
	me.cw[gopdf.Chr(241)]=547
	me.cw[gopdf.Chr(242)]=547
	me.cw[gopdf.Chr(243)]=547
	me.cw[gopdf.Chr(244)]=547
	me.cw[gopdf.Chr(245)]=547
	me.cw[gopdf.Chr(246)]=547
	me.cw[gopdf.Chr(247)]=566
	me.cw[gopdf.Chr(248)]=547
	me.cw[gopdf.Chr(249)]=547
	me.cw[gopdf.Chr(250)]=547
	me.cw[gopdf.Chr(251)]=547
	me.cw[gopdf.Chr(252)]=547
	me.cw[gopdf.Chr(253)]=478
	me.cw[gopdf.Chr(254)]=547
	me.cw[gopdf.Chr(255)]=478
	me.up = -73
	me.ut = 49
	me.fonttype = "TrueType"
	me.name = "Roboto-Thin"
	me.enc = "cp1252"
	me.desc = make([]gopdf.FontDescItem,8)
	me.desc[0] =  gopdf.FontDescItem{ Key:"Ascent",Val : "750" }
	me.desc[1] =  gopdf.FontDescItem{ Key: "Descent", Val : "-250" }
	me.desc[2] =  gopdf.FontDescItem{ Key:"CapHeight", Val :  "711" }
	me.desc[3] =  gopdf.FontDescItem{ Key: "Flags", Val :  "32" }
	me.desc[4] =  gopdf.FontDescItem{ Key:"FontBBox", Val :  "[-473 -271 1127 1048]" }
	me.desc[5] =  gopdf.FontDescItem{ Key:"ItalicAngle", Val :  "0" }
	me.desc[6] =  gopdf.FontDescItem{ Key:"StemV", Val :  "70" }
 	me.desc[7] =  gopdf.FontDescItem{ Key:"MissingWidth", Val :  "239" } 
 }
func (me * Roboto-Thin)GetType() string{
	return me.fonttype
}
func (me * Roboto-Thin)GetName() string{
	return me.name
}	
func (me * Roboto-Thin)GetDesc() []gopdf.FontDescItem{
	return me.desc
}
func (me * Roboto-Thin)GetUp() int{
	return me.up
}
func (me * Roboto-Thin)GetUt()  int{
	return me.ut
}
func (me * Roboto-Thin)GetCw() gopdf.FontCw{
	return me.cw
}
func (me * Roboto-Thin)GetEnc() string{
	return me.enc
}
func (me * Roboto-Thin)GetDiff() string {
	return me.diff
}
func (me * Roboto-Thin) GetOriginalsize() int{
	return 98764
}
func (me * Roboto-Thin)  SetFamily(family string){
	me.family = family
}
func (me * Roboto-Thin) 	GetFamily() string{
	return me.family
}
