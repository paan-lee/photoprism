import {Selector, t} from 'testcafe';

export default class Page {
    constructor() {
        this.view = Selector('div.p-view-select', {timeout: 15000});
        this.camera = Selector('div.p-camera-select', {timeout: 15000});
        this.countries = Selector('div.p-countries-select', {timeout: 15000});
        this.time = Selector('div.p-time-select', {timeout: 15000});
        this.search1 = Selector('div.p-search-field input', {timeout: 15000});
    }

    async setFilter(filter, option) {
        let filterSelector = "";

        switch (filter) {
            case 'view':
                filterSelector = 'div.p-view-select';
                break;
            case 'camera':
                filterSelector = 'div.p-camera-select';
                break;
            case 'time':
                filterSelector = 'div.p-time-select';
                break;
            case 'countries':
                filterSelector = 'div.p-countries-select';
                break;
            default:
                throw "unknown filter";
        }

        await t
            .click(filterSelector, {timeout: 15000});

        if (option) {
            await t
                .click(Selector('div.menuable__content__active div.v-select-list a').withText(option), {timeout: 15000})
        } else {
            await t
                .click(Selector('div.menuable__content__active div.v-select-list a').nth(1), {timeout: 15000})
        }
    }

    async search(term) {
        await t
            .typeText(this.search1, term)
            .pressKey('enter')
    }

    async openNav() {
        if (await Selector('button.p-navigation-show').visible) {
            await t.click(Selector('button.p-navigation-show'));
        } else if (await Selector('div.p-navigation-expand').exists) {
            await t.click(Selector('div.p-navigation-expand i'));
        }
    }

    async selectPhotoFromUID(uid) {
        await t
        .hover(Selector('div').withAttribute('data-uid', uid))
        .click(Selector('.t-select.t-off'));
    }

    async unselectPhotoFromUID(uid) {
        await t
            .hover(Selector('div').withAttribute('data-uid', uid))
            .click(Selector('.t-select.t-on'));
    }

    async selectNthPhoto(nPhoto) {
        await t
        .hover(Selector('div[class="v-image__image v-image__image--cover"]', {timeout:4000}).nth(nPhoto))
        .click(Selector('.t-select.t-off'));
    }

    async unselectPhoto(nPhoto) {
        await t
            .hover(Selector('div[class="v-image__image v-image__image--cover"]', {timeout:4000}).nth(nPhoto))
            .click(Selector('.t-select.t-on'));
    }

    async likePhoto(uid) {
        await t
            .click(Selector('.t-like.t-off').withAttribute('data-uid', uid));
    }

    async dislikePhoto(uid) {
        await t
            .click(Selector('.t-like.t-on').withAttribute('data-uid', uid));
    }

    async archiveSelectedPhotos() {
        await t
            .click(Selector('button.p-photo-clipboard-menu'))
            .click(Selector('button.p-photo-clipboard-archive'))
            .click(Selector('button.p-photo-dialog-confirm'));
    }

    async restoreSelectedPhotos() {
        await t
            .click(Selector('button.p-photo-clipboard-menu'))
            .click(Selector('button.p-photo-clipboard-restore'));
    }

    async editSelectedPhotos() {
        if (await Selector('button.p-photo-clipboard-edit').exists) {
            await t.click(Selector('button.p-photo-clipboard-edit'));
        } else if (await Selector('button.p-photo-clipboard-menu').exists) {
            await t
                .click(Selector('button.p-photo-clipboard-menu'))
                .click(Selector('button.p-photo-clipboard-edit'));
        }
    }

    async login(password) {
        await t
            .typeText(Selector('input[type="password"]'), password)
            .pressKey('enter');
    }

    async logout() {
        await t
            .click(Selector('div.p-navigation-logout'));
    }
}
