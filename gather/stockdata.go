package gather

import "strconv"

// Price is a type used to represent monetary values as integers.
// This simplifies and speeds up code in many instances.
type Price uint32

func (p Price) String() string {
	s := strconv.FormatUint(uint64(p), 10)
	l := len(s)
	return s[0:l-2] + "." + s[l-2:]
}

// ToFloat64 converts a price to a float64.
func (p Price) ToFloat64() float64 {
	return float64(p) / 100.0
}

type stock struct {
	name              string
	symbol            string
	close             Price
	change            Price
	volume            float32
	high              Price
	low               Price
	high52            Price
	low52             Price
	marketcap         float32
	sharesoutstanding float32
}

var ExampleStock stock = stock{}

// Reference regarding using float32 values:
// https://neurostars.org/t/consequence-of-using-single-float32-or-double-float64-precision-for-saving-interpolated-data/224/3
//
// Question:
// When saving interpolated data (after linear or non-linear warps but also as well as internal representation) we often face the decision if single (float32) or double (float64) precision should be used. I was wondering if there are any comparisons using MRI (fMRI especially) data between the two approaches that would show how much precision is being lost when choosing float32 and if that precision difference has real world consequences.
//
// The reason I’m asking is that using float64 comes with performance and memory (RAM and HDD) penalties which are especially painful when dealing with high-resolution short TR data.
//
// Update:
// On the other hand for data coming from a multiband sequence (say x8) acquired using 16-bit ADCs, the last step of the reconstruction involves combining 32 int16s, so there is potentially quite a bit more than 16-bits of information there which could justify using float64.
//
// Best Answer:
// For your original question, interpolation does not increase the available precision, so you only need to worry about rounding error. If your input is float32 or worse, float32 is fine for the interpolation output.
//
// Regarding your update about combination of values in reconstruction: the ideal conditions for a combination of independent measurements is repeatedly measuring the same quantity, in which case you can gain log2(N) bits of precision - for 32 independent measurements of the same quantity at 16 bits each, you can gain 5 bits, leading to a total of 21 bits. float32 has 24 bits of precision, so it has room to spare.
//
// However, if they are independent measurements but there are some weighting factors involved, it will not approach this ideal. If they are not of the same quantity, such that there is cancellation involved, you can instead have less precision than the inputs. This also ignores the more practical issue of SNR.
//
// As a more intuitive guideline for this, float32 has a precision ratio (take a value and the closest next representable number, divide the difference between them by their average) in the range [2^-24, 2^-23] = [10^-7.22, 10^-6.92]. For comparison, 7 significant figures in scientific notation has a precision ratio in the range [10^-7, 10^-6] (because 1.000002 and 9.999998 both have the same absolute step to the next number, but differ by an order of magnitude).
//
// So you can just ask yourself: does it make sense to have more than 7 significant figures? If not, store it as float32.
