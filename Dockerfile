FROM photoprism/development:20190629

# Set up project directory
WORKDIR "/go/src/github.com/photoprism/photoprism"
COPY . .
