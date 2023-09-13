var board = [];
var boardSize = -1;
var isCurrentlyGuessing = false;

var MESSAGE_WELCOME = "Welcome! Please input the Point and Voltorb totals.";
var MESSAGE_LOADING = "Solving...";
var MESSAGE_ASK_SUFFIX_CLICK_ANY_OTHER_CARD = " (Or click any other Card!)"
var MESSAGE_ASK_SAFE = "What is this Card?" + MESSAGE_ASK_SUFFIX_CLICK_ANY_OTHER_CARD;
var MESSAGE_ASK_UNSAFE = "What is this Card? $% chance it's a <img src=\"/assets/images/volt.png\" />." + MESSAGE_ASK_SUFFIX_CLICK_ANY_OTHER_CARD
var MESSAGE_WIN = "Game clear! You've found all the hidden <img src=\"/assets/images/3.png\"> and <img src=\"/assets/images/2.png\"> cards.";
var MESSAGE_ERROR_IMPOSSIBLE_BOARD = "This board is not possible. Please check your input.";
var MESSAGE_ERROR_TIMEOUT = "This board is too complex for me to solve. Sorry. :(";
var MESSAGE_ERROR_UNKNOWN = "An error occurred.";

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

function createAskLink(c, enabled, f, x, y) {
  a = $(document.createElement("a"))
    .attr("href", "#")
    .addClass(c)
    .click(function(event) {
      // Call the provided function f with the captured x and y values
      f(event, x, y);
    });
  if (!enabled)
    a.append($(document.createElement("div")).addClass("buttonoff"));
  return a;
}

function guess3(event, x, y) {
  event.preventDefault();
  board[y][x] = [false, false, false, true];
  ajaxSolve();
}

function guess2(event, x, y) {
  event.preventDefault();
  board[y][x] = [false, false, true, false];
  ajaxSolve();
}

function guess1(event, x, y) {
  event.preventDefault();
  board[y][x] = [false, true, false, false];
  ajaxSolve();
}

function guessV(event, x, y) {
  event.preventDefault();
  board[y][x] = [true, false, false, false];
  updateBoardDisplay();
  isCurrentlyGuessing = false;
  showMessage("lose", "Oh no! You get 0 Coins!");
}

function ajaxError(x, t, m) {
  if (t == "timeout") {
    showMessage("error", MESSAGE_ERROR_TIMEOUT);
  } else {
    showMessage("error", MESSAGE_ERROR_UNKNOWN);
  }
}

function makeCardGuessable(guessX, guessY, safety) {
    isCurrentlyGuessing = true;

    $(("#card"+guessX)+guessY).removeClass("unknown");
    $(("#card"+guessX)+guessY).addClass("ask");
    $(("#card"+guessX)+guessY).html("");
    var row1 = $(document.createElement("tr"));
    var cell1 = $(document.createElement("td"));
    cell1.append(
        createAskLink(
            "three",
            board[guessY][guessX][3],
            guess3,
            guessX,
            guessY));
    row1.append(cell1);
    var cell2 = $(document.createElement("td"));
    cell2.append(
        createAskLink(
            "two",
            board[guessY][guessX][2],
            guess2,
            guessX,
            guessY));
    row1.append(cell2);
    var row2 = $(document.createElement("tr"));
    var cell1a = $(document.createElement("td"));
    cell1a.append(
        createAskLink(
            "one",
            board[guessY][guessX][1],
            guess1,
            guessX,
            guessY));
    row2.append(cell1a);

    var cell2a = $(document.createElement("td"));
    cell2a.append(
        createAskLink(
            "volt",
            board[guessY][guessX][0] && (safety < 1),
            guessV,
            guessX,
            guessY));
    row2.append(cell2a);
    var tbl = $(document.createElement("table"));
    tbl.append(row1);
    tbl.append(row2);
    $(("#card"+guessX)+guessY).append(tbl);
}

function ajaxSuccess(data) {
  if (!data) {
    showMessage("error", MESSAGE_ERROR_UNKNOWN);
    return;
  }

  if (!data["IsPossible"]) {
    showMessage("error", MESSAGE_ERROR_IMPOSSIBLE_BOARD);
    return;
  }

  board = data["Tiles"]
  updateBoardDisplay();

  if (data["IsWon"]) {
    showMessage("win", MESSAGE_WIN);
    return;
  }

  var guessX = parseInt(data["SafestPosition"]["X"]);
  var guessY = parseInt(data["SafestPosition"]["Y"]);
  var safety = parseFloat(data["Safety"]);

  makeCardGuessable(guessX, guessY, safety);

  if (safety >= 1)
    showMessage("win", MESSAGE_ASK_SAFE);
  else {
    safety = (1-safety)*100;
    if (safety < 1) {
      safetyString = "1";
    } else {
      safetyString = parseFloat(safety).toFixed(0);
    }
    showMessage("warn", MESSAGE_ASK_UNSAFE.replace("$", safetyString));
  }
}

function reset() {
  showMessage("win", MESSAGE_WELCOME);
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

  isCurrentlyGuessing = false;

  $("#solve").css("visibility", "visible");
}

function ajaxSolve(query) {
  isCurrentlyGuessing = false;

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

    showMessage("win", MESSAGE_LOADING);

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
    showMessage("lose", MESSAGE_ERROR_UNKNOWN);
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