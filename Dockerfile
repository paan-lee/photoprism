FROM photoprism/development:20210205

# Set up project directory
WORKDIR "/go/src/github.com/photoprism/photoprism"
COPY . .