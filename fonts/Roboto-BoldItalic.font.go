package fonts //change this
import (
	"github.com/yaslama/gopdf"
)
type Roboto-BoldItalic struct {
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
func (me * Roboto-BoldItalic) Init(){
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
	me.cw[gopdf.ToByte("!")]=263
	me.cw[gopdf.ToByte("\"")]=308
	me.cw[gopdf.ToByte("#")]=557
	me.cw[gopdf.ToByte("$")]=541
	me.cw[gopdf.ToByte("%")]=691
	me.cw[gopdf.ToByte("&")]=620
	me.cw[gopdf.ToByte("'")]=154
	me.cw[gopdf.ToByte("(")]=332
	me.cw[gopdf.ToByte(")")]=332
	me.cw[gopdf.ToByte("*")]=419
	me.cw[gopdf.ToByte("+")]=511
	me.cw[gopdf.ToByte(",")]=237
	me.cw[gopdf.ToByte("-")]=375
	me.cw[gopdf.ToByte(".")]=281
	me.cw[gopdf.ToByte("/")]=351
	me.cw[gopdf.ToByte("0")]=541
	me.cw[gopdf.ToByte("1")]=541
	me.cw[gopdf.ToByte("2")]=541
	me.cw[gopdf.ToByte("3")]=541
	me.cw[gopdf.ToByte("4")]=541
	me.cw[gopdf.ToByte("5")]=541
	me.cw[gopdf.ToByte("6")]=541
	me.cw[gopdf.ToByte("7")]=541
	me.cw[gopdf.ToByte("8")]=541
	me.cw[gopdf.ToByte("9")]=541
	me.cw[gopdf.ToByte(":")]=277
	me.cw[gopdf.ToByte(";")]=267
	me.cw[gopdf.ToByte("<")]=478
	me.cw[gopdf.ToByte("=")]=542
	me.cw[gopdf.ToByte(">")]=485
	me.cw[gopdf.ToByte("?")]=473
	me.cw[gopdf.ToByte("@")]=827
	me.cw[gopdf.ToByte("A")]=608
	me.cw[gopdf.ToByte("B")]=604
	me.cw[gopdf.ToByte("C")]=599
	me.cw[gopdf.ToByte("D")]=614
	me.cw[gopdf.ToByte("E")]=539
	me.cw[gopdf.ToByte("F")]=539
	me.cw[gopdf.ToByte("G")]=626
	me.cw[gopdf.ToByte("H")]=664
	me.cw[gopdf.ToByte("I")]=280
	me.cw[gopdf.ToByte("J")]=537
	me.cw[gopdf.ToByte("K")]=578
	me.cw[gopdf.ToByte("L")]=516
	me.cw[gopdf.ToByte("M")]=817
	me.cw[gopdf.ToByte("N")]=664
	me.cw[gopdf.ToByte("O")]=642
	me.cw[gopdf.ToByte("P")]=614
	me.cw[gopdf.ToByte("Q")]=649
	me.cw[gopdf.ToByte("R")]=613
	me.cw[gopdf.ToByte("S")]=578
	me.cw[gopdf.ToByte("T")]=598
	me.cw[gopdf.ToByte("U")]=645
	me.cw[gopdf.ToByte("V")]=604
	me.cw[gopdf.ToByte("W")]=814
	me.cw[gopdf.ToByte("X")]=599
	me.cw[gopdf.ToByte("Y")]=596
	me.cw[gopdf.ToByte("Z")]=531
	me.cw[gopdf.ToByte("[")]=268
	me.cw[gopdf.ToByte("\\")]=400
	me.cw[gopdf.ToByte("]")]=268
	me.cw[gopdf.ToByte("^")]=414
	me.cw[gopdf.ToByte("_")]=422
	me.cw[gopdf.ToByte("`")]=317
	me.cw[gopdf.ToByte("a")]=505
	me.cw[gopdf.ToByte("b")]=531
	me.cw[gopdf.ToByte("c")]=487
	me.cw[gopdf.ToByte("d")]=531
	me.cw[gopdf.ToByte("e")]=499
	me.cw[gopdf.ToByte("f")]=340
	me.cw[gopdf.ToByte("g")]=531
	me.cw[gopdf.ToByte("h")]=531
	me.cw[gopdf.ToByte("i")]=257
	me.cw[gopdf.ToByte("j")]=255
	me.cw[gopdf.ToByte("k")]=507
	me.cw[gopdf.ToByte("l")]=257
	me.cw[gopdf.ToByte("m")]=807
	me.cw[gopdf.ToByte("n")]=531
	me.cw[gopdf.ToByte("o")]=531
	me.cw[gopdf.ToByte("p")]=531
	me.cw[gopdf.ToByte("q")]=531
	me.cw[gopdf.ToByte("r")]=345
	me.cw[gopdf.ToByte("s")]=486
	me.cw[gopdf.ToByte("t")]=319
	me.cw[gopdf.ToByte("u")]=531
	me.cw[gopdf.ToByte("v")]=482
	me.cw[gopdf.ToByte("w")]=688
	me.cw[gopdf.ToByte("x")]=482
	me.cw[gopdf.ToByte("y")]=482
	me.cw[gopdf.ToByte("z")]=482
	me.cw[gopdf.ToByte("{")]=313
	me.cw[gopdf.ToByte("|")]=245
	me.cw[gopdf.ToByte("}")]=313
	me.cw[gopdf.ToByte("~")]=604
	me.cw[gopdf.Chr(127)]=237
	me.cw[gopdf.Chr(128)]=514
	me.cw[gopdf.Chr(129)]=237
	me.cw[gopdf.Chr(130)]=244
	me.cw[gopdf.Chr(131)]=344
	me.cw[gopdf.Chr(132)]=392
	me.cw[gopdf.Chr(133)]=704
	me.cw[gopdf.Chr(134)]=504
	me.cw[gopdf.Chr(135)]=545
	me.cw[gopdf.Chr(136)]=474
	me.cw[gopdf.Chr(137)]=896
	me.cw[gopdf.Chr(138)]=588
	me.cw[gopdf.Chr(139)]=296
	me.cw[gopdf.Chr(140)]=910
	me.cw[gopdf.Chr(141)]=237
	me.cw[gopdf.Chr(142)]=552
	me.cw[gopdf.Chr(143)]=237
	me.cw[gopdf.Chr(144)]=237
	me.cw[gopdf.Chr(145)]=229
	me.cw[gopdf.Chr(146)]=224
	me.cw[gopdf.Chr(147)]=395
	me.cw[gopdf.Chr(148)]=398
	me.cw[gopdf.Chr(149)]=343
	me.cw[gopdf.Chr(150)]=646
	me.cw[gopdf.Chr(151)]=763
	me.cw[gopdf.Chr(152)]=458
	me.cw[gopdf.Chr(153)]=594
	me.cw[gopdf.Chr(154)]=486
	me.cw[gopdf.Chr(155)]=287
	me.cw[gopdf.Chr(156)]=841
	me.cw[gopdf.Chr(157)]=237
	me.cw[gopdf.Chr(158)]=482
	me.cw[gopdf.Chr(159)]=596
	me.cw[gopdf.Chr(160)]=237
	me.cw[gopdf.Chr(161)]=274
	me.cw[gopdf.Chr(162)]=543
	me.cw[gopdf.Chr(163)]=560
	me.cw[gopdf.Chr(164)]=646
	me.cw[gopdf.Chr(165)]=576
	me.cw[gopdf.Chr(166)]=244
	me.cw[gopdf.Chr(167)]=591
	me.cw[gopdf.Chr(168)]=506
	me.cw[gopdf.Chr(169)]=732
	me.cw[gopdf.Chr(170)]=418
	me.cw[gopdf.Chr(171)]=475
	me.cw[gopdf.Chr(172)]=519
	me.cw[gopdf.Chr(173)]=375
	me.cw[gopdf.Chr(174)]=732
	me.cw[gopdf.Chr(175)]=479
	me.cw[gopdf.Chr(176)]=367
	me.cw[gopdf.Chr(177)]=506
	me.cw[gopdf.Chr(178)]=397
	me.cw[gopdf.Chr(179)]=395
	me.cw[gopdf.Chr(180)]=342
	me.cw[gopdf.Chr(181)]=584
	me.cw[gopdf.Chr(182)]=499
	me.cw[gopdf.Chr(183)]=291
	me.cw[gopdf.Chr(184)]=256
	me.cw[gopdf.Chr(185)]=269
	me.cw[gopdf.Chr(186)]=431
	me.cw[gopdf.Chr(187)]=475
	me.cw[gopdf.Chr(188)]=682
	me.cw[gopdf.Chr(189)]=714
	me.cw[gopdf.Chr(190)]=781
	me.cw[gopdf.Chr(191)]=480
	me.cw[gopdf.Chr(192)]=608
	me.cw[gopdf.Chr(193)]=608
	me.cw[gopdf.Chr(194)]=608
	me.cw[gopdf.Chr(195)]=608
	me.cw[gopdf.Chr(196)]=608
	me.cw[gopdf.Chr(197)]=608
	me.cw[gopdf.Chr(198)]=883
	me.cw[gopdf.Chr(199)]=602
	me.cw[gopdf.Chr(200)]=539
	me.cw[gopdf.Chr(201)]=539
	me.cw[gopdf.Chr(202)]=539
	me.cw[gopdf.Chr(203)]=539
	me.cw[gopdf.Chr(204)]=280
	me.cw[gopdf.Chr(205)]=280
	me.cw[gopdf.Chr(206)]=280
	me.cw[gopdf.Chr(207)]=280
	me.cw[gopdf.Chr(208)]=640
	me.cw[gopdf.Chr(209)]=664
	me.cw[gopdf.Chr(210)]=649
	me.cw[gopdf.Chr(211)]=649
	me.cw[gopdf.Chr(212)]=649
	me.cw[gopdf.Chr(213)]=649
	me.cw[gopdf.Chr(214)]=649
	me.cw[gopdf.Chr(215)]=500
	me.cw[gopdf.Chr(216)]=646
	me.cw[gopdf.Chr(217)]=645
	me.cw[gopdf.Chr(218)]=645
	me.cw[gopdf.Chr(219)]=645
	me.cw[gopdf.Chr(220)]=645
	me.cw[gopdf.Chr(221)]=596
	me.cw[gopdf.Chr(222)]=572
	me.cw[gopdf.Chr(223)]=596
	me.cw[gopdf.Chr(224)]=505
	me.cw[gopdf.Chr(225)]=505
	me.cw[gopdf.Chr(226)]=505
	me.cw[gopdf.Chr(227)]=505
	me.cw[gopdf.Chr(228)]=505
	me.cw[gopdf.Chr(229)]=505
	me.cw[gopdf.Chr(230)]=790
	me.cw[gopdf.Chr(231)]=487
	me.cw[gopdf.Chr(232)]=499
	me.cw[gopdf.Chr(233)]=499
	me.cw[gopdf.Chr(234)]=499
	me.cw[gopdf.Chr(235)]=499
	me.cw[gopdf.Chr(236)]=266
	me.cw[gopdf.Chr(237)]=266
	me.cw[gopdf.Chr(238)]=266
	me.cw[gopdf.Chr(239)]=266
	me.cw[gopdf.Chr(240)]=559
	me.cw[gopdf.Chr(241)]=531
	me.cw[gopdf.Chr(242)]=531
	me.cw[gopdf.Chr(243)]=531
	me.cw[gopdf.Chr(244)]=531
	me.cw[gopdf.Chr(245)]=531
	me.cw[gopdf.Chr(246)]=531
	me.cw[gopdf.Chr(247)]=536
	me.cw[gopdf.Chr(248)]=531
	me.cw[gopdf.Chr(249)]=531
	me.cw[gopdf.Chr(250)]=531
	me.cw[gopdf.Chr(251)]=531
	me.cw[gopdf.Chr(252)]=531
	me.cw[gopdf.Chr(253)]=482
	me.cw[gopdf.Chr(254)]=534
	me.cw[gopdf.Chr(255)]=482
	me.up = -73
	me.ut = 49
	me.fonttype = "TrueType"
	me.name = "Roboto-BoldItalic"
	me.enc = "cp1252"
	me.desc = make([]gopdf.FontDescItem,8)
	me.desc[0] =  gopdf.FontDescItem{ Key:"Ascent",Val : "750" }
	me.desc[1] =  gopdf.FontDescItem{ Key: "Descent", Val : "-250" }
	me.desc[2] =  gopdf.FontDescItem{ Key:"CapHeight", Val :  "711" }
	me.desc[3] =  gopdf.FontDescItem{ Key: "Flags", Val :  "32" }
	me.desc[4] =  gopdf.FontDescItem{ Key:"FontBBox", Val :  "[-706 -271 1192 1048]" }
	me.desc[5] =  gopdf.FontDescItem{ Key:"ItalicAngle", Val :  "0" }
	me.desc[6] =  gopdf.FontDescItem{ Key:"StemV", Val :  "120" }
 	me.desc[7] =  gopdf.FontDescItem{ Key:"MissingWidth", Val :  "237" } 
 }
func (me * Roboto-BoldItalic)GetType() string{
	return me.fonttype
}
func (me * Roboto-BoldItalic)GetName() string{
	return me.name
}	
func (me * Roboto-BoldItalic)GetDesc() []gopdf.FontDescItem{
	return me.desc
}
func (me * Roboto-BoldItalic)GetUp() int{
	return me.up
}
func (me * Roboto-BoldItalic)GetUt()  int{
	return me.ut
}
func (me * Roboto-BoldItalic)GetCw() gopdf.FontCw{
	return me.cw
}
func (me * Roboto-BoldItalic)GetEnc() string{
	return me.enc
}
func (me * Roboto-BoldItalic)GetDiff() string {
	return me.diff
}
func (me * Roboto-BoldItalic) GetOriginalsize() int{
	return 98764
}
func (me * Roboto-BoldItalic)  SetFamily(family string){
	me.family = family
}
func (me * Roboto-BoldItalic) 	GetFamily() string{
	return me.family
}
