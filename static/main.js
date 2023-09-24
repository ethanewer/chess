const path = window.location.pathname;
const id = path.split("/")[2];

const socket = new WebSocket(`ws://${window.location.host}/ws/${id}/`);

socket.onmessage = (event) => {
  let msg = JSON.parse(event.data);
  let id = msg.id;
  let content = msg.content
  let whitesTurn = msg.whitesTurn
  let checkmate = msg.checkmate
  console.log(id, content, whitesTurn, checkmate);
  print(content)
};

socket.onclose = function (e) {
  console.error('Game socket closed unexpectedly');
};

const imageUrls = {
  whitePawn: "/static/images/White Pawn.png",
  blackPawn: "/static/images/Black Pawn.png",
  whiteRook: "/static/images/White Rook.png",
  blackRook: "/static/images/Black Rook.png",
  whiteKnight: "/static/images/White Knight.png",
  blackKnight: "/static/images/Black Knight.png",
  whiteBishop: "/static/images/White Bishop.png",
  blackBishop: "/static/images/Black Bishop.png",
  whiteQueen: "/static/images/White Queen.png",
  blackQueen: "/static/images/Black Queen.png",
  whiteKing: "/static/images/White King.png",
  blackKing: "/static/images/Black King.png"
};

function print(boardString) {
  let boardArr = [];
  for (let i = 0; i < 8; i++) {
    let row = []
    for (let j = 0; j < 8; j++) {
      row.push(" ");
    }
    boardArr.push(row);
  }
  for (let k = 0; k < 64; k++) {
    boardArr[Math.floor(k / 8)][k % 8] = boardString[k];
  }
  const table = document.getElementById('chessboard');
  while (table.rows.length > 0) {
    table.deleteRow(0);
  }
  let clickCount = 0;

  for (let i = 0; i < 8; i++) {
    const row = document.createElement('tr');
    for (let j = 0; j < 8; j++) {
      const cell = document.createElement('td');
      if ((i + j) % 2 === 0) {
        cell.classList.add('white');
      } else {
        cell.classList.add('black');
      }
      if (i === 7) {
        let div = document.createElement("div");
        div.textContent = String.fromCharCode(j + 65);
        div.classList.add('letters');
        cell.appendChild(div);
      }
      if (j === 0) {
        let div = document.createElement("div");
        div.textContent = i;
        div.classList.add("numbers");
        cell.appendChild(div);
      }
      if (boardArr[i][j] != " ") {
        const img = document.createElement('img');
        img.classList.add('piece');
        if (boardArr[i][j] === "p") {
          img.src = imageUrls.whitePawn;
        } else if (boardArr[i][j] === "P") {
          img.src = imageUrls.blackPawn;
        } else if (boardArr[i][j] === "r") {
          img.src = imageUrls.whiteRook;
        } else if (boardArr[i][j] === "R") {
          img.src = imageUrls.blackRook;
        } else if (boardArr[i][j] === "n") {
          img.src = imageUrls.whiteKnight;
        } else if (boardArr[i][j] === "N") {
          img.src = imageUrls.blackKnight;
        } else if (boardArr[i][j] === "b") {
          img.src = imageUrls.whiteBishop;
        } else if (boardArr[i][j] === "B") {
          img.src = imageUrls.blackBishop;
        } else if (boardArr[i][j] === "q") {
          img.src = imageUrls.whiteQueen;
        } else if (boardArr[i][j] === "Q") {
          img.src = imageUrls.blackQueen;
        } else if (boardArr[i][j] === "k") {
          img.src = imageUrls.whiteKing;
        } else if (boardArr[i][j] === "K") {
          img.src = imageUrls.blackKing;
        }
        cell.appendChild(img);
      }

      var pieceI;
      var pieceJ;
      cell.setAttribute("id", i + " " + j);
      cell.addEventListener("click", e => {
        if (clickCount === 0) {
          pieceI = i;
          pieceJ = j;
          clickCount++;
          cell.classList.add("highlight");
        }
        else if (clickCount === 1) {
          let x = e.clientX;
          let y = e.clientY;

          for (let l = 0; l < 8; l++) {
            for (let m = 0; m < 8; m++) {
              let c = document.getElementById(l + " " + m);
              if (c) {
                let cBound = c.getBoundingClientRect();

                if (x >= cBound.left && x <= cBound.right && y >= cBound.top && y <= cBound.bottom) {
                  socket.send("" + pieceI + pieceJ + l + m);
                  document.getElementById(pieceI + " " + pieceJ).classList.remove("highlight");
                  break;
                }
              }

            }
          }
          clickCount = 0;
        }
      });
      row.appendChild(cell);
    }
    table.appendChild(row);
  }

  let whiteScore = 0, blackScore = 0
  for (let i = 0; i < boardString.length; i++) {
    if (boardString.charAt(i) === ' ') continue;
    if (boardString.charAt(i).toLowerCase() === boardString.charAt(i)) {
      switch (boardString.charAt(i)) {
        case 'p':
          whiteScore += 1;
          break;
        case 'n':
        case 'b':
          whiteScore += 3;
          break;
        case 'r':
          whiteScore += 5;
          break;
        case 'q':
          whiteScore += 9;
          break;
      }
    }
    else {
      switch (boardString.charAt(i)) {
        case 'P':
          blackScore += 1;
          break;
        case 'N':
        case 'B':
          blackScore += 3;
          break;
        case 'R':
          blackScore += 5;
          break;
        case 'Q':
          blackScore += 9;
          break;
      }
    }
  }

  let whiteScoreDif, blackScoreDif
  if (whiteScore > blackScore) {
    whiteScoreDif = "+" + String(whiteScore - blackScore);
    blackScoreDif = " ";
  } else if (blackScore > whiteScore) {
    blackScoreDif = "+" + String(blackScore - whiteScore);
    whiteScoreDif = " ";
  } else {
    whiteScoreDif = " ";
    blackScoreDif = " ";
  }

  let bScoreDiv = document.getElementById('blackScore');
  bScoreDiv.textContent = blackScoreDif;
  let wScoreDiv = document.getElementById('whiteScore');
  wScoreDiv.textContent = whiteScoreDif;
}

if (checkmate) {
  if (whites_turn) {
    alert("White is in checkmate,\nBlack wins!");
  } else {
    alert("Black is in checkmate,\nWhite wins!");
  }
}
