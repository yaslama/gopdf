package fonts //change this
import (
	"github.com/yaslama/gopdf"
)
type Roboto-Black struct {
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
func (me * Roboto-Black) Init(){
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
	me.cw[gopdf.ToByte("!")]=276
	me.cw[gopdf.ToByte("\"")]=317
	me.cw[gopdf.ToByte("#")]=581
	me.cw[gopdf.ToByte("$")]=579
	me.cw[gopdf.ToByte("%")]=743
	me.cw[gopdf.ToByte("&")]=674
	me.cw[gopdf.ToByte("'")]=154
	me.cw[gopdf.ToByte("(")]=355
	me.cw[gopdf.ToByte(")")]=354
	me.cw[gopdf.ToByte("*")]=449
	me.cw[gopdf.ToByte("+")]=534
	me.cw[gopdf.ToByte(",")]=270
	me.cw[gopdf.ToByte("-")]=450
	me.cw[gopdf.ToByte(".")]=302
	me.cw[gopdf.ToByte("/")]=351
	me.cw[gopdf.ToByte("0")]=579
	me.cw[gopdf.ToByte("1")]=579
	me.cw[gopdf.ToByte("2")]=579
	me.cw[gopdf.ToByte("3")]=579
	me.cw[gopdf.ToByte("4")]=579
	me.cw[gopdf.ToByte("5")]=579
	me.cw[gopdf.ToByte("6")]=579
	me.cw[gopdf.ToByte("7")]=579
	me.cw[gopdf.ToByte("8")]=579
	me.cw[gopdf.ToByte("9")]=579
	me.cw[gopdf.ToByte(":")]=299
	me.cw[gopdf.ToByte(";")]=283
	me.cw[gopdf.ToByte("<")]=510
	me.cw[gopdf.ToByte("=")]=583
	me.cw[gopdf.ToByte(">")]=513
	me.cw[gopdf.ToByte("?")]=509
	me.cw[gopdf.ToByte("@")]=883
	me.cw[gopdf.ToByte("A")]=634
	me.cw[gopdf.ToByte("B")]=642
	me.cw[gopdf.ToByte("C")]=645
	me.cw[gopdf.ToByte("D")]=659
	me.cw[gopdf.ToByte("E")]=569
	me.cw[gopdf.ToByte("F")]=569
	me.cw[gopdf.ToByte("G")]=661
	me.cw[gopdf.ToByte("H")]=705
	me.cw[gopdf.ToByte("I")]=298
	me.cw[gopdf.ToByte("J")]=580
	me.cw[gopdf.ToByte("K")]=647
	me.cw[gopdf.ToByte("L")]=545
	me.cw[gopdf.ToByte("M")]=869
	me.cw[gopdf.ToByte("N")]=705
	me.cw[gopdf.ToByte("O")]=687
	me.cw[gopdf.ToByte("P")]=656
	me.cw[gopdf.ToByte("Q")]=687
	me.cw[gopdf.ToByte("R")]=649
	me.cw[gopdf.ToByte("S")]=631
	me.cw[gopdf.ToByte("T")]=654
	me.cw[gopdf.ToByte("U")]=692
	me.cw[gopdf.ToByte("V")]=636
	me.cw[gopdf.ToByte("W")]=869
	me.cw[gopdf.ToByte("X")]=638
	me.cw[gopdf.ToByte("Y")]=640
	me.cw[gopdf.ToByte("Z")]=584
	me.cw[gopdf.ToByte("[")]=282
	me.cw[gopdf.ToByte("\\")]=425
	me.cw[gopdf.ToByte("]")]=282
	me.cw[gopdf.ToByte("^")]=447
	me.cw[gopdf.ToByte("_")]=441
	me.cw[gopdf.ToByte("`")]=339
	me.cw[gopdf.ToByte("a")]=530
	me.cw[gopdf.ToByte("b")]=562
	me.cw[gopdf.ToByte("c")]=511
	me.cw[gopdf.ToByte("d")]=562
	me.cw[gopdf.ToByte("e")]=529
	me.cw[gopdf.ToByte("f")]=363
	me.cw[gopdf.ToByte("g")]=562
	me.cw[gopdf.ToByte("h")]=562
	me.cw[gopdf.ToByte("i")]=274
	me.cw[gopdf.ToByte("j")]=268
	me.cw[gopdf.ToByte("k")]=546
	me.cw[gopdf.ToByte("l")]=274
	me.cw[gopdf.ToByte("m")]=861
	me.cw[gopdf.ToByte("n")]=562
	me.cw[gopdf.ToByte("o")]=562
	me.cw[gopdf.ToByte("p")]=562
	me.cw[gopdf.ToByte("q")]=562
	me.cw[gopdf.ToByte("r")]=370
	me.cw[gopdf.ToByte("s")]=512
	me.cw[gopdf.ToByte("t")]=344
	me.cw[gopdf.ToByte("u")]=562
	me.cw[gopdf.ToByte("v")]=514
	me.cw[gopdf.ToByte("w")]=725
	me.cw[gopdf.ToByte("x")]=514
	me.cw[gopdf.ToByte("y")]=514
	me.cw[gopdf.ToByte("z")]=514
	me.cw[gopdf.ToByte("{")]=325
	me.cw[gopdf.ToByte("|")]=255
	me.cw[gopdf.ToByte("}")]=325
	me.cw[gopdf.ToByte("~")]=632
	me.cw[gopdf.Chr(127)]=249
	me.cw[gopdf.Chr(128)]=552
	me.cw[gopdf.Chr(129)]=249
	me.cw[gopdf.Chr(130)]=271
	me.cw[gopdf.Chr(131)]=370
	me.cw[gopdf.Chr(132)]=428
	me.cw[gopdf.Chr(133)]=775
	me.cw[gopdf.Chr(134)]=531
	me.cw[gopdf.Chr(135)]=583
	me.cw[gopdf.Chr(136)]=504
	me.cw[gopdf.Chr(137)]=963
	me.cw[gopdf.Chr(138)]=631
	me.cw[gopdf.Chr(139)]=317
	me.cw[gopdf.Chr(140)]=974
	me.cw[gopdf.Chr(141)]=249
	me.cw[gopdf.Chr(142)]=584
	me.cw[gopdf.Chr(143)]=249
	me.cw[gopdf.Chr(144)]=249
	me.cw[gopdf.Chr(145)]=250
	me.cw[gopdf.Chr(146)]=243
	me.cw[gopdf.Chr(147)]=427
	me.cw[gopdf.Chr(148)]=430
	me.cw[gopdf.Chr(149)]=370
	me.cw[gopdf.Chr(150)]=689
	me.cw[gopdf.Chr(151)]=818
	me.cw[gopdf.Chr(152)]=482
	me.cw[gopdf.Chr(153)]=634
	me.cw[gopdf.Chr(154)]=512
	me.cw[gopdf.Chr(155)]=303
	me.cw[gopdf.Chr(156)]=901
	me.cw[gopdf.Chr(157)]=249
	me.cw[gopdf.Chr(158)]=514
	me.cw[gopdf.Chr(159)]=640
	me.cw[gopdf.Chr(160)]=249
	me.cw[gopdf.Chr(161)]=300
	me.cw[gopdf.Chr(162)]=591
	me.cw[gopdf.Chr(163)]=599
	me.cw[gopdf.Chr(164)]=682
	me.cw[gopdf.Chr(165)]=614
	me.cw[gopdf.Chr(166)]=256
	me.cw[gopdf.Chr(167)]=635
	me.cw[gopdf.Chr(168)]=548
	me.cw[gopdf.Chr(169)]=785
	me.cw[gopdf.Chr(170)]=441
	me.cw[gopdf.Chr(171)]=515
	me.cw[gopdf.Chr(172)]=549
	me.cw[gopdf.Chr(173)]=450
	me.cw[gopdf.Chr(174)]=785
	me.cw[gopdf.Chr(175)]=520
	me.cw[gopdf.Chr(176)]=397
	me.cw[gopdf.Chr(177)]=539
	me.cw[gopdf.Chr(178)]=418
	me.cw[gopdf.Chr(179)]=413
	me.cw[gopdf.Chr(180)]=372
	me.cw[gopdf.Chr(181)]=641
	me.cw[gopdf.Chr(182)]=543
	me.cw[gopdf.Chr(183)]=321
	me.cw[gopdf.Chr(184)]=278
	me.cw[gopdf.Chr(185)]=283
	me.cw[gopdf.Chr(186)]=458
	me.cw[gopdf.Chr(187)]=515
	me.cw[gopdf.Chr(188)]=704
	me.cw[gopdf.Chr(189)]=719
	me.cw[gopdf.Chr(190)]=813
	me.cw[gopdf.Chr(191)]=511
	me.cw[gopdf.Chr(192)]=634
	me.cw[gopdf.Chr(193)]=634
	me.cw[gopdf.Chr(194)]=634
	me.cw[gopdf.Chr(195)]=634
	me.cw[gopdf.Chr(196)]=634
	me.cw[gopdf.Chr(197)]=634
	me.cw[gopdf.Chr(198)]=940
	me.cw[gopdf.Chr(199)]=645
	me.cw[gopdf.Chr(200)]=569
	me.cw[gopdf.Chr(201)]=569
	me.cw[gopdf.Chr(202)]=569
	me.cw[gopdf.Chr(203)]=569
	me.cw[gopdf.Chr(204)]=298
	me.cw[gopdf.Chr(205)]=298
	me.cw[gopdf.Chr(206)]=298
	me.cw[gopdf.Chr(207)]=298
	me.cw[gopdf.Chr(208)]=674
	me.cw[gopdf.Chr(209)]=705
	me.cw[gopdf.Chr(210)]=687
	me.cw[gopdf.Chr(211)]=687
	me.cw[gopdf.Chr(212)]=687
	me.cw[gopdf.Chr(213)]=687
	me.cw[gopdf.Chr(214)]=687
	me.cw[gopdf.Chr(215)]=529
	me.cw[gopdf.Chr(216)]=689
	me.cw[gopdf.Chr(217)]=692
	me.cw[gopdf.Chr(218)]=692
	me.cw[gopdf.Chr(219)]=692
	me.cw[gopdf.Chr(220)]=692
	me.cw[gopdf.Chr(221)]=640
	me.cw[gopdf.Chr(222)]=618
	me.cw[gopdf.Chr(223)]=649
	me.cw[gopdf.Chr(224)]=530
	me.cw[gopdf.Chr(225)]=530
	me.cw[gopdf.Chr(226)]=530
	me.cw[gopdf.Chr(227)]=530
	me.cw[gopdf.Chr(228)]=530
	me.cw[gopdf.Chr(229)]=530
	me.cw[gopdf.Chr(230)]=845
	me.cw[gopdf.Chr(231)]=511
	me.cw[gopdf.Chr(232)]=529
	me.cw[gopdf.Chr(233)]=529
	me.cw[gopdf.Chr(234)]=529
	me.cw[gopdf.Chr(235)]=529
	me.cw[gopdf.Chr(236)]=285
	me.cw[gopdf.Chr(237)]=285
	me.cw[gopdf.Chr(238)]=285
	me.cw[gopdf.Chr(239)]=285
	me.cw[gopdf.Chr(240)]=599
	me.cw[gopdf.Chr(241)]=562
	me.cw[gopdf.Chr(242)]=562
	me.cw[gopdf.Chr(243)]=562
	me.cw[gopdf.Chr(244)]=562
	me.cw[gopdf.Chr(245)]=562
	me.cw[gopdf.Chr(246)]=562
	me.cw[gopdf.Chr(247)]=569
	me.cw[gopdf.Chr(248)]=563
	me.cw[gopdf.Chr(249)]=562
	me.cw[gopdf.Chr(250)]=562
	me.cw[gopdf.Chr(251)]=562
	me.cw[gopdf.Chr(252)]=562
	me.cw[gopdf.Chr(253)]=514
	me.cw[gopdf.Chr(254)]=561
	me.cw[gopdf.Chr(255)]=514
	me.up = -73
	me.ut = 49
	me.fonttype = "TrueType"
	me.name = "Roboto-Black"
	me.enc = "cp1252"
	me.desc = make([]gopdf.FontDescItem,8)
	me.desc[0] =  gopdf.FontDescItem{ Key:"Ascent",Val : "750" }
	me.desc[1] =  gopdf.FontDescItem{ Key: "Descent", Val : "-250" }
	me.desc[2] =  gopdf.FontDescItem{ Key:"CapHeight", Val :  "711" }
	me.desc[3] =  gopdf.FontDescItem{ Key: "Flags", Val :  "32" }
	me.desc[4] =  gopdf.FontDescItem{ Key:"FontBBox", Val :  "[-721 -271 1205 1064]" }
	me.desc[5] =  gopdf.FontDescItem{ Key:"ItalicAngle", Val :  "0" }
	me.desc[6] =  gopdf.FontDescItem{ Key:"StemV", Val :  "120" }
 	me.desc[7] =  gopdf.FontDescItem{ Key:"MissingWidth", Val :  "249" } 
 }
func (me * Roboto-Black)GetType() string{
	return me.fonttype
}
func (me * Roboto-Black)GetName() string{
	return me.name
}	
func (me * Roboto-Black)GetDesc() []gopdf.FontDescItem{
	return me.desc
}
func (me * Roboto-Black)GetUp() int{
	return me.up
}
func (me * Roboto-Black)GetUt()  int{
	return me.ut
}
func (me * Roboto-Black)GetCw() gopdf.FontCw{
	return me.cw
}
func (me * Roboto-Black)GetEnc() string{
	return me.enc
}
func (me * Roboto-Black)GetDiff() string {
	return me.diff
}
func (me * Roboto-Black) GetOriginalsize() int{
	return 98764
}
func (me * Roboto-Black)  SetFamily(family string){
	me.family = family
}
func (me * Roboto-Black) 	GetFamily() string{
	return me.family
}
