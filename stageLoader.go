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
)

const (
	stageWidth  int32 = 25
	stageHeight int32 = 15
)

var tiles [][]Tile

var difficulty int

type PreStageData struct {
	stage          string
	px, py         int32
	difficultyData []*PreDifficultyData
}

type PreDifficultyData struct {
	playerSpeed int32
	snakes      []PreSnakeData
}

type PreSnakeData struct {
	x, y                 int32
	length               int
	ai                   AI
	moveTimerMax         int
	growTimerMax         int
	minLength, maxLength int
}

func GetPreStageData(stage string, px, py int32,
	snakes []PreSnakeData, speed [][]int32) *PreStageData {

	diffData := make([]*PreDifficultyData, len(speed))
	for i := 0; i < len(diffData); i++ {
		diffData[i] = &PreDifficultyData{
			0, make([]PreSnakeData, len(snakes)),
		}
		for j := 0; j < len(snakes); j++ {
			diffData[i].snakes[j] = snakes[j]
		}
	}

	for i := 0; i < len(speed); i++ {
		diffData[i].playerSpeed = speed[i][0]
		for j := 1; j < len(speed[i]); j++ {
			diffData[i].snakes[j-1].moveTimerMax = int(speed[i][j])
		}
	}

	return &PreStageData{stage, px, py, diffData}
}

