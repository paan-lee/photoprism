package entity

import (
	"testing"

	"github.com/photoprism/photoprism/internal/face"

	"github.com/stretchr/testify/assert"
)

func TestFace_TableName(t *testing.T) {
	m := &Face{}
	assert.Contains(t, m.TableName(), "faces")
}

func TestFace_Match(t *testing.T) {
	t.Run("1000003-4", func(t *testing.T) {
		m := FaceFixtures.Get("joe-biden")
		match, dist := m.Match(MarkerFixtures.Pointer("1000003-4").Embeddings())

		assert.True(t, match)
		assert.Greater(t, dist, 1.31)
		assert.Less(t, dist, 1.32)
	})

	t.Run("1000003-6", func(t *testing.T) {
		m := FaceFixtures.Get("joe-biden")
		match, dist := m.Match(MarkerFixtures.Pointer("1000003-6").Embeddings())

		assert.True(t, match)
		assert.Greater(t, dist, 1.27)
		assert.Less(t, dist, 1.28)
	})

	t.Run("len(embeddings) == 0", func(t *testing.T) {
		m := FaceFixtures.Get("joe-biden")
		match, dist := m.Match(face.Embeddings{})

		assert.False(t, match)
		assert.Equal(t, dist, float64(-1))
	})
	t.Run("len(efacEmbeddings) == 0", func(t *testing.T) {
		m := NewFace("12345", SrcAuto, face.Embeddings{})
		match, dist := m.Match(MarkerFixtures.Pointer("1000003-6").Embeddings())

		assert.False(t, match)
		assert.Equal(t, dist, float64(-1))
	})
	t.Run("jane doe- no match", func(t *testing.T) {
		m := FaceFixtures.Get("jane-doe")
		match, _ := m.Match(MarkerFixtures.Pointer("1000003-5").Embeddings())

		assert.False(t, match)
	})
}

func TestFace_ResolveCollision(t *testing.T) {
	t.Run("collision", func(t *testing.T) {
		m := FaceFixtures.Get("joe-biden")

		assert.Zero(t, m.Collisions)
		assert.Zero(t, m.CollisionRadius)

		if reported, err := m.ResolveCollision(MarkerFixtures.Pointer("1000003-4").Embeddings()); err != nil {
			t.Fatal(err)
		} else {
			assert.True(t, reported)
		}

		// Number of collisions must have increased by one.
		assert.Equal(t, 1, m.Collisions)

		// Actual distance is ~1.314040
		assert.Greater(t, m.CollisionRadius, 1.2)
		assert.Less(t, m.CollisionRadius, 1.314)

		if reported, err := m.ResolveCollision(MarkerFixtures.Pointer("1000003-6").Embeddings()); err != nil {
			t.Fatal(err)
		} else {
			assert.True(t, reported)
		}

		// Number of collisions must not have increased.
		assert.Equal(t, 2, m.Collisions)

		// Actual distance is ~1.272604
		assert.Greater(t, m.CollisionRadius, 1.1)
		assert.Less(t, m.CollisionRadius, 1.272)
	})
	t.Run("subject id empty", func(t *testing.T) {
		m := NewFace("", SrcAuto, face.Embeddings{})
		if reported, err := m.ResolveCollision(MarkerFixtures.Pointer("1000003-4").Embeddings()); err != nil {
			t.Fatal(err)
		} else {
			assert.False(t, reported)
		}
	})
	t.Run("invalid face id", func(t *testing.T) {
		m := NewFace("123", SrcAuto, face.Embeddings{})
		m.ID = ""
		if reported, err := m.ResolveCollision(MarkerFixtures.Pointer("1000003-4").Embeddings()); err == nil {
			t.Fatal(err)
		} else {
			assert.False(t, reported)
			assert.Equal(t, "invalid face id", err.Error())
		}
	})
	t.Run("embedding empty", func(t *testing.T) {
		m := NewFace("123", SrcAuto, face.Embeddings{})
		m.EmbeddingJSON = []byte("")
		if reported, err := m.ResolveCollision(MarkerFixtures.Pointer("1000003-4").Embeddings()); err == nil {
			t.Fatal(err)
		} else {
			assert.False(t, reported)
			assert.Equal(t, "embedding must not be empty", err.Error())
		}
	})
}

func TestFace_ReviseMatches(t *testing.T) {
	m := FaceFixtures.Get("joe-biden")
	removed, err := m.ReviseMatches()

	if err != nil {
		t.Fatal(err)
	}

	assert.Empty(t, removed)
}

