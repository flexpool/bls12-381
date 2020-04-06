package bls12381

var isogenyConstansG1 = [4][16]*fe{
	[16]*fe{
		&fe{5555391298090832668, 1871845530032595596, 4551034694774233518, 2584197799339864836, 15085749040064757844, 654075415717002996},
		&fe{9910598932128054667, 4357765064159749802, 1555960221863322426, 9671461638228026285, 1275132148248838779, 507072521670460589},
		&fe{11908177372061066827, 18190436643933350086, 6603102998733829542, 6581045210674032871, 16099974426311393401, 541581077397919012},
		&fe{5282195870529824577, 12365729195083706401, 2807246122435955773, 332702220601507168, 7339422050895811209, 1050416448951884523},
		&fe{10443415753526299973, 8852419397684277637, 1088333252544296036, 1174353327457337436, 1626144519293139599, 716651429285276662},
		&fe{7916646322956281527, 11909818257232418749, 1455301921509471421, 3317627683558310107, 12693445337245173919, 1798273032850409769},
		&fe{2577731109215284733, 8810166123993386985, 3186592767751348067, 15050850291391518479, 18435652654155870871, 1330813445865859326},
		&fe{8787912969482053798, 9653629252694769025, 1358451377919714320, 16331599695590198629, 13519934665722691825, 628078949001449512},
		&fe{16605411443261943819, 9536014432113026165, 8685402948537367476, 16291074259433785035, 407185289045737198, 713426768049972652},
		&fe{1001421273809907975, 724433776290697394, 16309429154639760781, 10003715605277815375, 307249038158020985, 688008371043525493},
		&fe{16622893420529658311, 18333652517857227637, 2139173376235292830, 16496634502105693419, 5355299366650241487, 382770009771704860},
		&fe{8276255265012938363, 9997870203437298645, 16819210142450232135, 5062450688048499179, 12776432501206859311, 1778476024187613533},
		&fe{0, 0, 0, 0, 0, 0},
		&fe{0, 0, 0, 0, 0, 0},
		&fe{0, 0, 0, 0, 0, 0},
		&fe{0, 0, 0, 0, 0, 0},
	},
	[16]*fe{
		&fe{13358415881952098629, 12009257493157516192, 13928884382876484932, 12988314785833227070, 11244145530317148182, 100673949996487007},
		&fe{2533162896381624793, 10578896196504721258, 4263020647280931071, 1255899686249737875, 17097124965295857733, 590960935246623182},
		&fe{10990404485039254780, 5344458621503091696, 1718862119039451458, 11600049052019063549, 18389973225607751698, 1092616849767867362},
		&fe{16377845895484993601, 15314247056264135931, 14543008873173635408, 4875476272346940127, 2030129768648768484, 1297689274107773964},
		&fe{6376927397170316667, 1460555178565443615, 18156708192400235081, 14761117739963869762, 8361091377443400626, 1421233557303902229},
		&fe{18127417459170613536, 5353764292720778676, 858818813615405862, 3528937506143354306, 12604964186779349896, 489837025077541867},
		&fe{15285065477075910543, 3650488990300576179, 7274499670465195193, 16100555180954076900, 7580582425312971905, 896074979407586822},
		&fe{7582945168915351799, 2506680954090651888, 10272835934257987876, 9924916350558121763, 13577194922650729507, 1698254565890367778},
		&fe{2009730524583761661, 11053280693947850663, 14409256190409559425, 3658799329773368860, 13529638021208614900, 869243908766415668},
		&fe{11058048790650732295, 7059501760293999296, 6596812464094265283, 14567744481299745071, 1591898617514919697, 1344004358835331304},
		&fe{8505329371266088957, 17002214543764226050, 6865905132761471162, 8632934651105793861, 6631298214892334189, 1582556514881692819},
		&fe{0, 0, 0, 0, 0, 0},
		&fe{0, 0, 0, 0, 0, 0},
		&fe{0, 0, 0, 0, 0, 0},
		&fe{0, 0, 0, 0, 0, 0},
		&fe{0, 0, 0, 0, 0, 0},
	},
	[16]*fe{
		&fe{3122824077082063463, 2111517899915568999, 14844585557031220083, 14713720721132803039, 9041847780307969683, 950267513573868304},
		&fe{11079511902567680319, 18338468344530008184, 6769016392463638666, 1504264063027988936, 8098359051856762276, 760455062874047829},
		&fe{1430247552210236986, 3854575382974307965, 14917507996414511245, 207936139448560, 9498310774218301406, 1438631746617682181},
		&fe{6654065794071117243, 2928282753802966791, 4144383358731160429, 12673586709493869907, 12918170109018188791, 844088361957958231},
		&fe{6416330705244672319, 3552017270878949117, 7777490944331917312, 7917192495177481567, 7271851377118683537, 253926972271069325},
		&fe{11903306495973637341, 11622313950541285762, 17991208474928993001, 12280964980743791783, 14941570282955772167, 143516344770893715},
		&fe{7324386472845891920, 16310961984705608217, 14050364318273732029, 410622978843904432, 13407944087243235067, 570579643952782879},
		&fe{10655681039374273828, 3913226275392147601, 9613292388335178165, 11852815148890010639, 17652581670569921892, 780578093363976825},
		&fe{10454026283255684948, 15005802245309313587, 4420421943175638630, 18052347756729021570, 12181908985148691767, 1485233717472293779},
		&fe{5056344670784885274, 15896288289018563095, 11120951801157184493, 7250506164525313606, 9295677455526059106, 1757175036496698059},
		&fe{417067620545670182, 113740147118943311, 7666319924200602156, 1469963335415292317, 13482947512490784447, 1353298443678343909},
		&fe{13069093794065563159, 18364685236451803588, 2235996605706292724, 1007629142299662669, 4077244143222018961, 162586537120788900},
		&fe{12976751790971550752, 10256454045927919861, 8968423978443605586, 91636529236982767, 9459527627289574163, 949550897353139410},
		&fe{10595118024452621845, 8010256778549625402, 10333144214150401956, 17682229685967587631, 8235697699445463546, 317883997785997129},
		&fe{16894283457285346118, 10513943172407809423, 4685513162956315481, 11558261883362075118, 574375951146893083, 1159440548124233311},
		&fe{9739780494108151959, 17207219630538774058, 553911396609642498, 6085929320386029624, 14175410874026216616, 1183751611824804793},
	},
	[16]*fe{
		&fe{16963992846030154524, 1796759822929186144, 15995221960860457854, 8232142361908220707, 5977498266010213481, 759868220591477233},
		&fe{7019489280640006651, 8025136855967848721, 17464762292772824538, 4490335113250743896, 7652702793653159798, 1129822927746498110},
		&fe{3164260796573156764, 2639884922337322818, 1251365706181388855, 13142429936036186189, 359878619957828340, 126848055205862465},
		&fe{17472832885692408710, 9911075278795900735, 2614390623136861791, 14474775734428698630, 6462878218464609418, 1225960780180864957},
		&fe{3586995257703132870, 2143554115308730112, 15207899356205612465, 4372523065560113828, 12811868595146042778, 307251632623424763},
		&fe{14298637377310410728, 10963101290308221781, 8192510423058716701, 1175370967867267532, 1029599188863854120, 678981456155013844},
		&fe{11149806480082726900, 3664985661428410608, 18095361538178773836, 14174906593575241395, 15305104369759711886, 901234928011491053},
		&fe{4727074327869776987, 15736954329525418288, 14642679026711520511, 11429849039208981702, 17333567062758618213, 951235897335772166},
		&fe{9130114290642375589, 14069725355798443159, 6621984191700563591, 270173975669947883, 6218390495944243859, 1077419361593130421},
		&fe{9144875514986933294, 16561351410666797616, 8591333879886582656, 15059370240386191395, 7834396448114781869, 946553772269403391},
		&fe{17809450171377747225, 15896956440537434491, 8451524482089653422, 1694507265233574136, 18224201536921880842, 317503425606567070},
		&fe{13940503876759740187, 8772047862193200131, 6080360161890657205, 7935486160089058373, 9407473295146243021, 1255078947940629503},
		&fe{1160821217138360586, 13542760608074182996, 11595911004531652098, 18158686636947034451, 13330657138280564947, 1773960737279760188},
		&fe{9132548444917292754, 16464415422105000789, 6319313500251671073, 12727658548847517900, 10985275115076354035, 1431541893474124246},
		&fe{662485641082390837, 260809847827618849, 6177381409359357075, 18231947741742261351, 18128540110746580014, 1079107229429227022},
		&fe{8505329371266088957, 17002214543764226050, 6865905132761471162, 8632934651105793861, 6631298214892334189, 1582556514881692819},
	},
}

