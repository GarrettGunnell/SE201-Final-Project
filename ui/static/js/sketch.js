var tiles = []
var shipcoords = [5, 5]

function setup() {
  createCanvas(windowWidth, windowHeight);
  initializeGrid();
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
    text("ship", tile.x + 10, tile.y + tile.size / 2);
}

function moveShip(direction) {
    shiptile = tiles[shipcoords[0]][shipcoords[1]]
    var i = shipcoords[0]
    var j = shipcoords[1]

    if (direction == 'west') {
        if (legalMovement(i, j, direction)) {
            shipcoords = [i, j - 1]
        }
    } else if (direction == 'east') {
        if (legalMovement(i, j, direction)) {
            shipcoords = [i, j + 1]
        }
    } else if (direction == 'north') {
        if (legalMovement(i, j, direction)) {
            shipcoords = [i - 1, j]
        }
    } else if (direction == 'south') {
        if (legalMovement(i, j, direction)) {
            shipcoords = [i + 1, j]
        }
    }
}

function legalMovement(i, j, direction) {
    if (direction == 'west') {
        return (j - 1) >= 0
    } else if (direction == 'east') {
        return (j + 1) <= 10
    } else if (direction == 'north') {
        return (i - 1) >= 0
    } else if (direction == 'south') {
        return (i + 1) <= 10
    }
}

function keyTyped() {
    if (key == 'a') {
        moveShip('west')
    } else if (key == 'd') {
        moveShip('east')
    } else if (key == 'w') {
        moveShip('north')
    } else if (key == 's') {
        moveShip('south')
     }
}