func TestNewFace(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		marker := MarkerFixtures.Get("1000003-4")
		e := marker.Embeddings()

		r := NewFace("123", SrcAuto, e)
		assert.Equal(t, "", r.FaceSrc)
		assert.Equal(t, "123", r.SubjUID)
	})
}

func TestFace_Unsuitable(t *testing.T) {
	t.Run("True", func(t *testing.T) {
		m := FaceFixtures.Get("joe-biden")
		assert.False(t, m.OmitMatch())
	})
	t.Run("False", func(t *testing.T) {
		m := NewFace("", SrcImage, face.Embeddings{{-0.00959064718335867, 0.03787063807249069, -0.0030881548300385475, 0.02789853885769844, 0.017454572021961212, 0.0396987721323967, -0.03091704286634922, 0.005318029318004847, 0.021617550402879715, -0.08214963972568512, -0.003952134400606155, 0.0269720908254385, 0.048880551010370255, -0.03537372127175331, -0.042236171662807465, 0.021553633734583855, 0.03937383368611336, 0.01815507560968399, 0.08373168110847473, -0.11838400363922119, -0.038254253566265106, -0.04993032291531563, 0.07148619741201401, 0.006384310312569141, 0.05344310402870178, -0.027579499408602715, 0.021648988127708435, -0.07013172656297684, -0.06400937587022781, 0.10622639954090118, -0.01507984846830368, -0.02844894863665104, -0.013048898428678513, -0.03571505844593048, -0.022063886746764183, 0.022826166823506355, 0.01703103445470333, 0.00679031852632761, -0.09583312273025513, 0.03446732088923454, -0.045221585780382156, 0.03292521834373474, -0.012820744886994362, 0.06122862547636032, 0.01973198726773262, -0.013975882902741432, 0.027514882385730743, 0.12478502094745636, -0.09630053490400314, -0.008597812615334988, -0.019534612074494362, 0.03927983343601227, 0.04311678186058998, 0.025297729298472404, -0.035719674080610275, 0.05421024188399315, 0.07541341334581375, 0.040334682911634445, -0.0632546916604042, -0.004164006095379591, 0.027950556948781013, 0.017827920615673065, 0.02774866297841072, -0.025094853714108467, 0.00012262807285878807, 0.04165732488036156, -0.03155842795968056, 0.03801475837826729, 0.0031508952379226685, -0.011753040365874767, 0.06262513995170593, 0.05895991623401642, -0.02384188584983349, -0.025149181485176086, -0.016906173899769783, -0.03138834610581398, -0.06759334355592728, 0.018074069172143936, 0.028748946264386177, 0.03350280225276947, 0.001738330232910812, -0.035873714834451675, 0.0050230612978339195, -0.005394259933382273, -0.035111431032419205, 0.005703517701476812, -0.060869812965393066, 0.044046416878700256, 0.05451945215463638, -0.0012109529925510287, 0.04929054155945778, 0.03312966600060463, -0.02503111958503723, -0.0699458047747612, 0.09152142703533173, -0.035196661949157715, -0.02000804804265499, 0.003603762947022915, -0.0549810416996479, 0.041149843484163284, 0.019640415906906128, -0.06913350522518158, -0.08494774252176285, 0.047828249633312225, 0.011485084891319275, 0.11441357433795929, 0.012079037725925446, 0.026444999501109123, 0.008605830371379852, -0.014796323142945766, 0.042191699147224426, 0.0360623262822628, -0.01067506056278944, -0.02117612026631832, -0.0003311904729343951, 0.020912105217576027, 0.02051572874188423, 0.04119933396577835, 0.011461400426924229, 0.02468070574104786, -0.030830683186650276, -0.024522947147488594, 0.07760800421237946, -0.044838037341833115, 0.007875975221395493, 0.03662760183215141, -0.031315844506025314, 0.028968002647161484, -0.007360775955021381, -0.052097514271736145, 0.004892056342214346, 0.0051552411168813705, 0.058972474187612534, -0.05307154729962349, -0.02330617979168892, 0.0560041144490242, -0.06173492223024368, 0.00004632262425730005, 0.007912986911833286, 0.0031768144108355045, -0.08211413770914078, -0.02641596458852291, -0.07240095734596252, -0.04998013749718666, 0.016048355028033257, -0.023686233907938004, 0.08416120707988739, 0.002466161735355854, 0.0017551603959873319, 0.000651281327009201, 0.018105899915099144, -0.05974912270903587, -0.03980677202343941, 0.019075721502304077, 0.0014616637490689754, 0.06682229787111282, 0.02257758192718029, 0.04021807014942169, 0.09144134074449539, 0.020396307110786438, 0.055604636669158936, 0.026022544130682945, -0.03050902672111988, 0.011569516733288765, -0.014519683085381985, 0.0038184933364391327, -0.03115340694785118, 0.029596896842122078, -0.055038318037986755, -0.005584381986409426, -0.015937503427267075, -0.01591162569820881, 0.034234486520290375, 0.010233158245682716, 0.0364360548555851, 0.02957785315811634, 0.038372594863176346, -0.04782934859395027, -0.03462134674191475, -0.0432763509452343, -0.041607096791267395, 0.019871780648827553, -0.026665959507226944, 0.046689242124557495, 0.020541366189718246, 0.03362491726875305, 0.04561452195048332, 0.12613892555236816, 0.02306310087442398, 0.0048497817479074, -0.027223020792007446, -0.0762500986456871, 0.06465625762939453, -0.020680397748947144, -0.02472679689526558, -0.0469549298286438, 0.05494922026991844, 0.011157477274537086, -0.05097919702529907, 0.05126889795064926, 0.03758222982287407, -0.06554574519395828, 0.00288044149056077, 0.014015591703355312, 0.013589163310825825, 0.03634551167488098, 0.0031862170435488224, -0.03541851416230202, -0.011984468437731266, -0.04591989517211914, -0.04950973764061928, 0.014266318641602993, 0.014613134786486626, 0.004269343335181475, 0.0013365329941734672, -0.010044350288808346, 0.025745976716279984, 0.029322613030672073, 0.08641400188207626, 0.00042273724102415144, -0.1199660375714302, -0.11129316687583923, -0.03984867036342621, -0.05681384354829788, 0.009998883120715618, 0.030147377401590347, -0.0286977831274271, -0.0003513149276841432, -0.08627857267856598, -0.023915421217679977, 0.025925707072019577, 0.08490575850009918, 0.031879108399152756, -0.0023629055358469486, 0.0480312779545784, 0.0021763548720628023, 0.020024623721837997, -0.01996619999408722, -0.001396739506162703, -0.026282500475645065, -0.040633674710989, -0.019956767559051514, 0.004316484089940786, -0.031683146953582764, 0.06379353255033493, 0.03608919307589531, 0.008245682343840599, 0.02868475206196308, -0.0009207205730490386, -0.0003780983679462224, 0.02880168706178665, -0.04014896973967552, 0.017292466014623642, 0.049382057040929794, -0.015038374811410904, 0.024192562326788902, 0.03517518192529678, 0.019119804725050926, 0.021942559629678726, 0.07587131857872009, 0.0005678452434949577, -0.04390380531549454, 0.030292486771941185, 0.042298778891563416, -0.06521622836589813, 0.02252770960330963, -0.00466647045686841, -0.024906277656555176, -0.026186272501945496, 0.07858474552631378, -0.05505937710404396, -0.0008577461121603847, 0.00968341063708067, -0.036743391305208206, -0.08478125929832458, -0.025725962594151497, 0.07145383208990097, 0.029407603666186333, -0.0001950680452864617, -0.1292036920785904, 0.02494245208799839, -0.008491290733218193, 0.050918228924274445, -0.011559431441128254, -0.04706485942006111, -0.013783150352537632, 0.009277299977838993, -0.07522283494472504, -0.036186907440423965, -0.06634241342544556, 0.010219116695225239, -0.08408123254776001, -0.014987781643867493, 0.010251465253531933, -0.01592072658240795, -0.035617098212242126, 0.020554568618535995, 0.05061344802379608, 0.0494505874812603, 0.02590356394648552, 0.01528799906373024, -0.00029076842474751174, 0.02353300340473652, 0.0015167297096922994, 0.05400843173265457, 0.04550565034151077, -0.04259566590189934, -0.0060416329652071, -0.00677477428689599, 0.05933074653148651, -0.005193949677050114, 0.014253835193812847, 0.042284123599529266, 0.06695422530174255, 0.04029611125588417, 0.015430709347128868, -0.06947603821754456, 0.0339425727725029, -0.06005615368485451, 0.01404648832976818, 0.06008269265294075, -0.011060234159231186, -0.04977267608046532, 0.05691606178879738, -0.013345426879823208, 0.10078923404216766, 0.08031554520130157, 0.0425117127597332, 0.09008562564849854, 0.04609135910868645, 0.06102297827601433, 0.022890515625476837, -0.03089219331741333, 0.0332498624920845, -0.031279049813747406, -0.009156256914138794, 0.027170570567250252, 0.04901871085166931, 0.07207565009593964, 0.04074881598353386, -0.027857864275574684, -0.025717025622725487, 0.032386474311351776, 0.036552079021930695, -0.055537834763526917, -0.0229702889919281, 0.03349658474326134, 0.03683074936270714, 0.015133108012378216, -0.0632123202085495, -0.030310358852148056, 0.09408748149871826, 0.011012745089828968, -0.10027626156806946, -0.056098587810993195, 0.007266550324857235, 0.09435073286294937, -0.005252359434962273, 0.0414881557226181, -0.07797796279191971, 0.0054626669734716415, -0.027152489870786667, -0.06476820260286331, -0.04554128646850586, -0.020997364073991776, 0.03704288229346275, -0.0041465735994279385, -0.08224689960479736, 0.019587524235248566, 0.05182863399386406, -0.09750733524560928, 0.012806789949536324, 0.014560981653630733, -0.012063717469573021, 0.10477723181247711, 0.04364655539393425, 0.05573931708931923, -0.08249012380838394, 0.002664536237716675, 0.016965137794613838, 0.016157248988747597, -0.07265286147594452, -0.0025825295597314835, -0.011157424189150333, -0.053293049335479736, 0.01613083854317665, 0.003192639909684658, -0.02518875151872635, 0.025411557406187057, -0.04756153002381325, -0.008369989693164825, 0.0018538516014814377, -0.001305201556533575, 0.006403622217476368, 0.020627789199352264, -0.024054545909166336, 0.05217380076646805, 0.0469573549926281, 0.01885838247835636, 0.020833401009440422, -0.04654202610254288, 0.044648706912994385, -0.004453012719750404, -0.021127738058567047, 0.007881376892328262, -0.04722931608557701, -0.009467313066124916, 0.013864696025848389, 0.014279618859291077, 0.01670973189175129, 0.006757605355232954, 0.03243840113282204, -0.08637244999408722, 0.014409483410418034, 0.014930488541722298, -0.021012697368860245, -0.00746690621599555, 0.04036633297801018, 0.0766197144985199, -0.002925584791228175, -0.037694621831178665, 0.01753336563706398, -0.0129204411059618, 0.058751046657562256, -0.003414733335375786, 0.009327893145382404, 0.006946941372007132, 0.036547087132930756, 0.01600072905421257, 0.027991879731416702, -0.024807672947645187, 0.013996168039739132, -0.024033015593886375, 0.020035816356539726, -0.06689176708459854, -0.021769963204860687, 0.019834108650684357, 0.007396597880870104, -0.03514741361141205, -0.038449011743068695, -0.0027228370308876038, -0.060723625123500824, 0.05235403776168823, 0.005773501005023718, 0.022514579817652702, 0.03794749826192856, -0.06979167461395264, -0.0036482769064605236, -0.011913052760064602, -0.01920865662395954, -0.04111270606517792, 0.05357895419001579, -0.023412834852933884, -0.0893779918551445, -0.02306830696761608, -0.03236269950866699, 0.007966117933392525, 0.10357413440942764, 0.02653438411653042, 0.004998756106942892, -0.015604768879711628, -0.022902334108948708, -0.10633908212184906, 0.03903093561530113, 0.05978463217616081, -0.011735391803085804, -0.06194228678941727, 0.03223072364926338, 0.04556537792086601, 0.007720542140305042, 0.039454445242881775, -0.04189905524253845, -0.004674337804317474, -0.01275805663317442, -0.12497187405824661, -0.07940814644098282, 0.023411696776747704, 0.02147858962416649, -0.03503002971410751, 0.016921473667025566, -0.016184881329536438, -0.045962586998939514, 0.08095240592956543, -0.004070675931870937, -0.05266023054718971, 0.13639050722122192, -0.02151007391512394, 0.006739250384271145, 0.03182916343212128, -0.027000118046998978, -0.0030197608284652233, 0.031326163560152054, -0.10159225016832352, -0.06630226224660873, 0.0699416846036911, 0.01672203093767166, -0.04788779094815254, 0.039929479360580444, 0.027769070118665695, 0.01937052048742771, -0.06442618370056152, -0.06701736897230148, 0.039595261216163635, 0.05279085412621498, 0.007269475143402815, 0.06969842314720154, 0.048928432166576385, 0.0164470374584198, -0.014216633513569832, -0.015720434486865997, -0.007112122140824795, -0.10834096372127533}})
		assert.True(t, m.OmitMatch())
	})
}

