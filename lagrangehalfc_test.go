package tfhe

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func printTorusPolynomial(t *TorusPolynomial) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d, ", t.CoefsT[i])
	}
	fmt.Printf("\n")
}

func printLagrange(p *LagrangeHalfCPolynomial) {
	for i := 0; i < 10; i++ {
		fmt.Printf("(%f, %f i), ", real(p.coefsC[i]), imag(p.coefsC[i]))
	}
	fmt.Printf("\n")
}

func createTestTorusPolynomial() *TorusPolynomial {
	c := []Torus32{-1865008400, 470211269, -689632771, 1115438162, -1009961146, 143542609, 1474833166, 1998097154, -743203371, 1505795332, -2045554382, 1624379146, -37472977, 156091742, 1899894088, -1210297292, -1557125705, 357571487, 1927702193, 685118021, -1952636241, 824938978, -185075636, 997389811, -669037148, -1489662526, 1102246879, 1816731563, 1807130334, -1255430505, 1004844894, -1764527821, -919864291, 1070477901, 1912844172, 1808266295, -1690603250, -1557810092, -1257795641, -623157681, -1465572687, -684303797, 832633818, 316824709, 1815859898, 2051724828, 318153054, -934372970, -1098406643, 2063936095, -1718508330, -796138426, -748926889, 1724586123, -1124354143, 436476767, -211154555, 304987841, -245568255, -1799164911, -1362924059, -1857338490, 971307214, 462242382, -298801229, 1086531965, 1755699912, 1771024149, 1104740030, -1258364252, 1882410544, 1567121207, 140002773, -179027348, 298918981, 2124639786, -1189655634, -1042260880, 1732267503, 454233499, 1323602328, -2145596011, -1250428800, -351285635, -1971953469, -1688895389, -1312492104, 578134253, 605925147, -2123182237, -36000852, -387135268, -1292476584, 379847696, -1929219042, -784438327, -862823205, -1462913364, 1262696573, -637825383, -1800261938, 1847320611, -1012434887, -884507452, 925238585, -1352428776, 959637301, 832861197, -1979414689, 545293944, 1564674543, 2018984923, -1158708792, 1420250111, 1300601357, 1122721012, -766697006, 1971062964, -64238380, -335060346, -629320190, 1614285129, -528575729, -2138348418, 171620200, -207659129, 13618673, 1113423931, -979203770, 463482571, 1346174539, 1523252768, 601915876, -1924580678, 1666190315, -2028773350, -1803298028, -1121406691, 1116666768, 322791557, -357164898, -28168569, 3559271, 1051519022, 1586901780, -1884402099, 1318824508, 270438663, 2048760590, 1799809562, -1236229297, -514070449, -718325862, -1697452024, -972864132, -1769581807, 927515971, -671272276, 1795739908, -57238595, -1974926720, 1652354186, 1651659857, 1609005589, -1671456710, 1401494898, 797716613, -1645634978, 356326668, -725863409, -109184196, 1333685347, 1550207530, 115307824, 1114919528, 2098478512, 706642598, 1596168658, -1779449325, 2018044933, -1881520433, 343775970, 225882717, 702715824, 2008416043, 777568663, 977260517, 516341988, -1242150391, 622268741, -1682797931, -619410635, -1325380731, -1160150620, -439476922, 6933751, -720075253, 401203080, -846105773, -1401250974, -1053000939, -2058533934, -500224070, -1032288485, -1956373360, 562165086, -1647962611, 1008240335, 73536493, 67943785, -995297965, 866683781, -1828817673, 2001271628, 1230092636, -117574146, -1409579542, -1940169284, 48114451, 1840127447, 1448530829, -1735201672, -1845310184, 1258459314, -2090919601, 1987502654, 1785108618, 346645552, -1006252222, -969081451, 1434444853, 1227464300, 947712911, -1068856883, -1026360897, -2072193930, -1917315239, 1185044666, 1312349940, 1045139813, 1496786968, 1045501367, 1525679888, 594094830, 601518174, 1283070431, 494500519, -215776759, -748366836, -137576522, -303297865, -102153779, 264032578, 1508545799, 231367329, 739972374, -520504639, 710298147, -916582708, 434826717, -55087537, 1728536149, -1804788635, -1826968731, 817350098, 1398644498, -2018328688, 67439093, 1743532051, 2018358640, 1245245247, 407052491, 192958057, -1252420281, -698591802, 1746004784, 1925270568, -1336841778, 1476562290, -1982510178, -1121757885, 2095688077, 1462863339, 157863572, 1039108997, 1073005944, -1833530916, -1974993388, -79112871, -1725916635, 1236261659, 1257726607, 790709215, 1288775052, 319518552, -1117166935, 1388337876, 1715914097, 768094282, 717918281, -715420582, -854162921, -170017286, 1109242887, 567222275, -357237888, -784648809, -1163837649, 334734504, -440490218, 619880736, 688507581, 1544374105, 801783025, -1988891279, -931615702, -1379703097, 994411016, 2028195353, 108408484, 1279743896, -1922901183, 1789665271, 1414318906, -601736182, 1797619586, 1148637871, 694558228, 1796102643, 1150957460, 522329227, -1617111699, -352193657, 1338403327, 1871067776, -695644568, 976132619, 69297953, 1836275588, 1420400193, 128762701, 378859172, -220665233, -663719361, 1222648125, 585066580, -82815267, -467132268, -85781621, 1231367541, -1870366054, 1084605920, 254256216, -1317035803, 1784766667, -1882389039, -324924288, -1920790029, -1372311988, 1247622087, 337486257, -1784323721, -705767268, 1903766465, 734112719, 1502463514, -1782495308, 157160833, -1794008747, 1660439613, -1215629703, -151780977, -1575637126, 1702482602, -1688103016, -1330092096, -153485724, 765221730, -448522391, 821225630, -1436040373, -183893003, 2023078856, 756044142, -919356522, 1755174367, 934961289, 749808925, -173691506, 1598973560, -480139377, -833099640, 270415766, 286245354, -1876834163, -535743833, -1161973895, -678720683, -313024772, -1889406237, -2018071186, 1286759524, 115377426, 68372448, -54831670, -1658367758, -1185810424, 2003613926, -1044753096, 1845784855, -807442881, 1383816523, -2091648430, 1203288089, 259181367, 295741959, 1165225314, 1903958410, -1368821078, -1504932614, 1490194923, -1285151394, 1111527893, -216321785, 1479904488, -1268565251, -1733777998, 1359225232, 743955676, 1057007320, 684079389, 830889292, 1643966959, -81774830, 1165706775, -1714514604, 1103289072, 314974531, -1022664974, -919924094, 1457517396, -333631556, -549595602, 864971327, 1192825662, 1319631679, -2062307408, 737637588, -1779418889, 910200909, 1568463427, 785214197, 185984464, 898003985, -1448987963, 236507133, 127446719, 1971340602, 1843638688, -330013787, -720812366, 2089916416, 342194327, -1539031274, 606086944, 925188408, 1630361893, -1785669895, -2139283639, -1764905476, -712377014, -1144678828, -1946385236, -747224906, 437900412, -2112963758, -1913229280, 1205024131, -2082972743, -1101556202, 1178767756, 1617778547, -1777131463, 816959682, -1500254468, 1912233248, 1521446977, 1130220827, 196566711, 1679612133, -1486361640, -522640730, 1385869607, -1816272040, -1169950782, 348811263, -1788241072, 2116197139, -1431147768, -1386573619, -1455671213, 1377343942, 80923230, 167963453, 1054676514, -1604774868, 1414362802, -2143909489, -95529705, 355207552, -83846636, -1600814131, 521861726, -968024567, -3672039, -1342091098, -546486888, 1377965110, 220048633, -293430655, 1999253740, 2094655067, 1783848209, -1740273717, 806244816, -865548671, -478923087, -250042940, -11465099, -1074228984, 1956041512, -209393385, -969878495, 565116024, 720707383, 972559126, 1432595546, -388254451, 1163498600, 873488584, -639349618, -727616188, 593550211, 647386829, -55318916, -1267105619, -438004945, -854684762, 156848783, 823304037, 839521933, 1200676935, 23929709, -1315585028, 576783051, -1800406452, 1735832448, -1735873317, 1391900801, 156454402, 813799940, -1088227110, -1683099952, 758332962, 943068409, -909163350, 307614737, 1073387843, -1776125600, 804493646, -1793825853, -966767714, 755632434, -181059755, 1731634140, 93216329, 279385575, 2081384199, 1783878627, -1490221282, 714829746, 176941270, 438415345, -1795432692, -1723430944, -1616999934, 1863997492, -866074623, 902053218, 633063276, 1018594747, -1150980701, -734300668, -939940441, -2020358703, -562291915, -1131315141, 1017739959, -1060019272, -1296032455, 1366488945, 1593437398, -455895709, 1806282460, 2028965131, 2031439868, -367853639, -1656567124, -297578155, -97385615, 1610836118, -1280528784, -1369971539, -1709555540, 1768623320, 891421354, 1295705219, 279554414, -1638888753, 1461885892, -1833966133, -574240014, -1684814391, -345401437, 581326625, -810000923, 982972978, -590054257, -1260924640, 219596121, 935814237, -1626765834, -790612357, 1169606049, -1137457388, 1833517716, 390801190, 1140164826, 1557551345, -1636927422, 1863776403, 83021269, -200345498, -876786565, 453854174, -897396754, 332985597, -1487078532, 791581614, -2072384081, -492645527, 650410713, 2077160139, 1498568005, 331637302, 1214317403, -2120722755, 1729784607, 1236203844, -1050494202, 1019571745, 1256668085, -589222717, 741287253, 70774699, 878627491, 349037104, -113750905, -1433147874, 1793222863, 947876488, 126799937, 1383594056, 448578094, 788169082, 265693516, 1746407389, -986649596, -573620773, 1180433794, 1517882900, 1548224371, 395903722, 653136422, 502534869, -1014279796, 1448716159, -866000904, 1762940467, 111398397, -435118068, 36996359, 1346094042, -1477495195, 688525090, 1748669805, -1923935623, 1897969707, -824361698, -261540348, -1337964770, 1726137534, 1391296199, 341137843, -1992904927, -1636687070, 739586033, -1466705302, -42647065, -725467789, -152681049, -459052136, -748957174, 389036159, 52167504, -1108696112, 1901564463, 1866766830, -1735503092, -1662647073, 1756525911, -979867914, -1754371011, 1919088457, 600168340, -1392065194, -658787938, -2054802037, 524035687, -921962104, 753972431, 1707280684, 693437779, -2020388915, 1107834617, -980916674, -40181211, 1381308112, 410827174, -1778801382, 1994557577, -1294102953, -921423029, 180901983, 1597972592, -730798154, 1088358880, 575542596, 976308948, 374079592, 1595884376, 2067859119, -1923269146, 300298038, -1939465820, -1887630863, -902854019, 2137463216, 138796480, -1340307736, -1635393010, -1859477220, -769503844, 1408000937, 753979194, -1779004749, 264290317, 258663569, 287957427, 905358966, 1223911129, -1815959087, 1289804406, 208657545, -1410259941, 1999939611, -1620367314, 1718056160, -887436691, 1743424056, 823466734, -945703872, 1507034379, 566900172, 109900667, -2046010696, -442434321, 601535870, -1056267244, -1504876028, -1910654715, 314958325, -1661860726, 350123869, -646390314, -1970640793, -1874408651, -1644715754, -1355867706, 431830534, 1481024359, -507613707, -2119260746, -264791994, 116601980, 2012035562, 691321848, 377054187, -764510235, 2015119922, -477336123, -1196532292, 282442304, 731933548, 786651956, 1557674462, -287680585, 902895426, 1407106869, 1339878050, 1016172229, 1215750015, -1843322987, 1545417695, -532095622, 138584224, 1484148406, 2139020932, -1146002430, -1689625518, 1793372301, 398087471, 1755482408, -1144237941, 699723789, -191493434, -1050600136, 323681477, -124325344, -1555382801, -2046152069, -209911586, -998555083, -1360389207, -84912258, 1416294890, 1253448195, 798670174, 95293709, -394683599, -785482526, 360623719, 1597375813, 1438868374, -558931037, 2000444073, -2016186608, 779284472, -1115077497, -400762032, 4369624, -1995891363, 13453, -59104648, 1280276408, -421700979, 488783982, 572739354, 434477294, 1386848196, -1225895230, -99915850, -1388066431, 286810812, -1853660625, 1465649963, 2141370368, -1725178204, -1812228072, -817929041, 755644605, 314484208, -1028883489, 1853260907, -99661539, -300493789, -199165629, -106787643, -210810941, -1358163025, 35956990, 703299413, -895456568, -179327627, 1835095292, -1854929466, 308917854, -835505639, -1013885354, 401868521, -491088885, 1177855792, 55642690, 1639709502, -2084700262, 1842790971, 1567987267, 713912908, 514943161, -403718961, 218352659, 498508974, 438285374, -2005591359, 2107043748, 1132626518, 783035737, -1553387872, 970181002, -1566919750, 2011126219, 369810414, -1536563284, -979862875, -2143249261, -2079494391, -827178339, -1823047339, -835949657, 820890845, 972569747, -1841991017, 2079712808, 283353455, 274958550, -844109847, 500079511, -2000779582, -1106411782, 303528880, 1948596365, 495053397, 534195434, -745274232, -1662799491, -1951304761, 16977870, -663079456, 553324186, 2034034825, -1543427451, 427543547, -210717103, -369746838, -2105337206, 1378800754, -287970603, -1945186946, -297909622, 1113746948, 1043108726, -1790246499, 1482120112, 1480894385, -1565198121, -1243617027, -1242313988, -1927925765, 925036887, -789077306, 1627507449, 907527235, 1215898384, -1218002134, 1295059758, 1020392974, -1041935114, 2101995236, 422967956, -467529901, -1835039641, -577875528, -1099745180, 60393210, 1410637151, -1911838180, -476270969, 953510824, 331015785, 1275577861, -354453943, -1099822892, 1503614724, -213478807, -1771988184, 1044295370, -1003813813, -604429609, 873975952, -2123836862, -2021485248, -49143650, 517100454, -1370952606, 1762495479, 716794258, 793755781, -1818478793, -842124041, -928248662, -1252536917, -787621892, -1619839321, 1785433620, 1386973241, 2072476161, 57561407, 1189280587, 1068367964, 648950707, -529749685, -2135362049, 943010066, 1445244852}

	return &TorusPolynomial{
		N:      1024,
		CoefsT: c,
	}
}