func GetPreStageDatas() ([]*PreStageData, [5][][2]int) {
	pSpeed := PrecisionMax / 4
	data := []*PreStageData{
		GetPreStageData(""+
			"#########################"+
			"#########################"+
			"#########################"+
			"#########################"+
			"#########****############"+
			"#########*##*############"+
			"#########*##*############"+
			"#########***0***#########"+
			"############*##*#########"+
			"############*##*#########"+
			"############****#########"+
			"#########################"+
			"#########################"+
			"#########################"+
			"#########################",
			stageWidth/2, stageHeight/2,
			nil, [][]int32{{pSpeed}, {pSpeed * 3 / 2}, {pSpeed * 2}}),
		GetPreStageData(""+
			"#########################"+
			"#########################"+
			"#########################"+
			"###0*****################"+
			"########**###############"+
			"#########**#0#0#0########"+
			"##########*#0#0#0########"+
			"##########***000000######"+
			"############*#0#0########"+
			"##########00***0000######"+
			"############0#*#0########"+
			"##########0000***00######"+
			"############0#0#0########"+
			"############0#0#0########"+
			"#########################",
			3, 3,
			nil, [][]int32{{pSpeed}, {pSpeed * 3 / 2}, {pSpeed * 2}}),
		GetPreStageData(""+
			"#########################"+
			"#########################"+
			"#########################"+
			"#########################"+
			"#########################"+
			"#####0*************3#####"+
			"#####*#############*#####"+
			"#####*#############*#####"+
			"#####*#############*#####"+
			"#####4*************3#####"+
			"#########################"+
			"#########################"+
			"#########################"+
			"#########################"+
			"#########################",
			5, 5,
			[]PreSnakeData{{stageWidth - 6, stageHeight - 6,
				1, &SimpleAI{}, 20, 1, 1, 6}},
			[][]int32{{pSpeed}, {pSpeed * 3 / 2}, {pSpeed * 2}}),
		nil,
		GetPreStageData(""+
			"#########################"+
			"#########################"+
			"#########################"+
			"####*****************####"+
			"####*####0#####0####*####"+
			"####*####4#####0####*####"+
			"####*####0#####0####*####"+
			"####*****************####"+
			"####*####0#####0####*####"+
			"####*####4#####0####*####"+
			"####*####0#####0####*####"+
			"####*****************####"+
			"#########################"+
			"#########################"+
			"#########################",
			4, 3,
			[]PreSnakeData{{stageWidth - 5, stageHeight - 4, 6,
				&SimpleAI{0, false, 0, 0, Left},
				0, 10 * 4, 2, 16}},
			[][]int32{{pSpeed, 9},
				{pSpeed * 3 / 2, 4},
				{pSpeed * 2, 0}}),
		GetPreStageData(""+
			"#########################"+
			"########*********########"+
			"#######00#0###0#00#######"+
			"######00##00000##00######"+
			"######0####0#0####0######"+
			"######***##4#4##***######"+
			"######*#**00000**#*######"+
			"######*##*##0##*##*######"+
			"######*##*##4##*##*######"+
			"######****##0##****######"+
			"######0####000####0######"+
			"######0###00#00###0######"+
			"######0###0###0###0######"+
			"######0000003000000######"+
			"#########################",
			stageWidth/2, 1,
			[]PreSnakeData{{stageWidth / 2, stageHeight - 2,
				6, &SimpleAI{0, true, 0, 5, Up}, 0, 10 * 4, 2, 16}},
			[][]int32{{pSpeed, 9},
				{pSpeed * 3 / 2, 4},
				{pSpeed * 2, 0}}),
		GetPreStageData(""+
			"#########################"+
			"#########################"+
			"#########################"+
			"###########***###########"+
			"###########*#*###########"+
			"##########*****##########"+
			"########***#%#***########"+
			"########*#*%0%*#*########"+
			"########***#%#***########"+
			"##########*****##########"+
			"###########*#*###########"+
			"###########***###########"+
			"#########################"+
			"#########################"+
			"#########################",
			stageWidth/2, stageHeight/2,
			[]PreSnakeData{{12, 3, 6, &SimpleAI{},
				0, 10 * 4, 2, 16}},
			[][]int32{{pSpeed, 9},
				{pSpeed * 3 / 2, 4},
				{pSpeed * 2, 1}}),
		GetPreStageData(""+
			"#########################"+
			"#00000000000000000000003#"+
			"#0#####0#########0#####0#"+
			"#0#000000000**********#0#"+
			"#0#0########*########*#0#"+
			"#0#0#0000000********#*#0#"+
			"#0#0#0######%######*#*#0#"+
			"#000#0%00000000000%*#***#"+
			"#0#0#0######%######*#0#*#"+
			"#0#0#0000000********#0#*#"+
			"#0#0########*########0#*#"+
			"#0#000000000******0000#*#"+
			"#0#####0#########*#####*#"+
			"#0000000000000000*******#"+
			"#########################",
			stageWidth/2, stageHeight/2,
			[]PreSnakeData{{stageWidth - 2, 1, 6, &SimpleAI{},
				0, 10 * 4, 2, 16}},
			[][]int32{{pSpeed, 9},
				{pSpeed * 3 / 2, 4},
				{pSpeed * 2, 0}}),
		nil,
		GetPreStageData(""+
			"#########################"+
			"#*******0000400000000003#"+
			"#*#####*#########0#####0#"+
			"#*#0000***************#0#"+
			"#*#0########*########*#0#"+
			"#*#0#***************#*#0#"+
			"#*#0#*######%######*#*#0#"+
			"#***#*%*****0*****%*#***#"+
			"#0#*#*######%######*#0#*#"+
			"#0#*#***************#0#*#"+
			"#0#*########*########0#*#"+
			"#0#***************0000#*#"+
			"#0#####0#########*#####*#"+
			"#3000000000040000*******#"+
			"#########################",
			stageWidth/2, stageHeight/2,
			[]PreSnakeData{{1, stageHeight - 2, 6, &SimpleAI{},
				0, 10 * 4, 2, 16},
				{stageWidth - 2, 1, 6, &SimpleAI{},
					0, 10 * 4, 2, 16}},
			[][]int32{{pSpeed, 9, 9},
				{pSpeed * 3 / 2, 4, 4},
				{pSpeed * 2, 0, 0}}),
		GetPreStageData(""+
			"#########################"+
			"#*********************00#"+
			"#*###########0#######0#0#"+
			"#*###########0##0000#0#0#"+
			"#*###########0##0##0#0#4#"+
			"#*###########0##0##000#0#"+
			"#*###########0##0######0#"+
			"#***********00**********#"+
			"#0######0##*###########*#"+
			"#0#000##0##*###########*#"+
			"#0#0#0##0##*###########*#"+
			"#0#0#0000##*###########*#"+
			"#0#0#######*###########*#"+
			"#0000000000*************#"+
			"#########################",
			stageWidth/2, stageHeight/2,
			[]PreSnakeData{{1, stageHeight - 2, 6, &SimpleAI{0, false, 0, 4, 0},
				0, 10 * 4, 2, 6},
				{stageWidth - 2, 1, 6, &SimpleAI{0, false, 0, 4, 0},
					0, 10 * 4, 2, 6}},
			[][]int32{{pSpeed, 9, 9},
				{pSpeed * 3 / 2, 4, 4},
				{pSpeed * 2, 0, 0}}),
		GetPreStageData(""+
			"#########################"+
			"##00000******0000000000##"+
			"##0####*####*####0####0##"+
			"##0####*####*####0####0##"+
			"##0####*####*####0####0##"+
			"##0####*####*00000####0##"+
			"##40000*####*####0####0##"+
			"##0####*####*####0####0##"+
			"##0####*####*####000300##"+
			"##0####*0000*####0####0##"+
			"##0####*####*####0####0##"+
			"##0####*####*####0####0##"+
			"##0####*####*####0####0##"+
			"##00000******0000000000##"+
			"#########################",
			stageWidth/2, stageHeight/2,
			[]PreSnakeData{{2, stageHeight - 2, 6,
				&SimpleAI{0, false, 0, 0, Up}, 0, 10 * 4, 2, 6},
				{stageWidth - 3, 1, 6, &SimpleAI{0, false, 0, 0, Down},
					0, 10 * 4, 2, 6}},
			[][]int32{{pSpeed, 9, 9},
				{pSpeed * 3 / 2, 4, 4},
				{pSpeed * 2, 0, 0}}),
		GetPreStageData(""+
			"#########################"+
			"#########################"+
			"############%############"+
			"###########*4*###########"+
			"##########**#**##########"+
			"#########**###**#########"+
			"########**#####**########"+
			"#######**##*0*##**#######"+
			"######**##**#**##**######"+
			"#####**###*###*###**#####"+
			"#####*####*****####*#####"+
			"#####**#####%#####**#####"+
			"######*************######"+
			"#########################"+
			"#########################",
			stageWidth/2, stageHeight/2,
			[]PreSnakeData{{stageWidth - 7, stageHeight - 3, 6, &SimpleAI{},
				0, 10 * 4, 2, 4},
				{5, stageHeight - 4, 6, &SimpleAI{0, false, 0, 0, Right},
					0, 10 * 4, 2, 4},
				{stageWidth / 2, 3, 6, &SimpleAI{},
					0, 10 * 4, 2, 4}},
			[][]int32{{pSpeed, 10, 10, 10},
				{pSpeed * 3 / 2, 6, 6, 6},
				{pSpeed * 2, 2, 2, 2}}),
		nil,
		GetPreStageData(""+
			"#########################"+
			"#########################"+
			"#########################"+
			"#########################"+
			"######*****000*****######"+
			"######*###########*######"+
			"######*###########*######"+
			"######*###########*######"+
			"######*######0440#*######"+
			"######*######0##0#*######"+
			"######*************######"+
			"#########################"+
			"#########################"+
			"#########################"+
			"#########################",
			stageWidth/2, stageHeight/2-3,
			[]PreSnakeData{{stageWidth/2 + 1, stageHeight - 5, 4,
				&ApproximatedAI{0, 1000}, 1, 10 * 4, 2, 4}},
			[][]int32{{pSpeed, 9},
				{pSpeed * 3 / 2, 4},
				{pSpeed * 3 / 2, 2}}),
		GetPreStageData(""+
			"#########################"+
			"#########################"+
			"#########################"+
			"####*****************####"+
			"####*####*#####*####*####"+
			"####*####*#####*####*####"+
			"####*####*#####*####*####"+
			"####*****************####"+
			"####*####*#####*####*####"+
			"####*####*#####*####*####"+
			"####*####*#####*####*####"+
			"####*****************####"+
			"#########################"+
			"#########################"+
			"#########################",
			stageWidth-5, stageHeight-5,
			[]PreSnakeData{{4, 4, 6, &ApproximatedAI{0, 3},
				1, 10 * 4, 2, 12}},
			[][]int32{{pSpeed, 9},
				{pSpeed * 3 / 2, 4},
				{pSpeed * 2, 1}}),
		GetPreStageData(""+
			"#########################"+
			"#########################"+
			"#########################"+
			"####*******###3000005####"+
			"####*#####*###0#0#0#0####"+
			"####*#####*###0#000#0####"+
			"####*#####*###0#0#0#0####"+
			"####*******%0%000#000####"+
			"####*#####*###0#0#0#0####"+
			"####*#####*###0#000#0####"+
			"####*#####*###0#0#0#0####"+
			"####*******###0000005####"+
			"#########################"+
			"#########################"+
			"#########################",
			stageWidth/2, stageHeight/2,
			[]PreSnakeData{{stageWidth/2 + 2, stageHeight / 2, 6,
				&ApproximatedAI{0, 3}, 1, 10 * 4, 2, 10},
				{stageWidth/2 - 2, stageHeight / 2, 6,
					&ApproximatedAI{0, 3}, 1, 10 * 4, 2, 14}},
			[][]int32{{pSpeed, 9, 7},
				{pSpeed * 3 / 2, 4, 3},
				{pSpeed * 2, 2, 1}}),
		GetPreStageData(""+
			"#########################"+
			"#########################"+
			"#########################"+
			"########3*******2########"+
			"###########*#*###########"+
			"###########*#*###########"+
			"###########*0*###########"+
			"###########*#*###########"+
			"###########*0*###########"+
			"###########*#*###########"+
			"###########*#*###########"+
			"###########***###########"+
			"#########################"+
			"#########################"+
			"#########################",
			stageWidth/2, stageHeight/2-1,
			[]PreSnakeData{{stageWidth / 2, stageHeight/2 + 4, 6,
				&ApproximatedAI{0, 19}, 1, 10 * 4, 2, 5}},
			[][]int32{{pSpeed, 9},
				{pSpeed * 3 / 2, 6},
				{pSpeed * 2, 3}}),
		GetPreStageData(""+
			"#########################"+
			"#0**********#***********#"+
			"#0#*#######*#*#*#4#4#4#*#"+
			"#0#*0003000***#*********#"+
			"#0#*#######*#*#########*#"+
			"#4#*********#***********#"+
			"#0#######*#*#*#*#########"+
			"#00000000*#*0*#*********#"+
			"#########*#*#*#*#######*#"+
			"#***********#***000000#*#"+
			"#0#########*#*#######0#*#"+
			"#000000000#***00000000#*#"+
			"#0#0#0#0#0#0#0#######0#*#"+
			"#00000000000#00000000003#"+
			"#########################",
			stageWidth/2, stageHeight/2,
			[]PreSnakeData{{1, stageHeight - 2, 6, &ApproximatedAI{0, 4},
				1, 10 * 4, 2, 16},
				{stageWidth - 2, 1, 6, &ApproximatedAI{0, 4},
					1, 10 * 4, 2, 16}},
			[][]int32{{pSpeed, 9, 9},
				{pSpeed * 3 / 2, 5, 5},
				{pSpeed * 2, 3, 3}}),
		nil,
		GetPreStageData(""+
			"#########################"+
			"#0**********#00000000000#"+
			"#0#*#######*#0#0#6#6#6#0#"+
			"#0#*0003000*00#000000000#"+
			"#0#*#######*#0#########0#"+
			"#4#*********#00000000000#"+
			"#0#######*#*#0#0#########"+
			"#00000000*#***#000000000#"+
			"#########*#*#*#0#######0#"+
			"#***********#*00000000#0#"+
			"#0#########*#*#######0#0#"+
			"#000000000#***00000000#0#"+
			"#0#0#0#0#0#0#0#######0#0#"+
			"#00000000000#00000000003#"+
			"#########################",
			stageWidth/2, stageHeight/2,
			[]PreSnakeData{{1, stageHeight - 2, 3, &ApproximatedAI{0, 3},
				1, 10 * 4, 1, 3},
				{stageWidth - 2, 1, 3, &ApproximatedAI{0, 3},
					1, 10 * 4, 1, 3},
				{stageWidth - 2, stageHeight - 2, 6, &SimpleAI{},
					1, 10 * 4, 2, 16},
				{1, 1, 6, &SimpleAI{},
					1, 10 * 4, 2, 16}},
			[][]int32{{pSpeed, 11, 11, 9, 9},
				{pSpeed * 3 / 2, 7, 7, 5, 5},
				{pSpeed * 2, 5, 5, 3, 3}}),
		GetPreStageData(""+
			"#########################"+
			"#***********************#"+
			"#*#0#0#0#0#0#0#0#0#0#0#0#"+
			"#***********************#"+
			"#0#0#0#0#0#0#0#0#0#0#0#*#"+
			"#***********************#"+
			"#*#0#0#0#0#0#0#0#0#0#0#0#"+
			"#***********000000000000#"+
			"#0#0#0#0#0#0#0#0#0#0#0#0#"+
			"#00000000000000000000000#"+
			"#0#0#5#0#0#0#0#0#0#5#0#0#"+
			"#00000000040004000000000#"+
			"#0#0#0#0#4#0#0#4#0#0#0#0#"+
			"#40000000000000000000004#"+
			"#########################",
			stageWidth/2, stageHeight/2,
			[]PreSnakeData{{stageWidth - 2, stageHeight - 2, 3,
				&ApproximatedAI{0, 3}, 1, 10 * 4, 1, 3},
				{1, stageHeight - 2, 3, &ApproximatedAI{0, 3},
					1, 10 * 4, 1, 3},
				{1, 1, 6, &SimpleAI{},
					1, 10 * 4, 2, 16},
				{stageWidth - 2, 1, 6, &SimpleAI{},
					1, 10 * 4, 2, 16}},
			[][]int32{{pSpeed, 11, 11, 9, 9},
				{pSpeed * 3 / 2, 7, 7, 5, 5},
				{pSpeed * 2, 5, 5, 3, 3}}),
		GetPreStageData(""+
			"#########################"+
			"#***********#***00000000#"+
			"#*#######*#*#*#*#######0#"+
			"#*#######*#*#*#*#######0#"+
			"#*********#***#*0*******#"+
			"#0###0#0#**###**#*#*###*#"+
			"#4###0#00#**#**#**#*###*#"+
			"#0###5##0%#*0*#%*##*###*#"+
			"#4###0#00#00#**#**#*###*#"+
			"#0###0#0#00###**#*#*###*#"+
			"#000000000#000#*********#"+
			"#0#######0#0#0#0#######0#"+
			"#0#######0#0#0#0#######0#"+
			"#50000000000#00400000003#"+
			"#########################",
			stageWidth/2, stageHeight/2,
			[]PreSnakeData{{1, stageHeight - 2, 3, &ApproximatedAI{0, 4},
				1, 10 * 4, 1, 3},
				{stageWidth - 2, 1, 3, &ApproximatedAI{0, 4},
					1, 10 * 4, 1, 3},
				{stageWidth - 2, stageHeight - 2, 6, &SimpleAI{},
					1, 10 * 4, 2, 16},
				{1, 1, 6, &SimpleAI{},
					1, 10 * 4, 2, 16}},
			[][]int32{{pSpeed, 11, 11, 9, 9},
				{pSpeed * 3 / 2, 7, 7, 5, 5},
				{pSpeed * 2, 5, 5, 3, 3}}),
		GetPreStageData(""+
			"#########################"+
			"#****####0004000####****#"+
			"#*##**###0#####0###**##*#"+
			"#*###**##0#####0##**###*#"+
			"#*####**#0004000#**####*#"+
			"#*#####*#0#####0#*#####*#"+
			"#***###***********###***#"+
			"###***######*######***###"+
			"###*#***************#*###"+
			"###*########*########*###"+
			"###*##****##*##****##*###"+
			"###*##*##*******##*##*###"+
			"###*##*#####*#####*##*###"+
			"#00*******************00#"+
			"#########################",
			stageWidth/2, stageHeight/2,
			[]PreSnakeData{{stageWidth - 2, stageHeight - 2, 3, &ApproximatedAI{0, 3},
				1, 10 * 4, 1, 3},
				{stageWidth - 2, 1, 3, &ApproximatedAI{0, 3},
					1, 10 * 4, 1, 3},
				{1, stageHeight - 2, 6, &SimpleAI{},
					1, 10 * 4, 2, 16},
				{1, 1, 6, &SimpleAI{},
					1, 10 * 4, 2, 16}},
			[][]int32{{pSpeed, 11, 11, 9, 9},
				{pSpeed * 3 / 2, 7, 7, 5, 5},
				{pSpeed * 2, 5, 5, 3, 3}}),
		GetPreStageData(""+
			"#########################"+
			"#***********************#"+
			"#*###*###0#####0###0###*#"+
			"#*###*###0#####0###0###*#"+
			"#*********0000000000000*#"+
			"#*###*###*#####0###0###*#"+
			"#*###*###*#####0###0###*#"+
			"#***************0000000*#"+
			"#*###*###*#####*###0###*#"+
			"#*###*###*#####*###0###*#"+
			"#*******************000*#"+
			"#*###*###*#####*###*###*#"+
			"#*###*###*#####*###*###*#"+
			"#***********************#"+
			"#########################",
			stageWidth/2, stageHeight/2,
			[]PreSnakeData{{1, stageHeight - 2, 3, &ApproximatedAI{0, 4},
				1, 10 * 4, 1, 3},
				{stageWidth - 2, 1, 3, &ApproximatedAI{0, 4},
					1, 10 * 4, 1, 3},
				{stageWidth - 2, stageHeight - 2, 6, &SimpleAI{},
					1, 10 * 4, 2, 16},
				{1, 1, 6, &SimpleAI{},
					1, 10 * 4, 2, 16}},
			[][]int32{{pSpeed, 11, 11, 9, 9},
				{pSpeed * 3 / 2, 7, 7, 5, 5},
				{pSpeed * 2, 5, 5, 3, 3}}),
		nil,
		GetPreStageData(""+
			"#########################"+
			"#*****000000000000000004#"+
			"#*###*###0#####5###0###0#"+
			"#*###*###0#####0###0###0#"+
			"#*********00000000000600#"+
			"#*###*###*#####0###0###0#"+
			"#*###*###*#####0###5###0#"+
			"#***************00000500#"+
			"#*###*###*#####*###0###0#"+
			"#*###*###*#####*###0###0#"+
			"#*******************0000#"+
			"#0###*###*#####*###*###0#"+
			"#0###*###*#####*###*###0#"+
			"#3000***************0004#"+
			"#########################",
			stageWidth/2, stageHeight/2,
			[]PreSnakeData{{1, stageHeight - 2, 3, &ApproximatedAI{0, 4},
				1, 10 * 4, 1, 3},
				{stageWidth - 2, stageHeight - 2, 3, &ApproximatedAI{0, 4},
					1, 10 * 4, 1, 3},
				{stageWidth / 2, 1, 4, &ApproximatedAI{0, 2},
					1, 10 * 4, 2, 5}},
			[][]int32{{pSpeed, 11, 11, 9},
				{pSpeed * 3 / 2, 7, 7, 5},
				{pSpeed * 2, 5, 5, 4}}),
		GetPreStageData(""+
			"#########################"+
			"#050#######***#######050#"+
			"#0#0#####***#***#####0#0#"+
			"#0#0###***#####***###0#0#"+
			"#0#0#***#########***#0#0#"+
			"#0#0#*#############*#0#0#"+
			"#0#0#***************#0#0#"+
			"#0#0#*#############*#0#0#"+
			"#0#0#***************#0#0#"+
			"#0#0#*#############*#0#0#"+
			"#0#0#*##004000400##*#0#0#"+
			"#00000##0###0###0##*0000#"+
			"#0######003000300######0#"+
			"#00000000#######00000000#"+
			"#########################",
			stageWidth/2, stageHeight/2-1,
			[]PreSnakeData{{1, stageHeight - 2, 4, &ApproximatedAI{0, 4},
				1, 10 * 4, 1, 5},
				{stageWidth - 2, stageHeight - 2, 4, &ApproximatedAI{0, 4},
					1, 10 * 4, 1, 5},
				{stageWidth / 2, 1, 5, &ApproximatedAI{0, 2},
					1, 10 * 4, 2, 7}},
			[][]int32{{pSpeed, 8, 8, 8},
				{pSpeed * 3 / 2, 5, 5, 5},
				{pSpeed * 2, 3, 3, 3}}),
		GetPreStageData(""+
			"#########################"+
			"#050#######***#######050#"+
			"#0#0#####***#***#####0#0#"+
			"#0#0###***#####***###0#0#"+
			"#0#0#***#########***#0#0#"+
			"#0#0#*#############*#0#0#"+
			"#0#0#***************#0#0#"+
			"#0#0#*#############*#0#0#"+
			"#0#0#***************#0#0#"+
			"#0#0#*#############*#0#0#"+
			"#0#0#*##004000400##*#0#0#"+
			"#00000##0###0###0##*0000#"+
			"#0######003000300######0#"+
			"#00000000#######00000000#"+
			"#########################",
			stageWidth/2, stageHeight/2-1,
			[]PreSnakeData{{1, stageHeight - 2, 3, &ApproximatedAI{0, 3},
				1, 10 * 4, 1, 4},
				{stageWidth - 2, stageHeight - 2, 3, &ApproximatedAI{0, 3},
					1, 10 * 4, 1, 4},
				{stageWidth - 2, 1, 6, &ApproximatedAI{0, 5},
					1, 10 * 4, 2, 6},
				{1, 1, 6, &ApproximatedAI{0, 5},
					1, 10 * 4, 2, 10}},
			[][]int32{{pSpeed, 11, 11, 9, 9},
				{pSpeed * 3 / 2, 7, 7, 5, 5},
				{pSpeed * 2, 5, 5, 3, 3}}),
		GetPreStageData(""+
			"#########################"+
			"#50000000000%00000000005#"+
			"#0#########0#0#########0#"+
			"#0#*******************#0#"+
			"#0#*###*###*#*###*###*#0#"+
			"#0#*#***###*#*###***#*#0#"+
			"#0#*#*#*###*#*###*#*#*#0#"+
			"#0#***#***********#***#0#"+
			"#0#*###*#########*###*#0#"+
			"#0#*******************#0#"+
			"#0#*#####*#####*#####*#0#"+
			"#0#*******************#0#"+
			"#0#####################0#"+
			"#50000000000000000000005#"+
			"#########################",
			stageWidth/2, stageHeight/2,
			[]PreSnakeData{{3, stageHeight - 4, 3, &ApproximatedAI{0, 3},
				1, 10 * 4, 1, 3},
				{stageWidth - 4, stageHeight - 4, 3, &ApproximatedAI{0, 3},
					1, 10 * 4, 1, 3},
				{stageWidth - 2, 1, 6, &ApproximatedAI{0, 5},
					1, 10 * 4, 2, 6},
				{1, 1, 6, &ApproximatedAI{0, 5},
					1, 10 * 4, 2, 10}},
			[][]int32{{pSpeed, 11, 11, 9, 9},
				{pSpeed * 3 / 2, 7, 7, 5, 5},
				{pSpeed * 2, 5, 5, 3, 3}}),
		GetPreStageData(""+
			"#########################"+
			"###########060###########"+
			"#########000#000#########"+
			"#######000#####000#######"+
			"#####000#########000#####"+
			"###000#############000###"+
			"#400#***************#004#"+
			"#0#0#*#############*#0#0#"+
			"#0#0#***************#0#0#"+
			"#0#0#*#############*#0#0#"+
			"#504#***************#405#"+
			"#0######0###0###0######0#"+
			"#0######0*******0######0#"+
			"#00000000#######00000000#"+
			"#########################",
			stageWidth/2, stageHeight/2+1,
			[]PreSnakeData{
				{1, stageHeight - 2, 4, &ApproximatedAI{0, 4},
					1, 10 * 4, 1, 5},
				{stageWidth - 2, stageHeight - 2, 4, &ApproximatedAI{0, 4},
					1, 10 * 4, 1, 5},
				{stageWidth - 2, stageHeight/2 - 1, 6, &ApproximatedAI{0, 5},
					1, 10 * 4, 2, 6},
				{1, stageHeight/2 - 1, 6, &ApproximatedAI{0, 5},
					1, 10 * 4, 2, 6}},
			[][]int32{{pSpeed, 11, 11, 9, 9},
				{pSpeed * 3 / 2, 7, 7, 5, 5},
				{pSpeed * 2, 5, 5, 5, 5}}),
		GetPreStageData(""+
			"#########################"+
			"###########070###########"+
			"#########000#000#########"+
			"#######000#####000#######"+
			"#####000#########000#####"+
			"###000######%######000###"+
			"#400#***************#004#"+
			"#0#0#*######%######*#0#0#"+
			"#0#0#***************#0#0#"+
			"#0#0#*#############*#0#0#"+
			"#504#***************#405#"+
			"#0######0###0###0######0#"+
			"#0######000000000######0#"+
			"#00000000#######00000000#"+
			"#########################",
			stageWidth/2, stageHeight/2-1,
			[]PreSnakeData{{1, stageHeight - 2, 3, &ApproximatedAI{0, 4},
				1, 10 * 4, 1, 3},
				{stageWidth - 2, stageHeight - 2, 3, &ApproximatedAI{0, 4},
					1, 10 * 4, 1, 3},
				{stageWidth / 2, 1, 4, &ApproximatedAI{0, 2},
					1, 10 * 4, 2, 5},
				{stageWidth - 2, stageHeight/2 - 1, 6, &ApproximatedAI{0, 5},
					1, 10 * 4, 2, 3},
				{1, stageHeight/2 - 1, 6, &ApproximatedAI{0, 5},
					1, 10 * 4, 2, 3}},
			[][]int32{{pSpeed, 11, 11, 9, 15, 15},
				{pSpeed * 3 / 2, 7, 7, 5, 11, 11},
				{pSpeed * 2, 5, 5, 4, 8, 8}}),
		GetPreStageData(""+
			"#########################"+
			"#00*********************#"+
			"#0#*#0#0#0#0#0#0#0#0#0#*#"+
			"#00***00000070007000000*#"+
			"#0#0#*#0#0#0#0#0#0#0#0#*#"+
			"#0700***0000000000000***#"+
			"#0#0#0#*#0#0#0#0#0#0#*#0#"+
			"#000000***000000000***00#"+
			"#0#0#0#0#*#0#0#0#0#*#0#0#"+
			"#00000000***00000***0000#"+
			"#0#0#0#0#0#*#0#0#*#0#0#0#"+
			"#0700000000***0***000070#"+
			"#0#7#7#0#0#0#*#*#0#7#7#0#"+
			"#000000000000***00000000#"+
			"#########################",
			stageWidth/2, stageHeight/2,
			[]PreSnakeData{{1, stageHeight - 2, 3, &ApproximatedAI{0, 4},
				1, 10 * 4, 1, 2},
				{stageWidth - 2, stageHeight - 2, 3, &ApproximatedAI{0, 4},
					1, 10 * 4, 1, 2},
				{stageWidth / 2, 1, 4, &ApproximatedAI{0, 2},
					1, 10 * 4, 2, 3},
				{stageWidth - 2, stageHeight/2 - 1, 6, &ApproximatedAI{0, 5},
					1, 10 * 4, 2, 2},
				{1, stageHeight/2 - 1, 6, &ApproximatedAI{0, 5},
					1, 10 * 4, 2, 2}},
			[][]int32{{pSpeed, 11, 11, 9, 15, 15},
				{pSpeed * 3 / 2, 7, 7, 5, 11, 11},
				{pSpeed * 2, 5, 5, 4, 8, 8}}),
		nil,
		GetPreStageData(""+
			"#########################"+
			"#########################"+
			"########*********########"+
			"######***#######***######"+
			"#####**##8#8#8#8##**#####"+
			"#####*####8###8####*#####"+
			"#####*###8#8#8#8###*#####"+
			"#####*#############*#####"+
			"#####**###########**#####"+
			"######**###888###**######"+
			"#######**#8###8#**#######"+
			"########*#######*########"+
			"########**#####**########"+
			"#########*******#########"+
			"#########################",
			stageWidth/2, stageHeight-2,
			nil, [][]int32{{pSpeed}, {pSpeed * 3 / 2}, {pSpeed * 2}}),
	}
	fmt.Println(len(data))
	firstNonIntro := 0
	for i := 0; data[i] != nil; i++ {
		firstNonIntro++
	}

	single := 0
	double := 0
	triple := 0

	parts := 1
	for i := firstNonIntro + 1; i < len(data); i++ {
		if data[i] != nil {
			v := len(data[i].difficultyData)
			if v == 1 {
				single++
			} else if v == 2 {
				double++
			} else {
				triple++
			}
		} else {
			parts++
		}
	}

	levels := [5][][2]int{
		make([][2]int, firstNonIntro+single+double*2+triple*3),
		make([][2]int, firstNonIntro+single+double+triple*2),
		make([][2]int, firstNonIntro+single+double+triple),
		make([][2]int, firstNonIntro+single+double+triple*2),
		make([][2]int, firstNonIntro+single+double*2+triple*3),
	}

	for i := 0; i < firstNonIntro; i++ {
		levels[0][i] = [2]int{i, 0}
		levels[1][i] = [2]int{i, 1}
		levels[2][i] = [2]int{i, 2}
		levels[3][i] = [2]int{i, 0}
		levels[4][i] = [2]int{i, 0}
	}

	diff := make([]int, parts)
	c1 := firstNonIntro
	c2 := firstNonIntro
	c3 := firstNonIntro

	for c1 < len(levels[3]) || c2 < len(levels[4]) || c3 < len(levels[2]) {
		level := 0
		for i := firstNonIntro + 1; i < len(data); i++ {
			if data[i] != nil {
				v := len(data[i].difficultyData)
				if v == 1 {
					if diff[level] == 0 {
						levels[3][c1] = [2]int{i, 0}
						levels[4][c2] = [2]int{i, 0}
						levels[2][c3] = [2]int{i, 0}
						c1++
						c2++
						c3++
					}
				} else {
					mhDiff := diff[level] / 2
					if diff[level]%2 == 0 && mhDiff < v {
						levels[4][c2] = [2]int{i, mhDiff}
						c2++
						if diff[level]/2 == v-1 {
							levels[2][c3] = [2]int{i, mhDiff}
							c3++
						}
					}
					eDiff := diff[level] / 3
					if diff[level]%3 == 0 && eDiff < v-1 {
						levels[3][c1] = [2]int{i, eDiff}
						c1++
					}
				}
			} else {
				if diff[level] == 0 {
					diff[level]++
					level = -1
					i = firstNonIntro
				} else {
					diff[level]++
				}
				level++
			}
		}
		diff[level]++
	}

	c1, c2 = firstNonIntro, firstNonIntro
	for diff := 0; diff < 3; diff++ {
		for i := firstNonIntro; i < len(data); i++ {
			if data[i] != nil {
				l := len(data[i].difficultyData)
				if l == 1 && diff == 0 {
					levels[0][c1] = [2]int{i, 0}
					levels[1][c2] = [2]int{i, 0}
				} else if l == 2 {
					if diff == 0 {
						levels[0][c1] = [2]int{i, 0}
						c1++
					} else if diff == 3 {
						levels[0][c1] = [2]int{i, 1}
						levels[1][c2] = [2]int{i, 1}
						c1++
						c2++
					}
				} else {
					levels[0][c1] = [2]int{i, 0}
					c1++
					if diff != 0 {
						levels[1][c2] = [2]int{i, diff}
						c2++
					}
				}
			}
		}
	}

	return data, levels
}

