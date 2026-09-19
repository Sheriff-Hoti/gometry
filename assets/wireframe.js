/* Reusable 3D wireframe renderer. Lessons supply cube parameters; this
 * draws a rotating cube on a 2D canvas with sliders for angles and focal
 * length, plus a live readout tracing one vertex from 3D to pixels.
 *
 *   Wireframe.mount(el, {
 *     size: 10, depth: 32,       // cube half-edge, center Z (camera at origin)
 *     focal: 48,                 // focal length in pixels
 *     rotX: 0.35, rotY: 0.5,     // starting angles, radians
 *     auto: false,               // play rotation on load
 *     sliders: true,             // show RotX/RotY/focal sliders
 *     w: 340, h: 260,            // canvas pixels
 *   });
 *
 * The math mirrors shapes/vec3.go and shapes/cube.go exactly (Y
 * down-screen, rotate X then Y, perspective divide, Z <= 0 rejected), so
 * anything learned here transfers to the Go code line-for-line. Pure math
 * is exposed as Wireframe.math for testing.
 * Expects ../assets/style.css (.wire ...) for styling.
 */
(function () {
  "use strict";

  var EDGES = [
    [0, 1], [1, 3], [3, 2], [2, 0],
    [4, 5], [5, 7], [7, 6], [6, 4],
    [0, 4], [1, 5], [2, 6], [3, 7]
  ];

  function rotateX(v, a) {
    var c = Math.cos(a), s = Math.sin(a);
    return { x: v.x, y: v.y * c - v.z * s, z: v.y * s + v.z * c };
  }

  function rotateY(v, a) {
    var c = Math.cos(a), s = Math.sin(a);
    return { x: v.x * c + v.z * s, y: v.y, z: -v.x * s + v.z * c };
  }

  // Returns pixel coords or null when behind the camera.
  function project(v, focal, cx, cy) {
    if (v.z <= 0) return null;
    return {
      x: Math.round(cx + v.x * focal / v.z),
      y: Math.round(cy + v.y * focal / v.z)
    };
  }

  function vertices(size, depth, rotX, rotY) {
    var vs = [];
    for (var i = 0; i < 8; i++) {
      var v = {
        x: size * ((i & 1) * 2 - 1),
        y: size * (((i >> 1) & 1) * 2 - 1),
        z: size * (((i >> 2) & 1) * 2 - 1)
      };
      v = rotateY(rotateX(v, rotX), rotY);
      vs.push({ x: v.x, y: v.y, z: v.z + depth });
    }
    return vs;
  }

  function mount(el, opts) {
    var size = opts.size, depth = opts.depth, focal = opts.focal;
    var rotX = opts.rotX, rotY = opts.rotY;
    var W = opts.w || 340, H = opts.h || 260;
    var timer = null;

    el.classList.add("wire");

    var canvas = document.createElement("canvas");
    canvas.width = W;
    canvas.height = H;
    el.appendChild(canvas);
    var ctx = canvas.getContext("2d");

    var readout = document.createElement("div");
    readout.className = "readout";
    el.appendChild(readout);

    function frame() {
      ctx.clearRect(0, 0, W, H);
      var vs = vertices(size, depth, rotX, rotY);
      var ps = vs.map(function (v) { return project(v, focal, W / 2, H / 2); });
      ctx.strokeStyle = "#0b5fff";
      ctx.lineWidth = 1.5;
      EDGES.forEach(function (e) {
        var a = ps[e[0]], b = ps[e[1]];
        if (!a || !b) return; // edge behind the camera: skipped like Lines()
        ctx.beginPath();
        ctx.moveTo(a.x, a.y);
        ctx.lineTo(b.x, b.y);
        ctx.stroke();
      });
      ctx.fillStyle = "#1a1a1a";
      ps.forEach(function (p) {
        if (!p) return;
        ctx.beginPath();
        ctx.arc(p.x, p.y, 2.5, 0, 2 * Math.PI);
        ctx.fill();
      });
      var v = vs[7], p = ps[7];
      readout.textContent = p
        ? "V7 (" + v.x.toFixed(1) + ", " + v.y.toFixed(1) + ", " + v.z.toFixed(1) +
          ") → pixel (" + p.x + ", " + p.y + ")   rotX=" + rotX.toFixed(2) +
          " rotY=" + rotY.toFixed(2) + " focal=" + focal
        : "V7 behind the camera (z ≤ 0) — edge skipped, like Project() returning false";
      if (slRotX) { slRotX.value = rotX; slRotY.value = rotY; }
    }

    var controls = document.createElement("div");
    controls.className = "controls";
    el.appendChild(controls);

    function slider(label, min, max, step, get, set) {
      var wrap = document.createElement("label");
      wrap.textContent = label + " ";
      var input = document.createElement("input");
      input.type = "range";
      input.min = min; input.max = max; input.step = step;
      input.value = get();
      input.oninput = function () { stop(); set(parseFloat(input.value)); frame(); };
      wrap.appendChild(input);
      controls.appendChild(wrap);
      return input;
    }

    var slRotX = null, slRotY = null;
    if (opts.sliders !== false) {
      slRotX = slider("rotX", 0, (2 * Math.PI).toFixed(3), 0.01,
        function () { return rotX; }, function (v) { rotX = v; });
      slRotY = slider("rotY", 0, (2 * Math.PI).toFixed(3), 0.01,
        function () { return rotY; }, function (v) { rotY = v; });
      slider("focal", 10, 120, 1,
        function () { return focal; }, function (v) { focal = v; });
    }

    var playBtn = document.createElement("button");
    playBtn.textContent = "Play";
    playBtn.onclick = function () {
      if (timer) { stop(); return; }
      playBtn.textContent = "Stop";
      timer = setInterval(function () {
        rotY = (rotY + 0.1) % (2 * Math.PI); // same rate as Spin(0.06, 0.1)'s Y
        rotX = (rotX + 0.06) % (2 * Math.PI);
        frame();
      }, 100);
    };
    controls.appendChild(playBtn);

    function stop() {
      if (timer) {
        clearInterval(timer);
        timer = null;
        playBtn.textContent = "Play";
      }
    }

    frame();
    if (opts.auto) playBtn.onclick();
  }

  window.Wireframe = {
    mount: mount,
    math: { rotateX: rotateX, rotateY: rotateY, project: project, vertices: vertices, edges: EDGES }
  };
})();
