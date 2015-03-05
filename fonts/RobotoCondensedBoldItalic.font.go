package fonts //change this
import (
	"github.com/yaslama/gopdf"
)
type RobotoCondensedBoldItalic struct {
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
func (me * RobotoCondensedBoldItalic) Init(){
	me.cw = make(gopdf.FontCw)
	me.cw[gopdf.Chr(0)]=208
	me.cw[gopdf.Chr(1)]=208
	me.cw[gopdf.Chr(2)]=208
	me.cw[gopdf.Chr(3)]=208
	me.cw[gopdf.Chr(4)]=208
	me.cw[gopdf.Chr(5)]=208
	me.cw[gopdf.Chr(6)]=208
	me.cw[gopdf.Chr(7)]=208
	me.cw[gopdf.Chr(8)]=208
	me.cw[gopdf.Chr(9)]=208
	me.cw[gopdf.Chr(10)]=208
	me.cw[gopdf.Chr(11)]=208
	me.cw[gopdf.Chr(12)]=208
	me.cw[gopdf.Chr(13)]=208
	me.cw[gopdf.Chr(14)]=208
	me.cw[gopdf.Chr(15)]=208
	me.cw[gopdf.Chr(16)]=208
	me.cw[gopdf.Chr(17)]=208
	me.cw[gopdf.Chr(18)]=208
	me.cw[gopdf.Chr(19)]=208
	me.cw[gopdf.Chr(20)]=208
	me.cw[gopdf.Chr(21)]=208
	me.cw[gopdf.Chr(22)]=208
	me.cw[gopdf.Chr(23)]=208
	me.cw[gopdf.Chr(24)]=208
	me.cw[gopdf.Chr(25)]=208
	me.cw[gopdf.Chr(26)]=208
	me.cw[gopdf.Chr(27)]=208
	me.cw[gopdf.Chr(28)]=208
	me.cw[gopdf.Chr(29)]=208
	me.cw[gopdf.Chr(30)]=208
	me.cw[gopdf.Chr(31)]=208
	me.cw[gopdf.ToByte(" ")]=208
	me.cw[gopdf.ToByte("!")]=239
	me.cw[gopdf.ToByte("\"")]=281
	me.cw[gopdf.ToByte("#")]=463
	me.cw[gopdf.ToByte("$")]=461
	me.cw[gopdf.ToByte("%")]=607
	me.cw[gopdf.ToByte("&")]=553
	me.cw[gopdf.ToByte("'")]=143
	me.cw[gopdf.ToByte("(")]=293
	me.cw[gopdf.ToByte(")")]=292
	me.cw[gopdf.ToByte("*")]=376
	me.cw[gopdf.ToByte("+")]=447
	me.cw[gopdf.ToByte(",")]=221
	me.cw[gopdf.ToByte("-")]=347
	me.cw[gopdf.ToByte(".")]=256
	me.cw[gopdf.ToByte("/")]=294
	me.cw[gopdf.ToByte("0")]=461
	me.cw[gopdf.ToByte("1")]=461
	me.cw[gopdf.ToByte("2")]=461
	me.cw[gopdf.ToByte("3")]=461
	me.cw[gopdf.ToByte("4")]=461
	me.cw[gopdf.ToByte("5")]=461
	me.cw[gopdf.ToByte("6")]=461
	me.cw[gopdf.ToByte("7")]=461
	me.cw[gopdf.ToByte("8")]=461
	me.cw[gopdf.ToByte("9")]=461
	me.cw[gopdf.ToByte(":")]=257
	me.cw[gopdf.ToByte(";")]=246
	me.cw[gopdf.ToByte("<")]=405
	me.cw[gopdf.ToByte("=")]=462
	me.cw[gopdf.ToByte(">")]=411
	me.cw[gopdf.ToByte("?")]=411
	me.cw[gopdf.ToByte("@")]=719
	me.cw[gopdf.ToByte("A")]=515
	me.cw[gopdf.ToByte("B")]=514
	me.cw[gopdf.ToByte("C")]=504
	me.cw[gopdf.ToByte("D")]=517
	me.cw[gopdf.ToByte("E")]=455
	me.cw[gopdf.ToByte("F")]=454
	me.cw[gopdf.ToByte("G")]=516
	me.cw[gopdf.ToByte("H")]=558
	me.cw[gopdf.ToByte("I")]=248
	me.cw[gopdf.ToByte("J")]=459
	me.cw[gopdf.ToByte("K")]=494
	me.cw[gopdf.ToByte("L")]=440
	me.cw[gopdf.ToByte("M")]=664
	me.cw[gopdf.ToByte("N")]=557
	me.cw[gopdf.ToByte("O")]=542
	me.cw[gopdf.ToByte("P")]=524
	me.cw[gopdf.ToByte("Q")]=548
	me.cw[gopdf.ToByte("R")]=515
	me.cw[gopdf.ToByte("S")]=491
	me.cw[gopdf.ToByte("T")]=499
	me.cw[gopdf.ToByte("U")]=529
	me.cw[gopdf.ToByte("V")]=513
	me.cw[gopdf.ToByte("W")]=654
	me.cw[gopdf.ToByte("X")]=510
	me.cw[gopdf.ToByte("Y")]=494
	me.cw[gopdf.ToByte("Z")]=453
	me.cw[gopdf.ToByte("[")]=249
	me.cw[gopdf.ToByte("\\")]=348
	me.cw[gopdf.ToByte("]")]=249
	me.cw[gopdf.ToByte("^")]=358
	me.cw[gopdf.ToByte("_")]=362
	me.cw[gopdf.ToByte("`")]=293
	me.cw[gopdf.ToByte("a")]=427
	me.cw[gopdf.ToByte("b")]=450
	me.cw[gopdf.ToByte("c")]=418
	me.cw[gopdf.ToByte("d")]=450
	me.cw[gopdf.ToByte("e")]=424
	me.cw[gopdf.ToByte("f")]=299
	me.cw[gopdf.ToByte("g")]=450
	me.cw[gopdf.ToByte("h")]=450
	me.cw[gopdf.ToByte("i")]=233
	me.cw[gopdf.ToByte("j")]=228
	me.cw[gopdf.ToByte("k")]=439
	me.cw[gopdf.ToByte("l")]=233
	me.cw[gopdf.ToByte("m")]=671
	me.cw[gopdf.ToByte("n")]=450
	me.cw[gopdf.ToByte("o")]=450
	me.cw[gopdf.ToByte("p")]=450
	me.cw[gopdf.ToByte("q")]=450
	me.cw[gopdf.ToByte("r")]=301
	me.cw[gopdf.ToByte("s")]=403
	me.cw[gopdf.ToByte("t")]=279
	me.cw[gopdf.ToByte("u")]=450
	me.cw[gopdf.ToByte("v")]=414
	me.cw[gopdf.ToByte("w")]=573
	me.cw[gopdf.ToByte("x")]=414
	me.cw[gopdf.ToByte("y")]=414
	me.cw[gopdf.ToByte("z")]=402
	me.cw[gopdf.ToByte("{")]=269
	me.cw[gopdf.ToByte("|")]=221
	me.cw[gopdf.ToByte("}")]=269
	me.cw[gopdf.ToByte("~")]=501
	me.cw[gopdf.Chr(127)]=208
	me.cw[gopdf.Chr(128)]=440
	me.cw[gopdf.Chr(129)]=208
	me.cw[gopdf.Chr(130)]=231
	me.cw[gopdf.Chr(131)]=303
	me.cw[gopdf.Chr(132)]=361
	me.cw[gopdf.Chr(133)]=615
	me.cw[gopdf.Chr(134)]=424
	me.cw[gopdf.Chr(135)]=464
	me.cw[gopdf.Chr(136)]=417
	me.cw[gopdf.Chr(137)]=746
	me.cw[gopdf.Chr(138)]=502
	me.cw[gopdf.Chr(139)]=260
	me.cw[gopdf.Chr(140)]=768
	me.cw[gopdf.Chr(141)]=208
	me.cw[gopdf.Chr(142)]=449
	me.cw[gopdf.Chr(143)]=208
	me.cw[gopdf.Chr(144)]=208
	me.cw[gopdf.Chr(145)]=213
	me.cw[gopdf.Chr(146)]=209
	me.cw[gopdf.Chr(147)]=362
	me.cw[gopdf.Chr(148)]=364
	me.cw[gopdf.Chr(149)]=314
	me.cw[gopdf.Chr(150)]=544
	me.cw[gopdf.Chr(151)]=640
	me.cw[gopdf.Chr(152)]=399
	me.cw[gopdf.Chr(153)]=504
	me.cw[gopdf.Chr(154)]=403
	me.cw[gopdf.Chr(155)]=249
	me.cw[gopdf.Chr(156)]=699
	me.cw[gopdf.Chr(157)]=208
	me.cw[gopdf.Chr(158)]=402
	me.cw[gopdf.Chr(159)]=494
	me.cw[gopdf.Chr(160)]=208
	me.cw[gopdf.Chr(161)]=253
	me.cw[gopdf.Chr(162)]=467
	me.cw[gopdf.Chr(163)]=478
	me.cw[gopdf.Chr(164)]=562
	me.cw[gopdf.Chr(165)]=489
	me.cw[gopdf.Chr(166)]=220
	me.cw[gopdf.Chr(167)]=504
	me.cw[gopdf.Chr(168)]=458
	me.cw[gopdf.Chr(169)]=640
	me.cw[gopdf.Chr(170)]=356
	me.cw[gopdf.Chr(171)]=413
	me.cw[gopdf.Chr(172)]=439
	me.cw[gopdf.Chr(173)]=347
	me.cw[gopdf.Chr(174)]=639
	me.cw[gopdf.Chr(175)]=438
	me.cw[gopdf.Chr(176)]=331
	me.cw[gopdf.Chr(177)]=448
	me.cw[gopdf.Chr(178)]=342
	me.cw[gopdf.Chr(179)]=337
	me.cw[gopdf.Chr(180)]=320
	me.cw[gopdf.Chr(181)]=507
	me.cw[gopdf.Chr(182)]=435
	me.cw[gopdf.Chr(183)]=265
	me.cw[gopdf.Chr(184)]=229
	me.cw[gopdf.Chr(185)]=240
	me.cw[gopdf.Chr(186)]=369
	me.cw[gopdf.Chr(187)]=414
	me.cw[gopdf.Chr(188)]=564
	me.cw[gopdf.Chr(189)]=596
	me.cw[gopdf.Chr(190)]=652
	me.cw[gopdf.Chr(191)]=416
	me.cw[gopdf.Chr(192)]=515
	me.cw[gopdf.Chr(193)]=515
	me.cw[gopdf.Chr(194)]=515
	me.cw[gopdf.Chr(195)]=515
	me.cw[gopdf.Chr(196)]=515
	me.cw[gopdf.Chr(197)]=515
	me.cw[gopdf.Chr(198)]=744
	me.cw[gopdf.Chr(199)]=495
	me.cw[gopdf.Chr(200)]=455
	me.cw[gopdf.Chr(201)]=455
	me.cw[gopdf.Chr(202)]=455
	me.cw[gopdf.Chr(203)]=455
	me.cw[gopdf.Chr(204)]=248
	me.cw[gopdf.Chr(205)]=248
	me.cw[gopdf.Chr(206)]=248
	me.cw[gopdf.Chr(207)]=248
	me.cw[gopdf.Chr(208)]=540
	me.cw[gopdf.Chr(209)]=557
	me.cw[gopdf.Chr(210)]=546
	me.cw[gopdf.Chr(211)]=546
	me.cw[gopdf.Chr(212)]=546
	me.cw[gopdf.Chr(213)]=546
	me.cw[gopdf.Chr(214)]=546
	me.cw[gopdf.Chr(215)]=443
	me.cw[gopdf.Chr(216)]=546
	me.cw[gopdf.Chr(217)]=529
	me.cw[gopdf.Chr(218)]=529
	me.cw[gopdf.Chr(219)]=529
	me.cw[gopdf.Chr(220)]=529
	me.cw[gopdf.Chr(221)]=494
	me.cw[gopdf.Chr(222)]=486
	me.cw[gopdf.Chr(223)]=513
	me.cw[gopdf.Chr(224)]=427
	me.cw[gopdf.Chr(225)]=427
	me.cw[gopdf.Chr(226)]=427
	me.cw[gopdf.Chr(227)]=427
	me.cw[gopdf.Chr(228)]=427
	me.cw[gopdf.Chr(229)]=427
	me.cw[gopdf.Chr(230)]=662
	me.cw[gopdf.Chr(231)]=418
	me.cw[gopdf.Chr(232)]=424
	me.cw[gopdf.Chr(233)]=424
	me.cw[gopdf.Chr(234)]=424
	me.cw[gopdf.Chr(235)]=424
	me.cw[gopdf.Chr(236)]=243
	me.cw[gopdf.Chr(237)]=243
	me.cw[gopdf.Chr(238)]=243
	me.cw[gopdf.Chr(239)]=243
	me.cw[gopdf.Chr(240)]=475
	me.cw[gopdf.Chr(241)]=450
	me.cw[gopdf.Chr(242)]=450
	me.cw[gopdf.Chr(243)]=450
	me.cw[gopdf.Chr(244)]=450
	me.cw[gopdf.Chr(245)]=450
	me.cw[gopdf.Chr(246)]=450
	me.cw[gopdf.Chr(247)]=453
	me.cw[gopdf.Chr(248)]=451
	me.cw[gopdf.Chr(249)]=450
	me.cw[gopdf.Chr(250)]=450
	me.cw[gopdf.Chr(251)]=450
	me.cw[gopdf.Chr(252)]=450
	me.cw[gopdf.Chr(253)]=414
	me.cw[gopdf.Chr(254)]=452
	me.cw[gopdf.Chr(255)]=414
	me.up = -73
	me.ut = 49
	me.fonttype = "TrueType"
	me.name = "RobotoCondensed-BoldItalic"
	me.enc = "cp1252"
	me.desc = make([]gopdf.FontDescItem,8)
	me.desc[0] =  gopdf.FontDescItem{ Key:"Ascent",Val : "750" }
	me.desc[1] =  gopdf.FontDescItem{ Key: "Descent", Val : "-250" }
	me.desc[2] =  gopdf.FontDescItem{ Key:"CapHeight", Val :  "711" }
	me.desc[3] =  gopdf.FontDescItem{ Key: "Flags", Val :  "32" }
	me.desc[4] =  gopdf.FontDescItem{ Key:"FontBBox", Val :  "[-636 -271 1018 1048]" }
	me.desc[5] =  gopdf.FontDescItem{ Key:"ItalicAngle", Val :  "0" }
	me.desc[6] =  gopdf.FontDescItem{ Key:"StemV", Val :  "120" }
 	me.desc[7] =  gopdf.FontDescItem{ Key:"MissingWidth", Val :  "208" } 
 }
func (me * RobotoCondensedBoldItalic)GetType() string{
	return me.fonttype
}
func (me * RobotoCondensedBoldItalic)GetName() string{
	return me.name
}	
func (me * RobotoCondensedBoldItalic)GetDesc() []gopdf.FontDescItem{
	return me.desc
}
func (me * RobotoCondensedBoldItalic)GetUp() int{
	return me.up
}
func (me * RobotoCondensedBoldItalic)GetUt()  int{
	return me.ut
}
func (me * RobotoCondensedBoldItalic)GetCw() gopdf.FontCw{
	return me.cw
}
func (me * RobotoCondensedBoldItalic)GetEnc() string{
	return me.enc
}
func (me * RobotoCondensedBoldItalic)GetDiff() string {
	return me.diff
}
func (me * RobotoCondensedBoldItalic) GetOriginalsize() int{
	return 98764
}
func (me * RobotoCondensedBoldItalic)  SetFamily(family string){
	me.family = family
}
func (me * RobotoCondensedBoldItalic) 	GetFamily() string{
	return me.family
}
