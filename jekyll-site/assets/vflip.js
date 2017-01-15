var guessX = -1;
var guessY = -1;
var board = [];
var boardSize = -1;

function isNumber(n) {
  return !isNaN(parseInt(n)) && isFinite(n);
}

function isInvalidNumber(e,min,max) {
  s = e.val();

  if (!s) {
    return "This number is required.";
  } else if (!isNumber(s)) {
    return "This number is not valid.";
  }

  var i = parseInt(s);

  if (i < min) {
    return "This number must be more than {0}.".format(min);
  } else if (i > max) {
    return "This number must be less than {0}.".format(max);
  }

  return false;
}

function showMessage(t, s) {
  $("#console").removeClass("error");
  $("#console").removeClass("win");
  $("#console").removeClass("warn");
  $("#console").removeClass("lose");
  $("#console").addClass(t);
  $("#consoletext").html(s);
}

function newBoard(n) {
  boardSize = n;
  board = [];
  for (var i = 0; i < n; i++) {
    board[i] = [];
    for (var j = 0; j < n; j++) {
      board[i][j] = [true, true, true, true];
    }
  }
}

function resetBoard() {
  for (var i = 0; i < boardSize; i++) {
    for (var j = 0; j < boardSize; j++) {
      board[i][j] = [true, true, true, true];
    }
  }
}

function setInputsReadonly(readonly) {
  for (var i = 0; i < boardSize; i++) {
    $("#c" + i + "p").attr("readonly", readonly);
    if (readonly)
      $("#c" + i + "p").addClass("readonly");
    else
      $("#c" + i + "p").removeClass("readonly");

    $("#c" + i + "v").attr("readonly", readonly);
    if (readonly)
      $("#c" + i + "v").addClass("readonly");
    else
      $("#c" + i + "v").removeClass("readonly");

    $("#r" + i + "p").attr("readonly", readonly);
    if (readonly)
      $("#r" + i + "p").addClass("readonly");
    else
      $("#r" + i + "p").removeClass("readonly");

    $("#r" + i + "v").attr("readonly", readonly);
    if (readonly)
      $("#r" + i + "v").addClass("readonly");
    else
      $("#r" + i + "v").removeClass("readonly");
  }
}

var ROW_COLORS = [ "red", "grn", "yel", "blu", "pur" ]
function createHintTd(type, i) {
  var hintTd = $(document.createElement("td"));
  hintTd.addClass(ROW_COLORS[i % ROW_COLORS.length]);

  var pInput = $(document.createElement("input"));
  pInput.attr("id", type + i + "p");
  pInput.attr("type", "text");
  pInput.attr("autocomplete", "off");
  pInput.attr("maxlength", "2");
  hintTd.append(pInput);

  hintTd.append($(document.createElement("div")).addClass("whitesep"));

  var vInput = $(document.createElement("input"));
  vInput.attr("id", type + i + "v");
  vInput.attr("type", "text");
  vInput.attr("autocomplete", "off");
  vInput.attr("maxlength", "1");
  hintTd.append(vInput);

  hintTd.append($(document.createElement("div")).addClass("voltorb").append(vInput));

  return hintTd;
}

function createBoardDisplay(n) {
  for (var i = 0; i < n; i++) {
      var row = $(document.createElement("tr"));
      for (var j = 0; j < n; j++) {
	row.append($(document.createElement("td")).attr("id", ("card" + i) + j));
      }
      row.append(createHintTd("r", i));
      $("#board").append(row);
  }

  var row = $(document.createElement("tr"));
  for (var i = 0; i < n; i++) {
    row.append(createHintTd("c", i));
  }
  row.append($(document.createElement("td")).addClass("hide"));
  $("#board").append(row);
}


function updateBoardDisplay() {
  for (var i = 0; i < boardSize; i++) {
    for (var j = 0; j < boardSize; j++) {
      updateCellDisplay(i, j);
    }
  }
}

function arraysEqual(a, b) {
  if (a === b) return true;
  if (a == null || b == null) return false;
  if (a.length != b.length) return false;

  // If you don't care about the order of the elements inside
  // the array, you should sort both arrays here.

  for (var i = 0; i < a.length; ++i) {
    if (a[i] !== b[i]) return false;
  }
  return true;
}

