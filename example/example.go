package example

import (
	"math"

	"github.com/wangzun/gogame/engine/geometry"
	"github.com/wangzun/gogame/engine/graphic"
	"github.com/wangzun/gogame/engine/light"
	"github.com/wangzun/gogame/engine/material"
	"github.com/wangzun/gogame/engine/math32"
)

//
// PointLightMesh
//
type PointLightMesh struct {
	*graphic.Mesh
	Light *light.Point
}

func NewPointLightMesh(color *math32.Color) *PointLightMesh {

	l := new(PointLightMesh)

	geom := geometry.NewSphere(0.05, 32, 32, 0, math.Pi*2, 0, math.Pi)
	mat := material.NewStandard(color)
	mat.SetUseLights(0)
	mat.SetEmissiveColor(color)
	l.Mesh = graphic.NewMesh(geom, mat)
	l.Mesh.SetVisible(true)

	l.Light = light.NewPoint(color, 1.0)
	l.Light.SetPosition(0, 0, 0)
	l.Light.SetLinearDecay(1)
	l.Light.SetQuadraticDecay(1)
	l.Light.SetVisible(true)

	l.Mesh.Add(l.Light)

	return l
}

//
// SpotLightMesh
//
type SpotLightMesh struct {
	*graphic.Mesh
	Light *light.Spot
}

func NewSpotLightMesh(color *math32.Color) *SpotLightMesh {

	l := new(SpotLightMesh)

	geom := geometry.NewCylinder(0, 0.3, 0.5, 16, 16, 0, 2*math32.Pi, true, true)
	mat1 := material.NewStandard(color)
	mat2 := material.NewStandard(color)
	mat2.SetEmissiveColor(color)
	l.Mesh = graphic.NewMesh(geom, nil)
	l.Mesh.AddGroupMaterial(mat1, 0)
	l.Mesh.AddGroupMaterial(mat2, 1)

	l.Light = light.NewSpot(color, 2.0)
	l.Light.SetDirection(0, -1, 0)
	l.Light.SetCutoffAngle(45)
	l.Light.SetLinearDecay(0)
	l.Light.SetQuadraticDecay(0)

	l.Mesh.Add(l.Light)

	return l
}
