package fonts //change this
import (
	"github.com/yaslama/gopdf"
)
type RobotoCondensed-Italic struct {
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
func (me * RobotoCondensed-Italic) Init(){
	me.cw = make(gopdf.FontCw)
	me.cw[gopdf.Chr(0)]=207
	me.cw[gopdf.Chr(1)]=207
	me.cw[gopdf.Chr(2)]=207
	me.cw[gopdf.Chr(3)]=207
	me.cw[gopdf.Chr(4)]=207
	me.cw[gopdf.Chr(5)]=207
	me.cw[gopdf.Chr(6)]=207
	me.cw[gopdf.Chr(7)]=207
	me.cw[gopdf.Chr(8)]=207
	me.cw[gopdf.Chr(9)]=207
	me.cw[gopdf.Chr(10)]=207
	me.cw[gopdf.Chr(11)]=207
	me.cw[gopdf.Chr(12)]=207
	me.cw[gopdf.Chr(13)]=207
	me.cw[gopdf.Chr(14)]=207
	me.cw[gopdf.Chr(15)]=207
	me.cw[gopdf.Chr(16)]=207
	me.cw[gopdf.Chr(17)]=207
	me.cw[gopdf.Chr(18)]=207
	me.cw[gopdf.Chr(19)]=207
	me.cw[gopdf.Chr(20)]=207
	me.cw[gopdf.Chr(21)]=207
	me.cw[gopdf.Chr(22)]=207
	me.cw[gopdf.Chr(23)]=207
	me.cw[gopdf.Chr(24)]=207
	me.cw[gopdf.Chr(25)]=207
	me.cw[gopdf.Chr(26)]=207
	me.cw[gopdf.Chr(27)]=207
	me.cw[gopdf.Chr(28)]=207
	me.cw[gopdf.Chr(29)]=207
	me.cw[gopdf.Chr(30)]=207
	me.cw[gopdf.Chr(31)]=207
	me.cw[gopdf.ToByte(" ")]=207
	me.cw[gopdf.ToByte("!")]=225
	me.cw[gopdf.ToByte("\"")]=275
	me.cw[gopdf.ToByte("#")]=480
	me.cw[gopdf.ToByte("$")]=452
	me.cw[gopdf.ToByte("%")]=604
	me.cw[gopdf.ToByte("&")]=530
	me.cw[gopdf.ToByte("'")]=150
	me.cw[gopdf.ToByte("(")]=279
	me.cw[gopdf.ToByte(")")]=280
	me.cw[gopdf.ToByte("*")]=367
	me.cw[gopdf.ToByte("+")]=460
	me.cw[gopdf.ToByte(",")]=189
	me.cw[gopdf.ToByte("-")]=277
	me.cw[gopdf.ToByte(".")]=234
	me.cw[gopdf.ToByte("/")]=316
	me.cw[gopdf.ToByte("0")]=452
	me.cw[gopdf.ToByte("1")]=452
	me.cw[gopdf.ToByte("2")]=452
	me.cw[gopdf.ToByte("3")]=452
	me.cw[gopdf.ToByte("4")]=452
	me.cw[gopdf.ToByte("5")]=452
	me.cw[gopdf.ToByte("6")]=452
	me.cw[gopdf.ToByte("7")]=452
	me.cw[gopdf.ToByte("8")]=452
	me.cw[gopdf.ToByte("9")]=452
	me.cw[gopdf.ToByte(":")]=227
	me.cw[gopdf.ToByte(";")]=224
	me.cw[gopdf.ToByte("<")]=405
	me.cw[gopdf.ToByte("=")]=453
	me.cw[gopdf.ToByte(">")]=414
	me.cw[gopdf.ToByte("?")]=392
	me.cw[gopdf.ToByte("@")]=729
	me.cw[gopdf.ToByte("A")]=520
	me.cw[gopdf.ToByte("B")]=508
	me.cw[gopdf.ToByte("C")]=502
	me.cw[gopdf.ToByte("D")]=519
	me.cw[gopdf.ToByte("E")]=459
	me.cw[gopdf.ToByte("F")]=458
	me.cw[gopdf.ToByte("G")]=527
	me.cw[gopdf.ToByte("H")]=560
	me.cw[gopdf.ToByte("I")]=238
	me.cw[gopdf.ToByte("J")]=447
	me.cw[gopdf.ToByte("K")]=481
	me.cw[gopdf.ToByte("L")]=438
	me.cw[gopdf.ToByte("M")]=665
	me.cw[gopdf.ToByte("N")]=560
	me.cw[gopdf.ToByte("O")]=539
	me.cw[gopdf.ToByte("P")]=513
	me.cw[gopdf.ToByte("Q")]=548
	me.cw[gopdf.ToByte("R")]=519
	me.cw[gopdf.ToByte("S")]=484
	me.cw[gopdf.ToByte("T")]=474
	me.cw[gopdf.ToByte("U")]=521
	me.cw[gopdf.ToByte("V")]=511
	me.cw[gopdf.ToByte("W")]=662
	me.cw[gopdf.ToByte("X")]=503
	me.cw[gopdf.ToByte("Y")]=482
	me.cw[gopdf.ToByte("Z")]=447
	me.cw[gopdf.ToByte("[")]=236
	me.cw[gopdf.ToByte("\\")]=338
	me.cw[gopdf.ToByte("]")]=236
	me.cw[gopdf.ToByte("^")]=345
	me.cw[gopdf.ToByte("_")]=362
	me.cw[gopdf.ToByte("`")]=275
	me.cw[gopdf.ToByte("a")]=432
	me.cw[gopdf.ToByte("b")]=450
	me.cw[gopdf.ToByte("c")]=419
	me.cw[gopdf.ToByte("d")]=450
	me.cw[gopdf.ToByte("e")]=421
	me.cw[gopdf.ToByte("f")]=287
	me.cw[gopdf.ToByte("g")]=450
	me.cw[gopdf.ToByte("h")]=450
	me.cw[gopdf.ToByte("i")]=218
	me.cw[gopdf.ToByte("j")]=219
	me.cw[gopdf.ToByte("k")]=420
	me.cw[gopdf.ToByte("l")]=218
	me.cw[gopdf.ToByte("m")]=680
	me.cw[gopdf.ToByte("n")]=450
	me.cw[gopdf.ToByte("o")]=450
	me.cw[gopdf.ToByte("p")]=450
	me.cw[gopdf.ToByte("q")]=450
	me.cw[gopdf.ToByte("r")]=291
	me.cw[gopdf.ToByte("s")]=404
	me.cw[gopdf.ToByte("t")]=268
	me.cw[gopdf.ToByte("u")]=450
	me.cw[gopdf.ToByte("v")]=406
	me.cw[gopdf.ToByte("w")]=586
	me.cw[gopdf.ToByte("x")]=406
	me.cw[gopdf.ToByte("y")]=406
	me.cw[gopdf.ToByte("z")]=394
	me.cw[gopdf.ToByte("{")]=273
	me.cw[gopdf.ToByte("|")]=211
	me.cw[gopdf.ToByte("}")]=273
	me.cw[gopdf.ToByte("~")]=521
	me.cw[gopdf.Chr(127)]=207
	me.cw[gopdf.Chr(128)]=430
	me.cw[gopdf.Chr(129)]=207
	me.cw[gopdf.Chr(130)]=195
	me.cw[gopdf.Chr(131)]=288
	me.cw[gopdf.Chr(132)]=314
	me.cw[gopdf.Chr(133)]=564
	me.cw[gopdf.Chr(134)]=432
	me.cw[gopdf.Chr(135)]=457
	me.cw[gopdf.Chr(136)]=394
	me.cw[gopdf.Chr(137)]=748
	me.cw[gopdf.Chr(138)]=491
	me.cw[gopdf.Chr(139)]=251
	me.cw[gopdf.Chr(140)]=755
	me.cw[gopdf.Chr(141)]=207
	me.cw[gopdf.Chr(142)]=455
	me.cw[gopdf.Chr(143)]=207
	me.cw[gopdf.Chr(144)]=207
	me.cw[gopdf.Chr(145)]=188
	me.cw[gopdf.Chr(146)]=186
	me.cw[gopdf.Chr(147)]=319
	me.cw[gopdf.Chr(148)]=322
	me.cw[gopdf.Chr(149)]=298
	me.cw[gopdf.Chr(150)]=545
	me.cw[gopdf.Chr(151)]=638
	me.cw[gopdf.Chr(152)]=386
	me.cw[gopdf.Chr(153)]=498
	me.cw[gopdf.Chr(154)]=404
	me.cw[gopdf.Chr(155)]=247
	me.cw[gopdf.Chr(156)]=706
	me.cw[gopdf.Chr(157)]=207
	me.cw[gopdf.Chr(158)]=394
	me.cw[gopdf.Chr(159)]=482
	me.cw[gopdf.Chr(160)]=207
	me.cw[gopdf.Chr(161)]=224
	me.cw[gopdf.Chr(162)]=449
	me.cw[gopdf.Chr(163)]=468
	me.cw[gopdf.Chr(164)]=577
	me.cw[gopdf.Chr(165)]=483
	me.cw[gopdf.Chr(166)]=208
	me.cw[gopdf.Chr(167)]=493
	me.cw[gopdf.Chr(168)]=430
	me.cw[gopdf.Chr(169)]=643
	me.cw[gopdf.Chr(170)]=357
	me.cw[gopdf.Chr(171)]=391
	me.cw[gopdf.Chr(172)]=439
	me.cw[gopdf.Chr(173)]=277
	me.cw[gopdf.Chr(174)]=643
	me.cw[gopdf.Chr(175)]=405
	me.cw[gopdf.Chr(176)]=322
	me.cw[gopdf.Chr(177)]=446
	me.cw[gopdf.Chr(178)]=341
	me.cw[gopdf.Chr(179)]=341
	me.cw[gopdf.Chr(180)]=288
	me.cw[gopdf.Chr(181)]=473
	me.cw[gopdf.Chr(182)]=409
	me.cw[gopdf.Chr(183)]=236
	me.cw[gopdf.Chr(184)]=215
	me.cw[gopdf.Chr(185)]=230
	me.cw[gopdf.Chr(186)]=366
	me.cw[gopdf.Chr(187)]=391
	me.cw[gopdf.Chr(188)]=591
	me.cw[gopdf.Chr(189)]=626
	me.cw[gopdf.Chr(190)]=668
	me.cw[gopdf.Chr(191)]=402
	me.cw[gopdf.Chr(192)]=520
	me.cw[gopdf.Chr(193)]=520
	me.cw[gopdf.Chr(194)]=520
	me.cw[gopdf.Chr(195)]=520
	me.cw[gopdf.Chr(196)]=520
	me.cw[gopdf.Chr(197)]=520
	me.cw[gopdf.Chr(198)]=736
	me.cw[gopdf.Chr(199)]=493
	me.cw[gopdf.Chr(200)]=459
	me.cw[gopdf.Chr(201)]=459
	me.cw[gopdf.Chr(202)]=459
	me.cw[gopdf.Chr(203)]=459
	me.cw[gopdf.Chr(204)]=238
	me.cw[gopdf.Chr(205)]=238
	me.cw[gopdf.Chr(206)]=238
	me.cw[gopdf.Chr(207)]=238
	me.cw[gopdf.Chr(208)]=545
	me.cw[gopdf.Chr(209)]=560
	me.cw[gopdf.Chr(210)]=546
	me.cw[gopdf.Chr(211)]=546
	me.cw[gopdf.Chr(212)]=546
	me.cw[gopdf.Chr(213)]=546
	me.cw[gopdf.Chr(214)]=546
	me.cw[gopdf.Chr(215)]=443
	me.cw[gopdf.Chr(216)]=541
	me.cw[gopdf.Chr(217)]=521
	me.cw[gopdf.Chr(218)]=521
	me.cw[gopdf.Chr(219)]=521
	me.cw[gopdf.Chr(220)]=521
	me.cw[gopdf.Chr(221)]=482
	me.cw[gopdf.Chr(222)]=475
	me.cw[gopdf.Chr(223)]=489
	me.cw[gopdf.Chr(224)]=432
	me.cw[gopdf.Chr(225)]=432
	me.cw[gopdf.Chr(226)]=432
	me.cw[gopdf.Chr(227)]=432
	me.cw[gopdf.Chr(228)]=432
	me.cw[gopdf.Chr(229)]=432
	me.cw[gopdf.Chr(230)]=662
	me.cw[gopdf.Chr(231)]=419
	me.cw[gopdf.Chr(232)]=421
	me.cw[gopdf.Chr(233)]=421
	me.cw[gopdf.Chr(234)]=421
	me.cw[gopdf.Chr(235)]=421
	me.cw[gopdf.Chr(236)]=221
	me.cw[gopdf.Chr(237)]=221
	me.cw[gopdf.Chr(238)]=221
	me.cw[gopdf.Chr(239)]=221
	me.cw[gopdf.Chr(240)]=469
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
	me.cw[gopdf.Chr(253)]=406
	me.cw[gopdf.Chr(254)]=456
	me.cw[gopdf.Chr(255)]=406
	me.up = -73
	me.ut = 49
	me.fonttype = "TrueType"
	me.name = "RobotoCondensed-Italic"
	me.enc = "cp1252"
	me.desc = make([]gopdf.FontDescItem,8)
	me.desc[0] =  gopdf.FontDescItem{ Key:"Ascent",Val : "750" }
	me.desc[1] =  gopdf.FontDescItem{ Key: "Descent", Val : "-250" }
	me.desc[2] =  gopdf.FontDescItem{ Key:"CapHeight", Val :  "711" }
	me.desc[3] =  gopdf.FontDescItem{ Key: "Flags", Val :  "32" }
	me.desc[4] =  gopdf.FontDescItem{ Key:"FontBBox", Val :  "[-563 -271 1002 1048]" }
	me.desc[5] =  gopdf.FontDescItem{ Key:"ItalicAngle", Val :  "0" }
	me.desc[6] =  gopdf.FontDescItem{ Key:"StemV", Val :  "70" }
 	me.desc[7] =  gopdf.FontDescItem{ Key:"MissingWidth", Val :  "207" } 
 }
func (me * RobotoCondensed-Italic)GetType() string{
	return me.fonttype
}
func (me * RobotoCondensed-Italic)GetName() string{
	return me.name
}	
func (me * RobotoCondensed-Italic)GetDesc() []gopdf.FontDescItem{
	return me.desc
}
func (me * RobotoCondensed-Italic)GetUp() int{
	return me.up
}
func (me * RobotoCondensed-Italic)GetUt()  int{
	return me.ut
}
func (me * RobotoCondensed-Italic)GetCw() gopdf.FontCw{
	return me.cw
}
func (me * RobotoCondensed-Italic)GetEnc() string{
	return me.enc
}
func (me * RobotoCondensed-Italic)GetDiff() string {
	return me.diff
}
func (me * RobotoCondensed-Italic) GetOriginalsize() int{
	return 98764
}
func (me * RobotoCondensed-Italic)  SetFamily(family string){
	me.family = family
}
func (me * RobotoCondensed-Italic) 	GetFamily() string{
	return me.family
}
