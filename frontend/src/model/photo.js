import Abstract from 'model/abstract';

class Photo extends Abstract {
    getEntityName() {
        return this.PhotoTitle;
    }

    getId() {
        return this.ID;
    }

    getGoogleMapsLink() {
        return 'https://www.google.com/maps/place/' + this.PhotoLat + ',' + this.PhotoLong;
    }

    getThumbnailUrl(type, size) {
        return '/api/v1/thumbnails/' + type + '/' + size + '/' + this.FileHash;
    }

    static getCollectionResource() {
        return 'photos';
    }

    static getModelName() {
        return 'Photo';
    }
}

export default Photo;
