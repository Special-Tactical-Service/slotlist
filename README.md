# Slotlist

An interactive slot list for Arma 3, other military games or for general organization.

## Development

Install the following tools:

* npm
* Golang

1. run `npm run build` inside the `public` directory to build the web application and CSS
2. run `dev.cmd` (Windows) or `./dev.sh` inside the main directory to start the web server
3. visit the application in your browser on `localhost:8888`

To watch files for changes when editing the web application, you can run `npm run dev` and `npm run style` inside the `public` directory. This will watch for file changes and rebuild the web application as needed.

## Build and deploy

1. install Docker
2. run `docker build -t slotlist .` inside the main directory
3. tag the image: `docker tag slotlist username/slotlist:version` where `username` is your username and `version` can be set to a valid version number (`latest` for example)
4. push the Docker image to the registry: `docker push username/slotlist:version` (type in your username and password if asked for)
5. deploy the application by updating the version number in your `docker-compose.yml` and call `docker-compuse up -d` inside the configuration directory
6. visit the application in your browser on `yourdomainorip:port`

## License

MIT
