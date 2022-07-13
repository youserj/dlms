package common_data_type

import "fmt"

type LengthError struct{
	got int
	expect int
}

func (e *LengthError) Error() string{
	return fmt.Sprintf("got length %d bytes, expected %d", e.got, e.expect)
}

type TagError struct{
	got byte
	expect byte
}

func (e *TagError) Error() string{
	return fmt.Sprintf("got tag %d, expected %d", e.got, e.expect)
}

type UnsupLengthError struct{
	length int
}

func (e *UnsupLengthError) Error() string{
	return fmt.Sprintf("unsupport length %d", e.length)
}

type NotEnoughValueTagError struct{
	tag byte
}

func (e *NotEnoughValueTagError) Error() string{
	return fmt.Sprintf("not enough value with tag: %d", e.tag)
}
