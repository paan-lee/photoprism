package entity

import (
	"time"
)

var FileFixtures = map[string]File{
	"exampleFileName.jpg": {
		ID:              1000000,
		Photo:           PhotoFixtures.Pointer("19800101_000002_D640C559"),
		PhotoID:         PhotoFixtures.Pointer("19800101_000002_D640C559").ID,
		PhotoUID:        PhotoFixtures.Pointer("19800101_000002_D640C559").PhotoUID,
		FileUID:         "ft8es39w45bnlqdw",
		FileName:        "exampleFileName.jpg",
		FileRoot:        RootOriginals,
		OriginalName:    "exampleFileNameOriginal.jpg",
		FileHash:        "2cad9168fa6acc5c5c2965ddf6ec465ca42fd818",
		ModTime:         time.Date(2020, 3, 6, 2, 6, 51, 0, time.UTC).Unix(),
		FileSize:        4278906,
		FileType:        "jpg",
		FileMime:        "image/jpg",
		FilePrimary:     true,
		FileSidecar:     false,
		FileVideo:       false,
		FileMissing:     false,
		FilePortrait:    true,
		FileWidth:       3648,
		FileHeight:      2736,
		FileOrientation: 0,
		FileAspectRatio: 1.33333,
		FileMainColor:   "green",
		FileColors:      "929299991",
		FileLuminance:   "8836BD496",
		FileDiff:        968,
		FileChroma:      25,
		FileError:       "",
		Share: []FileShare{
			FileShareFixtures.Get("FileShare1", 0, 0, ""),
			FileShareFixtures.Get("FileShare2", 0, 0, ""),
		},
		Sync:      []FileSync{},
		CreatedAt: time.Date(2020, 3, 6, 2, 6, 51, 0, time.UTC),
		CreatedIn: 2,
		UpdatedAt: time.Date(2020, 3, 28, 14, 6, 0, 0, time.UTC),
		UpdatedIn: 0,
		DeletedAt: nil,
	},
	"exampleDNGFile.dng": {
		ID:              1000001,
		Photo:           PhotoFixtures.Pointer("Photo01"),
		PhotoID:         PhotoFixtures.Pointer("Photo01").ID,
		PhotoUID:        PhotoFixtures.Pointer("Photo01").PhotoUID,
		FileUID:         "ft9es39w45bnlqdw",
		FileName:        "exampleDNGFile.dng",
		FileRoot:        RootOriginals,
		OriginalName:    "exampleDNGFile.dng",
		FileHash:        "3cad9168fa6acc5c5c2965ddf6ec465ca42fd818",
		ModTime:         time.Date(2019, 3, 6, 2, 6, 51, 0, time.UTC).Unix(),
		FileSize:        661858,
		FileType:        "dng",
		FileMime:        "image/dng",
		FilePrimary:     false,
		FileSidecar:     false,
		FileVideo:       false,
		FileMissing:     false,
		FilePortrait:    false,
		FileWidth:       1200,
		FileHeight:      1600,
		FileOrientation: 6,
		FileAspectRatio: 0.75,
		FileMainColor:   "gold",
		FileColors:      "5552E2222",
		FileLuminance:   "444428399",
		FileDiff:        747,
		FileChroma:      12,
		FileError:       "",
		Share:           []FileShare{},
		Sync:            []FileSync{},
		CreatedAt:       time.Date(2019, 3, 6, 2, 6, 51, 0, time.UTC),
		CreatedIn:       2,
		UpdatedAt:       time.Date(2020, 3, 28, 14, 6, 0, 0, time.UTC),
		UpdatedIn:       0,
		DeletedAt:       nil,
	},
	"exampleXmpFile.xmp": {
		ID:              1000002,
		Photo:           PhotoFixtures.Pointer("Photo01"),
		PhotoID:         PhotoFixtures.Pointer("Photo01").ID,
		PhotoUID:        PhotoFixtures.Pointer("Photo01").PhotoUID,
		FileUID:         "ft1es39w45bnlqdw",
		FileName:        "exampleXmpFile.xmp",
		FileRoot:        RootOriginals,
		OriginalName:    "exampleXmpFile.xmp",
		FileHash:        "ocad9168fa6acc5c5c2965ddf6ec465ca42fd818",
		ModTime:         time.Date(2019, 3, 6, 2, 6, 51, 0, time.UTC).Unix(),
		FileSize:        858,
		FileType:        "xmp",
		FileMime:        "text/xmp",
		FilePrimary:     false,
		FileSidecar:     true,
		FileVideo:       false,
		FileMissing:     false,
		FilePortrait:    false,
		FileWidth:       0,
		FileHeight:      0,
		FileOrientation: 0,
		FileAspectRatio: 0,
		FileMainColor:   "",
		FileColors:      "",
		FileLuminance:   "",
		FileDiff:        0,
		FileChroma:      0,
		FileError:       "",
		Share:           []FileShare{},
		Sync:            []FileSync{},
		CreatedAt:       time.Date(2019, 3, 6, 2, 6, 51, 0, time.UTC),
		CreatedIn:       2,
		UpdatedAt:       time.Date(2020, 3, 28, 14, 6, 0, 0, time.UTC),
		UpdatedIn:       0,
		DeletedAt:       nil,
	},
	"bridge.jpg": {
		ID:              1000003,
		Photo:           PhotoFixtures.Pointer("Photo04"),
		PhotoID:         PhotoFixtures.Pointer("Photo04").ID,
		PhotoUID:        PhotoFixtures.Pointer("Photo04").PhotoUID,
		FileUID:         "ft2es39w45bnlqdw",
		FileName:        "bridge.jpg",
		FileRoot:        RootOriginals,
		OriginalName:    "bridgeOriginal.jpg",
		FileHash:        "pcad9168fa6acc5c5c2965ddf6ec465ca42fd818",
		ModTime:         time.Date(2017, 2, 6, 2, 6, 51, 0, time.UTC).Unix(),
		FileSize:        961858,
		FileType:        "jpg",
		FileMime:        "image/jpg",
		FilePrimary:     true,
		FileSidecar:     false,
		FileVideo:       false,
		FileMissing:     false,
		FilePortrait:    false,
		FileWidth:       1200,
		FileHeight:      1600,
		FileOrientation: 6,
		FileAspectRatio: 0.75,
		FileMainColor:   "magenta",
		FileColors:      "225221C1E",
		FileLuminance:   "DC42844C8",
		FileDiff:        986,
		FileChroma:      32,
		FileError:       "",
		Share:           []FileShare{},
		Sync:            []FileSync{},
		CreatedAt:       time.Date(2019, 1, 1, 2, 6, 51, 0, time.UTC),
		CreatedIn:       2,
		UpdatedAt:       time.Date(2020, 3, 28, 14, 6, 0, 0, time.UTC),
		UpdatedIn:       0,
		DeletedAt:       nil,
	},
	"reunion.jpg": {
		ID:              1000004,
		Photo:           PhotoFixtures.Pointer("Photo05"),
		PhotoID:         PhotoFixtures.Pointer("Photo05").ID,
		PhotoUID:        PhotoFixtures.Pointer("Photo05").PhotoUID,
		FileUID:         "ft3es39w45bnlqdw",
		FileName:        "reunion.jpg",
		FileRoot:        RootOriginals,
		OriginalName:    "reunionOriginal.jpg",
		FileHash:        "acad9168fa6acc5c5c2965ddf6ec465ca42fd818",
		ModTime:         time.Date(2017, 1, 6, 2, 6, 51, 0, time.UTC).Unix(),
		FileSize:        81858,
		FileType:        "jpg",
		FileMime:        "image/jpg",
		FilePrimary:     true,
		FileSidecar:     false,
		FileVideo:       false,
		FileMissing:     false,
		FilePortrait:    false,
		FileWidth:       1200,
		FileHeight:      1600,
		FileOrientation: 6,
		FileAspectRatio: 0.75,
		FileMainColor:   "blue",
		FileColors:      "266111000",
		FileLuminance:   "DC42844C8",
		FileDiff:        800,
		FileChroma:      4,
		FileError:       "Error",
		Share:           []FileShare{},
		Sync:            []FileSync{},
		CreatedAt:       time.Date(2018, 1, 1, 2, 6, 51, 0, time.UTC),
		CreatedIn:       2,
		UpdatedAt:       time.Date(2029, 3, 28, 14, 6, 0, 0, time.UTC),
		UpdatedIn:       0,
		DeletedAt:       nil,
	},
	"Quality1FavoriteTrue.jpg": {
		ID:              1000005,
		Photo:           PhotoFixtures.Pointer("Photo17"),
		PhotoID:         PhotoFixtures.Pointer("Photo17").ID,
		PhotoUID:        PhotoFixtures.Pointer("Photo17").PhotoUID,
		FileUID:         "ft4es39w45bnlqdw",
		FileName:        "Quality1FavoriteTrue.jpg",
		FileRoot:        RootOriginals,
		OriginalName:    "Quality1FavoriteTrue.jpg",
		FileHash:        "acad9168fa6acc5c5c2965ddf6ec465ca42fd819",
		ModTime:         time.Date(2017, 1, 6, 2, 6, 51, 0, time.UTC).Unix(),
		FileSize:        500,
		FileType:        "jpg",
		FileMime:        "image/jpg",
		FilePrimary:     true,
		FileSidecar:     false,
		FileVideo:       false,
		FileMissing:     false,
		FilePortrait:    false,
		FileWidth:       1200,
		FileHeight:      1600,
		FileOrientation: 6,
		FileAspectRatio: 0.75,
		FileMainColor:   "blue",
		FileColors:      "266111000",
		FileLuminance:   "DC42844C8",
		FileDiff:        800,
		FileChroma:      26,
		FileError:       "",
		Share:           []FileShare{},
		Sync:            []FileSync{},
		CreatedAt:       time.Date(2018, 1, 1, 2, 6, 51, 0, time.UTC),
		CreatedIn:       2,
		UpdatedAt:       time.Date(2029, 3, 28, 14, 6, 0, 0, time.UTC),
		UpdatedIn:       0,
		DeletedAt:       nil,
	},
	"missing.jpg": {
		ID:              1000006,
		Photo:           PhotoFixtures.Pointer("Photo15"),
		PhotoID:         PhotoFixtures.Pointer("Photo15").ID,
		PhotoUID:        PhotoFixtures.Pointer("Photo15").PhotoUID,
		FileUID:         "ft5es39w45bnlqdw",
		FileName:        "missing.jpg",
		FileRoot:        RootOriginals,
		OriginalName:    "missing.jpg",
		FileHash:        "acad9168fa6acc5c5c2965ddf6ec465ca42fd819",
		ModTime:         time.Date(2017, 1, 6, 2, 6, 51, 0, time.UTC).Unix(),
		FileSize:        500,
		FileType:        "jpg",
		FileMime:        "image/jpg",
		FilePrimary:     true,
		FileSidecar:     false,
		FileVideo:       false,
		FileMissing:     true,
		FilePortrait:    false,
		FileWidth:       1200,
		FileHeight:      1600,
		FileOrientation: 6,
		FileAspectRatio: 0.75,
		FileMainColor:   "blue",
		FileColors:      "266111000",
		FileLuminance:   "DC42844C8",
		FileDiff:        800,
		FileChroma:      26,
		FileError:       "",
		Share:           []FileShare{},
		Sync:            []FileSync{},
		CreatedAt:       time.Date(2018, 1, 1, 2, 6, 51, 0, time.UTC),
		CreatedIn:       2,
		UpdatedAt:       time.Date(2029, 3, 28, 14, 6, 0, 0, time.UTC),
		UpdatedIn:       0,
		DeletedAt:       nil,
	},
	"Photo18.jpg": {
		ID:              1000007,
		Photo:           nil, // no pointer here because related photo is deleted
		PhotoID:         1000018,
		PhotoUID:        "pt9jtdre2lvl0y25",
		FileUID:         "ft6es39w45bnlqdw",
		FileName:        "Photo18.jpg",
		FileRoot:        RootOriginals,
		OriginalName:    "Photo18.jpg",
		FileHash:        "acad9168fa6acc5c5c2965ddf6ec465ca42fd820",
		ModTime:         time.Date(2017, 1, 6, 2, 6, 51, 0, time.UTC).Unix(),
		FileSize:        500,
		FileType:        "jpg",
		FileMime:        "image/jpg",
		FilePrimary:     true,
		FileSidecar:     false,
		FileVideo:       false,
		FileMissing:     false,
		FilePortrait:    false,
		FileWidth:       1200,
		FileHeight:      1600,
		FileOrientation: 6,
		FileAspectRatio: 0.75,
		FileMainColor:   "green",
		FileColors:      "266111000",
		FileLuminance:   "DC42844C8",
		FileDiff:        800,
		FileChroma:      0,
		FileError:       "",
		Share:           []FileShare{},
		Sync:            []FileSync{},
		CreatedAt:       time.Date(2018, 1, 1, 2, 6, 51, 0, time.UTC),
		CreatedIn:       2,
		UpdatedAt:       time.Date(2029, 3, 28, 14, 6, 0, 0, time.UTC),
		UpdatedIn:       0,
		DeletedAt:       nil,
	},
	"Video.mp4": {
		ID:              1000008,
		Photo:           PhotoFixtures.Pointer("Photo10"),
		PhotoID:         PhotoFixtures.Pointer("Photo10").ID,
		PhotoUID:        PhotoFixtures.Pointer("Photo10").PhotoUID,
		FileUID:         "ft71s39w45bnlqdw",
		FileName:        "Video.mp4",
		FileRoot:        RootOriginals,
		OriginalName:    "Video.mp4",
		FileHash:        "acad9168fa6acc5c5c2965ddf6ec465ca42fd831",
		ModTime:         time.Date(2017, 1, 6, 2, 6, 51, 0, time.UTC).Unix(),
		FileSize:        500,
		FileType:        "mp4",
		FileMime:        "video/mp4",
		FilePrimary:     true,
		FileSidecar:     false,
		FileVideo:       true,
		FileMissing:     false,
		FilePortrait:    false,
		FileWidth:       1200,
		FileHeight:      1600,
		FileOrientation: 6,
		FileAspectRatio: 0.75,
		FileMainColor:   "green",
		FileColors:      "266111000",
		FileLuminance:   "DC42844C8",
		FileDiff:        800,
		FileChroma:      0,
		FileError:       "",
		Share:           []FileShare{},
		Sync:            []FileSync{},
		CreatedAt:       time.Date(2018, 1, 1, 2, 6, 51, 0, time.UTC),
		CreatedIn:       2,
		UpdatedAt:       time.Date(2029, 3, 28, 14, 6, 0, 0, time.UTC),
		UpdatedIn:       0,
		DeletedAt:       nil,
	},
	"VideoWithError.mp4": {
		ID:              1000009,
		Photo:           PhotoFixtures.Pointer("Photo10"),
		PhotoID:         PhotoFixtures.Pointer("Photo10").ID,
		PhotoUID:        PhotoFixtures.Pointer("Photo10").PhotoUID,
		FileUID:         "ft72s39w45bnlqdw",
		FileName:        "VideoError.mp4",
		FileRoot:        RootOriginals,
		OriginalName:    "VideoError.mp4",
		FileHash:        "acad9168fa6acc5c5c2965ddf6ec465ca42fd832",
		ModTime:         time.Date(2017, 1, 6, 2, 6, 51, 0, time.UTC).Unix(),
		FileSize:        500,
		FileType:        "mp4",
		FileMime:        "video/mp4",
		FilePrimary:     false,
		FileSidecar:     false,
		FileVideo:       true,
		FileMissing:     false,
		FilePortrait:    false,
		FileWidth:       1200,
		FileHeight:      1600,
		FileOrientation: 6,
		FileAspectRatio: 0.75,
		FileMainColor:   "green",
		FileColors:      "266111000",
		FileLuminance:   "DC42844C8",
		FileDiff:        800,
		FileChroma:      0,
		FileError:       "Error",
		Share:           []FileShare{},
		Sync:            []FileSync{},
		CreatedAt:       time.Date(2018, 1, 1, 2, 6, 51, 0, time.UTC),
		CreatedIn:       2,
		UpdatedAt:       time.Date(2029, 3, 28, 14, 6, 0, 0, time.UTC),
		UpdatedIn:       0,
		DeletedAt:       nil,
	},
	"bridge1.jpg": {
		ID:              1000010,
		Photo:           PhotoFixtures.Pointer("Photo02"),
		PhotoID:         PhotoFixtures.Pointer("Photo02").ID,
		PhotoUID:        PhotoFixtures.Pointer("Photo02").PhotoUID,
		FileUID:         "ft2es39q45bnlqd0",
		FileName:        "bridge1.jpg",
		FileRoot:        RootOriginals,
		OriginalName:    "bridgeOriginal1.jpg",
		FileHash:        "pcad9168fa6acc5c5c2965ddf6ec465ca42fd828",
		ModTime:         time.Date(2017, 2, 6, 2, 6, 51, 0, time.UTC).Unix(),
		FileSize:        961851,
		FileType:        "jpg",
		FileMime:        "image/jpg",
		FilePrimary:     true,
		FileSidecar:     false,
		FileVideo:       false,
		FileMissing:     false,
		FilePortrait:    false,
		FileWidth:       1200,
		FileHeight:      1600,
		FileOrientation: 6,
		FileAspectRatio: 0.75,
		FileMainColor:   "magenta",
		FileColors:      "225221C1E",
		FileLuminance:   "DC42844C8",
		FileDiff:        986,
		FileChroma:      32,
		FileError:       "",
		Share:           []FileShare{},
		Sync:            []FileSync{},
		CreatedAt:       time.Date(2019, 1, 1, 2, 6, 51, 0, time.UTC),
		CreatedIn:       2,
		UpdatedAt:       time.Date(2020, 3, 28, 14, 6, 0, 0, time.UTC),
		UpdatedIn:       0,
		DeletedAt:       nil,
	},
	"bridge2.jpg": {
		ID:              1000011,
		Photo:           PhotoFixtures.Pointer("Photo03"),
		PhotoID:         PhotoFixtures.Pointer("Photo03").ID,
		PhotoUID:        PhotoFixtures.Pointer("Photo03").PhotoUID,
		FileUID:         "ft2es49w15bnlqdw",
		FileName:        "bridge2.jpg",
		FileRoot:        RootOriginals,
		OriginalName:    "bridgeOriginal2.jpg",
		FileHash:        "pcad9168fa6acc5c5c2965adf6ec465ca42fd818",
		ModTime:         time.Date(2017, 2, 6, 2, 6, 51, 0, time.UTC).Unix(),
		FileSize:        921858,
		FileType:        "jpg",
		FileMime:        "image/jpg",
		FilePrimary:     true,
		FileSidecar:     false,
		FileVideo:       false,
		FileMissing:     false,
		FilePortrait:    false,
		FileWidth:       1200,
		FileHeight:      1600,
		FileOrientation: 6,
		FileAspectRatio: 0.75,
		FileMainColor:   "magenta",
		FileColors:      "225221C1E",
		FileLuminance:   "DC42844C8",
		FileDiff:        986,
		FileChroma:      32,
		FileError:       "",
		Share:           []FileShare{},
		Sync:            []FileSync{},
		CreatedAt:       time.Date(2019, 1, 1, 2, 6, 51, 0, time.UTC),
		CreatedIn:       2,
		UpdatedAt:       time.Date(2020, 3, 28, 14, 6, 0, 0, time.UTC),
		UpdatedIn:       0,
		DeletedAt:       nil,
	},
	"bridge3.jpg": {
		ID:              1000012,
		Photo:           PhotoFixtures.Pointer("Photo03"),
		PhotoID:         PhotoFixtures.Pointer("Photo03").ID,
		PhotoUID:        PhotoFixtures.Pointer("Photo03").PhotoUID,
		FileUID:         "ft2es49whhbnlqdn",
		FileName:        "bridge3.jpg",
		FileRoot:        RootOriginals,
		OriginalName:    "bridgeOriginal.jpg",
		FileHash:        "pcad9168fa6acc5c5ba965adf6ec465ca42fd818",
		ModTime:         time.Date(2017, 2, 6, 2, 6, 51, 0, time.UTC).Unix(),
		FileSize:        921851,
		FileType:        "jpg",
		FileMime:        "image/jpg",
		FilePrimary:     false,
		FileSidecar:     false,
		FileVideo:       false,
		FileMissing:     false,
		FilePortrait:    false,
		FileWidth:       1200,
		FileHeight:      1600,
		FileOrientation: 6,
		FileAspectRatio: 0.75,
		FileMainColor:   "magenta",
		FileColors:      "225221C1E",
		FileLuminance:   "DC42844C8",
		FileDiff:        986,
		FileChroma:      32,
		FileError:       "",
		Share:           []FileShare{},
		Sync:            []FileSync{},
		CreatedAt:       time.Date(2019, 1, 1, 2, 6, 51, 0, time.UTC),
		CreatedIn:       2,
		UpdatedAt:       time.Date(2020, 3, 28, 14, 6, 0, 0, time.UTC),
		UpdatedIn:       0,
		DeletedAt:       nil,
	},
	"bridge.mp4": {
		ID:              1000013,
		Photo:           PhotoFixtures.Pointer("Photo03"),
		PhotoID:         1000003,
		PhotoUID:        PhotoFixtures.Pointer("Photo03").PhotoUID,
		FileUID:         "ft2es49whhbnlqdy",
		FileName:        "bridge.mp4",
		FileRoot:        RootOriginals,
		OriginalName:    "bridgeOriginal.mp4",
		FileHash:        "pcad9168fa6acc5c5ba965adf6ec465ca42fd819",
		ModTime:         time.Date(2017, 2, 6, 2, 6, 51, 0, time.UTC).Unix(),
		FileSize:        921851,
		FileType:        "mp4",
		FileMime:        "image/mp4",
		FilePrimary:     false,
		FileSidecar:     false,
		FileVideo:       true,
		FileMissing:     false,
		FilePortrait:    false,
		FileWidth:       1200,
		FileHeight:      1600,
		FileOrientation: 6,
		FileAspectRatio: 0.75,
		FileMainColor:   "magenta",
		FileColors:      "225221C1E",
		FileLuminance:   "DC42844C8",
		FileDiff:        986,
		FileChroma:      32,
		FileError:       "",
		Share:           []FileShare{},
		Sync:            []FileSync{},
		CreatedAt:       time.Date(2019, 1, 1, 2, 6, 51, 0, time.UTC),
		CreatedIn:       2,
		UpdatedAt:       time.Date(2020, 3, 28, 14, 6, 0, 0, time.UTC),
		UpdatedIn:       0,
		DeletedAt:       nil,
	},
	"Photo19.jpg": {
		ID:              1000019,
		Photo:           PhotoFixtures.Pointer("Photo19"),
		PhotoID:         PhotoFixtures.Pointer("Photo19").ID,
		PhotoUID:        PhotoFixtures.Pointer("Photo19").PhotoUID,
		FileUID:         "ft2es49qhhinlqdn",
		FileName:        "Photo19.jpg",
		FileRoot:        RootOriginals,
		OriginalName:    "Photo19.jpg",
		FileHash:        "pcad9a68fa6acc5c5ba965adf6ec465ca42fd811",
		ModTime:         time.Date(2020, 2, 6, 2, 6, 51, 0, time.UTC).Unix(),
		FileSize:        921831,
		FileType:        "jpg",
		FileMime:        "image/jpg",
		FilePrimary:     true,
		FileSidecar:     false,
		FileVideo:       false,
		FileMissing:     false,
		FilePortrait:    false,
		FileWidth:       1200,
		FileHeight:      1600,
		FileOrientation: 6,
		FileAspectRatio: 0.75,
		FileMainColor:   "magenta",
		FileColors:      "225221C1E",
		FileLuminance:   "DC42844C8",
		FileDiff:        986,
		FileChroma:      32,
		FileError:       "",
		Share:           []FileShare{},
		Sync:            []FileSync{},
		CreatedAt:       time.Date(2019, 1, 1, 2, 6, 51, 0, time.UTC),
		CreatedIn:       2,
		UpdatedAt:       time.Date(2020, 3, 28, 14, 6, 0, 0, time.UTC),
		UpdatedIn:       0,
		DeletedAt:       nil,
	},
	"Photo25.jpg": {
		ID:              1000020,
		Photo:           PhotoFixtures.Pointer("Photo25"),
		PhotoID:         PhotoFixtures.Pointer("Photo25").ID,
		PhotoUID:        PhotoFixtures.Pointer("Photo25").PhotoUID,
		FileUID:         "ft2es49qhhinlplk",
		FileName:        "Photo25.jpg",
		FileRoot:        RootOriginals,
		OriginalName:    "Photo25.jpg",
		FileHash:        "pcad9a68fa6acc5c5ba965adf6ec465ca42fd887",
		ModTime:         time.Date(2008, 2, 6, 2, 6, 51, 0, time.UTC).Unix(),
		FileSize:        921831,
		FileType:        "jpg",
		FileMime:        "image/jpg",
		FilePrimary:     true,
		FileSidecar:     false,
		FileVideo:       false,
		FileMissing:     false,
		FilePortrait:    false,
		FileWidth:       1200,
		FileHeight:      1600,
		FileOrientation: 6,
		FileAspectRatio: 0.75,
		FileMainColor:   "magenta",
		FileColors:      "225221C1E",
		FileLuminance:   "DC42844C8",
		FileDiff:        986,
		FileChroma:      32,
		FileError:       "",
		Share:           []FileShare{},
		Sync:            []FileSync{},
		CreatedAt:       time.Date(2019, 1, 1, 2, 6, 51, 0, time.UTC),
		CreatedIn:       2,
		UpdatedAt:       time.Date(2020, 3, 28, 14, 6, 0, 0, time.UTC),
		UpdatedIn:       0,
		DeletedAt:       nil,
	},
}

var FileFixturesExampleJPG = FileFixtures["exampleFileName.jpg"]
var FileFixturesExampleXMP = FileFixtures["exampleXmpFile.xmp"]
var FileFixturesExampleBridge = FileFixtures["bridge.jpg"]
var FileFixturesExampleBridgeVideo = FileFixtures["bridge.mp4"]

// CreateFileFixtures inserts known entities into the database for testing.
func CreateFileFixtures() {
	for _, entity := range FileFixtures {
		Db().Create(&entity)
	}
}
