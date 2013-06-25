package mangle

type Vector4d [4]float64

// Dot computes the dot product of two Vector4d's
func Dot(x, y *Vector4d) float64 {
	return x[0]*y[0] + x[1]*y[1] + x[2]*y[2] + x[3]*y[3]
}

// Caps are the building blocks of mangle masks. 
//
// The canonical description uses 3d vectors for the direction, 
// and then decides whether a point is in or out of a cap with the
// following logic :
//   cdot = 1 - x . v
//   if cm > 0 
//       return cdot < cm
//   else 
//       return cdot > cm
//
// We store the caps as a 4 vector
//     (1, -x, -y, -z) (cm > 0)
//     (-1, x, y, z) (cm < 0)
// and dot with (1, x1, y1, z1) and compare with cm 
type Cap struct {
	V  Vector4d
	CM float64
}

func InCap(c *Cap, v *Vector4d) bool {
	return Dot(&c.V, v) < c.CM
}
