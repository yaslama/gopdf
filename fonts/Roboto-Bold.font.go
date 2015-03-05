package fonts //change this
import (
	"github.com/yaslama/gopdf"
)
type Roboto-Bold struct {
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
func (me * Roboto-Bold) Init(){
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
	me.cw[gopdf.ToByte("!")]=272
	me.cw[gopdf.ToByte("\"")]=320
	me.cw[gopdf.ToByte("#")]=595
	me.cw[gopdf.ToByte("$")]=574
	me.cw[gopdf.ToByte("%")]=739
	me.cw[gopdf.ToByte("&")]=657
	me.cw[gopdf.ToByte("'")]=161
	me.cw[gopdf.ToByte("(")]=347
	me.cw[gopdf.ToByte(")")]=348
	me.cw[gopdf.ToByte("*")]=443
	me.cw[gopdf.ToByte("+")]=545
	me.cw[gopdf.ToByte(",")]=246
	me.cw[gopdf.ToByte("-")]=391
	me.cw[gopdf.ToByte(".")]=291
	me.cw[gopdf.ToByte("/")]=373
	me.cw[gopdf.ToByte("0")]=574
	me.cw[gopdf.ToByte("1")]=574
	me.cw[gopdf.ToByte("2")]=574
	me.cw[gopdf.ToByte("3")]=574
	me.cw[gopdf.ToByte("4")]=574
	me.cw[gopdf.ToByte("5")]=574
	me.cw[gopdf.ToByte("6")]=574
	me.cw[gopdf.ToByte("7")]=574
	me.cw[gopdf.ToByte("8")]=574
	me.cw[gopdf.ToByte("9")]=574
	me.cw[gopdf.ToByte(":")]=284
	me.cw[gopdf.ToByte(";")]=274
	me.cw[gopdf.ToByte("<")]=509
	me.cw[gopdf.ToByte("=")]=577
	me.cw[gopdf.ToByte(">")]=517
	me.cw[gopdf.ToByte("?")]=499
	me.cw[gopdf.ToByte("@")]=887
	me.cw[gopdf.ToByte("A")]=645
	me.cw[gopdf.ToByte("B")]=642
	me.cw[gopdf.ToByte("C")]=643
	me.cw[gopdf.ToByte("D")]=666
	me.cw[gopdf.ToByte("E")]=574
	me.cw[gopdf.ToByte("F")]=574
	me.cw[gopdf.ToByte("G")]=669
	me.cw[gopdf.ToByte("H")]=708
	me.cw[gopdf.ToByte("I")]=293
	me.cw[gopdf.ToByte("J")]=571
	me.cw[gopdf.ToByte("K")]=646
	me.cw[gopdf.ToByte("L")]=547
	me.cw[gopdf.ToByte("M")]=873
	me.cw[gopdf.ToByte("N")]=708
	me.cw[gopdf.ToByte("O")]=691
	me.cw[gopdf.ToByte("P")]=651
	me.cw[gopdf.ToByte("Q")]=691
	me.cw[gopdf.ToByte("R")]=654
	me.cw[gopdf.ToByte("S")]=625
	me.cw[gopdf.ToByte("T")]=635
	me.cw[gopdf.ToByte("U")]=687
	me.cw[gopdf.ToByte("V")]=641
	me.cw[gopdf.ToByte("W")]=873
	me.cw[gopdf.ToByte("X")]=636
	me.cw[gopdf.ToByte("Y")]=633
	me.cw[gopdf.ToByte("Z")]=589
	me.cw[gopdf.ToByte("[")]=278
	me.cw[gopdf.ToByte("\\")]=421
	me.cw[gopdf.ToByte("]")]=278
	me.cw[gopdf.ToByte("^")]=438
	me.cw[gopdf.ToByte("_")]=446
	me.cw[gopdf.ToByte("`")]=331
	me.cw[gopdf.ToByte("a")]=537
	me.cw[gopdf.ToByte("b")]=564
	me.cw[gopdf.ToByte("c")]=518
	me.cw[gopdf.ToByte("d")]=564
	me.cw[gopdf.ToByte("e")]=529
	me.cw[gopdf.ToByte("f")]=357
	me.cw[gopdf.ToByte("g")]=564
	me.cw[gopdf.ToByte("h")]=564
	me.cw[gopdf.ToByte("i")]=267
	me.cw[gopdf.ToByte("j")]=265
	me.cw[gopdf.ToByte("k")]=536
	me.cw[gopdf.ToByte("l")]=267
	me.cw[gopdf.ToByte("m")]=865
	me.cw[gopdf.ToByte("n")]=564
	me.cw[gopdf.ToByte("o")]=564
	me.cw[gopdf.ToByte("p")]=564
	me.cw[gopdf.ToByte("q")]=564
	me.cw[gopdf.ToByte("r")]=363
	me.cw[gopdf.ToByte("s")]=516
	me.cw[gopdf.ToByte("t")]=336
	me.cw[gopdf.ToByte("u")]=564
	me.cw[gopdf.ToByte("v")]=511
	me.cw[gopdf.ToByte("w")]=736
	me.cw[gopdf.ToByte("x")]=511
	me.cw[gopdf.ToByte("y")]=511
	me.cw[gopdf.ToByte("z")]=511
	me.cw[gopdf.ToByte("{")]=330
	me.cw[gopdf.ToByte("|")]=253
	me.cw[gopdf.ToByte("}")]=330
	me.cw[gopdf.ToByte("~")]=648
	me.cw[gopdf.Chr(127)]=249
	me.cw[gopdf.Chr(128)]=545
	me.cw[gopdf.Chr(129)]=249
	me.cw[gopdf.Chr(130)]=249
	me.cw[gopdf.Chr(131)]=361
	me.cw[gopdf.Chr(132)]=404
	me.cw[gopdf.Chr(133)]=742
	me.cw[gopdf.Chr(134)]=538
	me.cw[gopdf.Chr(135)]=579
	me.cw[gopdf.Chr(136)]=496
	me.cw[gopdf.Chr(137)]=960
	me.cw[gopdf.Chr(138)]=625
	me.cw[gopdf.Chr(139)]=312
	me.cw[gopdf.Chr(140)]=969
	me.cw[gopdf.Chr(141)]=249
	me.cw[gopdf.Chr(142)]=589
	me.cw[gopdf.Chr(143)]=249
	me.cw[gopdf.Chr(144)]=249
	me.cw[gopdf.Chr(145)]=234
	me.cw[gopdf.Chr(146)]=230
	me.cw[gopdf.Chr(147)]=407
	me.cw[gopdf.Chr(148)]=410
	me.cw[gopdf.Chr(149)]=360
	me.cw[gopdf.Chr(150)]=690
	me.cw[gopdf.Chr(151)]=815
	me.cw[gopdf.Chr(152)]=481
	me.cw[gopdf.Chr(153)]=632
	me.cw[gopdf.Chr(154)]=516
	me.cw[gopdf.Chr(155)]=302
	me.cw[gopdf.Chr(156)]=902
	me.cw[gopdf.Chr(157)]=249
	me.cw[gopdf.Chr(158)]=511
	me.cw[gopdf.Chr(159)]=633
	me.cw[gopdf.Chr(160)]=249
	me.cw[gopdf.Chr(161)]=283
	me.cw[gopdf.Chr(162)]=577
	me.cw[gopdf.Chr(163)]=594
	me.cw[gopdf.Chr(164)]=692
	me.cw[gopdf.Chr(165)]=612
	me.cw[gopdf.Chr(166)]=252
	me.cw[gopdf.Chr(167)]=628
	me.cw[gopdf.Chr(168)]=532
	me.cw[gopdf.Chr(169)]=784
	me.cw[gopdf.Chr(170)]=444
	me.cw[gopdf.Chr(171)]=500
	me.cw[gopdf.Chr(172)]=551
	me.cw[gopdf.Chr(173)]=391
	me.cw[gopdf.Chr(174)]=784
	me.cw[gopdf.Chr(175)]=502
	me.cw[gopdf.Chr(176)]=389
	me.cw[gopdf.Chr(177)]=538
	me.cw[gopdf.Chr(178)]=420
	me.cw[gopdf.Chr(179)]=418
	me.cw[gopdf.Chr(180)]=355
	me.cw[gopdf.Chr(181)]=617
	me.cw[gopdf.Chr(182)]=526
	me.cw[gopdf.Chr(183)]=302
	me.cw[gopdf.Chr(184)]=268
	me.cw[gopdf.Chr(185)]=280
	me.cw[gopdf.Chr(186)]=458
	me.cw[gopdf.Chr(187)]=500
	me.cw[gopdf.Chr(188)]=729
	me.cw[gopdf.Chr(189)]=758
	me.cw[gopdf.Chr(190)]=833
	me.cw[gopdf.Chr(191)]=506
	me.cw[gopdf.Chr(192)]=645
	me.cw[gopdf.Chr(193)]=645
	me.cw[gopdf.Chr(194)]=645
	me.cw[gopdf.Chr(195)]=645
	me.cw[gopdf.Chr(196)]=645
	me.cw[gopdf.Chr(197)]=645
	me.cw[gopdf.Chr(198)]=940
	me.cw[gopdf.Chr(199)]=643
	me.cw[gopdf.Chr(200)]=574
	me.cw[gopdf.Chr(201)]=574
	me.cw[gopdf.Chr(202)]=574
	me.cw[gopdf.Chr(203)]=574
	me.cw[gopdf.Chr(204)]=293
	me.cw[gopdf.Chr(205)]=293
	me.cw[gopdf.Chr(206)]=293
	me.cw[gopdf.Chr(207)]=293
	me.cw[gopdf.Chr(208)]=681
	me.cw[gopdf.Chr(209)]=708
	me.cw[gopdf.Chr(210)]=691
	me.cw[gopdf.Chr(211)]=691
	me.cw[gopdf.Chr(212)]=691
	me.cw[gopdf.Chr(213)]=691
	me.cw[gopdf.Chr(214)]=691
	me.cw[gopdf.Chr(215)]=531
	me.cw[gopdf.Chr(216)]=687
	me.cw[gopdf.Chr(217)]=687
	me.cw[gopdf.Chr(218)]=687
	me.cw[gopdf.Chr(219)]=687
	me.cw[gopdf.Chr(220)]=687
	me.cw[gopdf.Chr(221)]=633
	me.cw[gopdf.Chr(222)]=609
	me.cw[gopdf.Chr(223)]=632
	me.cw[gopdf.Chr(224)]=537
	me.cw[gopdf.Chr(225)]=537
	me.cw[gopdf.Chr(226)]=537
	me.cw[gopdf.Chr(227)]=537
	me.cw[gopdf.Chr(228)]=537
	me.cw[gopdf.Chr(229)]=537
	me.cw[gopdf.Chr(230)]=844
	me.cw[gopdf.Chr(231)]=518
	me.cw[gopdf.Chr(232)]=529
	me.cw[gopdf.Chr(233)]=529
	me.cw[gopdf.Chr(234)]=529
	me.cw[gopdf.Chr(235)]=529
	me.cw[gopdf.Chr(236)]=274
	me.cw[gopdf.Chr(237)]=274
	me.cw[gopdf.Chr(238)]=274
	me.cw[gopdf.Chr(239)]=274
	me.cw[gopdf.Chr(240)]=595
	me.cw[gopdf.Chr(241)]=564
	me.cw[gopdf.Chr(242)]=564
	me.cw[gopdf.Chr(243)]=564
	me.cw[gopdf.Chr(244)]=564
	me.cw[gopdf.Chr(245)]=564
	me.cw[gopdf.Chr(246)]=564
	me.cw[gopdf.Chr(247)]=570
	me.cw[gopdf.Chr(248)]=564
	me.cw[gopdf.Chr(249)]=564
	me.cw[gopdf.Chr(250)]=564
	me.cw[gopdf.Chr(251)]=564
	me.cw[gopdf.Chr(252)]=564
	me.cw[gopdf.Chr(253)]=511
	me.cw[gopdf.Chr(254)]=567
	me.cw[gopdf.Chr(255)]=511
	me.up = -73
	me.ut = 49
	me.fonttype = "TrueType"
	me.name = "Roboto-Bold"
	me.enc = "cp1252"
	me.desc = make([]gopdf.FontDescItem,8)
	me.desc[0] =  gopdf.FontDescItem{ Key:"Ascent",Val : "750" }
	me.desc[1] =  gopdf.FontDescItem{ Key: "Descent", Val : "-250" }
	me.desc[2] =  gopdf.FontDescItem{ Key:"CapHeight", Val :  "711" }
	me.desc[3] =  gopdf.FontDescItem{ Key: "Flags", Val :  "32" }
	me.desc[4] =  gopdf.FontDescItem{ Key:"FontBBox", Val :  "[-726 -271 1199 1048]" }
	me.desc[5] =  gopdf.FontDescItem{ Key:"ItalicAngle", Val :  "0" }
	me.desc[6] =  gopdf.FontDescItem{ Key:"StemV", Val :  "120" }
 	me.desc[7] =  gopdf.FontDescItem{ Key:"MissingWidth", Val :  "249" } 
 }
func (me * Roboto-Bold)GetType() string{
	return me.fonttype
}
func (me * Roboto-Bold)GetName() string{
	return me.name
}	
func (me * Roboto-Bold)GetDesc() []gopdf.FontDescItem{
	return me.desc
}
func (me * Roboto-Bold)GetUp() int{
	return me.up
}
func (me * Roboto-Bold)GetUt()  int{
	return me.ut
}
func (me * Roboto-Bold)GetCw() gopdf.FontCw{
	return me.cw
}
func (me * Roboto-Bold)GetEnc() string{
	return me.enc
}
func (me * Roboto-Bold)GetDiff() string {
	return me.diff
}
func (me * Roboto-Bold) GetOriginalsize() int{
	return 98764
}
func (me * Roboto-Bold)  SetFamily(family string){
	me.family = family
}
func (me * Roboto-Bold) 	GetFamily() string{
	return me.family
}
