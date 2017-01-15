package voltorbflip

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func NewSolvedRow() (vfPSolLineData, [5]VfPSolTile) {
	lineData := vfPSolLineData{0, 0, 0}
	tiles := [5]VfPSolTile{NewSolvedVfPSolTile(0), NewSolvedVfPSolTile(0), NewSolvedVfPSolTile(1), NewSolvedVfPSolTile(2), NewSolvedVfPSolTile(3)}
	return lineData, tiles
}

func TestHeuristic0(t *testing.T) {
	Convey("For Heuristic 0", t, func() {
		Convey("Given an row where RV+RP!=NUT", func() {
			lineData := vfPSolLineData{5, 1, 5}
			tiles := [5]VfPSolTile{
				NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile(),
				NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile()}

			So(heuristic0(lineData, &tiles), ShouldBeFalse)
			So(tiles, ShouldEqual,
				[5]VfPSolTile{
					{true, true, true, true},
					{true, true, true, true},
					{true, true, true, true},
					{true, true, true, true},
					{true, true, true, true},
				})
		})

		Convey("Given an unsolved row where RV+RP=NUT", func() {
			lineData := vfPSolLineData{3, 2, 5}
			tiles := [5]VfPSolTile{
				NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile(),
				NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile()}

			So(heuristic0(lineData, &tiles), ShouldBeTrue)
			So(tiles, ShouldEqual,
				[5]VfPSolTile{
					{true, true, false, false},
					{true, true, false, false},
					{true, true, false, false},
					{true, true, false, false},
					{true, true, false, false},
				})
		})

		Convey("Given a partially solved row where RV+RP=NUT", func() {
			lineData := vfPSolLineData{3, 1, 4}
			tiles := [5]VfPSolTile{
				NewSolvedVfPSolTile(3), NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile(),
				NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile()}

			So(heuristic0(lineData, &tiles), ShouldBeTrue)
			So(tiles, ShouldEqual,
				[5]VfPSolTile{
					{false, false, false, true},
					{true, true, false, false},
					{true, true, false, false},
					{true, true, false, false},
					{true, true, false, false},
				})
		})
	})
}

func TestHeuristic1(t *testing.T) {
	Convey("For Heuristic 1", t, func() {
		Convey("Given an row where NUT-RV!=RP-1", func() {
			lineData := vfPSolLineData{5, 0, 5}
			tiles := [5]VfPSolTile{
				NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile(),
				NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile()}

			So(heuristic1(lineData, &tiles), ShouldBeFalse)
			So(tiles, ShouldEqual,
				[5]VfPSolTile{
					{true, true, true, true},
					{true, true, true, true},
					{true, true, true, true},
					{true, true, true, true},
					{true, true, true, true},
				})
		})

		Convey("Given an unsolved row where NUT-RV=RP-1", func() {
			lineData := vfPSolLineData{6, 0, 5}
			tiles := [5]VfPSolTile{
				NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile(),
				NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile()}

			So(heuristic1(lineData, &tiles), ShouldBeTrue)
			So(tiles, ShouldEqual,
				[5]VfPSolTile{
					{true, true, true, false},
					{true, true, true, false},
					{true, true, true, false},
					{true, true, true, false},
					{true, true, true, false},
				})
		})

		Convey("Given an partially solved row where NUT-RV=RP-1", func() {
			lineData := vfPSolLineData{5, 0, 4}
			tiles := [5]VfPSolTile{
				{false, false, true, true},
				{false, false, false, true},
				NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile()}

			So(heuristic1(lineData, &tiles), ShouldBeTrue)
			So(tiles, ShouldEqual,
				[5]VfPSolTile{
					{false, false, true, false},
					{false, false, false, true},
					{true, true, true, false},
					{true, true, true, false},
					{true, true, true, false},
				})
		})
	})
}

