---
layout: default
---

<h1>{{site.meta.heading}}</h1>

<div class="text console">
  <div id="console" class="spacing win">
    <div id="consoletext" class="inner">
      Welcome! Please input the Point and Voltorb totals.
    </div>
  </div>
</div>

<div class="board-wrapper">
<table class="legacy-table" id="board">
  <tr>
    <td id="card00"/><td id="card10"/><td id="card20"/><td id="card30"/><td id="card40"/>
    <td class="red">
      {% include input_board_points.html id="r0p" %}
      <div class="whitesep"/>
      <div class="voltorb">
        {% include input_board_voltorbs.html id="r0v" %}
      </div>
    </td>
  </tr>

  <tr>
    <td id="card01"/><td id="card11"/><td id="card21"/><td id="card31"/><td id="card41"/>
    <td class="grn">
      {% include input_board_points.html id="r1p" %}
      <div class="whitesep"/>
      <div class="voltorb">
        {% include input_board_voltorbs.html id="r1v" %}
      </div>
    </td>
  </tr>

  <tr>
    <td id="card02"/><td id="card12"/><td id="card22"/><td id="card32"/><td id="card42"/>
    <td class="yel">
      {% include input_board_points.html id="r2p" %}
      <div class="whitesep"/>
      <div class="voltorb">
        {% include input_board_voltorbs.html id="r2v" %}
      </div>
    </td>
  </tr>

  <tr>
    <td id="card03"/><td id="card13"/><td id="card23"/><td id="card33"/><td id="card43"/>
    <td class="blu">
      {% include input_board_points.html id="r3p" %}
      <div class="whitesep"/>
      <div class="voltorb">
        {% include input_board_voltorbs.html id="r3v" %}
      </div>
    </td>
  </tr>

  <tr>
    <td id="card04"/><td id="card14"/><td id="card24"/><td id="card34"/><td id="card44"/>
    <td class="pur">
      {% include input_board_points.html id="r4p" %}
      <div class="whitesep"/>
      <div class="voltorb">
        {% include input_board_voltorbs.html id="r4v" %}
      </div>
    </td>
  </tr>

  <tr>
    <td class="red">{% include input_board_points.html id="c0p" %}<div class="whitesep" /><div class="voltorb">{% include input_board_voltorbs.html id="c0v" %}</div></td>
    <td class="grn">{% include input_board_points.html id="c1p" %}<div class="whitesep" /><div class="voltorb">{% include input_board_voltorbs.html id="c1v" %}</div></td>
    <td class="yel">{% include input_board_points.html id="c2p" %}<div class="whitesep" /><div class="voltorb">{% include input_board_voltorbs.html id="c2v" %}</div></td>
    <td class="blu">{% include input_board_points.html id="c3p" %}<div class="whitesep" /><div class="voltorb">{% include input_board_voltorbs.html id="c3v" %}</div></td>
    <td class="pur">{% include input_board_points.html id="c4p" %}<div class="whitesep" /><div class="voltorb">{% include input_board_voltorbs.html id="c4v" %}</div></td>
    <td class="hide"/>
  </tr>
</table>

<div class="text-center">
<a href="#" id="solve" class="blue">Solve</a>
<a href="#" id="reset" class="blue">Reset</a>
</div>
</div>

<div class="text">
  <div class="spacing info">
    <div class="inner">
      <p>Voltorb Flip is a minigame in Pokemon Heart Gold and Soul Silver. The above helper can serve as a guide to help you "cheat" in this game.</p>
      <p>Start by simply filling the numbers at the end of each of the columns and rows, then clicking the blue "Solve" button.</p>
      <p>Please note that Voltorb Flip is ultimately game of chance, and therefore this solver cannot be perfect. However, it does provide you with the best possible path to the solution.</p>
    </div>
  </div>
</div>

<div class="mb-3 text-center">
    {% include social_media_button.html style="btn-dark" offsite=true url=site.social_media_links.github fa="fa-brands fa-github" %}
    {% include social_media_button.html style="btn-dark" offsite=true url=site.social_media_links.twitter fa="fa-brands fa-x-twitter" %}
</div>

<script type="text/javascript" src="static/vflip.js"></script>

<script type="text/javascript">
  $(document).ready(function() {
    newBoard(5);
    reset();

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

    for (var y = 0; y < 5; y++) {
        for (var x = 0; x < 5; x++) {

            (function(x, y) {
                $(("#card"+x)+y).click(function(event) {
            
                    if (canClickOnUnknownTileToGuess && $(this).hasClass('unknown')) {
                        event.preventDefault();
                        makeCardGuessable(x, y, 0.5);
                        return;
                    }
            
                });
            })(x, y);

        }
    }
  });
</script>
