package method

type MyString string

func (s MyString) AddHoge() MyString {

	return s + " Hoge"
}