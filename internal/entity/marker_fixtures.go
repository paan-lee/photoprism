package entity

type MarkerMap map[string]Marker

func (m MarkerMap) Get(name string) Marker {
	if result, ok := m[name]; ok {
		return result
	}

	return *UnknownMarker
}

func (m MarkerMap) Pointer(name string) *Marker {
	if result, ok := m[name]; ok {
		return &result
	}

	return UnknownMarker
}

var MarkerFixtures = MarkerMap{
	"1000003-1": Marker{
		ID:         1,
		FileID:     1000003,
		SubjectUID: "lt9k3pw1wowuy3c3",
		MarkerSrc:  SrcImage,
		MarkerType: MarkerLabel,
		X:          0.308333,
		Y:          0.206944,
		W:          0.355556,
		H:          .355556,
	},
	"1000003-2": Marker{
		ID:         2,
		FileID:     1000003,
		SubjectUID: "",
		MarkerName: "Unknown",
		MarkerSrc:  SrcImage,
		MarkerType: MarkerLabel,
		X:          0.208333,
		Y:          0.106944,
		W:          0.05,
		H:          0.05,
	},
	"1000003-3": Marker{
		ID:         3,
		FileID:     1000003,
		SubjectUID: "",
		MarkerSrc:  SrcImage,
		MarkerType: MarkerLabel,
		MarkerName: "Center",
		X:          0.5,
		Y:          0.5,
		W:          0,
		H:          0,
	},
	"1000003-4": Marker{
		ID:             4,
		FileID:         1000003,
		SubjectUID:     "",
		MarkerSrc:      SrcImage,
		MarkerType:     MarkerFace,
		MarkerName:     "Jens Mander",
		LandmarksJSON:  []byte("[{\"name\":\"lp46\",\"x\":-0.08359375,\"y\":-0.027083334,\"h\":0.044791665,\"w\":0.03359375},{\"name\":\"lp46_v\",\"x\":0.08671875,\"y\":-0.009375,\"h\":0.045833334,\"w\":0.034375},{\"name\":\"lp44\",\"x\":-0.0546875,\"y\":-0.048958335,\"h\":0.044791665,\"w\":0.03359375},{\"name\":\"lp44_v\",\"x\":0.06328125,\"y\":-0.033333335,\"h\":0.045833334,\"w\":0.034375},{\"name\":\"lp42\",\"x\":-0.021875,\"y\":-0.03125,\"h\":0.044791665,\"w\":0.03359375},{\"name\":\"lp42_v\",\"x\":0.03203125,\"y\":-0.025,\"h\":0.045833334,\"w\":0.034375},{\"name\":\"lp38\",\"x\":-0.0265625,\"h\":0.044791665,\"w\":0.03359375},{\"name\":\"lp38_v\",\"x\":0.03125,\"y\":0.0052083335,\"h\":0.044791665,\"w\":0.03359375},{\"name\":\"lp312\",\"x\":-0.06796875,\"y\":-0.008333334,\"h\":0.045833334,\"w\":0.034375},{\"name\":\"lp312_v\",\"x\":0.06953125,\"y\":0.008333334,\"h\":0.044791665,\"w\":0.03359375},{\"name\":\"mouth_lp93\",\"x\":-0.00703125,\"y\":0.09375,\"h\":0.044791665,\"w\":0.03359375},{\"name\":\"mouth_lp84\",\"x\":-0.04921875,\"y\":0.128125,\"h\":0.045833334,\"w\":0.034375},{\"name\":\"mouth_lp82\",\"x\":-0.01328125,\"y\":0.16145833,\"h\":0.045833334,\"w\":0.034375},{\"name\":\"mouth_lp81\",\"x\":-0.0078125,\"y\":0.13333334,\"h\":0.044791665,\"w\":0.03359375},{\"name\":\"lp84\",\"x\":0.034375,\"y\":0.14479166,\"h\":0.044791665,\"w\":0.03359375},{\"name\":\"eye_l\",\"x\":-0.0484375,\"y\":-0.004166667,\"h\":0.030208332,\"w\":0.02265625},{\"name\":\"eye_r\",\"x\":0.0484375,\"y\":0.0052083335,\"h\":0.030208332,\"w\":0.02265625}]"),
		EmbeddingsJSON: []byte("[[0.019231493,-0.028809275,-0.083006255,-0.015598502,0.08550906,0.001886255,0.09019353,0.07551488,-0.011814484,0.03680722,-0.08332401,0.014950869,0.055766843,-0.0073407963,-0.0552993,0.05847239,0.026178556,-0.045643847,0.012550894,-0.012022383,-0.0040185535,0.00023647904,-0.01580684,-0.008704283,0.04575994,-0.046812356,0.056398943,-0.026091263,0.059522673,-0.044217024,0.014755385,-0.01350486,-0.049488768,-0.03139871,-0.006726978,-0.02347069,-0.0059445584,-0.004682308,-0.057403754,-0.06537466,0.013326172,-0.009667708,0.022370687,0.0037015954,0.03744496,0.052890837,-0.0077360696,-0.049944617,-0.03868134,0.001521219,0.03840492,-0.10928545,0.023024736,-0.055707198,-0.13260484,0.009903039,-0.04250921,-0.0040567834,0.03343564,-0.01785736,0.0043026204,-0.031062575,-0.0019649328,0.06487235,-0.14464019,-0.017717961,-0.0033534314,0.029505186,0.008849258,-0.0026131037,-0.06479913,-0.111862205,0.05469594,0.049985956,-0.00067700783,0.003068928,0.0018148758,0.0073374105,0.025748348,-0.0424614,0.062650666,0.058194485,-0.04309207,-0.020790769,-0.030982763,0.008360668,0.01289988,-0.019662105,0.0122521445,-0.00342255,-0.056044392,0.034414552,0.04604621,0.0074918787,0.033526078,-0.036619328,-0.047758896,-0.032501936,-0.08068566,-0.02964604,0.04137439,0.06888022,0.04018322,0.0023792675,-0.026837967,-0.049688686,-0.057930365,-0.064863205,0.004485477,-0.026958624,-0.025907256,0.009216111,-0.014622554,-0.037213538,0.078393415,-0.054682203,-0.009757617,0.03503295,0.027951613,-0.0014038666,0.06851987,0.020453943,-0.00996363,-0.12156495,-0.017301193,-0.032558206,0.07816977,0.029640608,0.03150378,0.047710016,-0.009032255,0.013137696,-0.056541067,-0.0075266357,0.007981631,-0.025004586,0.030189874,-0.007878598,0.03279407,0.059711676,-0.003458093,0.01758815,-0.010460049,0.012645757,0.006580964,-0.019848932,0.02724201,0.001597322,0.0998137,-0.03012735,0.011577833,0.0028034616,0.029448563,-0.011648439,0.05161503,-0.028912308,-0.07435124,-0.0033618107,0.030619241,0.04033193,-0.0049585714,-0.016229536,0.05426284,-0.031509526,0.014953531,-0.025750436,0.000924462,-0.008974987,0.03417905,0.0011432022,0.007551494,0.001544253,0.016510574,0.044189014,0.054653972,0.0811879,-0.0036358214,0.035907324,0.047312576,-0.0068604983,-0.029337948,0.027281888,-0.08003064,-0.054304074,0.0073654098,0.019096468,0.004265153,-0.03670545,0.0049256124,-0.00017001168,0.04502574,-0.00074317795,0.024450991,-0.052183975,0.15243798,-0.010411264,-0.016080985,0.017625313,0.022576308,-0.053998422,0.0023750497,-0.057889163,0.056488033,-0.025041293,0.011416959,0.08541408,0.08386572,0.0396809,0.0920811,-0.048865005,-0.026257817,0.032705985,-0.053586084,0.032277677,0.07027237,0.016478207,-0.08976395,-0.006354579,-0.038090594,0.024028458,0.013410826,0.08724634,-0.028266877,0.07450935,0.03037237,-0.018165091,-0.029061696,-0.0006272013,-0.005710233,0.037058197,-0.085284434,-0.04267883,0.029252399,-0.033832725,-0.07989745,0.025662571,0.019242961,-0.10760675,0.0021055646,0.0021679339,0.0382236,0.01998776,-0.040100973,-0.041129857,0.0025446033,-0.03981458,-0.031129645,-0.061818115,-0.031583495,0.08146585,-0.042177603,0.05506262,-0.045823902,-0.031090494,-0.04478884,-0.06997962,-0.0024934823,0.0020288285,-0.074601755,-0.029107075,-0.03953502,0.0551208,-0.05685647,-0.0010502128,0.04371499,-0.031777892,0.030995583,-0.056923013,0.04047056,-0.05058555,-0.020007737,-0.0034168877,0.040992126,-0.0065717557,-0.060389657,-0.015070164,0.0300555,-0.0049498156,0.035197794,-0.010181344,-0.01548949,0.09214268,0.06594641,-0.04095856,0.031043377,0.016846823,0.011824804,0.010486359,0.006952729,-0.0423542,-0.038642637,0.037917823,-0.042705245,0.002531653,0.049575835,-0.008132699,0.060207773,0.050799962,0.029537434,-0.011404661,0.07556842,0.044129632,-0.025151744,-0.04391785,-0.057073284,0.066147126,-0.03833207,-0.11469866,0.018607577,0.03958792,0.005069568,-0.0022620235,-0.06963503,-0.027407918,-0.01658652,-0.08928287,0.04213307,0.0653322,0.070556566,-0.08890351,0.05253341,-0.03952343,-0.03277939,-0.07603576,0.023622843,0.01869451,0.0012659469,-0.016868249,-0.048114307,-0.12678534,-0.023940234,0.03518056,0.030265998,-0.027942419,-0.030615386,-0.027904578,0.04184872,-0.06871633,0.018810445,-0.0050423085,0.01902196,0.036709674,0.05114107,-0.008238005,-0.033530205,0.009237725,-0.019793255,-0.011098562,-0.040874466,0.028055958,0.052516278,0.002427102,-0.018204305,-0.029648917,-0.071319304,-0.0468887,0.061347924,0.031454504,-0.031552624,-0.08102158,0.074395515,-0.048840746,0.030817837,0.07313089,0.062438965,-0.00383427,0.05508265,-0.0077238525,0.026816545,-0.050590646,-0.02294128,0.0017293892,-0.018405396,0.03635994,-0.0025451558,-0.005002223,-0.059726283,0.008561757,-0.05424462,-0.009393793,-0.040830627,0.02030258,0.003973316,0.020119876,-0.017317869,0.01453977,-0.025234057,0.072000384,-0.0413773,-0.111050114,0.06299958,0.00016407811,-0.050199028,-0.013726295,0.0100843245,0.022901619,0.056635447,0.038121402,0.024791898,-0.011317786,-0.059741378,-0.0004985886,-0.0129778,0.04761788,0.018754557,0.035193384,-0.021030819,0.042050865,0.08013355,-0.042123444,0.065898865,0.03751851,-0.049639545,-0.03844303,0.006924326,0.022647297,-0.008048924,0.015995622,-0.033804823,0.007260562,-0.046132274,-0.0064429105,0.031663842,-0.006572463,-0.06781134,0.013187006,0.013765757,-0.03454214,-0.015666338,-0.0023530773,-0.07751217,-0.008714079,-0.010440672,0.026577305,0.066843055,-0.00440322,-0.0071120844,0.008546605,0.087584786,-0.027872834,-0.043215286,-0.0657154,0.042205643,0.003758622,-0.029912153,0.020607445,-0.0034470167,0.040666997,0.081489064,-0.044269655,-0.006542095,-0.054021314,-0.029242665,-0.027248744,-0.02227558,0.082040176,-0.030761424,-0.023510806,-0.06973938,-0.0032560013,0.055813447,0.03283226,0.068810284,0.029060816,-0.03417918,0.004436392,-0.018858656,-0.0056046643,-0.034613956,-0.024074866,0.029658342,-0.023564866,-0.011503043,0.04425076,-0.017220033,0.094600976,0.0664404,0.0784834,0.0034832316,-0.056852203,-0.012945193,-0.050754715,-0.05909069,-0.046398517,0.024399279,-0.02930913,-0.10395105,-0.011431423,0.028347071,0.010558335,0.027873065,-0.109235674,0.032155566,-0.06912663,0.010209797,-0.038717333,0.014557831,-0.07610044,0.024598178,0.016510325,0.04507311,-0.056756962,0.013758646,-0.06844432,-0.006438681,0.090832904,0.051732734,-0.011350843,-0.025747566,-0.043756638,-0.028601611,0.0338011]]"),
		X:              0.6,
		Y:              0.7,
		W:              0.2,
		H:              0.05,
	},
	"1000003-5": Marker{
		ID:             5,
		FileID:         1000003,
		FaceID:         FaceFixtures.Get("unknown").ID,
		SubjectUID:     "",
		MarkerSrc:      SrcImage,
		MarkerType:     MarkerFace,
		MarkerName:     "Corn McCornface",
		LandmarksJSON:  []byte("[{\"name\":\"lp46\",\"x\":-0.10546875,\"y\":-0.045898438,\"h\":0.033203125,\"w\":0.044270832},{\"name\":\"lp46_v\",\"x\":0.11328125,\"y\":0.0126953125,\"h\":0.034179688,\"w\":0.045572918},{\"name\":\"lp44\",\"x\":-0.053385418,\"y\":-0.0546875,\"h\":0.033203125,\"w\":0.044270832},{\"name\":\"lp44_v\",\"x\":0.09375,\"y\":-0.0078125,\"h\":0.033203125,\"w\":0.044270832},{\"name\":\"lp42\",\"x\":-0.015625,\"y\":-0.030273438,\"h\":0.033203125,\"w\":0.044270832},{\"name\":\"lp42_v\",\"x\":0.0546875,\"y\":-0.0087890625,\"h\":0.033203125,\"w\":0.044270832},{\"name\":\"lp38\",\"x\":-0.033854168,\"y\":-0.0087890625,\"h\":0.033203125,\"w\":0.044270832},{\"name\":\"lp38_v\",\"x\":0.037760418,\"y\":0.01171875,\"h\":0.032226562,\"w\":0.04296875},{\"name\":\"lp312\",\"x\":-0.091145836,\"y\":-0.02734375,\"h\":0.033203125,\"w\":0.044270832},{\"name\":\"lp312_v\",\"x\":0.08984375,\"y\":0.021484375,\"h\":0.033203125,\"w\":0.044270832},{\"name\":\"mouth_lp93\",\"x\":-0.026041666,\"y\":0.07714844,\"h\":0.033203125,\"w\":0.044270832},{\"name\":\"mouth_lp84\",\"x\":-0.102864586,\"y\":0.08496094,\"h\":0.033203125,\"w\":0.044270832},{\"name\":\"mouth_lp82\",\"x\":-0.05859375,\"y\":0.12109375,\"h\":0.033203125,\"w\":0.044270832},{\"name\":\"mouth_lp81\",\"x\":-0.045572918,\"y\":0.10058594,\"h\":0.033203125,\"w\":0.044270832},{\"name\":\"lp84\",\"x\":-0.0065104165,\"y\":0.1171875,\"h\":0.033203125,\"w\":0.044270832},{\"name\":\"eye_l\",\"x\":-0.059895832,\"y\":-0.015625,\"h\":0.022460938,\"w\":0.029947916},{\"name\":\"eye_r\",\"x\":0.059895832,\"y\":0.015625,\"h\":0.022460938,\"w\":0.029947916}]"),
		EmbeddingsJSON: []byte("[[-0.01598889,-0.013085627,-0.091266885,0.01090832,0.059863407,-0.018543985,0.06513644,0.072262324,-0.0024608355,0.023480255,-0.034981046,0.028824307,0.031145066,0.020279909,-0.038303424,0.063686624,0.00016155555,-0.040199693,-0.004059554,-0.038183372,0.010284749,0.0393553,-0.013954249,0.014198846,0.047331642,-0.07308497,0.008753774,-0.044199772,0.0351775,-0.009616884,-0.011886778,-0.03133512,-0.008360827,-0.0021931753,-0.0031518617,-0.0007841898,0.0012749507,-0.013682331,-0.0093261255,-0.0646958,0.028137255,-0.051393177,-0.010831488,-0.0019370695,0.026701374,0.08734394,-0.03148508,-0.072140485,-0.008645494,0.03283726,0.025486251,-0.14762828,0.0016828487,-0.049219336,-0.090523295,-0.017858343,-0.0433293,-0.03822806,0.03775215,-0.030626448,-0.005236273,-0.025029438,0.011486794,0.08866543,-0.11626933,-0.012919138,0.011206989,0.029296853,0.029712738,-0.0035943172,-0.0625837,-0.08751456,0.06506425,0.08434424,0.018379156,0.006281598,-0.019832052,0.013404299,0.050819624,-0.0025536602,0.108513094,0.043542076,-0.03385126,-0.013718123,-0.020935653,0.026902547,0.023695294,-0.032848295,0.02122507,-0.06577069,-0.049782418,0.07434279,0.011499832,0.0274455,0.023498816,-0.0024784799,-0.0408338,-0.01835984,-0.07471391,-0.020153865,0.020164149,0.08315784,0.057026222,0.0064989394,-0.030934982,-0.068170995,-0.0447578,-0.08320161,-0.02340753,-0.052578453,-0.014655025,0.011221938,0.020638045,-0.046642996,0.07793852,-0.07613048,-0.0052055195,0.050180323,-0.008933726,-0.006827808,0.046653792,0.046993263,-0.022871166,-0.10476393,-0.019975707,-0.013084017,0.054878037,0.016760936,0.0547656,0.016288247,0.029425414,-0.0039762836,-0.06857062,0.011283167,0.031562343,-0.0301557,0.018833537,0.014527415,0.014445754,0.08944574,-0.0201444,0.07525237,0.00338875,-0.009128363,0.03199888,0.003338322,-0.034410495,0.036284383,0.061515816,-0.04700179,-0.0037704427,-0.014624933,0.00053029,-0.040128306,0.033271074,-0.034320436,-0.087529205,-0.015765699,-0.042281955,0.026603937,-0.03505665,0.0045038764,0.07849869,-0.020796757,-0.027909378,-0.030466244,-0.022126308,0.032068476,0.017038053,0.011572472,-0.023514347,-0.009388009,0.0471232,0.079254,0.038409125,0.09174638,-0.00019651726,0.031568147,0.061691687,0.0035580387,0.01834219,-0.034832783,-0.02472014,-0.0043187817,-0.012822156,-0.006406414,0.04589425,-0.029704314,0.0122809885,-0.003854935,0.034970216,-0.015017675,0.060136527,-0.0018005831,0.122445226,-0.02466541,-0.01479218,0.021981824,0.008459089,-0.024648905,-0.005501727,-0.03819929,0.09229076,-0.00051632634,0.026605012,0.07888228,0.108206324,0.052503213,0.11041478,-0.051830698,0.0057799574,0.06611321,-0.07650234,0.05879657,0.052348997,0.019547585,-0.07638257,-0.044862173,-0.043846298,-0.021011176,-0.01881608,0.107294865,-0.022127088,0.067301944,0.06469245,-0.00538971,-0.014842585,0.008351809,0.004484374,0.040771488,-0.10075861,-0.037090372,0.010828192,-0.014120988,-0.106951684,0.017808948,0.0014778537,-0.052681025,-0.005136359,0.03752494,0.021728097,0.054375395,-0.024591682,-0.028040286,-0.00003123716,-0.009580712,-0.01999454,-0.036310453,0.012808869,0.033562016,-0.048961494,0.049189627,-0.075262696,-0.025279071,-0.017653963,-0.07986348,-0.0058788625,-0.007859457,-0.07787707,-0.0141451955,-0.027267313,0.034545522,-0.066852376,-0.033797234,0.00662154,-0.0042986833,0.06258031,-0.027749743,0.035639953,-0.0860917,-0.021690408,-0.041379962,0.058219258,0.036375474,-0.028576316,0.009679853,0.0540403,-0.00078071933,0.046852123,-0.009272538,0.00014093994,0.08755382,0.029109005,-0.051301096,0.05482741,0.0515714,0.00093064347,-0.0030184423,0.006834386,0.018765705,-0.05297967,0.003963442,0.0058909976,0.0015918964,0.0059644254,-0.010626185,0.016774561,0.019620024,0.012779505,-0.0037311793,0.06090134,0.030533332,-0.0013029866,0.0078083063,-0.07804226,0.055505097,-0.0323863,-0.13843039,0.017691568,0.015463283,-0.0054733357,0.029853273,0.044509612,-0.017058503,-0.009644833,-0.08004177,0.018354492,0.08402592,0.08028074,-0.033064563,0.0071564782,-0.026635528,-0.020762388,-0.09037636,0.07283039,-0.0013951861,0.023556605,-0.025109852,-0.01911825,-0.072672606,0.023367604,0.0066802157,-0.029526467,-0.005831557,-0.056150723,-0.036622144,0.014390345,-0.06641906,0.027680077,-0.0027364918,0.07881769,0.0032321527,0.06692333,-0.023430921,-0.0156005705,-0.012797314,-0.06341876,-0.0075710113,-0.08883436,0.018736875,0.0515824,-0.03050761,-0.038169485,-0.019994874,-0.037897438,0.0030616417,0.016518306,0.017313045,-0.040592685,-0.121986456,0.036341745,-0.055392407,0.06128348,0.09128614,0.085432775,-0.04018598,0.020804984,0.022338325,-0.047893576,-0.04436214,0.010663377,-0.01539266,0.001970492,-0.02548427,-0.0010024207,0.012728738,-0.03458635,0.000458029,-0.07648158,0.017893706,-0.03620278,0.012510285,0.042211026,0.029836254,-0.01023813,-0.014964832,-0.036710255,0.072146155,-0.032623224,-0.054371897,0.059095327,-0.026043909,-0.026475005,-0.03756759,-0.0033844158,0.01665272,0.055144988,0.020163653,0.010291277,0.016448587,-0.08021163,-0.00711534,0.0014388722,0.047059905,0.055735916,0.04716966,0.002357553,0.028611615,0.053914364,-0.022333615,0.01227299,-0.006376796,-0.020971471,-0.031313244,0.004976049,0.027839795,0.032628387,0.014910606,-0.019821445,-0.023739582,-0.061071664,-0.03563204,0.04504174,-0.043931577,-0.09461471,0.035146907,0.018801821,-0.023486922,-0.014275421,-0.04465509,-0.02062559,0.0049236626,-0.018532282,0.0329802,0.08521481,-0.025042368,-0.0031619244,-0.012923802,0.062199675,0.008717194,-0.06812108,-0.059829578,0.012077271,0.04268468,-0.029709337,0.058341388,-0.023695359,0.026195621,0.089156136,-0.05210246,-0.006754805,-0.08180936,-0.016863013,0.01305001,-0.06608525,0.060481545,-0.008385101,-0.022109058,-0.060189016,0.029254012,0.030882897,-0.009892429,0.04769517,0.010766996,-0.032422658,-0.010280704,-0.035042368,-0.046891414,-0.021463934,-0.022753375,0.011999883,-0.027142627,-0.00023438907,0.016639061,-0.022108192,0.06750844,0.031127596,0.062595494,0.035271265,-0.00848079,0.0054323506,-0.038877886,-0.09272652,-0.028250217,-0.008661228,-0.08010225,-0.066937625,-0.022481173,0.023471648,0.0064566983,-0.0010692414,-0.11465305,0.041438285,-0.08444645,0.004585747,-0.028296055,0.04911384,-0.0049776305,0.043981623,-0.010605184,0.07764796,-0.022333274,0.050281037,-0.06543124,-0.00412494,0.08933364,0.043259624,0.0007392553,-0.0055968673,-0.041995358,-0.01275004,0.017657083]]"),
		X:              0.2,
		Y:              0.3,
		W:              0.1,
		H:              0.1,
	},
	"1000003-6": Marker{
		ID:             6,
		FileID:         1000003,
		FaceID:         FaceFixtures.Get("known").ID,
		SubjectUID:     "",
		MarkerSrc:      SrcImage,
		MarkerType:     MarkerFace,
		MarkerName:     "",
		LandmarksJSON:  []byte("[{\"name\":\"lp46\",\"x\":-0.08359375,\"y\":-0.027083334,\"h\":0.044791665,\"w\":0.03359375},{\"name\":\"lp46_v\",\"x\":0.08671875,\"y\":-0.009375,\"h\":0.045833334,\"w\":0.034375},{\"name\":\"lp44\",\"x\":-0.0546875,\"y\":-0.048958335,\"h\":0.044791665,\"w\":0.03359375},{\"name\":\"lp44_v\",\"x\":0.06328125,\"y\":-0.033333335,\"h\":0.045833334,\"w\":0.034375},{\"name\":\"lp42\",\"x\":-0.021875,\"y\":-0.03125,\"h\":0.044791665,\"w\":0.03359375},{\"name\":\"lp42_v\",\"x\":0.03203125,\"y\":-0.025,\"h\":0.045833334,\"w\":0.034375},{\"name\":\"lp38\",\"x\":-0.0265625,\"h\":0.044791665,\"w\":0.03359375},{\"name\":\"lp38_v\",\"x\":0.03125,\"y\":0.0052083335,\"h\":0.044791665,\"w\":0.03359375},{\"name\":\"lp312\",\"x\":-0.06796875,\"y\":-0.008333334,\"h\":0.045833334,\"w\":0.034375},{\"name\":\"lp312_v\",\"x\":0.06953125,\"y\":0.008333334,\"h\":0.044791665,\"w\":0.03359375},{\"name\":\"mouth_lp93\",\"x\":-0.00703125,\"y\":0.09375,\"h\":0.044791665,\"w\":0.03359375},{\"name\":\"mouth_lp84\",\"x\":-0.04921875,\"y\":0.128125,\"h\":0.045833334,\"w\":0.034375},{\"name\":\"mouth_lp82\",\"x\":-0.01328125,\"y\":0.16145833,\"h\":0.045833334,\"w\":0.034375},{\"name\":\"mouth_lp81\",\"x\":-0.0078125,\"y\":0.13333334,\"h\":0.044791665,\"w\":0.03359375},{\"name\":\"lp84\",\"x\":0.034375,\"y\":0.14479166,\"h\":0.044791665,\"w\":0.03359375},{\"name\":\"eye_l\",\"x\":-0.0484375,\"y\":-0.004166667,\"h\":0.030208332,\"w\":0.02265625},{\"name\":\"eye_r\",\"x\":0.0484375,\"y\":0.0052083335,\"h\":0.030208332,\"w\":0.02265625}]"),
		EmbeddingsJSON: []byte("[[0.019231493,-0.028809275,-0.083006255,-0.015598502,0.08550906,0.001886255,0.09019353,0.07551488,-0.011814484,0.03680722,-0.08332401,0.014950869,0.055766843,-0.0073407963,-0.0552993,0.05847239,0.026178556,-0.045643847,0.012550894,-0.012022383,-0.0040185535,0.00023647904,-0.01580684,-0.008704283,0.04575994,-0.046812356,0.056398943,-0.026091263,0.059522673,-0.044217024,0.014755385,-0.01350486,-0.049488768,-0.03139871,-0.006726978,-0.02347069,-0.0059445584,-0.004682308,-0.057403754,-0.06537466,0.013326172,-0.009667708,0.022370687,0.0037015954,0.03744496,0.052890837,-0.0077360696,-0.049944617,-0.03868134,0.001521219,0.03840492,-0.10928545,0.023024736,-0.055707198,-0.13260484,0.009903039,-0.04250921,-0.0040567834,0.03343564,-0.01785736,0.0043026204,-0.031062575,-0.0019649328,0.06487235,-0.14464019,-0.017717961,-0.0033534314,0.029505186,0.008849258,-0.0026131037,-0.06479913,-0.111862205,0.05469594,0.049985956,-0.00067700783,0.003068928,0.0018148758,0.0073374105,0.025748348,-0.0424614,0.062650666,0.058194485,-0.04309207,-0.020790769,-0.030982763,0.008360668,0.01289988,-0.019662105,0.0122521445,-0.00342255,-0.056044392,0.034414552,0.04604621,0.0074918787,0.033526078,-0.036619328,-0.047758896,-0.032501936,-0.08068566,-0.02964604,0.04137439,0.06888022,0.04018322,0.0023792675,-0.026837967,-0.049688686,-0.057930365,-0.064863205,0.004485477,-0.026958624,-0.025907256,0.009216111,-0.014622554,-0.037213538,0.078393415,-0.054682203,-0.009757617,0.03503295,0.027951613,-0.0014038666,0.06851987,0.020453943,-0.00996363,-0.12156495,-0.017301193,-0.032558206,0.07816977,0.029640608,0.03150378,0.047710016,-0.009032255,0.013137696,-0.056541067,-0.0075266357,0.007981631,-0.025004586,0.030189874,-0.007878598,0.03279407,0.059711676,-0.003458093,0.01758815,-0.010460049,0.012645757,0.006580964,-0.019848932,0.02724201,0.001597322,0.0998137,-0.03012735,0.011577833,0.0028034616,0.029448563,-0.011648439,0.05161503,-0.028912308,-0.07435124,-0.0033618107,0.030619241,0.04033193,-0.0049585714,-0.016229536,0.05426284,-0.031509526,0.014953531,-0.025750436,0.000924462,-0.008974987,0.03417905,0.0011432022,0.007551494,0.001544253,0.016510574,0.044189014,0.054653972,0.0811879,-0.0036358214,0.035907324,0.047312576,-0.0068604983,-0.029337948,0.027281888,-0.08003064,-0.054304074,0.0073654098,0.019096468,0.004265153,-0.03670545,0.0049256124,-0.00017001168,0.04502574,-0.00074317795,0.024450991,-0.052183975,0.15243798,-0.010411264,-0.016080985,0.017625313,0.022576308,-0.053998422,0.0023750497,-0.057889163,0.056488033,-0.025041293,0.011416959,0.08541408,0.08386572,0.0396809,0.0920811,-0.048865005,-0.026257817,0.032705985,-0.053586084,0.032277677,0.07027237,0.016478207,-0.08976395,-0.006354579,-0.038090594,0.024028458,0.013410826,0.08724634,-0.028266877,0.07450935,0.03037237,-0.018165091,-0.029061696,-0.0006272013,-0.005710233,0.037058197,-0.085284434,-0.04267883,0.029252399,-0.033832725,-0.07989745,0.025662571,0.019242961,-0.10760675,0.0021055646,0.0021679339,0.0382236,0.01998776,-0.040100973,-0.041129857,0.0025446033,-0.03981458,-0.031129645,-0.061818115,-0.031583495,0.08146585,-0.042177603,0.05506262,-0.045823902,-0.031090494,-0.04478884,-0.06997962,-0.0024934823,0.0020288285,-0.074601755,-0.029107075,-0.03953502,0.0551208,-0.05685647,-0.0010502128,0.04371499,-0.031777892,0.030995583,-0.056923013,0.04047056,-0.05058555,-0.020007737,-0.0034168877,0.040992126,-0.0065717557,-0.060389657,-0.015070164,0.0300555,-0.0049498156,0.035197794,-0.010181344,-0.01548949,0.09214268,0.06594641,-0.04095856,0.031043377,0.016846823,0.011824804,0.010486359,0.006952729,-0.0423542,-0.038642637,0.037917823,-0.042705245,0.002531653,0.049575835,-0.008132699,0.060207773,0.050799962,0.029537434,-0.011404661,0.07556842,0.044129632,-0.025151744,-0.04391785,-0.057073284,0.066147126,-0.03833207,-0.11469866,0.018607577,0.03958792,0.005069568,-0.0022620235,-0.06963503,-0.027407918,-0.01658652,-0.08928287,0.04213307,0.0653322,0.070556566,-0.08890351,0.05253341,-0.03952343,-0.03277939,-0.07603576,0.023622843,0.01869451,0.0012659469,-0.016868249,-0.048114307,-0.12678534,-0.023940234,0.03518056,0.030265998,-0.027942419,-0.030615386,-0.027904578,0.04184872,-0.06871633,0.018810445,-0.0050423085,0.01902196,0.036709674,0.05114107,-0.008238005,-0.033530205,0.009237725,-0.019793255,-0.011098562,-0.040874466,0.028055958,0.052516278,0.002427102,-0.018204305,-0.029648917,-0.071319304,-0.0468887,0.061347924,0.031454504,-0.031552624,-0.08102158,0.074395515,-0.048840746,0.030817837,0.07313089,0.062438965,-0.00383427,0.05508265,-0.0077238525,0.026816545,-0.050590646,-0.02294128,0.0017293892,-0.018405396,0.03635994,-0.0025451558,-0.005002223,-0.059726283,0.008561757,-0.05424462,-0.009393793,-0.040830627,0.02030258,0.003973316,0.020119876,-0.017317869,0.01453977,-0.025234057,0.072000384,-0.0413773,-0.111050114,0.06299958,0.00016407811,-0.050199028,-0.013726295,0.0100843245,0.022901619,0.056635447,0.038121402,0.024791898,-0.011317786,-0.059741378,-0.0004985886,-0.0129778,0.04761788,0.018754557,0.035193384,-0.021030819,0.042050865,0.08013355,-0.042123444,0.065898865,0.03751851,-0.049639545,-0.03844303,0.006924326,0.022647297,-0.008048924,0.015995622,-0.033804823,0.007260562,-0.046132274,-0.0064429105,0.031663842,-0.006572463,-0.06781134,0.013187006,0.013765757,-0.03454214,-0.015666338,-0.0023530773,-0.07751217,-0.008714079,-0.010440672,0.026577305,0.066843055,-0.00440322,-0.0071120844,0.008546605,0.087584786,-0.027872834,-0.043215286,-0.0657154,0.042205643,0.003758622,-0.029912153,0.020607445,-0.0034470167,0.040666997,0.081489064,-0.044269655,-0.006542095,-0.054021314,-0.029242665,-0.027248744,-0.02227558,0.082040176,-0.030761424,-0.023510806,-0.06973938,-0.0032560013,0.055813447,0.03283226,0.068810284,0.029060816,-0.03417918,0.004436392,-0.018858656,-0.0056046643,-0.034613956,-0.024074866,0.029658342,-0.023564866,-0.011503043,0.04425076,-0.017220033,0.094600976,0.0664404,0.0784834,0.0034832316,-0.056852203,-0.012945193,-0.050754715,-0.05909069,-0.046398517,0.024399279,-0.02930913,-0.10395105,-0.011431423,0.028347071,0.010558335,0.027873065,-0.109235674,0.032155566,-0.06912663,0.010209797,-0.038717333,0.014557831,-0.07610044,0.024598178,0.016510325,0.04507311,-0.056756962,0.013758646,-0.06844432,-0.006438681,0.090832904,0.051732734,-0.011350843,-0.025747566,-0.043756638,-0.028601611,0.0338011]]"),
		X:              0.494531,
		Y:              0.282292,
		W:              0.285937,
		H:              0.38125,
	},
}

// CreateMarkerFixtures inserts known entities into the database for testing.
func CreateMarkerFixtures() {
	for _, entity := range MarkerFixtures {
		Db().Create(&entity)
	}
}
