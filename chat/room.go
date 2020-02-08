package main

type room struct{
	//forwordは他クライアントに転送するためのメッセージを保持するチャネルです。
	forward chan []byte
}