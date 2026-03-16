# Traefik acme converter acme.json to certificates
This tool converts traefik acme.json to certificates. <br>
It is useful for converting existing traefik configurations to use Let's Encrypt certificates.

## Installation
1. Clone the repository: `git clone https://github.com/andrey4d/traefik-acme.git`
2. Navigate to the project directory: `cd traefik-acme`
3. Install dependencies using `go mod download`:
4. Build the binary using `go build -o build/traefik-acme .`
5. Run the tool with `./traefik-acme list` for show domains and certificates.
6. Run the tool with `./traefik-acme get` to convert acme.json to certificates.


## Docker 
1. Clone the repository: `git clone https://github.com/andrey4d/traefik-acme.git`
2. Navigate to the project directory: `cd traefik-acme`
3. Build the docker image: `docker build -t traefik-acme:v0.0.1 .`
4. Run the container: `docker run -it --rm -v ./acme:/acme -v ./certs:/certs traefik-acme:v0.0.1`


