* Every file in go has to belong to a package.
    In this case main.go file blongs to package main which indicates compiler that this file is an exacutable program not a library.
    It knows the main() entry point is in this file and looks for it.
* Importing is almost like any other language: import "fmt". if multible libraries being imported import {   "fmt", "..."}
* fmt is the standart format package in go like stdio.h or iostream