function updateCellDisplay(i, j) {
  e = $(("#card"+j)+i);
  e.removeClass("unknown");
  e.removeClass("known");
  e.removeClass("voltorb");
  e.removeClass("ask");
  if (arraysEqual(board[i][j], [false, false, false, true])) {
    e.html("3");
    e.addClass("known");
  } else if (arraysEqual(board[i][j], [false, false, true, false])) {
    e.html("2");
    e.addClass("known");
  } else if (arraysEqual(board[i][j], [false, true, false, false])) {
    e.html("1");
    e.addClass("known");
  } else if (arraysEqual(board[i][j], [true, false, false, false])) {
    e.html(".");
    e.addClass("voltorb");
  } else {
    e.html(".");
    e.addClass("unknown");
  }
}

function createAskLink(c, enabled, f) {
  a = $(document.createElement("a"))
    .attr("href", "#")
    .addClass(c)
    .click(f);
  if (!enabled)
    a.append($(document.createElement("div")).addClass("buttonoff"));
  return a;
}

function guess3(event) {
  event.preventDefault();
  board[guessY][guessX] = [false, false, false, true];
  ajaxSolve();
}

function guess2(event) {
  event.preventDefault();
  board[guessY][guessX] = [false, false, true, false];
  ajaxSolve();
}

function guess1(event) {
  event.preventDefault();
  board[guessY][guessX] = [false, true, false, false];
  ajaxSolve();
}

function guessV(event) {
  event.preventDefault();
  board[guessY][guessX] = [true, false, false, false];
  updateBoardDisplay();
  showMessage("lose", "Oh no! You get 0 Coins!");
}

function ajaxError(x, t, m) {
  if (t == "timeout") {
    showMessage("error", "This board is too complex for me to solve. Sorry. :(");
  } else {
    showMessage("error", "An error occurred.");
  }
}

function ajaxSuccess(data) {
  if (!data) {
    showMessage("error", "An error occurred.");
    return;
  }

  if (!data["IsPossible"]) {
    showMessage("error", "This board is not possible. Please check your input.");
    return;
  }

  board = data["Tiles"]
  updateBoardDisplay();

  if (data["IsWon"]) {
    showMessage("win", "Game clear! You've found all the hidden <img src=\"https://s3.amazonaws.com/vflip/images/3.png\"> and <img src=\"https://s3.amazonaws.com/vflip/images/2.png\"> cards.");
    return;
  }

  guessX = parseInt(data["SafestPosition"]["X"]);
  guessY = parseInt(data["SafestPosition"]["Y"]);
  $(("#card"+guessX)+guessY).removeClass("unknown");
  $(("#card"+guessX)+guessY).addClass("ask");
  $(("#card"+guessX)+guessY).html("");
  var row1 = $(document.createElement("tr"));
  var cell1 = $(document.createElement("td"));
  cell1.append(
      createAskLink("three",
	board[guessY][guessX][3],
	guess3));
  row1.append(cell1);
  var cell2 = $(document.createElement("td"));
  cell2.append(
      createAskLink("two",
        board[guessY][guessX][2],
	guess2));
  row1.append(cell2);
  var row2 = $(document.createElement("tr"));
  var cell1a = $(document.createElement("td"));
  cell1a.append(
      createAskLink("one",
        board[guessY][guessX][1],
        guess1));
  row2.append(cell1a)

  safety = parseFloat(data["Safety"]);

  var cell2a = $(document.createElement("td"));
  cell2a.append(
      createAskLink("volt",
        board[guessY][guessX][0] && (safety < 1),
        guessV));
  row2.append(cell2a);
  var tbl = $(document.createElement("table"));
  tbl.append(row1);
  tbl.append(row2);
  $(("#card"+guessX)+guessY).append(tbl);

  if (safety >= 1)
    showMessage("win", "What is this Card?");
  else {
    safety = (1-safety)*100;
    if (safety < 1) {
      safetyString = "1";
    } else {
      safetyString = parseFloat(safety).toFixed(0);
    }
    showMessage("warn", "What is this Card? There is a " + safetyString + "% chance it is a <img src=\"https://s3.amazonaws.com/vflip/images/volt.png\" />.");
  }
}

function reset() {
  showMessage("win", "Welcome! Please input the Point and Voltorb totals.");
  resetBoard();

  for (var i = 0; i < boardSize; i++) {
    //$("#c" + i + "p").val("");
    $("#c" + i + "p").removeClass("invalid");
    //$("#c" + i + "v").val("");
    $("#c" + i + "v").removeClass("invalid");
    //$("#r" + i + "p").val("");
    $("#r" + i + "p").removeClass("invalid");
    //$("#r" + i + "v").val("");
    $("#r" + i + "v").removeClass("invalid");
  }

  updateBoardDisplay();

  setInputsReadonly(false);

  $("#solve").css("visibility", "visible");
}

