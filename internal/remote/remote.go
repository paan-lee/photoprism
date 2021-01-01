/*

Package remote implements remote service sync and uploads.

See also:
  - RClone (https://rclone.org/), a popular Go tool for syncing data with remote services

Copyright (c) 2018 - 2021 Michael Mayer <hello@photoprism.org>

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU Affero General Public License as published
    by the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU Affero General Public License for more details.

    You should have received a copy of the GNU Affero General Public License
    along with this program.  If not, see <https://www.gnu.org/licenses/>.

    PhotoPrism® is a registered trademark of Michael Mayer.  You may use it as required
    to describe our software, run your own server, for educational purposes, but not for
    offering commercial goods, products, or services without prior written permission.
    In other words, please ask.

Feel free to send an e-mail to hello@photoprism.org if you have questions,
want to support our work, or just want to say hello.

Additional information can be found in our Developer Guide:
https://docs.photoprism.org/developer-guide/

*/
package remote

import (
	"net/http"
	"time"
)

var client = &http.Client{Timeout: 30 * time.Second} // TODO: Change timeout if needed

const (
	ServiceWebDAV    = "webdav"
	ServiceFacebook  = "facebook"
	ServiceTwitter   = "twitter"
	ServiceFlickr    = "flickr"
	ServiceInstagram = "instagram"
	ServiceEyeEm     = "eyeem"
	ServiceTelegram  = "telegram"
	ServiceWhatsApp  = "whatsapp"
	ServiceGPhotos   = "gphotos"
	ServiceGDrive    = "gdrive"
	ServiceOneDrive  = "onedrive"
)

func HttpOk(method, rawUrl string) bool {
	req, err := http.NewRequest(method, rawUrl, nil)

	if err != nil {
		return false
	}

	if resp, err := client.Do(req); err != nil {
		return false
	} else if resp.StatusCode < 400 {
		return true
	}

	return false
}
