var tiles = []
var shipcoords = [5, 5]
var callsign = "null"

function setup() {
  createCanvas(windowWidth, windowHeight);
  initializeGrid();
  //console.log(document.cookie)
  parseCookies();
}

function draw() {
  background(230)
  //line(windowWidth / 2, 0, windowWidth / 2, windowHeight)
  fill(0)
  textSize(25)
  text("Trade Wars (U w U)", windowWidth / 20, windowHeight / 15);
  text("Status: ", windowWidth / 2, windowHeight / 9)
  text("Credits: 6", 3 * (windowWidth / 4), windowHeight / 4);
  text("Cargo:", 3 * (windowWidth / 4), windowHeight / 2);
  fill(255)
  drawGrid();
}

function initializeGrid() {
  var y = windowHeight / 8;
  squareSize = windowHeight / 15;

  for (var i = 0; i < 11; ++i) {
    tiles[i] = []
    var x = (windowHeight / 15) * 10 - squareSize / 2;
    for (var j = 0; j < 11; ++j) {
        tiles[i][j] = initializeTile(x, y, squareSize)
        x += squareSize
    }
    y += squareSize;
  }
}

function initializeTile(x, y, size) {
  return {
    x: x,
    y: y,
    size: size,

    draw: function() {
      fill(255)
      square(this.x, this.y, this.size)
    }
  }
}

function drawGrid() {
  for (var i = 0; i < 11; ++i) {
    for (var j = 0; j < 11; ++j) {
        tiles[i][j].draw()
        if (shipIsHere(i, j)) {
            drawShip(tiles[i][j])
        }
    }
  }
}

function shipIsHere(x, y) {
    shipX = shipcoords[0]
    shipY = shipcoords[1]
    return x == shipX && y == shipY

}

function drawShip(tile) {
    fill(0)
    textSize(tile.size / 3)
    text(callsign, tile.x + 10, tile.y + tile.size / 2);
}

function parseCookies() {
    cookie = document.cookie
    cookies = []
    cookiename = ''
    cookievalue = 0;
    cookiestring = ''
    for (var i = 0; i < cookie.length; ++i) {
        if (cookie[i] == ";") {
            cookies.push(cookiestring)
            cookiestring = ''
            continue
        }
        cookiestring += cookie[i]
    }
    cookies.push(cookiestring);

    for (var i = 0; i < cookies.length; ++i) {
        name = cookies[i].split("=")[0].trim()
        value = cookies[i].split("=")[1].trim()

        switch(name) {
            case "XCoordinate":
                shipcoords[1] = parseInt(value)
                break
            case "YCoordinate":
                shipcoords[0] = parseInt(value)
                break
            case "Callsign":
                callsign = value
                break
        }
    }
}