func TestFace_SetEmbeddings(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		marker := MarkerFixtures.Get("1000003-4")
		e := marker.Embeddings()
		m := FaceFixtures.Get("joe-biden")
		assert.NotEqual(t, e[0][0], m.Embedding()[0])

		err := m.SetEmbeddings(e)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, e[0][0], m.Embedding()[0])
	})
}

func TestFace_Embedding(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		m := FaceFixtures.Get("joe-biden")

		assert.Equal(t, 0.10730543085474682, m.Embedding()[0])
	})
	t.Run("empty embedding", func(t *testing.T) {
		m := NewFace("12345", SrcAuto, face.Embeddings{})
		m.EmbeddingJSON = []byte("")

		assert.Empty(t, m.Embedding())
	})
	t.Run("invalid embedding json", func(t *testing.T) {
		m := NewFace("12345", SrcAuto, face.Embeddings{})
		m.EmbeddingJSON = []byte("[false]")

		assert.Equal(t, float64(0), m.Embedding()[0])
	})
}

func TestFace_UpdateMatchTime(t *testing.T) {
	m := NewFace("12345", SrcAuto, face.Embeddings{})
	initialMatchTime := m.MatchedAt
	assert.Equal(t, initialMatchTime, m.MatchedAt)
	m.Matched()
	assert.NotEqual(t, initialMatchTime, m.MatchedAt)
}

