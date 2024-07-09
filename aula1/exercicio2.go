package main
import ("errors" "fmt" )

func main ()
{
	str := "Hello, World!"
	lengh, err := CountString (str)

	if err != nil
	{
		fmt.Println ("Erro:", err)
	} else
	{
		fmt.Println ("Tamanho da String: ", lengh)
	}
}

func CountString (s string) (int, error)
{
	if s == " "
	{
		return 0, errors.New ("String está vazia")
	}
	return len (s), nil
}