# go-live-shapes
Just playing with golang and websockets. The shapes are draggable and the state of these shapes are shared via websockets so all users see these changes. A demo can be viewed here https://go-live-shapes.herokuapp.com

## Start program
### Go build
In the `src` folder run `go build .` and execute `./go-live-shapes` <br/>
You can specify a different port by running `./go-live-shapes -port=PORT_NUMBER` <br/>

### Docker
Build the image `docker build -t go-live-shapes .` <br/>
Start the container `docker run --name liveshapes -p 8080:8080 -e PORT=8080 -d go-live-shapes`
