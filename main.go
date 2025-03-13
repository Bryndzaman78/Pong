package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	Wheight = 600
	Wwidth  = 1000
)

var (
	speed float32 = 10

	rec_left = rl.NewRectangle(5, 5, 40, 100)

	rec_right = rl.NewRectangle(Wwidth-45, 5, 40, 100)

	ball       = rl.NewVector2(500, 300)
	ball_speed = rl.NewVector2(10, 6)

	line_up1, line_up2     = rl.NewVector2(0, 0), rl.NewVector2(1000, 0)
	line_down1, line_down2 = rl.NewVector2(0, 600), rl.NewVector2(1000, 600)
)

func init() {
	rl.InitWindow(Wwidth, Wheight, "Pong Golang/Raylib Exemple")
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

	if ball.X > 1000 || ball.X < 0 {
		ball.X, ball.Y = 500, 300
		ball_speed.Y = float32(rand.Intn(8))
	}
	ball.X += ball_speed.X
	ball.Y += ball_speed.Y
}

func collision_handling() {
	if rl.CheckCollisionCircleRec(ball, 30, rec_left) {
		ball_speed.X = 15
		ball_speed.Y = float32(rand.Intn(13))
	} else if rl.CheckCollisionCircleRec(ball, 30, rec_right) {
		ball_speed.X = -15
		ball_speed.Y = -float32(rand.Intn(13))
	}
	if rl.CheckCollisionCircleLine(ball, 30, line_up1, line_up2) {
		ball_speed.Y = float32(rand.Intn(12))
	} else if rl.CheckCollisionCircleLine(ball, 30, line_down1, line_down2) {
		ball_speed.Y = -float32(rand.Intn(12))
	}
}

func draw_on_screen() {
	rl.DrawCircle(int32(ball.X), int32(ball.Y), 30, rl.Black)

	rl.DrawRectangleRec(rec_left, rl.White)
	rl.DrawRectangleRec(rec_right, rl.White)
}

func rendering() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.NewColor(17, 137, 53, 255))

	draw_on_screen()

	rl.EndDrawing()
}

func quit() {
	defer rl.CloseWindow()
}

func main() {
	for !rl.WindowShouldClose() {
		input()
		collision_handling()
		rendering()
	}
	quit()
}