func TestFace_Save(t *testing.T) {
	m := NewFace("12345fde", SrcAuto, face.Embeddings{face.Embedding{1}, face.Embedding{2}})
	assert.Nil(t, FindFace(m.ID))
	m.Save()
	assert.NotNil(t, FindFace(m.ID))
	assert.Equal(t, "12345fde", FindFace(m.ID).SubjUID)
}

func TestFace_Update(t *testing.T) {
	m := NewFace("12345fdef", SrcAuto, face.Embeddings{face.Embedding{8}, face.Embedding{16}})
	assert.Nil(t, FindFace(m.ID))
	m.Save()
	assert.NotNil(t, FindFace(m.ID))
	assert.Equal(t, "12345fdef", FindFace(m.ID).SubjUID)

	m2 := FindFace(m.ID)
	m2.Update("SubjUID", "new")
	assert.Equal(t, "new", FindFace(m.ID).SubjUID)
}

func TestFace_RefreshPhotos(t *testing.T) {
	f := FaceFixtures.Get("joe-biden")

	if err := f.RefreshPhotos(); err != nil {
		t.Fatal(err)
	}
}

func TestFirstOrCreateFace(t *testing.T) {
	t.Run("create new face", func(t *testing.T) {
		m := NewFace("12345unique", SrcAuto, face.Embeddings{face.Embedding{99}, face.Embedding{2}})
		r := FirstOrCreateFace(m)
		assert.Equal(t, "12345unique", r.SubjUID)
	})
	t.Run("return existing entity", func(t *testing.T) {
		m := FaceFixtures.Pointer("joe-biden")
		r := FirstOrCreateFace(m)
		assert.Equal(t, "jqy3y652h8njw0sx", r.SubjUID)
		assert.Equal(t, 33, r.Samples)
	})
}

func TestFindFace(t *testing.T) {
	t.Run("existing face", func(t *testing.T) {
		assert.NotNil(t, FindFace("VF7ANLDET2BKZNT4VQWJMMC6HBEFDOG7"))
		assert.Equal(t, 3, FindFace("VF7ANLDET2BKZNT4VQWJMMC6HBEFDOG7").Samples)
	})
	t.Run("empty id", func(t *testing.T) {
		assert.Nil(t, FindFace(""))
	})
}

func TestFace_HideAndShow(t *testing.T) {
	f := FaceFixtures.Get("joe-biden")

	if err := f.Hide(); err != nil {
		t.Fatal(err)
	} else if err = f.Show(); err != nil {
		t.Fatal(err)
	}
}