var isogenyConstantsG2 = [4][4]*fe2{
	[4]*fe2{
		&fe2{
			fe{5185457120960601698, 494647221959407934, 8971396042087821730, 324544954362548322, 14214792730224113654, 1405280679127738945},
			fe{5185457120960601698, 494647221959407934, 8971396042087821730, 324544954362548322, 14214792730224113654, 1405280679127738945},
		},
		&fe2{
			fe{0, 0, 0, 0, 0, 0},
			fe{6910023028261548496, 9745789443900091043, 7668299866710145304, 2432656849393633605, 2897729527445498821, 776645607375592125},
		},
		&fe2{
			fe{724047465092313539, 15783990863276714670, 12824896677063784855, 15246381572572671516, 13186611051602728692, 1485475813959743803},
			fe{12678383550985550056, 4872894721950045521, 13057521970209848460, 10439700461551592610, 10672236800577525218, 388322803687796062},
		},
		&fe2{
			fe{4659755689450087917, 1804066951354704782, 15570919779568036803, 15592734958806855601, 7597208057374167129, 1841438384006890194},
			fe{0, 0, 0, 0, 0, 0},
		},
	},
	[4]*fe2{
		&fe2{
			fe{0, 0, 0, 0, 0, 0},
			fe{2250392438786206615, 17463829474098544446, 14571211649711714824, 4495761442775821336, 258811604141191305, 357646605018048850},
		},
		&fe2{
			fe{4933130441833534766, 15904462746612662304, 8034115857496836953, 12755092135412849606, 7007796720291435703, 252692002104915169},
			fe{8469300574244328829, 4752422838614097887, 17848302789776796362, 12930989898711414520, 16851051131888818207, 1621106615542624696},
		},
		&fe2{
			fe{8505329371266088957, 17002214543764226050, 6865905132761471162, 8632934651105793861, 6631298214892334189, 1582556514881692819},
			fe{0, 0, 0, 0, 0, 0},
		},
		&fe2{
			fe{0, 0, 0, 0, 0, 0},
			fe{0, 0, 0, 0, 0, 0},
		},
	},
	[4]*fe2{
		&fe2{
			fe{10869708750642247614, 13056187057366814946, 1750362034917495549, 6326189602300757217, 1140223926335695785, 632761649765668291},
			fe{10869708750642247614, 13056187057366814946, 1750362034917495549, 6326189602300757217, 1140223926335695785, 632761649765668291},
		},
		&fe2{
			fe{0, 0, 0, 0, 0, 0},
			fe{13765940311003083782, 5579209876153186557, 11349908400803699438, 11707848830955952341, 199199289641242246, 899896674917908607},
		},
		&fe2{
			fe{15562563812347550836, 2436447360975022760, 6528760985104924230, 5219850230775796305, 5336118400288762609, 194161401843898031},
			fe{16286611277439864375, 18220438224251737430, 906913588459157469, 2019487729638916206, 75985378181939686, 1679637215803641835},
		},
		&fe2{
			fe{11849179119594500956, 13906615243538674725, 14543197362847770509, 2041759640812427310, 2879701092679313252, 1259985822978576468},
			fe{0, 0, 0, 0, 0, 0},
		},
	},
	[4]*fe2{
		&fe2{
			fe{99923616639376095, 10339114964526300021, 6204619029868000785, 1288486622530663893, 14587509920085997152, 272081012460753233},
			fe{99923616639376095, 10339114964526300021, 6204619029868000785, 1288486622530663893, 14587509920085997152, 272081012460753233},
		},
		&fe2{
			fe{0, 0, 0, 0, 0, 0},
			fe{6751177316358619845, 15498000274876530106, 6820146801716041242, 13487284328327464010, 776434812423573915, 1072939815054146550},
		},
		&fe2{
			fe{7399695662750302149, 14633322083064217648, 12051173786245255430, 9909266166264498601, 1288323043582377747, 379038003157372754},
			fe{6002735353327561446, 6023563502162542543, 13831244861028377885, 15776815867859765525, 4123780734888324547, 1494760614490167112},
		},
		&fe2{
			fe{8505329371266088957, 17002214543764226050, 6865905132761471162, 8632934651105793861, 6631298214892334189, 1582556514881692819},
			fe{0, 0, 0, 0, 0, 0},
		},
	},
}

