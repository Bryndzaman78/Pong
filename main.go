// ! This is my first project am biginner programmer so if somethig is wrong or could be better I appreciate the help :)
package main

import (
	// Here are some standard golang liberaries
	"math/rand"
	"strconv"

	// Raylib
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	// Constances about window height and width
	Wheight = 600
	Wwidth  = 1000
)

var (
	// Speed of the two rectangles controlled by player.
	speed float32 = 10

	// Positions and sizes of the two rectangles.
	rec_right, rec_left = rl.NewRectangle(Wwidth-45, 250, 40, 100), rl.NewRectangle(5, 250, 40, 100)

	// Here are wariables about the ball like speed,positon and radius. Note: ball_increase_speed is one variable becouse its applied to both ways
	ball_speed, ball, ball_radius         = rl.NewVector2(10, 6), rl.NewVector2(500, 300), 30
	ball_increase_speed           float32 = 12

	// Here are variables about screen bottom and top lines for ball to collide.
	line_up1, line_up2     = rl.NewVector2(0, 0), rl.NewVector2(1000, 0)
	line_down1, line_down2 = rl.NewVector2(0, 600), rl.NewVector2(1000, 600)

	// Score count
	score_left, score_right = 0, 0
)

// Init main func.
func init() {
	rl.InitWindow(Wwidth, Wheight, "Pong Golang/Raylib Exemple")
	rl.SetTargetFPS(30)
}

// Input func for keys and movement.
func input() {
	// Restart score and speed if:
	if rl.IsKeyPressed(rl.KeyR) || score_left == 15 || score_right == 15 {
		ball.X, ball.Y = 500, 300
		score_left, score_right = 0, 0
		ball_increase_speed = 12
		rec_left.Y, rec_right.Y = 250, 250
	}

	// Moving the right rectangle up.
	if rl.IsKeyDown(rl.KeyW) && rec_left.Y > 10 {
		rec_left.Y -= speed
		// Moving the right rec down.
	} else if rl.IsKeyDown(rl.KeyS) && rec_left.Y < 490 {
		rec_left.Y += speed
	}

	// Moving the left rec up.
	if rl.IsKeyDown(rl.KeyUp) && rec_right.Y > 10 {
		rec_right.Y -= speed
		//Moving the left rec down.
	} else if rl.IsKeyDown(rl.KeyDown) && rec_right.Y < 490 {
		rec_right.Y += speed
	}

	// If ball is behind rectangles reset position of ball.
	if ball.X > 960 || ball.X < 40 {
		ball.X, ball.Y = 500, 300
		ball_speed.Y = float32(rand.Intn(8))
	}

	// Ball movement looped in frame loop.
	ball.X += ball_speed.X
	ball.Y += ball_speed.Y
}

// Func for collision handling.
func collision_handling() {
	// Collisons between rectangles
	if rl.CheckCollisionCircleRec(ball, float32(ball_radius), rec_left) {
		ball_speed.X = ball_increase_speed
		// Random Directon at Y after impact.
		ball_speed.Y = float32(rand.Intn(16))
		// Adding speed to ball.
		if ball_increase_speed <= 24 {
			ball_increase_speed += 1
		}
	} else if rl.CheckCollisionCircleRec(ball, float32(ball_radius), rec_right) {
		ball_speed.X = -ball_increase_speed
		// -||-
		ball_speed.Y = -float32(rand.Intn(16))
		// -||-
		if ball_increase_speed <= 24 {
			ball_increase_speed += 1
		}
	}

	// checking collision between top and bottom of screean and bouncing the ball away.
	if rl.CheckCollisionCircleLine(ball, float32(ball_radius), line_up1, line_up2) {
		ball_speed.Y = float32(rand.Intn(15))
	} else if rl.CheckCollisionCircleLine(ball, float32(ball_radius), line_down1, line_down2) {
		ball_speed.Y = -float32(rand.Intn(15))
	}

	// If ball behing rectangles adding score.
	if ball.X > 960 {
		score_left += 1
	} else if ball.X < 40 {
		score_right += 1
	}
}

// Draws on Screen.
func draw_on_screen() {
	// Ball black colour.
	rl.DrawCircle(int32(ball.X), int32(ball.Y), 20, rl.Black)

	// Rectangles white colour.
	rl.DrawRectangleRec(rec_left, rl.White)
	rl.DrawRectangleRec(rec_right, rl.White)

	// Score count and speed count
	rl.DrawText("Ball Speed:"+strconv.Itoa(int(ball_increase_speed)), 60, 5, 20, rl.White)
	rl.DrawText(strconv.Itoa(score_left)+"|"+strconv.Itoa(score_right), Wwidth/2-40, 5, 50, rl.White)
}

// Background and drawing on screen.
// ** NOTE: Func drawing on screen is not needed and all can be included here.
func rendering() {
	rl.BeginDrawing()

	// adds green background from rgb colour.
	rl.ClearBackground(rl.NewColor(17, 137, 53, 255))

	draw_on_screen()

	rl.EndDrawing()
}

// Function to close window. It is needed if you dont want your window to run infinitly.
func quit() {
	defer rl.CloseWindow()
}

// Main function every thing must be includet here or can directly be written here.
func main() {
	// Main window loop.
	for !rl.WindowShouldClose() {
		input()
		collision_handling()
		rendering()
	}
	quit()
}

//! I hope you enjoyed :)