//TEST(LagrangeHalfcTest, fftIsBijective) {
func TestFftIsBijective(t *testing.T) {
	assert := assert.New(t)

	NBTRIALS := 1
	toler := 1e-9
	var N int32 = 1024
	for trials := 0; trials < NBTRIALS; trials++ {
		//a := NewTorusPolynomial(N)
		acopy := NewTorusPolynomial(N)
		b := NewTorusPolynomial(N)
		afft := NewLagrangeHalfCPolynomial(N)

		//torusPolynomialUniform(a)
		a := createTestTorusPolynomial()
		fmt.Printf("torusPolynomialUniform(a):")
		printTorusPolynomial(a)

		TorusPolynomialCopy(acopy, a)
		fmt.Printf("torusPolynomialCopy(acopy, a):")
		printTorusPolynomial(acopy)

		torusPolynomialIfft(afft, a)
		fmt.Printf("torusPolynomialIfft(afft, a):")
		printLagrange(afft)

		torusPolynomialFft(b, afft)
		fmt.Printf("torusPolynomialFft(b, afft):")
		printTorusPolynomial(b)

		fmt.Printf("A: \n")
		printTorusPolynomial(a)

		fmt.Printf("B: \n")
		printTorusPolynomial(b)

		assert.EqualValues(torusPolynomialNormInftyDist(a, acopy), 0)
		assert.LessOrEqual(torusPolynomialNormInftyDist(a, b), toler)
	}
}