function ajaxSolve(query) {
  var query = {};
  var invalidInput = false;

  var rowTotals = Array(5);
  var columnTotals = Array(boardSize);
  for (var i = 0; i < boardSize; i++) {
    var es = [ $("#c" + i + "p"), $("#c" + i + "v"), $("#r" + i + "p"), $("#r" + i + "v") ];

    rowTotals[i] = {}
    columnTotals[i] = {}

    for (var j = 0; j < 4; j++) {
      var e = es[j];
      if (!invalidInput) {
        var check = isInvalidNumber(e);
        if (!check) {
          e.removeClass("invalid");

          if (j == 0) {
            columnTotals[i]["Points"] = parseInt(e.val());
          } else if (j == 1) {
            columnTotals[i]["Voltorbs"] = parseInt(e.val());
          } else if (j == 2) {
            rowTotals[i]["Points"] = parseInt(e.val());
          } else {
            rowTotals[i]["Voltorbs"] = parseInt(e.val());
          }
        } else {
          showMessage("error", check);
          e.addClass("invalid");
          invalidInput = true;
        }
      } else {
        e.removeClass("invalid");
      }
    }
  }

  if (!invalidInput) {
    setInputsReadonly(true);

    showMessage("win", "Solving...");

    var query = {
      BoardTotals: {
        RowTotals: rowTotals,
        ColumnTotals: columnTotals
      },
      Tiles: board
    };
    $.ajax({
      type: 'POST',
      url: "/api/solve",
      contentType:"application/json",
      dataType: "json",
      data: JSON.stringify(query),
      timeout: 15000,
      success: ajaxSuccess,
      error: ajaxError
    });
  } else {
    $("#solve").css("visibility", "visible");
  }
}

function isVoltorb(x) {
  return (x == 8);
}

function pointVal(x) {
  if (x == 1)
    return 3;
  else if (x == 2)
    return 2;
  else if (x == 4)
    return 1;
  else
    return 0;
}

function ajaxGetRandomBoardSuccess(data) {
  if (!data) {
    showMessage("lose", "An error occurred.");
    return;
  }
  // Display the row/column point/voltorb totals
  var cP = [];
  var cV = [];
  for (var i = 0; i < boardSize; i++) {
    cP[i] = 0;
    cV[i] = 0;
  }
  var rP = 0;
  var rV = 0;
  var val;
  for (var j = 0; j < boardSize; j++) {
    for (var i = 0; i < boardSize; i++) {
      val = parseInt(data[("card"+j)+i]);
      board[i][j] = val;
      if (isVoltorb(val)) {
	rV++;
	cV[i]++;
      } else {
	rP += pointVal(val);
	cP[i] += pointVal(val);
      }
    }
    $("#r" + j + "p").val(rP);
    $("#r" + j + "v").val(rV);

    rP = 0;
    rV = 0;
  }
  for (var i = 0; i < boardSize; i++) {
    $("#c" + i + "p").val(cP[i]);
    $("#c" + i + "v").val(cV[i]);
  }

  updateValuedCardsCount();

  playEnabled = true;
}

function ajaxGetRandomBoard(level) {
  query = {
    "command": "CREATE",
    "size": boardSize,
    "level": level
  };
  $.ajax({
    type: 'POST',
    url: "cgi-bin/vflip-ajax.cgi",
    data: query,
    success: ajaxGetRandomBoardSuccess,
    dataType: "json"
  });
}

// Play functions
var playEnabled = true; // Set when the player has not lost yet
var numValuedCards = 0;

function updateValuedCardsCount() {
  numValuedCards = 0;
  for (var i = 0; i < boardSize; i++) {
    for (var j = 0; j < boardSize; j++) {
      if (pointVal(board[i][j]) > 1)
	numValuedCards++;
    }
  }
}

function playNewGame() {
  playEnabled = false;
  resetBoard();
  updateBoardDisplay();
  ajaxGetRandomBoard(1);
  updateValuedCardsCount();
  showMessage("win", "Flip the cards and collect coins!");
}

function playWin() {
  if (playEnabled) {
    showMessage("win", "You win!");
    updateBoardDisplay();
    playEnabled = false;
  }
}

function playGuess(x, y) {
  if (playEnabled) {
    updateCellDisplay(y, x);
    if (isVoltorb(board[y][x])) {
      playReveal(boardSize);
      showMessage("lose", "Oh no! You get 0 Coins!");
    } else if (pointVal(board[y][x]) > 1) {
      numValuedCards--;
      if (numValuedCards <= 0) {
	playWin();
      }
    }
  }
}

function playReveal() {
  if (playEnabled) {
    showMessage("lose", "Game Over!");
    updateBoardDisplay();
    playEnabled = false;
  }
}
