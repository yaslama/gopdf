package fonts //change this
import (
	"github.com/signintech/gopdf"
)
type Roboto-BlackItalic struct {
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
func (me * Roboto-BlackItalic) Init(){
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
	me.cw[gopdf.ToByte("!")]=267
	me.cw[gopdf.ToByte("\"")]=304
	me.cw[gopdf.ToByte("#")]=542
	me.cw[gopdf.ToByte("$")]=546
	me.cw[gopdf.ToByte("%")]=695
	me.cw[gopdf.ToByte("&")]=636
	me.cw[gopdf.ToByte("'")]=147
	me.cw[gopdf.ToByte("(")]=339
	me.cw[gopdf.ToByte(")")]=338
	me.cw[gopdf.ToByte("*")]=424
	me.cw[gopdf.ToByte("+")]=500
	me.cw[gopdf.ToByte(",")]=261
	me.cw[gopdf.ToByte("-")]=434
	me.cw[gopdf.ToByte(".")]=292
	me.cw[gopdf.ToByte("/")]=329
	me.cw[gopdf.ToByte("0")]=546
	me.cw[gopdf.ToByte("1")]=546
	me.cw[gopdf.ToByte("2")]=546
	me.cw[gopdf.ToByte("3")]=546
	me.cw[gopdf.ToByte("4")]=546
	me.cw[gopdf.ToByte("5")]=546
	me.cw[gopdf.ToByte("6")]=546
	me.cw[gopdf.ToByte("7")]=546
	me.cw[gopdf.ToByte("8")]=546
	me.cw[gopdf.ToByte("9")]=546
	me.cw[gopdf.ToByte(":")]=292
	me.cw[gopdf.ToByte(";")]=275
	me.cw[gopdf.ToByte("<")]=479
	me.cw[gopdf.ToByte("=")]=549
	me.cw[gopdf.ToByte(">")]=481
	me.cw[gopdf.ToByte("?")]=483
	me.cw[gopdf.ToByte("@")]=823
	me.cw[gopdf.ToByte("A")]=597
	me.cw[gopdf.ToByte("B")]=605
	me.cw[gopdf.ToByte("C")]=602
	me.cw[gopdf.ToByte("D")]=612
	me.cw[gopdf.ToByte("E")]=534
	me.cw[gopdf.ToByte("F")]=534
	me.cw[gopdf.ToByte("G")]=618
	me.cw[gopdf.ToByte("H")]=661
	me.cw[gopdf.ToByte("I")]=285
	me.cw[gopdf.ToByte("J")]=546
	me.cw[gopdf.ToByte("K")]=584
	me.cw[gopdf.ToByte("L")]=514
	me.cw[gopdf.ToByte("M")]=814
	me.cw[gopdf.ToByte("N")]=661
	me.cw[gopdf.ToByte("O")]=642
	me.cw[gopdf.ToByte("P")]=619
	me.cw[gopdf.ToByte("Q")]=646
	me.cw[gopdf.ToByte("R")]=608
	me.cw[gopdf.ToByte("S")]=582
	me.cw[gopdf.ToByte("T")]=617
	me.cw[gopdf.ToByte("U")]=650
	me.cw[gopdf.ToByte("V")]=599
	me.cw[gopdf.ToByte("W")]=811
	me.cw[gopdf.ToByte("X")]=601
	me.cw[gopdf.ToByte("Y")]=603
	me.cw[gopdf.ToByte("Z")]=534
	me.cw[gopdf.ToByte("[")]=271
	me.cw[gopdf.ToByte("\\")]=404
	me.cw[gopdf.ToByte("]")]=271
	me.cw[gopdf.ToByte("^")]=423
	me.cw[gopdf.ToByte("_")]=417
	me.cw[gopdf.ToByte("`")]=326
	me.cw[gopdf.ToByte("a")]=498
	me.cw[gopdf.ToByte("b")]=529
	me.cw[gopdf.ToByte("c")]=480
	me.cw[gopdf.ToByte("d")]=529
	me.cw[gopdf.ToByte("e")]=499
	me.cw[gopdf.ToByte("f")]=346
	me.cw[gopdf.ToByte("g")]=528
	me.cw[gopdf.ToByte("h")]=529
	me.cw[gopdf.ToByte("i")]=264
	me.cw[gopdf.ToByte("j")]=257
	me.cw[gopdf.ToByte("k")]=518
	me.cw[gopdf.ToByte("l")]=264
	me.cw[gopdf.ToByte("m")]=803
	me.cw[gopdf.ToByte("n")]=529
	me.cw[gopdf.ToByte("o")]=529
	me.cw[gopdf.ToByte("p")]=529
	me.cw[gopdf.ToByte("q")]=529
	me.cw[gopdf.ToByte("r")]=351
	me.cw[gopdf.ToByte("s")]=482
	me.cw[gopdf.ToByte("t")]=327
	me.cw[gopdf.ToByte("u")]=529
	me.cw[gopdf.ToByte("v")]=486
	me.cw[gopdf.ToByte("w")]=677
	me.cw[gopdf.ToByte("x")]=486
	me.cw[gopdf.ToByte("y")]=486
	me.cw[gopdf.ToByte("z")]=486
	me.cw[gopdf.ToByte("{")]=307
	me.cw[gopdf.ToByte("|")]=247
	me.cw[gopdf.ToByte("}")]=307
	me.cw[gopdf.ToByte("~")]=589
	me.cw[gopdf.Chr(127)]=237
	me.cw[gopdf.Chr(128)]=521
	me.cw[gopdf.Chr(129)]=237
	me.cw[gopdf.Chr(130)]=266
	me.cw[gopdf.Chr(131)]=353
	me.cw[gopdf.Chr(132)]=416
	me.cw[gopdf.Chr(133)]=737
	me.cw[gopdf.Chr(134)]=498
	me.cw[gopdf.Chr(135)]=548
	me.cw[gopdf.Chr(136)]=482
	me.cw[gopdf.Chr(137)]=898
	me.cw[gopdf.Chr(138)]=595
	me.cw[gopdf.Chr(139)]=302
	me.cw[gopdf.Chr(140)]=915
	me.cw[gopdf.Chr(141)]=237
	me.cw[gopdf.Chr(142)]=547
	me.cw[gopdf.Chr(143)]=237
	me.cw[gopdf.Chr(144)]=237
	me.cw[gopdf.Chr(145)]=244
	me.cw[gopdf.Chr(146)]=237
	me.cw[gopdf.Chr(147)]=415
	me.cw[gopdf.Chr(148)]=418
	me.cw[gopdf.Chr(149)]=353
	me.cw[gopdf.Chr(150)]=646
	me.cw[gopdf.Chr(151)]=766
	me.cw[gopdf.Chr(152)]=459
	me.cw[gopdf.Chr(153)]=596
	me.cw[gopdf.Chr(154)]=482
	me.cw[gopdf.Chr(155)]=287
	me.cw[gopdf.Chr(156)]=840
	me.cw[gopdf.Chr(157)]=237
	me.cw[gopdf.Chr(158)]=486
	me.cw[gopdf.Chr(159)]=603
	me.cw[gopdf.Chr(160)]=237
	me.cw[gopdf.Chr(161)]=292
	me.cw[gopdf.Chr(162)]=558
	me.cw[gopdf.Chr(163)]=565
	me.cw[gopdf.Chr(164)]=636
	me.cw[gopdf.Chr(165)]=578
	me.cw[gopdf.Chr(166)]=247
	me.cw[gopdf.Chr(167)]=598
	me.cw[gopdf.Chr(168)]=522
	me.cw[gopdf.Chr(169)]=733
	me.cw[gopdf.Chr(170)]=416
	me.cw[gopdf.Chr(171)]=489
	me.cw[gopdf.Chr(172)]=516
	me.cw[gopdf.Chr(173)]=434
	me.cw[gopdf.Chr(174)]=733
	me.cw[gopdf.Chr(175)]=497
	me.cw[gopdf.Chr(176)]=375
	me.cw[gopdf.Chr(177)]=507
	me.cw[gopdf.Chr(178)]=396
	me.cw[gopdf.Chr(179)]=390
	me.cw[gopdf.Chr(180)]=359
	me.cw[gopdf.Chr(181)]=608
	me.cw[gopdf.Chr(182)]=516
	me.cw[gopdf.Chr(183)]=310
	me.cw[gopdf.Chr(184)]=266
	me.cw[gopdf.Chr(185)]=271
	me.cw[gopdf.Chr(186)]=432
	me.cw[gopdf.Chr(187)]=489
	me.cw[gopdf.Chr(188)]=656
	me.cw[gopdf.Chr(189)]=675
	me.cw[gopdf.Chr(190)]=761
	me.cw[gopdf.Chr(191)]=485
	me.cw[gopdf.Chr(192)]=597
	me.cw[gopdf.Chr(193)]=597
	me.cw[gopdf.Chr(194)]=597
	me.cw[gopdf.Chr(195)]=597
	me.cw[gopdf.Chr(196)]=597
	me.cw[gopdf.Chr(197)]=597
	me.cw[gopdf.Chr(198)]=883
	me.cw[gopdf.Chr(199)]=604
	me.cw[gopdf.Chr(200)]=534
	me.cw[gopdf.Chr(201)]=534
	me.cw[gopdf.Chr(202)]=534
	me.cw[gopdf.Chr(203)]=534
	me.cw[gopdf.Chr(204)]=285
	me.cw[gopdf.Chr(205)]=285
	me.cw[gopdf.Chr(206)]=285
	me.cw[gopdf.Chr(207)]=285
	me.cw[gopdf.Chr(208)]=633
	me.cw[gopdf.Chr(209)]=661
	me.cw[gopdf.Chr(210)]=646
	me.cw[gopdf.Chr(211)]=646
	me.cw[gopdf.Chr(212)]=646
	me.cw[gopdf.Chr(213)]=646
	me.cw[gopdf.Chr(214)]=646
	me.cw[gopdf.Chr(215)]=498
	me.cw[gopdf.Chr(216)]=648
	me.cw[gopdf.Chr(217)]=650
	me.cw[gopdf.Chr(218)]=650
	me.cw[gopdf.Chr(219)]=650
	me.cw[gopdf.Chr(220)]=650
	me.cw[gopdf.Chr(221)]=603
	me.cw[gopdf.Chr(222)]=582
	me.cw[gopdf.Chr(223)]=614
	me.cw[gopdf.Chr(224)]=498
	me.cw[gopdf.Chr(225)]=498
	me.cw[gopdf.Chr(226)]=498
	me.cw[gopdf.Chr(227)]=498
	me.cw[gopdf.Chr(228)]=498
	me.cw[gopdf.Chr(229)]=498
	me.cw[gopdf.Chr(230)]=791
	me.cw[gopdf.Chr(231)]=480
	me.cw[gopdf.Chr(232)]=499
	me.cw[gopdf.Chr(233)]=499
	me.cw[gopdf.Chr(234)]=499
	me.cw[gopdf.Chr(235)]=499
	me.cw[gopdf.Chr(236)]=276
	me.cw[gopdf.Chr(237)]=276
	me.cw[gopdf.Chr(238)]=276
	me.cw[gopdf.Chr(239)]=276
	me.cw[gopdf.Chr(240)]=563
	me.cw[gopdf.Chr(241)]=529
	me.cw[gopdf.Chr(242)]=529
	me.cw[gopdf.Chr(243)]=529
	me.cw[gopdf.Chr(244)]=529
	me.cw[gopdf.Chr(245)]=529
	me.cw[gopdf.Chr(246)]=529
	me.cw[gopdf.Chr(247)]=535
	me.cw[gopdf.Chr(248)]=529
	me.cw[gopdf.Chr(249)]=529
	me.cw[gopdf.Chr(250)]=529
	me.cw[gopdf.Chr(251)]=529
	me.cw[gopdf.Chr(252)]=529
	me.cw[gopdf.Chr(253)]=486
	me.cw[gopdf.Chr(254)]=528
	me.cw[gopdf.Chr(255)]=486
	me.up = -73
	me.ut = 49
	me.fonttype = "TrueType"
	me.name = "Roboto-BlackItalic"
	me.enc = "cp1252"
	me.desc = make([]gopdf.FontDescItem,8)
	me.desc[0] =  gopdf.FontDescItem{ Key:"Ascent",Val : "750" }
	me.desc[1] =  gopdf.FontDescItem{ Key: "Descent", Val : "-250" }
	me.desc[2] =  gopdf.FontDescItem{ Key:"CapHeight", Val :  "711" }
	me.desc[3] =  gopdf.FontDescItem{ Key: "Flags", Val :  "32" }
	me.desc[4] =  gopdf.FontDescItem{ Key:"FontBBox", Val :  "[-700 -271 1198 1064]" }
	me.desc[5] =  gopdf.FontDescItem{ Key:"ItalicAngle", Val :  "0" }
	me.desc[6] =  gopdf.FontDescItem{ Key:"StemV", Val :  "120" }
 	me.desc[7] =  gopdf.FontDescItem{ Key:"MissingWidth", Val :  "237" } 
 }
func (me * Roboto-BlackItalic)GetType() string{
	return me.fonttype
}
func (me * Roboto-BlackItalic)GetName() string{
	return me.name
}	
func (me * Roboto-BlackItalic)GetDesc() []gopdf.FontDescItem{
	return me.desc
}
func (me * Roboto-BlackItalic)GetUp() int{
	return me.up
}
func (me * Roboto-BlackItalic)GetUt()  int{
	return me.ut
}
func (me * Roboto-BlackItalic)GetCw() gopdf.FontCw{
	return me.cw
}
func (me * Roboto-BlackItalic)GetEnc() string{
	return me.enc
}
func (me * Roboto-BlackItalic)GetDiff() string {
	return me.diff
}
func (me * Roboto-BlackItalic) GetOriginalsize() int{
	return 98764
}
func (me * Roboto-BlackItalic)  SetFamily(family string){
	me.family = family
}
func (me * Roboto-BlackItalic) 	GetFamily() string{
	return me.family
}
