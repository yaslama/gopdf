package fonts //change this
import (
	"github.com/signintech/gopdf"
)
type Roboto-Medium struct {
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
func (me * Roboto-Medium) Init(){
	me.cw = make(gopdf.FontCw)
	me.cw[gopdf.Chr(0)]=249
	me.cw[gopdf.Chr(1)]=249
	me.cw[gopdf.Chr(2)]=249
	me.cw[gopdf.Chr(3)]=249
	me.cw[gopdf.Chr(4)]=249
	me.cw[gopdf.Chr(5)]=249
	me.cw[gopdf.Chr(6)]=249
	me.cw[gopdf.Chr(7)]=249
	me.cw[gopdf.Chr(8)]=249
	me.cw[gopdf.Chr(9)]=249
	me.cw[gopdf.Chr(10)]=249
	me.cw[gopdf.Chr(11)]=249
	me.cw[gopdf.Chr(12)]=249
	me.cw[gopdf.Chr(13)]=249
	me.cw[gopdf.Chr(14)]=249
	me.cw[gopdf.Chr(15)]=249
	me.cw[gopdf.Chr(16)]=249
	me.cw[gopdf.Chr(17)]=249
	me.cw[gopdf.Chr(18)]=249
	me.cw[gopdf.Chr(19)]=249
	me.cw[gopdf.Chr(20)]=249
	me.cw[gopdf.Chr(21)]=249
	me.cw[gopdf.Chr(22)]=249
	me.cw[gopdf.Chr(23)]=249
	me.cw[gopdf.Chr(24)]=249
	me.cw[gopdf.Chr(25)]=249
	me.cw[gopdf.Chr(26)]=249
	me.cw[gopdf.Chr(27)]=249
	me.cw[gopdf.Chr(28)]=249
	me.cw[gopdf.Chr(29)]=249
	me.cw[gopdf.Chr(30)]=249
	me.cw[gopdf.Chr(31)]=249
	me.cw[gopdf.ToByte(" ")]=249
	me.cw[gopdf.ToByte("!")]=268
	me.cw[gopdf.ToByte("\"")]=324
	me.cw[gopdf.ToByte("#")]=610
	me.cw[gopdf.ToByte("$")]=568
	me.cw[gopdf.ToByte("%")]=734
	me.cw[gopdf.ToByte("&")]=639
	me.cw[gopdf.ToByte("'")]=169
	me.cw[gopdf.ToByte("(")]=339
	me.cw[gopdf.ToByte(")")]=342
	me.cw[gopdf.ToByte("*")]=437
	me.cw[gopdf.ToByte("+")]=557
	me.cw[gopdf.ToByte(",")]=220
	me.cw[gopdf.ToByte("-")]=328
	me.cw[gopdf.ToByte(".")]=279
	me.cw[gopdf.ToByte("/")]=396
	me.cw[gopdf.ToByte("0")]=568
	me.cw[gopdf.ToByte("1")]=568
	me.cw[gopdf.ToByte("2")]=568
	me.cw[gopdf.ToByte("3")]=568
	me.cw[gopdf.ToByte("4")]=568
	me.cw[gopdf.ToByte("5")]=568
	me.cw[gopdf.ToByte("6")]=568
	me.cw[gopdf.ToByte("7")]=568
	me.cw[gopdf.ToByte("8")]=568
	me.cw[gopdf.ToByte("9")]=568
	me.cw[gopdf.ToByte(":")]=268
	me.cw[gopdf.ToByte(";")]=266
	me.cw[gopdf.ToByte("<")]=508
	me.cw[gopdf.ToByte("=")]=569
	me.cw[gopdf.ToByte(">")]=521
	me.cw[gopdf.ToByte("?")]=486
	me.cw[gopdf.ToByte("@")]=891
	me.cw[gopdf.ToByte("A")]=657
	me.cw[gopdf.ToByte("B")]=641
	me.cw[gopdf.ToByte("C")]=641
	me.cw[gopdf.ToByte("D")]=673
	me.cw[gopdf.ToByte("E")]=580
	me.cw[gopdf.ToByte("F")]=579
	me.cw[gopdf.ToByte("G")]=678
	me.cw[gopdf.ToByte("H")]=711
	me.cw[gopdf.ToByte("I")]=288
	me.cw[gopdf.ToByte("J")]=561
	me.cw[gopdf.ToByte("K")]=645
	me.cw[gopdf.ToByte("L")]=549
	me.cw[gopdf.ToByte("M")]=876
	me.cw[gopdf.ToByte("N")]=711
	me.cw[gopdf.ToByte("O")]=695
	me.cw[gopdf.ToByte("P")]=646
	me.cw[gopdf.ToByte("Q")]=695
	me.cw[gopdf.ToByte("R")]=659
	me.cw[gopdf.ToByte("S")]=619
	me.cw[gopdf.ToByte("T")]=614
	me.cw[gopdf.ToByte("U")]=682
	me.cw[gopdf.ToByte("V")]=646
	me.cw[gopdf.ToByte("W")]=876
	me.cw[gopdf.ToByte("X")]=635
	me.cw[gopdf.ToByte("Y")]=626
	me.cw[gopdf.ToByte("Z")]=594
	me.cw[gopdf.ToByte("[")]=274
	me.cw[gopdf.ToByte("\\")]=417
	me.cw[gopdf.ToByte("]")]=274
	me.cw[gopdf.ToByte("^")]=427
	me.cw[gopdf.ToByte("_")]=451
	me.cw[gopdf.ToByte("`")]=322
	me.cw[gopdf.ToByte("a")]=544
	me.cw[gopdf.ToByte("b")]=566
	me.cw[gopdf.ToByte("c")]=525
	me.cw[gopdf.ToByte("d")]=566
	me.cw[gopdf.ToByte("e")]=529
	me.cw[gopdf.ToByte("f")]=351
	me.cw[gopdf.ToByte("g")]=566
	me.cw[gopdf.ToByte("h")]=566
	me.cw[gopdf.ToByte("i")]=259
	me.cw[gopdf.ToByte("j")]=262
	me.cw[gopdf.ToByte("k")]=523
	me.cw[gopdf.ToByte("l")]=259
	me.cw[gopdf.ToByte("m")]=870
	me.cw[gopdf.ToByte("n")]=566
	me.cw[gopdf.ToByte("o")]=566
	me.cw[gopdf.ToByte("p")]=566
	me.cw[gopdf.ToByte("q")]=566
	me.cw[gopdf.ToByte("r")]=356
	me.cw[gopdf.ToByte("s")]=520
	me.cw[gopdf.ToByte("t")]=327
	me.cw[gopdf.ToByte("u")]=566
	me.cw[gopdf.ToByte("v")]=507
	me.cw[gopdf.ToByte("w")]=747
	me.cw[gopdf.ToByte("x")]=507
	me.cw[gopdf.ToByte("y")]=507
	me.cw[gopdf.ToByte("z")]=507
	me.cw[gopdf.ToByte("{")]=335
	me.cw[gopdf.ToByte("|")]=251
	me.cw[gopdf.ToByte("}")]=335
	me.cw[gopdf.ToByte("~")]=665
	me.cw[gopdf.Chr(127)]=249
	me.cw[gopdf.Chr(128)]=538
	me.cw[gopdf.Chr(129)]=249
	me.cw[gopdf.Chr(130)]=226
	me.cw[gopdf.Chr(131)]=352
	me.cw[gopdf.Chr(132)]=379
	me.cw[gopdf.Chr(133)]=706
	me.cw[gopdf.Chr(134)]=545
	me.cw[gopdf.Chr(135)]=575
	me.cw[gopdf.Chr(136)]=487
	me.cw[gopdf.Chr(137)]=957
	me.cw[gopdf.Chr(138)]=619
	me.cw[gopdf.Chr(139)]=306
	me.cw[gopdf.Chr(140)]=963
	me.cw[gopdf.Chr(141)]=249
	me.cw[gopdf.Chr(142)]=594
	me.cw[gopdf.Chr(143)]=249
	me.cw[gopdf.Chr(144)]=249
	me.cw[gopdf.Chr(145)]=219
	me.cw[gopdf.Chr(146)]=217
	me.cw[gopdf.Chr(147)]=385
	me.cw[gopdf.Chr(148)]=388
	me.cw[gopdf.Chr(149)]=349
	me.cw[gopdf.Chr(150)]=691
	me.cw[gopdf.Chr(151)]=813
	me.cw[gopdf.Chr(152)]=480
	me.cw[gopdf.Chr(153)]=629
	me.cw[gopdf.Chr(154)]=520
	me.cw[gopdf.Chr(155)]=301
	me.cw[gopdf.Chr(156)]=903
	me.cw[gopdf.Chr(157)]=249
	me.cw[gopdf.Chr(158)]=507
	me.cw[gopdf.Chr(159)]=626
	me.cw[gopdf.Chr(160)]=249
	me.cw[gopdf.Chr(161)]=265
	me.cw[gopdf.Chr(162)]=561
	me.cw[gopdf.Chr(163)]=588
	me.cw[gopdf.Chr(164)]=702
	me.cw[gopdf.Chr(165)]=609
	me.cw[gopdf.Chr(166)]=248
	me.cw[gopdf.Chr(167)]=621
	me.cw[gopdf.Chr(168)]=515
	me.cw[gopdf.Chr(169)]=783
	me.cw[gopdf.Chr(170)]=446
	me.cw[gopdf.Chr(171)]=485
	me.cw[gopdf.Chr(172)]=553
	me.cw[gopdf.Chr(173)]=328
	me.cw[gopdf.Chr(174)]=783
	me.cw[gopdf.Chr(175)]=482
	me.cw[gopdf.Chr(176)]=380
	me.cw[gopdf.Chr(177)]=537
	me.cw[gopdf.Chr(178)]=422
	me.cw[gopdf.Chr(179)]=423
	me.cw[gopdf.Chr(180)]=336
	me.cw[gopdf.Chr(181)]=591
	me.cw[gopdf.Chr(182)]=508
	me.cw[gopdf.Chr(183)]=282
	me.cw[gopdf.Chr(184)]=258
	me.cw[gopdf.Chr(185)]=276
	me.cw[gopdf.Chr(186)]=457
	me.cw[gopdf.Chr(187)]=485
	me.cw[gopdf.Chr(188)]=756
	me.cw[gopdf.Chr(189)]=800
	me.cw[gopdf.Chr(190)]=853
	me.cw[gopdf.Chr(191)]=500
	me.cw[gopdf.Chr(192)]=657
	me.cw[gopdf.Chr(193)]=657
	me.cw[gopdf.Chr(194)]=657
	me.cw[gopdf.Chr(195)]=657
	me.cw[gopdf.Chr(196)]=657
	me.cw[gopdf.Chr(197)]=657
	me.cw[gopdf.Chr(198)]=940
	me.cw[gopdf.Chr(199)]=641
	me.cw[gopdf.Chr(200)]=580
	me.cw[gopdf.Chr(201)]=580
	me.cw[gopdf.Chr(202)]=580
	me.cw[gopdf.Chr(203)]=580
	me.cw[gopdf.Chr(204)]=288
	me.cw[gopdf.Chr(205)]=288
	me.cw[gopdf.Chr(206)]=288
	me.cw[gopdf.Chr(207)]=288
	me.cw[gopdf.Chr(208)]=688
	me.cw[gopdf.Chr(209)]=711
	me.cw[gopdf.Chr(210)]=695
	me.cw[gopdf.Chr(211)]=695
	me.cw[gopdf.Chr(212)]=695
	me.cw[gopdf.Chr(213)]=695
	me.cw[gopdf.Chr(214)]=695
	me.cw[gopdf.Chr(215)]=533
	me.cw[gopdf.Chr(216)]=685
	me.cw[gopdf.Chr(217)]=682
	me.cw[gopdf.Chr(218)]=682
	me.cw[gopdf.Chr(219)]=682
	me.cw[gopdf.Chr(220)]=682
	me.cw[gopdf.Chr(221)]=626
	me.cw[gopdf.Chr(222)]=599
	me.cw[gopdf.Chr(223)]=613
	me.cw[gopdf.Chr(224)]=544
	me.cw[gopdf.Chr(225)]=544
	me.cw[gopdf.Chr(226)]=544
	me.cw[gopdf.Chr(227)]=544
	me.cw[gopdf.Chr(228)]=544
	me.cw[gopdf.Chr(229)]=544
	me.cw[gopdf.Chr(230)]=844
	me.cw[gopdf.Chr(231)]=525
	me.cw[gopdf.Chr(232)]=529
	me.cw[gopdf.Chr(233)]=529
	me.cw[gopdf.Chr(234)]=529
	me.cw[gopdf.Chr(235)]=529
	me.cw[gopdf.Chr(236)]=263
	me.cw[gopdf.Chr(237)]=263
	me.cw[gopdf.Chr(238)]=263
	me.cw[gopdf.Chr(239)]=263
	me.cw[gopdf.Chr(240)]=591
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
	me.cw[gopdf.Chr(253)]=507
	me.cw[gopdf.Chr(254)]=574
	me.cw[gopdf.Chr(255)]=507
	me.up = -73
	me.ut = 49
	me.fonttype = "TrueType"
	me.name = "Roboto-Medium"
	me.enc = "cp1252"
	me.desc = make([]gopdf.FontDescItem,8)
	me.desc[0] =  gopdf.FontDescItem{ Key:"Ascent",Val : "750" }
	me.desc[1] =  gopdf.FontDescItem{ Key: "Descent", Val : "-250" }
	me.desc[2] =  gopdf.FontDescItem{ Key:"CapHeight", Val :  "711" }
	me.desc[3] =  gopdf.FontDescItem{ Key: "Flags", Val :  "32" }
	me.desc[4] =  gopdf.FontDescItem{ Key:"FontBBox", Val :  "[-732 -271 1193 1048]" }
	me.desc[5] =  gopdf.FontDescItem{ Key:"ItalicAngle", Val :  "0" }
	me.desc[6] =  gopdf.FontDescItem{ Key:"StemV", Val :  "70" }
 	me.desc[7] =  gopdf.FontDescItem{ Key:"MissingWidth", Val :  "249" } 
 }
func (me * Roboto-Medium)GetType() string{
	return me.fonttype
}
func (me * Roboto-Medium)GetName() string{
	return me.name
}	
func (me * Roboto-Medium)GetDesc() []gopdf.FontDescItem{
	return me.desc
}
func (me * Roboto-Medium)GetUp() int{
	return me.up
}
func (me * Roboto-Medium)GetUt()  int{
	return me.ut
}
func (me * Roboto-Medium)GetCw() gopdf.FontCw{
	return me.cw
}
func (me * Roboto-Medium)GetEnc() string{
	return me.enc
}
func (me * Roboto-Medium)GetDiff() string {
	return me.diff
}
func (me * Roboto-Medium) GetOriginalsize() int{
	return 98764
}
func (me * Roboto-Medium)  SetFamily(family string){
	me.family = family
}
func (me * Roboto-Medium) 	GetFamily() string{
	return me.family
}
