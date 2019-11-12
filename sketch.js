var tiles = []


function setup() {
  createCanvas(windowWidth, windowHeight);
  initializeGrid();
}

function draw() {
  background(255)
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
      tiles[i][j] = initializeSquare(x, y, squareSize)
      x += squareSize
    }
    y += squareSize;
  }
}

function initializeSquare(x, y, size) {
  return {
    x: x,
    y: y,
    size: size,

    draw: function() {
      square(this.x, this.y, this.size)
    }
  }
}

function drawGrid() {
  for (var i = 0; i < 11; ++i) {
    for (var j = 0; j < 11; ++j) {
      tiles[i][j].draw();
    }
  }
}