func TestHeuristic2(t *testing.T) {
	Convey("For Heuristic 2", t, func() {
		Convey("Given an row where RV!=0", func() {
			lineData := vfPSolLineData{5, 1, 5}
			tiles := [5]VfPSolTile{
				NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile(),
				NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile()}

			So(heuristic2(lineData, &tiles), ShouldBeFalse)
			So(tiles, ShouldEqual,
				[5]VfPSolTile{
					{true, true, true, true},
					{true, true, true, true},
					{true, true, true, true},
					{true, true, true, true},
					{true, true, true, true},
				})
		})

		Convey("Given an unsolved row where RV==0", func() {
			lineData := vfPSolLineData{6, 0, 5}
			tiles := [5]VfPSolTile{
				NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile(),
				NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile()}

			So(heuristic2(lineData, &tiles), ShouldBeTrue)
			So(tiles, ShouldEqual,
				[5]VfPSolTile{
					{false, true, true, true},
					{false, true, true, true},
					{false, true, true, true},
					{false, true, true, true},
					{false, true, true, true},
				})
		})

		Convey("Given an partially solved row where RV==0", func() {
			lineData := vfPSolLineData{5, 0, 3}
			tiles := [5]VfPSolTile{
				NewSolvedVfPSolTile(0),
				NewSolvedVfPSolTile(3),
				NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile()}

			So(heuristic2(lineData, &tiles), ShouldBeTrue)
			So(tiles, ShouldEqual,
				[5]VfPSolTile{
					{true, false, false, false},
					{false, false, false, true},
					{false, true, true, true},
					{false, true, true, true},
					{false, true, true, true},
				})
		})
	})
}

func TestHeuristic3(t *testing.T) {
	Convey("For Heuristic 3", t, func() {
		Convey("Given an row where NUT-1!=RV", func() {
			lineData := vfPSolLineData{5, 1, 5}
			tiles := [5]VfPSolTile{
				NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile(),
				NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile()}

			So(heuristic3(lineData, &tiles), ShouldBeFalse)
			So(tiles, ShouldEqual,
				[5]VfPSolTile{
					{true, true, true, true},
					{true, true, true, true},
					{true, true, true, true},
					{true, true, true, true},
					{true, true, true, true},
				})
		})

		Convey("Given an unsolved row where NUT-1==RV", func() {
			lineData := vfPSolLineData{3, 4, 5}
			tiles := [5]VfPSolTile{
				NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile(),
				NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile()}

			So(heuristic3(lineData, &tiles), ShouldBeTrue)
			So(tiles, ShouldEqual,
				[5]VfPSolTile{
					{true, false, false, true},
					{true, false, false, true},
					{true, false, false, true},
					{true, false, false, true},
					{true, false, false, true},
				})
		})

		Convey("Given an partially solved row where NUT-1==RV", func() {
			lineData := vfPSolLineData{2, 3, 4}
			tiles := [5]VfPSolTile{
				NewSolvedVfPSolTile(3),
				{false, true, true, true},
				NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile()}

			So(heuristic3(lineData, &tiles), ShouldBeTrue)
			So(tiles, ShouldEqual,
				[5]VfPSolTile{
					{false, false, false, true},
					{false, false, true, false},
					{true, false, true, false},
					{true, false, true, false},
					{true, false, true, false},
				})
		})
	})
}

func TestHeuristic6(t *testing.T) {
	Convey("For Heuristic 6", t, func() {
		Convey("Given an row where RV!=NUT", func() {
			lineData := vfPSolLineData{5, 1, 5}
			tiles := [5]VfPSolTile{
				NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile(),
				NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile()}

			So(heuristic6(lineData, &tiles), ShouldBeFalse)
			So(tiles, ShouldEqual,
				[5]VfPSolTile{
					{true, true, true, true},
					{true, true, true, true},
					{true, true, true, true},
					{true, true, true, true},
					{true, true, true, true},
				})
		})

		Convey("Given an unsolved row where RV==NUT", func() {
			lineData := vfPSolLineData{0, 5, 5}
			tiles := [5]VfPSolTile{
				NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile(),
				NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile()}

			So(heuristic6(lineData, &tiles), ShouldBeTrue)
			So(tiles, ShouldEqual,
				[5]VfPSolTile{
					{true, false, false, false},
					{true, false, false, false},
					{true, false, false, false},
					{true, false, false, false},
					{true, false, false, false},
				})
		})

		Convey("Given an partially solved row where RV==NUT", func() {
			lineData := vfPSolLineData{0, 2, 2}
			tiles := [5]VfPSolTile{
				NewSolvedVfPSolTile(3),
				NewSolvedVfPSolTile(2),
				NewSolvedVfPSolTile(1),
				NewUnsolvedVfPSolTile(), NewUnsolvedVfPSolTile()}

			So(heuristic6(lineData, &tiles), ShouldBeTrue)
			So(tiles, ShouldEqual,
				[5]VfPSolTile{
					{false, false, false, true},
					{false, false, true, false},
					{false, true, false, false},
					{true, false, false, false},
					{true, false, false, false},
				})
		})
	})
}
