package main

func checkNullError(err error) {
	if err != nil {
		panic(err)
	}
}
