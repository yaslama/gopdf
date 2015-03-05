package fonts //change this
import (
	"github.com/yaslama/gopdf"
)
type RobotoCondensed-Regular struct {
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
func (me * RobotoCondensed-Regular) Init(){
	me.cw = make(gopdf.FontCw)
	me.cw[gopdf.Chr(0)]=211
	me.cw[gopdf.Chr(1)]=211
	me.cw[gopdf.Chr(2)]=211
	me.cw[gopdf.Chr(3)]=211
	me.cw[gopdf.Chr(4)]=211
	me.cw[gopdf.Chr(5)]=211
	me.cw[gopdf.Chr(6)]=211
	me.cw[gopdf.Chr(7)]=211
	me.cw[gopdf.Chr(8)]=211
	me.cw[gopdf.Chr(9)]=211
	me.cw[gopdf.Chr(10)]=211
	me.cw[gopdf.Chr(11)]=211
	me.cw[gopdf.Chr(12)]=211
	me.cw[gopdf.Chr(13)]=211
	me.cw[gopdf.Chr(14)]=211
	me.cw[gopdf.Chr(15)]=211
	me.cw[gopdf.Chr(16)]=211
	me.cw[gopdf.Chr(17)]=211
	me.cw[gopdf.Chr(18)]=211
	me.cw[gopdf.Chr(19)]=211
	me.cw[gopdf.Chr(20)]=211
	me.cw[gopdf.Chr(21)]=211
	me.cw[gopdf.Chr(22)]=211
	me.cw[gopdf.Chr(23)]=211
	me.cw[gopdf.Chr(24)]=211
	me.cw[gopdf.Chr(25)]=211
	me.cw[gopdf.Chr(26)]=211
	me.cw[gopdf.Chr(27)]=211
	me.cw[gopdf.Chr(28)]=211
	me.cw[gopdf.Chr(29)]=211
	me.cw[gopdf.Chr(30)]=211
	me.cw[gopdf.Chr(31)]=211
	me.cw[gopdf.ToByte(" ")]=211
	me.cw[gopdf.ToByte("!")]=228
	me.cw[gopdf.ToByte("\"")]=281
	me.cw[gopdf.ToByte("#")]=500
	me.cw[gopdf.ToByte("$")]=469
	me.cw[gopdf.ToByte("%")]=629
	me.cw[gopdf.ToByte("&")]=549
	me.cw[gopdf.ToByte("'")]=152
	me.cw[gopdf.ToByte("(")]=286
	me.cw[gopdf.ToByte(")")]=287
	me.cw[gopdf.ToByte("*")]=379
	me.cw[gopdf.ToByte("+")]=478
	me.cw[gopdf.ToByte(",")]=192
	me.cw[gopdf.ToByte("-")]=284
	me.cw[gopdf.ToByte(".")]=238
	me.cw[gopdf.ToByte("/")]=327
	me.cw[gopdf.ToByte("0")]=469
	me.cw[gopdf.ToByte("1")]=469
	me.cw[gopdf.ToByte("2")]=469
	me.cw[gopdf.ToByte("3")]=469
	me.cw[gopdf.ToByte("4")]=469
	me.cw[gopdf.ToByte("5")]=469
	me.cw[gopdf.ToByte("6")]=469
	me.cw[gopdf.ToByte("7")]=469
	me.cw[gopdf.ToByte("8")]=469
	me.cw[gopdf.ToByte("9")]=469
	me.cw[gopdf.ToByte(":")]=229
	me.cw[gopdf.ToByte(";")]=227
	me.cw[gopdf.ToByte("<")]=421
	me.cw[gopdf.ToByte("=")]=471
	me.cw[gopdf.ToByte(">")]=430
	me.cw[gopdf.ToByte("?")]=404
	me.cw[gopdf.ToByte("@")]=761
	me.cw[gopdf.ToByte("A")]=539
	me.cw[gopdf.ToByte("B")]=527
	me.cw[gopdf.ToByte("C")]=514
	me.cw[gopdf.ToByte("D")]=552
	me.cw[gopdf.ToByte("E")]=477
	me.cw[gopdf.ToByte("F")]=476
	me.cw[gopdf.ToByte("G")]=549
	me.cw[gopdf.ToByte("H")]=583
	me.cw[gopdf.ToByte("I")]=244
	me.cw[gopdf.ToByte("J")]=464
	me.cw[gopdf.ToByte("K")]=530
	me.cw[gopdf.ToByte("L")]=453
	me.cw[gopdf.ToByte("M")]=694
	me.cw[gopdf.ToByte("N")]=583
	me.cw[gopdf.ToByte("O")]=568
	me.cw[gopdf.ToByte("P")]=532
	me.cw[gopdf.ToByte("Q")]=570
	me.cw[gopdf.ToByte("R")]=541
	me.cw[gopdf.ToByte("S")]=510
	me.cw[gopdf.ToByte("T")]=493
	me.cw[gopdf.ToByte("U")]=543
	me.cw[gopdf.ToByte("V")]=530
	me.cw[gopdf.ToByte("W")]=693
	me.cw[gopdf.ToByte("X")]=522
	me.cw[gopdf.ToByte("Y")]=501
	me.cw[gopdf.ToByte("Z")]=474
	me.cw[gopdf.ToByte("[")]=240
	me.cw[gopdf.ToByte("\\")]=348
	me.cw[gopdf.ToByte("]")]=240
	me.cw[gopdf.ToByte("^")]=357
	me.cw[gopdf.ToByte("_")]=374
	me.cw[gopdf.ToByte("`")]=281
	me.cw[gopdf.ToByte("a")]=448
	me.cw[gopdf.ToByte("b")]=467
	me.cw[gopdf.ToByte("c")]=435
	me.cw[gopdf.ToByte("d")]=467
	me.cw[gopdf.ToByte("e")]=437
	me.cw[gopdf.ToByte("f")]=295
	me.cw[gopdf.ToByte("g")]=467
	me.cw[gopdf.ToByte("h")]=467
	me.cw[gopdf.ToByte("i")]=221
	me.cw[gopdf.ToByte("j")]=223
	me.cw[gopdf.ToByte("k")]=434
	me.cw[gopdf.ToByte("l")]=221
	me.cw[gopdf.ToByte("m")]=710
	me.cw[gopdf.ToByte("n")]=467
	me.cw[gopdf.ToByte("o")]=467
	me.cw[gopdf.ToByte("p")]=467
	me.cw[gopdf.ToByte("q")]=467
	me.cw[gopdf.ToByte("r")]=299
	me.cw[gopdf.ToByte("s")]=419
	me.cw[gopdf.ToByte("t")]=275
	me.cw[gopdf.ToByte("u")]=467
	me.cw[gopdf.ToByte("v")]=420
	me.cw[gopdf.ToByte("w")]=611
	me.cw[gopdf.ToByte("x")]=420
	me.cw[gopdf.ToByte("y")]=420
	me.cw[gopdf.ToByte("z")]=408
	me.cw[gopdf.ToByte("{")]=281
	me.cw[gopdf.ToByte("|")]=214
	me.cw[gopdf.ToByte("}")]=281
	me.cw[gopdf.ToByte("~")]=544
	me.cw[gopdf.Chr(127)]=211
	me.cw[gopdf.Chr(128)]=445
	me.cw[gopdf.Chr(129)]=211
	me.cw[gopdf.Chr(130)]=196
	me.cw[gopdf.Chr(131)]=295
	me.cw[gopdf.Chr(132)]=319
	me.cw[gopdf.Chr(133)]=584
	me.cw[gopdf.Chr(134)]=449
	me.cw[gopdf.Chr(135)]=475
	me.cw[gopdf.Chr(136)]=405
	me.cw[gopdf.Chr(137)]=782
	me.cw[gopdf.Chr(138)]=510
	me.cw[gopdf.Chr(139)]=258
	me.cw[gopdf.Chr(140)]=787
	me.cw[gopdf.Chr(141)]=211
	me.cw[gopdf.Chr(142)]=474
	me.cw[gopdf.Chr(143)]=211
	me.cw[gopdf.Chr(144)]=211
	me.cw[gopdf.Chr(145)]=189
	me.cw[gopdf.Chr(146)]=188
	me.cw[gopdf.Chr(147)]=324
	me.cw[gopdf.Chr(148)]=327
	me.cw[gopdf.Chr(149)]=305
	me.cw[gopdf.Chr(150)]=567
	me.cw[gopdf.Chr(151)]=666
	me.cw[gopdf.Chr(152)]=397
	me.cw[gopdf.Chr(153)]=518
	me.cw[gopdf.Chr(154)]=419
	me.cw[gopdf.Chr(155)]=254
	me.cw[gopdf.Chr(156)]=738
	me.cw[gopdf.Chr(157)]=211
	me.cw[gopdf.Chr(158)]=408
	me.cw[gopdf.Chr(159)]=501
	me.cw[gopdf.Chr(160)]=211
	me.cw[gopdf.Chr(161)]=227
	me.cw[gopdf.Chr(162)]=465
	me.cw[gopdf.Chr(163)]=485
	me.cw[gopdf.Chr(164)]=601
	me.cw[gopdf.Chr(165)]=502
	me.cw[gopdf.Chr(166)]=211
	me.cw[gopdf.Chr(167)]=512
	me.cw[gopdf.Chr(168)]=443
	me.cw[gopdf.Chr(169)]=671
	me.cw[gopdf.Chr(170)]=370
	me.cw[gopdf.Chr(171)]=404
	me.cw[gopdf.Chr(172)]=456
	me.cw[gopdf.Chr(173)]=284
	me.cw[gopdf.Chr(174)]=671
	me.cw[gopdf.Chr(175)]=416
	me.cw[gopdf.Chr(176)]=332
	me.cw[gopdf.Chr(177)]=462
	me.cw[gopdf.Chr(178)]=352
	me.cw[gopdf.Chr(179)]=352
	me.cw[gopdf.Chr(180)]=293
	me.cw[gopdf.Chr(181)]=490
	me.cw[gopdf.Chr(182)]=422
	me.cw[gopdf.Chr(183)]=241
	me.cw[gopdf.Chr(184)]=220
	me.cw[gopdf.Chr(185)]=234
	me.cw[gopdf.Chr(186)]=379
	me.cw[gopdf.Chr(187)]=404
	me.cw[gopdf.Chr(188)]=616
	me.cw[gopdf.Chr(189)]=649
	me.cw[gopdf.Chr(190)]=695
	me.cw[gopdf.Chr(191)]=415
	me.cw[gopdf.Chr(192)]=539
	me.cw[gopdf.Chr(193)]=539
	me.cw[gopdf.Chr(194)]=539
	me.cw[gopdf.Chr(195)]=539
	me.cw[gopdf.Chr(196)]=539
	me.cw[gopdf.Chr(197)]=539
	me.cw[gopdf.Chr(198)]=767
	me.cw[gopdf.Chr(199)]=514
	me.cw[gopdf.Chr(200)]=477
	me.cw[gopdf.Chr(201)]=477
	me.cw[gopdf.Chr(202)]=477
	me.cw[gopdf.Chr(203)]=477
	me.cw[gopdf.Chr(204)]=244
	me.cw[gopdf.Chr(205)]=244
	me.cw[gopdf.Chr(206)]=244
	me.cw[gopdf.Chr(207)]=244
	me.cw[gopdf.Chr(208)]=566
	me.cw[gopdf.Chr(209)]=583
	me.cw[gopdf.Chr(210)]=568
	me.cw[gopdf.Chr(211)]=568
	me.cw[gopdf.Chr(212)]=568
	me.cw[gopdf.Chr(213)]=568
	me.cw[gopdf.Chr(214)]=568
	me.cw[gopdf.Chr(215)]=458
	me.cw[gopdf.Chr(216)]=563
	me.cw[gopdf.Chr(217)]=543
	me.cw[gopdf.Chr(218)]=543
	me.cw[gopdf.Chr(219)]=543
	me.cw[gopdf.Chr(220)]=543
	me.cw[gopdf.Chr(221)]=501
	me.cw[gopdf.Chr(222)]=494
	me.cw[gopdf.Chr(223)]=507
	me.cw[gopdf.Chr(224)]=448
	me.cw[gopdf.Chr(225)]=448
	me.cw[gopdf.Chr(226)]=448
	me.cw[gopdf.Chr(227)]=448
	me.cw[gopdf.Chr(228)]=448
	me.cw[gopdf.Chr(229)]=448
	me.cw[gopdf.Chr(230)]=691
	me.cw[gopdf.Chr(231)]=435
	me.cw[gopdf.Chr(232)]=437
	me.cw[gopdf.Chr(233)]=437
	me.cw[gopdf.Chr(234)]=437
	me.cw[gopdf.Chr(235)]=437
	me.cw[gopdf.Chr(236)]=224
	me.cw[gopdf.Chr(237)]=224
	me.cw[gopdf.Chr(238)]=224
	me.cw[gopdf.Chr(239)]=224
	me.cw[gopdf.Chr(240)]=487
	me.cw[gopdf.Chr(241)]=467
	me.cw[gopdf.Chr(242)]=467
	me.cw[gopdf.Chr(243)]=467
	me.cw[gopdf.Chr(244)]=467
	me.cw[gopdf.Chr(245)]=467
	me.cw[gopdf.Chr(246)]=467
	me.cw[gopdf.Chr(247)]=471
	me.cw[gopdf.Chr(248)]=468
	me.cw[gopdf.Chr(249)]=467
	me.cw[gopdf.Chr(250)]=467
	me.cw[gopdf.Chr(251)]=467
	me.cw[gopdf.Chr(252)]=467
	me.cw[gopdf.Chr(253)]=420
	me.cw[gopdf.Chr(254)]=473
	me.cw[gopdf.Chr(255)]=420
	me.up = -73
	me.ut = 49
	me.fonttype = "TrueType"
	me.name = "RobotoCondensed-Regular"
	me.enc = "cp1252"
	me.desc = make([]gopdf.FontDescItem,8)
	me.desc[0] =  gopdf.FontDescItem{ Key:"Ascent",Val : "750" }
	me.desc[1] =  gopdf.FontDescItem{ Key: "Descent", Val : "-250" }
	me.desc[2] =  gopdf.FontDescItem{ Key:"CapHeight", Val :  "711" }
	me.desc[3] =  gopdf.FontDescItem{ Key: "Flags", Val :  "32" }
	me.desc[4] =  gopdf.FontDescItem{ Key:"FontBBox", Val :  "[-579 -271 976 1048]" }
	me.desc[5] =  gopdf.FontDescItem{ Key:"ItalicAngle", Val :  "0" }
	me.desc[6] =  gopdf.FontDescItem{ Key:"StemV", Val :  "70" }
 	me.desc[7] =  gopdf.FontDescItem{ Key:"MissingWidth", Val :  "211" } 
 }
func (me * RobotoCondensed-Regular)GetType() string{
	return me.fonttype
}
func (me * RobotoCondensed-Regular)GetName() string{
	return me.name
}	
func (me * RobotoCondensed-Regular)GetDesc() []gopdf.FontDescItem{
	return me.desc
}
func (me * RobotoCondensed-Regular)GetUp() int{
	return me.up
}
func (me * RobotoCondensed-Regular)GetUt()  int{
	return me.ut
}
func (me * RobotoCondensed-Regular)GetCw() gopdf.FontCw{
	return me.cw
}
func (me * RobotoCondensed-Regular)GetEnc() string{
	return me.enc
}
func (me * RobotoCondensed-Regular)GetDiff() string {
	return me.diff
}
func (me * RobotoCondensed-Regular) GetOriginalsize() int{
	return 98764
}
func (me * RobotoCondensed-Regular)  SetFamily(family string){
	me.family = family
}
func (me * RobotoCondensed-Regular) 	GetFamily() string{
	return me.family
}
