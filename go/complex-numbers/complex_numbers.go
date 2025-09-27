package complexnumbers

import "math"

// Define the Number type here.
type Number struct {
	Re float64
	Im float64
}

func (n Number) Real() float64 {
	return n.Re
}

func (n Number) Imaginary() float64 {
	return n.Im
}

func (n1 Number) Add(n2 Number) Number {
	return Number{Re: n1.Re + n2.Re, Im: n1.Im + n2.Im}
}

func (n1 Number) Subtract(n2 Number) Number {
	return Number{Re: n1.Re - n2.Re, Im: n1.Im - n2.Im}
}

func (n1 Number) Multiply(n2 Number) Number {
	return Number{
		Re: n1.Re*n2.Re - n1.Im*n2.Im,
		Im: n1.Re*n2.Im + n1.Im*n2.Re,
	}
}

func (n Number) Times(factor float64) Number {
	return Number{Re: n.Re * factor, Im: n.Im * factor}
}

func (n1 Number) Divide(n2 Number) Number {
	den := n2.Re*n2.Re + n2.Im*n2.Im
	return Number{
		Re: (n1.Re*n2.Re + n1.Im*n2.Im) / den,
		Im: (n1.Im*n2.Re - n1.Re*n2.Im) / den,
	}
}

func (n Number) Conjugate() Number {
	return Number{Re: n.Re, Im: -n.Im}
}

func (n Number) Abs() float64 {
	return math.Sqrt(n.Re*n.Re + n.Im*n.Im)
}

func (n Number) Exp() Number {
	expRe := math.Exp(n.Re)
	return Number{
		Re: expRe * math.Cos(n.Im),
		Im: expRe * math.Sin(n.Im),
	}
}
