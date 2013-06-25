package mangle

import (
	"errors"
	"fmt"
	"io"
)

type vector4d [4]float64

// Dot computes the dot product of two vector4d's
func dot(x, y *vector4d) float64 {
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
	v  vector4d
	cm float64
}

// in tests to see if the point is in the cap or not.
func (c *Cap) In(v *vector4d) bool {
	return dot(&c.v, v) < c.cm
}

// read reads in cap coordinates from io.Reader
func (c *Cap) Read(r io.Reader) error {
	n, err := fmt.Fscanf(r, "%f %f %f %f \n", &c.v[1], &c.v[2], &c.v[3], &c.cm)
	if n != 4 {
		return errors.New("Unable to parse cap")
	}
	if err != nil {
		return err
	}
	if c.cm < 0 {
		c.v[0] = -1
	} else {
		c.v[0] = 1
		for i := 1; i < 4; i++ {
			c.v[i] = -c.v[i]
		}
	}
	return nil
}

// String allows access to the cap type
func (c *Cap) String() string {
	return fmt.Sprintf("%e %e %e %e \n", c.v[1], c.v[2], c.v[3], c.cm)
}

type CapList []Cap
