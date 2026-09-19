# Mission extended to 3D wireframe rendering

After the Bresenham arc landed in a spinning 3D cube, the user asked to be
taught how the 2D-to-3D projection and spinning work, and confirmed extending
the mission over a one-off lesson. This matters because the 3D lessons can now
assume the Bresenham arc as prior knowledge (integer pixels, DrawHalf) and run
as a second arc: projection first, rotation plus the full pipeline second.

Evidence: direct answer to the mission-update question (2026-09-19).
Implications: MISSION.md now carries 3D success criteria; future sessions may
add a third arc (camera position/target, filled faces) the same way.
