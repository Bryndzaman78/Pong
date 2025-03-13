package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	Wheight = 600
	Wwidth  = 1000
)

var (
	speed float32 = 10

	rec_left = rl.NewRectangle(5, 5, 40, 100)

	rec_right = rl.NewRectangle(Wwidth-45, 5, 40, 100)
)

func init() {
	rl.InitWindow(Wwidth, Wheight, "raylib [core] example - basic window")
	rl.SetTargetFPS(30)
}

func input() {
	if rl.IsKeyDown(rl.KeyW) && rec_left.Y > 5 {
		rec_left.Y -= speed
	} else if rl.IsKeyDown(rl.KeyS) && rec_left.Y < 495 {
		rec_left.Y += speed
	}

	if rl.IsKeyDown(rl.KeyUp) && rec_right.Y > 5 {
		rec_right.Y -= speed
	} else if rl.IsKeyDown(rl.KeyDown) && rec_right.Y < 495 {
		rec_right.Y += speed
	}
}

func collision_handling() {

}

func draw_on_screen() {
	rl.DrawRectangleRec(rec_left, rl.Lime)
	rl.DrawRectangleRec(rec_right, rl.Lime)
}

func rendering() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.RayWhite)

	draw_on_screen()

	rl.EndDrawing()
}

func quit() {
	defer rl.CloseWindow()
}

func main() {
	for !rl.WindowShouldClose() {
		input()
		rendering()
	}
	quit()
}
