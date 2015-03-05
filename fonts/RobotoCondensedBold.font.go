package fonts //change this
import (
	"github.com/yaslama/gopdf"
)
type RobotoCondensedBold struct {
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
func (me * RobotoCondensedBold) Init(){
	me.cw = make(gopdf.FontCw)
	me.cw[gopdf.Chr(0)]=213
	me.cw[gopdf.Chr(1)]=213
	me.cw[gopdf.Chr(2)]=213
	me.cw[gopdf.Chr(3)]=213
	me.cw[gopdf.Chr(4)]=213
	me.cw[gopdf.Chr(5)]=213
	me.cw[gopdf.Chr(6)]=213
	me.cw[gopdf.Chr(7)]=213
	me.cw[gopdf.Chr(8)]=213
	me.cw[gopdf.Chr(9)]=213
	me.cw[gopdf.Chr(10)]=213
	me.cw[gopdf.Chr(11)]=213
	me.cw[gopdf.Chr(12)]=213
	me.cw[gopdf.Chr(13)]=213
	me.cw[gopdf.Chr(14)]=213
	me.cw[gopdf.Chr(15)]=213
	me.cw[gopdf.Chr(16)]=213
	me.cw[gopdf.Chr(17)]=213
	me.cw[gopdf.Chr(18)]=213
	me.cw[gopdf.Chr(19)]=213
	me.cw[gopdf.Chr(20)]=213
	me.cw[gopdf.Chr(21)]=213
	me.cw[gopdf.Chr(22)]=213
	me.cw[gopdf.Chr(23)]=213
	me.cw[gopdf.Chr(24)]=213
	me.cw[gopdf.Chr(25)]=213
	me.cw[gopdf.Chr(26)]=213
	me.cw[gopdf.Chr(27)]=213
	me.cw[gopdf.Chr(28)]=213
	me.cw[gopdf.Chr(29)]=213
	me.cw[gopdf.Chr(30)]=213
	me.cw[gopdf.Chr(31)]=213
	me.cw[gopdf.ToByte(" ")]=213
	me.cw[gopdf.ToByte("!")]=242
	me.cw[gopdf.ToByte("\"")]=287
	me.cw[gopdf.ToByte("#")]=483
	me.cw[gopdf.ToByte("$")]=479
	me.cw[gopdf.ToByte("%")]=632
	me.cw[gopdf.ToByte("&")]=572
	me.cw[gopdf.ToByte("'")]=145
	me.cw[gopdf.ToByte("(")]=300
	me.cw[gopdf.ToByte(")")]=299
	me.cw[gopdf.ToByte("*")]=388
	me.cw[gopdf.ToByte("+")]=465
	me.cw[gopdf.ToByte(",")]=223
	me.cw[gopdf.ToByte("-")]=354
	me.cw[gopdf.ToByte(".")]=260
	me.cw[gopdf.ToByte("/")]=305
	me.cw[gopdf.ToByte("0")]=479
	me.cw[gopdf.ToByte("1")]=479
	me.cw[gopdf.ToByte("2")]=479
	me.cw[gopdf.ToByte("3")]=479
	me.cw[gopdf.ToByte("4")]=479
	me.cw[gopdf.ToByte("5")]=479
	me.cw[gopdf.ToByte("6")]=479
	me.cw[gopdf.ToByte("7")]=479
	me.cw[gopdf.ToByte("8")]=479
	me.cw[gopdf.ToByte("9")]=479
	me.cw[gopdf.ToByte(":")]=259
	me.cw[gopdf.ToByte(";")]=248
	me.cw[gopdf.ToByte("<")]=421
	me.cw[gopdf.ToByte("=")]=480
	me.cw[gopdf.ToByte(">")]=426
	me.cw[gopdf.ToByte("?")]=424
	me.cw[gopdf.ToByte("@")]=751
	me.cw[gopdf.ToByte("A")]=534
	me.cw[gopdf.ToByte("B")]=533
	me.cw[gopdf.ToByte("C")]=516
	me.cw[gopdf.ToByte("D")]=547
	me.cw[gopdf.ToByte("E")]=473
	me.cw[gopdf.ToByte("F")]=472
	me.cw[gopdf.ToByte("G")]=539
	me.cw[gopdf.ToByte("H")]=581
	me.cw[gopdf.ToByte("I")]=253
	me.cw[gopdf.ToByte("J")]=476
	me.cw[gopdf.ToByte("K")]=533
	me.cw[gopdf.ToByte("L")]=456
	me.cw[gopdf.ToByte("M")]=693
	me.cw[gopdf.ToByte("N")]=581
	me.cw[gopdf.ToByte("O")]=568
	me.cw[gopdf.ToByte("P")]=543
	me.cw[gopdf.ToByte("Q")]=570
	me.cw[gopdf.ToByte("R")]=536
	me.cw[gopdf.ToByte("S")]=521
	me.cw[gopdf.ToByte("T")]=518
	me.cw[gopdf.ToByte("U")]=551
	me.cw[gopdf.ToByte("V")]=532
	me.cw[gopdf.ToByte("W")]=685
	me.cw[gopdf.ToByte("X")]=529
	me.cw[gopdf.ToByte("Y")]=513
	me.cw[gopdf.ToByte("Z")]=468
	me.cw[gopdf.ToByte("[")]=253
	me.cw[gopdf.ToByte("\\")]=358
	me.cw[gopdf.ToByte("]")]=253
	me.cw[gopdf.ToByte("^")]=370
	me.cw[gopdf.ToByte("_")]=374
	me.cw[gopdf.ToByte("`")]=299
	me.cw[gopdf.ToByte("a")]=443
	me.cw[gopdf.ToByte("b")]=467
	me.cw[gopdf.ToByte("c")]=434
	me.cw[gopdf.ToByte("d")]=467
	me.cw[gopdf.ToByte("e")]=439
	me.cw[gopdf.ToByte("f")]=307
	me.cw[gopdf.ToByte("g")]=467
	me.cw[gopdf.ToByte("h")]=467
	me.cw[gopdf.ToByte("i")]=236
	me.cw[gopdf.ToByte("j")]=232
	me.cw[gopdf.ToByte("k")]=454
	me.cw[gopdf.ToByte("l")]=236
	me.cw[gopdf.ToByte("m")]=702
	me.cw[gopdf.ToByte("n")]=467
	me.cw[gopdf.ToByte("o")]=467
	me.cw[gopdf.ToByte("p")]=467
	me.cw[gopdf.ToByte("q")]=467
	me.cw[gopdf.ToByte("r")]=310
	me.cw[gopdf.ToByte("s")]=418
	me.cw[gopdf.ToByte("t")]=287
	me.cw[gopdf.ToByte("u")]=467
	me.cw[gopdf.ToByte("v")]=428
	me.cw[gopdf.ToByte("w")]=598
	me.cw[gopdf.ToByte("x")]=428
	me.cw[gopdf.ToByte("y")]=428
	me.cw[gopdf.ToByte("z")]=416
	me.cw[gopdf.ToByte("{")]=277
	me.cw[gopdf.ToByte("|")]=224
	me.cw[gopdf.ToByte("}")]=277
	me.cw[gopdf.ToByte("~")]=523
	me.cw[gopdf.Chr(127)]=213
	me.cw[gopdf.Chr(128)]=456
	me.cw[gopdf.Chr(129)]=213
	me.cw[gopdf.Chr(130)]=232
	me.cw[gopdf.Chr(131)]=311
	me.cw[gopdf.Chr(132)]=366
	me.cw[gopdf.Chr(133)]=634
	me.cw[gopdf.Chr(134)]=441
	me.cw[gopdf.Chr(135)]=481
	me.cw[gopdf.Chr(136)]=428
	me.cw[gopdf.Chr(137)]=780
	me.cw[gopdf.Chr(138)]=521
	me.cw[gopdf.Chr(139)]=267
	me.cw[gopdf.Chr(140)]=800
	me.cw[gopdf.Chr(141)]=213
	me.cw[gopdf.Chr(142)]=468
	me.cw[gopdf.Chr(143)]=213
	me.cw[gopdf.Chr(144)]=213
	me.cw[gopdf.Chr(145)]=215
	me.cw[gopdf.Chr(146)]=210
	me.cw[gopdf.Chr(147)]=367
	me.cw[gopdf.Chr(148)]=369
	me.cw[gopdf.Chr(149)]=322
	me.cw[gopdf.Chr(150)]=567
	me.cw[gopdf.Chr(151)]=668
	me.cw[gopdf.Chr(152)]=410
	me.cw[gopdf.Chr(153)]=523
	me.cw[gopdf.Chr(154)]=418
	me.cw[gopdf.Chr(155)]=256
	me.cw[gopdf.Chr(156)]=731
	me.cw[gopdf.Chr(157)]=213
	me.cw[gopdf.Chr(158)]=416
	me.cw[gopdf.Chr(159)]=513
	me.cw[gopdf.Chr(160)]=213
	me.cw[gopdf.Chr(161)]=256
	me.cw[gopdf.Chr(162)]=484
	me.cw[gopdf.Chr(163)]=495
	me.cw[gopdf.Chr(164)]=586
	me.cw[gopdf.Chr(165)]=508
	me.cw[gopdf.Chr(166)]=223
	me.cw[gopdf.Chr(167)]=523
	me.cw[gopdf.Chr(168)]=471
	me.cw[gopdf.Chr(169)]=667
	me.cw[gopdf.Chr(170)]=369
	me.cw[gopdf.Chr(171)]=426
	me.cw[gopdf.Chr(172)]=456
	me.cw[gopdf.Chr(173)]=354
	me.cw[gopdf.Chr(174)]=667
	me.cw[gopdf.Chr(175)]=448
	me.cw[gopdf.Chr(176)]=341
	me.cw[gopdf.Chr(177)]=464
	me.cw[gopdf.Chr(178)]=353
	me.cw[gopdf.Chr(179)]=349
	me.cw[gopdf.Chr(180)]=325
	me.cw[gopdf.Chr(181)]=524
	me.cw[gopdf.Chr(182)]=448
	me.cw[gopdf.Chr(183)]=270
	me.cw[gopdf.Chr(184)]=234
	me.cw[gopdf.Chr(185)]=245
	me.cw[gopdf.Chr(186)]=382
	me.cw[gopdf.Chr(187)]=426
	me.cw[gopdf.Chr(188)]=589
	me.cw[gopdf.Chr(189)]=619
	me.cw[gopdf.Chr(190)]=680
	me.cw[gopdf.Chr(191)]=429
	me.cw[gopdf.Chr(192)]=534
	me.cw[gopdf.Chr(193)]=534
	me.cw[gopdf.Chr(194)]=534
	me.cw[gopdf.Chr(195)]=534
	me.cw[gopdf.Chr(196)]=534
	me.cw[gopdf.Chr(197)]=534
	me.cw[gopdf.Chr(198)]=775
	me.cw[gopdf.Chr(199)]=516
	me.cw[gopdf.Chr(200)]=473
	me.cw[gopdf.Chr(201)]=473
	me.cw[gopdf.Chr(202)]=473
	me.cw[gopdf.Chr(203)]=473
	me.cw[gopdf.Chr(204)]=253
	me.cw[gopdf.Chr(205)]=253
	me.cw[gopdf.Chr(206)]=253
	me.cw[gopdf.Chr(207)]=253
	me.cw[gopdf.Chr(208)]=562
	me.cw[gopdf.Chr(209)]=581
	me.cw[gopdf.Chr(210)]=568
	me.cw[gopdf.Chr(211)]=568
	me.cw[gopdf.Chr(212)]=568
	me.cw[gopdf.Chr(213)]=568
	me.cw[gopdf.Chr(214)]=568
	me.cw[gopdf.Chr(215)]=458
	me.cw[gopdf.Chr(216)]=568
	me.cw[gopdf.Chr(217)]=551
	me.cw[gopdf.Chr(218)]=551
	me.cw[gopdf.Chr(219)]=551
	me.cw[gopdf.Chr(220)]=551
	me.cw[gopdf.Chr(221)]=513
	me.cw[gopdf.Chr(222)]=505
	me.cw[gopdf.Chr(223)]=531
	me.cw[gopdf.Chr(224)]=443
	me.cw[gopdf.Chr(225)]=443
	me.cw[gopdf.Chr(226)]=443
	me.cw[gopdf.Chr(227)]=443
	me.cw[gopdf.Chr(228)]=443
	me.cw[gopdf.Chr(229)]=443
	me.cw[gopdf.Chr(230)]=691
	me.cw[gopdf.Chr(231)]=434
	me.cw[gopdf.Chr(232)]=439
	me.cw[gopdf.Chr(233)]=439
	me.cw[gopdf.Chr(234)]=439
	me.cw[gopdf.Chr(235)]=439
	me.cw[gopdf.Chr(236)]=246
	me.cw[gopdf.Chr(237)]=246
	me.cw[gopdf.Chr(238)]=246
	me.cw[gopdf.Chr(239)]=246
	me.cw[gopdf.Chr(240)]=493
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
	me.cw[gopdf.Chr(253)]=428
	me.cw[gopdf.Chr(254)]=469
	me.cw[gopdf.Chr(255)]=428
	me.up = -73
	me.ut = 49
	me.fonttype = "TrueType"
	me.name = "RobotoCondensed-Bold"
	me.enc = "cp1252"
	me.desc = make([]gopdf.FontDescItem,8)
	me.desc[0] =  gopdf.FontDescItem{ Key:"Ascent",Val : "750" }
	me.desc[1] =  gopdf.FontDescItem{ Key: "Descent", Val : "-250" }
	me.desc[2] =  gopdf.FontDescItem{ Key:"CapHeight", Val :  "711" }
	me.desc[3] =  gopdf.FontDescItem{ Key: "Flags", Val :  "32" }
	me.desc[4] =  gopdf.FontDescItem{ Key:"FontBBox", Val :  "[-651 -271 992 1048]" }
	me.desc[5] =  gopdf.FontDescItem{ Key:"ItalicAngle", Val :  "0" }
	me.desc[6] =  gopdf.FontDescItem{ Key:"StemV", Val :  "120" }
 	me.desc[7] =  gopdf.FontDescItem{ Key:"MissingWidth", Val :  "213" } 
 }
func (me * RobotoCondensedBold)GetType() string{
	return me.fonttype
}
func (me * RobotoCondensedBold)GetName() string{
	return me.name
}	
func (me * RobotoCondensedBold)GetDesc() []gopdf.FontDescItem{
	return me.desc
}
func (me * RobotoCondensedBold)GetUp() int{
	return me.up
}
func (me * RobotoCondensedBold)GetUt()  int{
	return me.ut
}
func (me * RobotoCondensedBold)GetCw() gopdf.FontCw{
	return me.cw
}
func (me * RobotoCondensedBold)GetEnc() string{
	return me.enc
}
func (me * RobotoCondensedBold)GetDiff() string {
	return me.diff
}
func (me * RobotoCondensedBold) GetOriginalsize() int{
	return 98764
}
func (me * RobotoCondensedBold)  SetFamily(family string){
	me.family = family
}
func (me * RobotoCondensedBold) 	GetFamily() string{
	return me.family
}
