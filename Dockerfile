FROM photoprism/development:20210117

# Set up project directory
WORKDIR "/go/src/github.com/photoprism/photoprism"
COPY . .