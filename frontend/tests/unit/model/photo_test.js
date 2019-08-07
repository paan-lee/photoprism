import assert from "assert";
import Photo from "model/photo";

describe("model/photo", () => {
    it("should get photo entity name",  () => {
        const values = {id: 5, PhotoTitle: "Crazy Cat"};
        const photo = new Photo(values);
        const result = photo.getEntityName();
        assert.equal(result, "Crazy Cat");
    });

    it("should get photo id",  () => {
        const values = {ID: 5, PhotoTitle: "Crazy Cat", PhotoUUID: 789};
        const photo = new Photo(values);
        const result = photo.getId();
        assert.equal(result, 5);
    });

    it("should get photo title",  () => {
        const values = {ID: 5, PhotoTitle: "Crazy Cat", PhotoUUID: 789};
        const photo = new Photo(values);
        const result = photo.getTitle();
        assert.equal(result, "Crazy Cat");
    });

    it("should get photo color brown",  () => {
        const values = {ID: 5, PhotoTitle: "Crazy Cat", PhotoColor: "brown"};
        const photo = new Photo(values);
        const result = photo.getColor();
        assert.equal(result, "grey lighten-2");
    });

    it("should get photo color grey",  () => {
        const values = {ID: 5, PhotoTitle: "Crazy Cat", PhotoColor: "grey"};
        const photo = new Photo(values);
        const result = photo.getColor();
        assert.equal(result, "grey lighten-2");
    });

    it("should get photo color pink",  () => {
        const values = {ID: 5, PhotoTitle: "Crazy Cat", PhotoColor: "pink"};
        const photo = new Photo(values);
        const result = photo.getColor();
        assert.equal(result, "pink lighten-4");
    });

    it("should get photo maps link",  () => {
        const values = {ID: 5, PhotoTitle: "Crazy Cat", PhotoLat: 36.442881666666665, PhotoLong: 28.229493333333334};
        const photo = new Photo(values);
        const result = photo.getGoogleMapsLink();
        assert.equal(result, "https://www.google.com/maps/place/36.442881666666665,28.229493333333334");
    });

    it("should get photo thumbnail url",  () => {
        const values = {ID: 5, PhotoTitle: "Crazy Cat", FileHash: 345982};
        const photo = new Photo(values);
        const result = photo.getThumbnailUrl("tile500");
        assert.equal(result, "/api/v1/thumbnails/345982/tile500");
    });

    it("should get photo download url",  () => {
        const values = {ID: 5, PhotoTitle: "Crazy Cat", FileHash: 345982};
        const photo = new Photo(values);
        const result = photo.getDownloadUrl();
        assert.equal(result, "/api/v1/download/345982");
    });

    it("should get photo thumbnail src set",  () => {
        const values = {ID: 5, PhotoTitle: "Crazy Cat", FileHash: 345982};
        const photo = new Photo(values);
        const result = photo.getThumbnailSrcset();
        console.log(result);
        console.log(result[1]);
        assert.equal(result, "/api/v1/thumbnails/345982/fit_720 720w, /api/v1/thumbnails/345982/fit_1280 1280w, /api/v1/thumbnails/345982/fit_1920 1920w, /api/v1/thumbnails/345982/fit_2560 2560w, /api/v1/thumbnails/345982/fit_3840 3840w");
    });

    it("should calculate photo size",  () => {
        const values = {ID: 5, PhotoTitle: "Crazy Cat", FileWidth: 500, FileHeight: 200};
        const photo = new Photo(values);
        const result = photo.calculateSize(500, 200);
        assert.equal(result.width, 500);
        assert.equal(result.height, 200);
    });

    it("should calculate photo size with srcAspectRatio < maxAspectRatio",  () => {
        const values = {ID: 5, PhotoTitle: "Crazy Cat", FileWidth: 500, FileHeight: 200};
        const photo = new Photo(values);
        const result = photo.calculateSize(300, 50);
        assert.equal(result.width, 125);
        assert.equal(result.height, 50);
    });

    it("should calculate photo size with srcAspectRatio > maxAspectRatio",  () => {
        const values = {ID: 5, PhotoTitle: "Crazy Cat", FileWidth: 500, FileHeight: 200};
        const photo = new Photo(values);
        const result = photo.calculateSize(400, 300);
        assert.equal(result.width, 400);
        assert.equal(result.height, 160);
    });

    it("should get thumbnail sizes",  () => {
        const values = {ID: 5, PhotoTitle: "Crazy Cat", FileWidth: 500, FileHeight: 200};
        const photo = new Photo(values);
        const result = photo.getThumbnailSizes();
        assert.equal(result, "(min-width: 2560px) 3840px, (min-width: 1920px) 2560px, (min-width: 1280px) 1920px, (min-width: 720px) 1280px, 720px");
    });

    it("should get date string",  () => {
        const t = "2009-11-17 20:34:58.651387237 +0000 UTC";
        const values = {ID: 5, PhotoTitle: "Crazy Cat", TakenAt: t};
        const photo = new Photo(values);
        const result = photo.getDateString();
        assert.equal(result, "November 17, 2009 8:34 PM");
    });

    it("should test whether photo has location",  () => {
        const values = {ID: 5, PhotoTitle: "Crazy Cat", PhotoLat: 36.442881666666665, PhotoLong: 28.229493333333334};
        const photo = new Photo(values);
        const result = photo.hasLocation();
        assert.equal(result, true);
    });

    it("should test whether photo has location",  () => {
        const values = {ID: 5, PhotoTitle: "Crazy Cat", PhotoLat: 0, PhotoLong: 0};
        const photo = new Photo(values);
        const result = photo.hasLocation();
        assert.equal(result, false);
    });

    it("should get location",  () => {
        const values = {ID: 5, PhotoTitle: "Crazy Cat", LocationID: 6, LocType: "viewpoint", LocName: "Cape Point", LocCountry: "Africa"};
        const photo = new Photo(values);
        const result = photo.getLocation();
        assert.equal(result, "Cape Point, Africa");
    });

    it("should get location",  () => {
        const values = {ID: 5, PhotoTitle: "Crazy Cat", LocationID: 6, LocType: "viewpoint", LocCountry: "Africa", LocCity: "Cape Town", LocCounty: "County", LocState: "State"};
        const photo = new Photo(values);
        const result = photo.getLocation();
        assert.equal(result, "Cape Town, State, Africa");
    });

    it("should get location",  () => {
        const values = {ID: 5, PhotoTitle: "Crazy Cat", LocType: "viewpoint", LocName: "Cape Point", LocCountry: "Africa", LocCity: "Cape Town", LocCounty: "County", LocState: "State"};
        const photo = new Photo(values);
        const result = photo.getLocation();
        assert.equal(result, "Unknown");
    });

    it("should get location",  () => {
        const values = {ID: 5, PhotoTitle: "Crazy Cat", CountryName: "Africa", LocCity: "Cape Town"};
        const photo = new Photo(values);
        const result = photo.getLocation();
        assert.equal(result, "Africa");
    });

    it("should get full location",  () => {
        const values = {ID: 5, PhotoTitle: "Crazy Cat", LocationID: 55, LocName: "Cape Point", LocCountry: "Africa", LocCity: "Cape Town", LocCounty: "County", LocState: "State", LocPostcode: 12345};
        const photo = new Photo(values);
        const result = photo.getFullLocation();
        assert.equal(result, "Cape Point, Cape Town, 12345, County, State, Africa");
    });

    it("should get full location",  () => {
        const values = {ID: 5, PhotoTitle: "Crazy Cat", CountryName: "Africa"};
        const photo = new Photo(values);
        const result = photo.getFullLocation();
        assert.equal(result, "Africa");
    });

    it("should get full location",  () => {
        const values = {ID: 5, PhotoTitle: "Crazy Cat", LocCity: "Cape Town"};
        const photo = new Photo(values);
        const result = photo.getFullLocation();
        assert.equal(result, "Unknown");
    });

    it("should get camera",  () => {
        const values = {ID: 5, PhotoTitle: "Crazy Cat", CameraModel: "EOSD10", CameraMake: "Canon"};
        const photo = new Photo(values);
        const result = photo.getCamera();
        assert.equal(result, "Canon EOSD10");
    });

    it("should get camera",  () => {
        const values = {ID: 5, PhotoTitle: "Crazy Cat"};
        const photo = new Photo(values);
        const result = photo.getCamera();
        assert.equal(result, "Unknown");
    });

    it("should get collection resource",  () => {
        const result = Photo.getCollectionResource();
        assert.equal(result, "photos");
    });

    it("should get model name",  () => {
        const result = Photo.getModelName();
        assert.equal(result, "Photo");
    });

});