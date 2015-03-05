package fonts //change this
import (
	"github.com/signintech/gopdf"
)
type Roboto-MediumItalic struct {
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
func (me * Roboto-MediumItalic) Init(){
	me.cw = make(gopdf.FontCw)
	me.cw[gopdf.Chr(0)]=237
	me.cw[gopdf.Chr(1)]=237
	me.cw[gopdf.Chr(2)]=237
	me.cw[gopdf.Chr(3)]=237
	me.cw[gopdf.Chr(4)]=237
	me.cw[gopdf.Chr(5)]=237
	me.cw[gopdf.Chr(6)]=237
	me.cw[gopdf.Chr(7)]=237
	me.cw[gopdf.Chr(8)]=237
	me.cw[gopdf.Chr(9)]=237
	me.cw[gopdf.Chr(10)]=237
	me.cw[gopdf.Chr(11)]=237
	me.cw[gopdf.Chr(12)]=237
	me.cw[gopdf.Chr(13)]=237
	me.cw[gopdf.Chr(14)]=237
	me.cw[gopdf.Chr(15)]=237
	me.cw[gopdf.Chr(16)]=237
	me.cw[gopdf.Chr(17)]=237
	me.cw[gopdf.Chr(18)]=237
	me.cw[gopdf.Chr(19)]=237
	me.cw[gopdf.Chr(20)]=237
	me.cw[gopdf.Chr(21)]=237
	me.cw[gopdf.Chr(22)]=237
	me.cw[gopdf.Chr(23)]=237
	me.cw[gopdf.Chr(24)]=237
	me.cw[gopdf.Chr(25)]=237
	me.cw[gopdf.Chr(26)]=237
	me.cw[gopdf.Chr(27)]=237
	me.cw[gopdf.Chr(28)]=237
	me.cw[gopdf.Chr(29)]=237
	me.cw[gopdf.Chr(30)]=237
	me.cw[gopdf.Chr(31)]=237
	me.cw[gopdf.ToByte(" ")]=237
	me.cw[gopdf.ToByte("!")]=259
	me.cw[gopdf.ToByte("\"")]=312
	me.cw[gopdf.ToByte("#")]=572
	me.cw[gopdf.ToByte("$")]=535
	me.cw[gopdf.ToByte("%")]=687
	me.cw[gopdf.ToByte("&")]=602
	me.cw[gopdf.ToByte("'")]=162
	me.cw[gopdf.ToByte("(")]=324
	me.cw[gopdf.ToByte(")")]=325
	me.cw[gopdf.ToByte("*")]=413
	me.cw[gopdf.ToByte("+")]=522
	me.cw[gopdf.ToByte(",")]=211
	me.cw[gopdf.ToByte("-")]=313
	me.cw[gopdf.ToByte(".")]=269
	me.cw[gopdf.ToByte("/")]=374
	me.cw[gopdf.ToByte("0")]=535
	me.cw[gopdf.ToByte("1")]=535
	me.cw[gopdf.ToByte("2")]=535
	me.cw[gopdf.ToByte("3")]=535
	me.cw[gopdf.ToByte("4")]=535
	me.cw[gopdf.ToByte("5")]=535
	me.cw[gopdf.ToByte("6")]=535
	me.cw[gopdf.ToByte("7")]=535
	me.cw[gopdf.ToByte("8")]=535
	me.cw[gopdf.ToByte("9")]=535
	me.cw[gopdf.ToByte(":")]=261
	me.cw[gopdf.ToByte(";")]=258
	me.cw[gopdf.ToByte("<")]=477
	me.cw[gopdf.ToByte("=")]=535
	me.cw[gopdf.ToByte(">")]=489
	me.cw[gopdf.ToByte("?")]=460
	me.cw[gopdf.ToByte("@")]=831
	me.cw[gopdf.ToByte("A")]=620
	me.cw[gopdf.ToByte("B")]=604
	me.cw[gopdf.ToByte("C")]=595
	me.cw[gopdf.ToByte("D")]=616
	me.cw[gopdf.ToByte("E")]=544
	me.cw[gopdf.ToByte("F")]=543
	me.cw[gopdf.ToByte("G")]=635
	me.cw[gopdf.ToByte("H")]=667
	me.cw[gopdf.ToByte("I")]=274
	me.cw[gopdf.ToByte("J")]=527
	me.cw[gopdf.ToByte("K")]=570
	me.cw[gopdf.ToByte("L")]=518
	me.cw[gopdf.ToByte("M")]=821
	me.cw[gopdf.ToByte("N")]=667
	me.cw[gopdf.ToByte("O")]=641
	me.cw[gopdf.ToByte("P")]=609
	me.cw[gopdf.ToByte("Q")]=654
	me.cw[gopdf.ToByte("R")]=618
	me.cw[gopdf.ToByte("S")]=574
	me.cw[gopdf.ToByte("T")]=577
	me.cw[gopdf.ToByte("U")]=640
	me.cw[gopdf.ToByte("V")]=609
	me.cw[gopdf.ToByte("W")]=817
	me.cw[gopdf.ToByte("X")]=598
	me.cw[gopdf.ToByte("Y")]=589
	me.cw[gopdf.ToByte("Z")]=529
	me.cw[gopdf.ToByte("[")]=264
	me.cw[gopdf.ToByte("\\")]=396
	me.cw[gopdf.ToByte("]")]=264
	me.cw[gopdf.ToByte("^")]=403
	me.cw[gopdf.ToByte("_")]=427
	me.cw[gopdf.ToByte("`")]=309
	me.cw[gopdf.ToByte("a")]=512
	me.cw[gopdf.ToByte("b")]=533
	me.cw[gopdf.ToByte("c")]=494
	me.cw[gopdf.ToByte("d")]=533
	me.cw[gopdf.ToByte("e")]=499
	me.cw[gopdf.ToByte("f")]=334
	me.cw[gopdf.ToByte("g")]=533
	me.cw[gopdf.ToByte("h")]=533
	me.cw[gopdf.ToByte("i")]=250
	me.cw[gopdf.ToByte("j")]=252
	me.cw[gopdf.ToByte("k")]=495
	me.cw[gopdf.ToByte("l")]=250
	me.cw[gopdf.ToByte("m")]=812
	me.cw[gopdf.ToByte("n")]=533
	me.cw[gopdf.ToByte("o")]=533
	me.cw[gopdf.ToByte("p")]=533
	me.cw[gopdf.ToByte("q")]=533
	me.cw[gopdf.ToByte("r")]=338
	me.cw[gopdf.ToByte("s")]=490
	me.cw[gopdf.ToByte("t")]=310
	me.cw[gopdf.ToByte("u")]=533
	me.cw[gopdf.ToByte("v")]=479
	me.cw[gopdf.ToByte("w")]=699
	me.cw[gopdf.ToByte("x")]=479
	me.cw[gopdf.ToByte("y")]=479
	me.cw[gopdf.ToByte("z")]=479
	me.cw[gopdf.ToByte("{")]=318
	me.cw[gopdf.ToByte("|")]=242
	me.cw[gopdf.ToByte("}")]=318
	me.cw[gopdf.ToByte("~")]=621
	me.cw[gopdf.Chr(127)]=237
	me.cw[gopdf.Chr(128)]=506
	me.cw[gopdf.Chr(129)]=237
	me.cw[gopdf.Chr(130)]=220
	me.cw[gopdf.Chr(131)]=334
	me.cw[gopdf.Chr(132)]=367
	me.cw[gopdf.Chr(133)]=668
	me.cw[gopdf.Chr(134)]=512
	me.cw[gopdf.Chr(135)]=541
	me.cw[gopdf.Chr(136)]=465
	me.cw[gopdf.Chr(137)]=893
	me.cw[gopdf.Chr(138)]=582
	me.cw[gopdf.Chr(139)]=290
	me.cw[gopdf.Chr(140)]=904
	me.cw[gopdf.Chr(141)]=237
	me.cw[gopdf.Chr(142)]=557
	me.cw[gopdf.Chr(143)]=237
	me.cw[gopdf.Chr(144)]=237
	me.cw[gopdf.Chr(145)]=213
	me.cw[gopdf.Chr(146)]=211
	me.cw[gopdf.Chr(147)]=373
	me.cw[gopdf.Chr(148)]=376
	me.cw[gopdf.Chr(149)]=333
	me.cw[gopdf.Chr(150)]=647
	me.cw[gopdf.Chr(151)]=761
	me.cw[gopdf.Chr(152)]=458
	me.cw[gopdf.Chr(153)]=592
	me.cw[gopdf.Chr(154)]=490
	me.cw[gopdf.Chr(155)]=286
	me.cw[gopdf.Chr(156)]=842
	me.cw[gopdf.Chr(157)]=237
	me.cw[gopdf.Chr(158)]=479
	me.cw[gopdf.Chr(159)]=589
	me.cw[gopdf.Chr(160)]=237
	me.cw[gopdf.Chr(161)]=256
	me.cw[gopdf.Chr(162)]=528
	me.cw[gopdf.Chr(163)]=554
	me.cw[gopdf.Chr(164)]=655
	me.cw[gopdf.Chr(165)]=573
	me.cw[gopdf.Chr(166)]=239
	me.cw[gopdf.Chr(167)]=584
	me.cw[gopdf.Chr(168)]=489
	me.cw[gopdf.Chr(169)]=731
	me.cw[gopdf.Chr(170)]=420
	me.cw[gopdf.Chr(171)]=459
	me.cw[gopdf.Chr(172)]=521
	me.cw[gopdf.Chr(173)]=313
	me.cw[gopdf.Chr(174)]=731
	me.cw[gopdf.Chr(175)]=459
	me.cw[gopdf.Chr(176)]=358
	me.cw[gopdf.Chr(177)]=505
	me.cw[gopdf.Chr(178)]=400
	me.cw[gopdf.Chr(179)]=400
	me.cw[gopdf.Chr(180)]=324
	me.cw[gopdf.Chr(181)]=558
	me.cw[gopdf.Chr(182)]=480
	me.cw[gopdf.Chr(183)]=271
	me.cw[gopdf.Chr(184)]=246
	me.cw[gopdf.Chr(185)]=265
	me.cw[gopdf.Chr(186)]=430
	me.cw[gopdf.Chr(187)]=459
	me.cw[gopdf.Chr(188)]=708
	me.cw[gopdf.Chr(189)]=755
	me.cw[gopdf.Chr(190)]=801
	me.cw[gopdf.Chr(191)]=474
	me.cw[gopdf.Chr(192)]=620
	me.cw[gopdf.Chr(193)]=620
	me.cw[gopdf.Chr(194)]=620
	me.cw[gopdf.Chr(195)]=620
	me.cw[gopdf.Chr(196)]=620
	me.cw[gopdf.Chr(197)]=620
	me.cw[gopdf.Chr(198)]=883
	me.cw[gopdf.Chr(199)]=600
	me.cw[gopdf.Chr(200)]=544
	me.cw[gopdf.Chr(201)]=544
	me.cw[gopdf.Chr(202)]=544
	me.cw[gopdf.Chr(203)]=544
	me.cw[gopdf.Chr(204)]=274
	me.cw[gopdf.Chr(205)]=274
	me.cw[gopdf.Chr(206)]=274
	me.cw[gopdf.Chr(207)]=274
	me.cw[gopdf.Chr(208)]=646
	me.cw[gopdf.Chr(209)]=667
	me.cw[gopdf.Chr(210)]=653
	me.cw[gopdf.Chr(211)]=653
	me.cw[gopdf.Chr(212)]=653
	me.cw[gopdf.Chr(213)]=653
	me.cw[gopdf.Chr(214)]=653
	me.cw[gopdf.Chr(215)]=502
	me.cw[gopdf.Chr(216)]=643
	me.cw[gopdf.Chr(217)]=640
	me.cw[gopdf.Chr(218)]=640
	me.cw[gopdf.Chr(219)]=640
	me.cw[gopdf.Chr(220)]=640
	me.cw[gopdf.Chr(221)]=589
	me.cw[gopdf.Chr(222)]=562
	me.cw[gopdf.Chr(223)]=577
	me.cw[gopdf.Chr(224)]=512
	me.cw[gopdf.Chr(225)]=512
	me.cw[gopdf.Chr(226)]=512
	me.cw[gopdf.Chr(227)]=512
	me.cw[gopdf.Chr(228)]=512
	me.cw[gopdf.Chr(229)]=512
	me.cw[gopdf.Chr(230)]=790
	me.cw[gopdf.Chr(231)]=494
	me.cw[gopdf.Chr(232)]=499
	me.cw[gopdf.Chr(233)]=499
	me.cw[gopdf.Chr(234)]=499
	me.cw[gopdf.Chr(235)]=499
	me.cw[gopdf.Chr(236)]=254
	me.cw[gopdf.Chr(237)]=254
	me.cw[gopdf.Chr(238)]=254
	me.cw[gopdf.Chr(239)]=254
	me.cw[gopdf.Chr(240)]=555
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
	me.cw[gopdf.Chr(253)]=479
	me.cw[gopdf.Chr(254)]=541
	me.cw[gopdf.Chr(255)]=479
	me.up = -73
	me.ut = 49
	me.fonttype = "TrueType"
	me.name = "Roboto-MediumItalic"
	me.enc = "cp1252"
	me.desc = make([]gopdf.FontDescItem,8)
	me.desc[0] =  gopdf.FontDescItem{ Key:"Ascent",Val : "750" }
	me.desc[1] =  gopdf.FontDescItem{ Key: "Descent", Val : "-250" }
	me.desc[2] =  gopdf.FontDescItem{ Key:"CapHeight", Val :  "711" }
	me.desc[3] =  gopdf.FontDescItem{ Key: "Flags", Val :  "32" }
	me.desc[4] =  gopdf.FontDescItem{ Key:"FontBBox", Val :  "[-712 -271 1186 1048]" }
	me.desc[5] =  gopdf.FontDescItem{ Key:"ItalicAngle", Val :  "0" }
	me.desc[6] =  gopdf.FontDescItem{ Key:"StemV", Val :  "120" }
 	me.desc[7] =  gopdf.FontDescItem{ Key:"MissingWidth", Val :  "237" } 
 }
func (me * Roboto-MediumItalic)GetType() string{
	return me.fonttype
}
func (me * Roboto-MediumItalic)GetName() string{
	return me.name
}	
func (me * Roboto-MediumItalic)GetDesc() []gopdf.FontDescItem{
	return me.desc
}
func (me * Roboto-MediumItalic)GetUp() int{
	return me.up
}
func (me * Roboto-MediumItalic)GetUt()  int{
	return me.ut
}
func (me * Roboto-MediumItalic)GetCw() gopdf.FontCw{
	return me.cw
}
func (me * Roboto-MediumItalic)GetEnc() string{
	return me.enc
}
func (me * Roboto-MediumItalic)GetDiff() string {
	return me.diff
}
func (me * Roboto-MediumItalic) GetOriginalsize() int{
	return 98764
}
func (me * Roboto-MediumItalic)  SetFamily(family string){
	me.family = family
}
func (me * Roboto-MediumItalic) 	GetFamily() string{
	return me.family
}
