package fonts //change this
import (
	"github.com/signintech/gopdf"
)
type RobotoCondensed-Light struct {
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
func (me * RobotoCondensed-Light) Init(){
	me.cw = make(gopdf.FontCw)
	me.cw[gopdf.Chr(0)]=205
	me.cw[gopdf.Chr(1)]=205
	me.cw[gopdf.Chr(2)]=205
	me.cw[gopdf.Chr(3)]=205
	me.cw[gopdf.Chr(4)]=205
	me.cw[gopdf.Chr(5)]=205
	me.cw[gopdf.Chr(6)]=205
	me.cw[gopdf.Chr(7)]=205
	me.cw[gopdf.Chr(8)]=205
	me.cw[gopdf.Chr(9)]=205
	me.cw[gopdf.Chr(10)]=205
	me.cw[gopdf.Chr(11)]=205
	me.cw[gopdf.Chr(12)]=205
	me.cw[gopdf.Chr(13)]=205
	me.cw[gopdf.Chr(14)]=205
	me.cw[gopdf.Chr(15)]=205
	me.cw[gopdf.Chr(16)]=205
	me.cw[gopdf.Chr(17)]=205
	me.cw[gopdf.Chr(18)]=205
	me.cw[gopdf.Chr(19)]=205
	me.cw[gopdf.Chr(20)]=205
	me.cw[gopdf.Chr(21)]=205
	me.cw[gopdf.Chr(22)]=205
	me.cw[gopdf.Chr(23)]=205
	me.cw[gopdf.Chr(24)]=205
	me.cw[gopdf.Chr(25)]=205
	me.cw[gopdf.Chr(26)]=205
	me.cw[gopdf.Chr(27)]=205
	me.cw[gopdf.Chr(28)]=205
	me.cw[gopdf.Chr(29)]=205
	me.cw[gopdf.Chr(30)]=205
	me.cw[gopdf.Chr(31)]=205
	me.cw[gopdf.ToByte(" ")]=205
	me.cw[gopdf.ToByte("!")]=201
	me.cw[gopdf.ToByte("\"")]=260
	me.cw[gopdf.ToByte("#")]=499
	me.cw[gopdf.ToByte("$")]=450
	me.cw[gopdf.ToByte("%")]=618
	me.cw[gopdf.ToByte("&")]=521
	me.cw[gopdf.ToByte("'")]=154
	me.cw[gopdf.ToByte("(")]=264
	me.cw[gopdf.ToByte(")")]=268
	me.cw[gopdf.ToByte("*")]=364
	me.cw[gopdf.ToByte("+")]=478
	me.cw[gopdf.ToByte(",")]=165
	me.cw[gopdf.ToByte("-")]=233
	me.cw[gopdf.ToByte(".")]=209
	me.cw[gopdf.ToByte("/")]=330
	me.cw[gopdf.ToByte("0")]=450
	me.cw[gopdf.ToByte("1")]=450
	me.cw[gopdf.ToByte("2")]=450
	me.cw[gopdf.ToByte("3")]=450
	me.cw[gopdf.ToByte("4")]=450
	me.cw[gopdf.ToByte("5")]=450
	me.cw[gopdf.ToByte("6")]=450
	me.cw[gopdf.ToByte("7")]=450
	me.cw[gopdf.ToByte("8")]=450
	me.cw[gopdf.ToByte("9")]=450
	me.cw[gopdf.ToByte(":")]=192
	me.cw[gopdf.ToByte(";")]=194
	me.cw[gopdf.ToByte("<")]=412
	me.cw[gopdf.ToByte("=")]=452
	me.cw[gopdf.ToByte(">")]=420
	me.cw[gopdf.ToByte("?")]=375
	me.cw[gopdf.ToByte("@")]=758
	me.cw[gopdf.ToByte("A")]=517
	me.cw[gopdf.ToByte("B")]=504
	me.cw[gopdf.ToByte("C")]=501
	me.cw[gopdf.ToByte("D")]=538
	me.cw[gopdf.ToByte("E")]=466
	me.cw[gopdf.ToByte("F")]=466
	me.cw[gopdf.ToByte("G")]=542
	me.cw[gopdf.ToByte("H")]=568
	me.cw[gopdf.ToByte("I")]=229
	me.cw[gopdf.ToByte("J")]=444
	me.cw[gopdf.ToByte("K")]=515
	me.cw[gopdf.ToByte("L")]=435
	me.cw[gopdf.ToByte("M")]=671
	me.cw[gopdf.ToByte("N")]=570
	me.cw[gopdf.ToByte("O")]=549
	me.cw[gopdf.ToByte("P")]=504
	me.cw[gopdf.ToByte("Q")]=552
	me.cw[gopdf.ToByte("R")]=529
	me.cw[gopdf.ToByte("S")]=487
	me.cw[gopdf.ToByte("T")]=463
	me.cw[gopdf.ToByte("U")]=521
	me.cw[gopdf.ToByte("V")]=508
	me.cw[gopdf.ToByte("W")]=682
	me.cw[gopdf.ToByte("X")]=499
	me.cw[gopdf.ToByte("Y")]=476
	me.cw[gopdf.ToByte("Z")]=464
	me.cw[gopdf.ToByte("[")]=219
	me.cw[gopdf.ToByte("\\")]=328
	me.cw[gopdf.ToByte("]")]=219
	me.cw[gopdf.ToByte("^")]=339
	me.cw[gopdf.ToByte("_")]=358
	me.cw[gopdf.ToByte("`")]=255
	me.cw[gopdf.ToByte("a")]=438
	me.cw[gopdf.ToByte("b")]=453
	me.cw[gopdf.ToByte("c")]=422
	me.cw[gopdf.ToByte("d")]=453
	me.cw[gopdf.ToByte("e")]=421
	me.cw[gopdf.ToByte("f")]=276
	me.cw[gopdf.ToByte("g")]=452
	me.cw[gopdf.ToByte("h")]=453
	me.cw[gopdf.ToByte("i")]=199
	me.cw[gopdf.ToByte("j")]=205
	me.cw[gopdf.ToByte("k")]=404
	me.cw[gopdf.ToByte("l")]=199
	me.cw[gopdf.ToByte("m")]=703
	me.cw[gopdf.ToByte("n")]=453
	me.cw[gopdf.ToByte("o")]=453
	me.cw[gopdf.ToByte("p")]=453
	me.cw[gopdf.ToByte("q")]=453
	me.cw[gopdf.ToByte("r")]=283
	me.cw[gopdf.ToByte("s")]=405
	me.cw[gopdf.ToByte("t")]=262
	me.cw[gopdf.ToByte("u")]=453
	me.cw[gopdf.ToByte("v")]=400
	me.cw[gopdf.ToByte("w")]=605
	me.cw[gopdf.ToByte("x")]=400
	me.cw[gopdf.ToByte("y")]=400
	me.cw[gopdf.ToByte("z")]=387
	me.cw[gopdf.ToByte("{")]=274
	me.cw[gopdf.ToByte("|")]=195
	me.cw[gopdf.ToByte("}")]=274
	me.cw[gopdf.ToByte("~")]=547
	me.cw[gopdf.Chr(127)]=205
	me.cw[gopdf.Chr(128)]=425
	me.cw[gopdf.Chr(129)]=205
	me.cw[gopdf.Chr(130)]=158
	me.cw[gopdf.Chr(131)]=274
	me.cw[gopdf.Chr(132)]=260
	me.cw[gopdf.Chr(133)]=525
	me.cw[gopdf.Chr(134)]=444
	me.cw[gopdf.Chr(135)]=457
	me.cw[gopdf.Chr(136)]=365
	me.cw[gopdf.Chr(137)]=770
	me.cw[gopdf.Chr(138)]=487
	me.cw[gopdf.Chr(139)]=247
	me.cw[gopdf.Chr(140)]=751
	me.cw[gopdf.Chr(141)]=205
	me.cw[gopdf.Chr(142)]=464
	me.cw[gopdf.Chr(143)]=205
	me.cw[gopdf.Chr(144)]=205
	me.cw[gopdf.Chr(145)]=161
	me.cw[gopdf.Chr(146)]=161
	me.cw[gopdf.Chr(147)]=268
	me.cw[gopdf.Chr(148)]=270
	me.cw[gopdf.Chr(149)]=284
	me.cw[gopdf.Chr(150)]=554
	me.cw[gopdf.Chr(151)]=649
	me.cw[gopdf.Chr(152)]=367
	me.cw[gopdf.Chr(153)]=499
	me.cw[gopdf.Chr(154)]=405
	me.cw[gopdf.Chr(155)]=247
	me.cw[gopdf.Chr(156)]=730
	me.cw[gopdf.Chr(157)]=205
	me.cw[gopdf.Chr(158)]=387
	me.cw[gopdf.Chr(159)]=476
	me.cw[gopdf.Chr(160)]=205
	me.cw[gopdf.Chr(161)]=195
	me.cw[gopdf.Chr(162)]=440
	me.cw[gopdf.Chr(163)]=464
	me.cw[gopdf.Chr(164)]=604
	me.cw[gopdf.Chr(165)]=483
	me.cw[gopdf.Chr(166)]=191
	me.cw[gopdf.Chr(167)]=490
	me.cw[gopdf.Chr(168)]=406
	me.cw[gopdf.Chr(169)]=666
	me.cw[gopdf.Chr(170)]=361
	me.cw[gopdf.Chr(171)]=375
	me.cw[gopdf.Chr(172)]=443
	me.cw[gopdf.Chr(173)]=233
	me.cw[gopdf.Chr(174)]=667
	me.cw[gopdf.Chr(175)]=374
	me.cw[gopdf.Chr(176)]=322
	me.cw[gopdf.Chr(177)]=450
	me.cw[gopdf.Chr(178)]=337
	me.cw[gopdf.Chr(179)]=343
	me.cw[gopdf.Chr(180)]=255
	me.cw[gopdf.Chr(181)]=453
	me.cw[gopdf.Chr(182)]=390
	me.cw[gopdf.Chr(183)]=211
	me.cw[gopdf.Chr(184)]=205
	me.cw[gopdf.Chr(185)]=216
	me.cw[gopdf.Chr(186)]=368
	me.cw[gopdf.Chr(187)]=374
	me.cw[gopdf.Chr(188)]=615
	me.cw[gopdf.Chr(189)]=633
	me.cw[gopdf.Chr(190)]=678
	me.cw[gopdf.Chr(191)]=388
	me.cw[gopdf.Chr(192)]=517
	me.cw[gopdf.Chr(193)]=517
	me.cw[gopdf.Chr(194)]=517
	me.cw[gopdf.Chr(195)]=517
	me.cw[gopdf.Chr(196)]=517
	me.cw[gopdf.Chr(197)]=517
	me.cw[gopdf.Chr(198)]=734
	me.cw[gopdf.Chr(199)]=501
	me.cw[gopdf.Chr(200)]=466
	me.cw[gopdf.Chr(201)]=466
	me.cw[gopdf.Chr(202)]=466
	me.cw[gopdf.Chr(203)]=466
	me.cw[gopdf.Chr(204)]=229
	me.cw[gopdf.Chr(205)]=229
	me.cw[gopdf.Chr(206)]=229
	me.cw[gopdf.Chr(207)]=229
	me.cw[gopdf.Chr(208)]=552
	me.cw[gopdf.Chr(209)]=570
	me.cw[gopdf.Chr(210)]=549
	me.cw[gopdf.Chr(211)]=549
	me.cw[gopdf.Chr(212)]=549
	me.cw[gopdf.Chr(213)]=549
	me.cw[gopdf.Chr(214)]=549
	me.cw[gopdf.Chr(215)]=446
	me.cw[gopdf.Chr(216)]=543
	me.cw[gopdf.Chr(217)]=521
	me.cw[gopdf.Chr(218)]=521
	me.cw[gopdf.Chr(219)]=521
	me.cw[gopdf.Chr(220)]=521
	me.cw[gopdf.Chr(221)]=476
	me.cw[gopdf.Chr(222)]=476
	me.cw[gopdf.Chr(223)]=475
	me.cw[gopdf.Chr(224)]=438
	me.cw[gopdf.Chr(225)]=438
	me.cw[gopdf.Chr(226)]=438
	me.cw[gopdf.Chr(227)]=438
	me.cw[gopdf.Chr(228)]=438
	me.cw[gopdf.Chr(229)]=438
	me.cw[gopdf.Chr(230)]=674
	me.cw[gopdf.Chr(231)]=422
	me.cw[gopdf.Chr(232)]=421
	me.cw[gopdf.Chr(233)]=421
	me.cw[gopdf.Chr(234)]=421
	me.cw[gopdf.Chr(235)]=421
	me.cw[gopdf.Chr(236)]=196
	me.cw[gopdf.Chr(237)]=196
	me.cw[gopdf.Chr(238)]=196
	me.cw[gopdf.Chr(239)]=196
	me.cw[gopdf.Chr(240)]=470
	me.cw[gopdf.Chr(241)]=453
	me.cw[gopdf.Chr(242)]=453
	me.cw[gopdf.Chr(243)]=453
	me.cw[gopdf.Chr(244)]=453
	me.cw[gopdf.Chr(245)]=453
	me.cw[gopdf.Chr(246)]=453
	me.cw[gopdf.Chr(247)]=459
	me.cw[gopdf.Chr(248)]=453
	me.cw[gopdf.Chr(249)]=453
	me.cw[gopdf.Chr(250)]=453
	me.cw[gopdf.Chr(251)]=453
	me.cw[gopdf.Chr(252)]=453
	me.cw[gopdf.Chr(253)]=400
	me.cw[gopdf.Chr(254)]=458
	me.cw[gopdf.Chr(255)]=400
	me.up = -73
	me.ut = 49
	me.fonttype = "TrueType"
	me.name = "RobotoCondensed-Light"
	me.enc = "cp1252"
	me.desc = make([]gopdf.FontDescItem,8)
	me.desc[0] =  gopdf.FontDescItem{ Key:"Ascent",Val : "750" }
	me.desc[1] =  gopdf.FontDescItem{ Key: "Descent", Val : "-250" }
	me.desc[2] =  gopdf.FontDescItem{ Key:"CapHeight", Val :  "711" }
	me.desc[3] =  gopdf.FontDescItem{ Key: "Flags", Val :  "32" }
	me.desc[4] =  gopdf.FontDescItem{ Key:"FontBBox", Val :  "[-421 -271 931 1048]" }
	me.desc[5] =  gopdf.FontDescItem{ Key:"ItalicAngle", Val :  "0" }
	me.desc[6] =  gopdf.FontDescItem{ Key:"StemV", Val :  "70" }
 	me.desc[7] =  gopdf.FontDescItem{ Key:"MissingWidth", Val :  "205" } 
 }
func (me * RobotoCondensed-Light)GetType() string{
	return me.fonttype
}
func (me * RobotoCondensed-Light)GetName() string{
	return me.name
}	
func (me * RobotoCondensed-Light)GetDesc() []gopdf.FontDescItem{
	return me.desc
}
func (me * RobotoCondensed-Light)GetUp() int{
	return me.up
}
func (me * RobotoCondensed-Light)GetUt()  int{
	return me.ut
}
func (me * RobotoCondensed-Light)GetCw() gopdf.FontCw{
	return me.cw
}
func (me * RobotoCondensed-Light)GetEnc() string{
	return me.enc
}
func (me * RobotoCondensed-Light)GetDiff() string {
	return me.diff
}
func (me * RobotoCondensed-Light) GetOriginalsize() int{
	return 98764
}
func (me * RobotoCondensed-Light)  SetFamily(family string){
	me.family = family
}
func (me * RobotoCondensed-Light) 	GetFamily() string{
	return me.family
}