//TEST(LagrangeHalfcTest, LagrangeHalfCPolynomialClear) {
func TestLagrangeHalfCPolynomialClear(t *testing.T) {
	assert := assert.New(t)
	NBTRIALS := 10
	var N int32 = 1024
	for trials := 0; trials < NBTRIALS; trials++ {
		a := NewTorusPolynomial(N)
		zero := NewTorusPolynomial(N)
		afft := NewLagrangeHalfCPolynomial(N)
		LagrangeHalfCPolynomialClear(afft)
		torusPolynomialUniform(a)
		torusPolynomialClear(zero)
		torusPolynomialFft(a, afft)
		assert.EqualValues(torusPolynomialNormInftyDist(zero, a), 0)
	}
}

/** sets to this torus32 constant */
//EXPORT void LagrangeHalfCPolynomialSetTorusConstant(LagrangeHalfCPolynomial* result, const Torus32 mu);
func TestLagrangeHalfCPolynomialSetTorusConstant(t *testing.T) {
	assert := assert.New(t)
	NBTRIALS := 10
	var N int32 = 1024
	for trials := 0; trials < NBTRIALS; trials++ {
		mu := UniformTorus32Dist()
		a := NewTorusPolynomial(N)
		cste := NewTorusPolynomial(N)
		afft := NewLagrangeHalfCPolynomial(N)
		torusPolynomialUniform(a)

		//tested function
		LagrangeHalfCPolynomialSetTorusConstant(afft, mu)
		torusPolynomialFft(a, afft)

		//expected result
		torusPolynomialClear(cste)
		cste.CoefsT[0] = mu

		assert.EqualValues(torusPolynomialNormInftyDist(cste, a), 0)
	}
}

