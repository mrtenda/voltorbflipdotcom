---
layout: default
---

<div class="text console">
  <div id="console" class="spacing win">
    <div id="consoletext" class="inner">
      Welcome! Please input the Point and Voltorb totals.
    </div>
  </div>
</div>

<table id="board">
  <tr>
    <td id="card00"/><td id="card10"/><td id="card20"/><td id="card30"/><td id="card40"/>
    <td class="red">
      <input type="text" autocomplete="off" id="r0p" maxlength="2" />
      <div class="whitesep"/>
      <div class="voltorb">
        <input type="text" autocomplete="off" id="r0v" maxlength="1"/>
      </div>
    </td>
  </tr>

  <tr>
    <td id="card01"/><td id="card11"/><td id="card21"/><td id="card31"/><td id="card41"/>
    <td class="grn">
      <input type="text" autocomplete="off" id="r1p" maxlength="2" />
      <div class="whitesep"/>
      <div class="voltorb">
        <input type="text" autocomplete="off" id="r1v" maxlength="1"/>
      </div>
    </td>
  </tr>

  <tr>
    <td id="card02"/><td id="card12"/><td id="card22"/><td id="card32"/><td id="card42"/>
    <td class="yel">
      <input type="text" autocomplete="off" id="r2p" maxlength="2" />
      <div class="whitesep"/>
      <div class="voltorb">
        <input type="text" autocomplete="off" id="r2v" maxlength="1"/>
      </div>
    </td>
  </tr>

  <tr>
    <td id="card03"/><td id="card13"/><td id="card23"/><td id="card33"/><td id="card43"/>
    <td class="blu">
      <input type="text" autocomplete="off" id="r3p" maxlength="2" />
      <div class="whitesep"/>
      <div class="voltorb">
        <input type="text" autocomplete="off" id="r3v" maxlength="1"/>
      </div>
    </td>
  </tr>

  <tr>
    <td id="card04"/><td id="card14"/><td id="card24"/><td id="card34"/><td id="card44"/>
    <td class="pur">
      <input type="text" autocomplete="off" id="r4p" maxlength="2" />
      <div class="whitesep"/>
      <div class="voltorb">
        <input type="text" autocomplete="off" id="r4v" maxlength="1"/>
      </div>
    </td>
  </tr>



  <tr>
    <td class="red"><input type="text" autocomplete="off" id="c0p" maxlength="2"/><div class="whitesep" /><div class="voltorb"><input type="text" autocomplete="off" id="c0v" maxlength="1"/></div></td>
    <td class="grn"><input type="text" autocomplete="off" id="c1p" maxlength="2"/><div class="whitesep" /><div class="voltorb"><input type="text" autocomplete="off" id="c1v" maxlength="1"/></div></td>
    <td class="yel"><input type="text" autocomplete="off" id="c2p" maxlength="2"/><div class="whitesep" /><div class="voltorb"><input type="text" autocomplete="off" id="c2v" maxlength="1"/></div></td>
    <td class="blu"><input type="text" autocomplete="off" id="c3p" maxlength="2"/><div class="whitesep" /><div class="voltorb"><input type="text" autocomplete="off" id="c3v" maxlength="1"/></div></td>
    <td class="pur"><input type="text" autocomplete="off" id="c4p" maxlength="2"/><div class="whitesep" /><div class="voltorb"><input type="text" autocomplete="off" id="c4v" maxlength="1"/></div></td>
    <td class="hide"/>
  </tr>
</table>

<a href="#" id="solve" class="blue">Solve</a>
<a href="#" id="reset" class="blue right">Reset</a>

<div class="text">
  <div class="spacing info">
    <div class="inner">
      <p>Voltorb Flip is a minigame in Pokemon Heart Gold and Soul Silver. The above helper can serve as a guide to help you "cheat" in this game.</p>
      <p>Start by simply filling the numbers at the end of each of the columns and rows, then clicking the blue "Solve" button.</p>
      <p>Please note that Voltorb Flip is ultimately game of chance, and therefore this solver cannot be perfect. However, it does provide you with the best possible path to the solution.</p>
    </div>
  </div>
</div>

<script type="text/javascript" src="assets/vflip.js"></script>

<script type="text/javascript">
  $(document).ready(function() {
    newBoard(5);

    $("#solve").click(function(event) {
      event.preventDefault();
      resetBoard();

      $("#solve").css("visibility", "hidden");
      ajaxSolve();
    });

    $("#reset").click(function(event) {
      event.preventDefault();
      reset();
    });

    reset();
  });
</script>
