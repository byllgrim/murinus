/*
   This file is part of Murinus.

   Murinus is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   Murinus is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with Murinus.  If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	sizeMult int32 = 3
	sizeDiv  int32 = 2

	timeExitHasToBeHeldBeforeGameEnd   int = 60 * 5
	timeExitHasToBeHeldBeforeCloseGame int = 90
)

const (
	screenWidthD       int32 = (1280 * sizeMult) / sizeDiv
	screenHeightD      int32 = (800 * sizeMult) / sizeDiv
	blockSizeD         int32 = (48 * sizeMult) / sizeDiv
	blockSizeBigBoardD int32 = (24 * sizeMult) / sizeDiv
	gSize              int32 = 12
)

var screenWidth, screenHeight, blockSize, blockSizeBigBoard int32

var quit, lostLife bool

var (
	window      *sdl.Window
	renderer    *sdl.Renderer
	input       *Input
	menus       []*Menu
	stage       *Stage
	highscores  Highscores
	defaultName string
)

func main() {
	Init()
	defer CleanUp()

	for !quit {
		difficulty = -1
		menuChoice := -1
	menuLoop:
		for difficulty == -1 && !quit {
			menuChoice = menus[0].Run(renderer, input)
			switch menuChoice {
			case -1:
				if !quit {
					quit = Arcade
				}
				break menuLoop
			case 0:
				fallthrough
			case 1:
				StartGameSession(menuChoice)
			case 2:
				fmt.Println("Not made yet") //Training
			case 3:
				highscores.Display(-1, false, renderer, input)
			case 4:
				DoSettings(menus[3], renderer, input)
			case 5:
				ShowCredits()
			case 6:
				quit = true
			default:
				panic("Unknown menu option")
			}
		}
	}
	fmt.Println("Quit")
}

func Init() {
	screenWidth, screenHeight, blockSize, blockSizeBigBoard =
		screenWidthD, screenHeightD, blockSizeD, blockSizeBigBoardD
	newScreenWidth, newScreenHeight = screenWidth, screenHeight

	runtime.LockOSThread()
	err := sdl.Init(sdl.INIT_EVERYTHING)
	PanicOnError(err)
	fmt.Println("Init SDL")

	window, err = sdl.CreateWindow("Murinus", sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED, int(screenWidth), int(screenHeight),
		sdl.WINDOW_SHOWN|sdl.WINDOW_RESIZABLE|sdl.RENDERER_PRESENTVSYNC)
	PanicOnError(err)
	fmt.Println("Created window")

	renderer, err = sdl.CreateRenderer(window, -1,
		sdl.RENDERER_ACCELERATED)
	PanicOnError(err)
	renderer.Clear()
	fmt.Println("Created renderer")

	InitText(renderer)
	fmt.Println("Initiated text")
	InitNumbers(renderer)
	fmt.Println("Initiated numbers")

	input = GetInput()
	fmt.Println("Got inputs")

	ReadOptions("options.xml", input)
	fmt.Println("Created options")

	menus = GetMenus(renderer)
	fmt.Println("Created menus")

	stage = LoadTextures(renderer, input)
	fmt.Println("Loaded stage-basis")

	highscores = Read("singleplayer.hs", "multiplayer.hs")

	fmt.Println("Loaded Highscores")

	defaultName = "\\\\\\\\\\"
}

func CleanUp() {
	highscores.Write("singleplayer.hs", "multiplayer.hs")
	if !Arcade {
		SaveOptions("options.xml", input)
	}
	numbers.Free()
	for i := 0; i < len(menus); i++ {
		menus[i].Free()
	}
	stage.Free()
	renderer.Destroy()
	window.Destroy()
}

func StartGameSession(menuChoice int) {
	difficulty = menus[1].Run(renderer, input)
	stage.ID = -1
	for !quit {
		levelsCleared := 0
		score := -ScoreMult(500)

		RunGame(menuChoice, &levelsCleared, &score);

		fmt.Printf("Game Over. Final score %d\n", score)
		stage.lostOnce = true
		input.exit.timeHeld = 0

		if !GameOverMenu(levelsCleared, score) {
			break
		}
	}
}

func ShowCredits() {
	txt := "Made by ITR   -   Source available on github.com/ITR13/murinus"
	creds, src, dst := GetText(txt, sdl.Color{255, 255, 255, 255},
		newScreenWidth/2, newScreenHeight/2, renderer)
	input.mono.a.down = false
	input.mono.b.down = false
	dst.X -= dst.W
	dst.Y -= dst.H / 2
	dst.W *= 2
	dst.H *= 2

	renderer.SetRenderTarget(nil)
	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.Clear()
	PanicOnError(renderer.Copy(creds, src, dst))

	for !input.mono.a.down && !input.mono.b.down && !quit {
		renderer.SetRenderTarget(nil)
		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()
		PanicOnError(renderer.Copy(creds, src, dst))
		renderer.Present()
		input.Poll()
	}
	creds.Destroy()
}

func RunGame(menuChoice int, levelsCleared *int, score *int64) {
	lostLife = false
	lives := 3
	wonInARow := -2
	extraLives := 0
	extraLivesCounter := int64(25000)

	for !quit && (lives != 1 || !lostLife) {
		var engine *Engine
		if lostLife {
			wonInARow = -1
			lostLife = false
			lives--
			if lives < extraLives {
				extraLives = lives
			}
			if lives == 0 {
				panic("Should not reach this statement")
			}
			engine = stage.Load(stage.ID, false, *score, menuChoice)
			window.SetTitle("Score: " + strconv.Itoa(int(*score)) +
				" Lives: " + strconv.Itoa(lives))
		} else {
			*levelsCleared++
			wonInARow++
			if wonInARow == 3 {
				if lives-extraLives < 4 {
					wonInARow = 0
					lives++
				}
			}
			fmt.Printf("Won in a row counter: %d\n", wonInARow)
			engine = stage.Load(stage.ID+1, true,
				*score + ScoreMult(500), menuChoice)
		}
		fmt.Printf("Lives: %d\n", lives)
		if engine == nil {
			fmt.Println("Engine nil, game was won")
			break
		}
		PlayStage(engine, window, renderer, int32(lives))
		*score = engine.Score
		if engine.Input.exit.timeHeld >
			timeExitHasToBeHeldBeforeCloseGame {
			fmt.Println("Game was quit with exit key")
			break
		}
		for *score > extraLivesCounter &&
			extraLivesCounter*2 > extraLivesCounter {
			extraLivesCounter *= 2
			//extraLives++
			//lives++
		}
		fmt.Printf("Score: %d\n", *score)
	}
}

func PlayStage(engine *Engine, window *sdl.Window, renderer *sdl.Renderer,
	lives int32) {
	p1C, p2C := options.CharacterP1, options.CharacterP2
	if engine.p1 == nil {
		p1C = p2C
	} else if engine.p2 == nil {
		p2C = p1C
	}

	quit = false
	lostLife = false
	score := int32(0)
	engine.Stage.scores.score, engine.Stage.scores.lives = engine.Score, lives
	for i := 0; i < 30 && !quit; i++ {
		engine.Stage.Render(p1C, p2C, renderer, false)
		engine.Input.Poll()
		if engine.Input.exit.timeHeld > timeExitHasToBeHeldBeforeCloseGame {
			fmt.Println("Round was quit with exit key")
			return
		}
	}

	for noKeysTouched >= 5 && !quit {
		engine.Stage.Render(p1C, p2C, renderer, false)
		engine.Input.Poll()
	}
	fmt.Println("Finished starting animation")

	engine.Stage.tiles.renderedOnce = false
	for !quit {
		engine.Input.Poll()
		if engine.Input.exit.timeHeld > timeExitHasToBeHeldBeforeCloseGame {
			fmt.Println("Round was quit with exit key")
			return
		}

		engine.Advance()
		window.SetTitle("Murinus (score: " +
			strconv.Itoa(int(score)) +
			", left " + strconv.Itoa(engine.Stage.pointsLeft) + ")")

		engine.Stage.scores.score = engine.Score
		engine.Stage.Render(p1C, p2C, renderer, true)
		if engine.Stage.pointsLeft <= 0 || lostLife {
			break
		}
	}
	fmt.Println("Exited play loop")

	for i := 0; i < len(engine.snakes); i++ {
		snake := engine.snakes[i]
		snake.head.display = true
		for j := 0; j < len(snake.body); j++ {
			snake.body[j].display = true
		}
		snake.tail.display = true
	}

	if lostLife {
		for i := 0; i < 90 && !quit; i++ {
			engine.Stage.tiles.renderedOnce = false
			engine.Stage.scores.lives = lives - int32(i/15%2)
			p1C, p2C := options.CharacterP1, options.CharacterP2
			if engine.p1 != nil {
				engine.p1.entity.display = (i / 15 % 2) == 0
			} else {
				p1C = p2C
			}

			if engine.p2 != nil {
				engine.p2.entity.display = (i / 15 % 2) == 0
			} else {
				p2C = p1C
			}

			engine.Stage.Render(p1C, p2C, renderer, false)
			engine.Input.Poll()
			if engine.Input.exit.timeHeld > timeExitHasToBeHeldBeforeCloseGame {
				fmt.Println("Round was quit with exit key")
				return
			}
		}
	} else {
		for i := 0; i < 30 && !quit; i++ {
			engine.Stage.Render(p1C, p2C, renderer, false)
			engine.Input.Poll()
		}
	}
	fmt.Println("Finished exit animation")
}

func DoSettings(menu *Menu, renderer *sdl.Renderer, input *Input) {
	for v := menu.Run(renderer, input); v != -1 &&
		!quit; v = menu.Run(renderer, input) {
		ReadOptions("", input)
		menu.menuItems[0].SetNumber(int32(options.CharacterP1), renderer)
		menu.menuItems[1].SetNumber(int32(options.CharacterP2), renderer)
		menu.menuItems[2].SetNumber(int32(options.EdgeSlip), renderer)
		menu.menuItems[3].SetNumber(int32(options.BetterSlip), renderer)
		menu.menuItems[4].SetNumber(int32(options.ShowDivert), renderer)
	}
	if quit {
		return
	}
	options.CharacterP1 = uint8(menu.menuItems[0].numberField.Value)
	options.CharacterP2 = uint8(menu.menuItems[1].numberField.Value)
	options.EdgeSlip = int(menu.menuItems[2].numberField.Value)
	options.BetterSlip = menu.menuItems[3].numberField.Value
	options.ShowDivert = uint8(menu.menuItems[4].numberField.Value)
	options.showDivert = options.ShowDivert != 0
	redrawTextures = true
}

func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func LogOnError(err error) bool {
	if err != nil {
		fmt.Println(err)
	}
	return err != nil
}

func GameOverMenu(levelsCleared int, score int64) (resume bool) {
	menuChoice := -1
	var scoreData *ScoreData
	menus[2].selectedElement = 0
	for !quit && menuChoice < 2 {
		menuChoice = menus[2].Run(renderer, input)
		if menuChoice == 0 { // Set name
			name := GetName(defaultName, renderer, input)
			if name != "" {
				defaultName = name
				if scoreData == nil {
					scoreData = &ScoreData{score, name,
						levelsCleared, difficulty,
						time.Now()}
					highscores.Add(scoreData,
						menuChoice != 0, true)
				} else {
					scoreData.Name = name
				}
			}
		} else if menuChoice == 1 { // Highscores
			highscores.Display(difficulty, menuChoice != 0,
				renderer, input)
		} else if menuChoice == -1 {
			menuChoice = 4
		}
	}

	resume = true
	if menuChoice == 2 { // Continue
		stage.ID--
	} else if menuChoice == 3 { // Restart
		stage.ID = -1
	} else if menuChoice == 4 { // Exit to menu
		resume = false
	} else {
		panic("Unknown menu option")
	}

	return
}