//EXPORT void LagrangeHalfCPolynomialAddTorusConstant(LagrangeHalfCPolynomial* result, const Torus32 cst);
func TestLagrangeHalfCPolynomialAddTorusConstant(t *testing.T) {
	assert := assert.New(t)
	NBTRIALS := 1
	var N int32 = 1024
	toler := 1e-9
	for trials := 0; trials < NBTRIALS; trials++ {
		var mu int32 = 42424242 //UniformTorus32Dist()
		//a := NewTorusPolynomial(N)
		a := createTestTorusPolynomial()
		aPlusCste := NewTorusPolynomial(N)
		b := NewTorusPolynomial(N)
		afft := NewLagrangeHalfCPolynomial(N)

		//torusPolynomialUniform(a)
		torusPolynomialIfft(afft, a)
		LagrangeHalfCPolynomialAddTorusConstant(afft, mu)
		torusPolynomialFft(b, afft)

		TorusPolynomialCopy(aPlusCste, a)
		aPlusCste.CoefsT[0] += mu

		assert.LessOrEqual(torusPolynomialNormInftyDistSkipFirst(aPlusCste, b), toler)
	}
}

func MockIntPolynomial() *IntPolynomial {
	//return &IntPolynomial{N: 4, Coefs: []int32{-1865008400, 470211269, -689632771, 1115438162}}
	return &IntPolynomial{N: 4, Coefs: []int32{9, -10, 7, 6}}
}