var swuParamsForG1 = struct {
	z           *fe
	zInv        *fe
	a           *fe
	b           *fe
	minusBOverA *fe
}{
	a:           &fe{3415322872136444497, 9675504606121301699, 13284745414851768802, 2873609449387478652, 2897906769629812789, 1536947672689614213},
	b:           &fe{18129637713272545760, 11144507692959411567, 10108153527111632324, 9745270364868568433, 14587922135379007624, 469008097655535723},
	z:           &fe{9830232086645309404, 1112389714365644829, 8603885298299447491, 11361495444721768256, 5788602283869803809, 543934104870762216},
	zInv:        &fe{1047701040585522704, 6568704757426767313, 7461573184509654906, 5499015922318795030, 11226104418450030905, 1048548528059189658},
	minusBOverA: &fe{370847444534405118, 4269648997187665026, 1978763176675559811, 2677363437243537255, 11096866317338941469, 683609622716391635},
}

var swuParamsForG2 = struct {
	z           *fe2
	zInv        *fe2
	a           *fe2
	b           *fe2
	minusBOverA *fe2
}{
	a: &fe2{
		fe{0, 0, 0, 0, 0, 0},
		fe{16517514583386313282, 74322656156451461, 16683759486841714365, 815493829203396097, 204518332920448171, 1306242806803223655},
	},
	b: &fe2{
		fe{2515823342057463218, 7982686274772798116, 7934098172177393262, 8484566552980779962, 4455086327883106868, 1323173589274087377},
		fe{2515823342057463218, 7982686274772798116, 7934098172177393262, 8484566552980779962, 4455086327883106868, 1323173589274087377},
	},
	z: &fe2{
		fe{9794203289623549276, 7309342082925068282, 1139538881605221074, 15659550692327388916, 16008355200866287827, 582484205531694093},
		fe{4897101644811774638, 3654671041462534141, 569769440802610537, 17053147383018470266, 17227549637287919721, 291242102765847046},
	},
	zInv: &fe2{
		fe{949978046398372251, 9282594348372276018, 12783089135259591524, 8269126545290330608, 594742981125487701, 491256564635846792},
		fe{15449598521694521480, 14910517655282017894, 6549664756007020895, 17931849781271742567, 2408680398672607296, 691271026505846537},
	},
	minusBOverA: &fe2{
		fe{10393275865055580083, 6888480573845999877, 11497223857339693790, 14306043441748627554, 5078453791572287059, 1040691004897901061},
		fe{3009155151022283512, 13768405011380760314, 14385194789933939525, 11380038592375636572, 333649986898415235, 833107612749638805},
	},
}