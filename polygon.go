package mangle

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
)

// Store polygon meta data. 
// This only makes sense if there is a corresponding cap class.
type Polygon struct {
	id, pixel, ncaps int64
	weight           float64
	clist            CapList
}

// Parse reads in a poly string and fills in the key pieces
func (p *Polygon) Parse(b []byte) error {
	r := bytes.NewReader(b)

	// Attempt to read the polygon id
	_, err := fmt.Fscanf(r, "polygon %d", &p.id)
	if err != nil {
		return err
	}

	// Grab the region in parens
	n1 := bytes.Index(b, []byte("("))
	if n1 == -1 {
		return errors.New("Error parsing polygon line : no opening parens")
	}
	n2 := bytes.Index(b, []byte(")"))
	if n2 == -1 {
		return errors.New("Error parsing polygon line : no closing parens")
	}
	if n2 <= (n1 + 1) {
		return errors.New("Opening and closing parens next to one another")
	}

	// Grab the substring
	b1 := b[n1+1 : n2]
	var op, arg string
	bsub := bytes.Split(b1, []byte(","))
	for i := range bsub {
		r := bytes.NewReader(bsub[i])
		_, err = fmt.Fscanf(r, "%s %s", &arg, &op)
		if err != nil {
			return err
		}
		switch op {
		case "pixel":
			p.pixel, err = strconv.ParseInt(arg, 0, 64)
			if err != nil {
				return err
			}
		case "caps":
			p.ncaps, err = strconv.ParseInt(arg, 0, 64)
			if err != nil {
				return err
			}
			p.clist = make(CapList, p.ncaps)
		case "weight":
			p.weight, err = strconv.ParseFloat(arg, 64)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// In checks to see if a point is in a polygon or not
func (p *Polygon) In(v *vector4d) bool {
	var tmp bool
	for i := range p.clist {
		tmp = p.clist[i].In(v)
		if !tmp {
			return false
		}
	}
	return true
}