func MockTorusPolynomial() *TorusPolynomial {
	// return &TorusPolynomial{N: 4, CoefsT: []int32{156091742, 1899894088, -1210297292, -1557125705}}
	return &TorusPolynomial{N: 4, CoefsT: []int32{-5, 4, 0, -2}}
}

func TestTorusPolynomialMultFFT(t *testing.T) {
	assert := assert.New(t)
	NBTRIALS := 1
	var N int32 = 1024
	toler := 1e-9
	for trials := 0; trials < NBTRIALS; trials++ {
		a := NewIntPolynomial(N)
		b := NewTorusPolynomial(N)
		//b := createTestTorusPolynomial()
		//a := MockIntPolynomial()
		//b := MockTorusPolynomial()
		aB := NewTorusPolynomial(N)
		aBref := NewTorusPolynomial(N)

		for i := int32(0); i < N; i++ {
			a.Coefs[i] = UniformTorus32Dist()%1000 - 500
		}
		torusPolynomialUniform(b)

		//torusPolynomialMultKaratsuba(aBref, a, b)
		torusPolynomialMultNaive(aBref, a, b)

		//torusPolynomialMultNaive(aB, a, b)
		torusPolynomialMultFFT(aB, a, b)
		assert.LessOrEqual(torusPolynomialNormInftyDist(aB, aBref), toler)
	}
}

