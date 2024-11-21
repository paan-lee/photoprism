## Install Tools/Packages ##
1. Git - `https://www.linuxtechi.com/how-to-install-git-on-ubuntu`
2. Make - `https://www.geeksforgeeks.org/how-to-install-make-on-ubuntu/`
3. Docker - `https://docs.docker.com/engine/install/ubuntu/`
4. Docker Compose - `https://docs.photoprism.app/getting-started/troubleshooting/docker/` or `https://linuxgenie.net/install-docker-compose-ubuntu-24-04/`
5. Go Lang - `https://linux.how2shout.com/your-first-steps-with-golang-installing-go-on-ubuntu-linux/`
6. Tensor Flow - `https://linux.how2shout.com/2-ways-for-installing-tensorflow-on-ubuntu-24-04-lts-linux/`
7. Nodejs & Npm - `https://linux.how2shout.com/installing-node-js-and-npm-on-ubuntu-24-04-lts-linux/`

## Setup and Clone Photoprism ##
1. Open terminal
2. `git config --global core.autocrlf false`
3. `cd Projects`
4. `git clone -b release --single-branch git@github.com:paan-lee/photoprism.git`

## Step to Launch Photoprism ##
1. Open terminal
2. Go to photoprism directory, `cd photoprism`
3. `sudo make docker-build`
4. `sudo docker compose up`
5. Open new terminal
6. `sudo make terminal`
7. `sudo make dep`
8. `sudo make build-js`
9. `git config --global --add safe.directory /go/src/github.com/photoprism/photoprism`
9. `sudo make build-go`
10. `./photoprism start`
11. Launch browser, the url `http://localhost:2342/` or `https://app.localssl.dev/`
12. User: `admin`, Pass: `photoprism`