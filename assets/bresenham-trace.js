/* Faithful port of shapes/line.go Points(), recording each decision.
 * Shared by every lesson that steps through a line.
 *
 *   BresenhamTrace.trace(ax, ay, bx, by)
 *     -> [{x, y, note}]  plotted pixels in order, with per-step notes.
 *
 * Conventions: y=0 is the TOP row (screen coordinates, matching tcell).
 * The `note` on each step names the e2/err values and the moves taken,
 * plus the true line's y at the new pixel for comparison.
 */
(function () {
  "use strict";
  var abs = Math.abs;
  function sign(n) { return n < 0 ? -1 : n > 0 ? 1 : 0; }

  function trace(ax, ay, bx, by) {
    var dx = abs(bx - ax), dy = -abs(by - ay);
    var sx = sign(bx - ax), sy = sign(by - ay);
    var err = dx + dy, x = ax, y = ay;
    var run = (bx - ax) === 0 ? 0 : (by - ay) / (bx - ax);
    function note(nx, ny, e2, moves) {
      var ideal = (ay + run * (nx - ax)).toFixed(2);
      return "e2=" + e2 + ", err=" + err +
        " → " + (moves || "stop") + "  (true y=" + ideal + ")";
    }
    var steps = [{ x: x, y: y, note: "start, err=" + err }];
    for (;;) {
      if (x === bx && y === by) break;
      var e2 = 2 * err, moves = [];
      if (e2 >= dy) { err += dy; x += sx; moves.push("x+=" + sx); }
      if (e2 <= dx) { err += dx; y += sy; moves.push("y+=" + sy); }
      steps.push({ x: x, y: y, note: note(x, y, e2, moves.join(", ")) });
    }
    return steps;
  }

  window.BresenhamTrace = { trace: trace };
})();