func (stage *Stage) Load(ID int, loadTiles bool, score int64,
	players int) *Engine {
	if ID == 0 {
		stage.lostOnce = false
	}
	stage.ID = ID
	hideSnakes, hideWalls := false, false
	if ID >= len(stage.levels[difficulty]) {
		if stage.lostOnce {
			return nil
		}
		ID = len(stage.levels[difficulty])*2 - ID - 1
		hideSnakes = true
		if ID < len(stage.levels[difficulty])*2/3 {
			hideWalls = true
		}
	}
	if ID < 0 {
		return nil
	}

	fmt.Printf("Loading level %d, Tiles: %t\n", ID, loadTiles)

	levelIndex := stage.levels[difficulty][ID][0]
	diffIndex := stage.levels[difficulty][ID][1]

	return stage.LoadSingleLevel(levelIndex, diffIndex,
		hideSnakes, hideWalls, loadTiles, score, players)
}

func (stage *Stage) LoadSingleLevel(levelIndex, diffIndex int, hideSnakes,
	hideWalls, loadTiles bool, score int64, players int) *Engine {
	var p1, p2 *Player
	var snakes []*Snake
	stage.sprites.entities = make([][]*Entity, 4)
	for i := 0; i < 4; i++ {
		stage.sprites.entities[i] = make([]*Entity, 0)
	}

	level := stage.stages[levelIndex]
	diffData := level.difficultyData[diffIndex]

	if loadTiles {
		ConvertStringToTiles(level.stage)
	}
	p1 = &Player{stage.sprites.GetEntity(level.px, level.py, Player1),
		diffData.playerSpeed}
	if players != 0 {
		p2 = &Player{stage.sprites.GetEntity(level.px, level.py, Player2),
			diffData.playerSpeed}
	}
	if diffData.snakes != nil {
		snakes = make([]*Snake, len(diffData.snakes))
		for i := 0; i < len(snakes); i++ {
			snake := diffData.snakes[i]
			ai := snake.ai
			if hideSnakes {
				ai = &HiddenAI{0, 200 / (snake.moveTimerMax + 5), false, ai}
			}
			snakes[i] = stage.sprites.GetSnake(snake.x, snake.y,
				snake.length, ai, snake.moveTimerMax,
				snake.growTimerMax, snake.minLength, snake.maxLength)
			snakes[i].ai.Reset()
			stage.sprites.AlertSnakes(snakes[i], false)
		}
	}
	fmt.Println("Exited set-up of stage")

	if loadTiles {
		fmt.Println("Calculating points left")
		stage.tiles.renderedOnce = false
		stage.hideWalls = hideWalls
		stage.pointsLeft = 0
		for x := int32(0); x < stageWidth; x++ {
			for y := int32(0); y < stageHeight; y++ {
				if tiles[x][y] == Point {
					stage.pointsLeft++
				}
			}
		}
		stage.tiles.tiles = tiles
		fmt.Printf("Replacing tiles\tPoints: %d\n", stage.pointsLeft)
	}

	fmt.Println("Getting engine")

	engine := GetEngine(p1, p2, score, snakes, stage)
	fmt.Println("Finished loading level ", stage.ID)
	fmt.Printf("Stage: %d\tDifficulty: %d\n", levelIndex, diffIndex)
	return engine
}

//Uses a global variable so it won't have to generate a ton of arrays
//TODO Consider naming [][]Tile and pass it as a variable
func ConvertStringToTiles(s string) {
	if tiles == nil {
		tiles = make([][]Tile, stageWidth)
		for x := int32(0); x < stageWidth; x++ {
			tiles[x] = make([]Tile, stageHeight)
		}
	}
	for x := int32(0); x < stageWidth; x++ {
		for y := int32(0); y < stageHeight; y++ {
			if x == 0 || y == 0 || x == stageWidth-1 ||
				y == stageHeight-1 {
				tiles[x][y] = Wall
			} else {
				tiles[x][y] = Point
			}
		}
	}

	for y := int32(0); y < stageHeight; y++ {
		for x := int32(0); x < stageWidth; x++ {
			c := s[y*stageWidth+x]
			if c == '#' {
				tiles[x][y] = Wall
			} else if c == '*' {
				tiles[x][y] = Point
			} else if c == '%' {
				tiles[x][y] = SnakeWall
			} else if c >= '0' && c <= '9' {
				tiles[x][y] = Tile(c - '0')
			}
		}
	}
}
