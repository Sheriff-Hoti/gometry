/* Reusable pixel-grid stepper. Lessons supply a trace; this renders it.
 *
 *   GridSim.mount(el, {
 *     cols: 6, rows: 3,          // grid size in cells
 *     steps: [{x, y, note}],     // plotted pixels in order, with per-step note
 *     ideal: [{x, y}],           // faint background cells (e.g. the true line)
 *     xlabel: x => x, ylabel: y => y,
 *   });
 *
 * Conventions: y=0 is the TOP row (screen coordinates, matching tcell).
 * Expects ../assets/style.css (.sim .cell.on/.head/.ideal) for styling.
 */
(function () {
  "use strict";

  function cellKey(x, y) {
    return x + "," + y;
  }

  function mount(el, opts) {
    var cols = opts.cols;
    var rows = opts.rows;
    var steps = opts.steps;
    var ideal = {};
    (opts.ideal || []).forEach(function (p) {
      ideal[cellKey(p.x, p.y)] = true;
    });
    var xlabel = opts.xlabel || function (x) { return x; };
    var ylabel = opts.ylabel || function (y) { return y; };

    var idx = -1; // nothing plotted yet
    var timer = null;

    el.classList.add("sim");

    var grid = document.createElement("div");
    grid.className = "grid";
    grid.style.gridTemplateColumns = "repeat(" + cols + ", var(--cell))";
    el.appendChild(grid);

    var cells = {};
    for (var y = 0; y < rows; y++) {
      for (var x = 0; x < cols; x++) {
        var c = document.createElement("div");
        c.className = "cell" + (ideal[cellKey(x, y)] ? " ideal" : "");
        c.textContent = xlabel(x) + "," + ylabel(y);
        c.dataset.pos = cellKey(x, y);
        grid.appendChild(c);
        cells[cellKey(x, y)] = c;
      }
    }

    var controls = document.createElement("div");
    controls.className = "controls";
    el.appendChild(controls);

    var meta = document.createElement("div");
    meta.className = "meta";
    el.appendChild(meta);

    function render() {
      Object.keys(cells).forEach(function (k) {
        cells[k].classList.remove("on", "head");
      });
      for (var i = 0; i <= idx && i < steps.length; i++) {
        var s = steps[i];
        var node = cells[cellKey(s.x, s.y)];
        if (node) node.classList.add(i === idx ? "head" : "on");
      }
      if (idx < 0) {
        meta.textContent = "Press Next — but first, guess which pixel lights up.";
      } else {
        var cur = steps[Math.min(idx, steps.length - 1)];
        meta.textContent =
          "step " + (idx + 1) + "/" + steps.length +
          "  pixel (" + cur.x + "," + cur.y + ")  " + (cur.note || "");
      }
      prevBtn.disabled = idx < 0;
      nextBtn.disabled = idx >= steps.length - 1;
    }

    function stop() {
      if (timer) {
        clearInterval(timer);
        timer = null;
        playBtn.textContent = "Play";
      }
    }

    var startBtn = document.createElement("button");
    startBtn.textContent = "|◀";
    startBtn.onclick = function () { stop(); idx = -1; render(); };

    var prevBtn = document.createElement("button");
    prevBtn.textContent = "◀ Prev";
    prevBtn.onclick = function () { stop(); if (idx >= 0) idx--; render(); };

    var nextBtn = document.createElement("button");
    nextBtn.textContent = "Next ▶";
    nextBtn.onclick = function () {
      stop();
      if (idx < steps.length - 1) idx++;
      render();
    };

    var playBtn = document.createElement("button");
    playBtn.textContent = "Play";
    playBtn.onclick = function () {
      if (timer) { stop(); return; }
      if (idx >= steps.length - 1) idx = -1;
      playBtn.textContent = "Stop";
      timer = setInterval(function () {
        if (idx >= steps.length - 1) { stop(); return; }
        idx++;
        render();
      }, 600);
    };

    controls.appendChild(startBtn);
    controls.appendChild(prevBtn);
    controls.appendChild(nextBtn);
    controls.appendChild(playBtn);

    render();
  }

  window.GridSim = { mount: mount };
})();