func TestTorusPolynomialAddMulRFFT(t *testing.T) {
	assert := assert.New(t)
	NBTRIALS := 1
	var N int32 = 1024
	toler := 1e-9
	for trials := 0; trials < NBTRIALS; trials++ {
		a := NewIntPolynomial(N)
		b := NewTorusPolynomial(N)
		aB := NewTorusPolynomial(N)
		aBref := NewTorusPolynomial(N)
		for i := int32(0); i < N; i++ {
			a.Coefs[i] = UniformTorus32Dist()%1000 - 500
		}
		torusPolynomialUniform(b)
		torusPolynomialUniform(aB)
		TorusPolynomialCopy(aBref, aB)
		torusPolynomialAddMulRKaratsuba(aBref, a, b)
		torusPolynomialAddMulRFFT(aB, a, b)
		assert.LessOrEqual(torusPolynomialNormInftyDist(aB, aBref), toler)
	}
}

func TestTorusPolynomialSubMulRFFT(t *testing.T) {
	assert := assert.New(t)
	NBTRIALS := 1
	var N int32 = 1024
	toler := 1e-9
	for trials := 0; trials < NBTRIALS; trials++ {
		a := NewIntPolynomial(N)
		b := NewTorusPolynomial(N)
		aB := NewTorusPolynomial(N)
		aBref := NewTorusPolynomial(N)

		for i := int32(0); i < N; i++ {
			a.Coefs[i] = UniformTorus32Dist()%1000 - 500
		}
		torusPolynomialUniform(b)
		torusPolynomialUniform(aB)
		TorusPolynomialCopy(aBref, aB)
		torusPolynomialSubMulRKaratsuba(aBref, a, b)
		torusPolynomialSubMulRFFT(aB, a, b)
		assert.LessOrEqual(torusPolynomialNormInftyDist(aB, aBref), toler)
	}
}

//TEST(LagrangeHalfcTest, LagrangeHalfCPolynomialAddTo) {
func TestLagrangeHalfCPolynomialAddTo(t *testing.T) {
	assert := assert.New(t)
	NBTRIALS := 1
	var N int32 = 1024
	toler := 1e-9
	for trials := 0; trials < NBTRIALS; trials++ {
		a := NewTorusPolynomial(N)
		b := NewTorusPolynomial(N)
		aPlusB := NewTorusPolynomial(N)
		aPlusBbis := NewTorusPolynomial(N)
		afft := NewLagrangeHalfCPolynomial(N)
		bfft := NewLagrangeHalfCPolynomial(N)
		torusPolynomialUniform(a)
		torusPolynomialIfft(afft, a)
		torusPolynomialUniform(b)
		torusPolynomialIfft(bfft, b)
		LagrangeHalfCPolynomialAddTo(afft, bfft)
		torusPolynomialFft(aPlusBbis, afft)
		TorusPolynomialAdd(aPlusB, b, a)
		assert.LessOrEqual(torusPolynomialNormInftyDist(aPlusBbis, aPlusB), toler)
	}
}
